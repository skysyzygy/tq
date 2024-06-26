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

// NewOrderCategoriesCreateParams creates a new OrderCategoriesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrderCategoriesCreateParams() *OrderCategoriesCreateParams {
	return &OrderCategoriesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrderCategoriesCreateParamsWithTimeout creates a new OrderCategoriesCreateParams object
// with the ability to set a timeout on a request.
func NewOrderCategoriesCreateParamsWithTimeout(timeout time.Duration) *OrderCategoriesCreateParams {
	return &OrderCategoriesCreateParams{
		timeout: timeout,
	}
}

// NewOrderCategoriesCreateParamsWithContext creates a new OrderCategoriesCreateParams object
// with the ability to set a context for a request.
func NewOrderCategoriesCreateParamsWithContext(ctx context.Context) *OrderCategoriesCreateParams {
	return &OrderCategoriesCreateParams{
		Context: ctx,
	}
}

// NewOrderCategoriesCreateParamsWithHTTPClient creates a new OrderCategoriesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrderCategoriesCreateParamsWithHTTPClient(client *http.Client) *OrderCategoriesCreateParams {
	return &OrderCategoriesCreateParams{
		HTTPClient: client,
	}
}

/*
OrderCategoriesCreateParams contains all the parameters to send to the API endpoint

	for the order categories create operation.

	Typically these are written to a http.Request.
*/
type OrderCategoriesCreateParams struct {

	// Data.
	Data *models.OrderCategory

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the order categories create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrderCategoriesCreateParams) WithDefaults() *OrderCategoriesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the order categories create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrderCategoriesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the order categories create params
func (o *OrderCategoriesCreateParams) WithTimeout(timeout time.Duration) *OrderCategoriesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the order categories create params
func (o *OrderCategoriesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the order categories create params
func (o *OrderCategoriesCreateParams) WithContext(ctx context.Context) *OrderCategoriesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the order categories create params
func (o *OrderCategoriesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the order categories create params
func (o *OrderCategoriesCreateParams) WithHTTPClient(client *http.Client) *OrderCategoriesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the order categories create params
func (o *OrderCategoriesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the order categories create params
func (o *OrderCategoriesCreateParams) WithData(data *models.OrderCategory) *OrderCategoriesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the order categories create params
func (o *OrderCategoriesCreateParams) SetData(data *models.OrderCategory) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *OrderCategoriesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
