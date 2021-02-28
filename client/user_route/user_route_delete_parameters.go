// Code generated by go-swagger; DO NOT EDIT.

package user_route

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

// NewUserRouteDeleteParams creates a new UserRouteDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserRouteDeleteParams() *UserRouteDeleteParams {
	return &UserRouteDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserRouteDeleteParamsWithTimeout creates a new UserRouteDeleteParams object
// with the ability to set a timeout on a request.
func NewUserRouteDeleteParamsWithTimeout(timeout time.Duration) *UserRouteDeleteParams {
	return &UserRouteDeleteParams{
		timeout: timeout,
	}
}

// NewUserRouteDeleteParamsWithContext creates a new UserRouteDeleteParams object
// with the ability to set a context for a request.
func NewUserRouteDeleteParamsWithContext(ctx context.Context) *UserRouteDeleteParams {
	return &UserRouteDeleteParams{
		Context: ctx,
	}
}

// NewUserRouteDeleteParamsWithHTTPClient creates a new UserRouteDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserRouteDeleteParamsWithHTTPClient(client *http.Client) *UserRouteDeleteParams {
	return &UserRouteDeleteParams{
		HTTPClient: client,
	}
}

/* UserRouteDeleteParams contains all the parameters to send to the API endpoint
   for the user route delete operation.

   Typically these are written to a http.Request.
*/
type UserRouteDeleteParams struct {

	/* Req.

	   request
	*/
	Req *models.UserRouteDeleteRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user route delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserRouteDeleteParams) WithDefaults() *UserRouteDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user route delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserRouteDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user route delete params
func (o *UserRouteDeleteParams) WithTimeout(timeout time.Duration) *UserRouteDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user route delete params
func (o *UserRouteDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user route delete params
func (o *UserRouteDeleteParams) WithContext(ctx context.Context) *UserRouteDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user route delete params
func (o *UserRouteDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user route delete params
func (o *UserRouteDeleteParams) WithHTTPClient(client *http.Client) *UserRouteDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user route delete params
func (o *UserRouteDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the user route delete params
func (o *UserRouteDeleteParams) WithReq(req *models.UserRouteDeleteRequest) *UserRouteDeleteParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the user route delete params
func (o *UserRouteDeleteParams) SetReq(req *models.UserRouteDeleteRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *UserRouteDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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