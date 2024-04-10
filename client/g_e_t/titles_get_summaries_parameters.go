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

// NewTitlesGetSummariesParams creates a new TitlesGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTitlesGetSummariesParams() *TitlesGetSummariesParams {
	return &TitlesGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTitlesGetSummariesParamsWithTimeout creates a new TitlesGetSummariesParams object
// with the ability to set a timeout on a request.
func NewTitlesGetSummariesParamsWithTimeout(timeout time.Duration) *TitlesGetSummariesParams {
	return &TitlesGetSummariesParams{
		timeout: timeout,
	}
}

// NewTitlesGetSummariesParamsWithContext creates a new TitlesGetSummariesParams object
// with the ability to set a context for a request.
func NewTitlesGetSummariesParamsWithContext(ctx context.Context) *TitlesGetSummariesParams {
	return &TitlesGetSummariesParams{
		Context: ctx,
	}
}

// NewTitlesGetSummariesParamsWithHTTPClient creates a new TitlesGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewTitlesGetSummariesParamsWithHTTPClient(client *http.Client) *TitlesGetSummariesParams {
	return &TitlesGetSummariesParams{
		HTTPClient: client,
	}
}

/*
TitlesGetSummariesParams contains all the parameters to send to the API endpoint

	for the titles get summaries operation.

	Typically these are written to a http.Request.
*/
type TitlesGetSummariesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the titles get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TitlesGetSummariesParams) WithDefaults() *TitlesGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the titles get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TitlesGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the titles get summaries params
func (o *TitlesGetSummariesParams) WithTimeout(timeout time.Duration) *TitlesGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the titles get summaries params
func (o *TitlesGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the titles get summaries params
func (o *TitlesGetSummariesParams) WithContext(ctx context.Context) *TitlesGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the titles get summaries params
func (o *TitlesGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the titles get summaries params
func (o *TitlesGetSummariesParams) WithHTTPClient(client *http.Client) *TitlesGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the titles get summaries params
func (o *TitlesGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *TitlesGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
