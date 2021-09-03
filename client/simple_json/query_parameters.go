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

// NewQueryParams creates a new QueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryParams() *QueryParams {
	return &QueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryParamsWithTimeout creates a new QueryParams object
// with the ability to set a timeout on a request.
func NewQueryParamsWithTimeout(timeout time.Duration) *QueryParams {
	return &QueryParams{
		timeout: timeout,
	}
}

// NewQueryParamsWithContext creates a new QueryParams object
// with the ability to set a context for a request.
func NewQueryParamsWithContext(ctx context.Context) *QueryParams {
	return &QueryParams{
		Context: ctx,
	}
}

// NewQueryParamsWithHTTPClient creates a new QueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryParamsWithHTTPClient(client *http.Client) *QueryParams {
	return &QueryParams{
		HTTPClient: client,
	}
}

/* QueryParams contains all the parameters to send to the API endpoint
   for the query operation.

   Typically these are written to a http.Request.
*/
type QueryParams struct {

	/* Options.

	   The options to query.
	*/
	Options *v1.Query

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryParams) WithDefaults() *QueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query params
func (o *QueryParams) WithTimeout(timeout time.Duration) *QueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query params
func (o *QueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query params
func (o *QueryParams) WithContext(ctx context.Context) *QueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query params
func (o *QueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query params
func (o *QueryParams) WithHTTPClient(client *http.Client) *QueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query params
func (o *QueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOptions adds the options to the query params
func (o *QueryParams) WithOptions(options *v1.Query) *QueryParams {
	o.SetOptions(options)
	return o
}

// SetOptions adds the options to the query params
func (o *QueryParams) SetOptions(options *v1.Query) {
	o.Options = options
}

// WriteToRequest writes these params to a swagger request
func (o *QueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
