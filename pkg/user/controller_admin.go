package user

import (
	"encoding/base64"
	"encoding/json"
	"github.com/hardstylez72/bzdacs/pkg/util"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

const (
	CookieSessionName = "session-token"
)

type Session struct {
	Token   string `json:"token"`
	IsAdmin bool   `json:"isAdmin"`
	Login   string `json:"login"`
}

func (c *controller) admin(w http.ResponseWriter, r *http.Request) {

	token, err := getTokenFromRequest(r)
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusUnauthorized).Send()
		return
	}

	isValid, login, err := validateToken(token)
	if err != nil || !isValid {
		util.NewResp(w, r).Error(err).Status(http.StatusUnauthorized).Send()
		return
	}

	err = setSessionInCookie(w, &Session{
		Token:   token,
		IsAdmin: true,
		Login:   login,
	})
	if err != nil {
		util.NewResp(w, r).Error(err).Status(http.StatusInternalServerError).Send()
		return
	}

	util.NewResp(w, r).Status(http.StatusOK).Send()
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
