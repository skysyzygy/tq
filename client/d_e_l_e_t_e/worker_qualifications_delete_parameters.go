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

// NewWorkerQualificationsDeleteParams creates a new WorkerQualificationsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkerQualificationsDeleteParams() *WorkerQualificationsDeleteParams {
	return &WorkerQualificationsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkerQualificationsDeleteParamsWithTimeout creates a new WorkerQualificationsDeleteParams object
// with the ability to set a timeout on a request.
func NewWorkerQualificationsDeleteParamsWithTimeout(timeout time.Duration) *WorkerQualificationsDeleteParams {
	return &WorkerQualificationsDeleteParams{
		timeout: timeout,
	}
}

// NewWorkerQualificationsDeleteParamsWithContext creates a new WorkerQualificationsDeleteParams object
// with the ability to set a context for a request.
func NewWorkerQualificationsDeleteParamsWithContext(ctx context.Context) *WorkerQualificationsDeleteParams {
	return &WorkerQualificationsDeleteParams{
		Context: ctx,
	}
}

// NewWorkerQualificationsDeleteParamsWithHTTPClient creates a new WorkerQualificationsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkerQualificationsDeleteParamsWithHTTPClient(client *http.Client) *WorkerQualificationsDeleteParams {
	return &WorkerQualificationsDeleteParams{
		HTTPClient: client,
	}
}

/*
WorkerQualificationsDeleteParams contains all the parameters to send to the API endpoint

	for the worker qualifications delete operation.

	Typically these are written to a http.Request.
*/
type WorkerQualificationsDeleteParams struct {

	// WorkerQualificationID.
	WorkerQualificationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the worker qualifications delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkerQualificationsDeleteParams) WithDefaults() *WorkerQualificationsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the worker qualifications delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkerQualificationsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the worker qualifications delete params
func (o *WorkerQualificationsDeleteParams) WithTimeout(timeout time.Duration) *WorkerQualificationsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the worker qualifications delete params
func (o *WorkerQualificationsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the worker qualifications delete params
func (o *WorkerQualificationsDeleteParams) WithContext(ctx context.Context) *WorkerQualificationsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the worker qualifications delete params
func (o *WorkerQualificationsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the worker qualifications delete params
func (o *WorkerQualificationsDeleteParams) WithHTTPClient(client *http.Client) *WorkerQualificationsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the worker qualifications delete params
func (o *WorkerQualificationsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkerQualificationID adds the workerQualificationID to the worker qualifications delete params
func (o *WorkerQualificationsDeleteParams) WithWorkerQualificationID(workerQualificationID string) *WorkerQualificationsDeleteParams {
	o.SetWorkerQualificationID(workerQualificationID)
	return o
}

// SetWorkerQualificationID adds the workerQualificationId to the worker qualifications delete params
func (o *WorkerQualificationsDeleteParams) SetWorkerQualificationID(workerQualificationID string) {
	o.WorkerQualificationID = workerQualificationID
}

// WriteToRequest writes these params to a swagger request
func (o *WorkerQualificationsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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