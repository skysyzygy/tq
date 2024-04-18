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

// EventControl event control
//
// swagger:model EventControl
type EventControl struct {

	// composite number
	CompositeNumber int32 `json:"CompositeNumber,omitempty"`

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// customer membership number
	CustomerMembershipNumber int32 `json:"CustomerMembershipNumber,omitempty"`

	// customer number
	CustomerNumber int32 `json:"CustomerNumber,omitempty"`

	// device name
	DeviceName string `json:"DeviceName,omitempty"`

	// entrance
	Entrance string `json:"Entrance,omitempty"`

	// error text
	ErrorText string `json:"ErrorText,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// membership level number
	MembershipLevelNumber int32 `json:"MembershipLevelNumber,omitempty"`

	// package number
	PackageNumber int32 `json:"PackageNumber,omitempty"`

	// performance code
	PerformanceCode string `json:"PerformanceCode,omitempty"`

	// performance date time
	// Format: date-time
	PerformanceDateTime strfmt.DateTime `json:"PerformanceDateTime,omitempty"`

	// performance number
	PerformanceNumber int32 `json:"PerformanceNumber,omitempty"`

	// scan string
	ScanString string `json:"ScanString,omitempty"`

	// seat number
	SeatNumber int32 `json:"SeatNumber,omitempty"`

	// status
	Status string `json:"Status,omitempty"`

	// sub line item number
	SubLineItemNumber int32 `json:"SubLineItemNumber,omitempty"`

	// ticket message
	TicketMessage string `json:"TicketMessage,omitempty"`

	// ticket number
	TicketNumber int32 `json:"TicketNumber,omitempty"`

	// ticket ok
	TicketOk string `json:"TicketOk,omitempty"`

	// ticket updated
	TicketUpdated string `json:"TicketUpdated,omitempty"`

	// update date time
	// Format: date-time
	UpdateDateTime strfmt.DateTime `json:"UpdateDateTime,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime strfmt.DateTime `json:"UpdatedDateTime,omitempty"`

	// user Id
	UserID string `json:"UserId,omitempty"`
}

// Validate validates this event control
func (m *EventControl) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerformanceDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventControl) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EventControl) validatePerformanceDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.PerformanceDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("PerformanceDateTime", "body", "date-time", m.PerformanceDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EventControl) validateUpdateDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdateDateTime", "body", "date-time", m.UpdateDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EventControl) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this event control based on context it is used
func (m *EventControl) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventControl) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventControl) UnmarshalBinary(b []byte) error {
	var res EventControl
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}