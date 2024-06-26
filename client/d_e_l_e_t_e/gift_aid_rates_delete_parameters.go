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

// NewGiftAidRatesDeleteParams creates a new GiftAidRatesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGiftAidRatesDeleteParams() *GiftAidRatesDeleteParams {
	return &GiftAidRatesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGiftAidRatesDeleteParamsWithTimeout creates a new GiftAidRatesDeleteParams object
// with the ability to set a timeout on a request.
func NewGiftAidRatesDeleteParamsWithTimeout(timeout time.Duration) *GiftAidRatesDeleteParams {
	return &GiftAidRatesDeleteParams{
		timeout: timeout,
	}
}

// NewGiftAidRatesDeleteParamsWithContext creates a new GiftAidRatesDeleteParams object
// with the ability to set a context for a request.
func NewGiftAidRatesDeleteParamsWithContext(ctx context.Context) *GiftAidRatesDeleteParams {
	return &GiftAidRatesDeleteParams{
		Context: ctx,
	}
}

// NewGiftAidRatesDeleteParamsWithHTTPClient creates a new GiftAidRatesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewGiftAidRatesDeleteParamsWithHTTPClient(client *http.Client) *GiftAidRatesDeleteParams {
	return &GiftAidRatesDeleteParams{
		HTTPClient: client,
	}
}

/*
GiftAidRatesDeleteParams contains all the parameters to send to the API endpoint

	for the gift aid rates delete operation.

	Typically these are written to a http.Request.
*/
type GiftAidRatesDeleteParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the gift aid rates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GiftAidRatesDeleteParams) WithDefaults() *GiftAidRatesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the gift aid rates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GiftAidRatesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the gift aid rates delete params
func (o *GiftAidRatesDeleteParams) WithTimeout(timeout time.Duration) *GiftAidRatesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gift aid rates delete params
func (o *GiftAidRatesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gift aid rates delete params
func (o *GiftAidRatesDeleteParams) WithContext(ctx context.Context) *GiftAidRatesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gift aid rates delete params
func (o *GiftAidRatesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gift aid rates delete params
func (o *GiftAidRatesDeleteParams) WithHTTPClient(client *http.Client) *GiftAidRatesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gift aid rates delete params
func (o *GiftAidRatesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the gift aid rates delete params
func (o *GiftAidRatesDeleteParams) WithID(id string) *GiftAidRatesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the gift aid rates delete params
func (o *GiftAidRatesDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GiftAidRatesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
