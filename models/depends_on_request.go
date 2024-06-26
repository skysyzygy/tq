// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DependsOnRequest depends on request
//
// swagger:model DependsOnRequest
type DependsOnRequest struct {

	// Id
	ID int32 `json:"Id,omitempty"`

	// is Url dependency
	IsURLDependency bool `json:"IsUrlDependency,omitempty"`

	// source path
	SourcePath string `json:"SourcePath,omitempty"`

	// target path
	TargetPath string `json:"TargetPath,omitempty"`
}

// Validate validates this depends on request
func (m *DependsOnRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this depends on request based on context it is used
func (m *DependsOnRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DependsOnRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DependsOnRequest) UnmarshalBinary(b []byte) error {
	var res DependsOnRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
