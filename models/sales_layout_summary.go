// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SalesLayoutSummary sales layout summary
//
// swagger:model SalesLayoutSummary
type SalesLayoutSummary struct {

	// description
	Description string `json:"Description,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// inactive
	Inactive bool `json:"Inactive,omitempty"`

	// primary indicator
	PrimaryIndicator bool `json:"PrimaryIndicator,omitempty"`
}

// Validate validates this sales layout summary
func (m *SalesLayoutSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sales layout summary based on context it is used
func (m *SalesLayoutSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SalesLayoutSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SalesLayoutSummary) UnmarshalBinary(b []byte) error {
	var res SalesLayoutSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
