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

// OrdersForDeliveryRequest orders for delivery request
//
// swagger:model OrdersForDeliveryRequest
type OrdersForDeliveryRequest struct {

	// delivery methods
	DeliveryMethods string `json:"DeliveryMethods,omitempty"`

	// max rows to return
	MaxRowsToReturn int32 `json:"MaxRowsToReturn,omitempty"`

	// modes of sale
	ModesOfSale string `json:"ModesOfSale,omitempty"`

	// order days in past
	OrderDaysInPast int32 `json:"OrderDaysInPast,omitempty"`

	// order end date time
	// Format: date-time
	OrderEndDateTime strfmt.DateTime `json:"OrderEndDateTime,omitempty"`

	// order start date time
	// Format: date-time
	OrderStartDateTime strfmt.DateTime `json:"OrderStartDateTime,omitempty"`

	// organization name
	OrganizationName string `json:"OrganizationName,omitempty"`

	// performance end date time
	// Format: date-time
	PerformanceEndDateTime strfmt.DateTime `json:"PerformanceEndDateTime,omitempty"`

	// performance start date time
	// Format: date-time
	PerformanceStartDateTime strfmt.DateTime `json:"PerformanceStartDateTime,omitempty"`

	// use primary email
	UsePrimaryEmail bool `json:"UsePrimaryEmail,omitempty"`
}

// Validate validates this orders for delivery request
func (m *OrdersForDeliveryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderEndDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderStartDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerformanceEndDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerformanceStartDateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrdersForDeliveryRequest) validateOrderEndDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderEndDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("OrderEndDateTime", "body", "date-time", m.OrderEndDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OrdersForDeliveryRequest) validateOrderStartDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderStartDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("OrderStartDateTime", "body", "date-time", m.OrderStartDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OrdersForDeliveryRequest) validatePerformanceEndDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.PerformanceEndDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("PerformanceEndDateTime", "body", "date-time", m.PerformanceEndDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OrdersForDeliveryRequest) validatePerformanceStartDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.PerformanceStartDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("PerformanceStartDateTime", "body", "date-time", m.PerformanceStartDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this orders for delivery request based on context it is used
func (m *OrdersForDeliveryRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrdersForDeliveryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrdersForDeliveryRequest) UnmarshalBinary(b []byte) error {
	var res OrdersForDeliveryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}