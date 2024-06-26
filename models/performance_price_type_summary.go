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

// PerformancePriceTypeSummary performance price type summary
//
// swagger:model PerformancePriceTypeSummary
type PerformancePriceTypeSummary struct {

	// base indicator
	BaseIndicator bool `json:"BaseIndicator,omitempty"`

	// designation code Id
	DesignationCodeID int32 `json:"DesignationCodeId,omitempty"`

	// effective date time
	// Format: date-time
	EffectiveDateTime *strfmt.DateTime `json:"EffectiveDateTime,omitempty"`

	// end date time
	// Format: date-time
	EndDateTime *strfmt.DateTime `json:"EndDateTime,omitempty"`

	// gl account Id
	GlAccountID int32 `json:"GlAccountId,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// is within date range
	IsWithinDateRange bool `json:"IsWithinDateRange,omitempty"`

	// performance Id
	PerformanceID int32 `json:"PerformanceId,omitempty"`

	// performance price layer Id
	PerformancePriceLayerID int32 `json:"PerformancePriceLayerId,omitempty"`

	// price type Id
	PriceTypeID int32 `json:"PriceTypeId,omitempty"`

	// prices
	Prices []string `json:"Prices"`

	// resale account Id
	ResaleAccountID int32 `json:"ResaleAccountId,omitempty"`

	// start date time
	// Format: date-time
	StartDateTime *strfmt.DateTime `json:"StartDateTime,omitempty"`

	// ticket design Id
	TicketDesignID int32 `json:"TicketDesignId,omitempty"`
}

// Validate validates this performance price type summary
func (m *PerformancePriceTypeSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEffectiveDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDateTime(formats); err != nil {
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

func (m *PerformancePriceTypeSummary) validateEffectiveDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EffectiveDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("EffectiveDateTime", "body", "date-time", m.EffectiveDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PerformancePriceTypeSummary) validateEndDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("EndDateTime", "body", "date-time", m.EndDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PerformancePriceTypeSummary) validateStartDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("StartDateTime", "body", "date-time", m.StartDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this performance price type summary based on context it is used
func (m *PerformancePriceTypeSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PerformancePriceTypeSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformancePriceTypeSummary) UnmarshalBinary(b []byte) error {
	var res PerformancePriceTypeSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
