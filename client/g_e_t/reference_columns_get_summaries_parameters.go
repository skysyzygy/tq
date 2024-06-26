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

// NewReferenceColumnsGetSummariesParams creates a new ReferenceColumnsGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReferenceColumnsGetSummariesParams() *ReferenceColumnsGetSummariesParams {
	return &ReferenceColumnsGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReferenceColumnsGetSummariesParamsWithTimeout creates a new ReferenceColumnsGetSummariesParams object
// with the ability to set a timeout on a request.
func NewReferenceColumnsGetSummariesParamsWithTimeout(timeout time.Duration) *ReferenceColumnsGetSummariesParams {
	return &ReferenceColumnsGetSummariesParams{
		timeout: timeout,
	}
}

// NewReferenceColumnsGetSummariesParamsWithContext creates a new ReferenceColumnsGetSummariesParams object
// with the ability to set a context for a request.
func NewReferenceColumnsGetSummariesParamsWithContext(ctx context.Context) *ReferenceColumnsGetSummariesParams {
	return &ReferenceColumnsGetSummariesParams{
		Context: ctx,
	}
}

// NewReferenceColumnsGetSummariesParamsWithHTTPClient creates a new ReferenceColumnsGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewReferenceColumnsGetSummariesParamsWithHTTPClient(client *http.Client) *ReferenceColumnsGetSummariesParams {
	return &ReferenceColumnsGetSummariesParams{
		HTTPClient: client,
	}
}

/*
ReferenceColumnsGetSummariesParams contains all the parameters to send to the API endpoint

	for the reference columns get summaries operation.

	Typically these are written to a http.Request.
*/
type ReferenceColumnsGetSummariesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the reference columns get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReferenceColumnsGetSummariesParams) WithDefaults() *ReferenceColumnsGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the reference columns get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReferenceColumnsGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the reference columns get summaries params
func (o *ReferenceColumnsGetSummariesParams) WithTimeout(timeout time.Duration) *ReferenceColumnsGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reference columns get summaries params
func (o *ReferenceColumnsGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reference columns get summaries params
func (o *ReferenceColumnsGetSummariesParams) WithContext(ctx context.Context) *ReferenceColumnsGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reference columns get summaries params
func (o *ReferenceColumnsGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reference columns get summaries params
func (o *ReferenceColumnsGetSummariesParams) WithHTTPClient(client *http.Client) *ReferenceColumnsGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reference columns get summaries params
func (o *ReferenceColumnsGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ReferenceColumnsGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
