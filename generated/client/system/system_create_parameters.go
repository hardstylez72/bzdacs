// Code generated by go-swagger; DO NOT EDIT.

package system

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

// NewSystemCreateParams creates a new SystemCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSystemCreateParams() *SystemCreateParams {
	return &SystemCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSystemCreateParamsWithTimeout creates a new SystemCreateParams object
// with the ability to set a timeout on a request.
func NewSystemCreateParamsWithTimeout(timeout time.Duration) *SystemCreateParams {
	return &SystemCreateParams{
		timeout: timeout,
	}
}

// NewSystemCreateParamsWithContext creates a new SystemCreateParams object
// with the ability to set a context for a request.
func NewSystemCreateParamsWithContext(ctx context.Context) *SystemCreateParams {
	return &SystemCreateParams{
		Context: ctx,
	}
}

// NewSystemCreateParamsWithHTTPClient creates a new SystemCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSystemCreateParamsWithHTTPClient(client *http.Client) *SystemCreateParams {
	return &SystemCreateParams{
		HTTPClient: client,
	}
}

/* SystemCreateParams contains all the parameters to send to the API endpoint
   for the system create operation.

   Typically these are written to a http.Request.
*/
type SystemCreateParams struct {

	/* Req.

	   request
	*/
	Req *models.SystemInsertRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the system create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemCreateParams) WithDefaults() *SystemCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the system create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the system create params
func (o *SystemCreateParams) WithTimeout(timeout time.Duration) *SystemCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system create params
func (o *SystemCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system create params
func (o *SystemCreateParams) WithContext(ctx context.Context) *SystemCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system create params
func (o *SystemCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system create params
func (o *SystemCreateParams) WithHTTPClient(client *http.Client) *SystemCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system create params
func (o *SystemCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the system create params
func (o *SystemCreateParams) WithReq(req *models.SystemInsertRequest) *SystemCreateParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the system create params
func (o *SystemCreateParams) SetReq(req *models.SystemInsertRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *SystemCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
