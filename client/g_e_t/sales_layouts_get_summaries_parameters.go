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

// NewSalesLayoutsGetSummariesParams creates a new SalesLayoutsGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSalesLayoutsGetSummariesParams() *SalesLayoutsGetSummariesParams {
	return &SalesLayoutsGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSalesLayoutsGetSummariesParamsWithTimeout creates a new SalesLayoutsGetSummariesParams object
// with the ability to set a timeout on a request.
func NewSalesLayoutsGetSummariesParamsWithTimeout(timeout time.Duration) *SalesLayoutsGetSummariesParams {
	return &SalesLayoutsGetSummariesParams{
		timeout: timeout,
	}
}

// NewSalesLayoutsGetSummariesParamsWithContext creates a new SalesLayoutsGetSummariesParams object
// with the ability to set a context for a request.
func NewSalesLayoutsGetSummariesParamsWithContext(ctx context.Context) *SalesLayoutsGetSummariesParams {
	return &SalesLayoutsGetSummariesParams{
		Context: ctx,
	}
}

// NewSalesLayoutsGetSummariesParamsWithHTTPClient creates a new SalesLayoutsGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewSalesLayoutsGetSummariesParamsWithHTTPClient(client *http.Client) *SalesLayoutsGetSummariesParams {
	return &SalesLayoutsGetSummariesParams{
		HTTPClient: client,
	}
}

/*
SalesLayoutsGetSummariesParams contains all the parameters to send to the API endpoint

	for the sales layouts get summaries operation.

	Typically these are written to a http.Request.
*/
type SalesLayoutsGetSummariesParams struct {

	// PrimaryOnly.
	PrimaryOnly *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the sales layouts get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SalesLayoutsGetSummariesParams) WithDefaults() *SalesLayoutsGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the sales layouts get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SalesLayoutsGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the sales layouts get summaries params
func (o *SalesLayoutsGetSummariesParams) WithTimeout(timeout time.Duration) *SalesLayoutsGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sales layouts get summaries params
func (o *SalesLayoutsGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sales layouts get summaries params
func (o *SalesLayoutsGetSummariesParams) WithContext(ctx context.Context) *SalesLayoutsGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sales layouts get summaries params
func (o *SalesLayoutsGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sales layouts get summaries params
func (o *SalesLayoutsGetSummariesParams) WithHTTPClient(client *http.Client) *SalesLayoutsGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sales layouts get summaries params
func (o *SalesLayoutsGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPrimaryOnly adds the primaryOnly to the sales layouts get summaries params
func (o *SalesLayoutsGetSummariesParams) WithPrimaryOnly(primaryOnly *string) *SalesLayoutsGetSummariesParams {
	o.SetPrimaryOnly(primaryOnly)
	return o
}

// SetPrimaryOnly adds the primaryOnly to the sales layouts get summaries params
func (o *SalesLayoutsGetSummariesParams) SetPrimaryOnly(primaryOnly *string) {
	o.PrimaryOnly = primaryOnly
}

// WriteToRequest writes these params to a swagger request
func (o *SalesLayoutsGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PrimaryOnly != nil {

		// query param primaryOnly
		var qrPrimaryOnly string

		if o.PrimaryOnly != nil {
			qrPrimaryOnly = *o.PrimaryOnly
		}
		qPrimaryOnly := qrPrimaryOnly
		if qPrimaryOnly != "" {

			if err := r.SetQueryParam("primaryOnly", qPrimaryOnly); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
