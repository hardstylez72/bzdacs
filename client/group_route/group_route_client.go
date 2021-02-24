// Code generated by go-swagger; DO NOT EDIT.

package group_route

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new group route API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for group route API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GroupRoutesCreate(params *GroupRoutesCreateParams, opts ...ClientOption) (*GroupRoutesCreateOK, error)

	GroupRoutesDelete(params *GroupRoutesDeleteParams, opts ...ClientOption) (*GroupRoutesDeleteOK, error)

	GroupRoutesList(params *GroupRoutesListParams, opts ...ClientOption) (*GroupRoutesListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GroupRoutesCreate Creates group-routes relations
*/
func (a *Client) GroupRoutesCreate(params *GroupRoutesCreateParams, opts ...ClientOption) (*GroupRoutesCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGroupRoutesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "group-routes.create",
		Method:             "POST",
		PathPattern:        "/v1/group/route/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GroupRoutesCreateReader{formats: a.formats},
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
	success, ok := result.(*GroupRoutesCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for group-routes.create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GroupRoutesDelete Deletes group-routes relations
*/
func (a *Client) GroupRoutesDelete(params *GroupRoutesDeleteParams, opts ...ClientOption) (*GroupRoutesDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGroupRoutesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "group-routes.delete",
		Method:             "POST",
		PathPattern:        "/v1/group/route/delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GroupRoutesDeleteReader{formats: a.formats},
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
	success, ok := result.(*GroupRoutesDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for group-routes.delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GroupRoutesList Gets list of group-routes relations
*/
func (a *Client) GroupRoutesList(params *GroupRoutesListParams, opts ...ClientOption) (*GroupRoutesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGroupRoutesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "group-routes.list",
		Method:             "POST",
		PathPattern:        "/v1/group/route/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GroupRoutesListReader{formats: a.formats},
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
	success, ok := result.(*GroupRoutesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for group-routes.list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
