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

// AuditDetail audit detail
//
// swagger:model AuditDetail
type AuditDetail struct {

	// action
	Action string `json:"Action,omitempty"`

	// audit entries
	AuditEntries []*AuditEntry `json:"AuditEntries"`

	// audit table
	AuditTable string `json:"AuditTable,omitempty"`

	// control group
	ControlGroup *ControlGroupSummary `json:"ControlGroup,omitempty"`

	// entity Id
	EntityID int32 `json:"EntityId,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// parent entity Id
	ParentEntityID int32 `json:"ParentEntityId,omitempty"`
}

// Validate validates this audit detail
func (m *AuditDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuditEntries(formats); err != nil {
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

func (m *AuditDetail) validateAuditEntries(formats strfmt.Registry) error {
	if swag.IsZero(m.AuditEntries) { // not required
		return nil
	}

	for i := 0; i < len(m.AuditEntries); i++ {
		if swag.IsZero(m.AuditEntries[i]) { // not required
			continue
		}

		if m.AuditEntries[i] != nil {
			if err := m.AuditEntries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AuditEntries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AuditEntries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AuditDetail) validateControlGroup(formats strfmt.Registry) error {
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

// ContextValidate validate this audit detail based on the context it is used
func (m *AuditDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuditEntries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateControlGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditDetail) contextValidateAuditEntries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AuditEntries); i++ {

		if m.AuditEntries[i] != nil {

			if swag.IsZero(m.AuditEntries[i]) { // not required
				return nil
			}

			if err := m.AuditEntries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AuditEntries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AuditEntries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AuditDetail) contextValidateControlGroup(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AuditDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditDetail) UnmarshalBinary(b []byte) error {
	var res AuditDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
