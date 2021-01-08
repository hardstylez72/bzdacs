package userroute

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/hardstylez72/bblog/ad/pkg/util"
	"net/http"
)

type params struct {
	RouteId    int  `json:"routeId" validate:"required"`
	UserId     int  `json:"userId" validate:"required"`
	IsExcluded bool `json:"isExcluded"`
}

type Repository interface {
	RoutesBelongUser(ctx context.Context, userId int) ([]RouteWithGroups, error)
	RoutesNotBelongUser(ctx context.Context, userId int) ([]RouteWithGroups, error)
	Insert(ctx context.Context, params []params) ([]Route, error)
	Delete(ctx context.Context, params []params) error
	Update(ctx context.Context, params params) (*Route, error)
}

type controller struct {
	rep       Repository
	validator *validator.Validate
}

func NewController(rep Repository) *controller {
	return &controller{
		rep:       rep,
		validator: validator.New(),
	}
}

func (c *controller) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req insertRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.NewResp(w).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	routes, err := c.rep.Insert(ctx, insertRequestConvert(req))
	if err != nil {
		util.NewResp(w).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w).Json(routes).Send()
}

func (c *controller) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req listRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.NewResp(w).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	if err := c.validator.Struct(req); err != nil {
		util.NewResp(w).Error(err).Status(http.StatusBadRequest).Send()
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	var list []RouteWithGroups
	var err error

	if req.BelongToUser {
		list, err = c.rep.RoutesBelongUser(ctx, req.UserId)
	} else {
		list, err = c.rep.RoutesNotBelongUser(ctx, req.UserId)
	}

	if err != nil {
		util.NewResp(w).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	render.JSON(w, r, newListResponse(list))
}

func (c *controller) delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req []params

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.NewResp(w).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	rr := deleteRequest{Params: req}

	if err := c.validator.Struct(rr); err != nil {
		util.NewResp(w).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	err := c.rep.Delete(ctx, req)
	if err != nil {
		util.NewResp(w).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w).Status(http.StatusOK).Send()
}

func (c *controller) update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req params

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.NewResp(w).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	if err := c.validator.Struct(req); err != nil {
		util.NewResp(w).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	route, err := c.rep.Update(ctx, req)
	if err != nil {
		util.NewResp(w).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w).Status(http.StatusOK).Json(route).Send()
}

func (c *controller) Mount(r chi.Router) {
	r.Post("/v1/user/route/list", c.list)
	r.Post("/v1/user/route/create", c.create)
	r.Post("/v1/user/route/delete", c.delete)
	r.Post("/v1/user/route/update", c.update)
}
