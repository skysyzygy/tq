// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConnectionStatus connection status
//
// swagger:model ConnectionStatus
type ConnectionStatus struct {

	// status
	Status string `json:"Status,omitempty"`

	// time stamp
	// Format: date-time
	TimeStamp strfmt.DateTime `json:"TimeStamp,omitempty"`
}

// Validate validates this connection status
func (m *ConnectionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimeStamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConnectionStatus) validateTimeStamp(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeStamp) { // not required
		return nil
	}

	if err := validate.FormatOf("TimeStamp", "body", "date-time", m.TimeStamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this connection status based on context it is used
func (m *ConnectionStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConnectionStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectionStatus) UnmarshalBinary(b []byte) error {
	var res ConnectionStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
