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

// AttendanceHistory attendance history
//
// swagger:model AttendanceHistory
type AttendanceHistory struct {

	// admission adult
	AdmissionAdult int32 `json:"AdmissionAdult,omitempty"`

	// admission child
	AdmissionChild int32 `json:"AdmissionChild,omitempty"`

	// admission other
	AdmissionOther int32 `json:"AdmissionOther,omitempty"`

	// area
	Area *EntitySummary `json:"Area,omitempty"`

	// attend date time
	// Format: date-time
	AttendDateTime *strfmt.DateTime `json:"AttendDateTime,omitempty"`

	// composite Id
	CompositeID int32 `json:"CompositeId,omitempty"`

	// constituent
	Constituent *Entity `json:"Constituent,omitempty"`

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime *strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// device
	Device string `json:"Device,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// membership
	Membership int32 `json:"Membership,omitempty"`

	// membership level
	MembershipLevel *HistoryMembershipLevelSummary `json:"MembershipLevel,omitempty"`

	// order Id
	OrderID int32 `json:"OrderId,omitempty"`

	// performance
	Performance *HistoryPerformanceSummary `json:"Performance,omitempty"`

	// recipient
	Recipient *ConstituentDisplaySummary `json:"Recipient,omitempty"`

	// scan source
	ScanSource string `json:"ScanSource,omitempty"`

	// ticket Id
	TicketID int32 `json:"TicketId,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime *strfmt.DateTime `json:"UpdatedDateTime,omitempty"`

	// zone
	Zone *HistoryZoneSummary `json:"Zone,omitempty"`
}

// Validate validates this attendance history
func (m *AttendanceHistory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArea(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttendDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConstituent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMembershipLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerformance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AttendanceHistory) validateArea(formats strfmt.Registry) error {
	if swag.IsZero(m.Area) { // not required
		return nil
	}

	if m.Area != nil {
		if err := m.Area.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Area")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Area")
			}
			return err
		}
	}

	return nil
}

func (m *AttendanceHistory) validateAttendDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.AttendDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("AttendDateTime", "body", "date-time", m.AttendDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AttendanceHistory) validateConstituent(formats strfmt.Registry) error {
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

func (m *AttendanceHistory) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AttendanceHistory) validateMembershipLevel(formats strfmt.Registry) error {
	if swag.IsZero(m.MembershipLevel) { // not required
		return nil
	}

	if m.MembershipLevel != nil {
		if err := m.MembershipLevel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MembershipLevel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MembershipLevel")
			}
			return err
		}
	}

	return nil
}

func (m *AttendanceHistory) validatePerformance(formats strfmt.Registry) error {
	if swag.IsZero(m.Performance) { // not required
		return nil
	}

	if m.Performance != nil {
		if err := m.Performance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Performance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Performance")
			}
			return err
		}
	}

	return nil
}

func (m *AttendanceHistory) validateRecipient(formats strfmt.Registry) error {
	if swag.IsZero(m.Recipient) { // not required
		return nil
	}

	if m.Recipient != nil {
		if err := m.Recipient.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Recipient")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Recipient")
			}
			return err
		}
	}

	return nil
}

func (m *AttendanceHistory) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AttendanceHistory) validateZone(formats strfmt.Registry) error {
	if swag.IsZero(m.Zone) { // not required
		return nil
	}

	if m.Zone != nil {
		if err := m.Zone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Zone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Zone")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this attendance history based on the context it is used
func (m *AttendanceHistory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArea(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConstituent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMembershipLevel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePerformance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecipient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateZone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AttendanceHistory) contextValidateArea(ctx context.Context, formats strfmt.Registry) error {

	if m.Area != nil {

		if swag.IsZero(m.Area) { // not required
			return nil
		}

		if err := m.Area.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Area")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Area")
			}
			return err
		}
	}

	return nil
}

func (m *AttendanceHistory) contextValidateConstituent(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AttendanceHistory) contextValidateMembershipLevel(ctx context.Context, formats strfmt.Registry) error {

	if m.MembershipLevel != nil {

		if swag.IsZero(m.MembershipLevel) { // not required
			return nil
		}

		if err := m.MembershipLevel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MembershipLevel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MembershipLevel")
			}
			return err
		}
	}

	return nil
}

func (m *AttendanceHistory) contextValidatePerformance(ctx context.Context, formats strfmt.Registry) error {

	if m.Performance != nil {

		if swag.IsZero(m.Performance) { // not required
			return nil
		}

		if err := m.Performance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Performance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Performance")
			}
			return err
		}
	}

	return nil
}

func (m *AttendanceHistory) contextValidateRecipient(ctx context.Context, formats strfmt.Registry) error {

	if m.Recipient != nil {

		if swag.IsZero(m.Recipient) { // not required
			return nil
		}

		if err := m.Recipient.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Recipient")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Recipient")
			}
			return err
		}
	}

	return nil
}

func (m *AttendanceHistory) contextValidateZone(ctx context.Context, formats strfmt.Registry) error {

	if m.Zone != nil {

		if swag.IsZero(m.Zone) { // not required
			return nil
		}

		if err := m.Zone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Zone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Zone")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AttendanceHistory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AttendanceHistory) UnmarshalBinary(b []byte) error {
	var res AttendanceHistory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
