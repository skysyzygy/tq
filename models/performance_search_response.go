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

// PerformanceSearchResponse performance search response
//
// swagger:model PerformanceSearchResponse
type PerformanceSearchResponse struct {

	// avail sale indicator
	AvailSaleIndicator string `json:"AvailSaleIndicator,omitempty"`

	// business unit Id
	BusinessUnitID int32 `json:"BusinessUnitId,omitempty"`

	// constituent interest weight
	ConstituentInterestWeight int32 `json:"ConstituentInterestWeight,omitempty"`

	// doors close
	// Format: date-time
	DoorsClose strfmt.DateTime `json:"DoorsClose,omitempty"`

	// doors open
	// Format: date-time
	DoorsOpen strfmt.DateTime `json:"DoorsOpen,omitempty"`

	// duration
	Duration int32 `json:"Duration,omitempty"`

	// facility
	Facility *EntitySummary `json:"Facility,omitempty"`

	// facility description
	FacilityDescription string `json:"FacilityDescription,omitempty"`

	// is restricted
	IsRestricted bool `json:"IsRestricted,omitempty"`

	// mode of sale
	ModeOfSale *EntitySummary `json:"ModeOfSale,omitempty"`

	// mode of sale description
	ModeOfSaleDescription string `json:"ModeOfSaleDescription,omitempty"`

	// mode of sale end date
	// Format: date-time
	ModeOfSaleEndDate strfmt.DateTime `json:"ModeOfSaleEndDate,omitempty"`

	// mode of sale start date
	// Format: date-time
	ModeOfSaleStartDate strfmt.DateTime `json:"ModeOfSaleStartDate,omitempty"`

	// performance code
	PerformanceCode string `json:"PerformanceCode,omitempty"`

	// performance date
	// Format: date-time
	PerformanceDate strfmt.DateTime `json:"PerformanceDate,omitempty"`

	// performance description
	PerformanceDescription string `json:"PerformanceDescription,omitempty"`

	// performance Id
	PerformanceID int32 `json:"PerformanceId,omitempty"`

	// performance short name
	PerformanceShortName string `json:"PerformanceShortName,omitempty"`

	// performance status
	PerformanceStatus *EntitySummary `json:"PerformanceStatus,omitempty"`

	// performance status description
	PerformanceStatusDescription string `json:"PerformanceStatusDescription,omitempty"`

	// performance type
	PerformanceType *PerformanceTypeSummary `json:"PerformanceType,omitempty"`

	// production season
	ProductionSeason *EntitySummary `json:"ProductionSeason,omitempty"`

	// publish client end date
	// Format: date-time
	PublishClientEndDate strfmt.DateTime `json:"PublishClientEndDate,omitempty"`

	// publish client start date
	// Format: date-time
	PublishClientStartDate strfmt.DateTime `json:"PublishClientStartDate,omitempty"`

	// publish web Api end date
	// Format: date-time
	PublishWebAPIEndDate strfmt.DateTime `json:"PublishWebApiEndDate,omitempty"`

	// publish web Api start date
	// Format: date-time
	PublishWebAPIStartDate strfmt.DateTime `json:"PublishWebApiStartDate,omitempty"`

	// rank type
	RankType *EntitySummary `json:"RankType,omitempty"`

	// rank type description
	RankTypeDescription string `json:"RankTypeDescription,omitempty"`

	// sales notes
	SalesNotes string `json:"SalesNotes,omitempty"`

	// sales notes required
	SalesNotesRequired bool `json:"SalesNotesRequired,omitempty"`

	// season
	Season *SeasonSummary `json:"Season,omitempty"`

	// text1
	Text1 string `json:"Text1,omitempty"`

	// text2
	Text2 string `json:"Text2,omitempty"`

	// text3
	Text3 string `json:"Text3,omitempty"`

	// text4
	Text4 string `json:"Text4,omitempty"`

	// time slot
	TimeSlot *EntitySummary `json:"TimeSlot,omitempty"`

	// time slot description
	TimeSlotDescription string `json:"TimeSlotDescription,omitempty"`

	// valid countries
	ValidCountries string `json:"ValidCountries,omitempty"`

	// zone map
	ZoneMap *EntitySummary `json:"ZoneMap,omitempty"`

	// zone map description
	ZoneMapDescription string `json:"ZoneMapDescription,omitempty"`
}

