// Code generated by go-swagger; DO NOT EDIT.

package group_route

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

// NewGroupRoutesDeleteParams creates a new GroupRoutesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGroupRoutesDeleteParams() *GroupRoutesDeleteParams {
	return &GroupRoutesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGroupRoutesDeleteParamsWithTimeout creates a new GroupRoutesDeleteParams object
// with the ability to set a timeout on a request.
func NewGroupRoutesDeleteParamsWithTimeout(timeout time.Duration) *GroupRoutesDeleteParams {
	return &GroupRoutesDeleteParams{
		timeout: timeout,
	}
}

// NewGroupRoutesDeleteParamsWithContext creates a new GroupRoutesDeleteParams object
// with the ability to set a context for a request.
func NewGroupRoutesDeleteParamsWithContext(ctx context.Context) *GroupRoutesDeleteParams {
	return &GroupRoutesDeleteParams{
		Context: ctx,
	}
}

// NewGroupRoutesDeleteParamsWithHTTPClient creates a new GroupRoutesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewGroupRoutesDeleteParamsWithHTTPClient(client *http.Client) *GroupRoutesDeleteParams {
	return &GroupRoutesDeleteParams{
		HTTPClient: client,
	}
}

/* GroupRoutesDeleteParams contains all the parameters to send to the API endpoint
   for the group routes delete operation.

   Typically these are written to a http.Request.
*/
type GroupRoutesDeleteParams struct {

	/* Req.

	   request
	*/
	Req *models.GroupRouteDeleteRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the group routes delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupRoutesDeleteParams) WithDefaults() *GroupRoutesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the group routes delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupRoutesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the group routes delete params
func (o *GroupRoutesDeleteParams) WithTimeout(timeout time.Duration) *GroupRoutesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group routes delete params
func (o *GroupRoutesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group routes delete params
func (o *GroupRoutesDeleteParams) WithContext(ctx context.Context) *GroupRoutesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group routes delete params
func (o *GroupRoutesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group routes delete params
func (o *GroupRoutesDeleteParams) WithHTTPClient(client *http.Client) *GroupRoutesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group routes delete params
func (o *GroupRoutesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the group routes delete params
func (o *GroupRoutesDeleteParams) WithReq(req *models.GroupRouteDeleteRequest) *GroupRoutesDeleteParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the group routes delete params
func (o *GroupRoutesDeleteParams) SetReq(req *models.GroupRouteDeleteRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *GroupRoutesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
