package main

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/hardstylez72/bzdacs/pkg/group"
	"github.com/hardstylez72/bzdacs/pkg/grouproute"
	"github.com/hardstylez72/bzdacs/pkg/route"
	"github.com/hardstylez72/bzdacs/pkg/user"
	"github.com/hardstylez72/bzdacs/pkg/usergroup"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

const (
	sysGroupCode        = "sys_group"
	sysGuestGroupCode   = "sys_guest"
	sysGroupDescription = "internal system group"
	sysGuestDescription = "internal system group for observing"
	sysRouteTag         = "sys"
)

type Option struct {
	Force bool
}

func (s *Server) initialize(ctx context.Context, force bool, r chi.Router) error {
	u, err := resolveUser(ctx, s.repository.user, &Option{Force: force})
	if err != nil {
		return err
	}
	g, err := resolveGroup(ctx, s.repository.group, &Option{Force: force})
	if err != nil {
		return err
	}
	err = resolveUserAndGroup(ctx, s.repository.usergroup, u.Id, g.Id, &Option{Force: force})
	if err != nil {
		return err
	}

	rs := buildRoutes(r)

	rs, err = resolveRoutes(ctx, s.repository.route, rs, &Option{Force: force})
	if err != nil {
		return err
	}

	err = resolveGroupAndRoutes(ctx, s.repository.grouproute, rs, g.Id, &Option{Force: force})
	if err != nil {
		return err
	}

	guestRs := filterGuestRoutes(rs)

	guestGroup, err := resolveGuestGroup(ctx, s.repository.group, &Option{Force: force})
	if err != nil {
		return err
	}

	err = resolveGroupAndRoutes(ctx, s.repository.grouproute, guestRs, guestGroup.Id, &Option{Force: force})
	if err != nil {
		return err
	}
	return nil
}

func resolveGroup(ctx context.Context, rep group.Repository, o *Option) (*group.Group, error) {

	g, err := rep.GetByCode(ctx, sysGroupCode)
	if err != nil {
		if err != group.ErrEntityNotFound {
			return nil, err
		}
	}
	if g != nil {
		if !o.Force {
			return g, nil
		}
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

func resolveGuestGroup(ctx context.Context, rep group.Repository, o *Option) (*group.Group, error) {

	g, err := rep.GetByCode(ctx, sysGuestGroupCode)
	if err != nil {
		if err != group.ErrEntityNotFound {
			return nil, err
		}
	}
	if g != nil {
		if !o.Force {
			return g, nil
		}
		err = rep.Delete(ctx, g.Id)
		if err != nil {
			return nil, err
		}
	}

	g, err = rep.Insert(ctx, &group.Group{Code: sysGuestGroupCode, Description: sysGuestDescription})
	if err != nil {
		return nil, err
	}

	return g, nil
}

func resolveUser(ctx context.Context, rep user.Repository, o *Option) (*user.User, error) {
	var login = viper.GetString("user.login")

	u, err := rep.GetByExternalId(ctx, login)
	if err != nil {
		if err != user.ErrEntityNotFound {
			return nil, err
		}
	}
	if u != nil {
		if !o.Force {
			return u, nil
		}
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

func resolveUserAndGroup(ctx context.Context, rep usergroup.Repository, userId, groupId int, o *Option) error {
	isExist, err := rep.IsPairExist(ctx, usergroup.Pair{
		GroupId: groupId,
		UserId:  userId,
	})
	if err != nil {
		if err != usergroup.ErrEntityNotFound {
			return err
		}
	}

	if isExist {
		if !o.Force {
			return nil
		}
		err = rep.Delete(ctx, []usergroup.Pair{{
			GroupId: groupId,
			UserId:  userId,
		}})
		if err != nil {
			return err
		}
	}
	_, err = rep.Insert(ctx, usergroup.Pair{
		GroupId: groupId,
		UserId:  userId,
	})
	if err != nil {
		return err
	}

	return nil
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

func filterGuestRoutes(in []route.Route) []route.Route {
	out := make([]route.Route, 0)

	for _, r := range in {
		if strings.Contains(r.Route, "/update") ||
			strings.Contains(r.Route, "/delete") ||
			strings.Contains(r.Route, "/create") {
			continue
		}
		out = append(out, r)
	}
	return out
}

func resolveRoutes(ctx context.Context, rep route.Repository, rs []route.Route, o *Option) ([]route.Route, error) {

	out := make([]route.Route, 0, len(rs))

	for _, r := range rs {

		rr, err := rep.GetByMethodAndRoute(ctx, r.Route, r.Method)
		if err != nil {
			if err != route.ErrEntityNotFound {
				return nil, err
			}
		}

		if rr != nil {
			if !o.Force {
				out = append(out, rr.Route)
				continue
			}
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

func resolveGroupAndRoutes(ctx context.Context, rep grouproute.Repository, rs []route.Route, groupId int, o *Option) error {

	for _, r := range rs {
		isExist, err := rep.IsPairExist(ctx, grouproute.Pair{
			GroupId: groupId,
			RouteId: r.Id,
		})
		if err != nil {
			if err != grouproute.ErrEntityNotFound {
				return err
			}
		}

		if isExist {
			if !o.Force {
				continue
			}
			err = rep.Delete(ctx, []grouproute.Pair{{
				GroupId: groupId,
				RouteId: r.Id,
			}})
			if err != nil {
				return err
			}
		}
		_, err = rep.Insert(ctx, grouproute.Pair{
			GroupId: groupId,
			RouteId: r.Id,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
