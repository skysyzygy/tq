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

// NewIntegrationsCreateParams creates a new IntegrationsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIntegrationsCreateParams() *IntegrationsCreateParams {
	return &IntegrationsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIntegrationsCreateParamsWithTimeout creates a new IntegrationsCreateParams object
// with the ability to set a timeout on a request.
func NewIntegrationsCreateParamsWithTimeout(timeout time.Duration) *IntegrationsCreateParams {
	return &IntegrationsCreateParams{
		timeout: timeout,
	}
}

// NewIntegrationsCreateParamsWithContext creates a new IntegrationsCreateParams object
// with the ability to set a context for a request.
func NewIntegrationsCreateParamsWithContext(ctx context.Context) *IntegrationsCreateParams {
	return &IntegrationsCreateParams{
		Context: ctx,
	}
}

// NewIntegrationsCreateParamsWithHTTPClient creates a new IntegrationsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIntegrationsCreateParamsWithHTTPClient(client *http.Client) *IntegrationsCreateParams {
	return &IntegrationsCreateParams{
		HTTPClient: client,
	}
}

/*
IntegrationsCreateParams contains all the parameters to send to the API endpoint

	for the integrations create operation.

	Typically these are written to a http.Request.
*/
type IntegrationsCreateParams struct {

	// Data.
	Data *models.Integration

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the integrations create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IntegrationsCreateParams) WithDefaults() *IntegrationsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the integrations create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IntegrationsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the integrations create params
func (o *IntegrationsCreateParams) WithTimeout(timeout time.Duration) *IntegrationsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the integrations create params
func (o *IntegrationsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the integrations create params
func (o *IntegrationsCreateParams) WithContext(ctx context.Context) *IntegrationsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the integrations create params
func (o *IntegrationsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the integrations create params
func (o *IntegrationsCreateParams) WithHTTPClient(client *http.Client) *IntegrationsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the integrations create params
func (o *IntegrationsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the integrations create params
func (o *IntegrationsCreateParams) WithData(data *models.Integration) *IntegrationsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the integrations create params
func (o *IntegrationsCreateParams) SetData(data *models.Integration) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IntegrationsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
