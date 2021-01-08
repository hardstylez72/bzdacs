package route

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/hardstylez72/bblog/ad/pkg/routetag"
	"github.com/hardstylez72/bblog/ad/pkg/tag"
	"github.com/hardstylez72/bblog/ad/pkg/util"
	"github.com/jmoiron/sqlx"
	"strings"
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

func (r *repository) UpdateWithTags(ctx context.Context, route *Route, tagNames []string) (*RouteWithTags, error) {
	tx, err := r.conn.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	updatedRoute, err := UpdateTx(ctx, tx, route)
	if err != nil {
		return nil, err
	}

	tagNames, err = routetag.Merge(ctx, r.conn, tx, updatedRoute.Id, tagNames)
	if err != nil {
		return nil, err
	}
	return &RouteWithTags{
		Route: *updatedRoute,
		Tags:  tagNames,
	}, nil
}
func (r *repository) Update(ctx context.Context, route *Route) (*Route, error) {
	query := `
	
			update ad.routes
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
						   deleted_at
`

	rows, err := r.conn.NamedQueryContext(ctx, query, route)
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

func UpdateTx(ctx context.Context, tx *sqlx.Tx, route *Route) (*Route, error) {
	query := `
	
			update ad.routes
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
						   deleted_at
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

func (r *repository) Insert(ctx context.Context, route *Route) (*Route, error) {
	query := `
insert into ad.routes (
                       route,
                       method,
                       description,
                       created_at,
                       updated_at,
                       deleted_at
                       )
                   values (
                       :route,
                       :method,
                       :description,
                       now(),
                       null,
                       null
                   ) returning id,
                               route,
                       		   method,
                               description,
                               created_at,
                               updated_at,
                               deleted_at;
`

	rows, err := r.conn.NamedQueryContext(ctx, query, route)
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

func (r *repository) InsertWithTags(ctx context.Context, route *Route, tagNames []string) (*RouteWithTags, error) {
	tx, err := r.conn.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	newRoute, err := InsertTx(ctx, tx, route)
	if err != nil {
		return nil, err
	}

	tagNames, err = routetag.Merge(ctx, r.conn, tx, newRoute.Id, tagNames)
	if err != nil {
		return nil, err
	}
	return &RouteWithTags{
		Route: *newRoute,
		Tags:  tagNames,
	}, nil
}

func InsertTx(ctx context.Context, tx *sqlx.Tx, route *Route) (*Route, error) {
	query := `
insert into ad.routes (
                       route,
                       method,
                       description,
                       created_at,
                       updated_at,
                       deleted_at
                       )
                   values (
                       :route,
                       :method,
                       :description,
                       now(),
                       null,
                       null
                   ) returning id,
                               route,
                       		   method,
                               description,
                               created_at,
                               updated_at,
                               deleted_at;
`
	var g Route
	rows, err := tx.NamedQuery(query, route)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.StructScan(&g)
		if err != nil {
			return nil, err
		}
	}
	if g.Id == 0 {
		return nil, ErrEntityAlreadyExists
	}

	return &g, nil
}
func (r *repository) GetByMethodAndRoute(ctx context.Context, route, method string) (*RouteWithTags, error) {
	rr, err := GetByMethodAndRouteDb(ctx, r.conn, method, route)
	if err != nil {
		return nil, err
	}

	tagNames, err := getTagsByRouteId(ctx, r.conn, rr.Id)
	if err != nil {
		return nil, err
	}

	return &RouteWithTags{
		Route: *rr,
		Tags:  tagNames,
	}, nil
}

func (r *repository) GetById(ctx context.Context, id int) (*RouteWithTags, error) {
	route, err := GetByIdDb(ctx, r.conn, id)
	if err != nil {
		return nil, err
	}

	tagNames, err := getTagsByRouteId(ctx, r.conn, route.Id)
	if err != nil {
		return nil, err
	}

	return &RouteWithTags{
		Route: *route,
		Tags:  tagNames,
	}, nil
}

func getTagsByRouteId(ctx context.Context, conn *sqlx.DB, id int) ([]string, error) {
	tagIds, err := routetag.GetRouteTags(ctx, conn, id)
	if err != nil {
		return nil, err
	}
	tagNames := make([]string, 0)

	for _, tagId := range tagIds {
		t, err := tag.GetByIdDb(ctx, conn, tagId)
		if err != nil {
			return nil, err
		}
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

		tagNames, err := getTagsByRouteId(ctx, r.conn, route.Id)
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
		from ad.routes
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
		from ad.routes
	   where deleted_at is null
  		 and id = $1
`
	var route Route
	err := conn.GetContext(ctx, &route, query, id)
	if err != nil {
		return nil, err
	}

	return &route, nil
}

func formArrayStringSql(arr []string) string {
	buf := make([]string, 0)

	for i := range arr {
		buf = append(buf, `'`+arr[i]+`'`)
	}

	return strings.Join(buf, ", ")
}

func printQuestions(n int) string {
	out := ""
	for i := 0; i < n; i++ {
		out = out + "?,"
	}

	out = out[:len(out)-1]
	return out
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
		From("ad.routes r").
		InnerJoin("ad.routes_tags rt on rt.route_id = r.id").
		InnerJoin("ad.tags t on rt.tag_id = t.id").
		Where("r.deleted_at is null").
		Where("t.deleted_at is null")

	if len(f.Tags.Names) > 0 {
		if f.Tags.Exclude {
			routes1 = routes1.Where(`rt.route_id not in (
													select route_id from ad.routes_tags where tag_id in (
												   	   select id from ad.tags where name in (` + formArrayStringSql(f.Tags.Names) + `)
													)
										     )`)
			routes1 = routes1.Where(`t.name not in (` + formArrayStringSql(f.Tags.Names) + `)`)
		} else {
			routes1 = routes1.Where(`t.name in (` + formArrayStringSql(f.Tags.Names) + `)`)
		}
	}

	routes1 = routes1.GroupBy("r.id")

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

func (r *repository) Delete(ctx context.Context, id int) error {
	query := `
		update ad.routes 
			set deleted_at = now()
		where id = $1;
`
	_, err := r.conn.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
