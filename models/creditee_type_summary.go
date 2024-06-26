// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CrediteeTypeSummary creditee type summary
//
// swagger:model CrediteeTypeSummary
type CrediteeTypeSummary struct {

	// description
	Description string `json:"Description,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`
}

// Validate validates this creditee type summary
func (m *CrediteeTypeSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this creditee type summary based on context it is used
func (m *CrediteeTypeSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CrediteeTypeSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CrediteeTypeSummary) UnmarshalBinary(b []byte) error {
	var res CrediteeTypeSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
