package user

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/hardstylez72/bzdacs/pkg/util"
	"net/http"
)

type Repository interface {
	GetById(ctx context.Context, id int) (*User, error)
	GetByExternalId(ctx context.Context, id string) (*User, error)
	List(ctx context.Context) ([]User, error)
	Insert(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, id int) error
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

	user, err := c.rep.GetById(ctx, req.Id)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w, r).Json(user).Status(http.StatusOK).Send()
}

func (c *controller) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	list, err := c.rep.List(ctx)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w, r).Json(newListResponse(list)).Status(http.StatusOK).Send()
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

func (c *controller) Mount(private, public chi.Router) {
	private.Post("/v1/user/list", c.list)
	private.Post("/v1/user/get", c.getById)
	private.Post("/v1/user/create", c.create)
	private.Post("/v1/user/delete", c.delete)
	public.Post("/v1/user/admin/login", c.admin)
	public.Post("/v1/user/guest/login", c.guest)
	public.Post("/v1/user/session/get", c.session)
}
