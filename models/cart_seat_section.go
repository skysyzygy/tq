// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CartSeatSection cart seat section
//
// swagger:model CartSeatSection
type CartSeatSection struct {

	// description
	Description string `json:"Description,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// print description
	PrintDescription string `json:"PrintDescription,omitempty"`

	// short description
	ShortDescription string `json:"ShortDescription,omitempty"`
}

// Validate validates this cart seat section
func (m *CartSeatSection) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cart seat section based on context it is used
func (m *CartSeatSection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CartSeatSection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CartSeatSection) UnmarshalBinary(b []byte) error {
	var res CartSeatSection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
