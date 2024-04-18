// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GiftCertificateRedemptionRequest gift certificate redemption request
//
// swagger:model GiftCertificateRedemptionRequest
type GiftCertificateRedemptionRequest struct {

	// batch Id
	BatchID int32 `json:"BatchId,omitempty"`

	// lock for batch
	LockForBatch bool `json:"LockForBatch,omitempty"`

	// number
	Number string `json:"Number,omitempty"`

	// payment method Id
	PaymentMethodID int32 `json:"PaymentMethodId,omitempty"`
}

// Validate validates this gift certificate redemption request
func (m *GiftCertificateRedemptionRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this gift certificate redemption request based on context it is used
func (m *GiftCertificateRedemptionRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GiftCertificateRedemptionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GiftCertificateRedemptionRequest) UnmarshalBinary(b []byte) error {
	var res GiftCertificateRedemptionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}