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

// ObjectPermissionSummary object permission summary
//
// swagger:model ObjectPermissionSummary
type ObjectPermissionSummary struct {

	// application object
	ApplicationObject *ApplicationObject `json:"ApplicationObject,omitempty"`

	// can create
	CanCreate string `json:"CanCreate,omitempty"`

	// can delete
	CanDelete string `json:"CanDelete,omitempty"`

	// can edit
	CanEdit string `json:"CanEdit,omitempty"`

	// can view
	CanView string `json:"CanView,omitempty"`

	// constituency
	Constituency *ConstituencyTypeSummary `json:"Constituency,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// user group
	UserGroup *UserGroup `json:"UserGroup,omitempty"`
}

// Validate validates this object permission summary
func (m *ObjectPermissionSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConstituency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectPermissionSummary) validateApplicationObject(formats strfmt.Registry) error {
	if swag.IsZero(m.ApplicationObject) { // not required
		return nil
	}

	if m.ApplicationObject != nil {
		if err := m.ApplicationObject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ApplicationObject")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ApplicationObject")
			}
			return err
		}
	}

	return nil
}

func (m *ObjectPermissionSummary) validateConstituency(formats strfmt.Registry) error {
	if swag.IsZero(m.Constituency) { // not required
		return nil
	}

	if m.Constituency != nil {
		if err := m.Constituency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Constituency")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Constituency")
			}
			return err
		}
	}

	return nil
}

func (m *ObjectPermissionSummary) validateUserGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.UserGroup) { // not required
		return nil
	}

	if m.UserGroup != nil {
		if err := m.UserGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("UserGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("UserGroup")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this object permission summary based on the context it is used
func (m *ObjectPermissionSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplicationObject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConstituency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectPermissionSummary) contextValidateApplicationObject(ctx context.Context, formats strfmt.Registry) error {

	if m.ApplicationObject != nil {

		if swag.IsZero(m.ApplicationObject) { // not required
			return nil
		}

		if err := m.ApplicationObject.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ApplicationObject")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ApplicationObject")
			}
			return err
		}
	}

	return nil
}

func (m *ObjectPermissionSummary) contextValidateConstituency(ctx context.Context, formats strfmt.Registry) error {

	if m.Constituency != nil {

		if swag.IsZero(m.Constituency) { // not required
			return nil
		}

		if err := m.Constituency.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Constituency")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Constituency")
			}
			return err
		}
	}

	return nil
}

func (m *ObjectPermissionSummary) contextValidateUserGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.UserGroup != nil {

		if swag.IsZero(m.UserGroup) { // not required
			return nil
		}

		if err := m.UserGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("UserGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("UserGroup")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObjectPermissionSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectPermissionSummary) UnmarshalBinary(b []byte) error {
	var res ObjectPermissionSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
