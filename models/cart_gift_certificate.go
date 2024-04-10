// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CartGiftCertificate cart gift certificate
//
// swagger:model CartGiftCertificate
type CartGiftCertificate struct {

	// amount
	Amount float64 `json:"Amount,omitempty"`

	// applied
	Applied bool `json:"Applied,omitempty"`

	// gift certificate number
	GiftCertificateNumber string `json:"GiftCertificateNumber,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// notes
	Notes string `json:"Notes,omitempty"`

	// payment Id
	PaymentID int32 `json:"PaymentId,omitempty"`
}

// Validate validates this cart gift certificate
func (m *CartGiftCertificate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cart gift certificate based on context it is used
func (m *CartGiftCertificate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CartGiftCertificate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CartGiftCertificate) UnmarshalBinary(b []byte) error {
	var res CartGiftCertificate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
