package user

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/hardstylez72/bzdacs/pkg/util"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type sessionResponse struct {
	Login string
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

	s, err := getSessionFromCookie(r.Cookies())
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusUnauthorized).Send()
		return
	}

	session, err := c.sessionService.GetByToken(s.Token)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusUnauthorized).Send()
		return
	}

	util.NewResp(w, r).Json(&sessionResponse{session.Login}).Status(http.StatusOK).Send()
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

func clearSession(w http.ResponseWriter) {
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

func getTokenFromRequest(r *http.Request) (string, error) {
	var token string

	token = r.Header.Get("token")
	if token != "" {
		return token, nil
	}

	session, _ := getSessionFromCookie(r.Cookies())
	if session != nil {
		return session.Token, nil
	}

	return "", errors.New("token not found")
}

func validateToken(token string) (bool, string, error) {
	var (
		login    = viper.GetString("user.login")
		password = viper.GetString("user.password")
	)
	tokenB, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return false, "", err
	}
	if string(tokenB) != login+":"+password {
		return false, "", errors.New("wrong credentials")
	}

	return true, login, nil
}

func parseToken(token string) (string, error) {
	isValid, login, err := validateToken(token)
	if isValid {
		return login, nil
	}
	return "", err
}

func ExtractLogin(r *http.Request) (string, error) {
	token, err := getTokenFromRequest(r)
	if err != nil {
		return "", err
	}

	login, err := parseToken(token)
	if err != nil {
		return "", err
	}
	return login, nil
}
