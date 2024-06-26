// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PriceUpdateRequest price update request
//
// swagger:model PriceUpdateRequest
type PriceUpdateRequest struct {

	// editable
	Editable bool `json:"Editable,omitempty"`

	// effective date time
	// Format: date-time
	EffectiveDateTime *strfmt.DateTime `json:"EffectiveDateTime,omitempty"`

	// enabled
	Enabled bool `json:"Enabled,omitempty"`

	// min price
	MinPrice float64 `json:"MinPrice,omitempty"`

	// performance ids
	PerformanceIds string `json:"PerformanceIds,omitempty"`

	// price
	Price float64 `json:"Price,omitempty"`

	// price layer type Id
	PriceLayerTypeID int32 `json:"PriceLayerTypeId,omitempty"`

	// price type Id
	PriceTypeID int32 `json:"PriceTypeId,omitempty"`

	// zone Id
	ZoneID int32 `json:"ZoneId,omitempty"`
}

// Validate validates this price update request
func (m *PriceUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEffectiveDateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PriceUpdateRequest) validateEffectiveDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EffectiveDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("EffectiveDateTime", "body", "date-time", m.EffectiveDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this price update request based on context it is used
func (m *PriceUpdateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PriceUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PriceUpdateRequest) UnmarshalBinary(b []byte) error {
	var res PriceUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
