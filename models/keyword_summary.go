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

// KeywordSummary keyword summary
//
// swagger:model KeywordSummary
type KeywordSummary struct {

	// category
	Category *KeywordCategorySummary `json:"Category,omitempty"`

	// custom Id
	CustomID int32 `json:"CustomId,omitempty"`

	// data type
	DataType string `json:"DataType,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// detail column
	DetailColumn string `json:"DetailColumn,omitempty"`

	// detail table
	DetailTable string `json:"DetailTable,omitempty"`

	// edit mask
	EditMask string `json:"EditMask,omitempty"`

	// extended description
	ExtendedDescription string `json:"ExtendedDescription,omitempty"`

	// help text
	HelpText string `json:"HelpText,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// keyword use
	KeywordUse string `json:"KeywordUse,omitempty"`

	// multiple value
	MultipleValue bool `json:"MultipleValue,omitempty"`

	// parent table
	ParentTable string `json:"ParentTable,omitempty"`

	// primary group default
	PrimaryGroupDefault string `json:"PrimaryGroupDefault,omitempty"`

	// reference table
	ReferenceTable string `json:"ReferenceTable,omitempty"`
}

// Validate validates this keyword summary
func (m *KeywordSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeywordSummary) validateCategory(formats strfmt.Registry) error {
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

// ContextValidate validate this keyword summary based on the context it is used
func (m *KeywordSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeywordSummary) contextValidateCategory(ctx context.Context, formats strfmt.Registry) error {

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
func (m *KeywordSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeywordSummary) UnmarshalBinary(b []byte) error {
	var res KeywordSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
