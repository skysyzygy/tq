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

// PerformancePriceLayerSummary performance price layer summary
//
// swagger:model PerformancePriceLayerSummary
type PerformancePriceLayerSummary struct {

	// default designation code Id
	DefaultDesignationCodeID int32 `json:"DefaultDesignationCodeId,omitempty"`

	// default gl account Id
	DefaultGlAccountID int32 `json:"DefaultGlAccountId,omitempty"`

	// default resale account Id
	DefaultResaleAccountID int32 `json:"DefaultResaleAccountId,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// effective date time
	// Format: date-time
	EffectiveDateTime *strfmt.DateTime `json:"EffectiveDateTime,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// performance Id
	PerformanceID int32 `json:"PerformanceId,omitempty"`

	// performance price types
	PerformancePriceTypes []*PerformancePriceTypeSummary `json:"PerformancePriceTypes"`

	// price layer type Id
	PriceLayerTypeID int32 `json:"PriceLayerTypeId,omitempty"`

	// template Id
	TemplateID int32 `json:"TemplateId,omitempty"`
}

// Validate validates this performance price layer summary
func (m *PerformancePriceLayerSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEffectiveDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerformancePriceTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformancePriceLayerSummary) validateEffectiveDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EffectiveDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("EffectiveDateTime", "body", "date-time", m.EffectiveDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PerformancePriceLayerSummary) validatePerformancePriceTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.PerformancePriceTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.PerformancePriceTypes); i++ {
		if swag.IsZero(m.PerformancePriceTypes[i]) { // not required
			continue
		}

		if m.PerformancePriceTypes[i] != nil {
			if err := m.PerformancePriceTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PerformancePriceTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("PerformancePriceTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this performance price layer summary based on the context it is used
func (m *PerformancePriceLayerSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePerformancePriceTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformancePriceLayerSummary) contextValidatePerformancePriceTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PerformancePriceTypes); i++ {

		if m.PerformancePriceTypes[i] != nil {

			if swag.IsZero(m.PerformancePriceTypes[i]) { // not required
				return nil
			}

			if err := m.PerformancePriceTypes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PerformancePriceTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("PerformancePriceTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformancePriceLayerSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformancePriceLayerSummary) UnmarshalBinary(b []byte) error {
	var res PerformancePriceLayerSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
