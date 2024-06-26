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

// ParameterValuesRequest parameter values request
//
// swagger:model ParameterValuesRequest
type ParameterValuesRequest struct {

	// parameter name
	ParameterName string `json:"ParameterName,omitempty"`

	// report Id
	ReportID string `json:"ReportId,omitempty"`

	// where dependencies
	WhereDependencies []*ParameterValueWhereDependencies `json:"WhereDependencies"`
}

// Validate validates this parameter values request
func (m *ParameterValuesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWhereDependencies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ParameterValuesRequest) validateWhereDependencies(formats strfmt.Registry) error {
	if swag.IsZero(m.WhereDependencies) { // not required
		return nil
	}

	for i := 0; i < len(m.WhereDependencies); i++ {
		if swag.IsZero(m.WhereDependencies[i]) { // not required
			continue
		}

		if m.WhereDependencies[i] != nil {
			if err := m.WhereDependencies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("WhereDependencies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("WhereDependencies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this parameter values request based on the context it is used
func (m *ParameterValuesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWhereDependencies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ParameterValuesRequest) contextValidateWhereDependencies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.WhereDependencies); i++ {

		if m.WhereDependencies[i] != nil {

			if swag.IsZero(m.WhereDependencies[i]) { // not required
				return nil
			}

			if err := m.WhereDependencies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("WhereDependencies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("WhereDependencies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ParameterValuesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ParameterValuesRequest) UnmarshalBinary(b []byte) error {
	var res ParameterValuesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
