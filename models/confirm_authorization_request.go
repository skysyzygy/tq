// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConfirmAuthorizationRequest confirm authorization request
//
// swagger:model ConfirmAuthorizationRequest
type ConfirmAuthorizationRequest struct {

	// amount
	Amount float64 `json:"Amount,omitempty"`

	// is e commerce
	IsECommerce bool `json:"IsECommerce,omitempty"`

	// payment Id
	PaymentID int32 `json:"PaymentId,omitempty"`

	// payment method Id
	PaymentMethodID int32 `json:"PaymentMethodId,omitempty"`

	// transaction origin
	TransactionOrigin string `json:"TransactionOrigin,omitempty"`

	// user data
	UserData string `json:"UserData,omitempty"`
}

// Validate validates this confirm authorization request
func (m *ConfirmAuthorizationRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this confirm authorization request based on context it is used
func (m *ConfirmAuthorizationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConfirmAuthorizationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfirmAuthorizationRequest) UnmarshalBinary(b []byte) error {
	var res ConfirmAuthorizationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
