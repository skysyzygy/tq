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

// NewPerformancePackageModeOfSalesGetParams creates a new PerformancePackageModeOfSalesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformancePackageModeOfSalesGetParams() *PerformancePackageModeOfSalesGetParams {
	return &PerformancePackageModeOfSalesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformancePackageModeOfSalesGetParamsWithTimeout creates a new PerformancePackageModeOfSalesGetParams object
// with the ability to set a timeout on a request.
func NewPerformancePackageModeOfSalesGetParamsWithTimeout(timeout time.Duration) *PerformancePackageModeOfSalesGetParams {
	return &PerformancePackageModeOfSalesGetParams{
		timeout: timeout,
	}
}

// NewPerformancePackageModeOfSalesGetParamsWithContext creates a new PerformancePackageModeOfSalesGetParams object
// with the ability to set a context for a request.
func NewPerformancePackageModeOfSalesGetParamsWithContext(ctx context.Context) *PerformancePackageModeOfSalesGetParams {
	return &PerformancePackageModeOfSalesGetParams{
		Context: ctx,
	}
}

// NewPerformancePackageModeOfSalesGetParamsWithHTTPClient creates a new PerformancePackageModeOfSalesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformancePackageModeOfSalesGetParamsWithHTTPClient(client *http.Client) *PerformancePackageModeOfSalesGetParams {
	return &PerformancePackageModeOfSalesGetParams{
		HTTPClient: client,
	}
}

/*
PerformancePackageModeOfSalesGetParams contains all the parameters to send to the API endpoint

	for the performance package mode of sales get operation.

	Typically these are written to a http.Request.
*/
type PerformancePackageModeOfSalesGetParams struct {

	// PerformancePackageModeOfSaleID.
	PerformancePackageModeOfSaleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the performance package mode of sales get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformancePackageModeOfSalesGetParams) WithDefaults() *PerformancePackageModeOfSalesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the performance package mode of sales get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformancePackageModeOfSalesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the performance package mode of sales get params
func (o *PerformancePackageModeOfSalesGetParams) WithTimeout(timeout time.Duration) *PerformancePackageModeOfSalesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the performance package mode of sales get params
func (o *PerformancePackageModeOfSalesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the performance package mode of sales get params
func (o *PerformancePackageModeOfSalesGetParams) WithContext(ctx context.Context) *PerformancePackageModeOfSalesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the performance package mode of sales get params
func (o *PerformancePackageModeOfSalesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the performance package mode of sales get params
func (o *PerformancePackageModeOfSalesGetParams) WithHTTPClient(client *http.Client) *PerformancePackageModeOfSalesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the performance package mode of sales get params
func (o *PerformancePackageModeOfSalesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPerformancePackageModeOfSaleID adds the performancePackageModeOfSaleID to the performance package mode of sales get params
func (o *PerformancePackageModeOfSalesGetParams) WithPerformancePackageModeOfSaleID(performancePackageModeOfSaleID string) *PerformancePackageModeOfSalesGetParams {
	o.SetPerformancePackageModeOfSaleID(performancePackageModeOfSaleID)
	return o
}

// SetPerformancePackageModeOfSaleID adds the performancePackageModeOfSaleId to the performance package mode of sales get params
func (o *PerformancePackageModeOfSalesGetParams) SetPerformancePackageModeOfSaleID(performancePackageModeOfSaleID string) {
	o.PerformancePackageModeOfSaleID = performancePackageModeOfSaleID
}

// WriteToRequest writes these params to a swagger request
func (o *PerformancePackageModeOfSalesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param performancePackageModeOfSaleId
	if err := r.SetPathParam("performancePackageModeOfSaleId", o.PerformancePackageModeOfSaleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}