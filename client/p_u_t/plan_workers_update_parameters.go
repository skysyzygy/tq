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

// NewPlanWorkersUpdateParams creates a new PlanWorkersUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPlanWorkersUpdateParams() *PlanWorkersUpdateParams {
	return &PlanWorkersUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPlanWorkersUpdateParamsWithTimeout creates a new PlanWorkersUpdateParams object
// with the ability to set a timeout on a request.
func NewPlanWorkersUpdateParamsWithTimeout(timeout time.Duration) *PlanWorkersUpdateParams {
	return &PlanWorkersUpdateParams{
		timeout: timeout,
	}
}

// NewPlanWorkersUpdateParamsWithContext creates a new PlanWorkersUpdateParams object
// with the ability to set a context for a request.
func NewPlanWorkersUpdateParamsWithContext(ctx context.Context) *PlanWorkersUpdateParams {
	return &PlanWorkersUpdateParams{
		Context: ctx,
	}
}

// NewPlanWorkersUpdateParamsWithHTTPClient creates a new PlanWorkersUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPlanWorkersUpdateParamsWithHTTPClient(client *http.Client) *PlanWorkersUpdateParams {
	return &PlanWorkersUpdateParams{
		HTTPClient: client,
	}
}

/*
PlanWorkersUpdateParams contains all the parameters to send to the API endpoint

	for the plan workers update operation.

	Typically these are written to a http.Request.
*/
type PlanWorkersUpdateParams struct {

	// PlanWorker.
	PlanWorker *models.PlanWorker

	// PlanWorkerID.
	PlanWorkerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plan workers update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PlanWorkersUpdateParams) WithDefaults() *PlanWorkersUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plan workers update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PlanWorkersUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plan workers update params
func (o *PlanWorkersUpdateParams) WithTimeout(timeout time.Duration) *PlanWorkersUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plan workers update params
func (o *PlanWorkersUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plan workers update params
func (o *PlanWorkersUpdateParams) WithContext(ctx context.Context) *PlanWorkersUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plan workers update params
func (o *PlanWorkersUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plan workers update params
func (o *PlanWorkersUpdateParams) WithHTTPClient(client *http.Client) *PlanWorkersUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plan workers update params
func (o *PlanWorkersUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPlanWorker adds the planWorker to the plan workers update params
func (o *PlanWorkersUpdateParams) WithPlanWorker(planWorker *models.PlanWorker) *PlanWorkersUpdateParams {
	o.SetPlanWorker(planWorker)
	return o
}

// SetPlanWorker adds the planWorker to the plan workers update params
func (o *PlanWorkersUpdateParams) SetPlanWorker(planWorker *models.PlanWorker) {
	o.PlanWorker = planWorker
}

// WithPlanWorkerID adds the planWorkerID to the plan workers update params
func (o *PlanWorkersUpdateParams) WithPlanWorkerID(planWorkerID string) *PlanWorkersUpdateParams {
	o.SetPlanWorkerID(planWorkerID)
	return o
}

// SetPlanWorkerID adds the planWorkerId to the plan workers update params
func (o *PlanWorkersUpdateParams) SetPlanWorkerID(planWorkerID string) {
	o.PlanWorkerID = planWorkerID
}

// WriteToRequest writes these params to a swagger request
func (o *PlanWorkersUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PlanWorker != nil {
		if err := r.SetBodyParam(o.PlanWorker); err != nil {
			return err
		}
	}

	// path param planWorkerId
	if err := r.SetPathParam("planWorkerId", o.PlanWorkerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
