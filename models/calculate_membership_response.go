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

// CalculateMembershipResponse calculate membership response
//
// swagger:model CalculateMembershipResponse
type CalculateMembershipResponse struct {

	// calculated membership
	CalculatedMembership *CalculatedMembership `json:"CalculatedMembership,omitempty"`

	// current membership summary
	CurrentMembershipSummary *CurrentMembership `json:"CurrentMembershipSummary,omitempty"`
}

// Validate validates this calculate membership response
func (m *CalculateMembershipResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCalculatedMembership(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentMembershipSummary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CalculateMembershipResponse) validateCalculatedMembership(formats strfmt.Registry) error {
	if swag.IsZero(m.CalculatedMembership) { // not required
		return nil
	}

	if m.CalculatedMembership != nil {
		if err := m.CalculatedMembership.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CalculatedMembership")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CalculatedMembership")
			}
			return err
		}
	}

	return nil
}

func (m *CalculateMembershipResponse) validateCurrentMembershipSummary(formats strfmt.Registry) error {
	if swag.IsZero(m.CurrentMembershipSummary) { // not required
		return nil
	}

	if m.CurrentMembershipSummary != nil {
		if err := m.CurrentMembershipSummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CurrentMembershipSummary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CurrentMembershipSummary")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this calculate membership response based on the context it is used
func (m *CalculateMembershipResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCalculatedMembership(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCurrentMembershipSummary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CalculateMembershipResponse) contextValidateCalculatedMembership(ctx context.Context, formats strfmt.Registry) error {

	if m.CalculatedMembership != nil {

		if swag.IsZero(m.CalculatedMembership) { // not required
			return nil
		}

		if err := m.CalculatedMembership.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CalculatedMembership")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CalculatedMembership")
			}
			return err
		}
	}

	return nil
}

func (m *CalculateMembershipResponse) contextValidateCurrentMembershipSummary(ctx context.Context, formats strfmt.Registry) error {

	if m.CurrentMembershipSummary != nil {

		if swag.IsZero(m.CurrentMembershipSummary) { // not required
			return nil
		}

		if err := m.CurrentMembershipSummary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CurrentMembershipSummary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CurrentMembershipSummary")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CalculateMembershipResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CalculateMembershipResponse) UnmarshalBinary(b []byte) error {
	var res CalculateMembershipResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
