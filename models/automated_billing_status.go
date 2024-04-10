// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AutomatedBillingStatus automated billing status
//
// swagger:model AutomatedBillingStatus
type AutomatedBillingStatus struct {

	// bill count
	BillCount int32 `json:"BillCount,omitempty"`

	// has errors
	HasErrors bool `json:"HasErrors,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// is complete
	IsComplete bool `json:"IsComplete,omitempty"`

	// message
	Message string `json:"Message,omitempty"`

	// processed count
	ProcessedCount int32 `json:"ProcessedCount,omitempty"`
}

// Validate validates this automated billing status
func (m *AutomatedBillingStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this automated billing status based on context it is used
func (m *AutomatedBillingStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AutomatedBillingStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutomatedBillingStatus) UnmarshalBinary(b []byte) error {
	var res AutomatedBillingStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
