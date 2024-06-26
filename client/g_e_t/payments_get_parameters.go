// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

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

// NewPaymentsGetParams creates a new PaymentsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPaymentsGetParams() *PaymentsGetParams {
	return &PaymentsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPaymentsGetParamsWithTimeout creates a new PaymentsGetParams object
// with the ability to set a timeout on a request.
func NewPaymentsGetParamsWithTimeout(timeout time.Duration) *PaymentsGetParams {
	return &PaymentsGetParams{
		timeout: timeout,
	}
}

// NewPaymentsGetParamsWithContext creates a new PaymentsGetParams object
// with the ability to set a context for a request.
func NewPaymentsGetParamsWithContext(ctx context.Context) *PaymentsGetParams {
	return &PaymentsGetParams{
		Context: ctx,
	}
}

// NewPaymentsGetParamsWithHTTPClient creates a new PaymentsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPaymentsGetParamsWithHTTPClient(client *http.Client) *PaymentsGetParams {
	return &PaymentsGetParams{
		HTTPClient: client,
	}
}

/*
PaymentsGetParams contains all the parameters to send to the API endpoint

	for the payments get operation.

	Typically these are written to a http.Request.
*/
type PaymentsGetParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the payments get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PaymentsGetParams) WithDefaults() *PaymentsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the payments get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PaymentsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the payments get params
func (o *PaymentsGetParams) WithTimeout(timeout time.Duration) *PaymentsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the payments get params
func (o *PaymentsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the payments get params
func (o *PaymentsGetParams) WithContext(ctx context.Context) *PaymentsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the payments get params
func (o *PaymentsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the payments get params
func (o *PaymentsGetParams) WithHTTPClient(client *http.Client) *PaymentsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the payments get params
func (o *PaymentsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the payments get params
func (o *PaymentsGetParams) WithID(id string) *PaymentsGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the payments get params
func (o *PaymentsGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PaymentsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
