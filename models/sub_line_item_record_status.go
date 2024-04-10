// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SubLineItemRecordStatus sub line item record status
//
// swagger:model SubLineItemRecordStatus
type SubLineItemRecordStatus struct {

	// failure reason
	FailureReason string `json:"FailureReason,omitempty"`

	// sub line item Id
	SubLineItemID int32 `json:"SubLineItemId,omitempty"`

	// ticket no
	TicketNo int32 `json:"TicketNo,omitempty"`

	// updated
	Updated bool `json:"Updated,omitempty"`
}

// Validate validates this sub line item record status
func (m *SubLineItemRecordStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sub line item record status based on context it is used
func (m *SubLineItemRecordStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SubLineItemRecordStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubLineItemRecordStatus) UnmarshalBinary(b []byte) error {
	var res SubLineItemRecordStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
