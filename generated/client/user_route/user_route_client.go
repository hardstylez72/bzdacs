// Code generated by go-swagger; DO NOT EDIT.

package user_route

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user route API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user route API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	UserRouteCreate(params *UserRouteCreateParams, opts ...ClientOption) (*UserRouteCreateOK, error)

	UserRouteDelete(params *UserRouteDeleteParams, opts ...ClientOption) (*UserRouteDeleteOK, error)

	UserRouteList(params *UserRouteListParams, opts ...ClientOption) (*UserRouteListOK, error)

	UserRouteUpdate(params *UserRouteUpdateParams, opts ...ClientOption) (*UserRouteUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  UserRouteCreate Creates pairs of user-route relations
*/
func (a *Client) UserRouteCreate(params *UserRouteCreateParams, opts ...ClientOption) (*UserRouteCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserRouteCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "user-route.create",
		Method:             "POST",
		PathPattern:        "/v1/user/route/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserRouteCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UserRouteCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for user-route.create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UserRouteDelete Deletes list of user-route relations
*/
func (a *Client) UserRouteDelete(params *UserRouteDeleteParams, opts ...ClientOption) (*UserRouteDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserRouteDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "user-route.delete",
		Method:             "POST",
		PathPattern:        "/v1/user/route/delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserRouteDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UserRouteDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for user-route.delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UserRouteList Gets list of user-route relations
*/
func (a *Client) UserRouteList(params *UserRouteListParams, opts ...ClientOption) (*UserRouteListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserRouteListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "user-route.list",
		Method:             "POST",
		PathPattern:        "/v1/user/route/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserRouteListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UserRouteListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for user-route.list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UserRouteUpdate Updates pair of user-route relations
*/
func (a *Client) UserRouteUpdate(params *UserRouteUpdateParams, opts ...ClientOption) (*UserRouteUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserRouteUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "user-route.update",
		Method:             "POST",
		PathPattern:        "/v1/user/route/update",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserRouteUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UserRouteUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for user-route.update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}