// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WebExpiringResponse web expiring response
//
// swagger:model WebExpiringResponse
type WebExpiringResponse struct {

	// message
	Message string `json:"Message,omitempty"`

	// sources
	Sources []*SourceSummary `json:"Sources"`
}

// Validate validates this web expiring response
func (m *WebExpiringResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebExpiringResponse) validateSources(formats strfmt.Registry) error {
	if swag.IsZero(m.Sources) { // not required
		return nil
	}

	for i := 0; i < len(m.Sources); i++ {
		if swag.IsZero(m.Sources[i]) { // not required
			continue
		}

		if m.Sources[i] != nil {
			if err := m.Sources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Sources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Sources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this web expiring response based on the context it is used
func (m *WebExpiringResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebExpiringResponse) contextValidateSources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Sources); i++ {

		if m.Sources[i] != nil {

			if swag.IsZero(m.Sources[i]) { // not required
				return nil
			}

			if err := m.Sources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Sources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Sources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebExpiringResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebExpiringResponse) UnmarshalBinary(b []byte) error {
	var res WebExpiringResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
