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

// ZoneMapSummary zone map summary
//
// swagger:model ZoneMapSummary
type ZoneMapSummary struct {

	// description
	Description string `json:"Description,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// inactive
	Inactive bool `json:"Inactive,omitempty"`

	// seat map
	SeatMap *SeatMapSummary `json:"SeatMap,omitempty"`
}

// Validate validates this zone map summary
func (m *ZoneMapSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSeatMap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZoneMapSummary) validateSeatMap(formats strfmt.Registry) error {
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

// ContextValidate validate this zone map summary based on the context it is used
func (m *ZoneMapSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSeatMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZoneMapSummary) contextValidateSeatMap(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *ZoneMapSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ZoneMapSummary) UnmarshalBinary(b []byte) error {
	var res ZoneMapSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
