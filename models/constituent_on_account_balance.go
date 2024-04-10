// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConstituentOnAccountBalance constituent on account balance
//
// swagger:model ConstituentOnAccountBalance
type ConstituentOnAccountBalance struct {

	// amount
	Amount float64 `json:"Amount,omitempty"`

	// payment method Id
	PaymentMethodID int32 `json:"PaymentMethodId,omitempty"`
}

// Validate validates this constituent on account balance
func (m *ConstituentOnAccountBalance) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this constituent on account balance based on context it is used
func (m *ConstituentOnAccountBalance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConstituentOnAccountBalance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConstituentOnAccountBalance) UnmarshalBinary(b []byte) error {
	var res ConstituentOnAccountBalance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
