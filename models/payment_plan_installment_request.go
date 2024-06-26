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

// PaymentPlanInstallmentRequest payment plan installment request
//
// swagger:model PaymentPlanInstallmentRequest
type PaymentPlanInstallmentRequest struct {

	// account Id
	AccountID int32 `json:"AccountId,omitempty"`

	// billing type Id
	BillingTypeID int32 `json:"BillingTypeId,omitempty"`

	// card
	Card *PaymentPlanCard `json:"Card,omitempty"`

	// payment plan installments
	PaymentPlanInstallments []*PaymentPlanInstallment `json:"PaymentPlanInstallments"`
}

// Validate validates this payment plan installment request
func (m *PaymentPlanInstallmentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentPlanInstallments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentPlanInstallmentRequest) validateCard(formats strfmt.Registry) error {
	if swag.IsZero(m.Card) { // not required
		return nil
	}

	if m.Card != nil {
		if err := m.Card.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Card")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Card")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentPlanInstallmentRequest) validatePaymentPlanInstallments(formats strfmt.Registry) error {
	if swag.IsZero(m.PaymentPlanInstallments) { // not required
		return nil
	}

	for i := 0; i < len(m.PaymentPlanInstallments); i++ {
		if swag.IsZero(m.PaymentPlanInstallments[i]) { // not required
			continue
		}

		if m.PaymentPlanInstallments[i] != nil {
			if err := m.PaymentPlanInstallments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PaymentPlanInstallments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("PaymentPlanInstallments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this payment plan installment request based on the context it is used
func (m *PaymentPlanInstallmentRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePaymentPlanInstallments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentPlanInstallmentRequest) contextValidateCard(ctx context.Context, formats strfmt.Registry) error {

	if m.Card != nil {

		if swag.IsZero(m.Card) { // not required
			return nil
		}

		if err := m.Card.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Card")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Card")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentPlanInstallmentRequest) contextValidatePaymentPlanInstallments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PaymentPlanInstallments); i++ {

		if m.PaymentPlanInstallments[i] != nil {

			if swag.IsZero(m.PaymentPlanInstallments[i]) { // not required
				return nil
			}

			if err := m.PaymentPlanInstallments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PaymentPlanInstallments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("PaymentPlanInstallments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentPlanInstallmentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentPlanInstallmentRequest) UnmarshalBinary(b []byte) error {
	var res PaymentPlanInstallmentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
