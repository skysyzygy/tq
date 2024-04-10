// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

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

	"github.com/skysyzygy/tq/models"
)

// NewPerformancePricesCreateParams creates a new PerformancePricesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformancePricesCreateParams() *PerformancePricesCreateParams {
	return &PerformancePricesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformancePricesCreateParamsWithTimeout creates a new PerformancePricesCreateParams object
// with the ability to set a timeout on a request.
func NewPerformancePricesCreateParamsWithTimeout(timeout time.Duration) *PerformancePricesCreateParams {
	return &PerformancePricesCreateParams{
		timeout: timeout,
	}
}

// NewPerformancePricesCreateParamsWithContext creates a new PerformancePricesCreateParams object
// with the ability to set a context for a request.
func NewPerformancePricesCreateParamsWithContext(ctx context.Context) *PerformancePricesCreateParams {
	return &PerformancePricesCreateParams{
		Context: ctx,
	}
}

// NewPerformancePricesCreateParamsWithHTTPClient creates a new PerformancePricesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformancePricesCreateParamsWithHTTPClient(client *http.Client) *PerformancePricesCreateParams {
	return &PerformancePricesCreateParams{
		HTTPClient: client,
	}
}

/*
PerformancePricesCreateParams contains all the parameters to send to the API endpoint

	for the performance prices create operation.

	Typically these are written to a http.Request.
*/
type PerformancePricesCreateParams struct {

	// PerformancePrice.
	PerformancePrice *models.PerformancePrice

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the performance prices create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformancePricesCreateParams) WithDefaults() *PerformancePricesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the performance prices create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformancePricesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the performance prices create params
func (o *PerformancePricesCreateParams) WithTimeout(timeout time.Duration) *PerformancePricesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the performance prices create params
func (o *PerformancePricesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the performance prices create params
func (o *PerformancePricesCreateParams) WithContext(ctx context.Context) *PerformancePricesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the performance prices create params
func (o *PerformancePricesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the performance prices create params
func (o *PerformancePricesCreateParams) WithHTTPClient(client *http.Client) *PerformancePricesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the performance prices create params
func (o *PerformancePricesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPerformancePrice adds the performancePrice to the performance prices create params
func (o *PerformancePricesCreateParams) WithPerformancePrice(performancePrice *models.PerformancePrice) *PerformancePricesCreateParams {
	o.SetPerformancePrice(performancePrice)
	return o
}

// SetPerformancePrice adds the performancePrice to the performance prices create params
func (o *PerformancePricesCreateParams) SetPerformancePrice(performancePrice *models.PerformancePrice) {
	o.PerformancePrice = performancePrice
}

// WriteToRequest writes these params to a swagger request
func (o *PerformancePricesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PerformancePrice != nil {
		if err := r.SetBodyParam(o.PerformancePrice); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
