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

// NewPronounsDeleteParams creates a new PronounsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPronounsDeleteParams() *PronounsDeleteParams {
	return &PronounsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPronounsDeleteParamsWithTimeout creates a new PronounsDeleteParams object
// with the ability to set a timeout on a request.
func NewPronounsDeleteParamsWithTimeout(timeout time.Duration) *PronounsDeleteParams {
	return &PronounsDeleteParams{
		timeout: timeout,
	}
}

// NewPronounsDeleteParamsWithContext creates a new PronounsDeleteParams object
// with the ability to set a context for a request.
func NewPronounsDeleteParamsWithContext(ctx context.Context) *PronounsDeleteParams {
	return &PronounsDeleteParams{
		Context: ctx,
	}
}

// NewPronounsDeleteParamsWithHTTPClient creates a new PronounsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPronounsDeleteParamsWithHTTPClient(client *http.Client) *PronounsDeleteParams {
	return &PronounsDeleteParams{
		HTTPClient: client,
	}
}

/*
PronounsDeleteParams contains all the parameters to send to the API endpoint

	for the pronouns delete operation.

	Typically these are written to a http.Request.
*/
type PronounsDeleteParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pronouns delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PronounsDeleteParams) WithDefaults() *PronounsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pronouns delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PronounsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pronouns delete params
func (o *PronounsDeleteParams) WithTimeout(timeout time.Duration) *PronounsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pronouns delete params
func (o *PronounsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pronouns delete params
func (o *PronounsDeleteParams) WithContext(ctx context.Context) *PronounsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pronouns delete params
func (o *PronounsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pronouns delete params
func (o *PronounsDeleteParams) WithHTTPClient(client *http.Client) *PronounsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pronouns delete params
func (o *PronounsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the pronouns delete params
func (o *PronounsDeleteParams) WithID(id string) *PronounsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the pronouns delete params
func (o *PronounsDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PronounsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
