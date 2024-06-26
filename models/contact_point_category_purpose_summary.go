// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContactPointCategoryPurposeSummary contact point category purpose summary
//
// swagger:model ContactPointCategoryPurposeSummary
type ContactPointCategoryPurposeSummary struct {

	// Id
	ID int32 `json:"Id,omitempty"`

	// purpose
	Purpose *ContactPointPurposeSummary `json:"Purpose,omitempty"`
}

// Validate validates this contact point category purpose summary
func (m *ContactPointCategoryPurposeSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePurpose(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactPointCategoryPurposeSummary) validatePurpose(formats strfmt.Registry) error {
	if swag.IsZero(m.Purpose) { // not required
		return nil
	}

	if m.Purpose != nil {
		if err := m.Purpose.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Purpose")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Purpose")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this contact point category purpose summary based on the context it is used
func (m *ContactPointCategoryPurposeSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePurpose(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactPointCategoryPurposeSummary) contextValidatePurpose(ctx context.Context, formats strfmt.Registry) error {

	if m.Purpose != nil {

		if swag.IsZero(m.Purpose) { // not required
			return nil
		}

		if err := m.Purpose.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Purpose")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Purpose")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactPointCategoryPurposeSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactPointCategoryPurposeSummary) UnmarshalBinary(b []byte) error {
	var res ContactPointCategoryPurposeSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
