// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HoldCodeDetail hold code detail
//
// swagger:model HoldCodeDetail
type HoldCodeDetail struct {

	// description
	Description string `json:"Description,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// is blackout
	IsBlackout bool `json:"IsBlackout,omitempty"`

	// legend
	Legend string `json:"Legend,omitempty"`
}

// Validate validates this hold code detail
func (m *HoldCodeDetail) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hold code detail based on context it is used
func (m *HoldCodeDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HoldCodeDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HoldCodeDetail) UnmarshalBinary(b []byte) error {
	var res HoldCodeDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
