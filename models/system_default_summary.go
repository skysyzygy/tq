// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SystemDefaultSummary system default summary
//
// swagger:model SystemDefaultSummary
type SystemDefaultSummary struct {

	// field name
	FieldName string `json:"FieldName,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// inactive
	Inactive bool `json:"Inactive,omitempty"`

	// parent table
	ParentTable string `json:"ParentTable,omitempty"`

	// value
	Value string `json:"Value,omitempty"`
}

// Validate validates this system default summary
func (m *SystemDefaultSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this system default summary based on context it is used
func (m *SystemDefaultSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SystemDefaultSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemDefaultSummary) UnmarshalBinary(b []byte) error {
	var res SystemDefaultSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
