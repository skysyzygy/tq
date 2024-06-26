// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustomDataItem custom data item
//
// swagger:model CustomDataItem
type CustomDataItem struct {

	// data type
	DataType string `json:"DataType,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// edit indicator
	EditIndicator bool `json:"EditIndicator,omitempty"`

	// index
	Index int32 `json:"Index,omitempty"`

	// is dropdown
	IsDropdown bool `json:"IsDropdown,omitempty"`

	// keyword Id
	KeywordID int32 `json:"KeywordId,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// value
	Value string `json:"Value,omitempty"`
}

// Validate validates this custom data item
func (m *CustomDataItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this custom data item based on context it is used
func (m *CustomDataItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomDataItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomDataItem) UnmarshalBinary(b []byte) error {
	var res CustomDataItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
