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

// NewCartRemoveOnAccountParams creates a new CartRemoveOnAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCartRemoveOnAccountParams() *CartRemoveOnAccountParams {
	return &CartRemoveOnAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCartRemoveOnAccountParamsWithTimeout creates a new CartRemoveOnAccountParams object
// with the ability to set a timeout on a request.
func NewCartRemoveOnAccountParamsWithTimeout(timeout time.Duration) *CartRemoveOnAccountParams {
	return &CartRemoveOnAccountParams{
		timeout: timeout,
	}
}

// NewCartRemoveOnAccountParamsWithContext creates a new CartRemoveOnAccountParams object
// with the ability to set a context for a request.
func NewCartRemoveOnAccountParamsWithContext(ctx context.Context) *CartRemoveOnAccountParams {
	return &CartRemoveOnAccountParams{
		Context: ctx,
	}
}

// NewCartRemoveOnAccountParamsWithHTTPClient creates a new CartRemoveOnAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewCartRemoveOnAccountParamsWithHTTPClient(client *http.Client) *CartRemoveOnAccountParams {
	return &CartRemoveOnAccountParams{
		HTTPClient: client,
	}
}

/*
CartRemoveOnAccountParams contains all the parameters to send to the API endpoint

	for the cart remove on account operation.

	Typically these are written to a http.Request.
*/
type CartRemoveOnAccountParams struct {

	// PaymentID.
	PaymentID string

	// SessionKey.
	SessionKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cart remove on account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartRemoveOnAccountParams) WithDefaults() *CartRemoveOnAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cart remove on account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartRemoveOnAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cart remove on account params
func (o *CartRemoveOnAccountParams) WithTimeout(timeout time.Duration) *CartRemoveOnAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cart remove on account params
func (o *CartRemoveOnAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cart remove on account params
func (o *CartRemoveOnAccountParams) WithContext(ctx context.Context) *CartRemoveOnAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cart remove on account params
func (o *CartRemoveOnAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cart remove on account params
func (o *CartRemoveOnAccountParams) WithHTTPClient(client *http.Client) *CartRemoveOnAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cart remove on account params
func (o *CartRemoveOnAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPaymentID adds the paymentID to the cart remove on account params
func (o *CartRemoveOnAccountParams) WithPaymentID(paymentID string) *CartRemoveOnAccountParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the cart remove on account params
func (o *CartRemoveOnAccountParams) SetPaymentID(paymentID string) {
	o.PaymentID = paymentID
}

// WithSessionKey adds the sessionKey to the cart remove on account params
func (o *CartRemoveOnAccountParams) WithSessionKey(sessionKey string) *CartRemoveOnAccountParams {
	o.SetSessionKey(sessionKey)
	return o
}

// SetSessionKey adds the sessionKey to the cart remove on account params
func (o *CartRemoveOnAccountParams) SetSessionKey(sessionKey string) {
	o.SessionKey = sessionKey
}

// WriteToRequest writes these params to a swagger request
func (o *CartRemoveOnAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param paymentId
	if err := r.SetPathParam("paymentId", o.PaymentID); err != nil {
		return err
	}

	// path param sessionKey
	if err := r.SetPathParam("sessionKey", o.SessionKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
