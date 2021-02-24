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

// NewUserListParams creates a new UserListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserListParams() *UserListParams {
	return &UserListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserListParamsWithTimeout creates a new UserListParams object
// with the ability to set a timeout on a request.
func NewUserListParamsWithTimeout(timeout time.Duration) *UserListParams {
	return &UserListParams{
		timeout: timeout,
	}
}

// NewUserListParamsWithContext creates a new UserListParams object
// with the ability to set a context for a request.
func NewUserListParamsWithContext(ctx context.Context) *UserListParams {
	return &UserListParams{
		Context: ctx,
	}
}

// NewUserListParamsWithHTTPClient creates a new UserListParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserListParamsWithHTTPClient(client *http.Client) *UserListParams {
	return &UserListParams{
		HTTPClient: client,
	}
}

/* UserListParams contains all the parameters to send to the API endpoint
   for the user list operation.

   Typically these are written to a http.Request.
*/
type UserListParams struct {

	/* Req.

	   request
	*/
	Req *models.UserListRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserListParams) WithDefaults() *UserListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user list params
func (o *UserListParams) WithTimeout(timeout time.Duration) *UserListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user list params
func (o *UserListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user list params
func (o *UserListParams) WithContext(ctx context.Context) *UserListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user list params
func (o *UserListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user list params
func (o *UserListParams) WithHTTPClient(client *http.Client) *UserListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user list params
func (o *UserListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the user list params
func (o *UserListParams) WithReq(req *models.UserListRequest) *UserListParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the user list params
func (o *UserListParams) SetReq(req *models.UserListRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *UserListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
