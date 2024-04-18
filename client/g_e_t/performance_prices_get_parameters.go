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

// NewPerformancePricesGetParams creates a new PerformancePricesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformancePricesGetParams() *PerformancePricesGetParams {
	return &PerformancePricesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformancePricesGetParamsWithTimeout creates a new PerformancePricesGetParams object
// with the ability to set a timeout on a request.
func NewPerformancePricesGetParamsWithTimeout(timeout time.Duration) *PerformancePricesGetParams {
	return &PerformancePricesGetParams{
		timeout: timeout,
	}
}

// NewPerformancePricesGetParamsWithContext creates a new PerformancePricesGetParams object
// with the ability to set a context for a request.
func NewPerformancePricesGetParamsWithContext(ctx context.Context) *PerformancePricesGetParams {
	return &PerformancePricesGetParams{
		Context: ctx,
	}
}

// NewPerformancePricesGetParamsWithHTTPClient creates a new PerformancePricesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformancePricesGetParamsWithHTTPClient(client *http.Client) *PerformancePricesGetParams {
	return &PerformancePricesGetParams{
		HTTPClient: client,
	}
}

/*
PerformancePricesGetParams contains all the parameters to send to the API endpoint

	for the performance prices get operation.

	Typically these are written to a http.Request.
*/
type PerformancePricesGetParams struct {

	// AsOfDateTime.
	AsOfDateTime *string

	// PerformancePriceID.
	PerformancePriceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the performance prices get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformancePricesGetParams) WithDefaults() *PerformancePricesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the performance prices get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformancePricesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the performance prices get params
func (o *PerformancePricesGetParams) WithTimeout(timeout time.Duration) *PerformancePricesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the performance prices get params
func (o *PerformancePricesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the performance prices get params
func (o *PerformancePricesGetParams) WithContext(ctx context.Context) *PerformancePricesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the performance prices get params
func (o *PerformancePricesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the performance prices get params
func (o *PerformancePricesGetParams) WithHTTPClient(client *http.Client) *PerformancePricesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the performance prices get params
func (o *PerformancePricesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAsOfDateTime adds the asOfDateTime to the performance prices get params
func (o *PerformancePricesGetParams) WithAsOfDateTime(asOfDateTime *string) *PerformancePricesGetParams {
	o.SetAsOfDateTime(asOfDateTime)
	return o
}

// SetAsOfDateTime adds the asOfDateTime to the performance prices get params
func (o *PerformancePricesGetParams) SetAsOfDateTime(asOfDateTime *string) {
	o.AsOfDateTime = asOfDateTime
}

// WithPerformancePriceID adds the performancePriceID to the performance prices get params
func (o *PerformancePricesGetParams) WithPerformancePriceID(performancePriceID string) *PerformancePricesGetParams {
	o.SetPerformancePriceID(performancePriceID)
	return o
}

// SetPerformancePriceID adds the performancePriceId to the performance prices get params
func (o *PerformancePricesGetParams) SetPerformancePriceID(performancePriceID string) {
	o.PerformancePriceID = performancePriceID
}

// WriteToRequest writes these params to a swagger request
func (o *PerformancePricesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AsOfDateTime != nil {

		// query param asOfDateTime
		var qrAsOfDateTime string

		if o.AsOfDateTime != nil {
			qrAsOfDateTime = *o.AsOfDateTime
		}
		qAsOfDateTime := qrAsOfDateTime
		if qAsOfDateTime != "" {

			if err := r.SetQueryParam("asOfDateTime", qAsOfDateTime); err != nil {
				return err
			}
		}
	}

	// path param performancePriceId
	if err := r.SetPathParam("performancePriceId", o.PerformancePriceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}