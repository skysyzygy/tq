// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SalesChannelSummary sales channel summary
//
// swagger:model SalesChannelSummary
type SalesChannelSummary struct {

	// description
	Description string `json:"Description,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// inactive
	Inactive bool `json:"Inactive,omitempty"`
}

// Validate validates this sales channel summary
func (m *SalesChannelSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sales channel summary based on context it is used
func (m *SalesChannelSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SalesChannelSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SalesChannelSummary) UnmarshalBinary(b []byte) error {
	var res SalesChannelSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
