// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApplicationObject application object
//
// swagger:model ApplicationObject
type ApplicationObject struct {

	// constituency based
	ConstituencyBased string `json:"ConstituencyBased,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// object type
	ObjectType string `json:"ObjectType,omitempty"`
}

// Validate validates this application object
func (m *ApplicationObject) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this application object based on context it is used
func (m *ApplicationObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationObject) UnmarshalBinary(b []byte) error {
	var res ApplicationObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
