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

// OutputSetElement output set element
//
// swagger:model OutputSetElement
type OutputSetElement struct {

	// description renamed
	DescriptionRenamed string `json:"DescriptionRenamed,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// inactive
	Inactive bool `json:"Inactive,omitempty"`

	// query element
	QueryElement *QueryElementSummary `json:"QueryElement,omitempty"`

	// rank
	Rank int32 `json:"Rank,omitempty"`
}

// Validate validates this output set element
func (m *OutputSetElement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQueryElement(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OutputSetElement) validateQueryElement(formats strfmt.Registry) error {
	if swag.IsZero(m.QueryElement) { // not required
		return nil
	}

	if m.QueryElement != nil {
		if err := m.QueryElement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("QueryElement")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("QueryElement")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this output set element based on the context it is used
func (m *OutputSetElement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateQueryElement(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OutputSetElement) contextValidateQueryElement(ctx context.Context, formats strfmt.Registry) error {

	if m.QueryElement != nil {

		if swag.IsZero(m.QueryElement) { // not required
			return nil
		}

		if err := m.QueryElement.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("QueryElement")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("QueryElement")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OutputSetElement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OutputSetElement) UnmarshalBinary(b []byte) error {
	var res OutputSetElement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
