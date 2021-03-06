package acs

import "github.com/hardstylez72/bzdacs/pkg/route"

type Route route.Route

type UserRoute struct {
	Route
	IsExcluded bool `json:"isExcluded" db:"is_excluded" validate:"required"`
}
