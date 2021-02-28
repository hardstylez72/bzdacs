package user // @name sys-user

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/hardstylez72/bzdacs/internal/session"
	"github.com/hardstylez72/bzdacs/pkg/group"
	"github.com/hardstylez72/bzdacs/pkg/relations/usergroup"
	"github.com/hardstylez72/bzdacs/pkg/util"
	"net/http"
)

type Repository interface {
	GetById(ctx context.Context, id, namespaceId int) (*User, error)
	GetByParams(ctx context.Context, login, password string) (*User, error)
	Insert(ctx context.Context, user *User) (*User, error)
}

type controller struct {
	userRepository      Repository
	groupRepository     group.Repository
	userGroupRepository usergroup.Repository
	validator           *validator.Validate
	sessionService      *session.Service
}

func NewController(rep Repository, groupRepository group.Repository, userGroupRepository usergroup.Repository) *controller {
	return &controller{
		userRepository:      rep,
		groupRepository:     groupRepository,
		userGroupRepository: userGroupRepository,
		validator:           validator.New(),
		sessionService:      session.NewSessionService(),
	}
}

// @tags sys-user
// @description Logout sys-user
// @id sys-user.logout
// @accept application/json
// @produce application/json
// @success 200
// @router /v1/sys-user/logout [post]
func (c *controller) logout(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	util.NewResp(w, r).Status(http.StatusOK).Send()
}

func (c *controller) Mount(private, public chi.Router) {
	public.Post("/v1/sys-user/register", c.register)
	public.Post("/v1/sys-user/login", c.login)
	public.Post("/v1/sys-user/session", c.session)
	private.Post("/v1/sys-user/logout", c.logout)
}
