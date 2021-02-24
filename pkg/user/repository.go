package user

import (
	"context"

	"github.com/hardstylez72/bzdacs/pkg/infra/storage"

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
func (r *repository) Update(ctx context.Context, user *User) (*User, error) {
	return UpdateLL(ctx, r.conn, user)
}
func UpdateLL(ctx context.Context, driver storage.SqlDriver, user *User) (*User, error) {
	query := `
	update users 
	   set external_id = :external_id,
	       updated_at = now()
	 where id = :id 
 returning id,
		   external_id,
		   created_at,
		   updated_at,
		   deleted_at,
     	   namespace_id
`

	rows, err := driver.NamedQueryContext(ctx, query, user)
	if err != nil {
		return nil, storage.PgError(err)
	}

	var u User
	for rows.Next() {
		err = rows.StructScan(&u)
		if err != nil {
			return nil, storage.PgError(err)
		}
	}

	return &u, nil
}
func (r *repository) Insert(ctx context.Context, u *User) (*User, error) {
	return InsertLL(ctx, r.conn, u)
}
func InsertLL(ctx context.Context, driver storage.SqlDriver, entity *User) (*User, error) {
	query := `
insert into users (
                       external_id,
                       created_at,
					   updated_at,
					   deleted_at,
					   namespace_id
                       )
                   values (
                       :external_id,
					   now(),
					   now(),
					   :deleted_at,
					   :namespace_id
                   ) returning id,
                               external_id,
							   created_at,
							   updated_at,
							   deleted_at,
							   namespace_id
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

func (r *repository) GetByLogin(ctx context.Context, login string, namespaceId int) (*User, error) {
	return GetByExternalId(ctx, r.conn, login, namespaceId)
}
func GetByExternalId(ctx context.Context, driver storage.SqlDriver, login string, namespaceId int) (*User, error) {
	query := `
		select id,
			   external_id,
		       created_at,
			   updated_at,
			   deleted_at,
			   namespace_id
		from users
	   where namespace_id = $2
 		 and external_id = $1
`
	var user User
	err := driver.GetContext(ctx, &user, query, login, namespaceId)
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

func (r *repository) Delete(ctx context.Context, id, namespaceId int) error {
	return DeleteLL(ctx, r.conn, id, namespaceId)
}
func DeleteLL(ctx context.Context, driver storage.SqlDriver, id, namespaceId int) error {
	query := `
		update users 
			set deleted_at = now()
		where id = $1 and namespace_id = $2
`
	_, err := driver.ExecContext(ctx, query, id, namespaceId)
	if err != nil {
		return err
	}

	return nil
}
