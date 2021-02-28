package user

import (
	"context"
	"github.com/hardstylez72/bzdacs/internal"
	"github.com/hardstylez72/bzdacs/pkg/group"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/namespace"
	"github.com/hardstylez72/bzdacs/pkg/relations/usergroup"
	"github.com/hardstylez72/bzdacs/pkg/system"
	"github.com/hardstylez72/bzdacs/pkg/user"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var ErrEntityNotFound = errors.New("entity not found")

type repository struct {
	conn *sqlx.DB
}

func NewRepository(conn *sqlx.DB) *repository {
	return &repository{conn: conn}
}

func (r *repository) Insert(ctx context.Context, u *User) (*User, error) {

	if u.Login == internal.Admin {
		// admin does not need to register, app does it for first start
		return u, nil
	}

	tx, err := r.conn.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	txx := storage.WrapSqlxTx(tx)

	sysUser, err := InsertLL(ctx, txx, u)
	if err != nil {
		return nil, err
	}

	if internal.HasGuest {
		s, err := system.GetLL(ctx, txx, 0, internal.SystemName)
		if err != nil {
			return nil, err
		}

		ns, err := namespace.GetLL(ctx, txx, s.Id, 0, internal.Namespace)
		if err != nil {
			return nil, err
		}

		appUser, err := user.InsertLL(ctx, txx, &user.User{
			ExternalId:  sysUser.Login,
			NamespaceId: ns.Id,
		})

		if err != nil {
			return nil, err
		}

		g, err := group.GetByCodeLL(ctx, txx, internal.GuestGroupName, ns.Id)
		if err != nil {
			return nil, err
		}

		_, err = usergroup.InsertPairLL(ctx, txx, g.Id, appUser.Id)
		if err != nil {
			return nil, err
		}
	}

	return sysUser, nil
}
func InsertLL(ctx context.Context, driver storage.SqlDriver, entity *User) (*User, error) {
	query := `
insert into sys_users (
                       login,
					   password,
                       created_at,
					   updated_at,
					   deleted_at
                       )
                   values (
                       :login,
					   :password,
					   now(),
					   now(),
					   :deleted_at
                   ) returning id,
                               login,
							   created_at,
							   password,
							   updated_at,
							   deleted_at
`

	rows, err := driver.NamedQueryContext(ctx, query, entity)
	if err != nil {
		return nil, storage.PgError(err)
	}

	var g User
	for rows.Next() {
		err = rows.StructScan(&g)
		if err != nil {
			return nil, storage.PgError(err)
		}
	}
	return &g, nil
}

func (r *repository) GetByParams(ctx context.Context, login, password string) (*User, error) {
	return GetByParams(ctx, r.conn, login, password)
}
func GetByParams(ctx context.Context, driver storage.SqlDriver, login, password string) (*User, error) {
	query := `
		select id,
			   login,
		       created_at,
			   updated_at,
			   deleted_at,
			   password
		from sys_users
	   where login = $1
 		 and password = $2
`
	var user User
	err := driver.GetContext(ctx, &user, query, login, password)
	if err != nil {
		return nil, storage.PgError(err)
	}

	return &user, nil
}

func (r *repository) GetById(ctx context.Context, id, namespaceId int) (*User, error) {
	return GetByIdLL(ctx, r.conn, id, namespaceId)
}

func GetByIdLL(ctx context.Context, driver storage.SqlDriver, id, namespaceId int) (*User, error) {
	query := `
		select id,
			   external_id,
		       created_at,
			   updated_at,
			   deleted_at,
			   namespace_id
		from users
	   where namespace_id = $2
 		 and id = $1
`
	var user User
	err := driver.GetContext(ctx, &user, query, id, namespaceId)
	if err != nil {
		return nil, storage.PgError(err)
	}

	return &user, nil
}
