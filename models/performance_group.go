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

// PerformanceGroup performance group
//
// swagger:model PerformanceGroup
type PerformanceGroup struct {

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// display by zone
	DisplayByZone bool `json:"DisplayByZone,omitempty"`

	// facility
	Facility *FacilitySummary `json:"Facility,omitempty"`

	// fixed seat indicator
	FixedSeatIndicator bool `json:"FixedSeatIndicator,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// season
	Season *SeasonSummary `json:"Season,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime strfmt.DateTime `json:"UpdatedDateTime,omitempty"`

	// zone map
	ZoneMap *ZoneMapSummary `json:"ZoneMap,omitempty"`
}

// Validate validates this performance group
func (m *PerformanceGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFacility(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZoneMap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceGroup) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceGroup) validateFacility(formats strfmt.Registry) error {
	if swag.IsZero(m.Facility) { // not required
		return nil
	}

	if m.Facility != nil {
		if err := m.Facility.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Facility")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Facility")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceGroup) validateSeason(formats strfmt.Registry) error {
	if swag.IsZero(m.Season) { // not required
		return nil
	}

	if m.Season != nil {
		if err := m.Season.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Season")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Season")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceGroup) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceGroup) validateZoneMap(formats strfmt.Registry) error {
	if swag.IsZero(m.ZoneMap) { // not required
		return nil
	}

	if m.ZoneMap != nil {
		if err := m.ZoneMap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ZoneMap")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ZoneMap")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this performance group based on the context it is used
func (m *PerformanceGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFacility(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSeason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateZoneMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceGroup) contextValidateFacility(ctx context.Context, formats strfmt.Registry) error {

	if m.Facility != nil {

		if swag.IsZero(m.Facility) { // not required
			return nil
		}

		if err := m.Facility.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Facility")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Facility")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceGroup) contextValidateSeason(ctx context.Context, formats strfmt.Registry) error {

	if m.Season != nil {

		if swag.IsZero(m.Season) { // not required
			return nil
		}

		if err := m.Season.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Season")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Season")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceGroup) contextValidateZoneMap(ctx context.Context, formats strfmt.Registry) error {

	if m.ZoneMap != nil {

		if swag.IsZero(m.ZoneMap) { // not required
			return nil
		}

		if err := m.ZoneMap.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ZoneMap")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ZoneMap")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceGroup) UnmarshalBinary(b []byte) error {
	var res PerformanceGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
