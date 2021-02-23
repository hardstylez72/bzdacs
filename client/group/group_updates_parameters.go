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

// NewGroupUpdatesParams creates a new GroupUpdatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGroupUpdatesParams() *GroupUpdatesParams {
	return &GroupUpdatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGroupUpdatesParamsWithTimeout creates a new GroupUpdatesParams object
// with the ability to set a timeout on a request.
func NewGroupUpdatesParamsWithTimeout(timeout time.Duration) *GroupUpdatesParams {
	return &GroupUpdatesParams{
		timeout: timeout,
	}
}

// NewGroupUpdatesParamsWithContext creates a new GroupUpdatesParams object
// with the ability to set a context for a request.
func NewGroupUpdatesParamsWithContext(ctx context.Context) *GroupUpdatesParams {
	return &GroupUpdatesParams{
		Context: ctx,
	}
}

// NewGroupUpdatesParamsWithHTTPClient creates a new GroupUpdatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGroupUpdatesParamsWithHTTPClient(client *http.Client) *GroupUpdatesParams {
	return &GroupUpdatesParams{
		HTTPClient: client,
	}
}

/* GroupUpdatesParams contains all the parameters to send to the API endpoint
   for the group updates operation.

   Typically these are written to a http.Request.
*/
type GroupUpdatesParams struct {

	/* Req.

	   request
	*/
	Req *models.GroupUpdateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the group updates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupUpdatesParams) WithDefaults() *GroupUpdatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the group updates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupUpdatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the group updates params
func (o *GroupUpdatesParams) WithTimeout(timeout time.Duration) *GroupUpdatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group updates params
func (o *GroupUpdatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group updates params
func (o *GroupUpdatesParams) WithContext(ctx context.Context) *GroupUpdatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group updates params
func (o *GroupUpdatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group updates params
func (o *GroupUpdatesParams) WithHTTPClient(client *http.Client) *GroupUpdatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group updates params
func (o *GroupUpdatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the group updates params
func (o *GroupUpdatesParams) WithReq(req *models.GroupUpdateRequest) *GroupUpdatesParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the group updates params
func (o *GroupUpdatesParams) SetReq(req *models.GroupUpdateRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *GroupUpdatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
