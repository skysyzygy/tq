// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

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

// NewDiagnosticsDeleteParams creates a new DiagnosticsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDiagnosticsDeleteParams() *DiagnosticsDeleteParams {
	return &DiagnosticsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDiagnosticsDeleteParamsWithTimeout creates a new DiagnosticsDeleteParams object
// with the ability to set a timeout on a request.
func NewDiagnosticsDeleteParamsWithTimeout(timeout time.Duration) *DiagnosticsDeleteParams {
	return &DiagnosticsDeleteParams{
		timeout: timeout,
	}
}

// NewDiagnosticsDeleteParamsWithContext creates a new DiagnosticsDeleteParams object
// with the ability to set a context for a request.
func NewDiagnosticsDeleteParamsWithContext(ctx context.Context) *DiagnosticsDeleteParams {
	return &DiagnosticsDeleteParams{
		Context: ctx,
	}
}

// NewDiagnosticsDeleteParamsWithHTTPClient creates a new DiagnosticsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDiagnosticsDeleteParamsWithHTTPClient(client *http.Client) *DiagnosticsDeleteParams {
	return &DiagnosticsDeleteParams{
		HTTPClient: client,
	}
}

/*
DiagnosticsDeleteParams contains all the parameters to send to the API endpoint

	for the diagnostics delete operation.

	Typically these are written to a http.Request.
*/
type DiagnosticsDeleteParams struct {

	// DiagnosticID.
	DiagnosticID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the diagnostics delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DiagnosticsDeleteParams) WithDefaults() *DiagnosticsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the diagnostics delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DiagnosticsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the diagnostics delete params
func (o *DiagnosticsDeleteParams) WithTimeout(timeout time.Duration) *DiagnosticsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the diagnostics delete params
func (o *DiagnosticsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the diagnostics delete params
func (o *DiagnosticsDeleteParams) WithContext(ctx context.Context) *DiagnosticsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the diagnostics delete params
func (o *DiagnosticsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the diagnostics delete params
func (o *DiagnosticsDeleteParams) WithHTTPClient(client *http.Client) *DiagnosticsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the diagnostics delete params
func (o *DiagnosticsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDiagnosticID adds the diagnosticID to the diagnostics delete params
func (o *DiagnosticsDeleteParams) WithDiagnosticID(diagnosticID string) *DiagnosticsDeleteParams {
	o.SetDiagnosticID(diagnosticID)
	return o
}

// SetDiagnosticID adds the diagnosticId to the diagnostics delete params
func (o *DiagnosticsDeleteParams) SetDiagnosticID(diagnosticID string) {
	o.DiagnosticID = diagnosticID
}

// WriteToRequest writes these params to a swagger request
func (o *DiagnosticsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param diagnosticId
	if err := r.SetPathParam("diagnosticId", o.DiagnosticID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
