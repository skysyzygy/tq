// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GenderSummary gender summary
//
// swagger:model GenderSummary
type GenderSummary struct {

	// description
	Description string `json:"Description,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// inactive
	Inactive bool `json:"Inactive,omitempty"`

	// short description
	ShortDescription string `json:"ShortDescription,omitempty"`
}

// Validate validates this gender summary
func (m *GenderSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this gender summary based on context it is used
func (m *GenderSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GenderSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenderSummary) UnmarshalBinary(b []byte) error {
	var res GenderSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
