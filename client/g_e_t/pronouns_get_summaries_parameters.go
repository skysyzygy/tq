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

// NewPronounsGetSummariesParams creates a new PronounsGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPronounsGetSummariesParams() *PronounsGetSummariesParams {
	return &PronounsGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPronounsGetSummariesParamsWithTimeout creates a new PronounsGetSummariesParams object
// with the ability to set a timeout on a request.
func NewPronounsGetSummariesParamsWithTimeout(timeout time.Duration) *PronounsGetSummariesParams {
	return &PronounsGetSummariesParams{
		timeout: timeout,
	}
}

// NewPronounsGetSummariesParamsWithContext creates a new PronounsGetSummariesParams object
// with the ability to set a context for a request.
func NewPronounsGetSummariesParamsWithContext(ctx context.Context) *PronounsGetSummariesParams {
	return &PronounsGetSummariesParams{
		Context: ctx,
	}
}

// NewPronounsGetSummariesParamsWithHTTPClient creates a new PronounsGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPronounsGetSummariesParamsWithHTTPClient(client *http.Client) *PronounsGetSummariesParams {
	return &PronounsGetSummariesParams{
		HTTPClient: client,
	}
}

/*
PronounsGetSummariesParams contains all the parameters to send to the API endpoint

	for the pronouns get summaries operation.

	Typically these are written to a http.Request.
*/
type PronounsGetSummariesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pronouns get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PronounsGetSummariesParams) WithDefaults() *PronounsGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pronouns get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PronounsGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pronouns get summaries params
func (o *PronounsGetSummariesParams) WithTimeout(timeout time.Duration) *PronounsGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pronouns get summaries params
func (o *PronounsGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pronouns get summaries params
func (o *PronounsGetSummariesParams) WithContext(ctx context.Context) *PronounsGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pronouns get summaries params
func (o *PronounsGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pronouns get summaries params
func (o *PronounsGetSummariesParams) WithHTTPClient(client *http.Client) *PronounsGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pronouns get summaries params
func (o *PronounsGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PronounsGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
