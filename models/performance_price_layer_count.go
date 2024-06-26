// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PerformancePriceLayerCount performance price layer count
//
// swagger:model PerformancePriceLayerCount
type PerformancePriceLayerCount struct {

	// layer count
	LayerCount int32 `json:"LayerCount,omitempty"`

	// performance Id
	PerformanceID int32 `json:"PerformanceId,omitempty"`
}

// Validate validates this performance price layer count
func (m *PerformancePriceLayerCount) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this performance price layer count based on context it is used
func (m *PerformancePriceLayerCount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PerformancePriceLayerCount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformancePriceLayerCount) UnmarshalBinary(b []byte) error {
	var res PerformancePriceLayerCount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
