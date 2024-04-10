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

// BookingFromTemplateRequest booking from template request
//
// swagger:model BookingFromTemplateRequest
type BookingFromTemplateRequest struct {

	// booking source
	BookingSource int32 `json:"BookingSource,omitempty"`

	// booking template Id
	BookingTemplateID int32 `json:"BookingTemplateId,omitempty"`

	// confirmation text
	ConfirmationText string `json:"ConfirmationText,omitempty"`

	// default count
	DefaultCount int32 `json:"DefaultCount,omitempty"`

	// default date time
	// Format: date-time
	DefaultDateTime strfmt.DateTime `json:"DefaultDateTime,omitempty"`

	// default duration
	DefaultDuration int32 `json:"DefaultDuration,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// notes
	Notes string `json:"Notes,omitempty"`
}

// Validate validates this booking from template request
func (m *BookingFromTemplateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultDateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BookingFromTemplateRequest) validateDefaultDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("DefaultDateTime", "body", "date-time", m.DefaultDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this booking from template request based on context it is used
func (m *BookingFromTemplateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BookingFromTemplateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BookingFromTemplateRequest) UnmarshalBinary(b []byte) error {
	var res BookingFromTemplateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
