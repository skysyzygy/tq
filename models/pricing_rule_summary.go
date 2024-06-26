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

// PricingRuleSummary pricing rule summary
//
// swagger:model PricingRuleSummary
type PricingRuleSummary struct {

	// description
	Description string `json:"Description,omitempty"`

	// end date time
	// Format: date-time
	EndDateTime *strfmt.DateTime `json:"EndDateTime,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// inactive
	Inactive bool `json:"Inactive,omitempty"`

	// rule action
	RuleAction int32 `json:"RuleAction,omitempty"`

	// rule category
	RuleCategory *PricingRuleCategorySummary `json:"RuleCategory,omitempty"`

	// rule type
	RuleType *PricingRuleTypeSummary `json:"RuleType,omitempty"`

	// start date time
	// Format: date-time
	StartDateTime *strfmt.DateTime `json:"StartDateTime,omitempty"`
}

// Validate validates this pricing rule summary
func (m *PricingRuleSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleType(formats); err != nil {
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

func (m *PricingRuleSummary) validateEndDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("EndDateTime", "body", "date-time", m.EndDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PricingRuleSummary) validateRuleCategory(formats strfmt.Registry) error {
	if swag.IsZero(m.RuleCategory) { // not required
		return nil
	}

	if m.RuleCategory != nil {
		if err := m.RuleCategory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RuleCategory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("RuleCategory")
			}
			return err
		}
	}

	return nil
}

func (m *PricingRuleSummary) validateRuleType(formats strfmt.Registry) error {
	if swag.IsZero(m.RuleType) { // not required
		return nil
	}

	if m.RuleType != nil {
		if err := m.RuleType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RuleType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("RuleType")
			}
			return err
		}
	}

	return nil
}

func (m *PricingRuleSummary) validateStartDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("StartDateTime", "body", "date-time", m.StartDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this pricing rule summary based on the context it is used
func (m *PricingRuleSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRuleCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRuleType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PricingRuleSummary) contextValidateRuleCategory(ctx context.Context, formats strfmt.Registry) error {

	if m.RuleCategory != nil {

		if swag.IsZero(m.RuleCategory) { // not required
			return nil
		}

		if err := m.RuleCategory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RuleCategory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("RuleCategory")
			}
			return err
		}
	}

	return nil
}

func (m *PricingRuleSummary) contextValidateRuleType(ctx context.Context, formats strfmt.Registry) error {

	if m.RuleType != nil {

		if swag.IsZero(m.RuleType) { // not required
			return nil
		}

		if err := m.RuleType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RuleType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("RuleType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PricingRuleSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PricingRuleSummary) UnmarshalBinary(b []byte) error {
	var res PricingRuleSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
