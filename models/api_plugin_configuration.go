// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIPluginConfiguration Api plugin configuration
//
// swagger:model ApiPluginConfiguration
type APIPluginConfiguration struct {

	// enabled
	Enabled bool `json:"Enabled,omitempty"`

	// registrations
	Registrations []*APIPluginRegistration `json:"Registrations"`

	// reload always
	ReloadAlways bool `json:"ReloadAlways,omitempty"`
}

// Validate validates this Api plugin configuration
func (m *APIPluginConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegistrations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIPluginConfiguration) validateRegistrations(formats strfmt.Registry) error {
	if swag.IsZero(m.Registrations) { // not required
		return nil
	}

	for i := 0; i < len(m.Registrations); i++ {
		if swag.IsZero(m.Registrations[i]) { // not required
			continue
		}

		if m.Registrations[i] != nil {
			if err := m.Registrations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Registrations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Registrations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this Api plugin configuration based on the context it is used
func (m *APIPluginConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRegistrations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIPluginConfiguration) contextValidateRegistrations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Registrations); i++ {

		if m.Registrations[i] != nil {

			if swag.IsZero(m.Registrations[i]) { // not required
				return nil
			}

			if err := m.Registrations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Registrations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Registrations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIPluginConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIPluginConfiguration) UnmarshalBinary(b []byte) error {
	var res APIPluginConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
