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

// CartMembershipSummary cart membership summary
//
// swagger:model CartMembershipSummary
type CartMembershipSummary struct {

	// amount
	Amount float64 `json:"Amount,omitempty"`

	// fund description
	FundDescription string `json:"FundDescription,omitempty"`

	// membership level description
	MembershipLevelDescription string `json:"MembershipLevelDescription,omitempty"`

	// membership organization
	MembershipOrganization *EntitySummary `json:"MembershipOrganization,omitempty"`

	// reference Id
	ReferenceID int32 `json:"ReferenceId,omitempty"`
}

// Validate validates this cart membership summary
func (m *CartMembershipSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMembershipOrganization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CartMembershipSummary) validateMembershipOrganization(formats strfmt.Registry) error {
	if swag.IsZero(m.MembershipOrganization) { // not required
		return nil
	}

	if m.MembershipOrganization != nil {
		if err := m.MembershipOrganization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MembershipOrganization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MembershipOrganization")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cart membership summary based on the context it is used
func (m *CartMembershipSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMembershipOrganization(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CartMembershipSummary) contextValidateMembershipOrganization(ctx context.Context, formats strfmt.Registry) error {

	if m.MembershipOrganization != nil {

		if swag.IsZero(m.MembershipOrganization) { // not required
			return nil
		}

		if err := m.MembershipOrganization.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MembershipOrganization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MembershipOrganization")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CartMembershipSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CartMembershipSummary) UnmarshalBinary(b []byte) error {
	var res CartMembershipSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
