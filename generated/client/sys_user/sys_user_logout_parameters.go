// Code generated by go-swagger; DO NOT EDIT.

package sys_user

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
)

// NewSysUserLogoutParams creates a new SysUserLogoutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSysUserLogoutParams() *SysUserLogoutParams {
	return &SysUserLogoutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSysUserLogoutParamsWithTimeout creates a new SysUserLogoutParams object
// with the ability to set a timeout on a request.
func NewSysUserLogoutParamsWithTimeout(timeout time.Duration) *SysUserLogoutParams {
	return &SysUserLogoutParams{
		timeout: timeout,
	}
}

// NewSysUserLogoutParamsWithContext creates a new SysUserLogoutParams object
// with the ability to set a context for a request.
func NewSysUserLogoutParamsWithContext(ctx context.Context) *SysUserLogoutParams {
	return &SysUserLogoutParams{
		Context: ctx,
	}
}

// NewSysUserLogoutParamsWithHTTPClient creates a new SysUserLogoutParams object
// with the ability to set a custom HTTPClient for a request.
func NewSysUserLogoutParamsWithHTTPClient(client *http.Client) *SysUserLogoutParams {
	return &SysUserLogoutParams{
		HTTPClient: client,
	}
}

/* SysUserLogoutParams contains all the parameters to send to the API endpoint
   for the sys user logout operation.

   Typically these are written to a http.Request.
*/
type SysUserLogoutParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the sys user logout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SysUserLogoutParams) WithDefaults() *SysUserLogoutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the sys user logout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SysUserLogoutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the sys user logout params
func (o *SysUserLogoutParams) WithTimeout(timeout time.Duration) *SysUserLogoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sys user logout params
func (o *SysUserLogoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sys user logout params
func (o *SysUserLogoutParams) WithContext(ctx context.Context) *SysUserLogoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sys user logout params
func (o *SysUserLogoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sys user logout params
func (o *SysUserLogoutParams) WithHTTPClient(client *http.Client) *SysUserLogoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sys user logout params
func (o *SysUserLogoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SysUserLogoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
