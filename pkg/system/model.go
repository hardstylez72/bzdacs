package system

import (
	"github.com/hardstylez72/bzdacs/pkg/util"
	"time"
)

type System struct {
	Id        int               `json:"id" db:"id" validate:"required"`
	Name      string            `json:"name" db:"name" validate:"required"`
	CreatedAt time.Time         `json:"createdAt" db:"created_at" swaggertype:"string" validate:"required"`
	UpdatedAt time.Time         `json:"updatedAt" db:"updated_at" swaggertype:"string" validate:"required"`
	DeletedAt util.JsonNullTime `json:"deletedAt" db:"deleted_at" swaggertype:"string" extensions:"x-nullable"`
} // @name System
