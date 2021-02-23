package group

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/util"
	"github.com/pkg/errors"
	"net/http"
)

type Repository interface {
	Insert(ctx context.Context, group *Group) (*Group, error)
	InsertGroupWithCopy(ctx context.Context, group *Group, groupId int) (*Group, error)
	Update(ctx context.Context, group *Group) (*Group, error)
	Delete(ctx context.Context, id, namespaceId int) error
	GetById(ctx context.Context, id, namespaceId int) (*Group, error)
	GetByCode(ctx context.Context, code string, namespaceId int) (*Group, error)
	List(ctx context.Context, filter filter) ([]Group, int, error)
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

// @tags group
// @description Creates group
// @id group.create
// @accept application/json
// @param req body insertRequest true "request"
// @produce application/json
// @success 200 {object} getResponse
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/group/create [post]
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
	var group *Group
	var err error

	if req.BaseGroupId != 0 {
		group, err = c.rep.InsertGroupWithCopy(ctx, insertRequestConvert(&req), req.BaseGroupId)
	} else {
		group, err = c.rep.Insert(ctx, insertRequestConvert(&req))
	}

	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}
	util.NewResp(w, r).Status(http.StatusOK).Json(newInsertResponse(group)).Send()
}

// @tags group
// @description Updates group
// @id group.update
// @accept application/json
// @param req body updateRequest true "request"
// @produce application/json
// @success 200 {object} getResponse
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/group/update [post]
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

// @tags group
// @description List groups
// @id group.list
// @accept application/json
// @param req body listRequest true "request"
// @produce application/json
// @success 200 {object} listResponse
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/group/list [post]
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

// @tags group
// @description Gets group by code
// @id group.getByCode
// @accept application/json
// @param req body getByCodeRequest true "request"
// @produce application/json
// @success 200 {object} getResponse
// @failure 400 {object} util.ResponseWithError
// @failure 404 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/group/getByCode [post]
func (c *controller) getByCode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req getByCodeRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	if err := c.validator.Struct(req); err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	group, err := c.rep.GetByCode(ctx, req.Code, req.NamespaceId)
	if err != nil {
		if errors.Is(err, storage.EntityNotFound) {
			util.NewResp(w, r).Error(err).Status(http.StatusNotFound).Send()
		} else {
			util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		}
		return
	}

	util.NewResp(w, r).Status(http.StatusOK).Json(newGetResponse(group)).Send()
}

// @tags group
// @description Gets group by id
// @id group.getById
// @accept application/json
// @param req body getByIdRequest true "request"
// @produce application/json
// @success 200 {object} getResponse
// @failure 400 {object} util.ResponseWithError
// @failure 404 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/group/getById [post]
func (c *controller) getById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req getByIdRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	if err := c.validator.Struct(req); err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	group, err := c.rep.GetById(ctx, req.Id, req.NamespaceId)
	if err != nil {
		if errors.Is(err, storage.EntityNotFound) {
			util.NewResp(w, r).Error(err).Status(http.StatusNotFound).Send()
		} else {
			util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		}
		return
	}

	util.NewResp(w, r).Status(http.StatusOK).Json(newGetResponse(group)).Send()
}

// @tags group
// @description Deletes group
// @id group.delete
// @accept application/json
// @param req body deleteRequest true "request"
// @produce application/json
// @success 200
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/group/delete [post]
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

	err := c.rep.Delete(ctx, req.Id, req.NamespaceId)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w, r).Status(http.StatusOK).Send()
}

func (c *controller) Mount(r chi.Router) {
	r.Post("/v1/group/list", c.list)
	r.Post("/v1/group/getById", c.getById)
	r.Post("/v1/group/getByCode", c.getByCode)
	r.Post("/v1/group/create", c.create)
	r.Post("/v1/group/delete", c.delete)
	r.Post("/v1/group/update", c.update)

}
