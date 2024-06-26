// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIPluginRegistration Api plugin registration
//
// swagger:model ApiPluginRegistration
type APIPluginRegistration struct {

	// full plugin name
	FullPluginName string `json:"FullPluginName,omitempty"`

	// plugin placement
	PluginPlacement string `json:"PluginPlacement,omitempty"`

	// Uri match
	URIMatch string `json:"UriMatch,omitempty"`

	// verb
	Verb string `json:"Verb,omitempty"`
}

// Validate validates this Api plugin registration
func (m *APIPluginRegistration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Api plugin registration based on context it is used
func (m *APIPluginRegistration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIPluginRegistration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIPluginRegistration) UnmarshalBinary(b []byte) error {
	var res APIPluginRegistration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
