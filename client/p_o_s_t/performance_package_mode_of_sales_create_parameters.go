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

// NewPerformancePackageModeOfSalesCreateParams creates a new PerformancePackageModeOfSalesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformancePackageModeOfSalesCreateParams() *PerformancePackageModeOfSalesCreateParams {
	return &PerformancePackageModeOfSalesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformancePackageModeOfSalesCreateParamsWithTimeout creates a new PerformancePackageModeOfSalesCreateParams object
// with the ability to set a timeout on a request.
func NewPerformancePackageModeOfSalesCreateParamsWithTimeout(timeout time.Duration) *PerformancePackageModeOfSalesCreateParams {
	return &PerformancePackageModeOfSalesCreateParams{
		timeout: timeout,
	}
}

// NewPerformancePackageModeOfSalesCreateParamsWithContext creates a new PerformancePackageModeOfSalesCreateParams object
// with the ability to set a context for a request.
func NewPerformancePackageModeOfSalesCreateParamsWithContext(ctx context.Context) *PerformancePackageModeOfSalesCreateParams {
	return &PerformancePackageModeOfSalesCreateParams{
		Context: ctx,
	}
}

// NewPerformancePackageModeOfSalesCreateParamsWithHTTPClient creates a new PerformancePackageModeOfSalesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformancePackageModeOfSalesCreateParamsWithHTTPClient(client *http.Client) *PerformancePackageModeOfSalesCreateParams {
	return &PerformancePackageModeOfSalesCreateParams{
		HTTPClient: client,
	}
}

/*
PerformancePackageModeOfSalesCreateParams contains all the parameters to send to the API endpoint

	for the performance package mode of sales create operation.

	Typically these are written to a http.Request.
*/
type PerformancePackageModeOfSalesCreateParams struct {

	// PerformancePackageModeOfSale.
	PerformancePackageModeOfSale *models.PerformancePackageModeOfSale

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the performance package mode of sales create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformancePackageModeOfSalesCreateParams) WithDefaults() *PerformancePackageModeOfSalesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the performance package mode of sales create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformancePackageModeOfSalesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the performance package mode of sales create params
func (o *PerformancePackageModeOfSalesCreateParams) WithTimeout(timeout time.Duration) *PerformancePackageModeOfSalesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the performance package mode of sales create params
func (o *PerformancePackageModeOfSalesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the performance package mode of sales create params
func (o *PerformancePackageModeOfSalesCreateParams) WithContext(ctx context.Context) *PerformancePackageModeOfSalesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the performance package mode of sales create params
func (o *PerformancePackageModeOfSalesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the performance package mode of sales create params
func (o *PerformancePackageModeOfSalesCreateParams) WithHTTPClient(client *http.Client) *PerformancePackageModeOfSalesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the performance package mode of sales create params
func (o *PerformancePackageModeOfSalesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPerformancePackageModeOfSale adds the performancePackageModeOfSale to the performance package mode of sales create params
func (o *PerformancePackageModeOfSalesCreateParams) WithPerformancePackageModeOfSale(performancePackageModeOfSale *models.PerformancePackageModeOfSale) *PerformancePackageModeOfSalesCreateParams {
	o.SetPerformancePackageModeOfSale(performancePackageModeOfSale)
	return o
}

// SetPerformancePackageModeOfSale adds the performancePackageModeOfSale to the performance package mode of sales create params
func (o *PerformancePackageModeOfSalesCreateParams) SetPerformancePackageModeOfSale(performancePackageModeOfSale *models.PerformancePackageModeOfSale) {
	o.PerformancePackageModeOfSale = performancePackageModeOfSale
}

// WriteToRequest writes these params to a swagger request
func (o *PerformancePackageModeOfSalesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PerformancePackageModeOfSale != nil {
		if err := r.SetBodyParam(o.PerformancePackageModeOfSale); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
