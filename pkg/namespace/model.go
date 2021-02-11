package namespace

import (
	"github.com/hardstylez72/bzdacs/pkg/util"
	"time"
)

type Namespace struct {
	Id        int               `json:"id" db:"id"`
	Name      string            `json:"name" db:"name"`
	CreatedAt time.Time         `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time         `json:"updatedAt" db:"updated_at"`
	DeletedAt util.JsonNullTime `json:"deletedAt" db:"deleted_at"`
}

type NamespaceExt struct {
	Namespace
	SystemId int
}