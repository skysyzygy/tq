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

// NewTimeSlotsCreateParams creates a new TimeSlotsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTimeSlotsCreateParams() *TimeSlotsCreateParams {
	return &TimeSlotsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTimeSlotsCreateParamsWithTimeout creates a new TimeSlotsCreateParams object
// with the ability to set a timeout on a request.
func NewTimeSlotsCreateParamsWithTimeout(timeout time.Duration) *TimeSlotsCreateParams {
	return &TimeSlotsCreateParams{
		timeout: timeout,
	}
}

// NewTimeSlotsCreateParamsWithContext creates a new TimeSlotsCreateParams object
// with the ability to set a context for a request.
func NewTimeSlotsCreateParamsWithContext(ctx context.Context) *TimeSlotsCreateParams {
	return &TimeSlotsCreateParams{
		Context: ctx,
	}
}

// NewTimeSlotsCreateParamsWithHTTPClient creates a new TimeSlotsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTimeSlotsCreateParamsWithHTTPClient(client *http.Client) *TimeSlotsCreateParams {
	return &TimeSlotsCreateParams{
		HTTPClient: client,
	}
}

/*
TimeSlotsCreateParams contains all the parameters to send to the API endpoint

	for the time slots create operation.

	Typically these are written to a http.Request.
*/
type TimeSlotsCreateParams struct {

	// Data.
	Data *models.TimeSlot

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the time slots create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TimeSlotsCreateParams) WithDefaults() *TimeSlotsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the time slots create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TimeSlotsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the time slots create params
func (o *TimeSlotsCreateParams) WithTimeout(timeout time.Duration) *TimeSlotsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the time slots create params
func (o *TimeSlotsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the time slots create params
func (o *TimeSlotsCreateParams) WithContext(ctx context.Context) *TimeSlotsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the time slots create params
func (o *TimeSlotsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the time slots create params
func (o *TimeSlotsCreateParams) WithHTTPClient(client *http.Client) *TimeSlotsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the time slots create params
func (o *TimeSlotsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the time slots create params
func (o *TimeSlotsCreateParams) WithData(data *models.TimeSlot) *TimeSlotsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the time slots create params
func (o *TimeSlotsCreateParams) SetData(data *models.TimeSlot) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *TimeSlotsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
