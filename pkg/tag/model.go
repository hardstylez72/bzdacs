package tag

import (
	"github.com/hardstylez72/bzdacs/pkg/util"
	"time"
)

type Tag struct {
	Id          int               `json:"id" db:"id"`
	Name        string            `json:"name" db:"name"`
	CreatedAt   time.Time         `json:"createdAt" db:"created_at" swaggertype:"string"`
	UpdatedAt   time.Time         `json:"updatedAt" db:"updated_at" swaggertype:"string"`
	DeletedAt   util.JsonNullTime `json:"deletedAt" db:"deleted_at" swaggertype:"string" extensions:"x-nullable"`
	NamespaceId int               `json:"namespaceId" db:"namespace_id"`
}
