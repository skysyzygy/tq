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

// CampaignFund campaign fund
//
// swagger:model CampaignFund
type CampaignFund struct {

	// campaign
	Campaign *CampaignSummary `json:"Campaign,omitempty"`

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime *strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// edit indicator
	EditIndicator bool `json:"EditIndicator,omitempty"`

	// end date time
	// Format: date-time
	EndDateTime *strfmt.DateTime `json:"EndDateTime,omitempty"`

	// fund
	Fund *FundSummary `json:"Fund,omitempty"`

	// goal amount
	GoalAmount float64 `json:"GoalAmount,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// start date time
	// Format: date-time
	StartDateTime *strfmt.DateTime `json:"StartDateTime,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime *strfmt.DateTime `json:"UpdatedDateTime,omitempty"`
}

// Validate validates this campaign fund
func (m *CampaignFund) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCampaign(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFund(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDateTime(formats); err != nil {
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

func (m *CampaignFund) validateCampaign(formats strfmt.Registry) error {
	if swag.IsZero(m.Campaign) { // not required
		return nil
	}

	if m.Campaign != nil {
		if err := m.Campaign.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Campaign")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Campaign")
			}
			return err
		}
	}

	return nil
}

func (m *CampaignFund) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CampaignFund) validateEndDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("EndDateTime", "body", "date-time", m.EndDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CampaignFund) validateFund(formats strfmt.Registry) error {
	if swag.IsZero(m.Fund) { // not required
		return nil
	}

	if m.Fund != nil {
		if err := m.Fund.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Fund")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Fund")
			}
			return err
		}
	}

	return nil
}

func (m *CampaignFund) validateStartDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("StartDateTime", "body", "date-time", m.StartDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CampaignFund) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this campaign fund based on the context it is used
func (m *CampaignFund) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCampaign(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFund(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CampaignFund) contextValidateCampaign(ctx context.Context, formats strfmt.Registry) error {

	if m.Campaign != nil {

		if swag.IsZero(m.Campaign) { // not required
			return nil
		}

		if err := m.Campaign.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Campaign")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Campaign")
			}
			return err
		}
	}

	return nil
}

func (m *CampaignFund) contextValidateFund(ctx context.Context, formats strfmt.Registry) error {

	if m.Fund != nil {

		if swag.IsZero(m.Fund) { // not required
			return nil
		}

		if err := m.Fund.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Fund")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Fund")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CampaignFund) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CampaignFund) UnmarshalBinary(b []byte) error {
	var res CampaignFund
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
