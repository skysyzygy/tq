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

// Phone phone
//
// swagger:model Phone
type Phone struct {

	// address
	Address *AddressSummary `json:"Address,omitempty"`

	// allow telemarketing
	AllowTelemarketing bool `json:"AllowTelemarketing,omitempty"`

	// constituent
	Constituent *Entity `json:"Constituent,omitempty"`

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// day evening indicator
	DayEveningIndicator string `json:"DayEveningIndicator,omitempty"`

	// edit indicator
	EditIndicator bool `json:"EditIndicator,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// inactive
	Inactive bool `json:"Inactive,omitempty"`

	// phone formatted
	PhoneFormatted string `json:"PhoneFormatted,omitempty"`

	// phone number
	PhoneNumber string `json:"PhoneNumber,omitempty"`

	// phone search
	PhoneSearch string `json:"PhoneSearch,omitempty"`

	// phone type
	PhoneType *PhoneTypeSummary `json:"PhoneType,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime strfmt.DateTime `json:"UpdatedDateTime,omitempty"`
}

// Validate validates this phone
func (m *Phone) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConstituent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneType(formats); err != nil {
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

func (m *Phone) validateAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Address")
			}
			return err
		}
	}

	return nil
}

func (m *Phone) validateConstituent(formats strfmt.Registry) error {
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

func (m *Phone) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Phone) validatePhoneType(formats strfmt.Registry) error {
	if swag.IsZero(m.PhoneType) { // not required
		return nil
	}

	if m.PhoneType != nil {
		if err := m.PhoneType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PhoneType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PhoneType")
			}
			return err
		}
	}

	return nil
}

func (m *Phone) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this phone based on the context it is used
func (m *Phone) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConstituent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhoneType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Phone) contextValidateAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.Address != nil {

		if swag.IsZero(m.Address) { // not required
			return nil
		}

		if err := m.Address.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Address")
			}
			return err
		}
	}

	return nil
}

func (m *Phone) contextValidateConstituent(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Phone) contextValidatePhoneType(ctx context.Context, formats strfmt.Registry) error {

	if m.PhoneType != nil {

		if swag.IsZero(m.PhoneType) { // not required
			return nil
		}

		if err := m.PhoneType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PhoneType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PhoneType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Phone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Phone) UnmarshalBinary(b []byte) error {
	var res Phone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
