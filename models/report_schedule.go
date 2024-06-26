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

// ReportSchedule report schedule
//
// swagger:model ReportSchedule
type ReportSchedule struct {

	// as of date time
	// Format: date-time
	AsOfDateTime *strfmt.DateTime `json:"AsOfDateTime,omitempty"`

	// control group
	ControlGroup *ControlGroupSummary `json:"ControlGroup,omitempty"`

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime *strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// day of week
	DayOfWeek int32 `json:"DayOfWeek,omitempty"`

	// day week number
	DayWeekNumber int32 `json:"DayWeekNumber,omitempty"`

	// end date
	// Format: date-time
	EndDate *strfmt.DateTime `json:"EndDate,omitempty"`

	// end time
	// Format: date-time
	EndTime *strfmt.DateTime `json:"EndTime,omitempty"`

	// header request
	HeaderRequest *ReportRequest `json:"HeaderRequest,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// inactive
	Inactive bool `json:"Inactive,omitempty"`

	// interval
	Interval int32 `json:"Interval,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// start date
	// Format: date-time
	StartDate *strfmt.DateTime `json:"StartDate,omitempty"`

	// start time
	// Format: date-time
	StartTime *strfmt.DateTime `json:"StartTime,omitempty"`

	// type
	Type string `json:"Type,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime *strfmt.DateTime `json:"UpdatedDateTime,omitempty"`
}

// Validate validates this report schedule
func (m *ReportSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsOfDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateControlGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeaderRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
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

func (m *ReportSchedule) validateAsOfDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.AsOfDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("AsOfDateTime", "body", "date-time", m.AsOfDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportSchedule) validateControlGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.ControlGroup) { // not required
		return nil
	}

	if m.ControlGroup != nil {
		if err := m.ControlGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ControlGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ControlGroup")
			}
			return err
		}
	}

	return nil
}

func (m *ReportSchedule) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportSchedule) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("EndDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportSchedule) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("EndTime", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportSchedule) validateHeaderRequest(formats strfmt.Registry) error {
	if swag.IsZero(m.HeaderRequest) { // not required
		return nil
	}

	if m.HeaderRequest != nil {
		if err := m.HeaderRequest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("HeaderRequest")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("HeaderRequest")
			}
			return err
		}
	}

	return nil
}

func (m *ReportSchedule) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("StartDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportSchedule) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("StartTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportSchedule) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this report schedule based on the context it is used
func (m *ReportSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateControlGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHeaderRequest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportSchedule) contextValidateControlGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.ControlGroup != nil {

		if swag.IsZero(m.ControlGroup) { // not required
			return nil
		}

		if err := m.ControlGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ControlGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ControlGroup")
			}
			return err
		}
	}

	return nil
}

func (m *ReportSchedule) contextValidateHeaderRequest(ctx context.Context, formats strfmt.Registry) error {

	if m.HeaderRequest != nil {

		if swag.IsZero(m.HeaderRequest) { // not required
			return nil
		}

		if err := m.HeaderRequest.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("HeaderRequest")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("HeaderRequest")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportSchedule) UnmarshalBinary(b []byte) error {
	var res ReportSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
