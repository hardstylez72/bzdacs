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

	"github.com/hardstylez72/bzdacs/generated/models"
)

// NewGroupDeletesParams creates a new GroupDeletesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGroupDeletesParams() *GroupDeletesParams {
	return &GroupDeletesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGroupDeletesParamsWithTimeout creates a new GroupDeletesParams object
// with the ability to set a timeout on a request.
func NewGroupDeletesParamsWithTimeout(timeout time.Duration) *GroupDeletesParams {
	return &GroupDeletesParams{
		timeout: timeout,
	}
}

// NewGroupDeletesParamsWithContext creates a new GroupDeletesParams object
// with the ability to set a context for a request.
func NewGroupDeletesParamsWithContext(ctx context.Context) *GroupDeletesParams {
	return &GroupDeletesParams{
		Context: ctx,
	}
}

// NewGroupDeletesParamsWithHTTPClient creates a new GroupDeletesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGroupDeletesParamsWithHTTPClient(client *http.Client) *GroupDeletesParams {
	return &GroupDeletesParams{
		HTTPClient: client,
	}
}

/* GroupDeletesParams contains all the parameters to send to the API endpoint
   for the group deletes operation.

   Typically these are written to a http.Request.
*/
type GroupDeletesParams struct {

	/* Req.

	   request
	*/
	Req *models.GroupDeleteRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the group deletes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupDeletesParams) WithDefaults() *GroupDeletesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the group deletes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupDeletesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the group deletes params
func (o *GroupDeletesParams) WithTimeout(timeout time.Duration) *GroupDeletesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group deletes params
func (o *GroupDeletesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group deletes params
func (o *GroupDeletesParams) WithContext(ctx context.Context) *GroupDeletesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group deletes params
func (o *GroupDeletesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group deletes params
func (o *GroupDeletesParams) WithHTTPClient(client *http.Client) *GroupDeletesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group deletes params
func (o *GroupDeletesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the group deletes params
func (o *GroupDeletesParams) WithReq(req *models.GroupDeleteRequest) *GroupDeletesParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the group deletes params
func (o *GroupDeletesParams) SetReq(req *models.GroupDeleteRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *GroupDeletesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
