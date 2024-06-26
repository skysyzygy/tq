// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VantivTran vantiv tran
//
// swagger:model VantivTran
type VantivTran struct {

	// tran type
	TranType string `json:"TranType,omitempty"`
}

// Validate validates this vantiv tran
func (m *VantivTran) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this vantiv tran based on context it is used
func (m *VantivTran) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VantivTran) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VantivTran) UnmarshalBinary(b []byte) error {
	var res VantivTran
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
