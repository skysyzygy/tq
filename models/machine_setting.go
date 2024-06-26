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

// MachineSetting machine setting
//
// swagger:model MachineSetting
type MachineSetting struct {

	// card reader host
	CardReaderHost string `json:"CardReaderHost,omitempty"`

	// card reader port
	CardReaderPort int32 `json:"CardReaderPort,omitempty"`

	// card reader type
	CardReaderType *CardReaderTypeSummary `json:"CardReaderType,omitempty"`

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime *strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// inactive
	Inactive bool `json:"Inactive,omitempty"`

	// merchant Id
	MerchantID string `json:"MerchantId,omitempty"`

	// p x station
	PXStation string `json:"PXStation,omitempty"`

	// p x user key
	PXUserKey string `json:"PXUserKey,omitempty"`

	// p x user name
	PXUserName string `json:"PXUserName,omitempty"`

	// tessitura payments pos device
	TessituraPaymentsPosDevice string `json:"TessituraPaymentsPosDevice,omitempty"`

	// tessitura payments pos device model
	TessituraPaymentsPosDeviceModel string `json:"TessituraPaymentsPosDeviceModel,omitempty"`

	// tnspay software terminal
	TnspaySoftwareTerminal bool `json:"TnspaySoftwareTerminal,omitempty"`

	// tri p o s cloud configuration
	TriPOSCloudConfiguration *TriPOSCloudConfiguration `json:"TriPOSCloudConfiguration,omitempty"`

	// tripos lane
	TriposLane int32 `json:"TriposLane,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime *strfmt.DateTime `json:"UpdatedDateTime,omitempty"`

	// workstation name
	WorkstationName string `json:"WorkstationName,omitempty"`
}

// Validate validates this machine setting
func (m *MachineSetting) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCardReaderType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTriPOSCloudConfiguration(formats); err != nil {
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

func (m *MachineSetting) validateCardReaderType(formats strfmt.Registry) error {
	if swag.IsZero(m.CardReaderType) { // not required
		return nil
	}

	if m.CardReaderType != nil {
		if err := m.CardReaderType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CardReaderType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CardReaderType")
			}
			return err
		}
	}

	return nil
}

func (m *MachineSetting) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MachineSetting) validateTriPOSCloudConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.TriPOSCloudConfiguration) { // not required
		return nil
	}

	if m.TriPOSCloudConfiguration != nil {
		if err := m.TriPOSCloudConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TriPOSCloudConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TriPOSCloudConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *MachineSetting) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this machine setting based on the context it is used
func (m *MachineSetting) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCardReaderType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTriPOSCloudConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MachineSetting) contextValidateCardReaderType(ctx context.Context, formats strfmt.Registry) error {

	if m.CardReaderType != nil {

		if swag.IsZero(m.CardReaderType) { // not required
			return nil
		}

		if err := m.CardReaderType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CardReaderType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CardReaderType")
			}
			return err
		}
	}

	return nil
}

func (m *MachineSetting) contextValidateTriPOSCloudConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.TriPOSCloudConfiguration != nil {

		if swag.IsZero(m.TriPOSCloudConfiguration) { // not required
			return nil
		}

		if err := m.TriPOSCloudConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TriPOSCloudConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TriPOSCloudConfiguration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MachineSetting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MachineSetting) UnmarshalBinary(b []byte) error {
	var res MachineSetting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
