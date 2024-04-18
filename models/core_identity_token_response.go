// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CoreIdentityTokenResponse core identity token response
//
// swagger:model CoreIdentityTokenResponse
type CoreIdentityTokenResponse struct {

	// environment
	Environment string `json:"Environment,omitempty"`

	// Id token
	IDToken string `json:"IdToken,omitempty"`

	// licensee
	Licensee string `json:"Licensee,omitempty"`
}

// Validate validates this core identity token response
func (m *CoreIdentityTokenResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this core identity token response based on context it is used
func (m *CoreIdentityTokenResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CoreIdentityTokenResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CoreIdentityTokenResponse) UnmarshalBinary(b []byte) error {
	var res CoreIdentityTokenResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}