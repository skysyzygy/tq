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

// NewReportsGetReportDefaultsParams creates a new ReportsGetReportDefaultsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReportsGetReportDefaultsParams() *ReportsGetReportDefaultsParams {
	return &ReportsGetReportDefaultsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReportsGetReportDefaultsParamsWithTimeout creates a new ReportsGetReportDefaultsParams object
// with the ability to set a timeout on a request.
func NewReportsGetReportDefaultsParamsWithTimeout(timeout time.Duration) *ReportsGetReportDefaultsParams {
	return &ReportsGetReportDefaultsParams{
		timeout: timeout,
	}
}

// NewReportsGetReportDefaultsParamsWithContext creates a new ReportsGetReportDefaultsParams object
// with the ability to set a context for a request.
func NewReportsGetReportDefaultsParamsWithContext(ctx context.Context) *ReportsGetReportDefaultsParams {
	return &ReportsGetReportDefaultsParams{
		Context: ctx,
	}
}

// NewReportsGetReportDefaultsParamsWithHTTPClient creates a new ReportsGetReportDefaultsParams object
// with the ability to set a custom HTTPClient for a request.
func NewReportsGetReportDefaultsParamsWithHTTPClient(client *http.Client) *ReportsGetReportDefaultsParams {
	return &ReportsGetReportDefaultsParams{
		HTTPClient: client,
	}
}

/*
ReportsGetReportDefaultsParams contains all the parameters to send to the API endpoint

	for the reports get report defaults operation.

	Typically these are written to a http.Request.
*/
type ReportsGetReportDefaultsParams struct {

	// UserGroup.
	UserGroup string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the reports get report defaults params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReportsGetReportDefaultsParams) WithDefaults() *ReportsGetReportDefaultsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the reports get report defaults params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReportsGetReportDefaultsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the reports get report defaults params
func (o *ReportsGetReportDefaultsParams) WithTimeout(timeout time.Duration) *ReportsGetReportDefaultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reports get report defaults params
func (o *ReportsGetReportDefaultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reports get report defaults params
func (o *ReportsGetReportDefaultsParams) WithContext(ctx context.Context) *ReportsGetReportDefaultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reports get report defaults params
func (o *ReportsGetReportDefaultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reports get report defaults params
func (o *ReportsGetReportDefaultsParams) WithHTTPClient(client *http.Client) *ReportsGetReportDefaultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reports get report defaults params
func (o *ReportsGetReportDefaultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserGroup adds the userGroup to the reports get report defaults params
func (o *ReportsGetReportDefaultsParams) WithUserGroup(userGroup string) *ReportsGetReportDefaultsParams {
	o.SetUserGroup(userGroup)
	return o
}

// SetUserGroup adds the userGroup to the reports get report defaults params
func (o *ReportsGetReportDefaultsParams) SetUserGroup(userGroup string) {
	o.UserGroup = userGroup
}

// WriteToRequest writes these params to a swagger request
func (o *ReportsGetReportDefaultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param userGroup
	if err := r.SetPathParam("userGroup", o.UserGroup); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
