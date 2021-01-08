package user

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/hardstylez72/bblog/ad/pkg/util"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type Repository interface {
	GetById(ctx context.Context, id int) (*User, error)
	GetByExternalId(ctx context.Context, id string) (*User, error)
	List(ctx context.Context) ([]User, error)
	Insert(ctx context.Context, group *User) (*User, error)
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
		util.NewResp(w).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	if err := c.validator.Struct(req); err != nil {
		util.NewResp(w).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	group, err := c.rep.Insert(ctx, insertRequestConvert(&req))
	if err != nil {
		util.NewResp(w).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w).Status(http.StatusOK).Json(newInsertResponse(group)).Send()
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

	user, err := c.rep.GetById(ctx, req.Id)
	if err != nil {
		util.NewResp(w).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w).Json(user).Status(http.StatusOK).Send()
}

func (c *controller) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	list, err := c.rep.List(ctx)
	if err != nil {
		util.NewResp(w).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w).Json(newListResponse(list)).Status(http.StatusOK).Send()
}

func (c *controller) login(w http.ResponseWriter, r *http.Request) {

	token, err := getTokenFromRequest(r)
	if err != nil {
		util.NewResp(w).Error(err).Status(http.StatusUnauthorized).Send()
		return
	}

	isValid, err := validateToken(token)
	if err != nil || !isValid {
		util.NewResp(w).Error(err).Status(http.StatusUnauthorized).Send()
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "bzdacs-token",
		Value:    token,
		Path:     "/",
		Domain:   "",
		Expires:  time.Now().Add(time.Hour * 24),
		SameSite: 2,
	})

	util.NewResp(w).Status(http.StatusOK).Send()
}

func getTokenFromRequest(r *http.Request) (string, error) {
	var token string

	for _, cookie := range r.Cookies() {
		if cookie.Name == "bzdacs-token" {
			token = cookie.Value
		}
	}

	if token != "" {
		return token, nil
	}

	token = r.Header.Get("token")
	if token != "" {
		return token, nil
	}

	return "", errors.New("token not found")
}

func validateToken(token string) (bool, error) {
	var (
		login    = viper.GetString("user.login")
		password = viper.GetString("user.password")
	)
	tokenB, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return false, err
	}
	if string(tokenB) != login+":"+password {
		return false, errors.New("wrong credentials")
	}

	return true, nil
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
	r.Post("/v1/user/list", c.list)
	r.Post("/v1/user/get", c.getById)
	r.Post("/v1/user/create", c.create)
	r.Post("/v1/user/delete", c.delete)
	r.Post("/v1/user/login", c.login)

}
