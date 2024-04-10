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

// Facility facility
//
// swagger:model Facility
type Facility struct {

	// control group
	ControlGroup *ControlGroupSummary `json:"ControlGroup,omitempty"`

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// default best seat map Id
	DefaultBestSeatMapID int32 `json:"DefaultBestSeatMapId,omitempty"`

	// default zone map Id
	DefaultZoneMapID int32 `json:"DefaultZoneMapId,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// seat map
	SeatMap *SeatMap `json:"SeatMap,omitempty"`

	// theater
	Theater *TheaterSummary `json:"Theater,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime strfmt.DateTime `json:"UpdatedDateTime,omitempty"`
}

// Validate validates this facility
func (m *Facility) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateControlGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeatMap(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTheater(formats); err != nil {
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

func (m *Facility) validateControlGroup(formats strfmt.Registry) error {
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

func (m *Facility) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Facility) validateSeatMap(formats strfmt.Registry) error {
	if swag.IsZero(m.SeatMap) { // not required
		return nil
	}

	if m.SeatMap != nil {
		if err := m.SeatMap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SeatMap")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("SeatMap")
			}
			return err
		}
	}

	return nil
}

func (m *Facility) validateTheater(formats strfmt.Registry) error {
	if swag.IsZero(m.Theater) { // not required
		return nil
	}

	if m.Theater != nil {
		if err := m.Theater.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Theater")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Theater")
			}
			return err
		}
	}

	return nil
}

func (m *Facility) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this facility based on the context it is used
func (m *Facility) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateControlGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSeatMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTheater(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Facility) contextValidateControlGroup(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Facility) contextValidateSeatMap(ctx context.Context, formats strfmt.Registry) error {

	if m.SeatMap != nil {

		if swag.IsZero(m.SeatMap) { // not required
			return nil
		}

		if err := m.SeatMap.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SeatMap")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("SeatMap")
			}
			return err
		}
	}

	return nil
}

func (m *Facility) contextValidateTheater(ctx context.Context, formats strfmt.Registry) error {

	if m.Theater != nil {

		if swag.IsZero(m.Theater) { // not required
			return nil
		}

		if err := m.Theater.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Theater")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Theater")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Facility) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Facility) UnmarshalBinary(b []byte) error {
	var res Facility
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
