// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpgradeCategorySummary upgrade category summary
//
// swagger:model UpgradeCategorySummary
type UpgradeCategorySummary struct {

	// description
	Description string `json:"Description,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`
}

// Validate validates this upgrade category summary
func (m *UpgradeCategorySummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this upgrade category summary based on context it is used
func (m *UpgradeCategorySummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpgradeCategorySummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpgradeCategorySummary) UnmarshalBinary(b []byte) error {
	var res UpgradeCategorySummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
