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

// Template template
//
// swagger:model Template
type Template struct {

	// category
	Category *TemplateCategorySummary `json:"Category,omitempty"`

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// edit indicator
	EditIndicator bool `json:"EditIndicator,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// inactive
	Inactive bool `json:"Inactive,omitempty"`

	// is Html body
	IsHTMLBody bool `json:"IsHtmlBody,omitempty"`

	// template body
	TemplateBody string `json:"TemplateBody,omitempty"`

	// template subject
	TemplateSubject string `json:"TemplateSubject,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime strfmt.DateTime `json:"UpdatedDateTime,omitempty"`
}

// Validate validates this template
func (m *Template) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDateTime(formats); err != nil {
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

func (m *Template) validateCategory(formats strfmt.Registry) error {
	if swag.IsZero(m.Category) { // not required
		return nil
	}

	if m.Category != nil {
		if err := m.Category.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Category")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Category")
			}
			return err
		}
	}

	return nil
}

func (m *Template) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Template) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this template based on the context it is used
func (m *Template) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Template) contextValidateCategory(ctx context.Context, formats strfmt.Registry) error {

	if m.Category != nil {

		if swag.IsZero(m.Category) { // not required
			return nil
		}

		if err := m.Category.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Category")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Category")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Template) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Template) UnmarshalBinary(b []byte) error {
	var res Template
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
