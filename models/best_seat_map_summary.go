// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BestSeatMapSummary best seat map summary
//
// swagger:model BestSeatMapSummary
type BestSeatMapSummary struct {

	// description
	Description string `json:"Description,omitempty"`

	// ga indicator
	GaIndicator string `json:"GaIndicator,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// is g a
	IsGA bool `json:"IsGA,omitempty"`
}

// Validate validates this best seat map summary
func (m *BestSeatMapSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this best seat map summary based on context it is used
func (m *BestSeatMapSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BestSeatMapSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BestSeatMapSummary) UnmarshalBinary(b []byte) error {
	var res BestSeatMapSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
