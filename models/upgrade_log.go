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

// UpgradeLog upgrade log
//
// swagger:model UpgradeLog
type UpgradeLog struct {

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// hot fix number
	HotFixNumber int32 `json:"HotFixNumber,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// release description
	ReleaseDescription string `json:"ReleaseDescription,omitempty"`

	// script Id
	ScriptID int32 `json:"ScriptId,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime strfmt.DateTime `json:"UpdatedDateTime,omitempty"`

	// upgrade category summary
	UpgradeCategorySummary *UpgradeCategorySummary `json:"UpgradeCategorySummary,omitempty"`

	// version
	Version string `json:"Version,omitempty"`
}

// Validate validates this upgrade log
func (m *UpgradeLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpgradeCategorySummary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpgradeLog) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpgradeLog) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpgradeLog) validateUpgradeCategorySummary(formats strfmt.Registry) error {
	if swag.IsZero(m.UpgradeCategorySummary) { // not required
		return nil
	}

	if m.UpgradeCategorySummary != nil {
		if err := m.UpgradeCategorySummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("UpgradeCategorySummary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("UpgradeCategorySummary")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this upgrade log based on the context it is used
func (m *UpgradeLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUpgradeCategorySummary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpgradeLog) contextValidateUpgradeCategorySummary(ctx context.Context, formats strfmt.Registry) error {

	if m.UpgradeCategorySummary != nil {

		if swag.IsZero(m.UpgradeCategorySummary) { // not required
			return nil
		}

		if err := m.UpgradeCategorySummary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("UpgradeCategorySummary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("UpgradeCategorySummary")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpgradeLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpgradeLog) UnmarshalBinary(b []byte) error {
	var res UpgradeLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
