// Code generated by go-swagger; DO NOT EDIT.

package relations_group_route

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new relations group route API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for relations group route API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GroupRoutesDelete(params *GroupRoutesDeleteParams, opts ...ClientOption) (*GroupRoutesDeleteOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GroupRoutesDelete Deletes group-routes
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

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
