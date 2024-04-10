// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OutputResultRequest output result request
//
// swagger:model OutputResultRequest
type OutputResultRequest struct {

	// address purpose Id
	AddressPurposeID int32 `json:"AddressPurposeId,omitempty"`

	// disable list generate
	DisableListGenerate bool `json:"DisableListGenerate,omitempty"`

	// e marketing indicator
	EMarketingIndicator bool `json:"EMarketingIndicator,omitempty"`

	// electronic address purpose Id
	ElectronicAddressPurposeID int32 `json:"ElectronicAddressPurposeId,omitempty"`

	// electronic address type Id
	ElectronicAddressTypeID int32 `json:"ElectronicAddressTypeId,omitempty"`

	// format date
	FormatDate bool `json:"FormatDate,omitempty"`

	// mailing date time
	// Format: date-time
	MailingDateTime strfmt.DateTime `json:"MailingDateTime,omitempty"`

	// membership organization Id
	MembershipOrganizationID int32 `json:"MembershipOrganizationId,omitempty"`

	// output set Id
	OutputSetID int32 `json:"OutputSetId,omitempty"`

	// page
	Page int32 `json:"Page,omitempty"`

	// page size
	PageSize int32 `json:"PageSize,omitempty"`

	// salutation type Id
	SalutationTypeID int32 `json:"SalutationTypeId,omitempty"`

	// search text
	SearchText string `json:"SearchText,omitempty"`

	// sort by
	SortBy string `json:"SortBy,omitempty"`

	// use label address
	UseLabelAddress bool `json:"UseLabelAddress,omitempty"`
}

// Validate validates this output result request
func (m *OutputResultRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMailingDateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OutputResultRequest) validateMailingDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.MailingDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("MailingDateTime", "body", "date-time", m.MailingDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this output result request based on context it is used
func (m *OutputResultRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OutputResultRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OutputResultRequest) UnmarshalBinary(b []byte) error {
	var res OutputResultRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
