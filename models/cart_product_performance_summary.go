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

// CartProductPerformanceSummary cart product performance summary
//
// swagger:model CartProductPerformanceSummary
type CartProductPerformanceSummary struct {

	// display lines
	DisplayLines []*DisplayLine `json:"DisplayLines"`

	// has zone time
	HasZoneTime bool `json:"HasZoneTime,omitempty"`

	// line item Id
	LineItemID int32 `json:"LineItemId,omitempty"`

	// performance code
	PerformanceCode string `json:"PerformanceCode,omitempty"`

	// performance date time
	// Format: date-time
	PerformanceDateTime strfmt.DateTime `json:"PerformanceDateTime,omitempty"`

	// performance description
	PerformanceDescription string `json:"PerformanceDescription,omitempty"`

	// performance Id
	PerformanceID int32 `json:"PerformanceId,omitempty"`

	// source
	Source *EntitySummary `json:"Source,omitempty"`

	// total amount
	TotalAmount float64 `json:"TotalAmount,omitempty"`
}

// Validate validates this cart product performance summary
func (m *CartProductPerformanceSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayLines(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerformanceDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CartProductPerformanceSummary) validateDisplayLines(formats strfmt.Registry) error {
	if swag.IsZero(m.DisplayLines) { // not required
		return nil
	}

	for i := 0; i < len(m.DisplayLines); i++ {
		if swag.IsZero(m.DisplayLines[i]) { // not required
			continue
		}

		if m.DisplayLines[i] != nil {
			if err := m.DisplayLines[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DisplayLines" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DisplayLines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CartProductPerformanceSummary) validatePerformanceDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.PerformanceDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("PerformanceDateTime", "body", "date-time", m.PerformanceDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CartProductPerformanceSummary) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Source")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cart product performance summary based on the context it is used
func (m *CartProductPerformanceSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisplayLines(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CartProductPerformanceSummary) contextValidateDisplayLines(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DisplayLines); i++ {

		if m.DisplayLines[i] != nil {

			if swag.IsZero(m.DisplayLines[i]) { // not required
				return nil
			}

			if err := m.DisplayLines[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DisplayLines" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DisplayLines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CartProductPerformanceSummary) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {

		if swag.IsZero(m.Source) { // not required
			return nil
		}

		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CartProductPerformanceSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CartProductPerformanceSummary) UnmarshalBinary(b []byte) error {
	var res CartProductPerformanceSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
