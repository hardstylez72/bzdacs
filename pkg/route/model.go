package route

import (
	"github.com/hardstylez72/bzdacs/pkg/util"
	"time"
)

type Route struct {
	Id          int               `json:"id" db:"id"`
	Route       string            `json:"route" db:"route"`
	Method      string            `json:"method" db:"method"`
	Description string            `json:"description" db:"description"`
	CreatedAt   time.Time         `json:"createdAt" db:"created_at" swaggertype:"string"`
	UpdatedAt   time.Time         `json:"updatedAt" db:"updated_at" swaggertype:"string"`
	DeletedAt   util.JsonNullTime `json:"deletedAt" db:"deleted_at" swaggertype:"string" extensions:"x-nullable"`
	NamespaceId int               `json:"namespaceId" db:"namespace_id"`
	Tags        []string          `json:"tags" `
}
