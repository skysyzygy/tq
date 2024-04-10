// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ParameterOption parameter option
//
// swagger:model ParameterOption
type ParameterOption struct {

	// display text
	DisplayText string `json:"DisplayText,omitempty"`

	// inactive
	Inactive bool `json:"Inactive,omitempty"`

	// value
	Value string `json:"Value,omitempty"`
}

// Validate validates this parameter option
func (m *ParameterOption) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this parameter option based on context it is used
func (m *ParameterOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ParameterOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ParameterOption) UnmarshalBinary(b []byte) error {
	var res ParameterOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
