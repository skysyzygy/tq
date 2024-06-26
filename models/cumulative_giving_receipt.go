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

// CumulativeGivingReceipt cumulative giving receipt
//
// swagger:model CumulativeGivingReceipt
type CumulativeGivingReceipt struct {

	// attach pdf
	AttachPdf bool `json:"AttachPdf,omitempty"`

	// constituent
	Constituent *ConstituentDisplaySummary `json:"Constituent,omitempty"`

	// contribution total amount
	ContributionTotalAmount float64 `json:"ContributionTotalAmount,omitempty"`

	// control group
	ControlGroup *ControlGroupSummary `json:"ControlGroup,omitempty"`

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime *strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// date sent
	// Format: date-time
	DateSent *strfmt.DateTime `json:"DateSent,omitempty"`

	// email address
	EmailAddress string `json:"EmailAddress,omitempty"`

	// end date time
	// Format: date-time
	EndDateTime *strfmt.DateTime `json:"EndDateTime,omitempty"`

	// format
	Format *Entity `json:"Format,omitempty"`

	// funds
	Funds string `json:"Funds,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// session Id
	SessionID int32 `json:"SessionId,omitempty"`

	// start date time
	// Format: date-time
	StartDateTime *strfmt.DateTime `json:"StartDateTime,omitempty"`

	// tax receipt Id
	TaxReceiptID int32 `json:"TaxReceiptId,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime *strfmt.DateTime `json:"UpdatedDateTime,omitempty"`

	// year
	Year int32 `json:"Year,omitempty"`
}

// Validate validates this cumulative giving receipt
func (m *CumulativeGivingReceipt) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConstituent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateControlGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateSent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormat(formats); err != nil {
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

func (m *CumulativeGivingReceipt) validateConstituent(formats strfmt.Registry) error {
	if swag.IsZero(m.Constituent) { // not required
		return nil
	}

	if m.Constituent != nil {
		if err := m.Constituent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Constituent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Constituent")
			}
			return err
		}
	}

	return nil
}

func (m *CumulativeGivingReceipt) validateControlGroup(formats strfmt.Registry) error {
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

func (m *CumulativeGivingReceipt) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CumulativeGivingReceipt) validateDateSent(formats strfmt.Registry) error {
	if swag.IsZero(m.DateSent) { // not required
		return nil
	}

	if err := validate.FormatOf("DateSent", "body", "date-time", m.DateSent.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CumulativeGivingReceipt) validateEndDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("EndDateTime", "body", "date-time", m.EndDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CumulativeGivingReceipt) validateFormat(formats strfmt.Registry) error {
	if swag.IsZero(m.Format) { // not required
		return nil
	}

	if m.Format != nil {
		if err := m.Format.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Format")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Format")
			}
			return err
		}
	}

	return nil
}

func (m *CumulativeGivingReceipt) validateStartDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("StartDateTime", "body", "date-time", m.StartDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CumulativeGivingReceipt) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cumulative giving receipt based on the context it is used
func (m *CumulativeGivingReceipt) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConstituent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateControlGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFormat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CumulativeGivingReceipt) contextValidateConstituent(ctx context.Context, formats strfmt.Registry) error {

	if m.Constituent != nil {

		if swag.IsZero(m.Constituent) { // not required
			return nil
		}

		if err := m.Constituent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Constituent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Constituent")
			}
			return err
		}
	}

	return nil
}

func (m *CumulativeGivingReceipt) contextValidateControlGroup(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CumulativeGivingReceipt) contextValidateFormat(ctx context.Context, formats strfmt.Registry) error {

	if m.Format != nil {

		if swag.IsZero(m.Format) { // not required
			return nil
		}

		if err := m.Format.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Format")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Format")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CumulativeGivingReceipt) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CumulativeGivingReceipt) UnmarshalBinary(b []byte) error {
	var res CumulativeGivingReceipt
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
