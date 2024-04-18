// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PrintOrderTicketsRequest print order tickets request
//
// swagger:model PrintOrderTicketsRequest
type PrintOrderTicketsRequest struct {

	// batch Id
	BatchID int32 `json:"BatchId,omitempty"`

	// header design Id
	HeaderDesignID int32 `json:"HeaderDesignId,omitempty"`

	// include receipts
	IncludeReceipts bool `json:"IncludeReceipts,omitempty"`

	// line items
	LineItems string `json:"LineItems,omitempty"`

	// new ticket no for reprints
	NewTicketNoForReprints bool `json:"NewTicketNoForReprints,omitempty"`

	// printer type
	PrinterType string `json:"PrinterType,omitempty"`

	// reprint tickets
	ReprintTickets bool `json:"ReprintTickets,omitempty"`

	// sub line items
	SubLineItems string `json:"SubLineItems,omitempty"`

	// ticket design Id
	TicketDesignID int32 `json:"TicketDesignId,omitempty"`
}

// Validate validates this print order tickets request
func (m *PrintOrderTicketsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this print order tickets request based on context it is used
func (m *PrintOrderTicketsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PrintOrderTicketsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrintOrderTicketsRequest) UnmarshalBinary(b []byte) error {
	var res PrintOrderTicketsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}