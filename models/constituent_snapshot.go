// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConstituentSnapshot constituent snapshot
//
// swagger:model ConstituentSnapshot
type ConstituentSnapshot struct {

	// address
	Address *Address `json:"Address,omitempty"`

	// affiliates
	Affiliates []*AffiliationInfo `json:"Affiliates"`

	// constituent type
	ConstituentType *ConstituentTypeSummary `json:"ConstituentType,omitempty"`

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime *strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// display name
	DisplayName string `json:"DisplayName,omitempty"`

	// edit indicator
	EditIndicator bool `json:"EditIndicator,omitempty"`

	// electronic address
	ElectronicAddress *ElectronicAddress `json:"ElectronicAddress,omitempty"`

	// emarket indicator
	EmarketIndicator *EmarketIndicatorSummary `json:"EmarketIndicator,omitempty"`

	// first name
	FirstName string `json:"FirstName,omitempty"`

	// gender
	Gender *GenderSummary `json:"Gender,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// inactive
	Inactive *ConstituentInactiveSummary `json:"Inactive,omitempty"`

	// inactive reason
	InactiveReason *InactiveReasonSummary `json:"InactiveReason,omitempty"`

	// last activity date
	// Format: date-time
	LastActivityDate *strfmt.DateTime `json:"LastActivityDate,omitempty"`

	// last name
	LastName string `json:"LastName,omitempty"`

	// mail indicator
	MailIndicator *MailIndicatorSummary `json:"MailIndicator,omitempty"`

	// middle name
	MiddleName string `json:"MiddleName,omitempty"`

	// name status
	NameStatus *NameStatusSummary `json:"NameStatus,omitempty"`

	// original source
	OriginalSource *OriginalSourceSummary `json:"OriginalSource,omitempty"`

	// phone indicator
	PhoneIndicator *PhoneIndicatorSummary `json:"PhoneIndicator,omitempty"`

	// prefix
	Prefix *PrefixSummary `json:"Prefix,omitempty"`

	// primary phone numbers
	PrimaryPhoneNumbers []*Phone `json:"PrimaryPhoneNumbers"`

	// protection type
	ProtectionType *ConstituentProtectionTypeSummary `json:"ProtectionType,omitempty"`

	// salutation
	Salutation *Salutation `json:"Salutation,omitempty"`

	// short display name
	ShortDisplayName string `json:"ShortDisplayName,omitempty"`

	// sort name
	SortName string `json:"SortName,omitempty"`

	// suffix
	Suffix *SuffixSummary `json:"Suffix,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime *strfmt.DateTime `json:"UpdatedDateTime,omitempty"`
}

