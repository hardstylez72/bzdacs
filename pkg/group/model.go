package group

import (
	"github.com/hardstylez72/bzdacs/pkg/util"
)

type Group struct {
	Id          int    `json:"id" db:"id" validate:"required"`
	Code        string `json:"code" db:"code" validate:"required"`
	Description string `json:"description" db:"description" validate:"required"`
	NamespaceId int    `json:"namespaceId" db:"namespace_id" validate:"required"`
	util.Times
} // @name Group
