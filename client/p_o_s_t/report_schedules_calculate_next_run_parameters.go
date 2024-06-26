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

// NewReportSchedulesCalculateNextRunParams creates a new ReportSchedulesCalculateNextRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReportSchedulesCalculateNextRunParams() *ReportSchedulesCalculateNextRunParams {
	return &ReportSchedulesCalculateNextRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReportSchedulesCalculateNextRunParamsWithTimeout creates a new ReportSchedulesCalculateNextRunParams object
// with the ability to set a timeout on a request.
func NewReportSchedulesCalculateNextRunParamsWithTimeout(timeout time.Duration) *ReportSchedulesCalculateNextRunParams {
	return &ReportSchedulesCalculateNextRunParams{
		timeout: timeout,
	}
}

// NewReportSchedulesCalculateNextRunParamsWithContext creates a new ReportSchedulesCalculateNextRunParams object
// with the ability to set a context for a request.
func NewReportSchedulesCalculateNextRunParamsWithContext(ctx context.Context) *ReportSchedulesCalculateNextRunParams {
	return &ReportSchedulesCalculateNextRunParams{
		Context: ctx,
	}
}

// NewReportSchedulesCalculateNextRunParamsWithHTTPClient creates a new ReportSchedulesCalculateNextRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewReportSchedulesCalculateNextRunParamsWithHTTPClient(client *http.Client) *ReportSchedulesCalculateNextRunParams {
	return &ReportSchedulesCalculateNextRunParams{
		HTTPClient: client,
	}
}

/*
ReportSchedulesCalculateNextRunParams contains all the parameters to send to the API endpoint

	for the report schedules calculate next run operation.

	Typically these are written to a http.Request.
*/
type ReportSchedulesCalculateNextRunParams struct {

	// ReportScheduleNextRequest.
	ReportScheduleNextRequest *models.ReportScheduleNextRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the report schedules calculate next run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReportSchedulesCalculateNextRunParams) WithDefaults() *ReportSchedulesCalculateNextRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the report schedules calculate next run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReportSchedulesCalculateNextRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the report schedules calculate next run params
func (o *ReportSchedulesCalculateNextRunParams) WithTimeout(timeout time.Duration) *ReportSchedulesCalculateNextRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the report schedules calculate next run params
func (o *ReportSchedulesCalculateNextRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the report schedules calculate next run params
func (o *ReportSchedulesCalculateNextRunParams) WithContext(ctx context.Context) *ReportSchedulesCalculateNextRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the report schedules calculate next run params
func (o *ReportSchedulesCalculateNextRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the report schedules calculate next run params
func (o *ReportSchedulesCalculateNextRunParams) WithHTTPClient(client *http.Client) *ReportSchedulesCalculateNextRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the report schedules calculate next run params
func (o *ReportSchedulesCalculateNextRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReportScheduleNextRequest adds the reportScheduleNextRequest to the report schedules calculate next run params
func (o *ReportSchedulesCalculateNextRunParams) WithReportScheduleNextRequest(reportScheduleNextRequest *models.ReportScheduleNextRequest) *ReportSchedulesCalculateNextRunParams {
	o.SetReportScheduleNextRequest(reportScheduleNextRequest)
	return o
}

// SetReportScheduleNextRequest adds the reportScheduleNextRequest to the report schedules calculate next run params
func (o *ReportSchedulesCalculateNextRunParams) SetReportScheduleNextRequest(reportScheduleNextRequest *models.ReportScheduleNextRequest) {
	o.ReportScheduleNextRequest = reportScheduleNextRequest
}

// WriteToRequest writes these params to a swagger request
func (o *ReportSchedulesCalculateNextRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ReportScheduleNextRequest != nil {
		if err := r.SetBodyParam(o.ReportScheduleNextRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
