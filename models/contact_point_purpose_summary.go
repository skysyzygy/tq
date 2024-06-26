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

// ContactPointPurposeSummary contact point purpose summary
//
// swagger:model ContactPointPurposeSummary
type ContactPointPurposeSummary struct {

	// description
	Description string `json:"Description,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// inactive
	Inactive bool `json:"Inactive,omitempty"`

	// purpose category
	PurposeCategory *PurposeCategorySummary `json:"PurposeCategory,omitempty"`
}

// Validate validates this contact point purpose summary
func (m *ContactPointPurposeSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePurposeCategory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactPointPurposeSummary) validatePurposeCategory(formats strfmt.Registry) error {
	if swag.IsZero(m.PurposeCategory) { // not required
		return nil
	}

	if m.PurposeCategory != nil {
		if err := m.PurposeCategory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PurposeCategory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PurposeCategory")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this contact point purpose summary based on the context it is used
func (m *ContactPointPurposeSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePurposeCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactPointPurposeSummary) contextValidatePurposeCategory(ctx context.Context, formats strfmt.Registry) error {

	if m.PurposeCategory != nil {

		if swag.IsZero(m.PurposeCategory) { // not required
			return nil
		}

		if err := m.PurposeCategory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PurposeCategory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PurposeCategory")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactPointPurposeSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactPointPurposeSummary) UnmarshalBinary(b []byte) error {
	var res ContactPointPurposeSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
