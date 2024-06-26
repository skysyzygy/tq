// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

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

	"github.com/skysyzygy/tq/models"
)

// NewDiagnosticsPutParams creates a new DiagnosticsPutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDiagnosticsPutParams() *DiagnosticsPutParams {
	return &DiagnosticsPutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDiagnosticsPutParamsWithTimeout creates a new DiagnosticsPutParams object
// with the ability to set a timeout on a request.
func NewDiagnosticsPutParamsWithTimeout(timeout time.Duration) *DiagnosticsPutParams {
	return &DiagnosticsPutParams{
		timeout: timeout,
	}
}

// NewDiagnosticsPutParamsWithContext creates a new DiagnosticsPutParams object
// with the ability to set a context for a request.
func NewDiagnosticsPutParamsWithContext(ctx context.Context) *DiagnosticsPutParams {
	return &DiagnosticsPutParams{
		Context: ctx,
	}
}

// NewDiagnosticsPutParamsWithHTTPClient creates a new DiagnosticsPutParams object
// with the ability to set a custom HTTPClient for a request.
func NewDiagnosticsPutParamsWithHTTPClient(client *http.Client) *DiagnosticsPutParams {
	return &DiagnosticsPutParams{
		HTTPClient: client,
	}
}

/*
DiagnosticsPutParams contains all the parameters to send to the API endpoint

	for the diagnostics put operation.

	Typically these are written to a http.Request.
*/
type DiagnosticsPutParams struct {

	// Diagnostic.
	Diagnostic *models.Diagnostic

	// DiagnosticID.
	DiagnosticID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the diagnostics put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DiagnosticsPutParams) WithDefaults() *DiagnosticsPutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the diagnostics put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DiagnosticsPutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the diagnostics put params
func (o *DiagnosticsPutParams) WithTimeout(timeout time.Duration) *DiagnosticsPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the diagnostics put params
func (o *DiagnosticsPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the diagnostics put params
func (o *DiagnosticsPutParams) WithContext(ctx context.Context) *DiagnosticsPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the diagnostics put params
func (o *DiagnosticsPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the diagnostics put params
func (o *DiagnosticsPutParams) WithHTTPClient(client *http.Client) *DiagnosticsPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the diagnostics put params
func (o *DiagnosticsPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDiagnostic adds the diagnostic to the diagnostics put params
func (o *DiagnosticsPutParams) WithDiagnostic(diagnostic *models.Diagnostic) *DiagnosticsPutParams {
	o.SetDiagnostic(diagnostic)
	return o
}

// SetDiagnostic adds the diagnostic to the diagnostics put params
func (o *DiagnosticsPutParams) SetDiagnostic(diagnostic *models.Diagnostic) {
	o.Diagnostic = diagnostic
}

// WithDiagnosticID adds the diagnosticID to the diagnostics put params
func (o *DiagnosticsPutParams) WithDiagnosticID(diagnosticID string) *DiagnosticsPutParams {
	o.SetDiagnosticID(diagnosticID)
	return o
}

// SetDiagnosticID adds the diagnosticId to the diagnostics put params
func (o *DiagnosticsPutParams) SetDiagnosticID(diagnosticID string) {
	o.DiagnosticID = diagnosticID
}

// WriteToRequest writes these params to a swagger request
func (o *DiagnosticsPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Diagnostic != nil {
		if err := r.SetBodyParam(o.Diagnostic); err != nil {
			return err
		}
	}

	// path param diagnosticId
	if err := r.SetPathParam("diagnosticId", o.DiagnosticID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
