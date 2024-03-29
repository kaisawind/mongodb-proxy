// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Annotation Annotation is the object passed by Grafana when it fetches annotations.
//
// swagger:model Annotation
type Annotation struct {

	// datasource
	Datasource string `json:"datasource,omitempty"`

	// enable
	Enable bool `json:"enable,omitempty"`

	// icon color
	IconColor string `json:"iconColor,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// query
	Query string `json:"query,omitempty"`

	// show line
	ShowLine bool `json:"showLine,omitempty"`
}

// Validate validates this annotation
func (m *Annotation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this annotation based on context it is used
func (m *Annotation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Annotation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Annotation) UnmarshalBinary(b []byte) error {
	var res Annotation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
