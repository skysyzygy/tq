// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EmailEmbeddedImage email embedded image
//
// swagger:model EmailEmbeddedImage
type EmailEmbeddedImage struct {

	// image bytes
	// Format: byte
	ImageBytes strfmt.Base64 `json:"ImageBytes,omitempty"`

	// media type
	MediaType string `json:"MediaType,omitempty"`

	// name
	Name string `json:"Name,omitempty"`
}

// Validate validates this email embedded image
func (m *EmailEmbeddedImage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this email embedded image based on context it is used
func (m *EmailEmbeddedImage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmailEmbeddedImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmailEmbeddedImage) UnmarshalBinary(b []byte) error {
	var res EmailEmbeddedImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
