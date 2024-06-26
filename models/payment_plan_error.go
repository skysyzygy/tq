// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PaymentPlanError payment plan error
//
// swagger:model PaymentPlanError
type PaymentPlanError struct {

	// error
	Error string `json:"Error,omitempty"`

	// order Id
	OrderID int32 `json:"OrderId,omitempty"`
}

// Validate validates this payment plan error
func (m *PaymentPlanError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this payment plan error based on context it is used
func (m *PaymentPlanError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentPlanError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentPlanError) UnmarshalBinary(b []byte) error {
	var res PaymentPlanError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
