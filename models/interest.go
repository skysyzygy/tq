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

// Interest interest
//
// swagger:model Interest
type Interest struct {

	// constituent
	Constituent *Entity `json:"Constituent,omitempty"`

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime *strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// edit indicator
	EditIndicator bool `json:"EditIndicator,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// interest type
	InterestType *InterestTypeSummary `json:"InterestType,omitempty"`

	// selected
	Selected bool `json:"Selected,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime *strfmt.DateTime `json:"UpdatedDateTime,omitempty"`

	// weight
	Weight int32 `json:"Weight,omitempty"`
}

// Validate validates this interest
func (m *Interest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConstituent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterestType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Interest) validateConstituent(formats strfmt.Registry) error {
	if swag.IsZero(m.Constituent) { // not required
		return nil
	}

	if m.Constituent != nil {
		if err := m.Constituent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Constituent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Constituent")
			}
			return err
		}
	}

	return nil
}

func (m *Interest) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Interest) validateInterestType(formats strfmt.Registry) error {
	if swag.IsZero(m.InterestType) { // not required
		return nil
	}

	if m.InterestType != nil {
		if err := m.InterestType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("InterestType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("InterestType")
			}
			return err
		}
	}

	return nil
}

func (m *Interest) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this interest based on the context it is used
func (m *Interest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConstituent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterestType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Interest) contextValidateConstituent(ctx context.Context, formats strfmt.Registry) error {

	if m.Constituent != nil {

		if swag.IsZero(m.Constituent) { // not required
			return nil
		}

		if err := m.Constituent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Constituent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Constituent")
			}
			return err
		}
	}

	return nil
}

func (m *Interest) contextValidateInterestType(ctx context.Context, formats strfmt.Registry) error {

	if m.InterestType != nil {

		if swag.IsZero(m.InterestType) { // not required
			return nil
		}

		if err := m.InterestType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("InterestType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("InterestType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Interest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Interest) UnmarshalBinary(b []byte) error {
	var res Interest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
