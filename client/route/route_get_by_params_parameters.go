// Code generated by go-swagger; DO NOT EDIT.

package route

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

// NewRouteGetByParamsParams creates a new RouteGetByParamsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRouteGetByParamsParams() *RouteGetByParamsParams {
	return &RouteGetByParamsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRouteGetByParamsParamsWithTimeout creates a new RouteGetByParamsParams object
// with the ability to set a timeout on a request.
func NewRouteGetByParamsParamsWithTimeout(timeout time.Duration) *RouteGetByParamsParams {
	return &RouteGetByParamsParams{
		timeout: timeout,
	}
}

// NewRouteGetByParamsParamsWithContext creates a new RouteGetByParamsParams object
// with the ability to set a context for a request.
func NewRouteGetByParamsParamsWithContext(ctx context.Context) *RouteGetByParamsParams {
	return &RouteGetByParamsParams{
		Context: ctx,
	}
}

// NewRouteGetByParamsParamsWithHTTPClient creates a new RouteGetByParamsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRouteGetByParamsParamsWithHTTPClient(client *http.Client) *RouteGetByParamsParams {
	return &RouteGetByParamsParams{
		HTTPClient: client,
	}
}

/* RouteGetByParamsParams contains all the parameters to send to the API endpoint
   for the route get by params operation.

   Typically these are written to a http.Request.
*/
type RouteGetByParamsParams struct {

	/* Req.

	   request
	*/
	Req *models.RouteGetByParamsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the route get by params params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RouteGetByParamsParams) WithDefaults() *RouteGetByParamsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the route get by params params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RouteGetByParamsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the route get by params params
func (o *RouteGetByParamsParams) WithTimeout(timeout time.Duration) *RouteGetByParamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the route get by params params
func (o *RouteGetByParamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the route get by params params
func (o *RouteGetByParamsParams) WithContext(ctx context.Context) *RouteGetByParamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the route get by params params
func (o *RouteGetByParamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the route get by params params
func (o *RouteGetByParamsParams) WithHTTPClient(client *http.Client) *RouteGetByParamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the route get by params params
func (o *RouteGetByParamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the route get by params params
func (o *RouteGetByParamsParams) WithReq(req *models.RouteGetByParamsRequest) *RouteGetByParamsParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the route get by params params
func (o *RouteGetByParamsParams) SetReq(req *models.RouteGetByParamsRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *RouteGetByParamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
