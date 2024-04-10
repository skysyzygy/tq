// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApplyInvoicePaymentRequest apply invoice payment request
//
// swagger:model ApplyInvoicePaymentRequest
type ApplyInvoicePaymentRequest struct {

	// amount
	Amount float64 `json:"Amount,omitempty"`

	// notes
	Notes string `json:"Notes,omitempty"`

	// payment method Id
	PaymentMethodID int32 `json:"PaymentMethodId,omitempty"`
}

// Validate validates this apply invoice payment request
func (m *ApplyInvoicePaymentRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this apply invoice payment request based on context it is used
func (m *ApplyInvoicePaymentRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApplyInvoicePaymentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplyInvoicePaymentRequest) UnmarshalBinary(b []byte) error {
	var res ApplyInvoicePaymentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
