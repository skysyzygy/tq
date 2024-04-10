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

// CreditResult credit result
//
// swagger:model CreditResult
type CreditResult struct {

	// credits
	Credits []*Credit `json:"Credits"`

	// requested owner
	RequestedOwner *Owner `json:"RequestedOwner,omitempty"`
}

// Validate validates this credit result
func (m *CreditResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedOwner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreditResult) validateCredits(formats strfmt.Registry) error {
	if swag.IsZero(m.Credits) { // not required
		return nil
	}

	for i := 0; i < len(m.Credits); i++ {
		if swag.IsZero(m.Credits[i]) { // not required
			continue
		}

		if m.Credits[i] != nil {
			if err := m.Credits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Credits" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Credits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreditResult) validateRequestedOwner(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestedOwner) { // not required
		return nil
	}

	if m.RequestedOwner != nil {
		if err := m.RequestedOwner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RequestedOwner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("RequestedOwner")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this credit result based on the context it is used
func (m *CreditResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCredits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequestedOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreditResult) contextValidateCredits(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Credits); i++ {

		if m.Credits[i] != nil {

			if swag.IsZero(m.Credits[i]) { // not required
				return nil
			}

			if err := m.Credits[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Credits" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Credits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreditResult) contextValidateRequestedOwner(ctx context.Context, formats strfmt.Registry) error {

	if m.RequestedOwner != nil {

		if swag.IsZero(m.RequestedOwner) { // not required
			return nil
		}

		if err := m.RequestedOwner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RequestedOwner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("RequestedOwner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreditResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreditResult) UnmarshalBinary(b []byte) error {
	var res CreditResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