// Validate validates this constituent snapshot
func (m *ConstituentSnapshot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAffiliates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConstituentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElectronicAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmarketIndicator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGender(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInactive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInactiveReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastActivityDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMailIndicator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginalSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneIndicator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrefix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryPhoneNumbers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtectionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSalutation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuffix(formats); err != nil {
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

func (m *ConstituentSnapshot) validateAddress(formats strfmt.Registry) error {
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

func (m *ConstituentSnapshot) validateAffiliates(formats strfmt.Registry) error {
	if swag.IsZero(m.Affiliates) { // not required
		return nil
	}

	for i := 0; i < len(m.Affiliates); i++ {
		if swag.IsZero(m.Affiliates[i]) { // not required
			continue
		}

		if m.Affiliates[i] != nil {
			if err := m.Affiliates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Affiliates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Affiliates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConstituentSnapshot) validateConstituentType(formats strfmt.Registry) error {
	if swag.IsZero(m.ConstituentType) { // not required
		return nil
	}

	if m.ConstituentType != nil {
		if err := m.ConstituentType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ConstituentType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ConstituentType")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConstituentSnapshot) validateElectronicAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.ElectronicAddress) { // not required
		return nil
	}

	if m.ElectronicAddress != nil {
		if err := m.ElectronicAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ElectronicAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ElectronicAddress")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) validateEmarketIndicator(formats strfmt.Registry) error {
	if swag.IsZero(m.EmarketIndicator) { // not required
		return nil
	}

	if m.EmarketIndicator != nil {
		if err := m.EmarketIndicator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EmarketIndicator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("EmarketIndicator")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) validateGender(formats strfmt.Registry) error {
	if swag.IsZero(m.Gender) { // not required
		return nil
	}

	if m.Gender != nil {
		if err := m.Gender.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Gender")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Gender")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) validateInactive(formats strfmt.Registry) error {
	if swag.IsZero(m.Inactive) { // not required
		return nil
	}

	if m.Inactive != nil {
		if err := m.Inactive.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Inactive")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Inactive")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) validateInactiveReason(formats strfmt.Registry) error {
	if swag.IsZero(m.InactiveReason) { // not required
		return nil
	}

	if m.InactiveReason != nil {
		if err := m.InactiveReason.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("InactiveReason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("InactiveReason")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) validateLastActivityDate(formats strfmt.Registry) error {
	if swag.IsZero(m.LastActivityDate) { // not required
		return nil
	}

	if err := validate.FormatOf("LastActivityDate", "body", "date-time", m.LastActivityDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConstituentSnapshot) validateMailIndicator(formats strfmt.Registry) error {
	if swag.IsZero(m.MailIndicator) { // not required
		return nil
	}

	if m.MailIndicator != nil {
		if err := m.MailIndicator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MailIndicator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MailIndicator")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) validateNameStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.NameStatus) { // not required
		return nil
	}

	if m.NameStatus != nil {
		if err := m.NameStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NameStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("NameStatus")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) validateOriginalSource(formats strfmt.Registry) error {
	if swag.IsZero(m.OriginalSource) { // not required
		return nil
	}

	if m.OriginalSource != nil {
		if err := m.OriginalSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OriginalSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OriginalSource")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) validatePhoneIndicator(formats strfmt.Registry) error {
	if swag.IsZero(m.PhoneIndicator) { // not required
		return nil
	}

	if m.PhoneIndicator != nil {
		if err := m.PhoneIndicator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PhoneIndicator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PhoneIndicator")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) validatePrefix(formats strfmt.Registry) error {
	if swag.IsZero(m.Prefix) { // not required
		return nil
	}

	if m.Prefix != nil {
		if err := m.Prefix.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Prefix")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Prefix")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) validatePrimaryPhoneNumbers(formats strfmt.Registry) error {
	if swag.IsZero(m.PrimaryPhoneNumbers) { // not required
		return nil
	}

	for i := 0; i < len(m.PrimaryPhoneNumbers); i++ {
		if swag.IsZero(m.PrimaryPhoneNumbers[i]) { // not required
			continue
		}

		if m.PrimaryPhoneNumbers[i] != nil {
			if err := m.PrimaryPhoneNumbers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PrimaryPhoneNumbers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("PrimaryPhoneNumbers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConstituentSnapshot) validateProtectionType(formats strfmt.Registry) error {
	if swag.IsZero(m.ProtectionType) { // not required
		return nil
	}

	if m.ProtectionType != nil {
		if err := m.ProtectionType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ProtectionType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ProtectionType")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) validateSalutation(formats strfmt.Registry) error {
	if swag.IsZero(m.Salutation) { // not required
		return nil
	}

	if m.Salutation != nil {
		if err := m.Salutation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Salutation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Salutation")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) validateSuffix(formats strfmt.Registry) error {
	if swag.IsZero(m.Suffix) { // not required
		return nil
	}

	if m.Suffix != nil {
		if err := m.Suffix.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Suffix")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Suffix")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this constituent snapshot based on the context it is used
func (m *ConstituentSnapshot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAffiliates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConstituentType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateElectronicAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmarketIndicator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGender(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInactive(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInactiveReason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMailIndicator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNameStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOriginalSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhoneIndicator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrefix(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrimaryPhoneNumbers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtectionType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSalutation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSuffix(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConstituentSnapshot) contextValidateAddress(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConstituentSnapshot) contextValidateAffiliates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Affiliates); i++ {

		if m.Affiliates[i] != nil {

			if swag.IsZero(m.Affiliates[i]) { // not required
				return nil
			}

			if err := m.Affiliates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Affiliates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Affiliates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConstituentSnapshot) contextValidateConstituentType(ctx context.Context, formats strfmt.Registry) error {

	if m.ConstituentType != nil {

		if swag.IsZero(m.ConstituentType) { // not required
			return nil
		}

		if err := m.ConstituentType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ConstituentType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ConstituentType")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) contextValidateElectronicAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.ElectronicAddress != nil {

		if swag.IsZero(m.ElectronicAddress) { // not required
			return nil
		}

		if err := m.ElectronicAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ElectronicAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ElectronicAddress")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) contextValidateEmarketIndicator(ctx context.Context, formats strfmt.Registry) error {

	if m.EmarketIndicator != nil {

		if swag.IsZero(m.EmarketIndicator) { // not required
			return nil
		}

		if err := m.EmarketIndicator.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EmarketIndicator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("EmarketIndicator")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) contextValidateGender(ctx context.Context, formats strfmt.Registry) error {

	if m.Gender != nil {

		if swag.IsZero(m.Gender) { // not required
			return nil
		}

		if err := m.Gender.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Gender")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Gender")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) contextValidateInactive(ctx context.Context, formats strfmt.Registry) error {

	if m.Inactive != nil {

		if swag.IsZero(m.Inactive) { // not required
			return nil
		}

		if err := m.Inactive.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Inactive")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Inactive")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) contextValidateInactiveReason(ctx context.Context, formats strfmt.Registry) error {

	if m.InactiveReason != nil {

		if swag.IsZero(m.InactiveReason) { // not required
			return nil
		}

		if err := m.InactiveReason.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("InactiveReason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("InactiveReason")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) contextValidateMailIndicator(ctx context.Context, formats strfmt.Registry) error {

	if m.MailIndicator != nil {

		if swag.IsZero(m.MailIndicator) { // not required
			return nil
		}

		if err := m.MailIndicator.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MailIndicator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MailIndicator")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) contextValidateNameStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.NameStatus != nil {

		if swag.IsZero(m.NameStatus) { // not required
			return nil
		}

		if err := m.NameStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NameStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("NameStatus")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) contextValidateOriginalSource(ctx context.Context, formats strfmt.Registry) error {

	if m.OriginalSource != nil {

		if swag.IsZero(m.OriginalSource) { // not required
			return nil
		}

		if err := m.OriginalSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OriginalSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OriginalSource")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) contextValidatePhoneIndicator(ctx context.Context, formats strfmt.Registry) error {

	if m.PhoneIndicator != nil {

		if swag.IsZero(m.PhoneIndicator) { // not required
			return nil
		}

		if err := m.PhoneIndicator.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PhoneIndicator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PhoneIndicator")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) contextValidatePrefix(ctx context.Context, formats strfmt.Registry) error {

	if m.Prefix != nil {

		if swag.IsZero(m.Prefix) { // not required
			return nil
		}

		if err := m.Prefix.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Prefix")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Prefix")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) contextValidatePrimaryPhoneNumbers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PrimaryPhoneNumbers); i++ {

		if m.PrimaryPhoneNumbers[i] != nil {

			if swag.IsZero(m.PrimaryPhoneNumbers[i]) { // not required
				return nil
			}

			if err := m.PrimaryPhoneNumbers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PrimaryPhoneNumbers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("PrimaryPhoneNumbers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConstituentSnapshot) contextValidateProtectionType(ctx context.Context, formats strfmt.Registry) error {

	if m.ProtectionType != nil {

		if swag.IsZero(m.ProtectionType) { // not required
			return nil
		}

		if err := m.ProtectionType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ProtectionType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ProtectionType")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) contextValidateSalutation(ctx context.Context, formats strfmt.Registry) error {

	if m.Salutation != nil {

		if swag.IsZero(m.Salutation) { // not required
			return nil
		}

		if err := m.Salutation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Salutation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Salutation")
			}
			return err
		}
	}

	return nil
}

func (m *ConstituentSnapshot) contextValidateSuffix(ctx context.Context, formats strfmt.Registry) error {

	if m.Suffix != nil {

		if swag.IsZero(m.Suffix) { // not required
			return nil
		}

		if err := m.Suffix.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Suffix")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Suffix")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConstituentSnapshot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConstituentSnapshot) UnmarshalBinary(b []byte) error {
	var res ConstituentSnapshot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
