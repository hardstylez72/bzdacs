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

// NewGroupRoutesCreateParams creates a new GroupRoutesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGroupRoutesCreateParams() *GroupRoutesCreateParams {
	return &GroupRoutesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGroupRoutesCreateParamsWithTimeout creates a new GroupRoutesCreateParams object
// with the ability to set a timeout on a request.
func NewGroupRoutesCreateParamsWithTimeout(timeout time.Duration) *GroupRoutesCreateParams {
	return &GroupRoutesCreateParams{
		timeout: timeout,
	}
}

// NewGroupRoutesCreateParamsWithContext creates a new GroupRoutesCreateParams object
// with the ability to set a context for a request.
func NewGroupRoutesCreateParamsWithContext(ctx context.Context) *GroupRoutesCreateParams {
	return &GroupRoutesCreateParams{
		Context: ctx,
	}
}

// NewGroupRoutesCreateParamsWithHTTPClient creates a new GroupRoutesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGroupRoutesCreateParamsWithHTTPClient(client *http.Client) *GroupRoutesCreateParams {
	return &GroupRoutesCreateParams{
		HTTPClient: client,
	}
}

/* GroupRoutesCreateParams contains all the parameters to send to the API endpoint
   for the group routes create operation.

   Typically these are written to a http.Request.
*/
type GroupRoutesCreateParams struct {

	/* Req.

	   request
	*/
	Req *models.GroupRouteInsertRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the group routes create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupRoutesCreateParams) WithDefaults() *GroupRoutesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the group routes create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupRoutesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the group routes create params
func (o *GroupRoutesCreateParams) WithTimeout(timeout time.Duration) *GroupRoutesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group routes create params
func (o *GroupRoutesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group routes create params
func (o *GroupRoutesCreateParams) WithContext(ctx context.Context) *GroupRoutesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group routes create params
func (o *GroupRoutesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group routes create params
func (o *GroupRoutesCreateParams) WithHTTPClient(client *http.Client) *GroupRoutesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group routes create params
func (o *GroupRoutesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the group routes create params
func (o *GroupRoutesCreateParams) WithReq(req *models.GroupRouteInsertRequest) *GroupRoutesCreateParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the group routes create params
func (o *GroupRoutesCreateParams) SetReq(req *models.GroupRouteInsertRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *GroupRoutesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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