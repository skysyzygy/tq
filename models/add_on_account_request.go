// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddOnAccountRequest add on account request
//
// swagger:model AddOnAccountRequest
type AddOnAccountRequest struct {

	// amount
	Amount float64 `json:"Amount,omitempty"`

	// notes
	Notes string `json:"Notes,omitempty"`

	// payment method Id
	PaymentMethodID int32 `json:"PaymentMethodId,omitempty"`
}

// Validate validates this add on account request
func (m *AddOnAccountRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add on account request based on context it is used
func (m *AddOnAccountRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AddOnAccountRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddOnAccountRequest) UnmarshalBinary(b []byte) error {
	var res AddOnAccountRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
