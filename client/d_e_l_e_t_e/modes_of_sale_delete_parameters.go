// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

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

// NewModesOfSaleDeleteParams creates a new ModesOfSaleDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewModesOfSaleDeleteParams() *ModesOfSaleDeleteParams {
	return &ModesOfSaleDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewModesOfSaleDeleteParamsWithTimeout creates a new ModesOfSaleDeleteParams object
// with the ability to set a timeout on a request.
func NewModesOfSaleDeleteParamsWithTimeout(timeout time.Duration) *ModesOfSaleDeleteParams {
	return &ModesOfSaleDeleteParams{
		timeout: timeout,
	}
}

// NewModesOfSaleDeleteParamsWithContext creates a new ModesOfSaleDeleteParams object
// with the ability to set a context for a request.
func NewModesOfSaleDeleteParamsWithContext(ctx context.Context) *ModesOfSaleDeleteParams {
	return &ModesOfSaleDeleteParams{
		Context: ctx,
	}
}

// NewModesOfSaleDeleteParamsWithHTTPClient creates a new ModesOfSaleDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewModesOfSaleDeleteParamsWithHTTPClient(client *http.Client) *ModesOfSaleDeleteParams {
	return &ModesOfSaleDeleteParams{
		HTTPClient: client,
	}
}

/*
ModesOfSaleDeleteParams contains all the parameters to send to the API endpoint

	for the modes of sale delete operation.

	Typically these are written to a http.Request.
*/
type ModesOfSaleDeleteParams struct {

	// ModeOfSaleID.
	ModeOfSaleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the modes of sale delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModesOfSaleDeleteParams) WithDefaults() *ModesOfSaleDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the modes of sale delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModesOfSaleDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the modes of sale delete params
func (o *ModesOfSaleDeleteParams) WithTimeout(timeout time.Duration) *ModesOfSaleDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modes of sale delete params
func (o *ModesOfSaleDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modes of sale delete params
func (o *ModesOfSaleDeleteParams) WithContext(ctx context.Context) *ModesOfSaleDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modes of sale delete params
func (o *ModesOfSaleDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modes of sale delete params
func (o *ModesOfSaleDeleteParams) WithHTTPClient(client *http.Client) *ModesOfSaleDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modes of sale delete params
func (o *ModesOfSaleDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModeOfSaleID adds the modeOfSaleID to the modes of sale delete params
func (o *ModesOfSaleDeleteParams) WithModeOfSaleID(modeOfSaleID string) *ModesOfSaleDeleteParams {
	o.SetModeOfSaleID(modeOfSaleID)
	return o
}

// SetModeOfSaleID adds the modeOfSaleId to the modes of sale delete params
func (o *ModesOfSaleDeleteParams) SetModeOfSaleID(modeOfSaleID string) {
	o.ModeOfSaleID = modeOfSaleID
}

// WriteToRequest writes these params to a swagger request
func (o *ModesOfSaleDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param modeOfSaleId
	if err := r.SetPathParam("modeOfSaleId", o.ModeOfSaleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
