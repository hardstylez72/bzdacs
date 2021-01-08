package userroute

import (
	"github.com/hardstylez72/bblog/ad/pkg/group"
	"github.com/hardstylez72/bblog/ad/pkg/route"
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
