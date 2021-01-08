package route

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/hardstylez72/bblog/ad/pkg/util"
	"net/http"
)

type Repository interface {
	List(ctx context.Context, f filter) ([]RouteWithTags, error)
	GetById(ctx context.Context, id int) (*RouteWithTags, error)
	GetByMethodAndRoute(ctx context.Context, route, method string) (*RouteWithTags, error)
	Insert(ctx context.Context, group *Route) (*Route, error)
	InsertWithTags(ctx context.Context, route *Route, tagNames []string) (*RouteWithTags, error)
	UpdateWithTags(ctx context.Context, route *Route, tagNames []string) (*RouteWithTags, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, group *Route) (*Route, error)
}

type controller struct {
	rep       Repository
	validator *validator.Validate
}

func NewController(rep Repository) *controller {
	return &controller{rep: rep, validator: validator.New()}
}

func (c *controller) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req insertRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.NewResp(w).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	if err := c.validator.Struct(req); err != nil {
		util.NewResp(w).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	route, err := c.rep.InsertWithTags(ctx, insertRequestConvert(&req), req.Tags)
	if err != nil {
		util.NewResp(w).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w).Status(http.StatusOK).Json(newInsertResponse(route)).Send()
}

func (c *controller) update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req updateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.NewResp(w).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	if err := c.validator.Struct(req); err != nil {
		util.NewResp(w).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	route, err := c.rep.UpdateWithTags(ctx, updateRequestConvert(&req), req.Tags)
	if err != nil {
		util.NewResp(w).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w).Status(http.StatusOK).Json(route).Send()
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
		return
	}

	list, err := c.rep.List(ctx, req.Filter)
	if err != nil {
		util.NewResp(w).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w).Status(http.StatusOK).Json(newListResponse(list)).Send()
}

func (c *controller) getById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req getRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.NewResp(w).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	if err := c.validator.Struct(req); err != nil {
		util.NewResp(w).Error(err).Status(http.StatusBadRequest).Send()
		return
	}
	list, err := c.rep.GetById(ctx, req.Id)
	if err != nil {
		util.NewResp(w).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w).Status(http.StatusOK).Json(list).Send()
}

func (c *controller) delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req deleteRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.NewResp(w).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	if err := c.validator.Struct(req); err != nil {
		util.NewResp(w).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	err := c.rep.Delete(ctx, req.Id)
	if err != nil {
		util.NewResp(w).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w).Status(http.StatusOK).Send()
}

func (c *controller) Mount(r chi.Router) {
	r.Post("/v1/route/list", c.list)
	r.Post("/v1/route/get", c.getById)
	r.Post("/v1/route/create", c.create)
	r.Post("/v1/route/delete", c.delete)
	r.Post("/v1/route/update", c.update)
}
