package usergroup

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/hardstylez72/bblog/ad/pkg/util"
	"net/http"
)

type Pair struct {
	GroupId int `json:"groupId" validate:"required"`
	UserId  int `json:"userId" validate:"required"`
}

type Repository interface {
	List(ctx context.Context, groupId int) ([]Group, error)
	ListNotInGroup(ctx context.Context, groupId int) ([]Group, error)
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
		return
	}

	var list []Group
	var err error

	if req.BelongToUser {
		list, err = c.rep.List(ctx, req.UserId)
	} else {
		list, err = c.rep.ListNotInGroup(ctx, req.UserId)
	}

	if err != nil {
		util.NewResp(w).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	render.JSON(w, r, newListResponse(list))
}

func (c *controller) delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req []Pair

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

	w.WriteHeader(http.StatusOK)
}

func (c *controller) Mount(r chi.Router) {
	r.Post("/v1/user/group/list", c.list)
	r.Post("/v1/user/group/create", c.create)
	r.Post("/v1/user/group/delete", c.delete)
}
