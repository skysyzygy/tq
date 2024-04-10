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

// ElectronicAddress electronic address
//
// swagger:model ElectronicAddress
type ElectronicAddress struct {

	// address
	Address string `json:"Address,omitempty"`

	// affiliated constituent
	AffiliatedConstituent *Entity `json:"AffiliatedConstituent,omitempty"`

	// allow Html format
	AllowHTMLFormat bool `json:"AllowHtmlFormat,omitempty"`

	// allow marketing
	AllowMarketing bool `json:"AllowMarketing,omitempty"`

	// alt salutation type
	AltSalutationType *SalutationTypeSummary `json:"AltSalutationType,omitempty"`

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

	// electronic address type
	ElectronicAddressType *ElectronicAddressTypeSummary `json:"ElectronicAddressType,omitempty"`

	// end date
	// Format: date-time
	EndDate strfmt.DateTime `json:"EndDate,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// inactive
	Inactive bool `json:"Inactive,omitempty"`

	// is email
	IsEmail bool `json:"IsEmail,omitempty"`

	// is from affiliation
	IsFromAffiliation bool `json:"IsFromAffiliation,omitempty"`

	// months
	Months string `json:"Months,omitempty"`

	// primary indicator
	PrimaryIndicator bool `json:"PrimaryIndicator,omitempty"`

	// start date
	// Format: date-time
	StartDate strfmt.DateTime `json:"StartDate,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime strfmt.DateTime `json:"UpdatedDateTime,omitempty"`
}

// Validate validates this electronic address
func (m *ElectronicAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAffiliatedConstituent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAltSalutationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConstituent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElectronicAddressType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
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

func (m *ElectronicAddress) validateAffiliatedConstituent(formats strfmt.Registry) error {
	if swag.IsZero(m.AffiliatedConstituent) { // not required
		return nil
	}

	if m.AffiliatedConstituent != nil {
		if err := m.AffiliatedConstituent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AffiliatedConstituent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("AffiliatedConstituent")
			}
			return err
		}
	}

	return nil
}

func (m *ElectronicAddress) validateAltSalutationType(formats strfmt.Registry) error {
	if swag.IsZero(m.AltSalutationType) { // not required
		return nil
	}

	if m.AltSalutationType != nil {
		if err := m.AltSalutationType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AltSalutationType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("AltSalutationType")
			}
			return err
		}
	}

	return nil
}

func (m *ElectronicAddress) validateConstituent(formats strfmt.Registry) error {
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

func (m *ElectronicAddress) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ElectronicAddress) validateElectronicAddressType(formats strfmt.Registry) error {
	if swag.IsZero(m.ElectronicAddressType) { // not required
		return nil
	}

	if m.ElectronicAddressType != nil {
		if err := m.ElectronicAddressType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ElectronicAddressType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ElectronicAddressType")
			}
			return err
		}
	}

	return nil
}

func (m *ElectronicAddress) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("EndDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ElectronicAddress) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("StartDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ElectronicAddress) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this electronic address based on the context it is used
func (m *ElectronicAddress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAffiliatedConstituent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAltSalutationType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConstituent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateElectronicAddressType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElectronicAddress) contextValidateAffiliatedConstituent(ctx context.Context, formats strfmt.Registry) error {

	if m.AffiliatedConstituent != nil {

		if swag.IsZero(m.AffiliatedConstituent) { // not required
			return nil
		}

		if err := m.AffiliatedConstituent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AffiliatedConstituent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("AffiliatedConstituent")
			}
			return err
		}
	}

	return nil
}

func (m *ElectronicAddress) contextValidateAltSalutationType(ctx context.Context, formats strfmt.Registry) error {

	if m.AltSalutationType != nil {

		if swag.IsZero(m.AltSalutationType) { // not required
			return nil
		}

		if err := m.AltSalutationType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AltSalutationType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("AltSalutationType")
			}
			return err
		}
	}

	return nil
}

func (m *ElectronicAddress) contextValidateConstituent(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ElectronicAddress) contextValidateElectronicAddressType(ctx context.Context, formats strfmt.Registry) error {

	if m.ElectronicAddressType != nil {

		if swag.IsZero(m.ElectronicAddressType) { // not required
			return nil
		}

		if err := m.ElectronicAddressType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ElectronicAddressType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ElectronicAddressType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ElectronicAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElectronicAddress) UnmarshalBinary(b []byte) error {
	var res ElectronicAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
