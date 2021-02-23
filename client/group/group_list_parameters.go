// Code generated by go-swagger; DO NOT EDIT.

package group

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

// NewGroupListParams creates a new GroupListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGroupListParams() *GroupListParams {
	return &GroupListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGroupListParamsWithTimeout creates a new GroupListParams object
// with the ability to set a timeout on a request.
func NewGroupListParamsWithTimeout(timeout time.Duration) *GroupListParams {
	return &GroupListParams{
		timeout: timeout,
	}
}

// NewGroupListParamsWithContext creates a new GroupListParams object
// with the ability to set a context for a request.
func NewGroupListParamsWithContext(ctx context.Context) *GroupListParams {
	return &GroupListParams{
		Context: ctx,
	}
}

// NewGroupListParamsWithHTTPClient creates a new GroupListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGroupListParamsWithHTTPClient(client *http.Client) *GroupListParams {
	return &GroupListParams{
		HTTPClient: client,
	}
}

/* GroupListParams contains all the parameters to send to the API endpoint
   for the group list operation.

   Typically these are written to a http.Request.
*/
type GroupListParams struct {

	/* Req.

	   request
	*/
	Req *models.GroupListRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the group list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupListParams) WithDefaults() *GroupListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the group list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the group list params
func (o *GroupListParams) WithTimeout(timeout time.Duration) *GroupListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group list params
func (o *GroupListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group list params
func (o *GroupListParams) WithContext(ctx context.Context) *GroupListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group list params
func (o *GroupListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group list params
func (o *GroupListParams) WithHTTPClient(client *http.Client) *GroupListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group list params
func (o *GroupListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the group list params
func (o *GroupListParams) WithReq(req *models.GroupListRequest) *GroupListParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the group list params
func (o *GroupListParams) SetReq(req *models.GroupListRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *GroupListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
