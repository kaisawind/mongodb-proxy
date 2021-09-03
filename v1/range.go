// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Range Range specifies the time range the request is valid for.
//
// swagger:model Range
type Range struct {

	// from
	// Format: date-time
	From strfmt.DateTime `json:"from,omitempty"`

	// to
	// Format: date-time
	To strfmt.DateTime `json:"to,omitempty"`
}

// Validate validates this range
func (m *Range) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Range) validateFrom(formats strfmt.Registry) error {
	if swag.IsZero(m.From) { // not required
		return nil
	}

	if err := validate.FormatOf("from", "body", "date-time", m.From.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Range) validateTo(formats strfmt.Registry) error {
	if swag.IsZero(m.To) { // not required
		return nil
	}

	if err := validate.FormatOf("to", "body", "date-time", m.To.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this range based on context it is used
func (m *Range) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Range) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Range) UnmarshalBinary(b []byte) error {
	var res Range
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
