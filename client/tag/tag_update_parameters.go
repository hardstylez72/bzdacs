// Code generated by go-swagger; DO NOT EDIT.

package tag

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

// NewTagUpdateParams creates a new TagUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTagUpdateParams() *TagUpdateParams {
	return &TagUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTagUpdateParamsWithTimeout creates a new TagUpdateParams object
// with the ability to set a timeout on a request.
func NewTagUpdateParamsWithTimeout(timeout time.Duration) *TagUpdateParams {
	return &TagUpdateParams{
		timeout: timeout,
	}
}

// NewTagUpdateParamsWithContext creates a new TagUpdateParams object
// with the ability to set a context for a request.
func NewTagUpdateParamsWithContext(ctx context.Context) *TagUpdateParams {
	return &TagUpdateParams{
		Context: ctx,
	}
}

// NewTagUpdateParamsWithHTTPClient creates a new TagUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTagUpdateParamsWithHTTPClient(client *http.Client) *TagUpdateParams {
	return &TagUpdateParams{
		HTTPClient: client,
	}
}

/* TagUpdateParams contains all the parameters to send to the API endpoint
   for the tag update operation.

   Typically these are written to a http.Request.
*/
type TagUpdateParams struct {

	/* Req.

	   request
	*/
	Req *models.TagUpdateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tag update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TagUpdateParams) WithDefaults() *TagUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tag update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TagUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tag update params
func (o *TagUpdateParams) WithTimeout(timeout time.Duration) *TagUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tag update params
func (o *TagUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tag update params
func (o *TagUpdateParams) WithContext(ctx context.Context) *TagUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tag update params
func (o *TagUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tag update params
func (o *TagUpdateParams) WithHTTPClient(client *http.Client) *TagUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tag update params
func (o *TagUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the tag update params
func (o *TagUpdateParams) WithReq(req *models.TagUpdateRequest) *TagUpdateParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the tag update params
func (o *TagUpdateParams) SetReq(req *models.TagUpdateRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *TagUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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