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

// NewPlanWorkersCreateParams creates a new PlanWorkersCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPlanWorkersCreateParams() *PlanWorkersCreateParams {
	return &PlanWorkersCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPlanWorkersCreateParamsWithTimeout creates a new PlanWorkersCreateParams object
// with the ability to set a timeout on a request.
func NewPlanWorkersCreateParamsWithTimeout(timeout time.Duration) *PlanWorkersCreateParams {
	return &PlanWorkersCreateParams{
		timeout: timeout,
	}
}

// NewPlanWorkersCreateParamsWithContext creates a new PlanWorkersCreateParams object
// with the ability to set a context for a request.
func NewPlanWorkersCreateParamsWithContext(ctx context.Context) *PlanWorkersCreateParams {
	return &PlanWorkersCreateParams{
		Context: ctx,
	}
}

// NewPlanWorkersCreateParamsWithHTTPClient creates a new PlanWorkersCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPlanWorkersCreateParamsWithHTTPClient(client *http.Client) *PlanWorkersCreateParams {
	return &PlanWorkersCreateParams{
		HTTPClient: client,
	}
}

/*
PlanWorkersCreateParams contains all the parameters to send to the API endpoint

	for the plan workers create operation.

	Typically these are written to a http.Request.
*/
type PlanWorkersCreateParams struct {

	// PlanWorker.
	PlanWorker *models.PlanWorker

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plan workers create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PlanWorkersCreateParams) WithDefaults() *PlanWorkersCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plan workers create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PlanWorkersCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plan workers create params
func (o *PlanWorkersCreateParams) WithTimeout(timeout time.Duration) *PlanWorkersCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plan workers create params
func (o *PlanWorkersCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plan workers create params
func (o *PlanWorkersCreateParams) WithContext(ctx context.Context) *PlanWorkersCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plan workers create params
func (o *PlanWorkersCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plan workers create params
func (o *PlanWorkersCreateParams) WithHTTPClient(client *http.Client) *PlanWorkersCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plan workers create params
func (o *PlanWorkersCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPlanWorker adds the planWorker to the plan workers create params
func (o *PlanWorkersCreateParams) WithPlanWorker(planWorker *models.PlanWorker) *PlanWorkersCreateParams {
	o.SetPlanWorker(planWorker)
	return o
}

// SetPlanWorker adds the planWorker to the plan workers create params
func (o *PlanWorkersCreateParams) SetPlanWorker(planWorker *models.PlanWorker) {
	o.PlanWorker = planWorker
}

// WriteToRequest writes these params to a swagger request
func (o *PlanWorkersCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PlanWorker != nil {
		if err := r.SetBodyParam(o.PlanWorker); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
