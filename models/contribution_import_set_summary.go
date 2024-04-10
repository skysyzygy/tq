// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContributionImportSetSummary contribution import set summary
//
// swagger:model ContributionImportSetSummary
type ContributionImportSetSummary struct {

	// description
	Description string `json:"Description,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`
}

// Validate validates this contribution import set summary
func (m *ContributionImportSetSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this contribution import set summary based on context it is used
func (m *ContributionImportSetSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContributionImportSetSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContributionImportSetSummary) UnmarshalBinary(b []byte) error {
	var res ContributionImportSetSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
