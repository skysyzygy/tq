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

// NewReportUserGroupsGetAllParams creates a new ReportUserGroupsGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReportUserGroupsGetAllParams() *ReportUserGroupsGetAllParams {
	return &ReportUserGroupsGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReportUserGroupsGetAllParamsWithTimeout creates a new ReportUserGroupsGetAllParams object
// with the ability to set a timeout on a request.
func NewReportUserGroupsGetAllParamsWithTimeout(timeout time.Duration) *ReportUserGroupsGetAllParams {
	return &ReportUserGroupsGetAllParams{
		timeout: timeout,
	}
}

// NewReportUserGroupsGetAllParamsWithContext creates a new ReportUserGroupsGetAllParams object
// with the ability to set a context for a request.
func NewReportUserGroupsGetAllParamsWithContext(ctx context.Context) *ReportUserGroupsGetAllParams {
	return &ReportUserGroupsGetAllParams{
		Context: ctx,
	}
}

// NewReportUserGroupsGetAllParamsWithHTTPClient creates a new ReportUserGroupsGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewReportUserGroupsGetAllParamsWithHTTPClient(client *http.Client) *ReportUserGroupsGetAllParams {
	return &ReportUserGroupsGetAllParams{
		HTTPClient: client,
	}
}

/*
ReportUserGroupsGetAllParams contains all the parameters to send to the API endpoint

	for the report user groups get all operation.

	Typically these are written to a http.Request.
*/
type ReportUserGroupsGetAllParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the report user groups get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReportUserGroupsGetAllParams) WithDefaults() *ReportUserGroupsGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the report user groups get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReportUserGroupsGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the report user groups get all params
func (o *ReportUserGroupsGetAllParams) WithTimeout(timeout time.Duration) *ReportUserGroupsGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the report user groups get all params
func (o *ReportUserGroupsGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the report user groups get all params
func (o *ReportUserGroupsGetAllParams) WithContext(ctx context.Context) *ReportUserGroupsGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the report user groups get all params
func (o *ReportUserGroupsGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the report user groups get all params
func (o *ReportUserGroupsGetAllParams) WithHTTPClient(client *http.Client) *ReportUserGroupsGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the report user groups get all params
func (o *ReportUserGroupsGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ReportUserGroupsGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
