package user

import (
	"github.com/hardstylez72/bzdacs/pkg/util"
	"time"
)

type User struct {
	Id        int               `json:"id" db:"id"  validate:"required"`
	Login     string            `json:"login" db:"login"  validate:"required"`
	Password  string            `json:"password" db:"password" validate:"required"`
	CreatedAt time.Time         `json:"createdAt" db:"created_at"  swaggertype:"string" validate:"required"`
	UpdatedAt time.Time         `json:"updatedAt" db:"updated_at" swaggertype:"string" validate:"required"`
	DeletedAt util.JsonNullTime `json:"deletedAt" db:"deleted_at" swaggertype:"string" extensions:"x-nullable"`
} // @name SysUser
