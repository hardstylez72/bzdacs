package acs

import (
	"context"
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/hardstylez72/bzdacs/pkg/namespace"
	"github.com/hardstylez72/bzdacs/pkg/system"
	"github.com/hardstylez72/bzdacs/pkg/user"
	"github.com/hardstylez72/bzdacs/pkg/util"
	"net/http"
)

const (
	loginHeaderName                = "login"
	requestedRouteHeaderName       = "route"
	requestedRouteMethodHeaderName = "method"
	requestedNamespaceHeaderName   = "namespace"
	requestedSystemHeaderName      = "system"
)

type inputParams struct {
	Login       string
	Route       string
	Namespace   string
	System      string
	RouteMethod string
}

type controller struct {
	userRepository      user.Repository
	acsRepository       Repository
	namespaceRepository namespace.Repository
	systemRepository    system.Repository
	validator           *validator.Validate
}

var (
	ErrLoginNotFoundInRequest                = errors.New("request does not contains login")
	ErrRequestedRouteNotFoundInRequest       = errors.New("request does not contains requested route")
	ErrRequestedRouteMethodNotFoundInRequest = errors.New("request does not contains requested route method")
	ErrRequestedSystemNotFoundInRequest      = errors.New("request does not contains requested system")
	ErrRequestedNamespaceNotFoundInRequest   = errors.New("request does not contains requested namespace")
)

type Repository interface {
	GetUserRoutes(ctx context.Context, userId, namespaceId int) ([]UserRoute, error)
}

func NewController(userRepository user.Repository, acsRepository Repository, namespaceRepository namespace.Repository, systemRepository system.Repository) *controller {
	return &controller{
		userRepository:      userRepository,
		acsRepository:       acsRepository,
		namespaceRepository: namespaceRepository,
		systemRepository:    systemRepository,
		validator:           validator.New(),
	}
}

func (c *controller) accessCheck(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params, err := getInputParams(r)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	s, err := c.systemRepository.Get(ctx, system.GetByName, params.System)
	if err != nil {
		if err == user.ErrEntityNotFound {
			util.NewResp(w, r).Error(err).Status(http.StatusNotFound).Send()
		} else {
			util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		}
		return
	}

	ns, err := c.namespaceRepository.Get(ctx, namespace.GetByName, s.Id, params.Namespace)
	if err != nil {
		if err == user.ErrEntityNotFound {
			util.NewResp(w, r).Error(err).Status(http.StatusNotFound).Send()
		} else {
			util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		}
		return
	}

	u, err := c.userRepository.GetByLogin(ctx, params.Login, ns.Id)
	if err != nil {
		if err == user.ErrEntityNotFound {
			util.NewResp(w, r).Error(err).Status(http.StatusNotFound).Send()
		} else {
			util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		}
		return
	}

	userRoutes, err := c.acsRepository.GetUserRoutes(ctx, u.Id, ns.Id)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	isAccessAllowed := AccessAllowed(params.Route, params.RouteMethod, userRoutes)

	util.NewResp(w, r).Json(newCheckAccessResponse(isAccessAllowed, u.ExternalId)).Status(http.StatusOK).Send()
}

func getInputParams(r *http.Request) (*inputParams, error) {
	login, err := getLoginFromRequest(r)
	if err != nil {
		return nil, err
	}

	route, err := getRequestedRouteFromRequest(r)
	if err != nil {
		return nil, err
	}

	ns, err := getRequestedNamespaceFromRequest(r)
	if err != nil {
		return nil, err
	}

	s, err := getRequestedSystemFromRequest(r)
	if err != nil {
		return nil, err
	}

	routeMethod, err := getRequestedMethodFromRequest(r)
	if err != nil {
		return nil, err
	}
	return &inputParams{
		Login:       login,
		Route:       route,
		Namespace:   ns,
		System:      s,
		RouteMethod: routeMethod,
	}, nil
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

func getRequestedNamespaceFromRequest(r *http.Request) (string, error) {
	ns := r.Header.Get(requestedNamespaceHeaderName)
	if ns != "" {
		return ns, nil
	}

	return "", ErrRequestedNamespaceNotFoundInRequest
}

func getRequestedSystemFromRequest(r *http.Request) (string, error) {
	s := r.Header.Get(requestedSystemHeaderName)
	if s != "" {
		return s, nil
	}

	return "", ErrRequestedSystemNotFoundInRequest
}

func getLoginFromRequest(r *http.Request) (string, error) {
	login := r.Header.Get(loginHeaderName)
	if login != "" {
		return login, nil
	}

	return "", ErrLoginNotFoundInRequest
}

func AccessAllowed(requestedRoute, requestedMethod string, routes []UserRoute) bool {
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
