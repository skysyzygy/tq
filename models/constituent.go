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

// Constituent constituent
//
// swagger:model Constituent
type Constituent struct {

	// constituent type
	ConstituentType *ConstituentTypeSummary `json:"ConstituentType,omitempty"`

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// display name
	DisplayName string `json:"DisplayName,omitempty"`

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
	LastActivityDate strfmt.DateTime `json:"LastActivityDate,omitempty"`

	// last gift date
	// Format: date-time
	LastGiftDate strfmt.DateTime `json:"LastGiftDate,omitempty"`

	// last name
	LastName string `json:"LastName,omitempty"`

	// last ticket date
	// Format: date-time
	LastTicketDate strfmt.DateTime `json:"LastTicketDate,omitempty"`

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

	// protection type
	ProtectionType *ConstituentProtectionTypeSummary `json:"ProtectionType,omitempty"`

	// sort name
	SortName string `json:"SortName,omitempty"`

	// suffix
	Suffix *SuffixSummary `json:"Suffix,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime strfmt.DateTime `json:"UpdatedDateTime,omitempty"`
}

// Validate validates this constituent
func (m *Constituent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConstituentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDateTime(formats); err != nil {
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

	if err := m.validateLastGiftDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastTicketDate(formats); err != nil {
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

	if err := m.validateProtectionType(formats); err != nil {
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

func (m *Constituent) validateConstituentType(formats strfmt.Registry) error {
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

func (m *Constituent) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Constituent) validateEmarketIndicator(formats strfmt.Registry) error {
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

func (m *Constituent) validateGender(formats strfmt.Registry) error {
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

func (m *Constituent) validateInactive(formats strfmt.Registry) error {
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

func (m *Constituent) validateInactiveReason(formats strfmt.Registry) error {
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

func (m *Constituent) validateLastActivityDate(formats strfmt.Registry) error {
	if swag.IsZero(m.LastActivityDate) { // not required
		return nil
	}

	if err := validate.FormatOf("LastActivityDate", "body", "date-time", m.LastActivityDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Constituent) validateLastGiftDate(formats strfmt.Registry) error {
	if swag.IsZero(m.LastGiftDate) { // not required
		return nil
	}

	if err := validate.FormatOf("LastGiftDate", "body", "date-time", m.LastGiftDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Constituent) validateLastTicketDate(formats strfmt.Registry) error {
	if swag.IsZero(m.LastTicketDate) { // not required
		return nil
	}

	if err := validate.FormatOf("LastTicketDate", "body", "date-time", m.LastTicketDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Constituent) validateMailIndicator(formats strfmt.Registry) error {
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

func (m *Constituent) validateNameStatus(formats strfmt.Registry) error {
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

func (m *Constituent) validateOriginalSource(formats strfmt.Registry) error {
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

func (m *Constituent) validatePhoneIndicator(formats strfmt.Registry) error {
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

func (m *Constituent) validatePrefix(formats strfmt.Registry) error {
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

func (m *Constituent) validateProtectionType(formats strfmt.Registry) error {
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

func (m *Constituent) validateSuffix(formats strfmt.Registry) error {
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

func (m *Constituent) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this constituent based on the context it is used
func (m *Constituent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConstituentType(ctx, formats); err != nil {
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

	if err := m.contextValidateProtectionType(ctx, formats); err != nil {
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

func (m *Constituent) contextValidateConstituentType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Constituent) contextValidateEmarketIndicator(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Constituent) contextValidateGender(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Constituent) contextValidateInactive(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Constituent) contextValidateInactiveReason(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Constituent) contextValidateMailIndicator(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Constituent) contextValidateNameStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Constituent) contextValidateOriginalSource(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Constituent) contextValidatePhoneIndicator(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Constituent) contextValidatePrefix(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Constituent) contextValidateProtectionType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Constituent) contextValidateSuffix(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Constituent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Constituent) UnmarshalBinary(b []byte) error {
	var res Constituent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
