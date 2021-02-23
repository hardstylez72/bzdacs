package acs

import (
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/hardstylez72/bzdacs/pkg/relations/usergroup"
	"github.com/hardstylez72/bzdacs/pkg/relations/userroute"
	"github.com/hardstylez72/bzdacs/pkg/user"
	"github.com/hardstylez72/bzdacs/pkg/util"
	"net/http"
)

const (
	loginHeaderName                = "login"
	requestedRouteHeaderName       = "route"
	requestedRouteMethodHeaderName = "method"
)

var (
	ErrLoginNotFoundInRequest                = errors.New("request does not contains login")
	ErrRequestedRouteNotFoundInRequest       = errors.New("request does not contains requested route")
	ErrRequestedRouteMethodNotFoundInRequest = errors.New("request does not contains requested route method")
)

type controller struct {
	userRouteRepository userroute.Repository
	userRepository      user.Repository
	userGroupRepository usergroup.Repository
	validator           *validator.Validate
}

func NewController(userRouteRepository userroute.Repository, userRepository user.Repository, userGroupRepository usergroup.Repository) *controller {
	return &controller{
		userRouteRepository: userRouteRepository,
		userRepository:      userRepository,
		userGroupRepository: userGroupRepository,
		validator:           validator.New(),
	}
}

func (c *controller) accessCheck(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	login, err := getLoginFromRequest(r)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
	}

	route, err := getRequestedRouteFromRequest(r)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
	}

	routeMethod, err := getRequestedMethodFromRequest(r)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
	}

	u, err := c.userRepository.GetByExternalId(ctx, login)
	if err != nil {
		if err == user.ErrEntityNotFound {
			util.NewResp(w, r).Error(err).Status(http.StatusNotFound).Send()
		} else {
			util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		}
		return
	}

	userRoutes, err := c.userRouteRepository.RoutesBelongUser(ctx, u.Id)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	isAccessAllowed := AccessAllowed(route, routeMethod, userRoutes)

	groups, err := c.userGroupRepository.ListUserGroups(ctx, u.Id)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
	}

	util.NewResp(w, r).Json(newCheckAccessResponse(groups, isAccessAllowed, login)).Status(http.StatusOK).Send()
}

func getRequestedMethodFromRequest(r *http.Request) (string, error) {
	method := r.Header.Get(requestedRouteMethodHeaderName)
	if method != "" {
		return method, nil
	}

	return "", ErrRequestedRouteMethodNotFoundInRequest
}

func getRequestedRouteFromRequest(r *http.Request) (string, error) {
	route := r.Header.Get(requestedRouteHeaderName)
	if route != "" {
		return route, nil
	}

	return "", ErrRequestedRouteNotFoundInRequest
}

func getLoginFromRequest(r *http.Request) (string, error) {
	login := r.Header.Get(loginHeaderName)
	if login != "" {
		return login, nil
	}

	return "", ErrLoginNotFoundInRequest
}

func AccessAllowed(requestedRoute, requestedMethod string, routes []userroute.RouteWithGroups) bool {
	for _, r := range routes {
		if r.Method == requestedMethod {
			routeMatched := isRouteAPartOfBaseRoutePrefix(requestedRoute, r.Route.Route)

			if r.IsExcluded {
				if routeMatched {
					return false
				}
			} else {
				if routeMatched {
					return true
				}
			}
		}
	}

	return false
}

func isRouteAPartOfBaseRoutePrefix(checked, base string) bool {

	if len(checked) > len(base) {
		return false
	}

	for checkedI := 0; checkedI < len(checked); checkedI++ {
		if checked[checkedI] != base[checkedI] {
			return false
		}
	}

	return true
}

func (c *controller) Mount(r chi.Router) {
	r.Post("/v1/user/access/check", c.accessCheck)
}
