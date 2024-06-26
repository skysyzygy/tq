// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CartContributionSummary cart contribution summary
//
// swagger:model CartContributionSummary
type CartContributionSummary struct {

	// amount
	Amount float64 `json:"Amount,omitempty"`

	// fund description
	FundDescription string `json:"FundDescription,omitempty"`

	// reference Id
	ReferenceID int32 `json:"ReferenceId,omitempty"`
}

// Validate validates this cart contribution summary
func (m *CartContributionSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cart contribution summary based on context it is used
func (m *CartContributionSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CartContributionSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CartContributionSummary) UnmarshalBinary(b []byte) error {
	var res CartContributionSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
