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

// NewWorkerQualificationsUpdateParams creates a new WorkerQualificationsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkerQualificationsUpdateParams() *WorkerQualificationsUpdateParams {
	return &WorkerQualificationsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkerQualificationsUpdateParamsWithTimeout creates a new WorkerQualificationsUpdateParams object
// with the ability to set a timeout on a request.
func NewWorkerQualificationsUpdateParamsWithTimeout(timeout time.Duration) *WorkerQualificationsUpdateParams {
	return &WorkerQualificationsUpdateParams{
		timeout: timeout,
	}
}

// NewWorkerQualificationsUpdateParamsWithContext creates a new WorkerQualificationsUpdateParams object
// with the ability to set a context for a request.
func NewWorkerQualificationsUpdateParamsWithContext(ctx context.Context) *WorkerQualificationsUpdateParams {
	return &WorkerQualificationsUpdateParams{
		Context: ctx,
	}
}

// NewWorkerQualificationsUpdateParamsWithHTTPClient creates a new WorkerQualificationsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkerQualificationsUpdateParamsWithHTTPClient(client *http.Client) *WorkerQualificationsUpdateParams {
	return &WorkerQualificationsUpdateParams{
		HTTPClient: client,
	}
}

/*
WorkerQualificationsUpdateParams contains all the parameters to send to the API endpoint

	for the worker qualifications update operation.

	Typically these are written to a http.Request.
*/
type WorkerQualificationsUpdateParams struct {

	// WorkerQualification.
	WorkerQualification *models.WorkerQualification

	// WorkerQualificationID.
	WorkerQualificationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the worker qualifications update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkerQualificationsUpdateParams) WithDefaults() *WorkerQualificationsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the worker qualifications update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkerQualificationsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the worker qualifications update params
func (o *WorkerQualificationsUpdateParams) WithTimeout(timeout time.Duration) *WorkerQualificationsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the worker qualifications update params
func (o *WorkerQualificationsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the worker qualifications update params
func (o *WorkerQualificationsUpdateParams) WithContext(ctx context.Context) *WorkerQualificationsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the worker qualifications update params
func (o *WorkerQualificationsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the worker qualifications update params
func (o *WorkerQualificationsUpdateParams) WithHTTPClient(client *http.Client) *WorkerQualificationsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the worker qualifications update params
func (o *WorkerQualificationsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkerQualification adds the workerQualification to the worker qualifications update params
func (o *WorkerQualificationsUpdateParams) WithWorkerQualification(workerQualification *models.WorkerQualification) *WorkerQualificationsUpdateParams {
	o.SetWorkerQualification(workerQualification)
	return o
}

// SetWorkerQualification adds the workerQualification to the worker qualifications update params
func (o *WorkerQualificationsUpdateParams) SetWorkerQualification(workerQualification *models.WorkerQualification) {
	o.WorkerQualification = workerQualification
}

// WithWorkerQualificationID adds the workerQualificationID to the worker qualifications update params
func (o *WorkerQualificationsUpdateParams) WithWorkerQualificationID(workerQualificationID string) *WorkerQualificationsUpdateParams {
	o.SetWorkerQualificationID(workerQualificationID)
	return o
}

// SetWorkerQualificationID adds the workerQualificationId to the worker qualifications update params
func (o *WorkerQualificationsUpdateParams) SetWorkerQualificationID(workerQualificationID string) {
	o.WorkerQualificationID = workerQualificationID
}

// WriteToRequest writes these params to a swagger request
func (o *WorkerQualificationsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.WorkerQualification != nil {
		if err := r.SetBodyParam(o.WorkerQualification); err != nil {
			return err
		}
	}

	// path param workerQualificationId
	if err := r.SetPathParam("workerQualificationId", o.WorkerQualificationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}