// Validate validates this performance search response
func (m *PerformanceSearchResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDoorsClose(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDoorsOpen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFacility(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModeOfSale(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModeOfSaleEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModeOfSaleStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerformanceDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerformanceStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerformanceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductionSeason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishClientEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishClientStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishWebAPIEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishWebAPIStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRankType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeSlot(formats); err != nil {
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

func (m *PerformanceSearchResponse) validateDoorsClose(formats strfmt.Registry) error {
	if swag.IsZero(m.DoorsClose) { // not required
		return nil
	}

	if err := validate.FormatOf("DoorsClose", "body", "date-time", m.DoorsClose.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceSearchResponse) validateDoorsOpen(formats strfmt.Registry) error {
	if swag.IsZero(m.DoorsOpen) { // not required
		return nil
	}

	if err := validate.FormatOf("DoorsOpen", "body", "date-time", m.DoorsOpen.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceSearchResponse) validateFacility(formats strfmt.Registry) error {
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

func (m *PerformanceSearchResponse) validateModeOfSale(formats strfmt.Registry) error {
	if swag.IsZero(m.ModeOfSale) { // not required
		return nil
	}

	if m.ModeOfSale != nil {
		if err := m.ModeOfSale.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ModeOfSale")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ModeOfSale")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSearchResponse) validateModeOfSaleEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ModeOfSaleEndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("ModeOfSaleEndDate", "body", "date-time", m.ModeOfSaleEndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceSearchResponse) validateModeOfSaleStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ModeOfSaleStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("ModeOfSaleStartDate", "body", "date-time", m.ModeOfSaleStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceSearchResponse) validatePerformanceDate(formats strfmt.Registry) error {
	if swag.IsZero(m.PerformanceDate) { // not required
		return nil
	}

	if err := validate.FormatOf("PerformanceDate", "body", "date-time", m.PerformanceDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceSearchResponse) validatePerformanceStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.PerformanceStatus) { // not required
		return nil
	}

	if m.PerformanceStatus != nil {
		if err := m.PerformanceStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PerformanceStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PerformanceStatus")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSearchResponse) validatePerformanceType(formats strfmt.Registry) error {
	if swag.IsZero(m.PerformanceType) { // not required
		return nil
	}

	if m.PerformanceType != nil {
		if err := m.PerformanceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PerformanceType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PerformanceType")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSearchResponse) validateProductionSeason(formats strfmt.Registry) error {
	if swag.IsZero(m.ProductionSeason) { // not required
		return nil
	}

	if m.ProductionSeason != nil {
		if err := m.ProductionSeason.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ProductionSeason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ProductionSeason")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSearchResponse) validatePublishClientEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.PublishClientEndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("PublishClientEndDate", "body", "date-time", m.PublishClientEndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceSearchResponse) validatePublishClientStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.PublishClientStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("PublishClientStartDate", "body", "date-time", m.PublishClientStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceSearchResponse) validatePublishWebAPIEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.PublishWebAPIEndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("PublishWebApiEndDate", "body", "date-time", m.PublishWebAPIEndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceSearchResponse) validatePublishWebAPIStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.PublishWebAPIStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("PublishWebApiStartDate", "body", "date-time", m.PublishWebAPIStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceSearchResponse) validateRankType(formats strfmt.Registry) error {
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

func (m *PerformanceSearchResponse) validateSeason(formats strfmt.Registry) error {
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

func (m *PerformanceSearchResponse) validateTimeSlot(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeSlot) { // not required
		return nil
	}

	if m.TimeSlot != nil {
		if err := m.TimeSlot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TimeSlot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TimeSlot")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSearchResponse) validateZoneMap(formats strfmt.Registry) error {
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

// ContextValidate validate this performance search response based on the context it is used
func (m *PerformanceSearchResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFacility(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModeOfSale(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePerformanceStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePerformanceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProductionSeason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRankType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSeason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimeSlot(ctx, formats); err != nil {
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

func (m *PerformanceSearchResponse) contextValidateFacility(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceSearchResponse) contextValidateModeOfSale(ctx context.Context, formats strfmt.Registry) error {

	if m.ModeOfSale != nil {

		if swag.IsZero(m.ModeOfSale) { // not required
			return nil
		}

		if err := m.ModeOfSale.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ModeOfSale")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ModeOfSale")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSearchResponse) contextValidatePerformanceStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.PerformanceStatus != nil {

		if swag.IsZero(m.PerformanceStatus) { // not required
			return nil
		}

		if err := m.PerformanceStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PerformanceStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PerformanceStatus")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSearchResponse) contextValidatePerformanceType(ctx context.Context, formats strfmt.Registry) error {

	if m.PerformanceType != nil {

		if swag.IsZero(m.PerformanceType) { // not required
			return nil
		}

		if err := m.PerformanceType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PerformanceType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PerformanceType")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSearchResponse) contextValidateProductionSeason(ctx context.Context, formats strfmt.Registry) error {

	if m.ProductionSeason != nil {

		if swag.IsZero(m.ProductionSeason) { // not required
			return nil
		}

		if err := m.ProductionSeason.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ProductionSeason")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ProductionSeason")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSearchResponse) contextValidateRankType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceSearchResponse) contextValidateSeason(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceSearchResponse) contextValidateTimeSlot(ctx context.Context, formats strfmt.Registry) error {

	if m.TimeSlot != nil {

		if swag.IsZero(m.TimeSlot) { // not required
			return nil
		}

		if err := m.TimeSlot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TimeSlot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TimeSlot")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceSearchResponse) contextValidateZoneMap(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PerformanceSearchResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceSearchResponse) UnmarshalBinary(b []byte) error {
	var res PerformanceSearchResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
