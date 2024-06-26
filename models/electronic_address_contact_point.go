// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ElectronicAddressContactPoint electronic address contact point
//
// swagger:model ElectronicAddressContactPoint
type ElectronicAddressContactPoint struct {

	// electronic address
	ElectronicAddress string `json:"ElectronicAddress,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// inactive
	Inactive bool `json:"Inactive,omitempty"`
}

// Validate validates this electronic address contact point
func (m *ElectronicAddressContactPoint) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this electronic address contact point based on context it is used
func (m *ElectronicAddressContactPoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ElectronicAddressContactPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElectronicAddressContactPoint) UnmarshalBinary(b []byte) error {
	var res ElectronicAddressContactPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
