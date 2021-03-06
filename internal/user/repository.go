package user

import (
	"context"
	"github.com/hardstylez72/bzdacs/pkg/group"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/namespace"
	"github.com/hardstylez72/bzdacs/pkg/relations/usergroup"
	"github.com/hardstylez72/bzdacs/pkg/system"
	"github.com/hardstylez72/bzdacs/pkg/user"
	"github.com/spf13/viper"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

func NewRepository(conn *sqlx.DB) *repository {
	return &repository{conn: conn}
}

func (r *repository) Insert(ctx context.Context, u *SysUser) (*SysUser, error) {

	tx, err := r.conn.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	txx := storage.WrapSqlxTx(tx)

	sysUser, err := Insert(ctx, txx, u)
	if err != nil {
		return nil, err
	}

	SystemName := viper.GetString("app.system")
	Namespace := viper.GetString("app.namespace")
	GuestGroupName := viper.GetString("app.guestGroupName")
	HasGuest := viper.GetBool("app.hasGuest")
	AdminLogin := viper.GetString("user.login")

	if AdminLogin == u.Login {
		return sysUser, nil
	}

	if HasGuest {
		s, err := system.Get(ctx, txx, 0, SystemName)
		if err != nil {
			return nil, err
		}

		ns, err := namespace.Get(ctx, txx, s.Id, 0, Namespace)
		if err != nil {
			return nil, err
		}

		appUser, err := user.Insert(ctx, txx, &user.User{
			ExternalId:  sysUser.Login,
			NamespaceId: ns.Id,
		})
		if err != nil {
			return nil, err
		}

		g, err := group.GetByCode(ctx, txx, GuestGroupName, ns.Id)
		if err != nil {
			return nil, err
		}

		_, err = usergroup.InsertPair(ctx, txx, g.Id, appUser.Id)
		if err != nil {
			return nil, err
		}
	}

	return sysUser, nil
}
func Insert(ctx context.Context, driver storage.SqlDriver, entity *SysUser) (*SysUser, error) {
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
							   deleted_at;
`

	rows, err := driver.NamedQueryContext(ctx, query, entity)
	if err != nil {
		return nil, storage.PgError(err)
	}

	var g SysUser
	for rows.Next() {
		err = rows.StructScan(&g)
		if err != nil {
			return nil, storage.PgError(err)
		}
	}
	return &g, nil
}

func (r *repository) GetByParams(ctx context.Context, login, password string) (*SysUser, error) {
	return GetByParams(ctx, r.conn, login, password)
}
func GetByParams(ctx context.Context, driver storage.SqlDriver, login, password string) (*SysUser, error) {
	query := `
		select id,
			   login,
		       created_at,
			   updated_at,
			   deleted_at,
			   password
		from sys_users
	   where login = $1
`
	var user SysUser
	err := driver.GetContext(ctx, &user, query, login)
	if err != nil {
		return nil, storage.PgError(err)
	}

	return &user, nil
}

func (r *repository) GetById(ctx context.Context, id, namespaceId int) (*SysUser, error) {
	return GetById(ctx, r.conn, id, namespaceId)
}

func GetById(ctx context.Context, driver storage.SqlDriver, id, namespaceId int) (*SysUser, error) {
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
	var user SysUser
	err := driver.GetContext(ctx, &user, query, id, namespaceId)
	if err != nil {
		return nil, storage.PgError(err)
	}

	return &user, nil
}
