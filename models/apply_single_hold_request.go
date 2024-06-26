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

// ApplySingleHoldRequest apply single hold request
//
// swagger:model ApplySingleHoldRequest
type ApplySingleHoldRequest struct {

	// hold code Id
	HoldCodeID int32 `json:"HoldCodeId,omitempty"`

	// hold until date
	// Format: date-time
	HoldUntilDate *strfmt.DateTime `json:"HoldUntilDate,omitempty"`

	// replace mode
	ReplaceMode int32 `json:"ReplaceMode,omitempty"`
}

// Validate validates this apply single hold request
func (m *ApplySingleHoldRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHoldUntilDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplySingleHoldRequest) validateHoldUntilDate(formats strfmt.Registry) error {
	if swag.IsZero(m.HoldUntilDate) { // not required
		return nil
	}

	if err := validate.FormatOf("HoldUntilDate", "body", "date-time", m.HoldUntilDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this apply single hold request based on context it is used
func (m *ApplySingleHoldRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApplySingleHoldRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplySingleHoldRequest) UnmarshalBinary(b []byte) error {
	var res ApplySingleHoldRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
