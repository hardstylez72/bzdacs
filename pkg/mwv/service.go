package mwv

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type service struct {
	host       string
	httpClient *http.Client
}

type CheckAccessResponse struct {
	Login         string   `json:"login"`
	AccessAllowed bool     `json:"accessAllowed"`
	Groups        []string `json:"groups"`
}

var (
	ErrLoginIsEmpty = errors.New("input param login is empty")
)

func NewService(host string) *service {
	return &service{
		host:       host,
		httpClient: &http.Client{},
	}
}

func (s *service) CheckAccess(ctx context.Context, login, route, method string) (*CheckAccessResponse, error) {
	if login == "" {
		return nil, ErrLoginIsEmpty
	}

	const r = "/api/v1/user/access/check"
	url := s.host + r

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error while building http request")
	}

	req.Header.Set("login", login)
	req.Header.Set("route", route)
	req.Header.Set("method", method)

	res, err := s.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error while doing http request")
	}

	defer func() {
		_ = res.Body.Close()
	}()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error while reading http request body")
	}

	var resBodyParsed CheckAccessResponse

	err = json.Unmarshal(resBody, &resBodyParsed)
	if err != nil {
		return nil, errors.New("error while decoding http request: " + err.Error())
	}

	return &resBodyParsed, nil

}
