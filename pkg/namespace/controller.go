package namespace

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/hardstylez72/bzdacs/pkg/util"
	"net/http"
)

type Repository interface {
	GetListBySystemId(ctx context.Context, systemId int) ([]Namespace, error)
	Insert(ctx context.Context, namespace *NamespaceExt) (*Namespace, error)

	GetById(ctx context.Context, id int) (*Namespace, error)
	List(ctx context.Context) ([]Namespace, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, namespace *Namespace) (*Namespace, error)
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

func (c *controller) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	list, err := c.rep.List(ctx)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}
	util.NewResp(w, r).Status(http.StatusOK).Json(newListResponse(list)).Send()
}

func (c *controller) getBySystemId(w http.ResponseWriter, r *http.Request) {
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

	namespaces, err := c.rep.GetListBySystemId(ctx, req.Id)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w, r).Status(http.StatusOK).Json(namespaces).Send()
}

func (c *controller) getById(w http.ResponseWriter, r *http.Request) {
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

	group, err := c.rep.GetById(ctx, req.Id)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w, r).Status(http.StatusOK).Json(group).Send()
}

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
	r.Post("/v1/system/namespace/create", c.create)
	r.Post("/v1/system/namespace/list", c.getBySystemId)


	r.Post("/v1/namespace/list", c.list)
	r.Post("/v1/namespace/get", c.getById)

	r.Post("/v1/namespace/delete", c.delete)
	r.Post("/v1/namespace/update", c.update)

}
