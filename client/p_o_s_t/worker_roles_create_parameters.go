// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

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

// NewWorkerRolesCreateParams creates a new WorkerRolesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkerRolesCreateParams() *WorkerRolesCreateParams {
	return &WorkerRolesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkerRolesCreateParamsWithTimeout creates a new WorkerRolesCreateParams object
// with the ability to set a timeout on a request.
func NewWorkerRolesCreateParamsWithTimeout(timeout time.Duration) *WorkerRolesCreateParams {
	return &WorkerRolesCreateParams{
		timeout: timeout,
	}
}

// NewWorkerRolesCreateParamsWithContext creates a new WorkerRolesCreateParams object
// with the ability to set a context for a request.
func NewWorkerRolesCreateParamsWithContext(ctx context.Context) *WorkerRolesCreateParams {
	return &WorkerRolesCreateParams{
		Context: ctx,
	}
}

// NewWorkerRolesCreateParamsWithHTTPClient creates a new WorkerRolesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkerRolesCreateParamsWithHTTPClient(client *http.Client) *WorkerRolesCreateParams {
	return &WorkerRolesCreateParams{
		HTTPClient: client,
	}
}

/*
WorkerRolesCreateParams contains all the parameters to send to the API endpoint

	for the worker roles create operation.

	Typically these are written to a http.Request.
*/
type WorkerRolesCreateParams struct {

	// Data.
	Data *models.WorkerRole

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the worker roles create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkerRolesCreateParams) WithDefaults() *WorkerRolesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the worker roles create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkerRolesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the worker roles create params
func (o *WorkerRolesCreateParams) WithTimeout(timeout time.Duration) *WorkerRolesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the worker roles create params
func (o *WorkerRolesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the worker roles create params
func (o *WorkerRolesCreateParams) WithContext(ctx context.Context) *WorkerRolesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the worker roles create params
func (o *WorkerRolesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the worker roles create params
func (o *WorkerRolesCreateParams) WithHTTPClient(client *http.Client) *WorkerRolesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the worker roles create params
func (o *WorkerRolesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the worker roles create params
func (o *WorkerRolesCreateParams) WithData(data *models.WorkerRole) *WorkerRolesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the worker roles create params
func (o *WorkerRolesCreateParams) SetData(data *models.WorkerRole) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *WorkerRolesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
