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

// NewOrderCategoriesUpdateParams creates a new OrderCategoriesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrderCategoriesUpdateParams() *OrderCategoriesUpdateParams {
	return &OrderCategoriesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrderCategoriesUpdateParamsWithTimeout creates a new OrderCategoriesUpdateParams object
// with the ability to set a timeout on a request.
func NewOrderCategoriesUpdateParamsWithTimeout(timeout time.Duration) *OrderCategoriesUpdateParams {
	return &OrderCategoriesUpdateParams{
		timeout: timeout,
	}
}

// NewOrderCategoriesUpdateParamsWithContext creates a new OrderCategoriesUpdateParams object
// with the ability to set a context for a request.
func NewOrderCategoriesUpdateParamsWithContext(ctx context.Context) *OrderCategoriesUpdateParams {
	return &OrderCategoriesUpdateParams{
		Context: ctx,
	}
}

// NewOrderCategoriesUpdateParamsWithHTTPClient creates a new OrderCategoriesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrderCategoriesUpdateParamsWithHTTPClient(client *http.Client) *OrderCategoriesUpdateParams {
	return &OrderCategoriesUpdateParams{
		HTTPClient: client,
	}
}

/*
OrderCategoriesUpdateParams contains all the parameters to send to the API endpoint

	for the order categories update operation.

	Typically these are written to a http.Request.
*/
type OrderCategoriesUpdateParams struct {

	// Data.
	Data *models.OrderCategory

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the order categories update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrderCategoriesUpdateParams) WithDefaults() *OrderCategoriesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the order categories update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrderCategoriesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the order categories update params
func (o *OrderCategoriesUpdateParams) WithTimeout(timeout time.Duration) *OrderCategoriesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the order categories update params
func (o *OrderCategoriesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the order categories update params
func (o *OrderCategoriesUpdateParams) WithContext(ctx context.Context) *OrderCategoriesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the order categories update params
func (o *OrderCategoriesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the order categories update params
func (o *OrderCategoriesUpdateParams) WithHTTPClient(client *http.Client) *OrderCategoriesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the order categories update params
func (o *OrderCategoriesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the order categories update params
func (o *OrderCategoriesUpdateParams) WithData(data *models.OrderCategory) *OrderCategoriesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the order categories update params
func (o *OrderCategoriesUpdateParams) SetData(data *models.OrderCategory) {
	o.Data = data
}

// WithID adds the id to the order categories update params
func (o *OrderCategoriesUpdateParams) WithID(id string) *OrderCategoriesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the order categories update params
func (o *OrderCategoriesUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *OrderCategoriesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
