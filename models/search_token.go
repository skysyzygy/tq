// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SearchToken search token
//
// swagger:model SearchToken
type SearchToken struct {

	// key
	Key string `json:"Key,omitempty"`

	// value
	Value string `json:"Value,omitempty"`
}

// Validate validates this search token
func (m *SearchToken) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search token based on context it is used
func (m *SearchToken) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SearchToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchToken) UnmarshalBinary(b []byte) error {
	var res SearchToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
