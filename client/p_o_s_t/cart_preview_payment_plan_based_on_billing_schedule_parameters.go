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

// NewCartPreviewPaymentPlanBasedOnBillingScheduleParams creates a new CartPreviewPaymentPlanBasedOnBillingScheduleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCartPreviewPaymentPlanBasedOnBillingScheduleParams() *CartPreviewPaymentPlanBasedOnBillingScheduleParams {
	return &CartPreviewPaymentPlanBasedOnBillingScheduleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCartPreviewPaymentPlanBasedOnBillingScheduleParamsWithTimeout creates a new CartPreviewPaymentPlanBasedOnBillingScheduleParams object
// with the ability to set a timeout on a request.
func NewCartPreviewPaymentPlanBasedOnBillingScheduleParamsWithTimeout(timeout time.Duration) *CartPreviewPaymentPlanBasedOnBillingScheduleParams {
	return &CartPreviewPaymentPlanBasedOnBillingScheduleParams{
		timeout: timeout,
	}
}

// NewCartPreviewPaymentPlanBasedOnBillingScheduleParamsWithContext creates a new CartPreviewPaymentPlanBasedOnBillingScheduleParams object
// with the ability to set a context for a request.
func NewCartPreviewPaymentPlanBasedOnBillingScheduleParamsWithContext(ctx context.Context) *CartPreviewPaymentPlanBasedOnBillingScheduleParams {
	return &CartPreviewPaymentPlanBasedOnBillingScheduleParams{
		Context: ctx,
	}
}

// NewCartPreviewPaymentPlanBasedOnBillingScheduleParamsWithHTTPClient creates a new CartPreviewPaymentPlanBasedOnBillingScheduleParams object
// with the ability to set a custom HTTPClient for a request.
func NewCartPreviewPaymentPlanBasedOnBillingScheduleParamsWithHTTPClient(client *http.Client) *CartPreviewPaymentPlanBasedOnBillingScheduleParams {
	return &CartPreviewPaymentPlanBasedOnBillingScheduleParams{
		HTTPClient: client,
	}
}

/*
CartPreviewPaymentPlanBasedOnBillingScheduleParams contains all the parameters to send to the API endpoint

	for the cart preview payment plan based on billing schedule operation.

	Typically these are written to a http.Request.
*/
type CartPreviewPaymentPlanBasedOnBillingScheduleParams struct {

	// Request.
	Request *models.PaymentPlanSchedulePreviewRequest

	// SessionKey.
	SessionKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cart preview payment plan based on billing schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartPreviewPaymentPlanBasedOnBillingScheduleParams) WithDefaults() *CartPreviewPaymentPlanBasedOnBillingScheduleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cart preview payment plan based on billing schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartPreviewPaymentPlanBasedOnBillingScheduleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cart preview payment plan based on billing schedule params
func (o *CartPreviewPaymentPlanBasedOnBillingScheduleParams) WithTimeout(timeout time.Duration) *CartPreviewPaymentPlanBasedOnBillingScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cart preview payment plan based on billing schedule params
func (o *CartPreviewPaymentPlanBasedOnBillingScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cart preview payment plan based on billing schedule params
func (o *CartPreviewPaymentPlanBasedOnBillingScheduleParams) WithContext(ctx context.Context) *CartPreviewPaymentPlanBasedOnBillingScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cart preview payment plan based on billing schedule params
func (o *CartPreviewPaymentPlanBasedOnBillingScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cart preview payment plan based on billing schedule params
func (o *CartPreviewPaymentPlanBasedOnBillingScheduleParams) WithHTTPClient(client *http.Client) *CartPreviewPaymentPlanBasedOnBillingScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cart preview payment plan based on billing schedule params
func (o *CartPreviewPaymentPlanBasedOnBillingScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the cart preview payment plan based on billing schedule params
func (o *CartPreviewPaymentPlanBasedOnBillingScheduleParams) WithRequest(request *models.PaymentPlanSchedulePreviewRequest) *CartPreviewPaymentPlanBasedOnBillingScheduleParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the cart preview payment plan based on billing schedule params
func (o *CartPreviewPaymentPlanBasedOnBillingScheduleParams) SetRequest(request *models.PaymentPlanSchedulePreviewRequest) {
	o.Request = request
}

// WithSessionKey adds the sessionKey to the cart preview payment plan based on billing schedule params
func (o *CartPreviewPaymentPlanBasedOnBillingScheduleParams) WithSessionKey(sessionKey string) *CartPreviewPaymentPlanBasedOnBillingScheduleParams {
	o.SetSessionKey(sessionKey)
	return o
}

// SetSessionKey adds the sessionKey to the cart preview payment plan based on billing schedule params
func (o *CartPreviewPaymentPlanBasedOnBillingScheduleParams) SetSessionKey(sessionKey string) {
	o.SessionKey = sessionKey
}

// WriteToRequest writes these params to a swagger request
func (o *CartPreviewPaymentPlanBasedOnBillingScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
