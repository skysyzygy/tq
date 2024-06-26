// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SendTemplateEmailRequest send template email request
//
// swagger:model SendTemplateEmailRequest
type SendTemplateEmailRequest struct {

	// email address
	EmailAddress string `json:"EmailAddress,omitempty"`

	// email profile Id
	EmailProfileID int32 `json:"EmailProfileId,omitempty"`

	// name values
	NameValues []*NameValue `json:"NameValues"`

	// template Id
	TemplateID int32 `json:"TemplateId,omitempty"`
}

// Validate validates this send template email request
func (m *SendTemplateEmailRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNameValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendTemplateEmailRequest) validateNameValues(formats strfmt.Registry) error {
	if swag.IsZero(m.NameValues) { // not required
		return nil
	}

	for i := 0; i < len(m.NameValues); i++ {
		if swag.IsZero(m.NameValues[i]) { // not required
			continue
		}

		if m.NameValues[i] != nil {
			if err := m.NameValues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NameValues" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NameValues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this send template email request based on the context it is used
func (m *SendTemplateEmailRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNameValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendTemplateEmailRequest) contextValidateNameValues(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NameValues); i++ {

		if m.NameValues[i] != nil {

			if swag.IsZero(m.NameValues[i]) { // not required
				return nil
			}

			if err := m.NameValues[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NameValues" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NameValues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SendTemplateEmailRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendTemplateEmailRequest) UnmarshalBinary(b []byte) error {
	var res SendTemplateEmailRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
