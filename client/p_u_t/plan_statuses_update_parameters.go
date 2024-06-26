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

// NewPlanStatusesUpdateParams creates a new PlanStatusesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPlanStatusesUpdateParams() *PlanStatusesUpdateParams {
	return &PlanStatusesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPlanStatusesUpdateParamsWithTimeout creates a new PlanStatusesUpdateParams object
// with the ability to set a timeout on a request.
func NewPlanStatusesUpdateParamsWithTimeout(timeout time.Duration) *PlanStatusesUpdateParams {
	return &PlanStatusesUpdateParams{
		timeout: timeout,
	}
}

// NewPlanStatusesUpdateParamsWithContext creates a new PlanStatusesUpdateParams object
// with the ability to set a context for a request.
func NewPlanStatusesUpdateParamsWithContext(ctx context.Context) *PlanStatusesUpdateParams {
	return &PlanStatusesUpdateParams{
		Context: ctx,
	}
}

// NewPlanStatusesUpdateParamsWithHTTPClient creates a new PlanStatusesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPlanStatusesUpdateParamsWithHTTPClient(client *http.Client) *PlanStatusesUpdateParams {
	return &PlanStatusesUpdateParams{
		HTTPClient: client,
	}
}

/*
PlanStatusesUpdateParams contains all the parameters to send to the API endpoint

	for the plan statuses update operation.

	Typically these are written to a http.Request.
*/
type PlanStatusesUpdateParams struct {

	// Data.
	Data *models.PlanStatus

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plan statuses update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PlanStatusesUpdateParams) WithDefaults() *PlanStatusesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plan statuses update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PlanStatusesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plan statuses update params
func (o *PlanStatusesUpdateParams) WithTimeout(timeout time.Duration) *PlanStatusesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plan statuses update params
func (o *PlanStatusesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plan statuses update params
func (o *PlanStatusesUpdateParams) WithContext(ctx context.Context) *PlanStatusesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plan statuses update params
func (o *PlanStatusesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plan statuses update params
func (o *PlanStatusesUpdateParams) WithHTTPClient(client *http.Client) *PlanStatusesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plan statuses update params
func (o *PlanStatusesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plan statuses update params
func (o *PlanStatusesUpdateParams) WithData(data *models.PlanStatus) *PlanStatusesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plan statuses update params
func (o *PlanStatusesUpdateParams) SetData(data *models.PlanStatus) {
	o.Data = data
}

// WithID adds the id to the plan statuses update params
func (o *PlanStatusesUpdateParams) WithID(id string) *PlanStatusesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plan statuses update params
func (o *PlanStatusesUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PlanStatusesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
