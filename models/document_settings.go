// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DocumentSettings document settings
//
// swagger:model DocumentSettings
type DocumentSettings struct {

	// max allowed content length
	MaxAllowedContentLength int32 `json:"MaxAllowedContentLength,omitempty"`

	// max array length
	MaxArrayLength int32 `json:"MaxArrayLength,omitempty"`

	// max received message size
	MaxReceivedMessageSize int32 `json:"MaxReceivedMessageSize,omitempty"`

	// max request length
	MaxRequestLength int32 `json:"MaxRequestLength,omitempty"`

	// max string content length
	MaxStringContentLength int32 `json:"MaxStringContentLength,omitempty"`
}

// Validate validates this document settings
func (m *DocumentSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this document settings based on context it is used
func (m *DocumentSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DocumentSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DocumentSettings) UnmarshalBinary(b []byte) error {
	var res DocumentSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
