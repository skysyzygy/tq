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

// NewSchedulePatternTypesGetSummariesParams creates a new SchedulePatternTypesGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSchedulePatternTypesGetSummariesParams() *SchedulePatternTypesGetSummariesParams {
	return &SchedulePatternTypesGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSchedulePatternTypesGetSummariesParamsWithTimeout creates a new SchedulePatternTypesGetSummariesParams object
// with the ability to set a timeout on a request.
func NewSchedulePatternTypesGetSummariesParamsWithTimeout(timeout time.Duration) *SchedulePatternTypesGetSummariesParams {
	return &SchedulePatternTypesGetSummariesParams{
		timeout: timeout,
	}
}

// NewSchedulePatternTypesGetSummariesParamsWithContext creates a new SchedulePatternTypesGetSummariesParams object
// with the ability to set a context for a request.
func NewSchedulePatternTypesGetSummariesParamsWithContext(ctx context.Context) *SchedulePatternTypesGetSummariesParams {
	return &SchedulePatternTypesGetSummariesParams{
		Context: ctx,
	}
}

// NewSchedulePatternTypesGetSummariesParamsWithHTTPClient creates a new SchedulePatternTypesGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewSchedulePatternTypesGetSummariesParamsWithHTTPClient(client *http.Client) *SchedulePatternTypesGetSummariesParams {
	return &SchedulePatternTypesGetSummariesParams{
		HTTPClient: client,
	}
}

/*
SchedulePatternTypesGetSummariesParams contains all the parameters to send to the API endpoint

	for the schedule pattern types get summaries operation.

	Typically these are written to a http.Request.
*/
type SchedulePatternTypesGetSummariesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the schedule pattern types get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchedulePatternTypesGetSummariesParams) WithDefaults() *SchedulePatternTypesGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the schedule pattern types get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchedulePatternTypesGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the schedule pattern types get summaries params
func (o *SchedulePatternTypesGetSummariesParams) WithTimeout(timeout time.Duration) *SchedulePatternTypesGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schedule pattern types get summaries params
func (o *SchedulePatternTypesGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schedule pattern types get summaries params
func (o *SchedulePatternTypesGetSummariesParams) WithContext(ctx context.Context) *SchedulePatternTypesGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schedule pattern types get summaries params
func (o *SchedulePatternTypesGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schedule pattern types get summaries params
func (o *SchedulePatternTypesGetSummariesParams) WithHTTPClient(client *http.Client) *SchedulePatternTypesGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schedule pattern types get summaries params
func (o *SchedulePatternTypesGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SchedulePatternTypesGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
