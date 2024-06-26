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

// NewPlanWorkersGetParams creates a new PlanWorkersGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPlanWorkersGetParams() *PlanWorkersGetParams {
	return &PlanWorkersGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPlanWorkersGetParamsWithTimeout creates a new PlanWorkersGetParams object
// with the ability to set a timeout on a request.
func NewPlanWorkersGetParamsWithTimeout(timeout time.Duration) *PlanWorkersGetParams {
	return &PlanWorkersGetParams{
		timeout: timeout,
	}
}

// NewPlanWorkersGetParamsWithContext creates a new PlanWorkersGetParams object
// with the ability to set a context for a request.
func NewPlanWorkersGetParamsWithContext(ctx context.Context) *PlanWorkersGetParams {
	return &PlanWorkersGetParams{
		Context: ctx,
	}
}

// NewPlanWorkersGetParamsWithHTTPClient creates a new PlanWorkersGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPlanWorkersGetParamsWithHTTPClient(client *http.Client) *PlanWorkersGetParams {
	return &PlanWorkersGetParams{
		HTTPClient: client,
	}
}

/*
PlanWorkersGetParams contains all the parameters to send to the API endpoint

	for the plan workers get operation.

	Typically these are written to a http.Request.
*/
type PlanWorkersGetParams struct {

	// PlanWorkerID.
	PlanWorkerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plan workers get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PlanWorkersGetParams) WithDefaults() *PlanWorkersGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plan workers get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PlanWorkersGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plan workers get params
func (o *PlanWorkersGetParams) WithTimeout(timeout time.Duration) *PlanWorkersGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plan workers get params
func (o *PlanWorkersGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plan workers get params
func (o *PlanWorkersGetParams) WithContext(ctx context.Context) *PlanWorkersGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plan workers get params
func (o *PlanWorkersGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plan workers get params
func (o *PlanWorkersGetParams) WithHTTPClient(client *http.Client) *PlanWorkersGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plan workers get params
func (o *PlanWorkersGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPlanWorkerID adds the planWorkerID to the plan workers get params
func (o *PlanWorkersGetParams) WithPlanWorkerID(planWorkerID string) *PlanWorkersGetParams {
	o.SetPlanWorkerID(planWorkerID)
	return o
}

// SetPlanWorkerID adds the planWorkerId to the plan workers get params
func (o *PlanWorkersGetParams) SetPlanWorkerID(planWorkerID string) {
	o.PlanWorkerID = planWorkerID
}

// WriteToRequest writes these params to a swagger request
func (o *PlanWorkersGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
