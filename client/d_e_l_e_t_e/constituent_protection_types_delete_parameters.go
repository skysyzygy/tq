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

// NewConstituentProtectionTypesDeleteParams creates a new ConstituentProtectionTypesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConstituentProtectionTypesDeleteParams() *ConstituentProtectionTypesDeleteParams {
	return &ConstituentProtectionTypesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConstituentProtectionTypesDeleteParamsWithTimeout creates a new ConstituentProtectionTypesDeleteParams object
// with the ability to set a timeout on a request.
func NewConstituentProtectionTypesDeleteParamsWithTimeout(timeout time.Duration) *ConstituentProtectionTypesDeleteParams {
	return &ConstituentProtectionTypesDeleteParams{
		timeout: timeout,
	}
}

// NewConstituentProtectionTypesDeleteParamsWithContext creates a new ConstituentProtectionTypesDeleteParams object
// with the ability to set a context for a request.
func NewConstituentProtectionTypesDeleteParamsWithContext(ctx context.Context) *ConstituentProtectionTypesDeleteParams {
	return &ConstituentProtectionTypesDeleteParams{
		Context: ctx,
	}
}

// NewConstituentProtectionTypesDeleteParamsWithHTTPClient creates a new ConstituentProtectionTypesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewConstituentProtectionTypesDeleteParamsWithHTTPClient(client *http.Client) *ConstituentProtectionTypesDeleteParams {
	return &ConstituentProtectionTypesDeleteParams{
		HTTPClient: client,
	}
}

/*
ConstituentProtectionTypesDeleteParams contains all the parameters to send to the API endpoint

	for the constituent protection types delete operation.

	Typically these are written to a http.Request.
*/
type ConstituentProtectionTypesDeleteParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the constituent protection types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConstituentProtectionTypesDeleteParams) WithDefaults() *ConstituentProtectionTypesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the constituent protection types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConstituentProtectionTypesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the constituent protection types delete params
func (o *ConstituentProtectionTypesDeleteParams) WithTimeout(timeout time.Duration) *ConstituentProtectionTypesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the constituent protection types delete params
func (o *ConstituentProtectionTypesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the constituent protection types delete params
func (o *ConstituentProtectionTypesDeleteParams) WithContext(ctx context.Context) *ConstituentProtectionTypesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the constituent protection types delete params
func (o *ConstituentProtectionTypesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the constituent protection types delete params
func (o *ConstituentProtectionTypesDeleteParams) WithHTTPClient(client *http.Client) *ConstituentProtectionTypesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the constituent protection types delete params
func (o *ConstituentProtectionTypesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the constituent protection types delete params
func (o *ConstituentProtectionTypesDeleteParams) WithID(id string) *ConstituentProtectionTypesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the constituent protection types delete params
func (o *ConstituentProtectionTypesDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ConstituentProtectionTypesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
