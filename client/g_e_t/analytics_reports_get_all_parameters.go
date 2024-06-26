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

// NewAnalyticsReportsGetAllParams creates a new AnalyticsReportsGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAnalyticsReportsGetAllParams() *AnalyticsReportsGetAllParams {
	return &AnalyticsReportsGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAnalyticsReportsGetAllParamsWithTimeout creates a new AnalyticsReportsGetAllParams object
// with the ability to set a timeout on a request.
func NewAnalyticsReportsGetAllParamsWithTimeout(timeout time.Duration) *AnalyticsReportsGetAllParams {
	return &AnalyticsReportsGetAllParams{
		timeout: timeout,
	}
}

// NewAnalyticsReportsGetAllParamsWithContext creates a new AnalyticsReportsGetAllParams object
// with the ability to set a context for a request.
func NewAnalyticsReportsGetAllParamsWithContext(ctx context.Context) *AnalyticsReportsGetAllParams {
	return &AnalyticsReportsGetAllParams{
		Context: ctx,
	}
}

// NewAnalyticsReportsGetAllParamsWithHTTPClient creates a new AnalyticsReportsGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewAnalyticsReportsGetAllParamsWithHTTPClient(client *http.Client) *AnalyticsReportsGetAllParams {
	return &AnalyticsReportsGetAllParams{
		HTTPClient: client,
	}
}

/*
AnalyticsReportsGetAllParams contains all the parameters to send to the API endpoint

	for the analytics reports get all operation.

	Typically these are written to a http.Request.
*/
type AnalyticsReportsGetAllParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the analytics reports get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AnalyticsReportsGetAllParams) WithDefaults() *AnalyticsReportsGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the analytics reports get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AnalyticsReportsGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the analytics reports get all params
func (o *AnalyticsReportsGetAllParams) WithTimeout(timeout time.Duration) *AnalyticsReportsGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the analytics reports get all params
func (o *AnalyticsReportsGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the analytics reports get all params
func (o *AnalyticsReportsGetAllParams) WithContext(ctx context.Context) *AnalyticsReportsGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the analytics reports get all params
func (o *AnalyticsReportsGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the analytics reports get all params
func (o *AnalyticsReportsGetAllParams) WithHTTPClient(client *http.Client) *AnalyticsReportsGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the analytics reports get all params
func (o *AnalyticsReportsGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AnalyticsReportsGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
