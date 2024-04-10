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

// NewConstituentInactivesCreateParams creates a new ConstituentInactivesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConstituentInactivesCreateParams() *ConstituentInactivesCreateParams {
	return &ConstituentInactivesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConstituentInactivesCreateParamsWithTimeout creates a new ConstituentInactivesCreateParams object
// with the ability to set a timeout on a request.
func NewConstituentInactivesCreateParamsWithTimeout(timeout time.Duration) *ConstituentInactivesCreateParams {
	return &ConstituentInactivesCreateParams{
		timeout: timeout,
	}
}

// NewConstituentInactivesCreateParamsWithContext creates a new ConstituentInactivesCreateParams object
// with the ability to set a context for a request.
func NewConstituentInactivesCreateParamsWithContext(ctx context.Context) *ConstituentInactivesCreateParams {
	return &ConstituentInactivesCreateParams{
		Context: ctx,
	}
}

// NewConstituentInactivesCreateParamsWithHTTPClient creates a new ConstituentInactivesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewConstituentInactivesCreateParamsWithHTTPClient(client *http.Client) *ConstituentInactivesCreateParams {
	return &ConstituentInactivesCreateParams{
		HTTPClient: client,
	}
}

/*
ConstituentInactivesCreateParams contains all the parameters to send to the API endpoint

	for the constituent inactives create operation.

	Typically these are written to a http.Request.
*/
type ConstituentInactivesCreateParams struct {

	// Data.
	Data *models.ConstituentInactive

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the constituent inactives create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConstituentInactivesCreateParams) WithDefaults() *ConstituentInactivesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the constituent inactives create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConstituentInactivesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the constituent inactives create params
func (o *ConstituentInactivesCreateParams) WithTimeout(timeout time.Duration) *ConstituentInactivesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the constituent inactives create params
func (o *ConstituentInactivesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the constituent inactives create params
func (o *ConstituentInactivesCreateParams) WithContext(ctx context.Context) *ConstituentInactivesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the constituent inactives create params
func (o *ConstituentInactivesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the constituent inactives create params
func (o *ConstituentInactivesCreateParams) WithHTTPClient(client *http.Client) *ConstituentInactivesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the constituent inactives create params
func (o *ConstituentInactivesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the constituent inactives create params
func (o *ConstituentInactivesCreateParams) WithData(data *models.ConstituentInactive) *ConstituentInactivesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the constituent inactives create params
func (o *ConstituentInactivesCreateParams) SetData(data *models.ConstituentInactive) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ConstituentInactivesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
