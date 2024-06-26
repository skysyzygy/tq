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

// NewPlanPrioritiesCreateParams creates a new PlanPrioritiesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPlanPrioritiesCreateParams() *PlanPrioritiesCreateParams {
	return &PlanPrioritiesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPlanPrioritiesCreateParamsWithTimeout creates a new PlanPrioritiesCreateParams object
// with the ability to set a timeout on a request.
func NewPlanPrioritiesCreateParamsWithTimeout(timeout time.Duration) *PlanPrioritiesCreateParams {
	return &PlanPrioritiesCreateParams{
		timeout: timeout,
	}
}

// NewPlanPrioritiesCreateParamsWithContext creates a new PlanPrioritiesCreateParams object
// with the ability to set a context for a request.
func NewPlanPrioritiesCreateParamsWithContext(ctx context.Context) *PlanPrioritiesCreateParams {
	return &PlanPrioritiesCreateParams{
		Context: ctx,
	}
}

// NewPlanPrioritiesCreateParamsWithHTTPClient creates a new PlanPrioritiesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPlanPrioritiesCreateParamsWithHTTPClient(client *http.Client) *PlanPrioritiesCreateParams {
	return &PlanPrioritiesCreateParams{
		HTTPClient: client,
	}
}

/*
PlanPrioritiesCreateParams contains all the parameters to send to the API endpoint

	for the plan priorities create operation.

	Typically these are written to a http.Request.
*/
type PlanPrioritiesCreateParams struct {

	// Data.
	Data *models.PlanPriority

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plan priorities create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PlanPrioritiesCreateParams) WithDefaults() *PlanPrioritiesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plan priorities create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PlanPrioritiesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plan priorities create params
func (o *PlanPrioritiesCreateParams) WithTimeout(timeout time.Duration) *PlanPrioritiesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plan priorities create params
func (o *PlanPrioritiesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plan priorities create params
func (o *PlanPrioritiesCreateParams) WithContext(ctx context.Context) *PlanPrioritiesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plan priorities create params
func (o *PlanPrioritiesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plan priorities create params
func (o *PlanPrioritiesCreateParams) WithHTTPClient(client *http.Client) *PlanPrioritiesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plan priorities create params
func (o *PlanPrioritiesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plan priorities create params
func (o *PlanPrioritiesCreateParams) WithData(data *models.PlanPriority) *PlanPrioritiesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plan priorities create params
func (o *PlanPrioritiesCreateParams) SetData(data *models.PlanPriority) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PlanPrioritiesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
