package route

import "github.com/hardstylez72/bzdacs/pkg/util"

type Route struct {
	Id          int      `json:"id" db:"id"`
	Route       string   `json:"route" db:"route"`
	Method      string   `json:"method" db:"method"`
	Description string   `json:"description" db:"description"`
	Tags        []string `json:"tags" `
	NamespaceId int      `json:"namespaceId" db:"namespace_id"`
	util.Times
}
