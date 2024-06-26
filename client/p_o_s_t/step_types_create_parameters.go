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

// NewStepTypesCreateParams creates a new StepTypesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStepTypesCreateParams() *StepTypesCreateParams {
	return &StepTypesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStepTypesCreateParamsWithTimeout creates a new StepTypesCreateParams object
// with the ability to set a timeout on a request.
func NewStepTypesCreateParamsWithTimeout(timeout time.Duration) *StepTypesCreateParams {
	return &StepTypesCreateParams{
		timeout: timeout,
	}
}

// NewStepTypesCreateParamsWithContext creates a new StepTypesCreateParams object
// with the ability to set a context for a request.
func NewStepTypesCreateParamsWithContext(ctx context.Context) *StepTypesCreateParams {
	return &StepTypesCreateParams{
		Context: ctx,
	}
}

// NewStepTypesCreateParamsWithHTTPClient creates a new StepTypesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewStepTypesCreateParamsWithHTTPClient(client *http.Client) *StepTypesCreateParams {
	return &StepTypesCreateParams{
		HTTPClient: client,
	}
}

/*
StepTypesCreateParams contains all the parameters to send to the API endpoint

	for the step types create operation.

	Typically these are written to a http.Request.
*/
type StepTypesCreateParams struct {

	// Data.
	Data *models.StepType

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the step types create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StepTypesCreateParams) WithDefaults() *StepTypesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the step types create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StepTypesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the step types create params
func (o *StepTypesCreateParams) WithTimeout(timeout time.Duration) *StepTypesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the step types create params
func (o *StepTypesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the step types create params
func (o *StepTypesCreateParams) WithContext(ctx context.Context) *StepTypesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the step types create params
func (o *StepTypesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the step types create params
func (o *StepTypesCreateParams) WithHTTPClient(client *http.Client) *StepTypesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the step types create params
func (o *StepTypesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the step types create params
func (o *StepTypesCreateParams) WithData(data *models.StepType) *StepTypesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the step types create params
func (o *StepTypesCreateParams) SetData(data *models.StepType) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *StepTypesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
