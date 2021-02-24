// Code generated by go-swagger; DO NOT EDIT.

package user_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user group API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user group API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	UserGroupCreate(params *UserGroupCreateParams, opts ...ClientOption) (*UserGroupCreateOK, error)

	UserGroupDelete(params *UserGroupDeleteParams, opts ...ClientOption) (*UserGroupDeleteOK, error)

	UserGroupList(params *UserGroupListParams, opts ...ClientOption) (*UserGroupListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  UserGroupCreate Creates user-group relations
*/
func (a *Client) UserGroupCreate(params *UserGroupCreateParams, opts ...ClientOption) (*UserGroupCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserGroupCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "user-group.create",
		Method:             "POST",
		PathPattern:        "/v1/user/group/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserGroupCreateReader{formats: a.formats},
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
	success, ok := result.(*UserGroupCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for user-group.create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UserGroupDelete Deletes user-group relations
*/
func (a *Client) UserGroupDelete(params *UserGroupDeleteParams, opts ...ClientOption) (*UserGroupDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserGroupDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "user-group.delete",
		Method:             "POST",
		PathPattern:        "/v1/user/group/delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserGroupDeleteReader{formats: a.formats},
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
	success, ok := result.(*UserGroupDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for user-group.delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UserGroupList Gets list of user-group relations
*/
func (a *Client) UserGroupList(params *UserGroupListParams, opts ...ClientOption) (*UserGroupListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserGroupListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "user-group.list",
		Method:             "POST",
		PathPattern:        "/v1/user/group/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserGroupListReader{formats: a.formats},
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
	success, ok := result.(*UserGroupListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for user-group.list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}