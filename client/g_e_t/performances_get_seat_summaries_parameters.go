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

// NewPerformancesGetSeatSummariesParams creates a new PerformancesGetSeatSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformancesGetSeatSummariesParams() *PerformancesGetSeatSummariesParams {
	return &PerformancesGetSeatSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformancesGetSeatSummariesParamsWithTimeout creates a new PerformancesGetSeatSummariesParams object
// with the ability to set a timeout on a request.
func NewPerformancesGetSeatSummariesParamsWithTimeout(timeout time.Duration) *PerformancesGetSeatSummariesParams {
	return &PerformancesGetSeatSummariesParams{
		timeout: timeout,
	}
}

// NewPerformancesGetSeatSummariesParamsWithContext creates a new PerformancesGetSeatSummariesParams object
// with the ability to set a context for a request.
func NewPerformancesGetSeatSummariesParamsWithContext(ctx context.Context) *PerformancesGetSeatSummariesParams {
	return &PerformancesGetSeatSummariesParams{
		Context: ctx,
	}
}

// NewPerformancesGetSeatSummariesParamsWithHTTPClient creates a new PerformancesGetSeatSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformancesGetSeatSummariesParamsWithHTTPClient(client *http.Client) *PerformancesGetSeatSummariesParams {
	return &PerformancesGetSeatSummariesParams{
		HTTPClient: client,
	}
}

/*
PerformancesGetSeatSummariesParams contains all the parameters to send to the API endpoint

	for the performances get seat summaries operation.

	Typically these are written to a http.Request.
*/
type PerformancesGetSeatSummariesParams struct {

	/* CheckPriceTypeIds.

	   checkPriceTypeIds must be either a list of valid price types or the token "All"
	*/
	CheckPriceTypeIds *string

	/* ConstituentID.

	   Required parameter. Must be a valid constituent ID
	*/
	ConstituentID *string

	/* ModeOfSaleID.

	   Required parameter. Must be a valid MOS id
	*/
	ModeOfSaleID *string

	/* PerformanceID.

	   ID of the performance
	*/
	PerformanceID string

	/* ScreenIds.

	   screenIds must be a comma separated list of valid screen ids for the specified performance
	*/
	ScreenIds *string

	/* SectionIds.

	   sectionIds must be a comma separated list of valid section ids
	*/
	SectionIds *string

	/* ZoneIds.

	   zoneIds must be a comma separated list of valid zone ids
	*/
	ZoneIds *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the performances get seat summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformancesGetSeatSummariesParams) WithDefaults() *PerformancesGetSeatSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the performances get seat summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformancesGetSeatSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the performances get seat summaries params
