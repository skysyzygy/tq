// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EMVSignatureRequest e m v signature request
//
// swagger:model EMVSignatureRequest
type EMVSignatureRequest struct {

	// machine Id
	MachineID int32 `json:"MachineId,omitempty"`

	// payment Id
	PaymentID int32 `json:"PaymentId,omitempty"`

	// signature matches
	SignatureMatches bool `json:"SignatureMatches,omitempty"`

	// store account
	StoreAccount bool `json:"StoreAccount,omitempty"`
}

// Validate validates this e m v signature request
func (m *EMVSignatureRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this e m v signature request based on context it is used
func (m *EMVSignatureRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EMVSignatureRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EMVSignatureRequest) UnmarshalBinary(b []byte) error {
	var res EMVSignatureRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
