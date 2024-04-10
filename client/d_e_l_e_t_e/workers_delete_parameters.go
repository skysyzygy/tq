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

// NewWorkersDeleteParams creates a new WorkersDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkersDeleteParams() *WorkersDeleteParams {
	return &WorkersDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkersDeleteParamsWithTimeout creates a new WorkersDeleteParams object
// with the ability to set a timeout on a request.
func NewWorkersDeleteParamsWithTimeout(timeout time.Duration) *WorkersDeleteParams {
	return &WorkersDeleteParams{
		timeout: timeout,
	}
}

// NewWorkersDeleteParamsWithContext creates a new WorkersDeleteParams object
// with the ability to set a context for a request.
func NewWorkersDeleteParamsWithContext(ctx context.Context) *WorkersDeleteParams {
	return &WorkersDeleteParams{
		Context: ctx,
	}
}

// NewWorkersDeleteParamsWithHTTPClient creates a new WorkersDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkersDeleteParamsWithHTTPClient(client *http.Client) *WorkersDeleteParams {
	return &WorkersDeleteParams{
		HTTPClient: client,
	}
}

/*
WorkersDeleteParams contains all the parameters to send to the API endpoint

	for the workers delete operation.

	Typically these are written to a http.Request.
*/
type WorkersDeleteParams struct {

	// WorkerID.
	WorkerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the workers delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkersDeleteParams) WithDefaults() *WorkersDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workers delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkersDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the workers delete params
func (o *WorkersDeleteParams) WithTimeout(timeout time.Duration) *WorkersDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workers delete params
func (o *WorkersDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workers delete params
func (o *WorkersDeleteParams) WithContext(ctx context.Context) *WorkersDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workers delete params
func (o *WorkersDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workers delete params
func (o *WorkersDeleteParams) WithHTTPClient(client *http.Client) *WorkersDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workers delete params
func (o *WorkersDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkerID adds the workerID to the workers delete params
func (o *WorkersDeleteParams) WithWorkerID(workerID string) *WorkersDeleteParams {
	o.SetWorkerID(workerID)
	return o
}

// SetWorkerID adds the workerId to the workers delete params
func (o *WorkersDeleteParams) SetWorkerID(workerID string) {
	o.WorkerID = workerID
}

// WriteToRequest writes these params to a swagger request
func (o *WorkersDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param workerId
	if err := r.SetPathParam("workerId", o.WorkerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
