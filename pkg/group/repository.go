package group

import (
	"context"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/relations/grouproute"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

func NewRepository(conn *sqlx.DB) *repository {
	return &repository{conn: conn}
}

func (r *repository) InsertGroupWithCopy(ctx context.Context, g *Group, groupBaseId int) (*Group, error) {

	tx, err := r.conn.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	txx := storage.WrapSqlxTx(tx)

	group, err := InsertLL(ctx, txx, g)
	if err != nil {
		return nil, err
	}

	routes, _, err := grouproute.ListLL(ctx, txx, grouproute.Filter{
		Page:          1,
		PageSize:      0,
		RoutePattern:  "",
		BelongToGroup: true,
		NamespaceId:   group.NamespaceId,
		GroupId:       groupBaseId,
	})
	if err != nil {
		return nil, err
	}

	groupRoutePairs := make([]grouproute.Pair, 0)
	for i := range routes {
		groupRoutePairs = append(groupRoutePairs, grouproute.Pair{
			GroupId: group.Id,
			RouteId: routes[i].Id,
		})
	}

	_, err = grouproute.InsertLL(ctx, txx, groupRoutePairs)
	if err != nil {
		return nil, err
	}

	return group, nil
}
func (r *repository) Update(ctx context.Context, group *Group) (*Group, error) {
	return UpdateLL(ctx, r.conn, group)
}
func UpdateLL(ctx context.Context, driver storage.SqlDriver, group *Group) (*Group, error) {
	query := `
	update groups 
	   set code = :code,
		   description = :description,
		   updated_at = now()
	 where id = :id returning  id,
                               code,
                               description,
                               created_at,
                               updated_at,
                               deleted_at,
	     					   namespace_id
`

	rows, err := driver.NamedQueryContext(ctx, query, group)
	if err != nil {
		return nil, storage.PgError(err)
	}

	var g Group
	for rows.Next() {
		err = rows.StructScan(&g)
		if err != nil {
			return nil, storage.PgError(err)
		}
	}

	return &g, nil
}
func (r *repository) Insert(ctx context.Context, group *Group) (*Group, error) {
	return InsertLL(ctx, r.conn, group)
}
func InsertLL(ctx context.Context, driver storage.SqlDriver, group *Group) (*Group, error) {
	query := `
insert into groups (
                       code,
                       description,
                       created_at,
                       updated_at,
                       deleted_at,
                       namespace_id
                       )
                   values (
                       :code,
                       :description,
                       now(),
                       now(),
                       null,
                       :namespace_id
                   ) returning id,
                               code,
                               description,
                               created_at,
                               updated_at,
                               deleted_at,
                       		   namespace_id
`

	rows, err := driver.NamedQueryContext(ctx, query, group)
	if err != nil {
		return nil, err
	}

	var g Group
	for rows.Next() {
		err = rows.StructScan(&g)
		if err != nil {
			return nil, storage.PgError(err)
		}
	}

	return &g, nil
}

func (r *repository) GetByCode(ctx context.Context, code string, namespaceId int) (*Group, error) {
	return GetByCodeLL(ctx, r.conn, code, namespaceId)
}

func GetByCodeLL(ctx context.Context, driver storage.SqlDriver, code string, namespaceId int) (*Group, error) {
	query := `
		select id,
			   code,
			   description,
			   created_at,
			   updated_at,
			   deleted_at
		from groups
	   where deleted_at is null
	     and code = $1
		 and namespace_id = $2
`
	var group Group
	err := driver.GetContext(ctx, &group, query, code, namespaceId)
	if err != nil {
		return nil, storage.PgError(err)
	}

	return &group, nil
}

func (r *repository) GetById(ctx context.Context, id, namespaceId int) (*Group, error) {
	return GetByIdLL(ctx, r.conn, id, namespaceId)
}

func GetByIdLL(ctx context.Context, driver storage.SqlDriver, id, namespaceId int) (*Group, error) {
	query := `
		select id,
			   code,
			   description,
			   created_at,
			   updated_at,
			   deleted_at,
			   namespace_id
		from groups
	   where id = $1
		 and namespace_id = $2
`
	var group Group
	err := driver.GetContext(ctx, &group, query, id, namespaceId)
	if err != nil {
		return nil, storage.PgError(err)
	}

	return &group, nil
}

func (r *repository) Delete(ctx context.Context, id, namespaceId int) error {
	return DeleteLL(ctx, r.conn, id, namespaceId)
}
func DeleteLL(ctx context.Context, driver storage.SqlDriver, id, namespaceId int) error {
	query := `
		update groups 
			set deleted_at = now()
		where id = $1
		  and namespace_id = $2
`
	_, err := driver.ExecContext(ctx, query, id, namespaceId)
	if err != nil {
		return storage.PgError(err)
	}

	return nil
}
