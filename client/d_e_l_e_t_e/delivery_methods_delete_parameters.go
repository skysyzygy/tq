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

// NewDeliveryMethodsDeleteParams creates a new DeliveryMethodsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeliveryMethodsDeleteParams() *DeliveryMethodsDeleteParams {
	return &DeliveryMethodsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeliveryMethodsDeleteParamsWithTimeout creates a new DeliveryMethodsDeleteParams object
// with the ability to set a timeout on a request.
func NewDeliveryMethodsDeleteParamsWithTimeout(timeout time.Duration) *DeliveryMethodsDeleteParams {
	return &DeliveryMethodsDeleteParams{
		timeout: timeout,
	}
}

// NewDeliveryMethodsDeleteParamsWithContext creates a new DeliveryMethodsDeleteParams object
// with the ability to set a context for a request.
func NewDeliveryMethodsDeleteParamsWithContext(ctx context.Context) *DeliveryMethodsDeleteParams {
	return &DeliveryMethodsDeleteParams{
		Context: ctx,
	}
}

// NewDeliveryMethodsDeleteParamsWithHTTPClient creates a new DeliveryMethodsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeliveryMethodsDeleteParamsWithHTTPClient(client *http.Client) *DeliveryMethodsDeleteParams {
	return &DeliveryMethodsDeleteParams{
		HTTPClient: client,
	}
}

/*
DeliveryMethodsDeleteParams contains all the parameters to send to the API endpoint

	for the delivery methods delete operation.

	Typically these are written to a http.Request.
*/
type DeliveryMethodsDeleteParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delivery methods delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeliveryMethodsDeleteParams) WithDefaults() *DeliveryMethodsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delivery methods delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeliveryMethodsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delivery methods delete params
func (o *DeliveryMethodsDeleteParams) WithTimeout(timeout time.Duration) *DeliveryMethodsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delivery methods delete params
func (o *DeliveryMethodsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delivery methods delete params
func (o *DeliveryMethodsDeleteParams) WithContext(ctx context.Context) *DeliveryMethodsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delivery methods delete params
func (o *DeliveryMethodsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delivery methods delete params
func (o *DeliveryMethodsDeleteParams) WithHTTPClient(client *http.Client) *DeliveryMethodsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delivery methods delete params
func (o *DeliveryMethodsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delivery methods delete params
func (o *DeliveryMethodsDeleteParams) WithID(id string) *DeliveryMethodsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delivery methods delete params
func (o *DeliveryMethodsDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeliveryMethodsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
