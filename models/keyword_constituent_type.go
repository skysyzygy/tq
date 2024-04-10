// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KeywordConstituentType keyword constituent type
//
// swagger:model KeywordConstituentType
type KeywordConstituentType struct {

	// inactive
	Inactive bool `json:"Inactive,omitempty"`

	// rank
	Rank int32 `json:"Rank,omitempty"`
}

// Validate validates this keyword constituent type
func (m *KeywordConstituentType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this keyword constituent type based on context it is used
func (m *KeywordConstituentType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KeywordConstituentType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeywordConstituentType) UnmarshalBinary(b []byte) error {
	var res KeywordConstituentType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
