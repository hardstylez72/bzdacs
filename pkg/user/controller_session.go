package user

import (
	"github.com/hardstylez72/bzdacs/pkg/util"
	"net/http"
	"time"
)

func (c *controller) session(w http.ResponseWriter, r *http.Request) {

	s, err := getSessionFromCookie(r.Cookies())
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusUnauthorized).Send()
		return
	}

	util.NewResp(w, r).Json(s).Status(http.StatusOK).Send()
}
func (c *controller) logout(w http.ResponseWriter, r *http.Request) {

	cookie := &http.Cookie{
		Name:     CookieSessionName,
		Value:    "",
		Path:     "/",
		Domain:   "",
		Expires:  time.Now().Add(-1 * time.Hour),
		SameSite: 2,
	}
	w.Header().Set("Set-Cookie", cookie.String())

	util.NewResp(w, r).Status(http.StatusOK).Send()
}
