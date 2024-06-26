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

// OrderProductViewProductPerformance order product view product performance
//
// swagger:model OrderProductViewProductPerformance
type OrderProductViewProductPerformance struct {

	// Id
	ID int32 `json:"Id,omitempty"`

	// line item
	LineItem *OrderProductViewLineItem `json:"LineItem,omitempty"`
}

// Validate validates this order product view product performance
func (m *OrderProductViewProductPerformance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLineItem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderProductViewProductPerformance) validateLineItem(formats strfmt.Registry) error {
	if swag.IsZero(m.LineItem) { // not required
		return nil
	}

	if m.LineItem != nil {
		if err := m.LineItem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LineItem")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("LineItem")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this order product view product performance based on the context it is used
func (m *OrderProductViewProductPerformance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLineItem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderProductViewProductPerformance) contextValidateLineItem(ctx context.Context, formats strfmt.Registry) error {

	if m.LineItem != nil {

		if swag.IsZero(m.LineItem) { // not required
			return nil
		}

		if err := m.LineItem.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LineItem")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("LineItem")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderProductViewProductPerformance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderProductViewProductPerformance) UnmarshalBinary(b []byte) error {
	var res OrderProductViewProductPerformance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
