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

// IssueAction issue action
//
// swagger:model IssueAction
type IssueAction struct {

	// action date
	// Format: date-time
	ActionDate strfmt.DateTime `json:"ActionDate,omitempty"`

	// action type
	ActionType *ActionTypeSummary `json:"ActionType,omitempty"`

	// constituent
	Constituent *Entity `json:"Constituent,omitempty"`

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// edit indicator
	EditIndicator bool `json:"EditIndicator,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// issue
	Issue *Entity `json:"Issue,omitempty"`

	// letter printed date
	// Format: date-time
	LetterPrintedDate strfmt.DateTime `json:"LetterPrintedDate,omitempty"`

	// notes
	Notes string `json:"Notes,omitempty"`

	// resolved
	Resolved bool `json:"Resolved,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime strfmt.DateTime `json:"UpdatedDateTime,omitempty"`
}

// Validate validates this issue action
func (m *IssueAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConstituent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLetterPrintedDate(formats); err != nil {
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

func (m *IssueAction) validateActionDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionDate) { // not required
		return nil
	}

	if err := validate.FormatOf("ActionDate", "body", "date-time", m.ActionDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IssueAction) validateActionType(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionType) { // not required
		return nil
	}

	if m.ActionType != nil {
		if err := m.ActionType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ActionType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ActionType")
			}
			return err
		}
	}

	return nil
}

func (m *IssueAction) validateConstituent(formats strfmt.Registry) error {
	if swag.IsZero(m.Constituent) { // not required
		return nil
	}

	if m.Constituent != nil {
		if err := m.Constituent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Constituent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Constituent")
			}
			return err
		}
	}

	return nil
}

func (m *IssueAction) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IssueAction) validateIssue(formats strfmt.Registry) error {
	if swag.IsZero(m.Issue) { // not required
		return nil
	}

	if m.Issue != nil {
		if err := m.Issue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Issue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Issue")
			}
			return err
		}
	}

	return nil
}

func (m *IssueAction) validateLetterPrintedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.LetterPrintedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("LetterPrintedDate", "body", "date-time", m.LetterPrintedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IssueAction) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this issue action based on the context it is used
func (m *IssueAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActionType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConstituent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIssue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IssueAction) contextValidateActionType(ctx context.Context, formats strfmt.Registry) error {

	if m.ActionType != nil {

		if swag.IsZero(m.ActionType) { // not required
			return nil
		}

		if err := m.ActionType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ActionType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ActionType")
			}
			return err
		}
	}

	return nil
}

func (m *IssueAction) contextValidateConstituent(ctx context.Context, formats strfmt.Registry) error {

	if m.Constituent != nil {

		if swag.IsZero(m.Constituent) { // not required
			return nil
		}

		if err := m.Constituent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Constituent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Constituent")
			}
			return err
		}
	}

	return nil
}

func (m *IssueAction) contextValidateIssue(ctx context.Context, formats strfmt.Registry) error {

	if m.Issue != nil {

		if swag.IsZero(m.Issue) { // not required
			return nil
		}

		if err := m.Issue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Issue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Issue")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IssueAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IssueAction) UnmarshalBinary(b []byte) error {
	var res IssueAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
