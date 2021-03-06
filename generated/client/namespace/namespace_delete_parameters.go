// Code generated by go-swagger; DO NOT EDIT.

package namespace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/hardstylez72/bzdacs/generated/models"
)

// NewNamespaceDeleteParams creates a new NamespaceDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNamespaceDeleteParams() *NamespaceDeleteParams {
	return &NamespaceDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNamespaceDeleteParamsWithTimeout creates a new NamespaceDeleteParams object
// with the ability to set a timeout on a request.
func NewNamespaceDeleteParamsWithTimeout(timeout time.Duration) *NamespaceDeleteParams {
	return &NamespaceDeleteParams{
		timeout: timeout,
	}
}

// NewNamespaceDeleteParamsWithContext creates a new NamespaceDeleteParams object
// with the ability to set a context for a request.
func NewNamespaceDeleteParamsWithContext(ctx context.Context) *NamespaceDeleteParams {
	return &NamespaceDeleteParams{
		Context: ctx,
	}
}

// NewNamespaceDeleteParamsWithHTTPClient creates a new NamespaceDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewNamespaceDeleteParamsWithHTTPClient(client *http.Client) *NamespaceDeleteParams {
	return &NamespaceDeleteParams{
		HTTPClient: client,
	}
}

/* NamespaceDeleteParams contains all the parameters to send to the API endpoint
   for the namespace delete operation.

   Typically these are written to a http.Request.
*/
type NamespaceDeleteParams struct {

	/* Req.

	   request
	*/
	Req *models.NamespaceDeleteRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the namespace delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NamespaceDeleteParams) WithDefaults() *NamespaceDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the namespace delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NamespaceDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the namespace delete params
func (o *NamespaceDeleteParams) WithTimeout(timeout time.Duration) *NamespaceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the namespace delete params
func (o *NamespaceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the namespace delete params
func (o *NamespaceDeleteParams) WithContext(ctx context.Context) *NamespaceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the namespace delete params
func (o *NamespaceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the namespace delete params
func (o *NamespaceDeleteParams) WithHTTPClient(client *http.Client) *NamespaceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the namespace delete params
func (o *NamespaceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the namespace delete params
func (o *NamespaceDeleteParams) WithReq(req *models.NamespaceDeleteRequest) *NamespaceDeleteParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the namespace delete params
func (o *NamespaceDeleteParams) SetReq(req *models.NamespaceDeleteRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *NamespaceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Req != nil {
		if err := r.SetBodyParam(o.Req); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
