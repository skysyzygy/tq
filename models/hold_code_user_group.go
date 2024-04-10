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

// HoldCodeUserGroup hold code user group
//
// swagger:model HoldCodeUserGroup
type HoldCodeUserGroup struct {

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// hold code
	HoldCode *HoldCodeSummary `json:"HoldCode,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime strfmt.DateTime `json:"UpdatedDateTime,omitempty"`

	// user group Id
	UserGroupID string `json:"UserGroupId,omitempty"`
}

// Validate validates this hold code user group
func (m *HoldCodeUserGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHoldCode(formats); err != nil {
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

func (m *HoldCodeUserGroup) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HoldCodeUserGroup) validateHoldCode(formats strfmt.Registry) error {
	if swag.IsZero(m.HoldCode) { // not required
		return nil
	}

	if m.HoldCode != nil {
		if err := m.HoldCode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("HoldCode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("HoldCode")
			}
			return err
		}
	}

	return nil
}

func (m *HoldCodeUserGroup) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this hold code user group based on the context it is used
func (m *HoldCodeUserGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHoldCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HoldCodeUserGroup) contextValidateHoldCode(ctx context.Context, formats strfmt.Registry) error {

	if m.HoldCode != nil {

		if swag.IsZero(m.HoldCode) { // not required
			return nil
		}

		if err := m.HoldCode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("HoldCode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("HoldCode")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HoldCodeUserGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HoldCodeUserGroup) UnmarshalBinary(b []byte) error {
	var res HoldCodeUserGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
