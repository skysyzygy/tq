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

// ExpireSeatHoldRequest expire seat hold request
//
// swagger:model ExpireSeatHoldRequest
type ExpireSeatHoldRequest struct {

	// expire as of date
	// Format: date-time
	ExpireAsOfDate strfmt.DateTime `json:"ExpireAsOfDate,omitempty"`

	// seat ids
	SeatIds string `json:"SeatIds,omitempty"`
}

// Validate validates this expire seat hold request
func (m *ExpireSeatHoldRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpireAsOfDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExpireSeatHoldRequest) validateExpireAsOfDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpireAsOfDate) { // not required
		return nil
	}

	if err := validate.FormatOf("ExpireAsOfDate", "body", "date-time", m.ExpireAsOfDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this expire seat hold request based on context it is used
func (m *ExpireSeatHoldRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExpireSeatHoldRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExpireSeatHoldRequest) UnmarshalBinary(b []byte) error {
	var res ExpireSeatHoldRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
