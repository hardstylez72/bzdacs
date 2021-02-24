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

// NewGroupRoutesListParams creates a new GroupRoutesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGroupRoutesListParams() *GroupRoutesListParams {
	return &GroupRoutesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGroupRoutesListParamsWithTimeout creates a new GroupRoutesListParams object
// with the ability to set a timeout on a request.
func NewGroupRoutesListParamsWithTimeout(timeout time.Duration) *GroupRoutesListParams {
	return &GroupRoutesListParams{
		timeout: timeout,
	}
}

// NewGroupRoutesListParamsWithContext creates a new GroupRoutesListParams object
// with the ability to set a context for a request.
func NewGroupRoutesListParamsWithContext(ctx context.Context) *GroupRoutesListParams {
	return &GroupRoutesListParams{
		Context: ctx,
	}
}

// NewGroupRoutesListParamsWithHTTPClient creates a new GroupRoutesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGroupRoutesListParamsWithHTTPClient(client *http.Client) *GroupRoutesListParams {
	return &GroupRoutesListParams{
		HTTPClient: client,
	}
}

/* GroupRoutesListParams contains all the parameters to send to the API endpoint
   for the group routes list operation.

   Typically these are written to a http.Request.
*/
type GroupRoutesListParams struct {

	/* Req.

	   request
	*/
	Req *models.GroupRouteListRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the group routes list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupRoutesListParams) WithDefaults() *GroupRoutesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the group routes list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupRoutesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the group routes list params
func (o *GroupRoutesListParams) WithTimeout(timeout time.Duration) *GroupRoutesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group routes list params
func (o *GroupRoutesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group routes list params
func (o *GroupRoutesListParams) WithContext(ctx context.Context) *GroupRoutesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group routes list params
func (o *GroupRoutesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group routes list params
func (o *GroupRoutesListParams) WithHTTPClient(client *http.Client) *GroupRoutesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group routes list params
func (o *GroupRoutesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the group routes list params
func (o *GroupRoutesListParams) WithReq(req *models.GroupRouteListRequest) *GroupRoutesListParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the group routes list params
func (o *GroupRoutesListParams) SetReq(req *models.GroupRouteListRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *GroupRoutesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
