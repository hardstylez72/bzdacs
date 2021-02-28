package userroute

import (
	"github.com/hardstylez72/bzdacs/pkg/group"
	"github.com/hardstylez72/bzdacs/pkg/route"
)

type Group group.Group
type RouteOrig route.Route

type Route struct {
	RouteOrig
	IsExcluded bool `json:"isExcluded" db:"is_excluded" validate:"required"`
} // @name userRouteRoute

type RouteWithGroups struct {
	Groups []Group `json:"groups" validate:"required"`
	Route
} // @name userRouteRouteWithGroups
