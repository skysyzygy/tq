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

// CoreIdentitySignRequest core identity sign request
//
// swagger:model CoreIdentitySignRequest
type CoreIdentitySignRequest struct {

	// body
	Body string `json:"Body,omitempty"`

	// headers
	Headers []*NameValue `json:"Headers"`

	// method
	Method string `json:"Method,omitempty"`

	// path
	Path string `json:"Path,omitempty"`

	// query parameters
	QueryParameters []*NameValue `json:"QueryParameters"`
}

// Validate validates this core identity sign request
func (m *CoreIdentitySignRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeaders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CoreIdentitySignRequest) validateHeaders(formats strfmt.Registry) error {
	if swag.IsZero(m.Headers) { // not required
		return nil
	}

	for i := 0; i < len(m.Headers); i++ {
		if swag.IsZero(m.Headers[i]) { // not required
			continue
		}

		if m.Headers[i] != nil {
			if err := m.Headers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Headers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Headers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CoreIdentitySignRequest) validateQueryParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.QueryParameters) { // not required
		return nil
	}

	for i := 0; i < len(m.QueryParameters); i++ {
		if swag.IsZero(m.QueryParameters[i]) { // not required
			continue
		}

		if m.QueryParameters[i] != nil {
			if err := m.QueryParameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("QueryParameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("QueryParameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this core identity sign request based on the context it is used
func (m *CoreIdentitySignRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHeaders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQueryParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CoreIdentitySignRequest) contextValidateHeaders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Headers); i++ {

		if m.Headers[i] != nil {

			if swag.IsZero(m.Headers[i]) { // not required
				return nil
			}

			if err := m.Headers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Headers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Headers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CoreIdentitySignRequest) contextValidateQueryParameters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.QueryParameters); i++ {

		if m.QueryParameters[i] != nil {

			if swag.IsZero(m.QueryParameters[i]) { // not required
				return nil
			}

			if err := m.QueryParameters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("QueryParameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("QueryParameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CoreIdentitySignRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CoreIdentitySignRequest) UnmarshalBinary(b []byte) error {
	var res CoreIdentitySignRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
