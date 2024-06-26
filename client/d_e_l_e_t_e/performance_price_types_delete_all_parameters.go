// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

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

// NewPerformancePriceTypesDeleteAllParams creates a new PerformancePriceTypesDeleteAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformancePriceTypesDeleteAllParams() *PerformancePriceTypesDeleteAllParams {
	return &PerformancePriceTypesDeleteAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformancePriceTypesDeleteAllParamsWithTimeout creates a new PerformancePriceTypesDeleteAllParams object
// with the ability to set a timeout on a request.
func NewPerformancePriceTypesDeleteAllParamsWithTimeout(timeout time.Duration) *PerformancePriceTypesDeleteAllParams {
	return &PerformancePriceTypesDeleteAllParams{
		timeout: timeout,
	}
}

// NewPerformancePriceTypesDeleteAllParamsWithContext creates a new PerformancePriceTypesDeleteAllParams object
// with the ability to set a context for a request.
func NewPerformancePriceTypesDeleteAllParamsWithContext(ctx context.Context) *PerformancePriceTypesDeleteAllParams {
	return &PerformancePriceTypesDeleteAllParams{
		Context: ctx,
	}
}

// NewPerformancePriceTypesDeleteAllParamsWithHTTPClient creates a new PerformancePriceTypesDeleteAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformancePriceTypesDeleteAllParamsWithHTTPClient(client *http.Client) *PerformancePriceTypesDeleteAllParams {
	return &PerformancePriceTypesDeleteAllParams{
		HTTPClient: client,
	}
}

/*
PerformancePriceTypesDeleteAllParams contains all the parameters to send to the API endpoint

	for the performance price types delete all operation.

	Typically these are written to a http.Request.
*/
type PerformancePriceTypesDeleteAllParams struct {

	// PerformanceIds.
	PerformanceIds string

	// PriceLayerTypeIds.
	PriceLayerTypeIds *string

	// PriceTypeIds.
	PriceTypeIds *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the performance price types delete all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformancePriceTypesDeleteAllParams) WithDefaults() *PerformancePriceTypesDeleteAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the performance price types delete all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformancePriceTypesDeleteAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the performance price types delete all params
func (o *PerformancePriceTypesDeleteAllParams) WithTimeout(timeout time.Duration) *PerformancePriceTypesDeleteAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the performance price types delete all params
func (o *PerformancePriceTypesDeleteAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the performance price types delete all params
func (o *PerformancePriceTypesDeleteAllParams) WithContext(ctx context.Context) *PerformancePriceTypesDeleteAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the performance price types delete all params
func (o *PerformancePriceTypesDeleteAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the performance price types delete all params
func (o *PerformancePriceTypesDeleteAllParams) WithHTTPClient(client *http.Client) *PerformancePriceTypesDeleteAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the performance price types delete all params
func (o *PerformancePriceTypesDeleteAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPerformanceIds adds the performanceIds to the performance price types delete all params
func (o *PerformancePriceTypesDeleteAllParams) WithPerformanceIds(performanceIds string) *PerformancePriceTypesDeleteAllParams {
	o.SetPerformanceIds(performanceIds)
	return o
}

// SetPerformanceIds adds the performanceIds to the performance price types delete all params
func (o *PerformancePriceTypesDeleteAllParams) SetPerformanceIds(performanceIds string) {
	o.PerformanceIds = performanceIds
}

// WithPriceLayerTypeIds adds the priceLayerTypeIds to the performance price types delete all params
func (o *PerformancePriceTypesDeleteAllParams) WithPriceLayerTypeIds(priceLayerTypeIds *string) *PerformancePriceTypesDeleteAllParams {
	o.SetPriceLayerTypeIds(priceLayerTypeIds)
	return o
}

// SetPriceLayerTypeIds adds the priceLayerTypeIds to the performance price types delete all params
func (o *PerformancePriceTypesDeleteAllParams) SetPriceLayerTypeIds(priceLayerTypeIds *string) {
	o.PriceLayerTypeIds = priceLayerTypeIds
}

// WithPriceTypeIds adds the priceTypeIds to the performance price types delete all params
func (o *PerformancePriceTypesDeleteAllParams) WithPriceTypeIds(priceTypeIds *string) *PerformancePriceTypesDeleteAllParams {
	o.SetPriceTypeIds(priceTypeIds)
	return o
}

// SetPriceTypeIds adds the priceTypeIds to the performance price types delete all params
func (o *PerformancePriceTypesDeleteAllParams) SetPriceTypeIds(priceTypeIds *string) {
	o.PriceTypeIds = priceTypeIds
}

// WriteToRequest writes these params to a swagger request
func (o *PerformancePriceTypesDeleteAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param performanceIds
	qrPerformanceIds := o.PerformanceIds
	qPerformanceIds := qrPerformanceIds
	if qPerformanceIds != "" {

		if err := r.SetQueryParam("performanceIds", qPerformanceIds); err != nil {
			return err
		}
	}

	if o.PriceLayerTypeIds != nil {

		// query param priceLayerTypeIds
		var qrPriceLayerTypeIds string

		if o.PriceLayerTypeIds != nil {
			qrPriceLayerTypeIds = *o.PriceLayerTypeIds
		}
		qPriceLayerTypeIds := qrPriceLayerTypeIds
		if qPriceLayerTypeIds != "" {

			if err := r.SetQueryParam("priceLayerTypeIds", qPriceLayerTypeIds); err != nil {
				return err
			}
		}
	}

	if o.PriceTypeIds != nil {

		// query param priceTypeIds
		var qrPriceTypeIds string

		if o.PriceTypeIds != nil {
			qrPriceTypeIds = *o.PriceTypeIds
		}
		qPriceTypeIds := qrPriceTypeIds
		if qPriceTypeIds != "" {

			if err := r.SetQueryParam("priceTypeIds", qPriceTypeIds); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
