package userroute

import (
	"github.com/hardstylez72/bzdacs/pkg/group"
	"github.com/hardstylez72/bzdacs/pkg/route"
)

type Group group.Group
type OriginalRoute route.Route

type Route struct {
	OriginalRoute
	IsExcluded bool `json:"isExcluded" db:"is_excluded" validate:"required"`
} // @name userRouteRoute

type RouteWithGroups struct {
	Groups []Group `json:"groups" validate:"required"`
	Route
} // @name userRouteRouteWithGroups
