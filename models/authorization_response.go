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

// AuthorizationResponse authorization response
//
// swagger:model AuthorizationResponse
type AuthorizationResponse struct {

	// a v s response code
	AVSResponseCode string `json:"AVSResponseCode,omitempty"`

	// account Id
	AccountID int32 `json:"AccountId,omitempty"`

	// authorization code
	AuthorizationCode string `json:"AuthorizationCode,omitempty"`

	// authorized amount
	AuthorizedAmount float64 `json:"AuthorizedAmount,omitempty"`

	// balance amount
	BalanceAmount float64 `json:"BalanceAmount,omitempty"`

	// balance currency code
	BalanceCurrencyCode string `json:"BalanceCurrencyCode,omitempty"`

	// c v v response code
	CVVResponseCode string `json:"CVVResponseCode,omitempty"`

	// is partial auth
	IsPartialAuth bool `json:"IsPartialAuth,omitempty"`

	// masked card number
	MaskedCardNumber string `json:"MaskedCardNumber,omitempty"`

	// payment method
	PaymentMethod *PaymentMethod `json:"PaymentMethod,omitempty"`

	// reference number
	ReferenceNumber string `json:"ReferenceNumber,omitempty"`

	// response code
	ResponseCode string `json:"ResponseCode,omitempty"`

	// response message
	ResponseMessage string `json:"ResponseMessage,omitempty"`

	// shopper Ip
	ShopperIP string `json:"ShopperIp,omitempty"`

	// succeeded
	Succeeded bool `json:"Succeeded,omitempty"`

	// tessitura merchant services action
	TessituraMerchantServicesAction string `json:"TessituraMerchantServicesAction,omitempty"`
}

// Validate validates this authorization response
func (m *AuthorizationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthorizationResponse) validatePaymentMethod(formats strfmt.Registry) error {
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

// ContextValidate validate this authorization response based on the context it is used
func (m *AuthorizationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePaymentMethod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthorizationResponse) contextValidatePaymentMethod(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AuthorizationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthorizationResponse) UnmarshalBinary(b []byte) error {
	var res AuthorizationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
