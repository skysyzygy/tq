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

// NewCrediteeTypesDeleteParams creates a new CrediteeTypesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCrediteeTypesDeleteParams() *CrediteeTypesDeleteParams {
	return &CrediteeTypesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCrediteeTypesDeleteParamsWithTimeout creates a new CrediteeTypesDeleteParams object
// with the ability to set a timeout on a request.
func NewCrediteeTypesDeleteParamsWithTimeout(timeout time.Duration) *CrediteeTypesDeleteParams {
	return &CrediteeTypesDeleteParams{
		timeout: timeout,
	}
}

// NewCrediteeTypesDeleteParamsWithContext creates a new CrediteeTypesDeleteParams object
// with the ability to set a context for a request.
func NewCrediteeTypesDeleteParamsWithContext(ctx context.Context) *CrediteeTypesDeleteParams {
	return &CrediteeTypesDeleteParams{
		Context: ctx,
	}
}

// NewCrediteeTypesDeleteParamsWithHTTPClient creates a new CrediteeTypesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewCrediteeTypesDeleteParamsWithHTTPClient(client *http.Client) *CrediteeTypesDeleteParams {
	return &CrediteeTypesDeleteParams{
		HTTPClient: client,
	}
}

/*
CrediteeTypesDeleteParams contains all the parameters to send to the API endpoint

	for the creditee types delete operation.

	Typically these are written to a http.Request.
*/
type CrediteeTypesDeleteParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the creditee types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CrediteeTypesDeleteParams) WithDefaults() *CrediteeTypesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the creditee types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CrediteeTypesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the creditee types delete params
func (o *CrediteeTypesDeleteParams) WithTimeout(timeout time.Duration) *CrediteeTypesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the creditee types delete params
func (o *CrediteeTypesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the creditee types delete params
func (o *CrediteeTypesDeleteParams) WithContext(ctx context.Context) *CrediteeTypesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the creditee types delete params
func (o *CrediteeTypesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the creditee types delete params
func (o *CrediteeTypesDeleteParams) WithHTTPClient(client *http.Client) *CrediteeTypesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the creditee types delete params
func (o *CrediteeTypesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the creditee types delete params
func (o *CrediteeTypesDeleteParams) WithID(id string) *CrediteeTypesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the creditee types delete params
func (o *CrediteeTypesDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CrediteeTypesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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