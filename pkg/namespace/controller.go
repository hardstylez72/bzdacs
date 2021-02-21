package namespace

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

type Repository interface {
	List(ctx context.Context, systemId int) ([]Namespace, error)
	Insert(ctx context.Context, namespace *Namespace, systemId int) (*Namespace, error)
	Delete(ctx context.Context, systemId, namespaceId int) error
	Update(ctx context.Context, namespace *Namespace) (*Namespace, error)
	Get(ctx context.Context, systemId, id int, name string) (*Namespace, error)
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

// @tags namespace
// @description Creates namespace
// @id namespace.create
// @accept application/json
// @param req body insertRequest true "request"
// @produce application/json
// @success 200 {object} insertResponse
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/namespace/create [post]
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

	ns, sysId := insertRequestConvert(&req)
	namespace, err := c.rep.Insert(ctx, ns, sysId)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}
	util.NewResp(w, r).Status(http.StatusOK).Json(newInsertResponse(namespace)).Send()
}

// @tags namespace
// @description Updates namespace
// @id namespace.update
// @accept application/json
// @param req body updateRequest true "request"
// @produce application/json
// @success 200 {object} updateResponse
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/namespace/update [post]
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

// @tags namespace
// @description Gets namespace with specified params
// @id namespace.get
// @accept application/json
// @param req body getRequest true "request"
// @produce application/json
// @success 200 {object} getResponse
// @failure 400 {object} util.ResponseWithError
// @failure 404 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/namespace/get [post]
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
		err := errors.New("namespace id or name+systemId must be set")
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	if req.Id <= 0 {
		if req.SystemId <= 0 && req.Name == "" {
			err := errors.New("name and systemId must be set")
			util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
			return
		}
	}

	namespace, err := c.rep.Get(ctx, req.SystemId, req.Id, req.Name)
	if err != nil {
		if err == storage.EntityNotFound {
			util.NewResp(w, r).Error(err).Status(http.StatusNotFound).Send()
		} else {
			util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		}
		return
	}

	util.NewResp(w, r).Status(http.StatusOK).Json(newGetResponse(namespace)).Send()
}

// @tags namespace
// @description List of namespaces belong to selected system
// @id namespace.list
// @accept application/json
// @param req body listRequest true "request"
// @produce application/json
// @success 200 {object} listResponse
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/namespace/list [post]
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

	namespaces, err := c.rep.List(ctx, req.Id)
	if err != nil {
		if err == storage.EntityNotFound {
			util.NewResp(w, r).Error(err).Status(http.StatusNotFound).Send()
		} else {
			util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		}
		return
	}

	util.NewResp(w, r).Status(http.StatusOK).Json(newListResponse(namespaces)).Send()
}

// @tags namespace
// @description Deletes selected namespace
// @id namespace.delete
// @accept application/json
// @param req body deleteRequest true "request"
// @produce application/json
// @success 200
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/namespace/delete [post]
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

	err := c.rep.Delete(ctx, req.SystemId, req.NamespaceId)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w, r).Status(http.StatusOK).Send()
}

func (c *controller) Mount(r chi.Router) {
	r.Post("/v1/namespace/create", c.create)
	r.Post("/v1/namespace/list", c.list)
	r.Post("/v1/namespace/delete", c.delete)
	r.Post("/v1/namespace/update", c.update)
	r.Post("/v1/namespace/get", c.get)

}
