// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PriceSummary price summary
//
// swagger:model PriceSummary
type PriceSummary struct {

	// editable min price
	EditableMinPrice float64 `json:"EditableMinPrice,omitempty"`

	// enabled
	Enabled bool `json:"Enabled,omitempty"`

	// is base
	IsBase bool `json:"IsBase,omitempty"`

	// is best
	IsBest bool `json:"IsBest,omitempty"`

	// is editable
	IsEditable bool `json:"IsEditable,omitempty"`

	// is editable for web
	IsEditableForWeb bool `json:"IsEditableForWeb,omitempty"`

	// layer type Id
	LayerTypeID int32 `json:"LayerTypeId,omitempty"`

	// min price
	MinPrice float64 `json:"MinPrice,omitempty"`

	// mode of sale offer Id
	ModeOfSaleOfferID int32 `json:"ModeOfSaleOfferId,omitempty"`

	// offer
	Offer bool `json:"Offer,omitempty"`

	// package Id
	PackageID int32 `json:"PackageId,omitempty"`

	// parent package Id
	ParentPackageID int32 `json:"ParentPackageId,omitempty"`

	// performance Id
	PerformanceID int32 `json:"PerformanceId,omitempty"`

	// performance price type Id
	PerformancePriceTypeID int32 `json:"PerformancePriceTypeId,omitempty"`

	// price
	Price float64 `json:"Price,omitempty"`

	// price type Id
	PriceTypeID int32 `json:"PriceTypeId,omitempty"`

	// zone Id
	ZoneID int32 `json:"ZoneId,omitempty"`
}

// Validate validates this price summary
func (m *PriceSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this price summary based on context it is used
func (m *PriceSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PriceSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PriceSummary) UnmarshalBinary(b []byte) error {
	var res PriceSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
