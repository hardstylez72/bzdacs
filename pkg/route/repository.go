package route

import (
	"context"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"

	"github.com/hardstylez72/bzdacs/pkg/relations/routetag"
	"github.com/jmoiron/sqlx"
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

	updatedRoute, err := Update(ctx, txx, route)
	if err != nil {
		return nil, err
	}

	tags, err := routetag.Merge(ctx, txx, updatedRoute.Id, route.Tags, route.NamespaceId)
	if err != nil {
		return nil, err
	}

	updatedRoute.Tags = getTagIdsFromArray(tags)

	return updatedRoute, nil
}

func Update(ctx context.Context, driver storage.SqlDriver, route *Route) (*Route, error) {
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

	rows, err := driver.NamedQueryContext(ctx, query, route)
	if err != nil {
		return nil, err
	}

	var r Route
	for rows.Next() {
		err = rows.StructScan(&r)
		if err != nil {
			return nil, err
		}
	}

	return &r, nil
}

func (r *repository) GetByParams(ctx context.Context, route, method string, namespaceId int) (*Route, error) {
	rr, err := GetByParams(ctx, r.conn, method, route, namespaceId)
	if err != nil {
		return nil, err
	}

	tagNames, err := GetTagNamesByRouteId(ctx, r.conn, rr.Id, namespaceId)
	if err != nil {
		return nil, err
	}

	rr.Tags = tagNames
	return rr, nil
}

func (r *repository) GetById(ctx context.Context, id, namespaceId int) (*Route, error) {
	route, err := GetById(ctx, r.conn, id)
	if err != nil {
		return nil, err
	}

	tagNames, err := GetTagNamesByRouteId(ctx, r.conn, route.Id, namespaceId)
	if err != nil {
		return nil, err
	}
	route.Tags = tagNames
	return route, nil
}

func GetTagNamesByRouteId(ctx context.Context, conn *sqlx.DB, routeId, namespaceId int) ([]string, error) {
	tags, err := routetag.GetRouteTags(ctx, conn, routeId, namespaceId)
	if err != nil {
		return nil, err
	}

	tagNames := make([]string, 0, len(tags))

	for _, t := range tags {
		tagNames = append(tagNames, t.Name)
	}
	return tagNames, nil
}

func GetByParams(ctx context.Context, driver storage.SqlDriver, method, route string, namespaceId int) (*Route, error) {
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
		 and namespace_id = $3
`
	var r Route
	err := driver.GetContext(ctx, &r, query, method, route, namespaceId)
	if err != nil {
		return nil, storage.PgError(err)
	}

	return &r, nil
}

func GetById(ctx context.Context, driver storage.SqlDriver, id int) (*Route, error) {
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
	err := driver.GetContext(ctx, &route, query, id)
	if err != nil {
		return nil, storage.PgError(err)
	}

	return &route, nil
}

func (r *repository) Delete(ctx context.Context, routeId, namespaceId int) error {
	return Delete(ctx, r.conn, routeId, namespaceId)
}
func Delete(ctx context.Context, driver storage.SqlDriver, routeId, namespaceId int) error {
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
