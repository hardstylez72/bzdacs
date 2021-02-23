package route

import "github.com/hardstylez72/bzdacs/pkg/util"

type Route struct {
	Id          int      `json:"id" db:"id" validate:"required"`
	Route       string   `json:"route" db:"route" validate:"required"`
	Method      string   `json:"method" db:"method" validate:"required"`
	Description string   `json:"description" db:"description" validate:"required"`
	Tags        []string `json:"tags" validate:"required"`
	NamespaceId int      `json:"namespaceId" db:"namespace_id" validate:"required"`
	util.Times
}
