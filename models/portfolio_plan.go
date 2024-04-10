// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PortfolioPlan portfolio plan
//
// swagger:model PortfolioPlan
type PortfolioPlan struct {

	// active workers
	ActiveWorkers string `json:"ActiveWorkers,omitempty"`

	// ask amount
	AskAmount float64 `json:"AskAmount,omitempty"`

	// campaign
	Campaign *CampaignSummary `json:"Campaign,omitempty"`

	// complete by date time
	// Format: date-time
	CompleteByDateTime strfmt.DateTime `json:"CompleteByDateTime,omitempty"`

	// constituent
	Constituent *ConstituentDisplaySummary `json:"Constituent,omitempty"`

	// contribution amount
	ContributionAmount float64 `json:"ContributionAmount,omitempty"`

	// contribution designation
	ContributionDesignation *ContributionDesignationSummary `json:"ContributionDesignation,omitempty"`

	// create location
	CreateLocation string `json:"CreateLocation,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// created date time
	// Format: date-time
	CreatedDateTime strfmt.DateTime `json:"CreatedDateTime,omitempty"`

	// custom plan elements
	CustomPlanElements []*PortfolioCustomDataItem `json:"CustomPlanElements"`

	// custom portfolio elements
	CustomPortfolioElements []*PortfolioCustomDataItem `json:"CustomPortfolioElements"`

	// edit indicator
	EditIndicator bool `json:"EditIndicator,omitempty"`

	// fund
	Fund *FundSummary `json:"Fund,omitempty"`

	// goal amount
	GoalAmount float64 `json:"GoalAmount,omitempty"`

	// has open steps
	HasOpenSteps bool `json:"HasOpenSteps,omitempty"`

	// has steps
	HasSteps bool `json:"HasSteps,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// last step date
	// Format: date-time
	LastStepDate strfmt.DateTime `json:"LastStepDate,omitempty"`

	// last step type
	LastStepType *StepTypeSummary `json:"LastStepType,omitempty"`

	// last worker display name
	LastWorkerDisplayName string `json:"LastWorkerDisplayName,omitempty"`

	// next step date
	// Format: date-time
	NextStepDate strfmt.DateTime `json:"NextStepDate,omitempty"`

	// next step type
	NextStepType *StepTypeSummary `json:"NextStepType,omitempty"`

	// notes
	Notes string `json:"Notes,omitempty"`

	// original source
	OriginalSource *PlanSourceSummary `json:"OriginalSource,omitempty"`

	// primary worker
	PrimaryWorker *ConstituentShortDisplayNameSummary `json:"PrimaryWorker,omitempty"`

	// priority
	Priority *PlanPrioritySummary `json:"Priority,omitempty"`

	// probability
	Probability float64 `json:"Probability,omitempty"`

	// recorded amount
	RecordedAmount float64 `json:"RecordedAmount,omitempty"`

	// start date time
	// Format: date-time
	StartDateTime strfmt.DateTime `json:"StartDateTime,omitempty"`

	// status
	Status *PlanStatusSummary `json:"Status,omitempty"`

	// type
	Type *PlanTypeSummary `json:"Type,omitempty"`

	// updated by
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// updated date time
	// Format: date-time
	UpdatedDateTime strfmt.DateTime `json:"UpdatedDateTime,omitempty"`
}

// Validate validates this portfolio plan
func (m *PortfolioPlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCampaign(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompleteByDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConstituent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContributionDesignation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomPlanElements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomPortfolioElements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFund(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastStepDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastStepType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextStepDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextStepType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginalSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryWorker(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
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

func (m *PortfolioPlan) validateCampaign(formats strfmt.Registry) error {
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

func (m *PortfolioPlan) validateCompleteByDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CompleteByDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CompleteByDateTime", "body", "date-time", m.CompleteByDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PortfolioPlan) validateConstituent(formats strfmt.Registry) error {
	if swag.IsZero(m.Constituent) { // not required
		return nil
	}

	if m.Constituent != nil {
		if err := m.Constituent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Constituent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Constituent")
			}
			return err
		}
	}

	return nil
}

func (m *PortfolioPlan) validateContributionDesignation(formats strfmt.Registry) error {
	if swag.IsZero(m.ContributionDesignation) { // not required
		return nil
	}

	if m.ContributionDesignation != nil {
		if err := m.ContributionDesignation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ContributionDesignation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ContributionDesignation")
			}
			return err
		}
	}

	return nil
}

