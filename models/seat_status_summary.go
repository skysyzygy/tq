// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SeatStatusSummary seat status summary
//
// swagger:model SeatStatusSummary
type SeatStatusSummary struct {

	// description
	Description string `json:"Description,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`
}

// Validate validates this seat status summary
func (m *SeatStatusSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this seat status summary based on context it is used
func (m *SeatStatusSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SeatStatusSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SeatStatusSummary) UnmarshalBinary(b []byte) error {
	var res SeatStatusSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
