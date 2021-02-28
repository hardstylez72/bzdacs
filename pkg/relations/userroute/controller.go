package userroute

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/hardstylez72/bzdacs/pkg/util"
	"net/http"
)

type PairToDelete struct {
	RouteId int `json:"routeId" validate:"required"`
	UserId  int `json:"userId" validate:"required"`
} // @name userRoutePairToDelete

type Pair struct {
	RouteId    int  `json:"routeId" validate:"required"`
	UserId     int  `json:"userId" validate:"required"`
	IsExcluded bool `json:"isExcluded"`
} // @name userRoutePair

type Repository interface {
	List(ctx context.Context, filter Filter) ([]RouteWithGroups, int, error)
	Insert(ctx context.Context, params []Pair) ([]Route, error)
	Delete(ctx context.Context, params []PairToDelete) error
	Update(ctx context.Context, params Pair) (*Route, error)
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

// @tags user-route
// @description Creates pairs of user-route relations
// @id user-route.create
// @accept application/json
// @param req body insertRequest true "request"
// @produce application/json
// @success 200	{object} insertResponse
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/user/route/create [post]
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

	util.NewResp(w, r).Json(routes).Send()
}

// @tags user-route
// @description Gets list of user-route relations
// @id user-route.list
// @accept application/json
// @param req body listRequest true "request"
// @produce application/json
// @success 200	{object} listResponse
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/user/route/list [post]
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

	render.JSON(w, r, newListResponse(list, total))
}

// @tags user-route
// @description Deletes list of user-route relations
// @id user-route.delete
// @accept application/json
// @param req body deleteRequest true "request"
// @produce application/json
// @success 200
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/user/route/delete [post]
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

// @tags user-route
// @description Updates pair of user-route relations
// @id user-route.update
// @accept application/json
// @param req body Pair true "request"
// @produce application/json
// @success 200	{object} updateResponse
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/user/route/update [post]
func (c *controller) update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req Pair

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	if err := c.validator.Struct(req); err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	route, err := c.rep.Update(ctx, req)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w, r).Status(http.StatusOK).Json(route).Send()
}

func (c *controller) Mount(r chi.Router) {
	r.Post("/v1/user/route/list", c.list)
	r.Post("/v1/user/route/create", c.create)
	r.Post("/v1/user/route/delete", c.delete)
	r.Post("/v1/user/route/update", c.update)
}
