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

// CartSummary cart summary
//
// swagger:model CartSummary
type CartSummary struct {

	// amount paid now
	AmountPaidNow float64 `json:"AmountPaidNow,omitempty"`

	// appeal
	Appeal *EntitySummary `json:"Appeal,omitempty"`

	// balance to charge
	BalanceToCharge float64 `json:"BalanceToCharge,omitempty"`

	// batch Id
	BatchID int32 `json:"BatchId,omitempty"`

	// booking Id
	BookingID int32 `json:"BookingId,omitempty"`

	// cart amount
	CartAmount float64 `json:"CartAmount,omitempty"`

	// cart was priced
	CartWasPriced bool `json:"CartWasPriced,omitempty"`

	// constituent
	Constituent *ConstituentDisplaySummary `json:"Constituent,omitempty"`

	// delivery date
	// Format: date-time
	DeliveryDate *strfmt.DateTime `json:"DeliveryDate,omitempty"`

	// delivery method
	DeliveryMethod *EntitySummary `json:"DeliveryMethod,omitempty"`

	// fees amount
	FeesAmount float64 `json:"FeesAmount,omitempty"`

	// first seat added date time
	// Format: date-time
	FirstSeatAddedDateTime *strfmt.DateTime `json:"FirstSeatAddedDateTime,omitempty"`

	// Id
	ID int32 `json:"Id,omitempty"`

	// initiator
	Initiator *ConstituentDisplaySummary `json:"Initiator,omitempty"`

	// messages
	Messages []*CartPricingRuleMessage `json:"Messages"`

	// mode of sale
	ModeOfSale *EntitySummary `json:"ModeOfSale,omitempty"`

	// order category
	OrderCategory *OrderCategorySummary `json:"OrderCategory,omitempty"`

	// order date time
	// Format: date-time
	OrderDateTime *strfmt.DateTime `json:"OrderDateTime,omitempty"`

	// order notes
	OrderNotes string `json:"OrderNotes,omitempty"`

	// payments
	Payments []*CartPayment `json:"Payments"`

	// products
	Products []*CartProductSummary `json:"Products"`

	// rule based fees
	RuleBasedFees []*RuleBasedFee `json:"RuleBasedFees"`

	// session key
	SessionKey string `json:"SessionKey,omitempty"`

	// solicitor
	Solicitor string `json:"Solicitor,omitempty"`

	// source
	Source *EntitySummary `json:"Source,omitempty"`

	// sub total
	SubTotal float64 `json:"SubTotal,omitempty"`
}

// Validate validates this cart summary
func (m *CartSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppeal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConstituent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeliveryDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeliveryMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstSeatAddedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitiator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModeOfSale(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProducts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleBasedFees(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CartSummary) validateAppeal(formats strfmt.Registry) error {
	if swag.IsZero(m.Appeal) { // not required
		return nil
	}

	if m.Appeal != nil {
		if err := m.Appeal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Appeal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Appeal")
			}
			return err
		}
	}

	return nil
}

func (m *CartSummary) validateConstituent(formats strfmt.Registry) error {
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

func (m *CartSummary) validateDeliveryDate(formats strfmt.Registry) error {
	if swag.IsZero(m.DeliveryDate) { // not required
		return nil
	}

	if err := validate.FormatOf("DeliveryDate", "body", "date-time", m.DeliveryDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CartSummary) validateDeliveryMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.DeliveryMethod) { // not required
		return nil
	}

	if m.DeliveryMethod != nil {
		if err := m.DeliveryMethod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DeliveryMethod")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DeliveryMethod")
			}
			return err
		}
	}

	return nil
}

func (m *CartSummary) validateFirstSeatAddedDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.FirstSeatAddedDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("FirstSeatAddedDateTime", "body", "date-time", m.FirstSeatAddedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CartSummary) validateInitiator(formats strfmt.Registry) error {
	if swag.IsZero(m.Initiator) { // not required
		return nil
	}

	if m.Initiator != nil {
		if err := m.Initiator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Initiator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Initiator")
			}
			return err
		}
	}

	return nil
}

func (m *CartSummary) validateMessages(formats strfmt.Registry) error {
	if swag.IsZero(m.Messages) { // not required
		return nil
	}

	for i := 0; i < len(m.Messages); i++ {
		if swag.IsZero(m.Messages[i]) { // not required
			continue
		}

		if m.Messages[i] != nil {
			if err := m.Messages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Messages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Messages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CartSummary) validateModeOfSale(formats strfmt.Registry) error {
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

func (m *CartSummary) validateOrderCategory(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderCategory) { // not required
		return nil
	}

	if m.OrderCategory != nil {
		if err := m.OrderCategory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OrderCategory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OrderCategory")
			}
			return err
		}
	}

	return nil
}

