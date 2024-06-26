// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Entity entity
//
// swagger:model Entity
type Entity struct {

	// Id
	ID int32 `json:"Id,omitempty"`
}

// Validate validates this entity
func (m *Entity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this entity based on context it is used
func (m *Entity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Entity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Entity) UnmarshalBinary(b []byte) error {
	var res Entity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
