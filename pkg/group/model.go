package group

import (
	"github.com/hardstylez72/bblog/ad/pkg/util"
	"time"
)

type Group struct {
	Id          int               `json:"id" db:"id"`
	Code        string            `json:"code" db:"code"`
	Description string            `json:"description" db:"description"`
	CreatedAt   time.Time         `json:"createdAt" db:"created_at"`
	UpdatedAt   util.JsonNullTime `json:"updatedAt" db:"updated_at"`
	DeletedAt   util.JsonNullTime `json:"deletedAt" db:"deleted_at"`
}
