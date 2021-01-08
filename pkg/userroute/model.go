package userroute

import (
	"github.com/hardstylez72/bzdacs/pkg/group"
	"github.com/hardstylez72/bzdacs/pkg/route"
)

type Route struct {
	IsExcluded bool `json:"isExcluded" db:"is_excluded"`
	route.Route
}

type RouteWithGroups struct {
	Groups []group.Group `json:"groups"`
	RouteExt
}

type RouteExt struct {
	Route
	IsOverwritten bool `json:"isOverwritten"`
	IsIndependent bool `json:"isIndependent"`
}
