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

// CartPerformance cart performance
//
// swagger:model CartPerformance
type CartPerformance struct {

	// best seat map
	BestSeatMap *CartBestSeatMap `json:"BestSeatMap,omitempty"`

	// code
	Code string `json:"Code,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// duration
	Duration int32 `json:"Duration,omitempty"`

	// facility
	Facility *EntitySummary `json:"Facility,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// performance date time
	// Format: date-time
	PerformanceDateTime strfmt.DateTime `json:"PerformanceDateTime,omitempty"`

	// production season
	ProductionSeason *EntitySummary `json:"ProductionSeason,omitempty"`

	// season
	Season *EntitySummary `json:"Season,omitempty"`

	// time slot
	TimeSlot *EntitySummary `json:"TimeSlot,omitempty"`

	// type
	Type *EntitySummary `json:"Type,omitempty"`

	// zone map
	ZoneMap *CartZoneMap `json:"ZoneMap,omitempty"`
}

// Validate validates this cart performance
func (m *CartPerformance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBestSeatMap(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFacility(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerformanceDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductionSeason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeSlot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
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

func (m *CartPerformance) validateBestSeatMap(formats strfmt.Registry) error {
	if swag.IsZero(m.BestSeatMap) { // not required
		return nil
	}

	if m.BestSeatMap != nil {
		if err := m.BestSeatMap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BestSeatMap")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("BestSeatMap")
			}
			return err
		}
	}

	return nil
}

func (m *CartPerformance) validateFacility(formats strfmt.Registry) error {
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

func (m *CartPerformance) validatePerformanceDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.PerformanceDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("PerformanceDateTime", "body", "date-time", m.PerformanceDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CartPerformance) validateProductionSeason(formats strfmt.Registry) error {
	if swag.IsZero(m.ProductionSeason) { // not required
		return nil
	}

	if m.ProductionSeason != nil {
		if err := m.ProductionSeason.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ProductionSeason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ProductionSeason")
			}
			return err
		}
	}

	return nil
}

func (m *CartPerformance) validateSeason(formats strfmt.Registry) error {
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

func (m *CartPerformance) validateTimeSlot(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeSlot) { // not required
		return nil
	}

	if m.TimeSlot != nil {
		if err := m.TimeSlot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TimeSlot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TimeSlot")
			}
			return err
		}
	}

	return nil
}

func (m *CartPerformance) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Type")
			}
			return err
		}
	}

	return nil
}

func (m *CartPerformance) validateZoneMap(formats strfmt.Registry) error {
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

// ContextValidate validate this cart performance based on the context it is used
func (m *CartPerformance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBestSeatMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFacility(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProductionSeason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSeason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimeSlot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
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

func (m *CartPerformance) contextValidateBestSeatMap(ctx context.Context, formats strfmt.Registry) error {

	if m.BestSeatMap != nil {

		if swag.IsZero(m.BestSeatMap) { // not required
			return nil
		}

		if err := m.BestSeatMap.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BestSeatMap")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("BestSeatMap")
			}
			return err
		}
	}

	return nil
}

func (m *CartPerformance) contextValidateFacility(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CartPerformance) contextValidateProductionSeason(ctx context.Context, formats strfmt.Registry) error {

	if m.ProductionSeason != nil {

		if swag.IsZero(m.ProductionSeason) { // not required
			return nil
		}

		if err := m.ProductionSeason.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ProductionSeason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ProductionSeason")
			}
			return err
		}
	}

	return nil
}

func (m *CartPerformance) contextValidateSeason(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CartPerformance) contextValidateTimeSlot(ctx context.Context, formats strfmt.Registry) error {

	if m.TimeSlot != nil {

		if swag.IsZero(m.TimeSlot) { // not required
			return nil
		}

		if err := m.TimeSlot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TimeSlot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TimeSlot")
			}
			return err
		}
	}

	return nil
}

func (m *CartPerformance) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {

		if swag.IsZero(m.Type) { // not required
			return nil
		}

		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Type")
			}
			return err
		}
	}

	return nil
}

func (m *CartPerformance) contextValidateZoneMap(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CartPerformance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CartPerformance) UnmarshalBinary(b []byte) error {
	var res CartPerformance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
