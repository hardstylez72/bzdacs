package user

import (
	"encoding/json"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/util"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type registerRequest struct {
	Login    string `json:"login"  validate:"required"`
	Password string `json:"password"  validate:"required"`
} // @name sysUserRegisterRequest

// @tags sys-user
// @description Creates sys-user
// @id sys-user.register
// @accept application/json
// @param req body registerRequest true "request"
// @produce application/json
// @success 200
// @failure 400 {object} util.ResponseWithError
// @failure 422 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/sys-user/register [post]
func (c *controller) register(w http.ResponseWriter, r *http.Request) {
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

	u, err := c.userRepository.Insert(ctx, &SysUser{
		Login:    req.Login,
		Password: req.Password,
	})
	if err != nil {
		if errors.Is(err, storage.EntityAlreadyExist) {
			util.NewResp(w, r).Error(err).Status(http.StatusUnprocessableEntity).Send()
		} else {
			util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		}
		return
	}
	expiresInSeconds := viper.GetInt("user.sessionExpirationInSeconds")

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
