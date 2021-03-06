package user

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/hardstylez72/bzdacs/internal/session"
	"github.com/hardstylez72/bzdacs/pkg/util"
	"net/http"
	"time"
)

type sessionResponse struct {
	Login string `json:"login" validate:"required"`
}

// @tags sys-user
// @description Gets sys-user session
// @id sys-user.session
// @accept application/json
// @produce application/json
// @success 200 {object} sessionResponse
// @failure 401 {object} util.ResponseWithError
// @failure 500 {object} util.ResponseWithError
// @router /v1/sys-user/session [post]
func (c *controller) session(w http.ResponseWriter, r *http.Request) {

	session, err := getSession(r, c.sessionService)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusUnauthorized).Send()
		return
	}

	util.NewResp(w, r).Json(&sessionResponse{session.Login}).Status(http.StatusOK).Send()
}

func getSession(r *http.Request, service *session.Service) (*session.Session, error) {
	s, err := getSessionFromCookie(r.Cookies())
	if err != nil {
		return nil, err
	}

	session, err := service.GetByToken(s.Token)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func getSessionFromCookie(cookies []*http.Cookie) (*Session, error) {
	var session Session
	var cookieValue string
	for _, cookie := range cookies {
		if cookie.Name == CookieSessionName {
			cookieValue = cookie.Value
		}
	}

	if cookieValue == "" {
		return nil, errors.New("session cookies not found")
	}

	s, err := base64.StdEncoding.DecodeString(cookieValue)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(s, &session)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func clearCookieSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     CookieSessionName,
		Value:    "",
		Path:     "/",
		Domain:   "",
		Expires:  time.Now().Add(-1 * time.Hour),
		SameSite: 2,
	}
	w.Header().Set("Set-Cookie", cookie.String())
}

func setSessionInCookie(w http.ResponseWriter, session *Session) error {

	jsonSession, err := json.Marshal(session)
	if err != nil {
		return err
	}

	s := base64.StdEncoding.EncodeToString(jsonSession)

	c := &http.Cookie{
		Name:     CookieSessionName,
		Value:    s,
		Path:     "/",
		Domain:   "",
		Expires:  time.Now().Add(time.Hour * 24),
		SameSite: 2,
	}
	w.Header().Set("Set-Cookie", c.String())

	return nil
}
