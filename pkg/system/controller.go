package system

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/util"
	"net/http"
)

type GetKind string
type ListKind string

type Repository interface {
	Insert(ctx context.Context, system *System) (*System, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, system *System) (*System, error)
	Get(ctx context.Context, kind GetKind, arg ...interface{}) (*System, error)
	List(ctx context.Context, kind ListKind, arg ...interface{}) ([]System, error)
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

// @tags system
// @description Creates system
// @id system.create
// @accept application/json
// @param req body insertRequest true "request"
// @produce application/json
// @success 200 {object} getResponse
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/system/create [post]
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

	group, err := c.rep.Insert(ctx, insertRequestConvert(&req))
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}
	util.NewResp(w, r).Status(http.StatusOK).Json(newInsertResponse(group)).Send()
}

// @tags system
// @description Updates system
// @id system.update
// @accept application/json
// @param req body updateRequest true "request"
// @produce application/json
// @success 200 {object} getResponse
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/system/update [post]
func (c *controller) update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req updateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	if err := c.validator.Struct(req); err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	group, err := c.rep.Update(ctx, updateRequestConvert(&req))
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}
	util.NewResp(w, r).Status(http.StatusOK).Json(newUpdateResponse(group)).Send()
}

// @tags system
// @description Gets system list
// @id system.list
// @accept application/json
// @produce application/json
// @success 200 {array} getResponse
// @failure 500 {object} util.ResponseWithError
// @router /v1/system/list [post]
func (c *controller) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	list, err := c.rep.List(ctx, ListNoFilter)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}
	util.NewResp(w, r).Status(http.StatusOK).Json(newListResponse(list)).Send()
}

// @tags system
// @description Gets system by multiple params
// @id system.get
// @accept application/json
// @param req body getRequest true "request"
// @produce application/json
// @success 200 {object} getResponse
// @failure 400 {object} util.ResponseWithError
// @failure 404 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/system/get [post]
func (c *controller) get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req getRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	if err := c.validator.Struct(req); err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
		return
	}
	if req.Id <= 0 && req.Name == "" {
		err := errors.New("system id or name must be set")
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	var system *System
	var err error
	if req.Id != 0 {
		system, err = c.rep.Get(ctx, GetById, req.Id)
	} else {
		system, err = c.rep.Get(ctx, GetByName, req.Name)
	}

	if err != nil {
		if errors.Is(err, storage.EntityNotFound) {
			util.NewResp(w, r).Error(err).Status(http.StatusNotFound).Send()
		} else {
			util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		}

		return
	}

	util.NewResp(w, r).Status(http.StatusOK).Json(newGetResponse(system)).Send()
}

// @tags system
// @description Deletes system by id
// @id system.delete
// @accept application/json
// @param req body deleteRequest true "request"
// @produce application/json
// @success 200
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/system/delete [post]
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

	err := c.rep.Delete(ctx, req.Id)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w, r).Status(http.StatusOK).Send()
}

func (c *controller) Mount(r chi.Router) {
	r.Post("/v1/system/list", c.list)
	r.Post("/v1/system/get", c.get)
	r.Post("/v1/system/create", c.create)
	r.Post("/v1/system/delete", c.delete)
	r.Post("/v1/system/update", c.update)
}
