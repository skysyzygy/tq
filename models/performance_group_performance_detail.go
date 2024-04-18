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

// PerformanceGroupPerformanceDetail performance group performance detail
//
// swagger:model PerformanceGroupPerformanceDetail
type PerformanceGroupPerformanceDetail struct {

	// available indicator
	AvailableIndicator bool `json:"AvailableIndicator,omitempty"`

	// display by zone
	DisplayByZone bool `json:"DisplayByZone,omitempty"`

	// facility
	Facility string `json:"Facility,omitempty"`

	// inactive
	Inactive bool `json:"Inactive,omitempty"`

	// performance code
	PerformanceCode string `json:"PerformanceCode,omitempty"`

	// performance date time
	// Format: date-time
	PerformanceDateTime strfmt.DateTime `json:"PerformanceDateTime,omitempty"`

	// performance description
	PerformanceDescription string `json:"PerformanceDescription,omitempty"`

	// performance Id
	PerformanceID int32 `json:"PerformanceId,omitempty"`

	// sales notes
	SalesNotes string `json:"SalesNotes,omitempty"`

	// zone map Id
	ZoneMapID int32 `json:"ZoneMapId,omitempty"`

	// zones
	Zones []*PerformanceZoneAvailability `json:"Zones"`
}

// Validate validates this performance group performance detail
func (m *PerformanceGroupPerformanceDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePerformanceDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZones(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceGroupPerformanceDetail) validatePerformanceDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.PerformanceDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("PerformanceDateTime", "body", "date-time", m.PerformanceDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceGroupPerformanceDetail) validateZones(formats strfmt.Registry) error {
	if swag.IsZero(m.Zones) { // not required
		return nil
	}

	for i := 0; i < len(m.Zones); i++ {
		if swag.IsZero(m.Zones[i]) { // not required
			continue
		}

		if m.Zones[i] != nil {
			if err := m.Zones[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Zones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Zones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this performance group performance detail based on the context it is used
func (m *PerformanceGroupPerformanceDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateZones(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceGroupPerformanceDetail) contextValidateZones(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Zones); i++ {

		if m.Zones[i] != nil {

			if swag.IsZero(m.Zones[i]) { // not required
				return nil
			}

			if err := m.Zones[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Zones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Zones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceGroupPerformanceDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceGroupPerformanceDetail) UnmarshalBinary(b []byte) error {
	var res PerformanceGroupPerformanceDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}