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

// Relationship relationship
//
// swagger:model Relationship
type Relationship struct {

	// associated constituent Id
	AssociatedConstituentID int32 `json:"AssociatedConstituentId,omitempty"`

	// associated constituent name
	AssociatedConstituentName string `json:"AssociatedConstituentName,omitempty"`

	// associated constituent sort name
	AssociatedConstituentSortName string `json:"AssociatedConstituentSortName,omitempty"`

	// birth date
	// Format: date-time
	BirthDate *strfmt.DateTime `json:"BirthDate,omitempty"`

	// constituent Id
	ConstituentID int32 `json:"ConstituentId,omitempty"`

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime *strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// end date
	// Format: date-time
	EndDate *strfmt.DateTime `json:"EndDate,omitempty"`

	// gender Id
	GenderID int32 `json:"GenderId,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// inactive
	Inactive bool `json:"Inactive,omitempty"`

	// is allowed to transact
	IsAllowedToTransact bool `json:"IsAllowedToTransact,omitempty"`

	// is editable
	IsEditable bool `json:"IsEditable,omitempty"`

	// is included in search results
	IsIncludedInSearchResults bool `json:"IsIncludedInSearchResults,omitempty"`

	// name indicator
	NameIndicator int32 `json:"NameIndicator,omitempty"`

	// note
	Note string `json:"Note,omitempty"`

	// primary indicator
	PrimaryIndicator bool `json:"PrimaryIndicator,omitempty"`

	// rank
	Rank int32 `json:"Rank,omitempty"`

	// reciprocal relationship Id
	ReciprocalRelationshipID int32 `json:"ReciprocalRelationshipId,omitempty"`

	// relationship category description
	RelationshipCategoryDescription string `json:"RelationshipCategoryDescription,omitempty"`

	// relationship category Id
	RelationshipCategoryID int32 `json:"RelationshipCategoryId,omitempty"`

	// relationship context
	RelationshipContext string `json:"RelationshipContext,omitempty"`

	// relationship type description
	RelationshipTypeDescription string `json:"RelationshipTypeDescription,omitempty"`

	// relationship type Id
	RelationshipTypeID int32 `json:"RelationshipTypeId,omitempty"`

	// salary
	Salary float64 `json:"Salary,omitempty"`

	// start date
	// Format: date-time
	StartDate *strfmt.DateTime `json:"StartDate,omitempty"`

	// title
	Title string `json:"Title,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime *strfmt.DateTime `json:"UpdatedDateTime,omitempty"`
}

// Validate validates this relationship
func (m *Relationship) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBirthDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
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

func (m *Relationship) validateBirthDate(formats strfmt.Registry) error {
	if swag.IsZero(m.BirthDate) { // not required
		return nil
	}

	if err := validate.FormatOf("BirthDate", "body", "date-time", m.BirthDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Relationship) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Relationship) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("EndDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Relationship) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("StartDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Relationship) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this relationship based on context it is used
func (m *Relationship) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Relationship) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Relationship) UnmarshalBinary(b []byte) error {
	var res Relationship
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
