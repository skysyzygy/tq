// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PerformancePriceResponse performance price response
//
// swagger:model PerformancePriceResponse
type PerformancePriceResponse struct {

	// layer types
	LayerTypes []*PriceLayerTypeSummary `json:"LayerTypes"`

	// performance price layers
	PerformancePriceLayers []*PerformancePriceLayerSummary `json:"PerformancePriceLayers"`
}

// Validate validates this performance price response
func (m *PerformancePriceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLayerTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerformancePriceLayers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformancePriceResponse) validateLayerTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.LayerTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.LayerTypes); i++ {
		if swag.IsZero(m.LayerTypes[i]) { // not required
			continue
		}

		if m.LayerTypes[i] != nil {
			if err := m.LayerTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("LayerTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("LayerTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PerformancePriceResponse) validatePerformancePriceLayers(formats strfmt.Registry) error {
	if swag.IsZero(m.PerformancePriceLayers) { // not required
		return nil
	}

	for i := 0; i < len(m.PerformancePriceLayers); i++ {
		if swag.IsZero(m.PerformancePriceLayers[i]) { // not required
			continue
		}

		if m.PerformancePriceLayers[i] != nil {
			if err := m.PerformancePriceLayers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PerformancePriceLayers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("PerformancePriceLayers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this performance price response based on the context it is used
func (m *PerformancePriceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLayerTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePerformancePriceLayers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformancePriceResponse) contextValidateLayerTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LayerTypes); i++ {

		if m.LayerTypes[i] != nil {

			if swag.IsZero(m.LayerTypes[i]) { // not required
				return nil
			}

			if err := m.LayerTypes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("LayerTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("LayerTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PerformancePriceResponse) contextValidatePerformancePriceLayers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PerformancePriceLayers); i++ {

		if m.PerformancePriceLayers[i] != nil {

			if swag.IsZero(m.PerformancePriceLayers[i]) { // not required
				return nil
			}

			if err := m.PerformancePriceLayers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PerformancePriceLayers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("PerformancePriceLayers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformancePriceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformancePriceResponse) UnmarshalBinary(b []byte) error {
	var res PerformancePriceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
