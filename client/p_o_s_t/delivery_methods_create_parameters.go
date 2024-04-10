// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

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

	"github.com/skysyzygy/tq/models"
)

// NewDeliveryMethodsCreateParams creates a new DeliveryMethodsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeliveryMethodsCreateParams() *DeliveryMethodsCreateParams {
	return &DeliveryMethodsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeliveryMethodsCreateParamsWithTimeout creates a new DeliveryMethodsCreateParams object
// with the ability to set a timeout on a request.
func NewDeliveryMethodsCreateParamsWithTimeout(timeout time.Duration) *DeliveryMethodsCreateParams {
	return &DeliveryMethodsCreateParams{
		timeout: timeout,
	}
}

// NewDeliveryMethodsCreateParamsWithContext creates a new DeliveryMethodsCreateParams object
// with the ability to set a context for a request.
func NewDeliveryMethodsCreateParamsWithContext(ctx context.Context) *DeliveryMethodsCreateParams {
	return &DeliveryMethodsCreateParams{
		Context: ctx,
	}
}

// NewDeliveryMethodsCreateParamsWithHTTPClient creates a new DeliveryMethodsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeliveryMethodsCreateParamsWithHTTPClient(client *http.Client) *DeliveryMethodsCreateParams {
	return &DeliveryMethodsCreateParams{
		HTTPClient: client,
	}
}

/*
DeliveryMethodsCreateParams contains all the parameters to send to the API endpoint

	for the delivery methods create operation.

	Typically these are written to a http.Request.
*/
type DeliveryMethodsCreateParams struct {

	// Data.
	Data *models.DeliveryMethod

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delivery methods create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeliveryMethodsCreateParams) WithDefaults() *DeliveryMethodsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delivery methods create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeliveryMethodsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delivery methods create params
func (o *DeliveryMethodsCreateParams) WithTimeout(timeout time.Duration) *DeliveryMethodsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delivery methods create params
func (o *DeliveryMethodsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delivery methods create params
func (o *DeliveryMethodsCreateParams) WithContext(ctx context.Context) *DeliveryMethodsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delivery methods create params
func (o *DeliveryMethodsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delivery methods create params
func (o *DeliveryMethodsCreateParams) WithHTTPClient(client *http.Client) *DeliveryMethodsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delivery methods create params
func (o *DeliveryMethodsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the delivery methods create params
func (o *DeliveryMethodsCreateParams) WithData(data *models.DeliveryMethod) *DeliveryMethodsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the delivery methods create params
func (o *DeliveryMethodsCreateParams) SetData(data *models.DeliveryMethod) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DeliveryMethodsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
