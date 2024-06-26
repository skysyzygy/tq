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

// NewPhoneTypesDeleteParams creates a new PhoneTypesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPhoneTypesDeleteParams() *PhoneTypesDeleteParams {
	return &PhoneTypesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPhoneTypesDeleteParamsWithTimeout creates a new PhoneTypesDeleteParams object
// with the ability to set a timeout on a request.
func NewPhoneTypesDeleteParamsWithTimeout(timeout time.Duration) *PhoneTypesDeleteParams {
	return &PhoneTypesDeleteParams{
		timeout: timeout,
	}
}

// NewPhoneTypesDeleteParamsWithContext creates a new PhoneTypesDeleteParams object
// with the ability to set a context for a request.
func NewPhoneTypesDeleteParamsWithContext(ctx context.Context) *PhoneTypesDeleteParams {
	return &PhoneTypesDeleteParams{
		Context: ctx,
	}
}

// NewPhoneTypesDeleteParamsWithHTTPClient creates a new PhoneTypesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPhoneTypesDeleteParamsWithHTTPClient(client *http.Client) *PhoneTypesDeleteParams {
	return &PhoneTypesDeleteParams{
		HTTPClient: client,
	}
}

/*
PhoneTypesDeleteParams contains all the parameters to send to the API endpoint

	for the phone types delete operation.

	Typically these are written to a http.Request.
*/
type PhoneTypesDeleteParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the phone types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PhoneTypesDeleteParams) WithDefaults() *PhoneTypesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the phone types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PhoneTypesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the phone types delete params
func (o *PhoneTypesDeleteParams) WithTimeout(timeout time.Duration) *PhoneTypesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the phone types delete params
func (o *PhoneTypesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the phone types delete params
func (o *PhoneTypesDeleteParams) WithContext(ctx context.Context) *PhoneTypesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the phone types delete params
func (o *PhoneTypesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the phone types delete params
func (o *PhoneTypesDeleteParams) WithHTTPClient(client *http.Client) *PhoneTypesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the phone types delete params
func (o *PhoneTypesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the phone types delete params
func (o *PhoneTypesDeleteParams) WithID(id string) *PhoneTypesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the phone types delete params
func (o *PhoneTypesDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PhoneTypesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
