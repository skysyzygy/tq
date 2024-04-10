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

// NewPerformanceStatusesCreateParams creates a new PerformanceStatusesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformanceStatusesCreateParams() *PerformanceStatusesCreateParams {
	return &PerformanceStatusesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformanceStatusesCreateParamsWithTimeout creates a new PerformanceStatusesCreateParams object
// with the ability to set a timeout on a request.
func NewPerformanceStatusesCreateParamsWithTimeout(timeout time.Duration) *PerformanceStatusesCreateParams {
	return &PerformanceStatusesCreateParams{
		timeout: timeout,
	}
}

// NewPerformanceStatusesCreateParamsWithContext creates a new PerformanceStatusesCreateParams object
// with the ability to set a context for a request.
func NewPerformanceStatusesCreateParamsWithContext(ctx context.Context) *PerformanceStatusesCreateParams {
	return &PerformanceStatusesCreateParams{
		Context: ctx,
	}
}

// NewPerformanceStatusesCreateParamsWithHTTPClient creates a new PerformanceStatusesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformanceStatusesCreateParamsWithHTTPClient(client *http.Client) *PerformanceStatusesCreateParams {
	return &PerformanceStatusesCreateParams{
		HTTPClient: client,
	}
}

/*
PerformanceStatusesCreateParams contains all the parameters to send to the API endpoint

	for the performance statuses create operation.

	Typically these are written to a http.Request.
*/
type PerformanceStatusesCreateParams struct {

	// Data.
	Data *models.PerformanceStatus

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the performance statuses create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformanceStatusesCreateParams) WithDefaults() *PerformanceStatusesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the performance statuses create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformanceStatusesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the performance statuses create params
func (o *PerformanceStatusesCreateParams) WithTimeout(timeout time.Duration) *PerformanceStatusesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the performance statuses create params
func (o *PerformanceStatusesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the performance statuses create params
func (o *PerformanceStatusesCreateParams) WithContext(ctx context.Context) *PerformanceStatusesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the performance statuses create params
func (o *PerformanceStatusesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the performance statuses create params
func (o *PerformanceStatusesCreateParams) WithHTTPClient(client *http.Client) *PerformanceStatusesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the performance statuses create params
func (o *PerformanceStatusesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the performance statuses create params
func (o *PerformanceStatusesCreateParams) WithData(data *models.PerformanceStatus) *PerformanceStatusesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the performance statuses create params
func (o *PerformanceStatusesCreateParams) SetData(data *models.PerformanceStatus) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PerformanceStatusesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
