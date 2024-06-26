// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProductSeatFee product seat fee
//
// swagger:model ProductSeatFee
type ProductSeatFee struct {

	// fee
	Fee *SeatFeeSummary `json:"Fee,omitempty"`

	// fee amount
	FeeAmount float64 `json:"FeeAmount,omitempty"`

	// price
	Price float64 `json:"Price,omitempty"`

	// price type Id
	PriceTypeID int32 `json:"PriceTypeId,omitempty"`

	// zone Id
	ZoneID int32 `json:"ZoneId,omitempty"`
}

// Validate validates this product seat fee
func (m *ProductSeatFee) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductSeatFee) validateFee(formats strfmt.Registry) error {
	if swag.IsZero(m.Fee) { // not required
		return nil
	}

	if m.Fee != nil {
		if err := m.Fee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Fee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Fee")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this product seat fee based on the context it is used
func (m *ProductSeatFee) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFee(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductSeatFee) contextValidateFee(ctx context.Context, formats strfmt.Registry) error {

	if m.Fee != nil {

		if swag.IsZero(m.Fee) { // not required
			return nil
		}

		if err := m.Fee.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Fee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Fee")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProductSeatFee) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductSeatFee) UnmarshalBinary(b []byte) error {
	var res ProductSeatFee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
