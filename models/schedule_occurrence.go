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

// ScheduleOccurrence schedule occurrence
//
// swagger:model ScheduleOccurrence
type ScheduleOccurrence struct {

	// booking
	Booking *EntitySummary `json:"Booking,omitempty"`

	// booking assignment Id
	BookingAssignmentID int32 `json:"BookingAssignmentId,omitempty"`

	// constituent Id
	ConstituentID int32 `json:"ConstituentId,omitempty"`

	// count
	Count int32 `json:"Count,omitempty"`

	// end date time
	// Format: date-time
	EndDateTime *strfmt.DateTime `json:"EndDateTime,omitempty"`

	// is shared
	IsShared bool `json:"IsShared,omitempty"`

	// recurrence pattern
	RecurrencePattern int32 `json:"RecurrencePattern,omitempty"`

	// resource
	Resource *EntitySummary `json:"Resource,omitempty"`

	// resource category
	ResourceCategory *EntitySummary `json:"ResourceCategory,omitempty"`

	// resource type
	ResourceType *EntitySummary `json:"ResourceType,omitempty"`

	// schedule description
	ScheduleDescription string `json:"ScheduleDescription,omitempty"`

	// schedule Id
	ScheduleID int32 `json:"ScheduleId,omitempty"`

	// schedule type Id
	ScheduleTypeID int32 `json:"ScheduleTypeId,omitempty"`

	// start date time
	// Format: date-time
	StartDateTime *strfmt.DateTime `json:"StartDateTime,omitempty"`
}

// Validate validates this schedule occurrence
func (m *ScheduleOccurrence) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBooking(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScheduleOccurrence) validateBooking(formats strfmt.Registry) error {
	if swag.IsZero(m.Booking) { // not required
		return nil
	}

	if m.Booking != nil {
		if err := m.Booking.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Booking")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Booking")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduleOccurrence) validateEndDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("EndDateTime", "body", "date-time", m.EndDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ScheduleOccurrence) validateResource(formats strfmt.Registry) error {
	if swag.IsZero(m.Resource) { // not required
		return nil
	}

	if m.Resource != nil {
		if err := m.Resource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Resource")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduleOccurrence) validateResourceCategory(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceCategory) { // not required
		return nil
	}

	if m.ResourceCategory != nil {
		if err := m.ResourceCategory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ResourceCategory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ResourceCategory")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduleOccurrence) validateResourceType(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceType) { // not required
		return nil
	}

	if m.ResourceType != nil {
		if err := m.ResourceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ResourceType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ResourceType")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduleOccurrence) validateStartDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("StartDateTime", "body", "date-time", m.StartDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this schedule occurrence based on the context it is used
func (m *ScheduleOccurrence) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBooking(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScheduleOccurrence) contextValidateBooking(ctx context.Context, formats strfmt.Registry) error {

	if m.Booking != nil {

		if swag.IsZero(m.Booking) { // not required
			return nil
		}

		if err := m.Booking.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Booking")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Booking")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduleOccurrence) contextValidateResource(ctx context.Context, formats strfmt.Registry) error {

	if m.Resource != nil {

		if swag.IsZero(m.Resource) { // not required
			return nil
		}

		if err := m.Resource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Resource")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduleOccurrence) contextValidateResourceCategory(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceCategory != nil {

		if swag.IsZero(m.ResourceCategory) { // not required
			return nil
		}

		if err := m.ResourceCategory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ResourceCategory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ResourceCategory")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduleOccurrence) contextValidateResourceType(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceType != nil {

		if swag.IsZero(m.ResourceType) { // not required
			return nil
		}

		if err := m.ResourceType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ResourceType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ResourceType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScheduleOccurrence) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduleOccurrence) UnmarshalBinary(b []byte) error {
	var res ScheduleOccurrence
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
