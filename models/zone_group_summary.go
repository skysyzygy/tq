// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ZoneGroupSummary zone group summary
//
// swagger:model ZoneGroupSummary
type ZoneGroupSummary struct {

	// alias description
	AliasDescription string `json:"AliasDescription,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// inactive
	Inactive bool `json:"Inactive,omitempty"`

	// rank
	Rank int32 `json:"Rank,omitempty"`
}

// Validate validates this zone group summary
func (m *ZoneGroupSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this zone group summary based on context it is used
func (m *ZoneGroupSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ZoneGroupSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ZoneGroupSummary) UnmarshalBinary(b []byte) error {
	var res ZoneGroupSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
