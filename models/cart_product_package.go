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

// CartProductPackage cart product package
//
// swagger:model CartProductPackage
type CartProductPackage struct {

	// Id
	ID int32 `json:"Id,omitempty"`

	// line items
	LineItems []*CartLineItem `json:"LineItems"`
}

// Validate validates this cart product package
func (m *CartProductPackage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLineItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CartProductPackage) validateLineItems(formats strfmt.Registry) error {
	if swag.IsZero(m.LineItems) { // not required
		return nil
	}

	for i := 0; i < len(m.LineItems); i++ {
		if swag.IsZero(m.LineItems[i]) { // not required
			continue
		}

		if m.LineItems[i] != nil {
			if err := m.LineItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("LineItems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("LineItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cart product package based on the context it is used
func (m *CartProductPackage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLineItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CartProductPackage) contextValidateLineItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LineItems); i++ {

		if m.LineItems[i] != nil {

			if swag.IsZero(m.LineItems[i]) { // not required
				return nil
			}

			if err := m.LineItems[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("LineItems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("LineItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CartProductPackage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CartProductPackage) UnmarshalBinary(b []byte) error {
	var res CartProductPackage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
