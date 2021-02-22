package route

import (
	"context"
	"database/sql"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"

	sq "github.com/Masterminds/squirrel"
	"github.com/hardstylez72/bzdacs/pkg/routetag"
	"github.com/hardstylez72/bzdacs/pkg/util"
	"github.com/jmoiron/sqlx"
)

var (
	ErrEntityAlreadyExists = util.ErrEntityAlreadyExists
	ErrEntityNotFound      = util.ErrEntityNotFound
)

type repository struct {
	conn *sqlx.DB
}

func NewRepository(conn *sqlx.DB) *repository {
	return &repository{conn: conn}
}

func (r *repository) Update(ctx context.Context, route *Route) (*Route, error) {
	tx, err := r.conn.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	txx := storage.WrapSqlxTx(tx)

	updatedRoute, err := UpdateTx(ctx, tx, route)
	if err != nil {
		return nil, err
	}

	tags, err := routetag.MergeLL(ctx, txx, updatedRoute.Id, route.Tags, route.NamespaceId)
	if err != nil {
		return nil, err
	}

	updatedRoute.Tags = getTagIdsFromArray(tags)

	return updatedRoute, nil
}

func UpdateTx(ctx context.Context, tx *sqlx.Tx, route *Route) (*Route, error) {
	query := `
	
			update routes
			   set route = :route,
				   method = :method,
				   description = :description,
				   updated_at = now()
			where id = :id returning id,
						   route,
						   method,
						   description,
						   created_at,
						   updated_at,
						   deleted_at,
			    		   namespace_id
`

	rows, err := tx.NamedQuery(query, route)
	if err != nil {
		return nil, err
	}

	var g Route
	for rows.Next() {
		err = rows.StructScan(&g)
		if err != nil {
			return nil, err
		}
	}

	return &g, nil
}

func (r *repository) Get(ctx context.Context, route, method string) (*RouteWithTags, error) {
	rr, err := GetByMethodAndRouteDb(ctx, r.conn, method, route)
	if err != nil {
		return nil, err
	}

	tagNames, err := getTagNamesByRouteId(ctx, r.conn, rr.Id, 11)
	if err != nil {
		return nil, err
	}

	return &RouteWithTags{
		Route: *rr,
		Tags:  tagNames,
	}, nil
}

func (r *repository) GetById(ctx context.Context, id, namespaceId int) (*Route, error) {
	route, err := GetByIdDb(ctx, r.conn, id)
	if err != nil {
		return nil, err
	}

	tagNames, err := getTagNamesByRouteId(ctx, r.conn, route.Id, namespaceId)
	if err != nil {
		return nil, err
	}
	route.Tags = tagNames
	return route, nil
}

func getTagNamesByRouteId(ctx context.Context, conn *sqlx.DB, routeId, namespaceId int) ([]string, error) {
	tags, err := routetag.GetRouteTagsLL(ctx, conn, routeId, namespaceId)
	if err != nil {
		return nil, err
	}

	tagNames := make([]string, 0, len(tags))

	for _, t := range tags {
		tagNames = append(tagNames, t.Name)
	}
	return tagNames, nil
}
func (r *repository) List(ctx context.Context, f filter) ([]RouteWithTags, error) {

	routes, err := ListDb(ctx, r.conn, f)
	if err != nil {
		return nil, err
	}

	routesWithTags := make([]RouteWithTags, 0)

	for _, route := range routes {

		tagNames, err := getTagNamesByRouteId(ctx, r.conn, route.Id, 11)
		if err != nil {
			return nil, err
		}
		routesWithTags = append(routesWithTags, RouteWithTags{
			Route: route,
			Tags:  tagNames,
		})
	}
	return routesWithTags, nil
}

func GetByMethodAndRouteDb(ctx context.Context, conn *sqlx.DB, method, route string) (*Route, error) {
	query := `
		select id,
			   route,
		       method,
			   description,
			   created_at,
			   updated_at,
			   deleted_at
		from routes
	   where deleted_at is null
  		 and method = $1
 		 and route = $2
`
	var r Route
	err := conn.GetContext(ctx, &r, query, method, route)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrEntityNotFound
		}
		return nil, err
	}

	return &r, nil
}

func GetByIdDb(ctx context.Context, conn *sqlx.DB, id int) (*Route, error) {
	query := `
		select id,
			   route,
		       method,
			   description,
			   created_at,
			   updated_at,
			   deleted_at
		from routes
	   where id = $1
`
	var route Route
	err := conn.GetContext(ctx, &route, query, id)
	if err != nil {
		return nil, err
	}

	return &route, nil
}

func ListDb(ctx context.Context, conn *sqlx.DB, f filter) ([]Route, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	routes1 := psql.Select(`
			   r.id,
			   r.route,
		       r.method,
			   r.description,
			   r.created_at,
			   r.updated_at,
			   r.deleted_at
			`).
		From("routes r")

	if len(f.Tags.Names) > 0 {
		if f.Tags.Exclude {

			subQuery, args, err := sqlx.In(`
			r.id not in (    
           select rt.route_id from routes_tags rt 
             join tags t on t.id = rt.tag_id 
			 			   and t.deleted_at is null
 				           and t.name in (`+sq.Placeholders(len(f.Tags.Names))+`)
         )`, f.Tags.Names)
			if err != nil {
				return nil, err
			}
			routes1 = routes1.
				Where(subQuery, args...)
		} else {
			subQuery, args, err := sqlx.In(`
			r.id  in (    
           select rt.route_id from routes_tags rt 
             join tags t on t.id = rt.tag_id
						   and t.deleted_at is null
             			   and t.name in (`+sq.Placeholders(len(f.Tags.Names))+`)
         )`, f.Tags.Names)
			if err != nil {
				return nil, err
			}
			routes1 = routes1.
				Where(subQuery, args...)
		}
	}

	routes1 = routes1.
		Where("r.deleted_at is null")

	query, args, err := routes1.ToSql()
	if err != nil {
		return nil, err
	}

	groups := make([]Route, 0)
	err = conn.SelectContext(ctx, &groups, query, args...)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (r *repository) Delete(ctx context.Context, routeId, namespaceId int) error {
	return DeleteLL(ctx, r.conn, routeId, namespaceId)
}
func DeleteLL(ctx context.Context, driver storage.SqlDriver, routeId, namespaceId int) error {
	query := `
		update routes 
			set deleted_at = now()
		where id = $1
		  and namespace_id = $2
`
	_, err := driver.ExecContext(ctx, query, routeId, namespaceId)
	if err != nil {
		return err
	}

	return nil
}
