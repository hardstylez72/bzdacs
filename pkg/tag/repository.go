package tag

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

var (
	ErrNotFound = errors.New("tag not found")
)

func NewRepository(conn *sqlx.DB) *repository {
	return &repository{conn: conn}
}

func (r *repository) Update(ctx context.Context, tag *Tag) (*Tag, error) {
	return UpdateLL(ctx, r.conn, tag)
}

func UpdateLL(ctx context.Context, driver storage.SqlDriver, tag *Tag) (*Tag, error) {
	query := `
	update tags 
	   set name = :name,
		   updated_at = now()
	 where id = :id returning  id,
                               name,
                               created_at,
                               updated_at,
                               deleted_at,
							   namespace_id
	`

	rows, err := driver.NamedQueryContext(ctx, query, tag)
	if err != nil {
		return nil, storage.PgError(err)
	}

	var t Tag
	for rows.Next() {
		err = rows.StructScan(&t)
		if err != nil {
			return nil, storage.PgError(err)
		}
	}

	return &t, nil
}

func (r *repository) Insert(ctx context.Context, tag *Tag) (*Tag, error) {
	return InsertLL(ctx, r.conn, tag)
}
func InsertLL(ctx context.Context, driver storage.SqlDriver, tag *Tag) (*Tag, error) {
	query := `
	insert into tags (
                       name,
                       created_at,
                       updated_at,
                       deleted_at,
					   namespace_id
                       )
                   values (
                       :name,
                       now(),
                       now(),
                       null,
					   :namespace_id
                   ) returning id,
                               name,
                               created_at,
                               updated_at,
                               deleted_at,
							   namespace_id
	`

	rows, err := driver.NamedQueryContext(ctx, query, tag)
	if err != nil {
		return nil, storage.PgError(err)
	}

	var t Tag
	for rows.Next() {
		err = rows.StructScan(&t)
		if err != nil {
			return nil, storage.PgError(err)
		}
	}

	return &t, nil
}

func (r *repository) GetById(ctx context.Context, id int) (*Tag, error) {
	return GetByIdLL(ctx, r.conn, id)
}

func GetByIdLL(ctx context.Context, driver storage.SqlDriver, id int) (*Tag, error) {
	query := `
		select id,
			   name,
			   created_at,
			   updated_at,
			   deleted_at
		from tags
	   where id = $1
`
	var tag Tag
	err := driver.GetContext(ctx, &tag, query, id)
	if err != nil {
		return nil, storage.PgError(err)
	}

	return &tag, nil
}
func (r *repository) List(ctx context.Context, pattern string, namespaceId int) ([]Tag, error) {
	return ListLL(ctx, r.conn, pattern, namespaceId)
}

func ListLL(ctx context.Context, driver storage.SqlDriver, pattern string, namespaceId int) ([]Tag, error) {

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	builder := psql.Select(`
 			   id,
			   name,
			   created_at,
			   updated_at,
			   deleted_at,
			   namespace_id
			`).From("tags")

	if pattern != "" {
		pattern = "%" + pattern + "%"
		builder = builder.Where(sq.Like{"name": pattern})
	}
	builder = builder.Where(sq.Eq{"deleted_at": nil})
	builder = builder.Where(sq.Eq{"namespace_id": namespaceId})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, storage.PgError(err)
	}
	tags := make([]Tag, 0)
	err = driver.SelectContext(ctx, &tags, query, args...)
	if err != nil {
		return nil, storage.PgError(err)
	}

	return tags, nil
}

func (r *repository) Delete(ctx context.Context, id, namespaceId int) error {
	return DeleteLL(ctx, r.conn, id, namespaceId)
}
func DeleteLL(ctx context.Context, driver storage.SqlDriver, id, namespaceId int) error {
	query := `
		update tags 
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
