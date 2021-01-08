package main

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/hardstylez72/bblog/ad/pkg/group"
	"github.com/hardstylez72/bblog/ad/pkg/grouproute"
	"github.com/hardstylez72/bblog/ad/pkg/route"
	"github.com/hardstylez72/bblog/ad/pkg/user"
	"github.com/hardstylez72/bblog/ad/pkg/usergroup"
	"github.com/spf13/viper"
	"net/http"
)

const (
	sysGroupCode        = "sys_group"
	sysGroupDescription = "internal system group"
	sysRouteTag         = "sys"
)

func resolveGroup(ctx context.Context, rep group.Repository) (*group.Group, error) {

	g, err := rep.GetByCodeDb(ctx, sysGroupCode)
	if err != nil {
		if err != group.ErrEntityNotFound {
			return nil, err
		}
	}
	if g != nil {
		err = rep.Delete(ctx, g.Id)
		if err != nil {
			return nil, err
		}
	}

	g, err = rep.Insert(ctx, &group.Group{Code: sysGroupCode, Description: sysGroupDescription})
	if err != nil {
		return nil, err
	}
	return g, nil
}

func resolveUser(ctx context.Context, rep user.Repository) (*user.User, error) {
	var login = viper.GetString("user.login")

	u, err := rep.GetByExternalId(ctx, login)
	if err != nil {
		if err != user.ErrEntityNotFound {
			return nil, err
		}
	}
	if u != nil {
		err = rep.Delete(ctx, u.Id)
		if err != nil {
			return nil, err
		}
	}

	u, err = rep.Insert(ctx, &user.User{ExternalId: login})
	if err != nil {
		return nil, err
	}
	return u, nil
}

func resolveUserAndGroup(ctx context.Context, rep usergroup.Repository, userId, groupId int) error {
	_, err := rep.Insert(ctx, []usergroup.Pair{{GroupId: groupId, UserId: userId}})
	return err
}

func buildRoutes(r chi.Router) []route.Route {
	rs := make([]route.Route, 0)
	for _, r := range r.Routes() {

		if apiPathPrefix+"/*" == r.Pattern {
			continue
		}

		rs = append(rs, route.Route{
			Route:       apiPathPrefix + r.Pattern,
			Method:      http.MethodPost,
			Description: "",
		})
	}
	return rs
}

func resolveRoutes(ctx context.Context, rep route.Repository, rs []route.Route) ([]route.Route, error) {

	out := make([]route.Route, 0, len(rs))

	for _, r := range rs {

		rr, err := rep.GetByMethodAndRoute(ctx, r.Route, r.Method)
		if err != nil {
			if err != route.ErrEntityNotFound {
				return nil, err
			}
		}

		if rr != nil {
			err := rep.Delete(ctx, rr.Id)
			if err != nil {
				return nil, err
			}
		}

		rr, err = rep.InsertWithTags(ctx, &r, []string{sysRouteTag})
		if err != nil {
			return nil, err
		}

		out = append(out, rr.Route)
	}
	return out, nil
}

func resolveGroupAndRoutes(ctx context.Context, rep grouproute.Repository, rs []route.Route, groupId int) error {
	groupRoutePairs := make([]grouproute.Pair, 0)
	for _, r := range rs {
		groupRoutePairs = append(groupRoutePairs, grouproute.Pair{
			GroupId: groupId,
			RouteId: r.Id,
		})
	}
	_, err := rep.Insert(ctx, groupRoutePairs)
	return err
}
