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

// BulkCopyDayForCopyPackage bulk copy day for copy package
//
// swagger:model BulkCopyDayForCopyPackage
type BulkCopyDayForCopyPackage struct {

	// code
	Code string `json:"Code,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// performances
	Performances []*BulkCopyDayForCopyPerformance `json:"Performances"`
}

// Validate validates this bulk copy day for copy package
func (m *BulkCopyDayForCopyPackage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePerformances(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkCopyDayForCopyPackage) validatePerformances(formats strfmt.Registry) error {
	if swag.IsZero(m.Performances) { // not required
		return nil
	}

	for i := 0; i < len(m.Performances); i++ {
		if swag.IsZero(m.Performances[i]) { // not required
			continue
		}

		if m.Performances[i] != nil {
			if err := m.Performances[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Performances" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Performances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this bulk copy day for copy package based on the context it is used
func (m *BulkCopyDayForCopyPackage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePerformances(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkCopyDayForCopyPackage) contextValidatePerformances(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Performances); i++ {

		if m.Performances[i] != nil {

			if swag.IsZero(m.Performances[i]) { // not required
				return nil
			}

			if err := m.Performances[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Performances" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Performances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BulkCopyDayForCopyPackage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BulkCopyDayForCopyPackage) UnmarshalBinary(b []byte) error {
	var res BulkCopyDayForCopyPackage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
