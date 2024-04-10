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

// ReportParameter report parameter
//
// swagger:model ReportParameter
type ReportParameter struct {

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// data column
	DataColumn string `json:"DataColumn,omitempty"`

	// data type
	DataType int32 `json:"DataType,omitempty"`

	// default value
	DefaultValue string `json:"DefaultValue,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// disable clause
	DisableClause string `json:"DisableClause,omitempty"`

	// display column
	DisplayColumn string `json:"DisplayColumn,omitempty"`

	// edit style
	EditStyle string `json:"EditStyle,omitempty"`

	// edit style flag
	EditStyleFlag string `json:"EditStyleFlag,omitempty"`

	// end of day
	EndOfDay bool `json:"EndOfDay,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// is dependent on other parameters
	IsDependentOnOtherParameters bool `json:"IsDependentOnOtherParameters,omitempty"`

	// length
	Length int32 `json:"Length,omitempty"`

	// multi select
	MultiSelect bool `json:"MultiSelect,omitempty"`

	// order clause
	OrderClause string `json:"OrderClause,omitempty"`

	// parameter name
	ParameterName string `json:"ParameterName,omitempty"`

	// parameter type
	ParameterType int32 `json:"ParameterType,omitempty"`

	// report Id
	ReportID string `json:"ReportId,omitempty"`

	// required
	Required bool `json:"Required,omitempty"`

	// sequence number
	SequenceNumber int32 `json:"SequenceNumber,omitempty"`

	// table name
	TableName string `json:"TableName,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime strfmt.DateTime `json:"UpdatedDateTime,omitempty"`

	// validation
	Validation string `json:"Validation,omitempty"`

	// where clause
	WhereClause string `json:"WhereClause,omitempty"`
}

// Validate validates this report parameter
func (m *ReportParameter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportParameter) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportParameter) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this report parameter based on context it is used
func (m *ReportParameter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReportParameter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportParameter) UnmarshalBinary(b []byte) error {
	var res ReportParameter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
