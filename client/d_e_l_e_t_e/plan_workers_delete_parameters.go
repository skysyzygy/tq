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

// NewPlanWorkersDeleteParams creates a new PlanWorkersDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPlanWorkersDeleteParams() *PlanWorkersDeleteParams {
	return &PlanWorkersDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPlanWorkersDeleteParamsWithTimeout creates a new PlanWorkersDeleteParams object
// with the ability to set a timeout on a request.
func NewPlanWorkersDeleteParamsWithTimeout(timeout time.Duration) *PlanWorkersDeleteParams {
	return &PlanWorkersDeleteParams{
		timeout: timeout,
	}
}

// NewPlanWorkersDeleteParamsWithContext creates a new PlanWorkersDeleteParams object
// with the ability to set a context for a request.
func NewPlanWorkersDeleteParamsWithContext(ctx context.Context) *PlanWorkersDeleteParams {
	return &PlanWorkersDeleteParams{
		Context: ctx,
	}
}

// NewPlanWorkersDeleteParamsWithHTTPClient creates a new PlanWorkersDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPlanWorkersDeleteParamsWithHTTPClient(client *http.Client) *PlanWorkersDeleteParams {
	return &PlanWorkersDeleteParams{
		HTTPClient: client,
	}
}

/*
PlanWorkersDeleteParams contains all the parameters to send to the API endpoint

	for the plan workers delete operation.

	Typically these are written to a http.Request.
*/
type PlanWorkersDeleteParams struct {

	// PlanWorkerID.
	PlanWorkerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plan workers delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PlanWorkersDeleteParams) WithDefaults() *PlanWorkersDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plan workers delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PlanWorkersDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plan workers delete params
func (o *PlanWorkersDeleteParams) WithTimeout(timeout time.Duration) *PlanWorkersDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plan workers delete params
func (o *PlanWorkersDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plan workers delete params
func (o *PlanWorkersDeleteParams) WithContext(ctx context.Context) *PlanWorkersDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plan workers delete params
func (o *PlanWorkersDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plan workers delete params
func (o *PlanWorkersDeleteParams) WithHTTPClient(client *http.Client) *PlanWorkersDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plan workers delete params
func (o *PlanWorkersDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPlanWorkerID adds the planWorkerID to the plan workers delete params
func (o *PlanWorkersDeleteParams) WithPlanWorkerID(planWorkerID string) *PlanWorkersDeleteParams {
	o.SetPlanWorkerID(planWorkerID)
	return o
}

// SetPlanWorkerID adds the planWorkerId to the plan workers delete params
func (o *PlanWorkersDeleteParams) SetPlanWorkerID(planWorkerID string) {
	o.PlanWorkerID = planWorkerID
}

// WriteToRequest writes these params to a swagger request
func (o *PlanWorkersDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param planWorkerId
	if err := r.SetPathParam("planWorkerId", o.PlanWorkerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
