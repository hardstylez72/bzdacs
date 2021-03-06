package user

import (
	"encoding/json"
	"github.com/hardstylez72/bzdacs/config"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/util"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

const (
	CookieSessionName = "session-token"
)

type Session struct {
	Token string `json:"token"`
}

type loginRequest struct {
	Login    string `json:"login"  validate:"required"`
	Password string `json:"password"  validate:"required"`
} // @name sysUserLoginRequest

// @tags sys-user
// @description Login sys-user
// @id sys-user.login
// @accept application/json
// @param req body loginRequest true "request"
// @produce application/json
// @success 200
// @failure 400 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/sys-user/login [post]
func (c *controller) login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	if err := c.validator.Struct(req); err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	// todo: add validation user already logged in
	_, err := getSession(r, c.sessionService)
	if err == nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	u, err := c.userRepository.GetByParams(ctx, req.Login, req.Password)
	if err != nil {
		if errors.Is(err, storage.EntityNotFound) {
			util.NewResp(w, r).Error(err).Status(http.StatusNotFound).Send()
		} else {
			util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		}
		return
	}
	expiresInSeconds := config.GetInternal().SessionExpirationInSeconds
	session := c.sessionService.GenerateSession(u.Login, req.Password, time.Now().Add(time.Duration(expiresInSeconds)*time.Second))
	token, err := c.sessionService.Set(session)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	err = setSessionInCookie(w, &Session{Token: token})
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w, r).Status(http.StatusOK).Send()
}
