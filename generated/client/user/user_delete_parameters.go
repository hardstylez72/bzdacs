// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewUserDeleteParams creates a new UserDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserDeleteParams() *UserDeleteParams {
	return &UserDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserDeleteParamsWithTimeout creates a new UserDeleteParams object
// with the ability to set a timeout on a request.
func NewUserDeleteParamsWithTimeout(timeout time.Duration) *UserDeleteParams {
	return &UserDeleteParams{
		timeout: timeout,
	}
}

// NewUserDeleteParamsWithContext creates a new UserDeleteParams object
// with the ability to set a context for a request.
func NewUserDeleteParamsWithContext(ctx context.Context) *UserDeleteParams {
	return &UserDeleteParams{
		Context: ctx,
	}
}

// NewUserDeleteParamsWithHTTPClient creates a new UserDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserDeleteParamsWithHTTPClient(client *http.Client) *UserDeleteParams {
	return &UserDeleteParams{
		HTTPClient: client,
	}
}

/* UserDeleteParams contains all the parameters to send to the API endpoint
   for the user delete operation.

   Typically these are written to a http.Request.
*/
type UserDeleteParams struct {

	/* Req.

	   request
	*/
	Req *models.UserDeleteRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserDeleteParams) WithDefaults() *UserDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user delete params
func (o *UserDeleteParams) WithTimeout(timeout time.Duration) *UserDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user delete params
func (o *UserDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user delete params
func (o *UserDeleteParams) WithContext(ctx context.Context) *UserDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user delete params
func (o *UserDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user delete params
func (o *UserDeleteParams) WithHTTPClient(client *http.Client) *UserDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user delete params
func (o *UserDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the user delete params
func (o *UserDeleteParams) WithReq(req *models.UserDeleteRequest) *UserDeleteParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the user delete params
func (o *UserDeleteParams) SetReq(req *models.UserDeleteRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *UserDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
