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

// NewSystemDefaultsGetSummariesParams creates a new SystemDefaultsGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSystemDefaultsGetSummariesParams() *SystemDefaultsGetSummariesParams {
	return &SystemDefaultsGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSystemDefaultsGetSummariesParamsWithTimeout creates a new SystemDefaultsGetSummariesParams object
// with the ability to set a timeout on a request.
func NewSystemDefaultsGetSummariesParamsWithTimeout(timeout time.Duration) *SystemDefaultsGetSummariesParams {
	return &SystemDefaultsGetSummariesParams{
		timeout: timeout,
	}
}

// NewSystemDefaultsGetSummariesParamsWithContext creates a new SystemDefaultsGetSummariesParams object
// with the ability to set a context for a request.
func NewSystemDefaultsGetSummariesParamsWithContext(ctx context.Context) *SystemDefaultsGetSummariesParams {
	return &SystemDefaultsGetSummariesParams{
		Context: ctx,
	}
}

// NewSystemDefaultsGetSummariesParamsWithHTTPClient creates a new SystemDefaultsGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewSystemDefaultsGetSummariesParamsWithHTTPClient(client *http.Client) *SystemDefaultsGetSummariesParams {
	return &SystemDefaultsGetSummariesParams{
		HTTPClient: client,
	}
}

/*
SystemDefaultsGetSummariesParams contains all the parameters to send to the API endpoint

	for the system defaults get summaries operation.

	Typically these are written to a http.Request.
*/
type SystemDefaultsGetSummariesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the system defaults get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemDefaultsGetSummariesParams) WithDefaults() *SystemDefaultsGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the system defaults get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemDefaultsGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the system defaults get summaries params
func (o *SystemDefaultsGetSummariesParams) WithTimeout(timeout time.Duration) *SystemDefaultsGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system defaults get summaries params
func (o *SystemDefaultsGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system defaults get summaries params
func (o *SystemDefaultsGetSummariesParams) WithContext(ctx context.Context) *SystemDefaultsGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system defaults get summaries params
func (o *SystemDefaultsGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system defaults get summaries params
func (o *SystemDefaultsGetSummariesParams) WithHTTPClient(client *http.Client) *SystemDefaultsGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system defaults get summaries params
func (o *SystemDefaultsGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SystemDefaultsGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
