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

// NewPaymentMethodGroupsUpdateParams creates a new PaymentMethodGroupsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPaymentMethodGroupsUpdateParams() *PaymentMethodGroupsUpdateParams {
	return &PaymentMethodGroupsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPaymentMethodGroupsUpdateParamsWithTimeout creates a new PaymentMethodGroupsUpdateParams object
// with the ability to set a timeout on a request.
func NewPaymentMethodGroupsUpdateParamsWithTimeout(timeout time.Duration) *PaymentMethodGroupsUpdateParams {
	return &PaymentMethodGroupsUpdateParams{
		timeout: timeout,
	}
}

// NewPaymentMethodGroupsUpdateParamsWithContext creates a new PaymentMethodGroupsUpdateParams object
// with the ability to set a context for a request.
func NewPaymentMethodGroupsUpdateParamsWithContext(ctx context.Context) *PaymentMethodGroupsUpdateParams {
	return &PaymentMethodGroupsUpdateParams{
		Context: ctx,
	}
}

// NewPaymentMethodGroupsUpdateParamsWithHTTPClient creates a new PaymentMethodGroupsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPaymentMethodGroupsUpdateParamsWithHTTPClient(client *http.Client) *PaymentMethodGroupsUpdateParams {
	return &PaymentMethodGroupsUpdateParams{
		HTTPClient: client,
	}
}

/*
PaymentMethodGroupsUpdateParams contains all the parameters to send to the API endpoint

	for the payment method groups update operation.

	Typically these are written to a http.Request.
*/
type PaymentMethodGroupsUpdateParams struct {

	// Data.
	Data *models.PaymentMethodGroup

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the payment method groups update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PaymentMethodGroupsUpdateParams) WithDefaults() *PaymentMethodGroupsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the payment method groups update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PaymentMethodGroupsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the payment method groups update params
func (o *PaymentMethodGroupsUpdateParams) WithTimeout(timeout time.Duration) *PaymentMethodGroupsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the payment method groups update params
func (o *PaymentMethodGroupsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the payment method groups update params
func (o *PaymentMethodGroupsUpdateParams) WithContext(ctx context.Context) *PaymentMethodGroupsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the payment method groups update params
func (o *PaymentMethodGroupsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the payment method groups update params
func (o *PaymentMethodGroupsUpdateParams) WithHTTPClient(client *http.Client) *PaymentMethodGroupsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the payment method groups update params
func (o *PaymentMethodGroupsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the payment method groups update params
func (o *PaymentMethodGroupsUpdateParams) WithData(data *models.PaymentMethodGroup) *PaymentMethodGroupsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the payment method groups update params
func (o *PaymentMethodGroupsUpdateParams) SetData(data *models.PaymentMethodGroup) {
	o.Data = data
}

// WithID adds the id to the payment method groups update params
func (o *PaymentMethodGroupsUpdateParams) WithID(id string) *PaymentMethodGroupsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the payment method groups update params
func (o *PaymentMethodGroupsUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PaymentMethodGroupsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
