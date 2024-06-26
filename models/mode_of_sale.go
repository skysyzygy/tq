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

// ModeOfSale mode of sale
//
// swagger:model ModeOfSale
type ModeOfSale struct {

	// allow unseated paid ind
	AllowUnseatedPaidInd bool `json:"AllowUnseatedPaidInd,omitempty"`

	// category
	Category *ModeOfSaleCategorySummary `json:"Category,omitempty"`

	// category required
	CategoryRequired bool `json:"CategoryRequired,omitempty"`

	// clear source no ind
	ClearSourceNoInd bool `json:"ClearSourceNoInd,omitempty"`

	// clear source on reload
	ClearSourceOnReload bool `json:"ClearSourceOnReload,omitempty"`

	// confirmation mode
	ConfirmationMode bool `json:"ConfirmationMode,omitempty"`

	// const link required
	ConstLinkRequired bool `json:"ConstLinkRequired,omitempty"`

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime *strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// default ack format Id
	DefaultAckFormatID int32 `json:"DefaultAckFormatId,omitempty"`

	// default channel Id
	DefaultChannelID int32 `json:"DefaultChannelId,omitempty"`

	// default delivery method Id
	DefaultDeliveryMethodID int32 `json:"DefaultDeliveryMethodId,omitempty"`

	// default header format Id
	DefaultHeaderFormatID int32 `json:"DefaultHeaderFormatId,omitempty"`

	// default sales layout Id
	DefaultSalesLayoutID int32 `json:"DefaultSalesLayoutId,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// edit date
	EditDate bool `json:"EditDate,omitempty"`

	// edit source on reload
	EditSourceOnReload bool `json:"EditSourceOnReload,omitempty"`

	// general public ind
	GeneralPublicInd bool `json:"GeneralPublicInd,omitempty"`

	// habo days
	HaboDays int32 `json:"HaboDays,omitempty"`

	// habo foreign
	HaboForeign bool `json:"HaboForeign,omitempty"`

	// hold until date
	// Format: date-time
	HoldUntilDate *strfmt.DateTime `json:"HoldUntilDate,omitempty"`

	// hold until days
	HoldUntilDays int32 `json:"HoldUntilDays,omitempty"`

	// hold until method
	HoldUntilMethod string `json:"HoldUntilMethod,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// inactive
	Inactive bool `json:"Inactive,omitempty"`

	// must be paid ind
	MustBePaidInd bool `json:"MustBePaidInd,omitempty"`

	// must be seated ind
	MustBeSeatedInd bool `json:"MustBeSeatedInd,omitempty"`

	// must be ticketed ind
	MustBeTicketedInd bool `json:"MustBeTicketedInd,omitempty"`

	// pricing rule set Id
	PricingRuleSetID int32 `json:"PricingRuleSetId,omitempty"`

	// sample for general public
	SampleForGeneralPublic int32 `json:"SampleForGeneralPublic,omitempty"`

	// sample for known constituent
	SampleForKnownConstituent int32 `json:"SampleForKnownConstituent,omitempty"`

	// sli auto delete ind
	SliAutoDeleteInd bool `json:"SliAutoDeleteInd,omitempty"`

	// start pkg or perf
	StartPkgOrPerf string `json:"StartPkgOrPerf,omitempty"`

	// subs summary required
	SubsSummaryRequired bool `json:"SubsSummaryRequired,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime *strfmt.DateTime `json:"UpdatedDateTime,omitempty"`
}

// Validate validates this mode of sale
func (m *ModeOfSale) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHoldUntilDate(formats); err != nil {
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

func (m *ModeOfSale) validateCategory(formats strfmt.Registry) error {
	if swag.IsZero(m.Category) { // not required
		return nil
	}

	if m.Category != nil {
		if err := m.Category.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Category")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Category")
			}
			return err
		}
	}

	return nil
}

func (m *ModeOfSale) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModeOfSale) validateHoldUntilDate(formats strfmt.Registry) error {
	if swag.IsZero(m.HoldUntilDate) { // not required
		return nil
	}

	if err := validate.FormatOf("HoldUntilDate", "body", "date-time", m.HoldUntilDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModeOfSale) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this mode of sale based on the context it is used
func (m *ModeOfSale) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModeOfSale) contextValidateCategory(ctx context.Context, formats strfmt.Registry) error {

	if m.Category != nil {

		if swag.IsZero(m.Category) { // not required
			return nil
		}

		if err := m.Category.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Category")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Category")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModeOfSale) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModeOfSale) UnmarshalBinary(b []byte) error {
	var res ModeOfSale
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
