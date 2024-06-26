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

// NewPerformancePricesDeleteParams creates a new PerformancePricesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformancePricesDeleteParams() *PerformancePricesDeleteParams {
	return &PerformancePricesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformancePricesDeleteParamsWithTimeout creates a new PerformancePricesDeleteParams object
// with the ability to set a timeout on a request.
func NewPerformancePricesDeleteParamsWithTimeout(timeout time.Duration) *PerformancePricesDeleteParams {
	return &PerformancePricesDeleteParams{
		timeout: timeout,
	}
}

// NewPerformancePricesDeleteParamsWithContext creates a new PerformancePricesDeleteParams object
// with the ability to set a context for a request.
func NewPerformancePricesDeleteParamsWithContext(ctx context.Context) *PerformancePricesDeleteParams {
	return &PerformancePricesDeleteParams{
		Context: ctx,
	}
}

// NewPerformancePricesDeleteParamsWithHTTPClient creates a new PerformancePricesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformancePricesDeleteParamsWithHTTPClient(client *http.Client) *PerformancePricesDeleteParams {
	return &PerformancePricesDeleteParams{
		HTTPClient: client,
	}
}

/*
PerformancePricesDeleteParams contains all the parameters to send to the API endpoint

	for the performance prices delete operation.

	Typically these are written to a http.Request.
*/
type PerformancePricesDeleteParams struct {

	// PerformancePriceID.
	PerformancePriceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the performance prices delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformancePricesDeleteParams) WithDefaults() *PerformancePricesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the performance prices delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformancePricesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the performance prices delete params
func (o *PerformancePricesDeleteParams) WithTimeout(timeout time.Duration) *PerformancePricesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the performance prices delete params
func (o *PerformancePricesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the performance prices delete params
func (o *PerformancePricesDeleteParams) WithContext(ctx context.Context) *PerformancePricesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the performance prices delete params
func (o *PerformancePricesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the performance prices delete params
func (o *PerformancePricesDeleteParams) WithHTTPClient(client *http.Client) *PerformancePricesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the performance prices delete params
func (o *PerformancePricesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPerformancePriceID adds the performancePriceID to the performance prices delete params
func (o *PerformancePricesDeleteParams) WithPerformancePriceID(performancePriceID string) *PerformancePricesDeleteParams {
	o.SetPerformancePriceID(performancePriceID)
	return o
}

// SetPerformancePriceID adds the performancePriceId to the performance prices delete params
func (o *PerformancePricesDeleteParams) SetPerformancePriceID(performancePriceID string) {
	o.PerformancePriceID = performancePriceID
}

// WriteToRequest writes these params to a swagger request
func (o *PerformancePricesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param performancePriceId
	if err := r.SetPathParam("performancePriceId", o.PerformancePriceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
