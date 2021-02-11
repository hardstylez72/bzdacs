package systemnamespace

import (
	"context"
	"errors"
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


func InsertConn(ctx context.Context, conn storage.SqlDriver, pair Pair) error {
	query := `insert into systems_namespaces (system_id, namespace_id) values ($1, $2)`
	_, err := conn.ExecContext(ctx, query, pair.SystemId, pair.NamespaceId)
	if err != nil {
		return err
	}
	return nil
}