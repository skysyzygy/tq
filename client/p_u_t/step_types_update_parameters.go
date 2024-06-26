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

// NewStepTypesUpdateParams creates a new StepTypesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStepTypesUpdateParams() *StepTypesUpdateParams {
	return &StepTypesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStepTypesUpdateParamsWithTimeout creates a new StepTypesUpdateParams object
// with the ability to set a timeout on a request.
func NewStepTypesUpdateParamsWithTimeout(timeout time.Duration) *StepTypesUpdateParams {
	return &StepTypesUpdateParams{
		timeout: timeout,
	}
}

// NewStepTypesUpdateParamsWithContext creates a new StepTypesUpdateParams object
// with the ability to set a context for a request.
func NewStepTypesUpdateParamsWithContext(ctx context.Context) *StepTypesUpdateParams {
	return &StepTypesUpdateParams{
		Context: ctx,
	}
}

// NewStepTypesUpdateParamsWithHTTPClient creates a new StepTypesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewStepTypesUpdateParamsWithHTTPClient(client *http.Client) *StepTypesUpdateParams {
	return &StepTypesUpdateParams{
		HTTPClient: client,
	}
}

/*
StepTypesUpdateParams contains all the parameters to send to the API endpoint

	for the step types update operation.

	Typically these are written to a http.Request.
*/
type StepTypesUpdateParams struct {

	// Data.
	Data *models.StepType

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the step types update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StepTypesUpdateParams) WithDefaults() *StepTypesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the step types update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StepTypesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the step types update params
func (o *StepTypesUpdateParams) WithTimeout(timeout time.Duration) *StepTypesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the step types update params
func (o *StepTypesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the step types update params
func (o *StepTypesUpdateParams) WithContext(ctx context.Context) *StepTypesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the step types update params
func (o *StepTypesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the step types update params
func (o *StepTypesUpdateParams) WithHTTPClient(client *http.Client) *StepTypesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the step types update params
func (o *StepTypesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the step types update params
func (o *StepTypesUpdateParams) WithData(data *models.StepType) *StepTypesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the step types update params
func (o *StepTypesUpdateParams) SetData(data *models.StepType) {
	o.Data = data
}

// WithID adds the id to the step types update params
func (o *StepTypesUpdateParams) WithID(id string) *StepTypesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the step types update params
func (o *StepTypesUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *StepTypesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
