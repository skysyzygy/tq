// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortfolioCustomDataItem portfolio custom data item
//
// swagger:model PortfolioCustomDataItem
type PortfolioCustomDataItem struct {

	// custom element Id
	CustomElementID int32 `json:"CustomElementId,omitempty"`

	// data type
	DataType string `json:"DataType,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// index
	Index int32 `json:"Index,omitempty"`

	// is dropdown
	IsDropdown bool `json:"IsDropdown,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// value
	Value string `json:"Value,omitempty"`
}

// Validate validates this portfolio custom data item
func (m *PortfolioCustomDataItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this portfolio custom data item based on context it is used
func (m *PortfolioCustomDataItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortfolioCustomDataItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortfolioCustomDataItem) UnmarshalBinary(b []byte) error {
	var res PortfolioCustomDataItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
