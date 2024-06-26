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

// NewPerformancePriceTypesGetParams creates a new PerformancePriceTypesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformancePriceTypesGetParams() *PerformancePriceTypesGetParams {
	return &PerformancePriceTypesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformancePriceTypesGetParamsWithTimeout creates a new PerformancePriceTypesGetParams object
// with the ability to set a timeout on a request.
func NewPerformancePriceTypesGetParamsWithTimeout(timeout time.Duration) *PerformancePriceTypesGetParams {
	return &PerformancePriceTypesGetParams{
		timeout: timeout,
	}
}

// NewPerformancePriceTypesGetParamsWithContext creates a new PerformancePriceTypesGetParams object
// with the ability to set a context for a request.
func NewPerformancePriceTypesGetParamsWithContext(ctx context.Context) *PerformancePriceTypesGetParams {
	return &PerformancePriceTypesGetParams{
		Context: ctx,
	}
}

// NewPerformancePriceTypesGetParamsWithHTTPClient creates a new PerformancePriceTypesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformancePriceTypesGetParamsWithHTTPClient(client *http.Client) *PerformancePriceTypesGetParams {
	return &PerformancePriceTypesGetParams{
		HTTPClient: client,
	}
}

/*
PerformancePriceTypesGetParams contains all the parameters to send to the API endpoint

	for the performance price types get operation.

	Typically these are written to a http.Request.
*/
type PerformancePriceTypesGetParams struct {

	// AsOfDateTime.
	AsOfDateTime *string

	// PerformancePriceTypeID.
	PerformancePriceTypeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the performance price types get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformancePriceTypesGetParams) WithDefaults() *PerformancePriceTypesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the performance price types get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformancePriceTypesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the performance price types get params
func (o *PerformancePriceTypesGetParams) WithTimeout(timeout time.Duration) *PerformancePriceTypesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the performance price types get params
func (o *PerformancePriceTypesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the performance price types get params
func (o *PerformancePriceTypesGetParams) WithContext(ctx context.Context) *PerformancePriceTypesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the performance price types get params
func (o *PerformancePriceTypesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the performance price types get params
func (o *PerformancePriceTypesGetParams) WithHTTPClient(client *http.Client) *PerformancePriceTypesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the performance price types get params
func (o *PerformancePriceTypesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAsOfDateTime adds the asOfDateTime to the performance price types get params
func (o *PerformancePriceTypesGetParams) WithAsOfDateTime(asOfDateTime *string) *PerformancePriceTypesGetParams {
	o.SetAsOfDateTime(asOfDateTime)
	return o
}

// SetAsOfDateTime adds the asOfDateTime to the performance price types get params
func (o *PerformancePriceTypesGetParams) SetAsOfDateTime(asOfDateTime *string) {
	o.AsOfDateTime = asOfDateTime
}

// WithPerformancePriceTypeID adds the performancePriceTypeID to the performance price types get params
func (o *PerformancePriceTypesGetParams) WithPerformancePriceTypeID(performancePriceTypeID string) *PerformancePriceTypesGetParams {
	o.SetPerformancePriceTypeID(performancePriceTypeID)
	return o
}

// SetPerformancePriceTypeID adds the performancePriceTypeId to the performance price types get params
func (o *PerformancePriceTypesGetParams) SetPerformancePriceTypeID(performancePriceTypeID string) {
	o.PerformancePriceTypeID = performancePriceTypeID
}

// WriteToRequest writes these params to a swagger request
func (o *PerformancePriceTypesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param performancePriceTypeId
	if err := r.SetPathParam("performancePriceTypeId", o.PerformancePriceTypeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
