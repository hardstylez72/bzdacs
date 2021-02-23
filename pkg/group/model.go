package group

import (
	"github.com/hardstylez72/bzdacs/pkg/util"
)

type Group struct {
	Id          int    `json:"id" db:"id"`
	Code        string `json:"code" db:"code"`
	Description string `json:"description" db:"description"`
	NamespaceId int    `json:"namespaceId" db:"namespace_id"`
	util.Times
}
