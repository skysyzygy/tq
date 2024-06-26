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

// NewIntegrationsGetSummariesParams creates a new IntegrationsGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIntegrationsGetSummariesParams() *IntegrationsGetSummariesParams {
	return &IntegrationsGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIntegrationsGetSummariesParamsWithTimeout creates a new IntegrationsGetSummariesParams object
// with the ability to set a timeout on a request.
func NewIntegrationsGetSummariesParamsWithTimeout(timeout time.Duration) *IntegrationsGetSummariesParams {
	return &IntegrationsGetSummariesParams{
		timeout: timeout,
	}
}

// NewIntegrationsGetSummariesParamsWithContext creates a new IntegrationsGetSummariesParams object
// with the ability to set a context for a request.
func NewIntegrationsGetSummariesParamsWithContext(ctx context.Context) *IntegrationsGetSummariesParams {
	return &IntegrationsGetSummariesParams{
		Context: ctx,
	}
}

// NewIntegrationsGetSummariesParamsWithHTTPClient creates a new IntegrationsGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewIntegrationsGetSummariesParamsWithHTTPClient(client *http.Client) *IntegrationsGetSummariesParams {
	return &IntegrationsGetSummariesParams{
		HTTPClient: client,
	}
}

/*
IntegrationsGetSummariesParams contains all the parameters to send to the API endpoint

	for the integrations get summaries operation.

	Typically these are written to a http.Request.
*/
type IntegrationsGetSummariesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the integrations get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IntegrationsGetSummariesParams) WithDefaults() *IntegrationsGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the integrations get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IntegrationsGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the integrations get summaries params
func (o *IntegrationsGetSummariesParams) WithTimeout(timeout time.Duration) *IntegrationsGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the integrations get summaries params
func (o *IntegrationsGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the integrations get summaries params
func (o *IntegrationsGetSummariesParams) WithContext(ctx context.Context) *IntegrationsGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the integrations get summaries params
func (o *IntegrationsGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the integrations get summaries params
func (o *IntegrationsGetSummariesParams) WithHTTPClient(client *http.Client) *IntegrationsGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the integrations get summaries params
func (o *IntegrationsGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *IntegrationsGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
