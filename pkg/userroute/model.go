package userroute

import (
	"github.com/hardstylez72/bzdacs/pkg/group"
	"github.com/hardstylez72/bzdacs/pkg/route"
)

type UserRoute struct {
	route.Route
	IsExcluded bool `json:"isExcluded" db:"is_excluded"`
}

type RouteWithGroups struct {
	Groups []group.Group `json:"groups"`
	RouteExt
}

type RouteExt struct {
	route.Route
	IsExcluded    bool `json:"isExcluded" db:"is_excluded"`
	IsOverwritten bool `json:"isOverwritten"`
	IsIndependent bool `json:"isIndependent"`
}
