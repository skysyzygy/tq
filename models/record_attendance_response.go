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

// RecordAttendanceResponse record attendance response
//
// swagger:model RecordAttendanceResponse
type RecordAttendanceResponse struct {

	// attendance status
	AttendanceStatus *AttendanceStatus `json:"AttendanceStatus,omitempty"`

	// nscan control status
	NscanControlStatus *NscanControlStatus `json:"NscanControlStatus,omitempty"`

	// sub line item record status
	SubLineItemRecordStatus *SubLineItemRecordStatus `json:"SubLineItemRecordStatus,omitempty"`
}

// Validate validates this record attendance response
func (m *RecordAttendanceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttendanceStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNscanControlStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubLineItemRecordStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecordAttendanceResponse) validateAttendanceStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.AttendanceStatus) { // not required
		return nil
	}

	if m.AttendanceStatus != nil {
		if err := m.AttendanceStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AttendanceStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("AttendanceStatus")
			}
			return err
		}
	}

	return nil
}

func (m *RecordAttendanceResponse) validateNscanControlStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.NscanControlStatus) { // not required
		return nil
	}

	if m.NscanControlStatus != nil {
		if err := m.NscanControlStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NscanControlStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("NscanControlStatus")
			}
			return err
		}
	}

	return nil
}

func (m *RecordAttendanceResponse) validateSubLineItemRecordStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.SubLineItemRecordStatus) { // not required
		return nil
	}

	if m.SubLineItemRecordStatus != nil {
		if err := m.SubLineItemRecordStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SubLineItemRecordStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("SubLineItemRecordStatus")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this record attendance response based on the context it is used
func (m *RecordAttendanceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttendanceStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNscanControlStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubLineItemRecordStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecordAttendanceResponse) contextValidateAttendanceStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.AttendanceStatus != nil {

		if swag.IsZero(m.AttendanceStatus) { // not required
			return nil
		}

		if err := m.AttendanceStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AttendanceStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("AttendanceStatus")
			}
			return err
		}
	}

	return nil
}

func (m *RecordAttendanceResponse) contextValidateNscanControlStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.NscanControlStatus != nil {

		if swag.IsZero(m.NscanControlStatus) { // not required
			return nil
		}

		if err := m.NscanControlStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NscanControlStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("NscanControlStatus")
			}
			return err
		}
	}

	return nil
}

func (m *RecordAttendanceResponse) contextValidateSubLineItemRecordStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.SubLineItemRecordStatus != nil {

		if swag.IsZero(m.SubLineItemRecordStatus) { // not required
			return nil
		}

		if err := m.SubLineItemRecordStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SubLineItemRecordStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("SubLineItemRecordStatus")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecordAttendanceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecordAttendanceResponse) UnmarshalBinary(b []byte) error {
	var res RecordAttendanceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
