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
	"github.com/go-openapi/validate"
)

// AuditLog audit log
//
// swagger:model AuditLog
type AuditLog struct {

	// action date time
	// Format: date-time
	ActionDateTime *strfmt.DateTime `json:"ActionDateTime,omitempty"`

	// audit details
	AuditDetails []*AuditDetail `json:"AuditDetails"`

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// user name
	UserName string `json:"UserName,omitempty"`
}

// Validate validates this audit log
func (m *AuditLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuditDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditLog) validateActionDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("ActionDateTime", "body", "date-time", m.ActionDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AuditLog) validateAuditDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.AuditDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.AuditDetails); i++ {
		if swag.IsZero(m.AuditDetails[i]) { // not required
			continue
		}

		if m.AuditDetails[i] != nil {
			if err := m.AuditDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AuditDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AuditDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this audit log based on the context it is used
func (m *AuditLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuditDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditLog) contextValidateAuditDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AuditDetails); i++ {

		if m.AuditDetails[i] != nil {

			if swag.IsZero(m.AuditDetails[i]) { // not required
				return nil
			}

			if err := m.AuditDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AuditDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AuditDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditLog) UnmarshalBinary(b []byte) error {
	var res AuditLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
