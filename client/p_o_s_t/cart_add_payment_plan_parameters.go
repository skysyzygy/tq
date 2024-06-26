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

// NewCartAddPaymentPlanParams creates a new CartAddPaymentPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCartAddPaymentPlanParams() *CartAddPaymentPlanParams {
	return &CartAddPaymentPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCartAddPaymentPlanParamsWithTimeout creates a new CartAddPaymentPlanParams object
// with the ability to set a timeout on a request.
func NewCartAddPaymentPlanParamsWithTimeout(timeout time.Duration) *CartAddPaymentPlanParams {
	return &CartAddPaymentPlanParams{
		timeout: timeout,
	}
}

// NewCartAddPaymentPlanParamsWithContext creates a new CartAddPaymentPlanParams object
// with the ability to set a context for a request.
func NewCartAddPaymentPlanParamsWithContext(ctx context.Context) *CartAddPaymentPlanParams {
	return &CartAddPaymentPlanParams{
		Context: ctx,
	}
}

// NewCartAddPaymentPlanParamsWithHTTPClient creates a new CartAddPaymentPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewCartAddPaymentPlanParamsWithHTTPClient(client *http.Client) *CartAddPaymentPlanParams {
	return &CartAddPaymentPlanParams{
		HTTPClient: client,
	}
}

/*
CartAddPaymentPlanParams contains all the parameters to send to the API endpoint

	for the cart add payment plan operation.

	Typically these are written to a http.Request.
*/
type CartAddPaymentPlanParams struct {

	// Request.
	Request *models.PaymentPlanRequest

	// SessionKey.
	SessionKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cart add payment plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartAddPaymentPlanParams) WithDefaults() *CartAddPaymentPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cart add payment plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartAddPaymentPlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cart add payment plan params
func (o *CartAddPaymentPlanParams) WithTimeout(timeout time.Duration) *CartAddPaymentPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cart add payment plan params
func (o *CartAddPaymentPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cart add payment plan params
func (o *CartAddPaymentPlanParams) WithContext(ctx context.Context) *CartAddPaymentPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cart add payment plan params
func (o *CartAddPaymentPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cart add payment plan params
func (o *CartAddPaymentPlanParams) WithHTTPClient(client *http.Client) *CartAddPaymentPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cart add payment plan params
func (o *CartAddPaymentPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the cart add payment plan params
func (o *CartAddPaymentPlanParams) WithRequest(request *models.PaymentPlanRequest) *CartAddPaymentPlanParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the cart add payment plan params
func (o *CartAddPaymentPlanParams) SetRequest(request *models.PaymentPlanRequest) {
	o.Request = request
}

// WithSessionKey adds the sessionKey to the cart add payment plan params
func (o *CartAddPaymentPlanParams) WithSessionKey(sessionKey string) *CartAddPaymentPlanParams {
	o.SetSessionKey(sessionKey)
	return o
}

// SetSessionKey adds the sessionKey to the cart add payment plan params
func (o *CartAddPaymentPlanParams) SetSessionKey(sessionKey string) {
	o.SessionKey = sessionKey
}

// WriteToRequest writes these params to a swagger request
func (o *CartAddPaymentPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
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