func (m *PortfolioPlan) validateCreatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PortfolioPlan) validateCustomPlanElements(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomPlanElements) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomPlanElements); i++ {
		if swag.IsZero(m.CustomPlanElements[i]) { // not required
			continue
		}

		if m.CustomPlanElements[i] != nil {
			if err := m.CustomPlanElements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CustomPlanElements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("CustomPlanElements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortfolioPlan) validateCustomPortfolioElements(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomPortfolioElements) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomPortfolioElements); i++ {
		if swag.IsZero(m.CustomPortfolioElements[i]) { // not required
			continue
		}

		if m.CustomPortfolioElements[i] != nil {
			if err := m.CustomPortfolioElements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CustomPortfolioElements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("CustomPortfolioElements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortfolioPlan) validateFund(formats strfmt.Registry) error {
	if swag.IsZero(m.Fund) { // not required
		return nil
	}

	if m.Fund != nil {
		if err := m.Fund.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Fund")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Fund")
			}
			return err
		}
	}

	return nil
}

func (m *PortfolioPlan) validateLastStepDate(formats strfmt.Registry) error {
	if swag.IsZero(m.LastStepDate) { // not required
		return nil
	}

	if err := validate.FormatOf("LastStepDate", "body", "date-time", m.LastStepDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PortfolioPlan) validateLastStepType(formats strfmt.Registry) error {
	if swag.IsZero(m.LastStepType) { // not required
		return nil
	}

	if m.LastStepType != nil {
		if err := m.LastStepType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LastStepType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("LastStepType")
			}
			return err
		}
	}

	return nil
}

func (m *PortfolioPlan) validateNextStepDate(formats strfmt.Registry) error {
	if swag.IsZero(m.NextStepDate) { // not required
		return nil
	}

	if err := validate.FormatOf("NextStepDate", "body", "date-time", m.NextStepDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PortfolioPlan) validateNextStepType(formats strfmt.Registry) error {
	if swag.IsZero(m.NextStepType) { // not required
		return nil
	}

	if m.NextStepType != nil {
		if err := m.NextStepType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NextStepType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("NextStepType")
			}
			return err
		}
	}

	return nil
}

func (m *PortfolioPlan) validateOriginalSource(formats strfmt.Registry) error {
	if swag.IsZero(m.OriginalSource) { // not required
		return nil
	}

	if m.OriginalSource != nil {
		if err := m.OriginalSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OriginalSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OriginalSource")
			}
			return err
		}
	}

	return nil
}

func (m *PortfolioPlan) validatePrimaryWorker(formats strfmt.Registry) error {
	if swag.IsZero(m.PrimaryWorker) { // not required
		return nil
	}

	if m.PrimaryWorker != nil {
		if err := m.PrimaryWorker.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PrimaryWorker")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PrimaryWorker")
			}
			return err
		}
	}

	return nil
}

func (m *PortfolioPlan) validatePriority(formats strfmt.Registry) error {
	if swag.IsZero(m.Priority) { // not required
		return nil
	}

	if m.Priority != nil {
		if err := m.Priority.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Priority")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Priority")
			}
			return err
		}
	}

	return nil
}

func (m *PortfolioPlan) validateStartDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("StartDateTime", "body", "date-time", m.StartDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PortfolioPlan) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Status")
			}
			return err
		}
	}

	return nil
}

