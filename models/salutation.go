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

// Salutation salutation
//
// swagger:model Salutation
type Salutation struct {

	// business title
	BusinessTitle string `json:"BusinessTitle,omitempty"`

	// constituent
	Constituent *Entity `json:"Constituent,omitempty"`

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// default indicator
	DefaultIndicator bool `json:"DefaultIndicator,omitempty"`

	// edit indicator
	EditIndicator bool `json:"EditIndicator,omitempty"`

	// envelope salutation1
	EnvelopeSalutation1 string `json:"EnvelopeSalutation1,omitempty"`

	// envelope salutation2
	EnvelopeSalutation2 string `json:"EnvelopeSalutation2,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// is from affiliation
	IsFromAffiliation bool `json:"IsFromAffiliation,omitempty"`

	// label
	Label bool `json:"Label,omitempty"`

	// letter salutation
	LetterSalutation string `json:"LetterSalutation,omitempty"`

	// salutation type
	SalutationType *SalutationTypeSummary `json:"SalutationType,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime strfmt.DateTime `json:"UpdatedDateTime,omitempty"`
}

// Validate validates this salutation
func (m *Salutation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConstituent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSalutationType(formats); err != nil {
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

func (m *Salutation) validateConstituent(formats strfmt.Registry) error {
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

func (m *Salutation) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Salutation) validateSalutationType(formats strfmt.Registry) error {
	if swag.IsZero(m.SalutationType) { // not required
		return nil
	}

	if m.SalutationType != nil {
		if err := m.SalutationType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SalutationType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("SalutationType")
			}
			return err
		}
	}

	return nil
}

func (m *Salutation) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this salutation based on the context it is used
func (m *Salutation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConstituent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSalutationType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Salutation) contextValidateConstituent(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Salutation) contextValidateSalutationType(ctx context.Context, formats strfmt.Registry) error {

	if m.SalutationType != nil {

		if swag.IsZero(m.SalutationType) { // not required
			return nil
		}

		if err := m.SalutationType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SalutationType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("SalutationType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Salutation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Salutation) UnmarshalBinary(b []byte) error {
	var res Salutation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
