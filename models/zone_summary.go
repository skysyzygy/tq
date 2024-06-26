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

// ZoneSummary zone summary
//
// swagger:model ZoneSummary
type ZoneSummary struct {

	// abbreviation
	Abbreviation string `json:"Abbreviation,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// rank
	Rank int32 `json:"Rank,omitempty"`

	// short description
	ShortDescription string `json:"ShortDescription,omitempty"`

	// zone group
	ZoneGroup *ZoneGroupSummary `json:"ZoneGroup,omitempty"`

	// zone legend
	ZoneLegend string `json:"ZoneLegend,omitempty"`

	// zone map Id
	ZoneMapID int32 `json:"ZoneMapId,omitempty"`

	// zone time
	ZoneTime string `json:"ZoneTime,omitempty"`
}

// Validate validates this zone summary
func (m *ZoneSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateZoneGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZoneSummary) validateZoneGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.ZoneGroup) { // not required
		return nil
	}

	if m.ZoneGroup != nil {
		if err := m.ZoneGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ZoneGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ZoneGroup")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this zone summary based on the context it is used
func (m *ZoneSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateZoneGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZoneSummary) contextValidateZoneGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.ZoneGroup != nil {

		if swag.IsZero(m.ZoneGroup) { // not required
			return nil
		}

		if err := m.ZoneGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ZoneGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ZoneGroup")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ZoneSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ZoneSummary) UnmarshalBinary(b []byte) error {
	var res ZoneSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