func (m *CartSummary) validateOrderDateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("OrderDateTime", "body", "date-time", m.OrderDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CartSummary) validatePayments(formats strfmt.Registry) error {
	if swag.IsZero(m.Payments) { // not required
		return nil
	}

	for i := 0; i < len(m.Payments); i++ {
		if swag.IsZero(m.Payments[i]) { // not required
			continue
		}

		if m.Payments[i] != nil {
			if err := m.Payments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Payments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Payments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CartSummary) validateProducts(formats strfmt.Registry) error {
	if swag.IsZero(m.Products) { // not required
		return nil
	}

	for i := 0; i < len(m.Products); i++ {
		if swag.IsZero(m.Products[i]) { // not required
			continue
		}

		if m.Products[i] != nil {
			if err := m.Products[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Products" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Products" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CartSummary) validateRuleBasedFees(formats strfmt.Registry) error {
	if swag.IsZero(m.RuleBasedFees) { // not required
		return nil
	}

	for i := 0; i < len(m.RuleBasedFees); i++ {
		if swag.IsZero(m.RuleBasedFees[i]) { // not required
			continue
		}

		if m.RuleBasedFees[i] != nil {
			if err := m.RuleBasedFees[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RuleBasedFees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RuleBasedFees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CartSummary) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Source")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cart summary based on the context it is used
func (m *CartSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppeal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConstituent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeliveryMethod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInitiator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModeOfSale(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrderCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePayments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProducts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRuleBasedFees(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CartSummary) contextValidateAppeal(ctx context.Context, formats strfmt.Registry) error {

	if m.Appeal != nil {

		if swag.IsZero(m.Appeal) { // not required
			return nil
		}

		if err := m.Appeal.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Appeal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Appeal")
			}
			return err
		}
	}

	return nil
}

func (m *CartSummary) contextValidateConstituent(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CartSummary) contextValidateDeliveryMethod(ctx context.Context, formats strfmt.Registry) error {

	if m.DeliveryMethod != nil {

		if swag.IsZero(m.DeliveryMethod) { // not required
			return nil
		}

		if err := m.DeliveryMethod.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DeliveryMethod")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DeliveryMethod")
			}
			return err
		}
	}

	return nil
}

func (m *CartSummary) contextValidateInitiator(ctx context.Context, formats strfmt.Registry) error {

	if m.Initiator != nil {

		if swag.IsZero(m.Initiator) { // not required
			return nil
		}

		if err := m.Initiator.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Initiator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Initiator")
			}
			return err
		}
	}

	return nil
}

func (m *CartSummary) contextValidateMessages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Messages); i++ {

		if m.Messages[i] != nil {

			if swag.IsZero(m.Messages[i]) { // not required
				return nil
			}

			if err := m.Messages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Messages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Messages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CartSummary) contextValidateModeOfSale(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CartSummary) contextValidateOrderCategory(ctx context.Context, formats strfmt.Registry) error {

	if m.OrderCategory != nil {

		if swag.IsZero(m.OrderCategory) { // not required
			return nil
		}

		if err := m.OrderCategory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OrderCategory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OrderCategory")
			}
			return err
		}
	}

	return nil
}

func (m *CartSummary) contextValidatePayments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Payments); i++ {

		if m.Payments[i] != nil {

			if swag.IsZero(m.Payments[i]) { // not required
				return nil
			}

			if err := m.Payments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Payments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Payments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CartSummary) contextValidateProducts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Products); i++ {

		if m.Products[i] != nil {

			if swag.IsZero(m.Products[i]) { // not required
				return nil
			}

			if err := m.Products[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Products" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Products" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CartSummary) contextValidateRuleBasedFees(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RuleBasedFees); i++ {

		if m.RuleBasedFees[i] != nil {

			if swag.IsZero(m.RuleBasedFees[i]) { // not required
				return nil
			}

			if err := m.RuleBasedFees[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RuleBasedFees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RuleBasedFees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CartSummary) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {

		if swag.IsZero(m.Source) { // not required
			return nil
		}

		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CartSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CartSummary) UnmarshalBinary(b []byte) error {
	var res CartSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
