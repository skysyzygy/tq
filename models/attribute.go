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

// Attribute attribute
//
// swagger:model Attribute
type Attribute struct {

	// constituent
	Constituent *Entity `json:"Constituent,omitempty"`

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// edit indicator
	EditIndicator bool `json:"EditIndicator,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// keyword
	Keyword *KeywordSummary `json:"Keyword,omitempty"`

	// keyword constituent type
	KeywordConstituentType *KeywordConstituentType `json:"KeywordConstituentType,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime strfmt.DateTime `json:"UpdatedDateTime,omitempty"`

	// value
	Value string `json:"Value,omitempty"`
}

// Validate validates this attribute
func (m *Attribute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConstituent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeywordConstituentType(formats); err != nil {
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

func (m *Attribute) validateConstituent(formats strfmt.Registry) error {
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

func (m *Attribute) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Attribute) validateKeyword(formats strfmt.Registry) error {
	if swag.IsZero(m.Keyword) { // not required
		return nil
	}

	if m.Keyword != nil {
		if err := m.Keyword.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Keyword")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Keyword")
			}
			return err
		}
	}

	return nil
}

func (m *Attribute) validateKeywordConstituentType(formats strfmt.Registry) error {
	if swag.IsZero(m.KeywordConstituentType) { // not required
		return nil
	}

	if m.KeywordConstituentType != nil {
		if err := m.KeywordConstituentType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("KeywordConstituentType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("KeywordConstituentType")
			}
			return err
		}
	}

	return nil
}

func (m *Attribute) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this attribute based on the context it is used
func (m *Attribute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConstituent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKeyword(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKeywordConstituentType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Attribute) contextValidateConstituent(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Attribute) contextValidateKeyword(ctx context.Context, formats strfmt.Registry) error {

	if m.Keyword != nil {

		if swag.IsZero(m.Keyword) { // not required
			return nil
		}

		if err := m.Keyword.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Keyword")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Keyword")
			}
			return err
		}
	}

	return nil
}

func (m *Attribute) contextValidateKeywordConstituentType(ctx context.Context, formats strfmt.Registry) error {

	if m.KeywordConstituentType != nil {

		if swag.IsZero(m.KeywordConstituentType) { // not required
			return nil
		}

		if err := m.KeywordConstituentType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("KeywordConstituentType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("KeywordConstituentType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Attribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Attribute) UnmarshalBinary(b []byte) error {
	var res Attribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
