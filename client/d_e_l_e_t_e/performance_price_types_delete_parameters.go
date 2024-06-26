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

// NewPerformancePriceTypesDeleteParams creates a new PerformancePriceTypesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformancePriceTypesDeleteParams() *PerformancePriceTypesDeleteParams {
	return &PerformancePriceTypesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformancePriceTypesDeleteParamsWithTimeout creates a new PerformancePriceTypesDeleteParams object
// with the ability to set a timeout on a request.
func NewPerformancePriceTypesDeleteParamsWithTimeout(timeout time.Duration) *PerformancePriceTypesDeleteParams {
	return &PerformancePriceTypesDeleteParams{
		timeout: timeout,
	}
}

// NewPerformancePriceTypesDeleteParamsWithContext creates a new PerformancePriceTypesDeleteParams object
// with the ability to set a context for a request.
func NewPerformancePriceTypesDeleteParamsWithContext(ctx context.Context) *PerformancePriceTypesDeleteParams {
	return &PerformancePriceTypesDeleteParams{
		Context: ctx,
	}
}

// NewPerformancePriceTypesDeleteParamsWithHTTPClient creates a new PerformancePriceTypesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformancePriceTypesDeleteParamsWithHTTPClient(client *http.Client) *PerformancePriceTypesDeleteParams {
	return &PerformancePriceTypesDeleteParams{
		HTTPClient: client,
	}
}

/*
PerformancePriceTypesDeleteParams contains all the parameters to send to the API endpoint

	for the performance price types delete operation.

	Typically these are written to a http.Request.
*/
type PerformancePriceTypesDeleteParams struct {

	// PerformancePriceTypeID.
	PerformancePriceTypeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the performance price types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformancePriceTypesDeleteParams) WithDefaults() *PerformancePriceTypesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the performance price types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformancePriceTypesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the performance price types delete params
func (o *PerformancePriceTypesDeleteParams) WithTimeout(timeout time.Duration) *PerformancePriceTypesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the performance price types delete params
func (o *PerformancePriceTypesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the performance price types delete params
func (o *PerformancePriceTypesDeleteParams) WithContext(ctx context.Context) *PerformancePriceTypesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the performance price types delete params
func (o *PerformancePriceTypesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the performance price types delete params
func (o *PerformancePriceTypesDeleteParams) WithHTTPClient(client *http.Client) *PerformancePriceTypesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the performance price types delete params
func (o *PerformancePriceTypesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPerformancePriceTypeID adds the performancePriceTypeID to the performance price types delete params
func (o *PerformancePriceTypesDeleteParams) WithPerformancePriceTypeID(performancePriceTypeID string) *PerformancePriceTypesDeleteParams {
	o.SetPerformancePriceTypeID(performancePriceTypeID)
	return o
}

// SetPerformancePriceTypeID adds the performancePriceTypeId to the performance price types delete params
func (o *PerformancePriceTypesDeleteParams) SetPerformancePriceTypeID(performancePriceTypeID string) {
	o.PerformancePriceTypeID = performancePriceTypeID
}

// WriteToRequest writes these params to a swagger request
func (o *PerformancePriceTypesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param performancePriceTypeId
	if err := r.SetPathParam("performancePriceTypeId", o.PerformancePriceTypeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
