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

// Package package
//
// swagger:model Package
type Package struct {

	// best seat map
	BestSeatMap *BestSeatMapSummary `json:"BestSeatMap,omitempty"`

	// campaign
	Campaign *CampaignSummary `json:"Campaign,omitempty"`

	// code
	Code string `json:"Code,omitempty"`

	// control group
	ControlGroup *ControlGroupSummary `json:"ControlGroup,omitempty"`

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime *strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// edit indicator
	EditIndicator bool `json:"EditIndicator,omitempty"`

	// facility
	Facility *FacilitySummary `json:"Facility,omitempty"`

	// fixed seat indicator
	FixedSeatIndicator bool `json:"FixedSeatIndicator,omitempty"`

	// flex indicator
	FlexIndicator bool `json:"FlexIndicator,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// maximum performances
	MaximumPerformances int32 `json:"MaximumPerformances,omitempty"`

	// minimum performances
	MinimumPerformances int32 `json:"MinimumPerformances,omitempty"`

	// package date time
	// Format: date-time
	PackageDateTime *strfmt.DateTime `json:"PackageDateTime,omitempty"`

	// rank type
	RankType *RankTypeSummary `json:"RankType,omitempty"`

	// season
	Season *SeasonSummary `json:"Season,omitempty"`

	// sub package indicator
	SubPackageIndicator bool `json:"SubPackageIndicator,omitempty"`

	// sub package order Id
	SubPackageOrderID int32 `json:"SubPackageOrderId,omitempty"`

	// super package Id
	SuperPackageID int32 `json:"SuperPackageId,omitempty"`

	// super package indicator
	SuperPackageIndicator bool `json:"SuperPackageIndicator,omitempty"`

	// text1
	Text1 string `json:"Text1,omitempty"`

	// text2
	Text2 string `json:"Text2,omitempty"`

	// text3
	Text3 string `json:"Text3,omitempty"`

	// text4
	Text4 string `json:"Text4,omitempty"`

	// ticket design
	TicketDesign *TicketDesign `json:"TicketDesign,omitempty"`

	// type
	Type *PackageTypeSummary `json:"Type,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime *strfmt.DateTime `json:"UpdatedDateTime,omitempty"`

	// zone map
	ZoneMap *ZoneMapSummary `json:"ZoneMap,omitempty"`
}

// Validate validates this package
func (m *Package) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBestSeatMap(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCampaign(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateControlGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFacility(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackageDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRankType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTicketDesign(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZoneMap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Package) validateBestSeatMap(formats strfmt.Registry) error {
	if swag.IsZero(m.BestSeatMap) { // not required
		return nil
	}

	if m.BestSeatMap != nil {
		if err := m.BestSeatMap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BestSeatMap")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("BestSeatMap")
			}
			return err
		}
	}

	return nil
}

func (m *Package) validateCampaign(formats strfmt.Registry) error {
	if swag.IsZero(m.Campaign) { // not required
		return nil
	}

	if m.Campaign != nil {
		if err := m.Campaign.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Campaign")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Campaign")
			}
			return err
		}
	}

	return nil
}

func (m *Package) validateControlGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.ControlGroup) { // not required
		return nil
	}

	if m.ControlGroup != nil {
		if err := m.ControlGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ControlGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ControlGroup")
			}
			return err
		}
	}

	return nil
}

func (m *Package) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Package) validateFacility(formats strfmt.Registry) error {
	if swag.IsZero(m.Facility) { // not required
		return nil
	}

	if m.Facility != nil {
		if err := m.Facility.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Facility")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Facility")
			}
			return err
		}
	}

	return nil
}

func (m *Package) validatePackageDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.PackageDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("PackageDateTime", "body", "date-time", m.PackageDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Package) validateRankType(formats strfmt.Registry) error {
	if swag.IsZero(m.RankType) { // not required
		return nil
	}

	if m.RankType != nil {
		if err := m.RankType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RankType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("RankType")
			}
			return err
		}
	}

	return nil
}

