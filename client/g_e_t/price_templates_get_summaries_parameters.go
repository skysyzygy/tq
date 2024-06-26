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

// NewPriceTemplatesGetSummariesParams creates a new PriceTemplatesGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPriceTemplatesGetSummariesParams() *PriceTemplatesGetSummariesParams {
	return &PriceTemplatesGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPriceTemplatesGetSummariesParamsWithTimeout creates a new PriceTemplatesGetSummariesParams object
// with the ability to set a timeout on a request.
func NewPriceTemplatesGetSummariesParamsWithTimeout(timeout time.Duration) *PriceTemplatesGetSummariesParams {
	return &PriceTemplatesGetSummariesParams{
		timeout: timeout,
	}
}

// NewPriceTemplatesGetSummariesParamsWithContext creates a new PriceTemplatesGetSummariesParams object
// with the ability to set a context for a request.
func NewPriceTemplatesGetSummariesParamsWithContext(ctx context.Context) *PriceTemplatesGetSummariesParams {
	return &PriceTemplatesGetSummariesParams{
		Context: ctx,
	}
}

// NewPriceTemplatesGetSummariesParamsWithHTTPClient creates a new PriceTemplatesGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPriceTemplatesGetSummariesParamsWithHTTPClient(client *http.Client) *PriceTemplatesGetSummariesParams {
	return &PriceTemplatesGetSummariesParams{
		HTTPClient: client,
	}
}

/*
PriceTemplatesGetSummariesParams contains all the parameters to send to the API endpoint

	for the price templates get summaries operation.

	Typically these are written to a http.Request.
*/
type PriceTemplatesGetSummariesParams struct {

	// FacilityID.
	FacilityID *string

	// ZoneMapID.
	ZoneMapID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the price templates get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PriceTemplatesGetSummariesParams) WithDefaults() *PriceTemplatesGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the price templates get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PriceTemplatesGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the price templates get summaries params
func (o *PriceTemplatesGetSummariesParams) WithTimeout(timeout time.Duration) *PriceTemplatesGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the price templates get summaries params
func (o *PriceTemplatesGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the price templates get summaries params
func (o *PriceTemplatesGetSummariesParams) WithContext(ctx context.Context) *PriceTemplatesGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the price templates get summaries params
func (o *PriceTemplatesGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the price templates get summaries params
func (o *PriceTemplatesGetSummariesParams) WithHTTPClient(client *http.Client) *PriceTemplatesGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the price templates get summaries params
func (o *PriceTemplatesGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFacilityID adds the facilityID to the price templates get summaries params
func (o *PriceTemplatesGetSummariesParams) WithFacilityID(facilityID *string) *PriceTemplatesGetSummariesParams {
	o.SetFacilityID(facilityID)
	return o
}

// SetFacilityID adds the facilityId to the price templates get summaries params
func (o *PriceTemplatesGetSummariesParams) SetFacilityID(facilityID *string) {
	o.FacilityID = facilityID
}

// WithZoneMapID adds the zoneMapID to the price templates get summaries params
func (o *PriceTemplatesGetSummariesParams) WithZoneMapID(zoneMapID *string) *PriceTemplatesGetSummariesParams {
	o.SetZoneMapID(zoneMapID)
	return o
}

// SetZoneMapID adds the zoneMapId to the price templates get summaries params
func (o *PriceTemplatesGetSummariesParams) SetZoneMapID(zoneMapID *string) {
	o.ZoneMapID = zoneMapID
}

// WriteToRequest writes these params to a swagger request
func (o *PriceTemplatesGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FacilityID != nil {

		// query param facilityId
		var qrFacilityID string

		if o.FacilityID != nil {
			qrFacilityID = *o.FacilityID
		}
		qFacilityID := qrFacilityID
		if qFacilityID != "" {

			if err := r.SetQueryParam("facilityId", qFacilityID); err != nil {
				return err
			}
		}
	}

	if o.ZoneMapID != nil {

		// query param zoneMapId
		var qrZoneMapID string

		if o.ZoneMapID != nil {
			qrZoneMapID = *o.ZoneMapID
		}
		qZoneMapID := qrZoneMapID
		if qZoneMapID != "" {

			if err := r.SetQueryParam("zoneMapId", qZoneMapID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
