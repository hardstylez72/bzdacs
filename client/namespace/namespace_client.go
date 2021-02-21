// Code generated by go-swagger; DO NOT EDIT.

package namespace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new namespace API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for namespace API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	NamespaceCreate(params *NamespaceCreateParams, opts ...ClientOption) (*NamespaceCreateOK, error)

	NamespaceDelete(params *NamespaceDeleteParams, opts ...ClientOption) (*NamespaceDeleteOK, error)

	NamespaceGet(params *NamespaceGetParams, opts ...ClientOption) (*NamespaceGetOK, error)

	NamespaceList(params *NamespaceListParams, opts ...ClientOption) (*NamespaceListOK, error)

	NamespaceUpdate(params *NamespaceUpdateParams, opts ...ClientOption) (*NamespaceUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  NamespaceCreate Creates namespace
*/
func (a *Client) NamespaceCreate(params *NamespaceCreateParams, opts ...ClientOption) (*NamespaceCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNamespaceCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "namespace.create",
		Method:             "POST",
		PathPattern:        "/v1/namespace/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NamespaceCreateReader{formats: a.formats},
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
	success, ok := result.(*NamespaceCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for namespace.create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  NamespaceDelete Deletes selected namespace
*/
func (a *Client) NamespaceDelete(params *NamespaceDeleteParams, opts ...ClientOption) (*NamespaceDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNamespaceDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "namespace.delete",
		Method:             "POST",
		PathPattern:        "/v1/namespace/delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NamespaceDeleteReader{formats: a.formats},
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
	success, ok := result.(*NamespaceDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for namespace.delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  NamespaceGet Gets namespace with specified params
*/
func (a *Client) NamespaceGet(params *NamespaceGetParams, opts ...ClientOption) (*NamespaceGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNamespaceGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "namespace.get",
		Method:             "POST",
		PathPattern:        "/v1/namespace/get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NamespaceGetReader{formats: a.formats},
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
	success, ok := result.(*NamespaceGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for namespace.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  NamespaceList List of namespaces belong to selected system
*/
func (a *Client) NamespaceList(params *NamespaceListParams, opts ...ClientOption) (*NamespaceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNamespaceListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "namespace.list",
		Method:             "POST",
		PathPattern:        "/v1/namespace/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NamespaceListReader{formats: a.formats},
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
	success, ok := result.(*NamespaceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for namespace.list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  NamespaceUpdate Updates namespace
*/
func (a *Client) NamespaceUpdate(params *NamespaceUpdateParams, opts ...ClientOption) (*NamespaceUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNamespaceUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "namespace.update",
		Method:             "POST",
		PathPattern:        "/v1/namespace/update",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NamespaceUpdateReader{formats: a.formats},
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
	success, ok := result.(*NamespaceUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for namespace.update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
