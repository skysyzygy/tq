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

// HTMLTicket Html ticket
//
// swagger:model HtmlTicket
type HTMLTicket struct {

	// Html
	HTML string `json:"Html,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// seat
	Seat *CartSeat `json:"Seat,omitempty"`
}

// Validate validates this Html ticket
func (m *HTMLTicket) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSeat(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HTMLTicket) validateSeat(formats strfmt.Registry) error {
	if swag.IsZero(m.Seat) { // not required
		return nil
	}

	if m.Seat != nil {
		if err := m.Seat.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Seat")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Seat")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Html ticket based on the context it is used
func (m *HTMLTicket) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSeat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HTMLTicket) contextValidateSeat(ctx context.Context, formats strfmt.Registry) error {

	if m.Seat != nil {

		if swag.IsZero(m.Seat) { // not required
			return nil
		}

		if err := m.Seat.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Seat")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Seat")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HTMLTicket) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HTMLTicket) UnmarshalBinary(b []byte) error {
	var res HTMLTicket
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
