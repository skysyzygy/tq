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

// NewErasGetSummariesParams creates a new ErasGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewErasGetSummariesParams() *ErasGetSummariesParams {
	return &ErasGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewErasGetSummariesParamsWithTimeout creates a new ErasGetSummariesParams object
// with the ability to set a timeout on a request.
func NewErasGetSummariesParamsWithTimeout(timeout time.Duration) *ErasGetSummariesParams {
	return &ErasGetSummariesParams{
		timeout: timeout,
	}
}

// NewErasGetSummariesParamsWithContext creates a new ErasGetSummariesParams object
// with the ability to set a context for a request.
func NewErasGetSummariesParamsWithContext(ctx context.Context) *ErasGetSummariesParams {
	return &ErasGetSummariesParams{
		Context: ctx,
	}
}

// NewErasGetSummariesParamsWithHTTPClient creates a new ErasGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewErasGetSummariesParamsWithHTTPClient(client *http.Client) *ErasGetSummariesParams {
	return &ErasGetSummariesParams{
		HTTPClient: client,
	}
}

/*
ErasGetSummariesParams contains all the parameters to send to the API endpoint

	for the eras get summaries operation.

	Typically these are written to a http.Request.
*/
type ErasGetSummariesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the eras get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ErasGetSummariesParams) WithDefaults() *ErasGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the eras get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ErasGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the eras get summaries params
func (o *ErasGetSummariesParams) WithTimeout(timeout time.Duration) *ErasGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the eras get summaries params
func (o *ErasGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the eras get summaries params
func (o *ErasGetSummariesParams) WithContext(ctx context.Context) *ErasGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the eras get summaries params
func (o *ErasGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the eras get summaries params
func (o *ErasGetSummariesParams) WithHTTPClient(client *http.Client) *ErasGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the eras get summaries params
func (o *ErasGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ErasGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
