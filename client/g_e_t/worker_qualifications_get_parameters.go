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

// NewWorkerQualificationsGetParams creates a new WorkerQualificationsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkerQualificationsGetParams() *WorkerQualificationsGetParams {
	return &WorkerQualificationsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkerQualificationsGetParamsWithTimeout creates a new WorkerQualificationsGetParams object
// with the ability to set a timeout on a request.
func NewWorkerQualificationsGetParamsWithTimeout(timeout time.Duration) *WorkerQualificationsGetParams {
	return &WorkerQualificationsGetParams{
		timeout: timeout,
	}
}

// NewWorkerQualificationsGetParamsWithContext creates a new WorkerQualificationsGetParams object
// with the ability to set a context for a request.
func NewWorkerQualificationsGetParamsWithContext(ctx context.Context) *WorkerQualificationsGetParams {
	return &WorkerQualificationsGetParams{
		Context: ctx,
	}
}

// NewWorkerQualificationsGetParamsWithHTTPClient creates a new WorkerQualificationsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkerQualificationsGetParamsWithHTTPClient(client *http.Client) *WorkerQualificationsGetParams {
	return &WorkerQualificationsGetParams{
		HTTPClient: client,
	}
}

/*
WorkerQualificationsGetParams contains all the parameters to send to the API endpoint

	for the worker qualifications get operation.

	Typically these are written to a http.Request.
*/
type WorkerQualificationsGetParams struct {

	// WorkerQualificationID.
	WorkerQualificationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the worker qualifications get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkerQualificationsGetParams) WithDefaults() *WorkerQualificationsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the worker qualifications get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkerQualificationsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the worker qualifications get params
func (o *WorkerQualificationsGetParams) WithTimeout(timeout time.Duration) *WorkerQualificationsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the worker qualifications get params
func (o *WorkerQualificationsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the worker qualifications get params
func (o *WorkerQualificationsGetParams) WithContext(ctx context.Context) *WorkerQualificationsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the worker qualifications get params
func (o *WorkerQualificationsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the worker qualifications get params
func (o *WorkerQualificationsGetParams) WithHTTPClient(client *http.Client) *WorkerQualificationsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the worker qualifications get params
func (o *WorkerQualificationsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkerQualificationID adds the workerQualificationID to the worker qualifications get params
func (o *WorkerQualificationsGetParams) WithWorkerQualificationID(workerQualificationID string) *WorkerQualificationsGetParams {
	o.SetWorkerQualificationID(workerQualificationID)
	return o
}

// SetWorkerQualificationID adds the workerQualificationId to the worker qualifications get params
func (o *WorkerQualificationsGetParams) SetWorkerQualificationID(workerQualificationID string) {
	o.WorkerQualificationID = workerQualificationID
}

// WriteToRequest writes these params to a swagger request
func (o *WorkerQualificationsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param workerQualificationId
	if err := r.SetPathParam("workerQualificationId", o.WorkerQualificationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}