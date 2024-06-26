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

// PaymentPlanSchedulePreviewRequest payment plan schedule preview request
//
// swagger:model PaymentPlanSchedulePreviewRequest
type PaymentPlanSchedulePreviewRequest struct {

	// billing schedule Id
	BillingScheduleID int32 `json:"BillingScheduleId,omitempty"`

	// end date
	// Format: date-time
	EndDate *strfmt.DateTime `json:"EndDate,omitempty"`

	// override amount to schedule
	OverrideAmountToSchedule float64 `json:"OverrideAmountToSchedule,omitempty"`

	// start date
	// Format: date-time
	StartDate *strfmt.DateTime `json:"StartDate,omitempty"`
}

// Validate validates this payment plan schedule preview request
func (m *PaymentPlanSchedulePreviewRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentPlanSchedulePreviewRequest) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("EndDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentPlanSchedulePreviewRequest) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("StartDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this payment plan schedule preview request based on context it is used
func (m *PaymentPlanSchedulePreviewRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentPlanSchedulePreviewRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentPlanSchedulePreviewRequest) UnmarshalBinary(b []byte) error {
	var res PaymentPlanSchedulePreviewRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
