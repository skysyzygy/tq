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

// NewReportSchedulesGetAllParams creates a new ReportSchedulesGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReportSchedulesGetAllParams() *ReportSchedulesGetAllParams {
	return &ReportSchedulesGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReportSchedulesGetAllParamsWithTimeout creates a new ReportSchedulesGetAllParams object
// with the ability to set a timeout on a request.
func NewReportSchedulesGetAllParamsWithTimeout(timeout time.Duration) *ReportSchedulesGetAllParams {
	return &ReportSchedulesGetAllParams{
		timeout: timeout,
	}
}

// NewReportSchedulesGetAllParamsWithContext creates a new ReportSchedulesGetAllParams object
// with the ability to set a context for a request.
func NewReportSchedulesGetAllParamsWithContext(ctx context.Context) *ReportSchedulesGetAllParams {
	return &ReportSchedulesGetAllParams{
		Context: ctx,
	}
}

// NewReportSchedulesGetAllParamsWithHTTPClient creates a new ReportSchedulesGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewReportSchedulesGetAllParamsWithHTTPClient(client *http.Client) *ReportSchedulesGetAllParams {
	return &ReportSchedulesGetAllParams{
		HTTPClient: client,
	}
}

/*
ReportSchedulesGetAllParams contains all the parameters to send to the API endpoint

	for the report schedules get all operation.

	Typically these are written to a http.Request.
*/
type ReportSchedulesGetAllParams struct {

	// ActiveOnly.
	ActiveOnly *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the report schedules get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReportSchedulesGetAllParams) WithDefaults() *ReportSchedulesGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the report schedules get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReportSchedulesGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the report schedules get all params
func (o *ReportSchedulesGetAllParams) WithTimeout(timeout time.Duration) *ReportSchedulesGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the report schedules get all params
func (o *ReportSchedulesGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the report schedules get all params
func (o *ReportSchedulesGetAllParams) WithContext(ctx context.Context) *ReportSchedulesGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the report schedules get all params
func (o *ReportSchedulesGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the report schedules get all params
func (o *ReportSchedulesGetAllParams) WithHTTPClient(client *http.Client) *ReportSchedulesGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the report schedules get all params
func (o *ReportSchedulesGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActiveOnly adds the activeOnly to the report schedules get all params
func (o *ReportSchedulesGetAllParams) WithActiveOnly(activeOnly *string) *ReportSchedulesGetAllParams {
	o.SetActiveOnly(activeOnly)
	return o
}

// SetActiveOnly adds the activeOnly to the report schedules get all params
func (o *ReportSchedulesGetAllParams) SetActiveOnly(activeOnly *string) {
	o.ActiveOnly = activeOnly
}

// WriteToRequest writes these params to a swagger request
func (o *ReportSchedulesGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ActiveOnly != nil {

		// query param activeOnly
		var qrActiveOnly string

		if o.ActiveOnly != nil {
			qrActiveOnly = *o.ActiveOnly
		}
		qActiveOnly := qrActiveOnly
		if qActiveOnly != "" {

			if err := r.SetQueryParam("activeOnly", qActiveOnly); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
