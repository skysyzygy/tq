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

// Title title
//
// swagger:model Title
type Title struct {

	// author
	Author string `json:"Author,omitempty"`

	// composer
	Composer *ComposerSummary `json:"Composer,omitempty"`

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime *strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// era
	Era *EraSummary `json:"Era,omitempty"`

	// fulltext
	Fulltext string `json:"Fulltext,omitempty"`

	// fulltext complete date time
	// Format: date-time
	FulltextCompleteDateTime *strfmt.DateTime `json:"FulltextCompleteDateTime,omitempty"`

	// fulltext request date time
	// Format: date-time
	FulltextRequestDateTime *strfmt.DateTime `json:"FulltextRequestDateTime,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// inventory type
	InventoryType string `json:"InventoryType,omitempty"`

	// lib
	Lib string `json:"Lib,omitempty"`

	// original language
	OriginalLanguage *LanguageSummary `json:"OriginalLanguage,omitempty"`

	// original synopsis
	OriginalSynopsis string `json:"OriginalSynopsis,omitempty"`

	// short name
	ShortName string `json:"ShortName,omitempty"`

	// text1
	Text1 string `json:"Text1,omitempty"`

	// text2
	Text2 string `json:"Text2,omitempty"`

	// text3
	Text3 string `json:"Text3,omitempty"`

	// text4
	Text4 string `json:"Text4,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime *strfmt.DateTime `json:"UpdatedDateTime,omitempty"`
}

// Validate validates this title
func (m *Title) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComposer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEra(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFulltextCompleteDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFulltextRequestDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginalLanguage(formats); err != nil {
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

func (m *Title) validateComposer(formats strfmt.Registry) error {
	if swag.IsZero(m.Composer) { // not required
		return nil
	}

	if m.Composer != nil {
		if err := m.Composer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Composer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Composer")
			}
			return err
		}
	}

	return nil
}

func (m *Title) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Title) validateEra(formats strfmt.Registry) error {
	if swag.IsZero(m.Era) { // not required
		return nil
	}

	if m.Era != nil {
		if err := m.Era.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Era")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Era")
			}
			return err
		}
	}

	return nil
}

func (m *Title) validateFulltextCompleteDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.FulltextCompleteDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("FulltextCompleteDateTime", "body", "date-time", m.FulltextCompleteDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Title) validateFulltextRequestDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.FulltextRequestDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("FulltextRequestDateTime", "body", "date-time", m.FulltextRequestDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Title) validateOriginalLanguage(formats strfmt.Registry) error {
	if swag.IsZero(m.OriginalLanguage) { // not required
		return nil
	}

	if m.OriginalLanguage != nil {
		if err := m.OriginalLanguage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OriginalLanguage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OriginalLanguage")
			}
			return err
		}
	}

	return nil
}

func (m *Title) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this title based on the context it is used
func (m *Title) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComposer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEra(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOriginalLanguage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Title) contextValidateComposer(ctx context.Context, formats strfmt.Registry) error {

	if m.Composer != nil {

		if swag.IsZero(m.Composer) { // not required
			return nil
		}

		if err := m.Composer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Composer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Composer")
			}
			return err
		}
	}

	return nil
}

func (m *Title) contextValidateEra(ctx context.Context, formats strfmt.Registry) error {

	if m.Era != nil {

		if swag.IsZero(m.Era) { // not required
			return nil
		}

		if err := m.Era.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Era")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Era")
			}
			return err
		}
	}

	return nil
}

func (m *Title) contextValidateOriginalLanguage(ctx context.Context, formats strfmt.Registry) error {

	if m.OriginalLanguage != nil {

		if swag.IsZero(m.OriginalLanguage) { // not required
			return nil
		}

		if err := m.OriginalLanguage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OriginalLanguage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OriginalLanguage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Title) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Title) UnmarshalBinary(b []byte) error {
	var res Title
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
