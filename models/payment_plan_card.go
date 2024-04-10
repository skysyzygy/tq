// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PaymentPlanCard payment plan card
//
// swagger:model PaymentPlanCard
type PaymentPlanCard struct {

	// expiry month
	ExpiryMonth int32 `json:"ExpiryMonth,omitempty"`

	// expiry year
	ExpiryYear int32 `json:"ExpiryYear,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// number
	Number string `json:"Number,omitempty"`

	// payment method group Id
	PaymentMethodGroupID int32 `json:"PaymentMethodGroupId,omitempty"`
}

// Validate validates this payment plan card
func (m *PaymentPlanCard) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this payment plan card based on context it is used
func (m *PaymentPlanCard) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentPlanCard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentPlanCard) UnmarshalBinary(b []byte) error {
	var res PaymentPlanCard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
