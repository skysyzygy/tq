// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SchedulePatternTypeSummary schedule pattern type summary
//
// swagger:model SchedulePatternTypeSummary
type SchedulePatternTypeSummary struct {

	// description
	Description string `json:"Description,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`
}

// Validate validates this schedule pattern type summary
func (m *SchedulePatternTypeSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this schedule pattern type summary based on context it is used
func (m *SchedulePatternTypeSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SchedulePatternTypeSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchedulePatternTypeSummary) UnmarshalBinary(b []byte) error {
	var res SchedulePatternTypeSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
