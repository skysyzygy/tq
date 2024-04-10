// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PricingRuleMessageSummary pricing rule message summary
//
// swagger:model PricingRuleMessageSummary
type PricingRuleMessageSummary struct {

	// Id
	ID int32 `json:"Id,omitempty"`

	// is from message only rule
	IsFromMessageOnlyRule bool `json:"IsFromMessageOnlyRule,omitempty"`

	// message
	Message string `json:"Message,omitempty"`

	// message type
	MessageType *PricingRuleMessageTypeSummary `json:"MessageType,omitempty"`

	// pricing rule
	PricingRule *EntitySummary `json:"PricingRule,omitempty"`
}

// Validate validates this pricing rule message summary
func (m *PricingRuleMessageSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessageType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePricingRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PricingRuleMessageSummary) validateMessageType(formats strfmt.Registry) error {
	if swag.IsZero(m.MessageType) { // not required
		return nil
	}

	if m.MessageType != nil {
		if err := m.MessageType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MessageType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MessageType")
			}
			return err
		}
	}

	return nil
}

func (m *PricingRuleMessageSummary) validatePricingRule(formats strfmt.Registry) error {
	if swag.IsZero(m.PricingRule) { // not required
		return nil
	}

	if m.PricingRule != nil {
		if err := m.PricingRule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PricingRule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PricingRule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this pricing rule message summary based on the context it is used
func (m *PricingRuleMessageSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMessageType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePricingRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PricingRuleMessageSummary) contextValidateMessageType(ctx context.Context, formats strfmt.Registry) error {

	if m.MessageType != nil {

		if swag.IsZero(m.MessageType) { // not required
			return nil
		}

		if err := m.MessageType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MessageType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MessageType")
			}
			return err
		}
	}

	return nil
}

func (m *PricingRuleMessageSummary) contextValidatePricingRule(ctx context.Context, formats strfmt.Registry) error {

	if m.PricingRule != nil {

		if swag.IsZero(m.PricingRule) { // not required
			return nil
		}

		if err := m.PricingRule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PricingRule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PricingRule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PricingRuleMessageSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PricingRuleMessageSummary) UnmarshalBinary(b []byte) error {
	var res PricingRuleMessageSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
