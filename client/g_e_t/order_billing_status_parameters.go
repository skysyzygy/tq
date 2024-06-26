// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

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

// NewOrderBillingStatusParams creates a new OrderBillingStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrderBillingStatusParams() *OrderBillingStatusParams {
	return &OrderBillingStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrderBillingStatusParamsWithTimeout creates a new OrderBillingStatusParams object
// with the ability to set a timeout on a request.
func NewOrderBillingStatusParamsWithTimeout(timeout time.Duration) *OrderBillingStatusParams {
	return &OrderBillingStatusParams{
		timeout: timeout,
	}
}

// NewOrderBillingStatusParamsWithContext creates a new OrderBillingStatusParams object
// with the ability to set a context for a request.
func NewOrderBillingStatusParamsWithContext(ctx context.Context) *OrderBillingStatusParams {
	return &OrderBillingStatusParams{
		Context: ctx,
	}
}

// NewOrderBillingStatusParamsWithHTTPClient creates a new OrderBillingStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrderBillingStatusParamsWithHTTPClient(client *http.Client) *OrderBillingStatusParams {
	return &OrderBillingStatusParams{
		HTTPClient: client,
	}
}

/*
OrderBillingStatusParams contains all the parameters to send to the API endpoint

	for the order billing status operation.

	Typically these are written to a http.Request.
*/
type OrderBillingStatusParams struct {

	// OrderBillingID.
	OrderBillingID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the order billing status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrderBillingStatusParams) WithDefaults() *OrderBillingStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the order billing status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrderBillingStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the order billing status params
func (o *OrderBillingStatusParams) WithTimeout(timeout time.Duration) *OrderBillingStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the order billing status params
func (o *OrderBillingStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the order billing status params
func (o *OrderBillingStatusParams) WithContext(ctx context.Context) *OrderBillingStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the order billing status params
func (o *OrderBillingStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the order billing status params
func (o *OrderBillingStatusParams) WithHTTPClient(client *http.Client) *OrderBillingStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the order billing status params
func (o *OrderBillingStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderBillingID adds the orderBillingID to the order billing status params
func (o *OrderBillingStatusParams) WithOrderBillingID(orderBillingID string) *OrderBillingStatusParams {
	o.SetOrderBillingID(orderBillingID)
	return o
}

// SetOrderBillingID adds the orderBillingId to the order billing status params
func (o *OrderBillingStatusParams) SetOrderBillingID(orderBillingID string) {
	o.OrderBillingID = orderBillingID
}

// WriteToRequest writes these params to a swagger request
func (o *OrderBillingStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orderBillingId
	if err := r.SetPathParam("orderBillingId", o.OrderBillingID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
