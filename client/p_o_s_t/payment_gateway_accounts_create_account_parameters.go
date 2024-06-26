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

// NewPaymentGatewayAccountsCreateAccountParams creates a new PaymentGatewayAccountsCreateAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPaymentGatewayAccountsCreateAccountParams() *PaymentGatewayAccountsCreateAccountParams {
	return &PaymentGatewayAccountsCreateAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPaymentGatewayAccountsCreateAccountParamsWithTimeout creates a new PaymentGatewayAccountsCreateAccountParams object
// with the ability to set a timeout on a request.
func NewPaymentGatewayAccountsCreateAccountParamsWithTimeout(timeout time.Duration) *PaymentGatewayAccountsCreateAccountParams {
	return &PaymentGatewayAccountsCreateAccountParams{
		timeout: timeout,
	}
}

// NewPaymentGatewayAccountsCreateAccountParamsWithContext creates a new PaymentGatewayAccountsCreateAccountParams object
// with the ability to set a context for a request.
func NewPaymentGatewayAccountsCreateAccountParamsWithContext(ctx context.Context) *PaymentGatewayAccountsCreateAccountParams {
	return &PaymentGatewayAccountsCreateAccountParams{
		Context: ctx,
	}
}

// NewPaymentGatewayAccountsCreateAccountParamsWithHTTPClient creates a new PaymentGatewayAccountsCreateAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewPaymentGatewayAccountsCreateAccountParamsWithHTTPClient(client *http.Client) *PaymentGatewayAccountsCreateAccountParams {
	return &PaymentGatewayAccountsCreateAccountParams{
		HTTPClient: client,
	}
}

/*
PaymentGatewayAccountsCreateAccountParams contains all the parameters to send to the API endpoint

	for the payment gateway accounts create account operation.

	Typically these are written to a http.Request.
*/
type PaymentGatewayAccountsCreateAccountParams struct {

	// Request.
	Request *models.AccountReferenceNumberRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the payment gateway accounts create account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PaymentGatewayAccountsCreateAccountParams) WithDefaults() *PaymentGatewayAccountsCreateAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the payment gateway accounts create account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PaymentGatewayAccountsCreateAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the payment gateway accounts create account params
func (o *PaymentGatewayAccountsCreateAccountParams) WithTimeout(timeout time.Duration) *PaymentGatewayAccountsCreateAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the payment gateway accounts create account params
func (o *PaymentGatewayAccountsCreateAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the payment gateway accounts create account params
func (o *PaymentGatewayAccountsCreateAccountParams) WithContext(ctx context.Context) *PaymentGatewayAccountsCreateAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the payment gateway accounts create account params
func (o *PaymentGatewayAccountsCreateAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the payment gateway accounts create account params
func (o *PaymentGatewayAccountsCreateAccountParams) WithHTTPClient(client *http.Client) *PaymentGatewayAccountsCreateAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the payment gateway accounts create account params
func (o *PaymentGatewayAccountsCreateAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the payment gateway accounts create account params
func (o *PaymentGatewayAccountsCreateAccountParams) WithRequest(request *models.AccountReferenceNumberRequest) *PaymentGatewayAccountsCreateAccountParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the payment gateway accounts create account params
func (o *PaymentGatewayAccountsCreateAccountParams) SetRequest(request *models.AccountReferenceNumberRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PaymentGatewayAccountsCreateAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
