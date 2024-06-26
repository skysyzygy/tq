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

// CardData card data
//
// swagger:model CardData
type CardData struct {

	// account Id
	AccountID int32 `json:"AccountId,omitempty"`

	// card holder name
	CardHolderName string `json:"CardHolderName,omitempty"`

	// card number
	CardNumber string `json:"CardNumber,omitempty"`

	// cvv2
	Cvv2 string `json:"Cvv2,omitempty"`

	// expiration date
	ExpirationDate string `json:"ExpirationDate,omitempty"`

	// tessitura merchant services data
	TessituraMerchantServicesData string `json:"TessituraMerchantServicesData,omitempty"`

	// track1
	Track1 string `json:"Track1,omitempty"`

	// track2
	Track2 string `json:"Track2,omitempty"`

	// vantiv encrypted card
	VantivEncryptedCard *VantivEncryptedCard `json:"VantivEncryptedCard,omitempty"`
}

// Validate validates this card data
func (m *CardData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVantivEncryptedCard(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CardData) validateVantivEncryptedCard(formats strfmt.Registry) error {
	if swag.IsZero(m.VantivEncryptedCard) { // not required
		return nil
	}

	if m.VantivEncryptedCard != nil {
		if err := m.VantivEncryptedCard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("VantivEncryptedCard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("VantivEncryptedCard")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this card data based on the context it is used
func (m *CardData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVantivEncryptedCard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CardData) contextValidateVantivEncryptedCard(ctx context.Context, formats strfmt.Registry) error {

	if m.VantivEncryptedCard != nil {

		if swag.IsZero(m.VantivEncryptedCard) { // not required
			return nil
		}

		if err := m.VantivEncryptedCard.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("VantivEncryptedCard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("VantivEncryptedCard")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CardData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CardData) UnmarshalBinary(b []byte) error {
	var res CardData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
