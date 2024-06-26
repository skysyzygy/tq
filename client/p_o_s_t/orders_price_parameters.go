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

// NewOrdersPriceParams creates a new OrdersPriceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrdersPriceParams() *OrdersPriceParams {
	return &OrdersPriceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrdersPriceParamsWithTimeout creates a new OrdersPriceParams object
// with the ability to set a timeout on a request.
func NewOrdersPriceParamsWithTimeout(timeout time.Duration) *OrdersPriceParams {
	return &OrdersPriceParams{
		timeout: timeout,
	}
}

// NewOrdersPriceParamsWithContext creates a new OrdersPriceParams object
// with the ability to set a context for a request.
func NewOrdersPriceParamsWithContext(ctx context.Context) *OrdersPriceParams {
	return &OrdersPriceParams{
		Context: ctx,
	}
}

// NewOrdersPriceParamsWithHTTPClient creates a new OrdersPriceParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrdersPriceParamsWithHTTPClient(client *http.Client) *OrdersPriceParams {
	return &OrdersPriceParams{
		HTTPClient: client,
	}
}

/*
OrdersPriceParams contains all the parameters to send to the API endpoint

	for the orders price operation.

	Typically these are written to a http.Request.
*/
type OrdersPriceParams struct {

	// Order.
	Order *models.Order

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the orders price params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrdersPriceParams) WithDefaults() *OrdersPriceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the orders price params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrdersPriceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the orders price params
func (o *OrdersPriceParams) WithTimeout(timeout time.Duration) *OrdersPriceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the orders price params
func (o *OrdersPriceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the orders price params
func (o *OrdersPriceParams) WithContext(ctx context.Context) *OrdersPriceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the orders price params
func (o *OrdersPriceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the orders price params
func (o *OrdersPriceParams) WithHTTPClient(client *http.Client) *OrdersPriceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the orders price params
func (o *OrdersPriceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrder adds the order to the orders price params
func (o *OrdersPriceParams) WithOrder(order *models.Order) *OrdersPriceParams {
	o.SetOrder(order)
	return o
}

// SetOrder adds the order to the orders price params
func (o *OrdersPriceParams) SetOrder(order *models.Order) {
	o.Order = order
}

// WriteToRequest writes these params to a swagger request
func (o *OrdersPriceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Order != nil {
		if err := r.SetBodyParam(o.Order); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
