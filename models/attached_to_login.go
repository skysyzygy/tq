// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AttachedToLogin attached to login
//
// swagger:model AttachedToLogin
type AttachedToLogin struct {

	// is attached to login
	IsAttachedToLogin bool `json:"IsAttachedToLogin,omitempty"`
}

// Validate validates this attached to login
func (m *AttachedToLogin) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this attached to login based on context it is used
func (m *AttachedToLogin) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AttachedToLogin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AttachedToLogin) UnmarshalBinary(b []byte) error {
	var res AttachedToLogin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
