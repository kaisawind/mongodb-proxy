// Code generated by go-swagger; DO NOT EDIT.

package simple_json

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
	v1 "github.com/kaisawind/mongodb-proxy/v1"
)

// NewAnnotationQueryParams creates a new AnnotationQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAnnotationQueryParams() *AnnotationQueryParams {
	return &AnnotationQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAnnotationQueryParamsWithTimeout creates a new AnnotationQueryParams object
// with the ability to set a timeout on a request.
func NewAnnotationQueryParamsWithTimeout(timeout time.Duration) *AnnotationQueryParams {
	return &AnnotationQueryParams{
		timeout: timeout,
	}
}

// NewAnnotationQueryParamsWithContext creates a new AnnotationQueryParams object
// with the ability to set a context for a request.
func NewAnnotationQueryParamsWithContext(ctx context.Context) *AnnotationQueryParams {
	return &AnnotationQueryParams{
		Context: ctx,
	}
}

// NewAnnotationQueryParamsWithHTTPClient creates a new AnnotationQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewAnnotationQueryParamsWithHTTPClient(client *http.Client) *AnnotationQueryParams {
	return &AnnotationQueryParams{
		HTTPClient: client,
	}
}

/* AnnotationQueryParams contains all the parameters to send to the API endpoint
   for the annotation query operation.

   Typically these are written to a http.Request.
*/
type AnnotationQueryParams struct {

	/* Options.

	   The options to query.
	*/
	Options *v1.Target

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the annotation query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AnnotationQueryParams) WithDefaults() *AnnotationQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the annotation query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AnnotationQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the annotation query params
func (o *AnnotationQueryParams) WithTimeout(timeout time.Duration) *AnnotationQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the annotation query params
func (o *AnnotationQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the annotation query params
func (o *AnnotationQueryParams) WithContext(ctx context.Context) *AnnotationQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the annotation query params
func (o *AnnotationQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the annotation query params
func (o *AnnotationQueryParams) WithHTTPClient(client *http.Client) *AnnotationQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the annotation query params
func (o *AnnotationQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOptions adds the options to the annotation query params
func (o *AnnotationQueryParams) WithOptions(options *v1.Target) *AnnotationQueryParams {
	o.SetOptions(options)
	return o
}

// SetOptions adds the options to the annotation query params
func (o *AnnotationQueryParams) SetOptions(options *v1.Target) {
	o.Options = options
}

// WriteToRequest writes these params to a swagger request
func (o *AnnotationQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Options != nil {
		if err := r.SetBodyParam(o.Options); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
