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

// OrderProductViewPayment order product view payment
//
// swagger:model OrderProductViewPayment
type OrderProductViewPayment struct {

	// amount
	Amount float64 `json:"Amount,omitempty"`

	// check number
	CheckNumber string `json:"CheckNumber,omitempty"`

	// gift certificate number
	GiftCertificateNumber string `json:"GiftCertificateNumber,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// last four credit card number
	LastFourCreditCardNumber string `json:"LastFourCreditCardNumber,omitempty"`

	// notes
	Notes string `json:"Notes,omitempty"`

	// payer name
	PayerName string `json:"PayerName,omitempty"`

	// payment method
	PaymentMethod *OrderProductViewPaymentMethod `json:"PaymentMethod,omitempty"`

	// tendered amount
	TenderedAmount float64 `json:"TenderedAmount,omitempty"`
}

// Validate validates this order product view payment
func (m *OrderProductViewPayment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderProductViewPayment) validatePaymentMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.PaymentMethod) { // not required
		return nil
	}

	if m.PaymentMethod != nil {
		if err := m.PaymentMethod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PaymentMethod")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PaymentMethod")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this order product view payment based on the context it is used
func (m *OrderProductViewPayment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePaymentMethod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderProductViewPayment) contextValidatePaymentMethod(ctx context.Context, formats strfmt.Registry) error {

	if m.PaymentMethod != nil {

		if swag.IsZero(m.PaymentMethod) { // not required
			return nil
		}

		if err := m.PaymentMethod.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PaymentMethod")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PaymentMethod")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderProductViewPayment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderProductViewPayment) UnmarshalBinary(b []byte) error {
	var res OrderProductViewPayment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
