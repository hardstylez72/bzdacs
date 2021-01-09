package user

import (
	"errors"
	"github.com/hardstylez72/bzdacs/pkg/usergroup"
	"github.com/hardstylez72/bzdacs/pkg/util"
	"github.com/spf13/viper"
	"net/http"
)

func (c *controller) guest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	login, err := GetLoginFromRequest(r)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusUnauthorized).Send()
		return
	}

	if login == viper.GetString("user.login") {
		util.NewResp(w, r).Error(err).Status(http.StatusUnauthorized).Send()
		return
	}

	if login == "" {
		util.NewResp(w, r).Error(err).Status(http.StatusBadRequest).Send()
		return
	}

	u, err := c.userRepository.GetByExternalId(ctx, login)
	if err != nil {
		println(err == util.ErrEntityNotFound)
		if err == ErrEntityNotFound {
			var insertUserErr error
			u, insertUserErr = c.userRepository.Insert(ctx, &User{ExternalId: login})
			if insertUserErr != nil {
				util.NewResp(w, r).Error(insertUserErr).Status(http.StatusInternalServerError).Send()
				return
			}
			g, getGuestGroupErr := c.groupRepository.GetByCode(ctx, "sys_guest")
			if getGuestGroupErr != nil {
				util.NewResp(w, r).Error(getGuestGroupErr).Status(http.StatusInternalServerError).Send()
				return
			}

			_, insertUserGroupPair := c.userGroupRepository.InsertMany(ctx, []usergroup.Pair{{
				GroupId: g.Id,
				UserId:  u.Id,
			}})
			if insertUserGroupPair != nil {
				util.NewResp(w, r).Error(insertUserGroupPair).Status(http.StatusInternalServerError).Send()
				return
			}

		} else {
			util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
			return
		}
	}

	err = setSessionInCookie(w, &Session{
		Token:   "",
		IsAdmin: false,
		Login:   u.ExternalId,
	})
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w, r).Status(http.StatusOK).Send()
}

func GetLoginFromRequest(r *http.Request) (string, error) {

	login := r.Header.Get("login")
	if login != "" {
		return login, nil
	}

	session, _ := getSessionFromCookie(r.Cookies())
	if session != nil {
		return session.Login, nil
	}

	return "", errors.New("param login is not found in request body")
}
