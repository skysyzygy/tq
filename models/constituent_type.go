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

// ConstituentType constituent type
//
// swagger:model ConstituentType
type ConstituentType struct {

	// address type Id
	AddressTypeID int32 `json:"AddressTypeId,omitempty"`

	// constituent group
	ConstituentGroup *ConstituentGroupSummary `json:"ConstituentGroup,omitempty"`

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// default affiliated constituent type Id
	DefaultAffiliatedConstituentTypeID int32 `json:"DefaultAffiliatedConstituentTypeId,omitempty"`

	// default affiliation type Id
	DefaultAffiliationTypeID int32 `json:"DefaultAffiliationTypeId,omitempty"`

	// default indicator
	DefaultIndicator bool `json:"DefaultIndicator,omitempty"`

	// default salutation Id
	DefaultSalutationID int32 `json:"DefaultSalutationId,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// electronic address type Id
	ElectronicAddressTypeID int32 `json:"ElectronicAddressTypeId,omitempty"`

	// gift aid indicator
	GiftAidIndicator bool `json:"GiftAidIndicator,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// inactive
	Inactive bool `json:"Inactive,omitempty"`

	// login type Id
	LoginTypeID int32 `json:"LoginTypeId,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime strfmt.DateTime `json:"UpdatedDateTime,omitempty"`
}

// Validate validates this constituent type
func (m *ConstituentType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConstituentGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDateTime(formats); err != nil {
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

func (m *ConstituentType) validateConstituentGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.ConstituentGroup) { // not required
		return nil
	}

	if m.ConstituentGroup != nil {
		if err := m.ConstituentGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ConstituentGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ConstituentGroup")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentType) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConstituentType) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this constituent type based on the context it is used
func (m *ConstituentType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConstituentGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConstituentType) contextValidateConstituentGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.ConstituentGroup != nil {

		if swag.IsZero(m.ConstituentGroup) { // not required
			return nil
		}

		if err := m.ConstituentGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ConstituentGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ConstituentGroup")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConstituentType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConstituentType) UnmarshalBinary(b []byte) error {
	var res ConstituentType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
