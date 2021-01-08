package group

import (
	"context"
	"database/sql"
	"github.com/hardstylez72/bblog/ad/pkg/grouproute"
	"github.com/hardstylez72/bblog/ad/pkg/util"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

var (
	ErrGroupAlreadyExists = util.ErrEntityAlreadyExists
	ErrEntityNotFound     = util.ErrEntityNotFound
)

func NewRepository(conn *sqlx.DB) *repository {
	return &repository{conn: conn}
}

func (r *repository) InsertGroupBasedOnAnother(ctx context.Context, g *Group, groupBaseId int) (*Group, error) {

	tx, err := r.conn.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	group, err := InsertTx(ctx, tx, g)
	if err != nil {
		return nil, err
	}

	routes, err := grouproute.ListDb(ctx, r.conn, groupBaseId)
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

	_, err = grouproute.InsertTx(ctx, tx, groupRoutePairs)
	if err != nil {
		return nil, err
	}

	return group, nil
}
func InsertTx(ctx context.Context, tx *sqlx.Tx, group *Group) (*Group, error) {
	query := `
insert into ad.groups (
                       code,
                       description,
                       created_at,
                       updated_at,
                       deleted_at
                       )
                   values (
                       $1,
                       $2,
                       now(),
                       null,
                       null
                   ) returning id,
                               code,
                               description,
                               created_at,
                               updated_at,
                               deleted_at
`
	var g Group
	err := tx.GetContext(ctx, &g, query, group.Code, group.Description)
	if err != nil {
		return nil, err
	}

	return &g, nil
}

func (r *repository) Update(ctx context.Context, group *Group) (*Group, error) {
	query := `
	update ad.groups 
	   set code = :code,
		   description = :description,
		   updated_at = now()
	 where id = :id returning  id,
                               code,
                               description,
                               created_at,
                               updated_at,
                               deleted_at;
`

	rows, err := r.conn.NamedQueryContext(ctx, query, group)
	if err != nil {
		return nil, err
	}

	var g Group
	for rows.Next() {
		err = rows.StructScan(&g)
		if err != nil {
			return nil, err
		}
	}

	return &g, nil
}

func (r *repository) Insert(ctx context.Context, group *Group) (*Group, error) {
	query := `
insert into ad.groups (
                       code,
                       description,
                       created_at,
                       updated_at,
                       deleted_at
                       )
                   values (
                       :code,
                       :description,
                       now(),
                       null,
                       null
                   ) returning id,
                               code,
                               description,
                               created_at,
                               updated_at,
                               deleted_at;
`

	rows, err := r.conn.NamedQueryContext(ctx, query, group)
	if err != nil {
		return nil, err
	}

	var g Group
	for rows.Next() {
		err = rows.StructScan(&g)
		if err != nil {
			return nil, err
		}
	}

	if g.Id == 0 {
		return nil, ErrGroupAlreadyExists
	}

	return &g, nil
}

func (r *repository) GetById(ctx context.Context, id int) (*Group, error) {
	return GetByIdDb(ctx, r.conn, id)
}

func (r *repository) GetByCodeDb(ctx context.Context, code string) (*Group, error) {
	return GetByCodeDb(ctx, r.conn, code)
}

func GetByCodeDb(ctx context.Context, conn *sqlx.DB, code string) (*Group, error) {
	query := `
		select id,
			   code,
			   description,
			   created_at,
			   updated_at,
			   deleted_at
		from ad.groups
	   where code = $1
`
	var group Group
	err := conn.GetContext(ctx, &group, query, code)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrEntityNotFound
		}
		return nil, err
	}

	return &group, nil
}

func GetByIdDb(ctx context.Context, conn *sqlx.DB, id int) (*Group, error) {
	query := `
		select id,
			   code,
			   description,
			   created_at,
			   updated_at,
			   deleted_at
		from ad.groups
	   where id = $1
`
	var group Group
	err := conn.GetContext(ctx, &group, query, id)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (r *repository) List(ctx context.Context) ([]Group, error) {
	query := `
		select id,
			   code,
			   description,
			   created_at,
			   updated_at,
			   deleted_at
		from ad.groups
	   where deleted_at is null;
`
	groups := make([]Group, 0)
	err := r.conn.SelectContext(ctx, &groups, query)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	query := `
		update ad.groups 
			set deleted_at = now()
		where id = $1;
`
	_, err := r.conn.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
