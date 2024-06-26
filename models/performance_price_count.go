// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PerformancePriceCount performance price count
//
// swagger:model PerformancePriceCount
type PerformancePriceCount struct {

	// count
	Count int32 `json:"Count,omitempty"`
}

// Validate validates this performance price count
func (m *PerformancePriceCount) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this performance price count based on context it is used
func (m *PerformancePriceCount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PerformancePriceCount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformancePriceCount) UnmarshalBinary(b []byte) error {
	var res PerformancePriceCount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
