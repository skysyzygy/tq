// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddUpdateFeeResponse add update fee response
//
// swagger:model AddUpdateFeeResponse
type AddUpdateFeeResponse struct {

	// item fee Id
	ItemFeeID int32 `json:"ItemFeeId,omitempty"`
}

// Validate validates this add update fee response
func (m *AddUpdateFeeResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add update fee response based on context it is used
func (m *AddUpdateFeeResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AddUpdateFeeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddUpdateFeeResponse) UnmarshalBinary(b []byte) error {
	var res AddUpdateFeeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
