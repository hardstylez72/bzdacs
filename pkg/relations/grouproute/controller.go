package grouproute

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/hardstylez72/bzdacs/pkg/util"
	"net/http"
)

type Pair struct {
	GroupId int `json:"groupId" validate:"required"`
	RouteId int `json:"routeId" validate:"required"`
} // @name groupRoutePair

type Repository interface {
	List(ctx context.Context, filter Filter) ([]Route, int, error)
	Insert(ctx context.Context, params []Pair) ([]Route, error)
	Delete(ctx context.Context, params []Pair) error
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

// @tags group-route
// @description Creates group-routes relations
// @id group-routes.create
// @accept application/json
// @param req body insertRequest true "request"
// @produce application/json
// @success 200	{object} insertResponse
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/group/route/create [post]
func (c *controller) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req insertRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	if err := c.validator.Struct(req); err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	routes, err := c.rep.Insert(ctx, req.Pairs)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}
	util.NewResp(w, r).Status(http.StatusOK).Json(routes).Send()
}

// @tags group-route
// @description Gets list of group-routes relations
// @id group-routes.list
// @accept application/json
// @param req body listRequest true "request"
// @produce application/json
// @success 200	{object} listResponse
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/group/route/list [post]
func (c *controller) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req listRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	if err := c.validator.Struct(req); err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	list, total, err := c.rep.List(ctx, req.Filter)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}
	util.NewResp(w, r).Status(http.StatusOK).Json(newListResponse(list, total)).Send()
}

// @tags group-route
// @description Deletes group-routes relations
// @id group-routes.delete
// @accept application/json
// @param req body deleteRequest true "request"
// @produce application/json
// @success 200
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/group/route/delete [post]
func (c *controller) delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req deleteRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	if err := c.validator.Struct(req); err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	err := c.rep.Delete(ctx, req.Pairs)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w, r).Status(http.StatusOK).Send()
}

func (c *controller) Mount(r chi.Router) {
	r.Post("/v1/group/route/list", c.list)
	r.Post("/v1/group/route/create", c.create)
	r.Post("/v1/group/route/delete", c.delete)
}