func (m *Package) validateSeason(formats strfmt.Registry) error {
	if swag.IsZero(m.Season) { // not required
		return nil
	}

	if m.Season != nil {
		if err := m.Season.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Season")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Season")
			}
			return err
		}
	}

	return nil
}

func (m *Package) validateTicketDesign(formats strfmt.Registry) error {
	if swag.IsZero(m.TicketDesign) { // not required
		return nil
	}

	if m.TicketDesign != nil {
		if err := m.TicketDesign.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TicketDesign")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TicketDesign")
			}
			return err
		}
	}

	return nil
}

func (m *Package) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Type")
			}
			return err
		}
	}

	return nil
}

func (m *Package) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Package) validateZoneMap(formats strfmt.Registry) error {
	if swag.IsZero(m.ZoneMap) { // not required
		return nil
	}

	if m.ZoneMap != nil {
		if err := m.ZoneMap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ZoneMap")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ZoneMap")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this package based on the context it is used
func (m *Package) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBestSeatMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCampaign(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateControlGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFacility(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRankType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSeason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTicketDesign(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateZoneMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Package) contextValidateBestSeatMap(ctx context.Context, formats strfmt.Registry) error {

	if m.BestSeatMap != nil {

		if swag.IsZero(m.BestSeatMap) { // not required
			return nil
		}

		if err := m.BestSeatMap.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BestSeatMap")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("BestSeatMap")
			}
			return err
		}
	}

	return nil
}

func (m *Package) contextValidateCampaign(ctx context.Context, formats strfmt.Registry) error {

	if m.Campaign != nil {

		if swag.IsZero(m.Campaign) { // not required
			return nil
		}

		if err := m.Campaign.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Campaign")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Campaign")
			}
			return err
		}
	}

	return nil
}

func (m *Package) contextValidateControlGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.ControlGroup != nil {

		if swag.IsZero(m.ControlGroup) { // not required
			return nil
		}

		if err := m.ControlGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ControlGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ControlGroup")
			}
			return err
		}
	}

	return nil
}

func (m *Package) contextValidateFacility(ctx context.Context, formats strfmt.Registry) error {

	if m.Facility != nil {

		if swag.IsZero(m.Facility) { // not required
			return nil
		}

		if err := m.Facility.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Facility")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Facility")
			}
			return err
		}
	}

	return nil
}

func (m *Package) contextValidateRankType(ctx context.Context, formats strfmt.Registry) error {

	if m.RankType != nil {

		if swag.IsZero(m.RankType) { // not required
			return nil
		}

		if err := m.RankType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RankType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("RankType")
			}
			return err
		}
	}

	return nil
}

func (m *Package) contextValidateSeason(ctx context.Context, formats strfmt.Registry) error {

	if m.Season != nil {

		if swag.IsZero(m.Season) { // not required
			return nil
		}

		if err := m.Season.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Season")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Season")
			}
			return err
		}
	}

	return nil
}

func (m *Package) contextValidateTicketDesign(ctx context.Context, formats strfmt.Registry) error {

	if m.TicketDesign != nil {

		if swag.IsZero(m.TicketDesign) { // not required
			return nil
		}

		if err := m.TicketDesign.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TicketDesign")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TicketDesign")
			}
			return err
		}
	}

	return nil
}

func (m *Package) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {

		if swag.IsZero(m.Type) { // not required
			return nil
		}

		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Type")
			}
			return err
		}
	}

	return nil
}

func (m *Package) contextValidateZoneMap(ctx context.Context, formats strfmt.Registry) error {

	if m.ZoneMap != nil {

		if swag.IsZero(m.ZoneMap) { // not required
			return nil
		}

		if err := m.ZoneMap.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ZoneMap")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ZoneMap")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Package) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Package) UnmarshalBinary(b []byte) error {
	var res Package
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
