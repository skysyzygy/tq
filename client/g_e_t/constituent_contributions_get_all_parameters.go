// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewConstituentContributionsGetAllParams creates a new ConstituentContributionsGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConstituentContributionsGetAllParams() *ConstituentContributionsGetAllParams {
	return &ConstituentContributionsGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConstituentContributionsGetAllParamsWithTimeout creates a new ConstituentContributionsGetAllParams object
// with the ability to set a timeout on a request.
func NewConstituentContributionsGetAllParamsWithTimeout(timeout time.Duration) *ConstituentContributionsGetAllParams {
	return &ConstituentContributionsGetAllParams{
		timeout: timeout,
	}
}

// NewConstituentContributionsGetAllParamsWithContext creates a new ConstituentContributionsGetAllParams object
// with the ability to set a context for a request.
func NewConstituentContributionsGetAllParamsWithContext(ctx context.Context) *ConstituentContributionsGetAllParams {
	return &ConstituentContributionsGetAllParams{
		Context: ctx,
	}
}

// NewConstituentContributionsGetAllParamsWithHTTPClient creates a new ConstituentContributionsGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewConstituentContributionsGetAllParamsWithHTTPClient(client *http.Client) *ConstituentContributionsGetAllParams {
	return &ConstituentContributionsGetAllParams{
		HTTPClient: client,
	}
}

/*
ConstituentContributionsGetAllParams contains all the parameters to send to the API endpoint

	for the constituent contributions get all operation.

	Typically these are written to a http.Request.
*/
type ConstituentContributionsGetAllParams struct {

	// CampaignIds.
	CampaignIds *string

	/* ConstituentID.

	   Limit results by constituent
	*/
	ConstituentID *string

	// EndDate.
	EndDate *string

	// FundIds.
	FundIds *string

	/* IncludeAffiliations.

	   Include all of the constituent's affiliates in the results (default: false)
	*/
	IncludeAffiliations *string

	// StartDate.
	StartDate *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the constituent contributions get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConstituentContributionsGetAllParams) WithDefaults() *ConstituentContributionsGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the constituent contributions get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConstituentContributionsGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the constituent contributions get all params
func (o *ConstituentContributionsGetAllParams) WithTimeout(timeout time.Duration) *ConstituentContributionsGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the constituent contributions get all params
func (o *ConstituentContributionsGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the constituent contributions get all params
func (o *ConstituentContributionsGetAllParams) WithContext(ctx context.Context) *ConstituentContributionsGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the constituent contributions get all params
func (o *ConstituentContributionsGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the constituent contributions get all params
func (o *ConstituentContributionsGetAllParams) WithHTTPClient(client *http.Client) *ConstituentContributionsGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the constituent contributions get all params
func (o *ConstituentContributionsGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCampaignIds adds the campaignIds to the constituent contributions get all params
func (o *ConstituentContributionsGetAllParams) WithCampaignIds(campaignIds *string) *ConstituentContributionsGetAllParams {
	o.SetCampaignIds(campaignIds)
	return o
}

// SetCampaignIds adds the campaignIds to the constituent contributions get all params
func (o *ConstituentContributionsGetAllParams) SetCampaignIds(campaignIds *string) {
	o.CampaignIds = campaignIds
}

// WithConstituentID adds the constituentID to the constituent contributions get all params
func (o *ConstituentContributionsGetAllParams) WithConstituentID(constituentID *string) *ConstituentContributionsGetAllParams {
	o.SetConstituentID(constituentID)
	return o
}

// SetConstituentID adds the constituentId to the constituent contributions get all params
func (o *ConstituentContributionsGetAllParams) SetConstituentID(constituentID *string) {
	o.ConstituentID = constituentID
}

// WithEndDate adds the endDate to the constituent contributions get all params
func (o *ConstituentContributionsGetAllParams) WithEndDate(endDate *string) *ConstituentContributionsGetAllParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the constituent contributions get all params
func (o *ConstituentContributionsGetAllParams) SetEndDate(endDate *string) {
	o.EndDate = endDate
}

// WithFundIds adds the fundIds to the constituent contributions get all params
func (o *ConstituentContributionsGetAllParams) WithFundIds(fundIds *string) *ConstituentContributionsGetAllParams {
	o.SetFundIds(fundIds)
	return o
}

// SetFundIds adds the fundIds to the constituent contributions get all params
func (o *ConstituentContributionsGetAllParams) SetFundIds(fundIds *string) {
	o.FundIds = fundIds
}

// WithIncludeAffiliations adds the includeAffiliations to the constituent contributions get all params
func (o *ConstituentContributionsGetAllParams) WithIncludeAffiliations(includeAffiliations *string) *ConstituentContributionsGetAllParams {
	o.SetIncludeAffiliations(includeAffiliations)
	return o
}

// SetIncludeAffiliations adds the includeAffiliations to the constituent contributions get all params
func (o *ConstituentContributionsGetAllParams) SetIncludeAffiliations(includeAffiliations *string) {
	o.IncludeAffiliations = includeAffiliations
}

// WithStartDate adds the startDate to the constituent contributions get all params
func (o *ConstituentContributionsGetAllParams) WithStartDate(startDate *string) *ConstituentContributionsGetAllParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the constituent contributions get all params
func (o *ConstituentContributionsGetAllParams) SetStartDate(startDate *string) {
	o.StartDate = startDate
}

// WriteToRequest writes these params to a swagger request
func (o *ConstituentContributionsGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CampaignIds != nil {

		// query param campaignIds
		var qrCampaignIds string

		if o.CampaignIds != nil {
			qrCampaignIds = *o.CampaignIds
		}
		qCampaignIds := qrCampaignIds
		if qCampaignIds != "" {

			if err := r.SetQueryParam("campaignIds", qCampaignIds); err != nil {
				return err
			}
		}
	}

	if o.ConstituentID != nil {

		// query param constituentId
		var qrConstituentID string

		if o.ConstituentID != nil {
			qrConstituentID = *o.ConstituentID
		}
		qConstituentID := qrConstituentID
		if qConstituentID != "" {

			if err := r.SetQueryParam("constituentId", qConstituentID); err != nil {
				return err
			}
		}
	}

	if o.EndDate != nil {

		// query param endDate
		var qrEndDate string

		if o.EndDate != nil {
			qrEndDate = *o.EndDate
		}
		qEndDate := qrEndDate
		if qEndDate != "" {

			if err := r.SetQueryParam("endDate", qEndDate); err != nil {
				return err
			}
		}
	}

	if o.FundIds != nil {

		// query param fundIds
		var qrFundIds string

		if o.FundIds != nil {
			qrFundIds = *o.FundIds
		}
		qFundIds := qrFundIds
		if qFundIds != "" {

			if err := r.SetQueryParam("fundIds", qFundIds); err != nil {
				return err
			}
		}
	}

	if o.IncludeAffiliations != nil {

		// query param includeAffiliations
		var qrIncludeAffiliations string

		if o.IncludeAffiliations != nil {
			qrIncludeAffiliations = *o.IncludeAffiliations
		}
		qIncludeAffiliations := qrIncludeAffiliations
		if qIncludeAffiliations != "" {

			if err := r.SetQueryParam("includeAffiliations", qIncludeAffiliations); err != nil {
				return err
			}
		}
	}

	if o.StartDate != nil {

		// query param startDate
		var qrStartDate string

		if o.StartDate != nil {
			qrStartDate = *o.StartDate
		}
		qStartDate := qrStartDate
		if qStartDate != "" {

			if err := r.SetQueryParam("startDate", qStartDate); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
