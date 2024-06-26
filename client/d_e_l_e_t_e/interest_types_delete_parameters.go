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

// NewInterestTypesDeleteParams creates a new InterestTypesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInterestTypesDeleteParams() *InterestTypesDeleteParams {
	return &InterestTypesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInterestTypesDeleteParamsWithTimeout creates a new InterestTypesDeleteParams object
// with the ability to set a timeout on a request.
func NewInterestTypesDeleteParamsWithTimeout(timeout time.Duration) *InterestTypesDeleteParams {
	return &InterestTypesDeleteParams{
		timeout: timeout,
	}
}

// NewInterestTypesDeleteParamsWithContext creates a new InterestTypesDeleteParams object
// with the ability to set a context for a request.
func NewInterestTypesDeleteParamsWithContext(ctx context.Context) *InterestTypesDeleteParams {
	return &InterestTypesDeleteParams{
		Context: ctx,
	}
}

// NewInterestTypesDeleteParamsWithHTTPClient creates a new InterestTypesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewInterestTypesDeleteParamsWithHTTPClient(client *http.Client) *InterestTypesDeleteParams {
	return &InterestTypesDeleteParams{
		HTTPClient: client,
	}
}

/*
InterestTypesDeleteParams contains all the parameters to send to the API endpoint

	for the interest types delete operation.

	Typically these are written to a http.Request.
*/
type InterestTypesDeleteParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the interest types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InterestTypesDeleteParams) WithDefaults() *InterestTypesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the interest types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InterestTypesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the interest types delete params
func (o *InterestTypesDeleteParams) WithTimeout(timeout time.Duration) *InterestTypesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the interest types delete params
func (o *InterestTypesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the interest types delete params
func (o *InterestTypesDeleteParams) WithContext(ctx context.Context) *InterestTypesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the interest types delete params
func (o *InterestTypesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the interest types delete params
func (o *InterestTypesDeleteParams) WithHTTPClient(client *http.Client) *InterestTypesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the interest types delete params
func (o *InterestTypesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the interest types delete params
func (o *InterestTypesDeleteParams) WithID(id string) *InterestTypesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the interest types delete params
func (o *InterestTypesDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *InterestTypesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
