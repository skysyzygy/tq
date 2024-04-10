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

// NewTimeSlotsGetSummariesParams creates a new TimeSlotsGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTimeSlotsGetSummariesParams() *TimeSlotsGetSummariesParams {
	return &TimeSlotsGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTimeSlotsGetSummariesParamsWithTimeout creates a new TimeSlotsGetSummariesParams object
// with the ability to set a timeout on a request.
func NewTimeSlotsGetSummariesParamsWithTimeout(timeout time.Duration) *TimeSlotsGetSummariesParams {
	return &TimeSlotsGetSummariesParams{
		timeout: timeout,
	}
}

// NewTimeSlotsGetSummariesParamsWithContext creates a new TimeSlotsGetSummariesParams object
// with the ability to set a context for a request.
func NewTimeSlotsGetSummariesParamsWithContext(ctx context.Context) *TimeSlotsGetSummariesParams {
	return &TimeSlotsGetSummariesParams{
		Context: ctx,
	}
}

// NewTimeSlotsGetSummariesParamsWithHTTPClient creates a new TimeSlotsGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewTimeSlotsGetSummariesParamsWithHTTPClient(client *http.Client) *TimeSlotsGetSummariesParams {
	return &TimeSlotsGetSummariesParams{
		HTTPClient: client,
	}
}

/*
TimeSlotsGetSummariesParams contains all the parameters to send to the API endpoint

	for the time slots get summaries operation.

	Typically these are written to a http.Request.
*/
type TimeSlotsGetSummariesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the time slots get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TimeSlotsGetSummariesParams) WithDefaults() *TimeSlotsGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the time slots get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TimeSlotsGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the time slots get summaries params
func (o *TimeSlotsGetSummariesParams) WithTimeout(timeout time.Duration) *TimeSlotsGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the time slots get summaries params
func (o *TimeSlotsGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the time slots get summaries params
func (o *TimeSlotsGetSummariesParams) WithContext(ctx context.Context) *TimeSlotsGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the time slots get summaries params
func (o *TimeSlotsGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the time slots get summaries params
func (o *TimeSlotsGetSummariesParams) WithHTTPClient(client *http.Client) *TimeSlotsGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the time slots get summaries params
func (o *TimeSlotsGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *TimeSlotsGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
