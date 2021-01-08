package route

import (
	"github.com/hardstylez72/bblog/ad/pkg/util"
	"time"
)

type Route struct {
	Id          int               `json:"id" db:"id"`
	Route       string            `json:"route" db:"route"`
	Method      string            `json:"method" db:"method"`
	Description string            `json:"description" db:"description"`
	CreatedAt   time.Time         `json:"createdAt" db:"created_at"`
	UpdatedAt   util.JsonNullTime `json:"updatedAt" db:"updated_at"`
	DeletedAt   util.JsonNullTime `json:"deletedAt" db:"deleted_at"`
}

type RouteWithTags struct {
	Route
	Tags []string `json:"tags" `
}
