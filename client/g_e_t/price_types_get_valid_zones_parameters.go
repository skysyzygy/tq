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

// NewPriceTypesGetValidZonesParams creates a new PriceTypesGetValidZonesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPriceTypesGetValidZonesParams() *PriceTypesGetValidZonesParams {
	return &PriceTypesGetValidZonesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPriceTypesGetValidZonesParamsWithTimeout creates a new PriceTypesGetValidZonesParams object
// with the ability to set a timeout on a request.
func NewPriceTypesGetValidZonesParamsWithTimeout(timeout time.Duration) *PriceTypesGetValidZonesParams {
	return &PriceTypesGetValidZonesParams{
		timeout: timeout,
	}
}

// NewPriceTypesGetValidZonesParamsWithContext creates a new PriceTypesGetValidZonesParams object
// with the ability to set a context for a request.
func NewPriceTypesGetValidZonesParamsWithContext(ctx context.Context) *PriceTypesGetValidZonesParams {
	return &PriceTypesGetValidZonesParams{
		Context: ctx,
	}
}

// NewPriceTypesGetValidZonesParamsWithHTTPClient creates a new PriceTypesGetValidZonesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPriceTypesGetValidZonesParamsWithHTTPClient(client *http.Client) *PriceTypesGetValidZonesParams {
	return &PriceTypesGetValidZonesParams{
		HTTPClient: client,
	}
}

/*
PriceTypesGetValidZonesParams contains all the parameters to send to the API endpoint

	for the price types get valid zones operation.

	Typically these are written to a http.Request.
*/
type PriceTypesGetValidZonesParams struct {

	// OrderDateTime.
	OrderDateTime *string

	// PackageID.
	PackageID *string

	// PerformanceID.
	PerformanceID *string

	// PriceTypeID.
	PriceTypeID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the price types get valid zones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PriceTypesGetValidZonesParams) WithDefaults() *PriceTypesGetValidZonesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the price types get valid zones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PriceTypesGetValidZonesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the price types get valid zones params
func (o *PriceTypesGetValidZonesParams) WithTimeout(timeout time.Duration) *PriceTypesGetValidZonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the price types get valid zones params
func (o *PriceTypesGetValidZonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the price types get valid zones params
func (o *PriceTypesGetValidZonesParams) WithContext(ctx context.Context) *PriceTypesGetValidZonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the price types get valid zones params
func (o *PriceTypesGetValidZonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the price types get valid zones params
func (o *PriceTypesGetValidZonesParams) WithHTTPClient(client *http.Client) *PriceTypesGetValidZonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the price types get valid zones params
func (o *PriceTypesGetValidZonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderDateTime adds the orderDateTime to the price types get valid zones params
func (o *PriceTypesGetValidZonesParams) WithOrderDateTime(orderDateTime *string) *PriceTypesGetValidZonesParams {
	o.SetOrderDateTime(orderDateTime)
	return o
}

// SetOrderDateTime adds the orderDateTime to the price types get valid zones params
func (o *PriceTypesGetValidZonesParams) SetOrderDateTime(orderDateTime *string) {
	o.OrderDateTime = orderDateTime
}

// WithPackageID adds the packageID to the price types get valid zones params
func (o *PriceTypesGetValidZonesParams) WithPackageID(packageID *string) *PriceTypesGetValidZonesParams {
	o.SetPackageID(packageID)
	return o
}

// SetPackageID adds the packageId to the price types get valid zones params
func (o *PriceTypesGetValidZonesParams) SetPackageID(packageID *string) {
	o.PackageID = packageID
}

// WithPerformanceID adds the performanceID to the price types get valid zones params
func (o *PriceTypesGetValidZonesParams) WithPerformanceID(performanceID *string) *PriceTypesGetValidZonesParams {
	o.SetPerformanceID(performanceID)
	return o
}

// SetPerformanceID adds the performanceId to the price types get valid zones params
func (o *PriceTypesGetValidZonesParams) SetPerformanceID(performanceID *string) {
	o.PerformanceID = performanceID
}

// WithPriceTypeID adds the priceTypeID to the price types get valid zones params
func (o *PriceTypesGetValidZonesParams) WithPriceTypeID(priceTypeID *string) *PriceTypesGetValidZonesParams {
	o.SetPriceTypeID(priceTypeID)
	return o
}

// SetPriceTypeID adds the priceTypeId to the price types get valid zones params
func (o *PriceTypesGetValidZonesParams) SetPriceTypeID(priceTypeID *string) {
	o.PriceTypeID = priceTypeID
}

// WriteToRequest writes these params to a swagger request
func (o *PriceTypesGetValidZonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.OrderDateTime != nil {

		// query param orderDateTime
		var qrOrderDateTime string

		if o.OrderDateTime != nil {
			qrOrderDateTime = *o.OrderDateTime
		}
		qOrderDateTime := qrOrderDateTime
		if qOrderDateTime != "" {

			if err := r.SetQueryParam("orderDateTime", qOrderDateTime); err != nil {
				return err
			}
		}
	}

	if o.PackageID != nil {

		// query param packageId
		var qrPackageID string

		if o.PackageID != nil {
			qrPackageID = *o.PackageID
		}
		qPackageID := qrPackageID
		if qPackageID != "" {

			if err := r.SetQueryParam("packageId", qPackageID); err != nil {
				return err
			}
		}
	}

	if o.PerformanceID != nil {

		// query param performanceId
		var qrPerformanceID string

		if o.PerformanceID != nil {
			qrPerformanceID = *o.PerformanceID
		}
		qPerformanceID := qrPerformanceID
		if qPerformanceID != "" {

			if err := r.SetQueryParam("performanceId", qPerformanceID); err != nil {
				return err
			}
		}
	}

	if o.PriceTypeID != nil {

		// query param priceTypeId
		var qrPriceTypeID string

		if o.PriceTypeID != nil {
			qrPriceTypeID = *o.PriceTypeID
		}
		qPriceTypeID := qrPriceTypeID
		if qPriceTypeID != "" {

			if err := r.SetQueryParam("priceTypeId", qPriceTypeID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}