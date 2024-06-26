// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CartZone cart zone
//
// swagger:model CartZone
type CartZone struct {

	// abbreviation
	Abbreviation string `json:"Abbreviation,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// rank
	Rank int32 `json:"Rank,omitempty"`

	// short description
	ShortDescription string `json:"ShortDescription,omitempty"`

	// zone group Id
	ZoneGroupID int32 `json:"ZoneGroupId,omitempty"`

	// zone time
	ZoneTime string `json:"ZoneTime,omitempty"`
}

// Validate validates this cart zone
func (m *CartZone) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cart zone based on context it is used
func (m *CartZone) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CartZone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CartZone) UnmarshalBinary(b []byte) error {
	var res CartZone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
