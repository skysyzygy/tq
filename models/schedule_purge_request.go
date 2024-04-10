// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SchedulePurgeRequest schedule purge request
//
// swagger:model SchedulePurgeRequest
type SchedulePurgeRequest struct {

	// ignore warnings
	IgnoreWarnings bool `json:"IgnoreWarnings,omitempty"`
}

// Validate validates this schedule purge request
func (m *SchedulePurgeRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this schedule purge request based on context it is used
func (m *SchedulePurgeRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SchedulePurgeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchedulePurgeRequest) UnmarshalBinary(b []byte) error {
	var res SchedulePurgeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
