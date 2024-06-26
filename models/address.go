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

// Address address
//
// swagger:model Address
type Address struct {

	// address type
	AddressType *AddressTypeSummary `json:"AddressType,omitempty"`

	// affiliated constituent
	AffiliatedConstituent *Entity `json:"AffiliatedConstituent,omitempty"`

	// alt salutation type
	AltSalutationType *SalutationTypeSummary `json:"AltSalutationType,omitempty"`

	// city
	City string `json:"City,omitempty"`

	// constituent
	Constituent *Entity `json:"Constituent,omitempty"`

	// country
	Country *CountrySummary `json:"Country,omitempty"`

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime *strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// delivery point
	DeliveryPoint string `json:"DeliveryPoint,omitempty"`

	// edit indicator
	EditIndicator bool `json:"EditIndicator,omitempty"`

	// end date
	// Format: date-time
	EndDate *strfmt.DateTime `json:"EndDate,omitempty"`

	// geo area
	GeoArea int32 `json:"GeoArea,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// inactive
	Inactive bool `json:"Inactive,omitempty"`

	// is from affiliation
	IsFromAffiliation bool `json:"IsFromAffiliation,omitempty"`

	// label
	Label bool `json:"Label,omitempty"`

	// months
	Months string `json:"Months,omitempty"`

	// ncoa action
	NcoaAction int32 `json:"NcoaAction,omitempty"`

	// ncoa session
	NcoaSession int32 `json:"NcoaSession,omitempty"`

	// postal code
	PostalCode string `json:"PostalCode,omitempty"`

	// postal code formatted
	PostalCodeFormatted string `json:"PostalCodeFormatted,omitempty"`

	// primary indicator
	PrimaryIndicator bool `json:"PrimaryIndicator,omitempty"`

	// start date
	// Format: date-time
	StartDate *strfmt.DateTime `json:"StartDate,omitempty"`

	// state
	State *StateSummary `json:"State,omitempty"`

	// street1
	Street1 string `json:"Street1,omitempty"`

	// street1 address
	Street1Address *StreetAddress `json:"Street1Address,omitempty"`

	// street2
	Street2 string `json:"Street2,omitempty"`

	// street3
	Street3 string `json:"Street3,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime *strfmt.DateTime `json:"UpdatedDateTime,omitempty"`
}

// Validate validates this address
func (m *Address) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAffiliatedConstituent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAltSalutationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConstituent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStreet1Address(formats); err != nil {
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

func (m *Address) validateAddressType(formats strfmt.Registry) error {
	if swag.IsZero(m.AddressType) { // not required
		return nil
	}

	if m.AddressType != nil {
		if err := m.AddressType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AddressType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("AddressType")
			}
			return err
		}
	}

	return nil
}

func (m *Address) validateAffiliatedConstituent(formats strfmt.Registry) error {
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

func (m *Address) validateAltSalutationType(formats strfmt.Registry) error {
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

func (m *Address) validateConstituent(formats strfmt.Registry) error {
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

func (m *Address) validateCountry(formats strfmt.Registry) error {
	if swag.IsZero(m.Country) { // not required
		return nil
	}

	if m.Country != nil {
		if err := m.Country.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Country")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Country")
			}
			return err
		}
	}

	return nil
}

func (m *Address) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Address) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("EndDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Address) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("StartDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Address) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("State")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("State")
			}
			return err
		}
	}

	return nil
}

func (m *Address) validateStreet1Address(formats strfmt.Registry) error {
	if swag.IsZero(m.Street1Address) { // not required
		return nil
	}

	if m.Street1Address != nil {
		if err := m.Street1Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Street1Address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Street1Address")
			}
			return err
		}
	}

	return nil
}

func (m *Address) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this address based on the context it is used
func (m *Address) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddressType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAffiliatedConstituent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAltSalutationType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConstituent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCountry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStreet1Address(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Address) contextValidateAddressType(ctx context.Context, formats strfmt.Registry) error {

	if m.AddressType != nil {

		if swag.IsZero(m.AddressType) { // not required
			return nil
		}

		if err := m.AddressType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AddressType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("AddressType")
			}
			return err
		}
	}

	return nil
}

func (m *Address) contextValidateAffiliatedConstituent(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Address) contextValidateAltSalutationType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Address) contextValidateConstituent(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Address) contextValidateCountry(ctx context.Context, formats strfmt.Registry) error {

	if m.Country != nil {

		if swag.IsZero(m.Country) { // not required
			return nil
		}

		if err := m.Country.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Country")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Country")
			}
			return err
		}
	}

	return nil
}

func (m *Address) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {

		if swag.IsZero(m.State) { // not required
			return nil
		}

		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("State")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("State")
			}
			return err
		}
	}

	return nil
}

func (m *Address) contextValidateStreet1Address(ctx context.Context, formats strfmt.Registry) error {

	if m.Street1Address != nil {

		if swag.IsZero(m.Street1Address) { // not required
			return nil
		}

		if err := m.Street1Address.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Street1Address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Street1Address")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Address) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Address) UnmarshalBinary(b []byte) error {
	var res Address
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
