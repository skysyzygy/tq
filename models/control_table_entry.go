// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ControlTableEntry control table entry
//
// swagger:model ControlTableEntry
type ControlTableEntry struct {

	// Id
	ID int32 `json:"Id,omitempty"`
}

// Validate validates this control table entry
func (m *ControlTableEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this control table entry based on context it is used
func (m *ControlTableEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ControlTableEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ControlTableEntry) UnmarshalBinary(b []byte) error {
	var res ControlTableEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
