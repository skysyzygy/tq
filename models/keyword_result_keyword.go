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

// KeywordResultKeyword keyword result keyword
//
// swagger:model KeywordResultKeyword
type KeywordResultKeyword struct {

	// category
	Category *EntitySummary `json:"Category,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// owner
	Owner *Owner `json:"Owner,omitempty"`
}

// Validate validates this keyword result keyword
func (m *KeywordResultKeyword) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeywordResultKeyword) validateCategory(formats strfmt.Registry) error {
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

func (m *KeywordResultKeyword) validateOwner(formats strfmt.Registry) error {
	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Owner")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this keyword result keyword based on the context it is used
func (m *KeywordResultKeyword) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeywordResultKeyword) contextValidateCategory(ctx context.Context, formats strfmt.Registry) error {

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

func (m *KeywordResultKeyword) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

	if m.Owner != nil {

		if swag.IsZero(m.Owner) { // not required
			return nil
		}

		if err := m.Owner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Owner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KeywordResultKeyword) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeywordResultKeyword) UnmarshalBinary(b []byte) error {
	var res KeywordResultKeyword
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
