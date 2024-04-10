// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

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

// NewDeliveryMethodsUpdateParams creates a new DeliveryMethodsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeliveryMethodsUpdateParams() *DeliveryMethodsUpdateParams {
	return &DeliveryMethodsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeliveryMethodsUpdateParamsWithTimeout creates a new DeliveryMethodsUpdateParams object
// with the ability to set a timeout on a request.
func NewDeliveryMethodsUpdateParamsWithTimeout(timeout time.Duration) *DeliveryMethodsUpdateParams {
	return &DeliveryMethodsUpdateParams{
		timeout: timeout,
	}
}

// NewDeliveryMethodsUpdateParamsWithContext creates a new DeliveryMethodsUpdateParams object
// with the ability to set a context for a request.
func NewDeliveryMethodsUpdateParamsWithContext(ctx context.Context) *DeliveryMethodsUpdateParams {
	return &DeliveryMethodsUpdateParams{
		Context: ctx,
	}
}

// NewDeliveryMethodsUpdateParamsWithHTTPClient creates a new DeliveryMethodsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeliveryMethodsUpdateParamsWithHTTPClient(client *http.Client) *DeliveryMethodsUpdateParams {
	return &DeliveryMethodsUpdateParams{
		HTTPClient: client,
	}
}

/*
DeliveryMethodsUpdateParams contains all the parameters to send to the API endpoint

	for the delivery methods update operation.

	Typically these are written to a http.Request.
*/
type DeliveryMethodsUpdateParams struct {

	// Data.
	Data *models.DeliveryMethod

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delivery methods update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeliveryMethodsUpdateParams) WithDefaults() *DeliveryMethodsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delivery methods update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeliveryMethodsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delivery methods update params
func (o *DeliveryMethodsUpdateParams) WithTimeout(timeout time.Duration) *DeliveryMethodsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delivery methods update params
func (o *DeliveryMethodsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delivery methods update params
func (o *DeliveryMethodsUpdateParams) WithContext(ctx context.Context) *DeliveryMethodsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delivery methods update params
func (o *DeliveryMethodsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delivery methods update params
func (o *DeliveryMethodsUpdateParams) WithHTTPClient(client *http.Client) *DeliveryMethodsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delivery methods update params
func (o *DeliveryMethodsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the delivery methods update params
func (o *DeliveryMethodsUpdateParams) WithData(data *models.DeliveryMethod) *DeliveryMethodsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the delivery methods update params
func (o *DeliveryMethodsUpdateParams) SetData(data *models.DeliveryMethod) {
	o.Data = data
}

// WithID adds the id to the delivery methods update params
func (o *DeliveryMethodsUpdateParams) WithID(id string) *DeliveryMethodsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delivery methods update params
func (o *DeliveryMethodsUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeliveryMethodsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
