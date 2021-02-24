package usergroup

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
	UserId  int `json:"userId" validate:"required"`
} // @name userGroupPair

type Repository interface {
	List(ctx context.Context, filter Filter) ([]Group, int, error)
	Insert(ctx context.Context, params []Pair) ([]Group, error)
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

// @tags user-group
// @description Creates user-group relations
// @id user-group.create
// @accept application/json
// @param req body insertRequest true "request"
// @produce application/json
// @success 200	{object} insertResponse
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/user/group/create [post]
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

	util.NewResp(w, r).Json(routes).Status(http.StatusOK).Send()
}

// @tags user-group
// @description Gets list of user-group relations
// @id user-group.list
// @accept application/json
// @param req body listRequest true "request"
// @produce application/json
// @success 200	{object} listResponse
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/user/group/list [post]
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

	util.NewResp(w, r).Json(newListResponse(list, total)).Status(http.StatusOK).Send()
}

// @tags user-group
// @description Deletes user-group relations
// @id user-group.delete
// @accept application/json
// @param req body deleteRequest true "request"
// @produce application/json
// @success 200
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/user/group/delete [post]
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

	w.WriteHeader(http.StatusOK)
}

func (c *controller) Mount(r chi.Router) {
	r.Post("/v1/user/group/list", c.list)
	r.Post("/v1/user/group/create", c.create)
	r.Post("/v1/user/group/delete", c.delete)
}
