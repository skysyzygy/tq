// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SpecialActivityStatusSummary special activity status summary
//
// swagger:model SpecialActivityStatusSummary
type SpecialActivityStatusSummary struct {

	// description
	Description string `json:"Description,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// inactive
	Inactive bool `json:"Inactive,omitempty"`
}

// Validate validates this special activity status summary
func (m *SpecialActivityStatusSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this special activity status summary based on context it is used
func (m *SpecialActivityStatusSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SpecialActivityStatusSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SpecialActivityStatusSummary) UnmarshalBinary(b []byte) error {
	var res SpecialActivityStatusSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