func (m *PortfolioPlan) validateType(formats strfmt.Registry) error {
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

func (m *PortfolioPlan) validateUpdatedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedDateTime", "body", "date-time", m.UpdatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this portfolio plan based on the context it is used
func (m *PortfolioPlan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCampaign(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConstituent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContributionDesignation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomPlanElements(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomPortfolioElements(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFund(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastStepType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNextStepType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOriginalSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrimaryWorker(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePriority(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortfolioPlan) contextValidateCampaign(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PortfolioPlan) contextValidateConstituent(ctx context.Context, formats strfmt.Registry) error {

	if m.Constituent != nil {

		if swag.IsZero(m.Constituent) { // not required
			return nil
		}

		if err := m.Constituent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Constituent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Constituent")
			}
			return err
		}
	}

	return nil
}

func (m *PortfolioPlan) contextValidateContributionDesignation(ctx context.Context, formats strfmt.Registry) error {

	if m.ContributionDesignation != nil {

		if swag.IsZero(m.ContributionDesignation) { // not required
			return nil
		}

		if err := m.ContributionDesignation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ContributionDesignation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ContributionDesignation")
			}
			return err
		}
	}

	return nil
}

func (m *PortfolioPlan) contextValidateCustomPlanElements(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomPlanElements); i++ {

		if m.CustomPlanElements[i] != nil {

			if swag.IsZero(m.CustomPlanElements[i]) { // not required
				return nil
			}

			if err := m.CustomPlanElements[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CustomPlanElements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("CustomPlanElements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortfolioPlan) contextValidateCustomPortfolioElements(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomPortfolioElements); i++ {

		if m.CustomPortfolioElements[i] != nil {

			if swag.IsZero(m.CustomPortfolioElements[i]) { // not required
				return nil
			}

			if err := m.CustomPortfolioElements[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CustomPortfolioElements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("CustomPortfolioElements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PortfolioPlan) contextValidateFund(ctx context.Context, formats strfmt.Registry) error {

	if m.Fund != nil {

		if swag.IsZero(m.Fund) { // not required
			return nil
		}

		if err := m.Fund.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Fund")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Fund")
			}
			return err
		}
	}

	return nil
}

func (m *PortfolioPlan) contextValidateLastStepType(ctx context.Context, formats strfmt.Registry) error {

	if m.LastStepType != nil {

		if swag.IsZero(m.LastStepType) { // not required
			return nil
		}

		if err := m.LastStepType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LastStepType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("LastStepType")
			}
			return err
		}
	}

	return nil
}

func (m *PortfolioPlan) contextValidateNextStepType(ctx context.Context, formats strfmt.Registry) error {

	if m.NextStepType != nil {

		if swag.IsZero(m.NextStepType) { // not required
			return nil
		}

		if err := m.NextStepType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NextStepType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("NextStepType")
			}
			return err
		}
	}

	return nil
}

func (m *PortfolioPlan) contextValidateOriginalSource(ctx context.Context, formats strfmt.Registry) error {

	if m.OriginalSource != nil {

		if swag.IsZero(m.OriginalSource) { // not required
			return nil
		}

		if err := m.OriginalSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OriginalSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OriginalSource")
			}
			return err
		}
	}

	return nil
}

func (m *PortfolioPlan) contextValidatePrimaryWorker(ctx context.Context, formats strfmt.Registry) error {

	if m.PrimaryWorker != nil {

		if swag.IsZero(m.PrimaryWorker) { // not required
			return nil
		}

		if err := m.PrimaryWorker.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PrimaryWorker")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PrimaryWorker")
			}
			return err
		}
	}

	return nil
}

func (m *PortfolioPlan) contextValidatePriority(ctx context.Context, formats strfmt.Registry) error {

	if m.Priority != nil {

		if swag.IsZero(m.Priority) { // not required
			return nil
		}

		if err := m.Priority.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Priority")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Priority")
			}
			return err
		}
	}

	return nil
}

func (m *PortfolioPlan) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if swag.IsZero(m.Status) { // not required
			return nil
		}

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Status")
			}
			return err
		}
	}

	return nil
}

func (m *PortfolioPlan) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *PortfolioPlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortfolioPlan) UnmarshalBinary(b []byte) error {
	var res PortfolioPlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
