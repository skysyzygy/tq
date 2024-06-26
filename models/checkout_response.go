// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CheckoutResponse checkout response
//
// swagger:model CheckoutResponse
type CheckoutResponse struct {

	// message
	Message string `json:"Message,omitempty"`

	// status
	Status string `json:"Status,omitempty"`
}

// Validate validates this checkout response
func (m *CheckoutResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this checkout response based on context it is used
func (m *CheckoutResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CheckoutResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckoutResponse) UnmarshalBinary(b []byte) error {
	var res CheckoutResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
