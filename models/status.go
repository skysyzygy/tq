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
	"github.com/go-openapi/validate"
)

// Status status
//
// swagger:model Status
type Status struct {

	// dependencies
	Dependencies []*Dependency `json:"Dependencies"`

	// success
	Success bool `json:"Success,omitempty"`

	// time stamp
	// Format: date-time
	TimeStamp *strfmt.DateTime `json:"TimeStamp,omitempty"`

	// time zone
	TimeZone string `json:"TimeZone,omitempty"`
}

// Validate validates this status
func (m *Status) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDependencies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeStamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Status) validateDependencies(formats strfmt.Registry) error {
	if swag.IsZero(m.Dependencies) { // not required
		return nil
	}

	for i := 0; i < len(m.Dependencies); i++ {
		if swag.IsZero(m.Dependencies[i]) { // not required
			continue
		}

		if m.Dependencies[i] != nil {
			if err := m.Dependencies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Dependencies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Dependencies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Status) validateTimeStamp(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeStamp) { // not required
		return nil
	}

	if err := validate.FormatOf("TimeStamp", "body", "date-time", m.TimeStamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this status based on the context it is used
func (m *Status) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDependencies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Status) contextValidateDependencies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Dependencies); i++ {

		if m.Dependencies[i] != nil {

			if swag.IsZero(m.Dependencies[i]) { // not required
				return nil
			}

			if err := m.Dependencies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Dependencies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Dependencies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Status) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Status) UnmarshalBinary(b []byte) error {
	var res Status
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
