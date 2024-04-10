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

// NewReportRequestsGetAllParams creates a new ReportRequestsGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReportRequestsGetAllParams() *ReportRequestsGetAllParams {
	return &ReportRequestsGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReportRequestsGetAllParamsWithTimeout creates a new ReportRequestsGetAllParams object
// with the ability to set a timeout on a request.
func NewReportRequestsGetAllParamsWithTimeout(timeout time.Duration) *ReportRequestsGetAllParams {
	return &ReportRequestsGetAllParams{
		timeout: timeout,
	}
}

// NewReportRequestsGetAllParamsWithContext creates a new ReportRequestsGetAllParams object
// with the ability to set a context for a request.
func NewReportRequestsGetAllParamsWithContext(ctx context.Context) *ReportRequestsGetAllParams {
	return &ReportRequestsGetAllParams{
		Context: ctx,
	}
}

// NewReportRequestsGetAllParamsWithHTTPClient creates a new ReportRequestsGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewReportRequestsGetAllParamsWithHTTPClient(client *http.Client) *ReportRequestsGetAllParams {
	return &ReportRequestsGetAllParams{
		HTTPClient: client,
	}
}

/*
ReportRequestsGetAllParams contains all the parameters to send to the API endpoint

	for the report requests get all operation.

	Typically these are written to a http.Request.
*/
type ReportRequestsGetAllParams struct {

	// ActiveOnly.
	ActiveOnly *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the report requests get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReportRequestsGetAllParams) WithDefaults() *ReportRequestsGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the report requests get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReportRequestsGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the report requests get all params
func (o *ReportRequestsGetAllParams) WithTimeout(timeout time.Duration) *ReportRequestsGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the report requests get all params
func (o *ReportRequestsGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the report requests get all params
func (o *ReportRequestsGetAllParams) WithContext(ctx context.Context) *ReportRequestsGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the report requests get all params
func (o *ReportRequestsGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the report requests get all params
func (o *ReportRequestsGetAllParams) WithHTTPClient(client *http.Client) *ReportRequestsGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the report requests get all params
func (o *ReportRequestsGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActiveOnly adds the activeOnly to the report requests get all params
func (o *ReportRequestsGetAllParams) WithActiveOnly(activeOnly *string) *ReportRequestsGetAllParams {
	o.SetActiveOnly(activeOnly)
	return o
}

// SetActiveOnly adds the activeOnly to the report requests get all params
func (o *ReportRequestsGetAllParams) SetActiveOnly(activeOnly *string) {
	o.ActiveOnly = activeOnly
}

// WriteToRequest writes these params to a swagger request
func (o *ReportRequestsGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
