package group

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/hardstylez72/bblog/ad/pkg/util"
	"net/http"
)

type Repository interface {
	GetById(ctx context.Context, id int) (*Group, error)
	GetByCodeDb(ctx context.Context, code string) (*Group, error)
	List(ctx context.Context) ([]Group, error)
	Insert(ctx context.Context, group *Group) (*Group, error)
	InsertGroupBasedOnAnother(ctx context.Context, group *Group, groupBaseId int) (*Group, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, group *Group) (*Group, error)
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

	if err := c.validator.Struct(req); err != nil {
		util.NewResp(w).Error(err).Status(http.StatusBadRequest).Send()
		return
	}
	var group *Group
	var err error

	if req.BaseGroupId != 0 {
		group, err = c.rep.InsertGroupBasedOnAnother(ctx, insertRequestConvert(&req), req.BaseGroupId)
	} else {
		group, err = c.rep.Insert(ctx, insertRequestConvert(&req))
	}

	if err != nil {
		util.NewResp(w).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}
	util.NewResp(w).Status(http.StatusOK).Json(newInsertResponse(group)).Send()
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

	group, err := c.rep.Update(ctx, updateRequestConvert(&req))
	if err != nil {
		util.NewResp(w).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}
	util.NewResp(w).Status(http.StatusOK).Json(newUpdateResponse(group)).Send()
}

func (c *controller) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	list, err := c.rep.List(ctx)
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

	group, err := c.rep.GetById(ctx, req.Id)
	if err != nil {
		util.NewResp(w).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w).Status(http.StatusOK).Json(group).Send()
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
	r.Post("/v1/group/list", c.list)
	r.Post("/v1/group/get", c.getById)
	r.Post("/v1/group/create", c.create)
	r.Post("/v1/group/delete", c.delete)
	r.Post("/v1/group/update", c.update)

}
