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

// NewWorkersUpdateParams creates a new WorkersUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkersUpdateParams() *WorkersUpdateParams {
	return &WorkersUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkersUpdateParamsWithTimeout creates a new WorkersUpdateParams object
// with the ability to set a timeout on a request.
func NewWorkersUpdateParamsWithTimeout(timeout time.Duration) *WorkersUpdateParams {
	return &WorkersUpdateParams{
		timeout: timeout,
	}
}

// NewWorkersUpdateParamsWithContext creates a new WorkersUpdateParams object
// with the ability to set a context for a request.
func NewWorkersUpdateParamsWithContext(ctx context.Context) *WorkersUpdateParams {
	return &WorkersUpdateParams{
		Context: ctx,
	}
}

// NewWorkersUpdateParamsWithHTTPClient creates a new WorkersUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkersUpdateParamsWithHTTPClient(client *http.Client) *WorkersUpdateParams {
	return &WorkersUpdateParams{
		HTTPClient: client,
	}
}

/*
WorkersUpdateParams contains all the parameters to send to the API endpoint

	for the workers update operation.

	Typically these are written to a http.Request.
*/
type WorkersUpdateParams struct {

	// Worker.
	Worker *models.Worker

	// WorkerID.
	WorkerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the workers update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkersUpdateParams) WithDefaults() *WorkersUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workers update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkersUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the workers update params
func (o *WorkersUpdateParams) WithTimeout(timeout time.Duration) *WorkersUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workers update params
func (o *WorkersUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workers update params
func (o *WorkersUpdateParams) WithContext(ctx context.Context) *WorkersUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workers update params
func (o *WorkersUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workers update params
func (o *WorkersUpdateParams) WithHTTPClient(client *http.Client) *WorkersUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workers update params
func (o *WorkersUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorker adds the worker to the workers update params
func (o *WorkersUpdateParams) WithWorker(worker *models.Worker) *WorkersUpdateParams {
	o.SetWorker(worker)
	return o
}

// SetWorker adds the worker to the workers update params
func (o *WorkersUpdateParams) SetWorker(worker *models.Worker) {
	o.Worker = worker
}

// WithWorkerID adds the workerID to the workers update params
func (o *WorkersUpdateParams) WithWorkerID(workerID string) *WorkersUpdateParams {
	o.SetWorkerID(workerID)
	return o
}

// SetWorkerID adds the workerId to the workers update params
func (o *WorkersUpdateParams) SetWorkerID(workerID string) {
	o.WorkerID = workerID
}

// WriteToRequest writes these params to a swagger request
func (o *WorkersUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Worker != nil {
		if err := r.SetBodyParam(o.Worker); err != nil {
			return err
		}
	}

	// path param workerId
	if err := r.SetPathParam("workerId", o.WorkerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