func (o *PerformancesGetSeatSummariesParams) WithTimeout(timeout time.Duration) *PerformancesGetSeatSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the performances get seat summaries params
func (o *PerformancesGetSeatSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the performances get seat summaries params
func (o *PerformancesGetSeatSummariesParams) WithContext(ctx context.Context) *PerformancesGetSeatSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the performances get seat summaries params
func (o *PerformancesGetSeatSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the performances get seat summaries params
func (o *PerformancesGetSeatSummariesParams) WithHTTPClient(client *http.Client) *PerformancesGetSeatSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the performances get seat summaries params
func (o *PerformancesGetSeatSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCheckPriceTypeIds adds the checkPriceTypeIds to the performances get seat summaries params
func (o *PerformancesGetSeatSummariesParams) WithCheckPriceTypeIds(checkPriceTypeIds *string) *PerformancesGetSeatSummariesParams {
	o.SetCheckPriceTypeIds(checkPriceTypeIds)
	return o
}

// SetCheckPriceTypeIds adds the checkPriceTypeIds to the performances get seat summaries params
func (o *PerformancesGetSeatSummariesParams) SetCheckPriceTypeIds(checkPriceTypeIds *string) {
	o.CheckPriceTypeIds = checkPriceTypeIds
}

// WithConstituentID adds the constituentID to the performances get seat summaries params
func (o *PerformancesGetSeatSummariesParams) WithConstituentID(constituentID *string) *PerformancesGetSeatSummariesParams {
	o.SetConstituentID(constituentID)
	return o
}

// SetConstituentID adds the constituentId to the performances get seat summaries params
func (o *PerformancesGetSeatSummariesParams) SetConstituentID(constituentID *string) {
	o.ConstituentID = constituentID
}

// WithModeOfSaleID adds the modeOfSaleID to the performances get seat summaries params
func (o *PerformancesGetSeatSummariesParams) WithModeOfSaleID(modeOfSaleID *string) *PerformancesGetSeatSummariesParams {
	o.SetModeOfSaleID(modeOfSaleID)
	return o
}

// SetModeOfSaleID adds the modeOfSaleId to the performances get seat summaries params
func (o *PerformancesGetSeatSummariesParams) SetModeOfSaleID(modeOfSaleID *string) {
	o.ModeOfSaleID = modeOfSaleID
}

// WithPerformanceID adds the performanceID to the performances get seat summaries params
func (o *PerformancesGetSeatSummariesParams) WithPerformanceID(performanceID string) *PerformancesGetSeatSummariesParams {
	o.SetPerformanceID(performanceID)
	return o
}

// SetPerformanceID adds the performanceId to the performances get seat summaries params
func (o *PerformancesGetSeatSummariesParams) SetPerformanceID(performanceID string) {
	o.PerformanceID = performanceID
}

// WithScreenIds adds the screenIds to the performances get seat summaries params
func (o *PerformancesGetSeatSummariesParams) WithScreenIds(screenIds *string) *PerformancesGetSeatSummariesParams {
	o.SetScreenIds(screenIds)
	return o
}

// SetScreenIds adds the screenIds to the performances get seat summaries params
func (o *PerformancesGetSeatSummariesParams) SetScreenIds(screenIds *string) {
	o.ScreenIds = screenIds
}

// WithSectionIds adds the sectionIds to the performances get seat summaries params
func (o *PerformancesGetSeatSummariesParams) WithSectionIds(sectionIds *string) *PerformancesGetSeatSummariesParams {
	o.SetSectionIds(sectionIds)
	return o
}

// SetSectionIds adds the sectionIds to the performances get seat summaries params
func (o *PerformancesGetSeatSummariesParams) SetSectionIds(sectionIds *string) {
	o.SectionIds = sectionIds
}

// WithZoneIds adds the zoneIds to the performances get seat summaries params
func (o *PerformancesGetSeatSummariesParams) WithZoneIds(zoneIds *string) *PerformancesGetSeatSummariesParams {
	o.SetZoneIds(zoneIds)
	return o
}

// SetZoneIds adds the zoneIds to the performances get seat summaries params
func (o *PerformancesGetSeatSummariesParams) SetZoneIds(zoneIds *string) {
	o.ZoneIds = zoneIds
}

// WriteToRequest writes these params to a swagger request
func (o *PerformancesGetSeatSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CheckPriceTypeIds != nil {

		// query param checkPriceTypeIds
		var qrCheckPriceTypeIds string

		if o.CheckPriceTypeIds != nil {
			qrCheckPriceTypeIds = *o.CheckPriceTypeIds
		}
		qCheckPriceTypeIds := qrCheckPriceTypeIds
		if qCheckPriceTypeIds != "" {

			if err := r.SetQueryParam("checkPriceTypeIds", qCheckPriceTypeIds); err != nil {
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

	if o.ModeOfSaleID != nil {

		// query param modeOfSaleId
		var qrModeOfSaleID string

		if o.ModeOfSaleID != nil {
			qrModeOfSaleID = *o.ModeOfSaleID
		}
		qModeOfSaleID := qrModeOfSaleID
		if qModeOfSaleID != "" {

			if err := r.SetQueryParam("modeOfSaleId", qModeOfSaleID); err != nil {
				return err
			}
		}
	}

	// path param performanceId
	if err := r.SetPathParam("performanceId", o.PerformanceID); err != nil {
		return err
	}

	if o.ScreenIds != nil {

		// query param screenIds
		var qrScreenIds string

		if o.ScreenIds != nil {
			qrScreenIds = *o.ScreenIds
		}
		qScreenIds := qrScreenIds
		if qScreenIds != "" {

			if err := r.SetQueryParam("screenIds", qScreenIds); err != nil {
				return err
			}
		}
	}

	if o.SectionIds != nil {

		// query param sectionIds
		var qrSectionIds string

		if o.SectionIds != nil {
			qrSectionIds = *o.SectionIds
		}
		qSectionIds := qrSectionIds
		if qSectionIds != "" {

			if err := r.SetQueryParam("sectionIds", qSectionIds); err != nil {
				return err
			}
		}
	}

	if o.ZoneIds != nil {

		// query param zoneIds
		var qrZoneIds string

		if o.ZoneIds != nil {
			qrZoneIds = *o.ZoneIds
		}
		qZoneIds := qrZoneIds
		if qZoneIds != "" {

			if err := r.SetQueryParam("zoneIds", qZoneIds); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}