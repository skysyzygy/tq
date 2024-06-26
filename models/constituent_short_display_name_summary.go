// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConstituentShortDisplayNameSummary constituent short display name summary
//
// swagger:model ConstituentShortDisplayNameSummary
type ConstituentShortDisplayNameSummary struct {

	// Id
	ID int32 `json:"Id,omitempty"`

	// short display name
	ShortDisplayName string `json:"ShortDisplayName,omitempty"`

	// sort name
	SortName string `json:"SortName,omitempty"`
}

// Validate validates this constituent short display name summary
func (m *ConstituentShortDisplayNameSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this constituent short display name summary based on context it is used
func (m *ConstituentShortDisplayNameSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConstituentShortDisplayNameSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConstituentShortDisplayNameSummary) UnmarshalBinary(b []byte) error {
	var res ConstituentShortDisplayNameSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
