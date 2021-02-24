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

	"github.com/hardstylez72/bzdacs/models"
)

// NewUserCreateParams creates a new UserCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserCreateParams() *UserCreateParams {
	return &UserCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserCreateParamsWithTimeout creates a new UserCreateParams object
// with the ability to set a timeout on a request.
func NewUserCreateParamsWithTimeout(timeout time.Duration) *UserCreateParams {
	return &UserCreateParams{
		timeout: timeout,
	}
}

// NewUserCreateParamsWithContext creates a new UserCreateParams object
// with the ability to set a context for a request.
func NewUserCreateParamsWithContext(ctx context.Context) *UserCreateParams {
	return &UserCreateParams{
		Context: ctx,
	}
}

// NewUserCreateParamsWithHTTPClient creates a new UserCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserCreateParamsWithHTTPClient(client *http.Client) *UserCreateParams {
	return &UserCreateParams{
		HTTPClient: client,
	}
}

/* UserCreateParams contains all the parameters to send to the API endpoint
   for the user create operation.

   Typically these are written to a http.Request.
*/
type UserCreateParams struct {

	/* Req.

	   request
	*/
	Req *models.UserCreateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserCreateParams) WithDefaults() *UserCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user create params
func (o *UserCreateParams) WithTimeout(timeout time.Duration) *UserCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user create params
func (o *UserCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user create params
func (o *UserCreateParams) WithContext(ctx context.Context) *UserCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user create params
func (o *UserCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user create params
func (o *UserCreateParams) WithHTTPClient(client *http.Client) *UserCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user create params
func (o *UserCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the user create params
func (o *UserCreateParams) WithReq(req *models.UserCreateRequest) *UserCreateParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the user create params
func (o *UserCreateParams) SetReq(req *models.UserCreateRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *UserCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
