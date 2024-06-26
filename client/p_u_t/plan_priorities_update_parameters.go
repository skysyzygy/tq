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

// NewPlanPrioritiesUpdateParams creates a new PlanPrioritiesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPlanPrioritiesUpdateParams() *PlanPrioritiesUpdateParams {
	return &PlanPrioritiesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPlanPrioritiesUpdateParamsWithTimeout creates a new PlanPrioritiesUpdateParams object
// with the ability to set a timeout on a request.
func NewPlanPrioritiesUpdateParamsWithTimeout(timeout time.Duration) *PlanPrioritiesUpdateParams {
	return &PlanPrioritiesUpdateParams{
		timeout: timeout,
	}
}

// NewPlanPrioritiesUpdateParamsWithContext creates a new PlanPrioritiesUpdateParams object
// with the ability to set a context for a request.
func NewPlanPrioritiesUpdateParamsWithContext(ctx context.Context) *PlanPrioritiesUpdateParams {
	return &PlanPrioritiesUpdateParams{
		Context: ctx,
	}
}

// NewPlanPrioritiesUpdateParamsWithHTTPClient creates a new PlanPrioritiesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPlanPrioritiesUpdateParamsWithHTTPClient(client *http.Client) *PlanPrioritiesUpdateParams {
	return &PlanPrioritiesUpdateParams{
		HTTPClient: client,
	}
}

/*
PlanPrioritiesUpdateParams contains all the parameters to send to the API endpoint

	for the plan priorities update operation.

	Typically these are written to a http.Request.
*/
type PlanPrioritiesUpdateParams struct {

	// Data.
	Data *models.PlanPriority

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plan priorities update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PlanPrioritiesUpdateParams) WithDefaults() *PlanPrioritiesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plan priorities update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PlanPrioritiesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plan priorities update params
func (o *PlanPrioritiesUpdateParams) WithTimeout(timeout time.Duration) *PlanPrioritiesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plan priorities update params
func (o *PlanPrioritiesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plan priorities update params
func (o *PlanPrioritiesUpdateParams) WithContext(ctx context.Context) *PlanPrioritiesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plan priorities update params
func (o *PlanPrioritiesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plan priorities update params
func (o *PlanPrioritiesUpdateParams) WithHTTPClient(client *http.Client) *PlanPrioritiesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plan priorities update params
func (o *PlanPrioritiesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plan priorities update params
func (o *PlanPrioritiesUpdateParams) WithData(data *models.PlanPriority) *PlanPrioritiesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plan priorities update params
func (o *PlanPrioritiesUpdateParams) SetData(data *models.PlanPriority) {
	o.Data = data
}

// WithID adds the id to the plan priorities update params
func (o *PlanPrioritiesUpdateParams) WithID(id string) *PlanPrioritiesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plan priorities update params
func (o *PlanPrioritiesUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PlanPrioritiesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
