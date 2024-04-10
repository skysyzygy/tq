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

// NewAddressTypesDeleteParams creates a new AddressTypesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddressTypesDeleteParams() *AddressTypesDeleteParams {
	return &AddressTypesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddressTypesDeleteParamsWithTimeout creates a new AddressTypesDeleteParams object
// with the ability to set a timeout on a request.
func NewAddressTypesDeleteParamsWithTimeout(timeout time.Duration) *AddressTypesDeleteParams {
	return &AddressTypesDeleteParams{
		timeout: timeout,
	}
}

// NewAddressTypesDeleteParamsWithContext creates a new AddressTypesDeleteParams object
// with the ability to set a context for a request.
func NewAddressTypesDeleteParamsWithContext(ctx context.Context) *AddressTypesDeleteParams {
	return &AddressTypesDeleteParams{
		Context: ctx,
	}
}

// NewAddressTypesDeleteParamsWithHTTPClient creates a new AddressTypesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddressTypesDeleteParamsWithHTTPClient(client *http.Client) *AddressTypesDeleteParams {
	return &AddressTypesDeleteParams{
		HTTPClient: client,
	}
}

/*
AddressTypesDeleteParams contains all the parameters to send to the API endpoint

	for the address types delete operation.

	Typically these are written to a http.Request.
*/
type AddressTypesDeleteParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the address types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressTypesDeleteParams) WithDefaults() *AddressTypesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the address types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressTypesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the address types delete params
func (o *AddressTypesDeleteParams) WithTimeout(timeout time.Duration) *AddressTypesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the address types delete params
func (o *AddressTypesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the address types delete params
func (o *AddressTypesDeleteParams) WithContext(ctx context.Context) *AddressTypesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the address types delete params
func (o *AddressTypesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the address types delete params
func (o *AddressTypesDeleteParams) WithHTTPClient(client *http.Client) *AddressTypesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the address types delete params
func (o *AddressTypesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the address types delete params
func (o *AddressTypesDeleteParams) WithID(id string) *AddressTypesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the address types delete params
func (o *AddressTypesDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AddressTypesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
