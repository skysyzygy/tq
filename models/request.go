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

// Request request
//
// swagger:model Request
type Request struct {

	// content
	Content interface{} `json:"Content,omitempty"`

	// continue on error
	ContinueOnError bool `json:"ContinueOnError,omitempty"`

	// depends on requests
	DependsOnRequests []*DependsOnRequest `json:"DependsOnRequests"`

	// Http method
	HTTPMethod string `json:"HttpMethod,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// return type string
	ReturnTypeString string `json:"ReturnTypeString,omitempty"`

	// Uri
	URI string `json:"Uri,omitempty"`
}

// Validate validates this request
func (m *Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDependsOnRequests(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Request) validateDependsOnRequests(formats strfmt.Registry) error {
	if swag.IsZero(m.DependsOnRequests) { // not required
		return nil
	}

	for i := 0; i < len(m.DependsOnRequests); i++ {
		if swag.IsZero(m.DependsOnRequests[i]) { // not required
			continue
		}

		if m.DependsOnRequests[i] != nil {
			if err := m.DependsOnRequests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DependsOnRequests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DependsOnRequests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this request based on the context it is used
func (m *Request) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDependsOnRequests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Request) contextValidateDependsOnRequests(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DependsOnRequests); i++ {

		if m.DependsOnRequests[i] != nil {

			if swag.IsZero(m.DependsOnRequests[i]) { // not required
				return nil
			}

			if err := m.DependsOnRequests[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DependsOnRequests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DependsOnRequests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Request) UnmarshalBinary(b []byte) error {
	var res Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
