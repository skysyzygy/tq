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

// AuditEntryDetail audit entry detail
//
// swagger:model AuditEntryDetail
type AuditEntryDetail struct {

	// action
	Action string `json:"Action,omitempty"`

	// action date time
	// Format: date-time
	ActionDateTime strfmt.DateTime `json:"ActionDateTime,omitempty"`

	// audit detail
	AuditDetail int32 `json:"AuditDetail,omitempty"`

	// audit log Id
	AuditLogID int32 `json:"AuditLogId,omitempty"`

	// audit table
	AuditTable string `json:"AuditTable,omitempty"`

	// control group
	ControlGroup *Entity `json:"ControlGroup,omitempty"`

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// entity Id
	EntityID int32 `json:"EntityId,omitempty"`

	// new
	New string `json:"New,omitempty"`

	// old
	Old string `json:"Old,omitempty"`

	// parent entity Id
	ParentEntityID int32 `json:"ParentEntityId,omitempty"`

	// property
	Property string `json:"Property,omitempty"`

	// user name
	UserName string `json:"UserName,omitempty"`
}

// Validate validates this audit entry detail
func (m *AuditEntryDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateControlGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditEntryDetail) validateActionDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("ActionDateTime", "body", "date-time", m.ActionDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AuditEntryDetail) validateControlGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.ControlGroup) { // not required
		return nil
	}

	if m.ControlGroup != nil {
		if err := m.ControlGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ControlGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ControlGroup")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this audit entry detail based on the context it is used
func (m *AuditEntryDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateControlGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditEntryDetail) contextValidateControlGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.ControlGroup != nil {

		if swag.IsZero(m.ControlGroup) { // not required
			return nil
		}

		if err := m.ControlGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ControlGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ControlGroup")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditEntryDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditEntryDetail) UnmarshalBinary(b []byte) error {
	var res AuditEntryDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
