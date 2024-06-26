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

// VantivEncryptedCard vantiv encrypted card
//
// swagger:model VantivEncryptedCard
type VantivEncryptedCard struct {

	// version
	Version string `json:"Version,omitempty"`

	// vantiv addr
	VantivAddr *VantivAddr `json:"vantivAddr,omitempty"`

	// vantiv card
	VantivCard *VantivCard `json:"vantivCard,omitempty"`

	// vantiv device
	VantivDevice *VantivDevice `json:"vantivDevice,omitempty"`

	// vantiv tran
	VantivTran *VantivTran `json:"vantivTran,omitempty"`
}

// Validate validates this vantiv encrypted card
func (m *VantivEncryptedCard) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVantivAddr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVantivCard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVantivDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVantivTran(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VantivEncryptedCard) validateVantivAddr(formats strfmt.Registry) error {
	if swag.IsZero(m.VantivAddr) { // not required
		return nil
	}

	if m.VantivAddr != nil {
		if err := m.VantivAddr.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vantivAddr")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vantivAddr")
			}
			return err
		}
	}

	return nil
}

func (m *VantivEncryptedCard) validateVantivCard(formats strfmt.Registry) error {
	if swag.IsZero(m.VantivCard) { // not required
		return nil
	}

	if m.VantivCard != nil {
		if err := m.VantivCard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vantivCard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vantivCard")
			}
			return err
		}
	}

	return nil
}

func (m *VantivEncryptedCard) validateVantivDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.VantivDevice) { // not required
		return nil
	}

	if m.VantivDevice != nil {
		if err := m.VantivDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vantivDevice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vantivDevice")
			}
			return err
		}
	}

	return nil
}

func (m *VantivEncryptedCard) validateVantivTran(formats strfmt.Registry) error {
	if swag.IsZero(m.VantivTran) { // not required
		return nil
	}

	if m.VantivTran != nil {
		if err := m.VantivTran.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vantivTran")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vantivTran")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this vantiv encrypted card based on the context it is used
func (m *VantivEncryptedCard) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVantivAddr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVantivCard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVantivDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVantivTran(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VantivEncryptedCard) contextValidateVantivAddr(ctx context.Context, formats strfmt.Registry) error {

	if m.VantivAddr != nil {

		if swag.IsZero(m.VantivAddr) { // not required
			return nil
		}

		if err := m.VantivAddr.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vantivAddr")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vantivAddr")
			}
			return err
		}
	}

	return nil
}

func (m *VantivEncryptedCard) contextValidateVantivCard(ctx context.Context, formats strfmt.Registry) error {

	if m.VantivCard != nil {

		if swag.IsZero(m.VantivCard) { // not required
			return nil
		}

		if err := m.VantivCard.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vantivCard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vantivCard")
			}
			return err
		}
	}

	return nil
}

func (m *VantivEncryptedCard) contextValidateVantivDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.VantivDevice != nil {

		if swag.IsZero(m.VantivDevice) { // not required
			return nil
		}

		if err := m.VantivDevice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vantivDevice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vantivDevice")
			}
			return err
		}
	}

	return nil
}

func (m *VantivEncryptedCard) contextValidateVantivTran(ctx context.Context, formats strfmt.Registry) error {

	if m.VantivTran != nil {

		if swag.IsZero(m.VantivTran) { // not required
			return nil
		}

		if err := m.VantivTran.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vantivTran")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vantivTran")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VantivEncryptedCard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VantivEncryptedCard) UnmarshalBinary(b []byte) error {
	var res VantivEncryptedCard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
