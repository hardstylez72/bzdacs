package acsmw

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
	Login         string `json:"login"`
	AccessAllowed bool   `json:"accessAllowed"`
}

var (
	ErrLoginIsEmpty       = errors.New("input param login is empty")
	ErrRouteIsEmpty       = errors.New("input param route is empty")
	ErrRouteMethodIsEmpty = errors.New("input param route method is empty")
	ErrSystemIsEmpty      = errors.New("input param system is empty")
	ErrNamespaceIsEmpty   = errors.New("input param namespace is empty")
)

func NewService(host string) *service {
	return &service{
		host:       host,
		httpClient: &http.Client{},
	}
}

func (s *service) CheckAccess(ctx context.Context, params *InputParams) (*CheckAccessResponse, error) {
	if params.Login == "" {
		return nil, ErrLoginIsEmpty
	}
	if params.Route == "" {
		return nil, ErrRouteIsEmpty
	}

	if params.RouteMethod == "" {
		return nil, ErrRouteMethodIsEmpty
	}

	if params.System == "" {
		return nil, ErrSystemIsEmpty
	}

	if params.Namespace == "" {
		return nil, ErrNamespaceIsEmpty
	}

	const r = "/api/v1/user/access/check"
	url := s.host + r

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error while building http request")
	}

	req.Header.Set("login", params.Login)
	req.Header.Set("route", params.Route)
	req.Header.Set("method", params.RouteMethod)
	req.Header.Set("namespace", params.Namespace)
	req.Header.Set("system", params.System)

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
