package user

import (
	"github.com/hardstylez72/bzdacs/pkg/util"
	"net/http"
)

func (c *controller) session(w http.ResponseWriter, r *http.Request) {

	s, err := getSessionFromCookie(r.Cookies())
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusUnauthorized).Send()
		return
	}

	util.NewResp(w, r).Json(s).Status(http.StatusOK).Send()
}
