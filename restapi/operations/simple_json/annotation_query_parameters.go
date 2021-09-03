// Code generated by go-swagger; DO NOT EDIT.

package simple_json

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"
	v1 "github.com/kaisawind/mongodb-proxy/v1"
)

// NewAnnotationQueryParams creates a new AnnotationQueryParams object
//
// There are no default values defined in the spec.
func NewAnnotationQueryParams() AnnotationQueryParams {

	return AnnotationQueryParams{}
}

// AnnotationQueryParams contains all the bound params for the annotation query operation
// typically these are obtained from a http.Request
//
// swagger:parameters AnnotationQuery
type AnnotationQueryParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The options to query.
	  In: body
	*/
	Options *v1.Target
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAnnotationQueryParams() beforehand.
func (o *AnnotationQueryParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body v1.Target
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("options", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Options = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
