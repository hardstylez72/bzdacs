package user

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/hardstylez72/bzdacs/pkg/group"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/relations/usergroup"
	"github.com/hardstylez72/bzdacs/pkg/util"
	"github.com/pkg/errors"
	"net/http"
)

type Repository interface {
	GetById(ctx context.Context, id, namespaceId int) (*User, error)
	GetByLogin(ctx context.Context, login string, namespaceId int) (*User, error)
	Insert(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, id, namespaceId int) error
	List(ctx context.Context, filter filter) ([]User, int, error)
}

type controller struct {
	userRepository      Repository
	groupRepository     group.Repository
	userGroupRepository usergroup.Repository
	validator           *validator.Validate
}

func NewController(rep Repository, groupRepository group.Repository, userGroupRepository usergroup.Repository) *controller {
	return &controller{
		userRepository:      rep,
		groupRepository:     groupRepository,
		userGroupRepository: userGroupRepository,
		validator:           validator.New(),
	}
}

// @tags user
// @description Updates user
// @id user.update
// @accept application/json
// @param req body updateRequest true "request"
// @produce application/json
// @success 200 {object} getResponse
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/user/update [post]
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

	u, err := c.userRepository.Update(ctx, updateRequestConvert(&req))
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w, r).Status(http.StatusOK).Json(newInsertResponse(u)).Send()
}

// @tags user
// @description Creates user
// @id user.create
// @accept application/json
// @param req body insertRequest true "request"
// @produce application/json
// @success 200 {object} getResponse
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/user/create [post]
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

	u, err := c.userRepository.Insert(ctx, insertRequestConvert(&req))
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w, r).Status(http.StatusOK).Json(newInsertResponse(u)).Send()
}

// @tags user
// @description Gets user by login
// @id user.getByLogin
// @accept application/json
// @param req body getByLoginRequest true "request"
// @produce application/json
// @success 200 {object} getResponse
// @failure 400 {object} util.ResponseWithError
// @failure 404 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/user/getByLogin [post]
func (c *controller) getByLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req getByLoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	if err := c.validator.Struct(req); err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	user, err := c.userRepository.GetByLogin(ctx, req.Login, req.NamespaceId)
	if err != nil {
		if errors.Is(err, storage.EntityNotFound) {
			util.NewResp(w, r).Error(err).Status(http.StatusNotFound).Send()
		} else {
			util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		}
		return
	}

	util.NewResp(w, r).Json(user).Status(http.StatusOK).Send()
}

// @tags user
// @description Gets user by id
// @id user.getById
// @accept application/json
// @param req body getByIdRequest true "request"
// @produce application/json
// @success 200 {object} getResponse
// @failure 400 {object} util.ResponseWithError
// @failure 404 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/user/getById [post]
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

	user, err := c.userRepository.GetById(ctx, req.Id, req.NamespaceId)
	if err != nil {
		if errors.Is(err, storage.EntityNotFound) {
			util.NewResp(w, r).Error(err).Status(http.StatusNotFound).Send()
		} else {
			util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		}
		return
	}

	util.NewResp(w, r).Json(user).Status(http.StatusOK).Send()
}

// @tags user
// @description Gets list of users
// @id user.list
// @accept application/json
// @param req body listRequest true "request"
// @produce application/json
// @success 200 {object} listResponse
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/user/list [post]
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

	list, total, err := c.userRepository.List(ctx, req.Filter)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w, r).Json(newListResponse(list, total)).Status(http.StatusOK).Send()
}

// @tags user
// @description Deletes user by id
// @id user.delete
// @accept application/json
// @param req body deleteRequest true "request"
// @produce application/json
// @success 200
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/user/delete [post]
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

	err := c.userRepository.Delete(ctx, req.Id, req.NamespaceId)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w, r).Status(http.StatusOK).Send()
}

func (c *controller) Mount(private, public chi.Router) {
	private.Post("/v1/user/list", c.list)
	private.Post("/v1/user/getById", c.getById)
	private.Post("/v1/user/getByLogin", c.getByLogin)
	private.Post("/v1/user/create", c.create)
	private.Post("/v1/user/delete", c.delete)
	private.Post("/v1/user/update", c.update)

	public.Post("/v1/user/login", c.login)
	public.Post("/v1/user/session/get", c.session)
	public.Post("/v1/user/logout", c.logout)
}
