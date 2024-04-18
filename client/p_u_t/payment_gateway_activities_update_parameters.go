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

// NewPaymentGatewayActivitiesUpdateParams creates a new PaymentGatewayActivitiesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPaymentGatewayActivitiesUpdateParams() *PaymentGatewayActivitiesUpdateParams {
	return &PaymentGatewayActivitiesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPaymentGatewayActivitiesUpdateParamsWithTimeout creates a new PaymentGatewayActivitiesUpdateParams object
// with the ability to set a timeout on a request.
func NewPaymentGatewayActivitiesUpdateParamsWithTimeout(timeout time.Duration) *PaymentGatewayActivitiesUpdateParams {
	return &PaymentGatewayActivitiesUpdateParams{
		timeout: timeout,
	}
}

// NewPaymentGatewayActivitiesUpdateParamsWithContext creates a new PaymentGatewayActivitiesUpdateParams object
// with the ability to set a context for a request.
func NewPaymentGatewayActivitiesUpdateParamsWithContext(ctx context.Context) *PaymentGatewayActivitiesUpdateParams {
	return &PaymentGatewayActivitiesUpdateParams{
		Context: ctx,
	}
}

// NewPaymentGatewayActivitiesUpdateParamsWithHTTPClient creates a new PaymentGatewayActivitiesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPaymentGatewayActivitiesUpdateParamsWithHTTPClient(client *http.Client) *PaymentGatewayActivitiesUpdateParams {
	return &PaymentGatewayActivitiesUpdateParams{
		HTTPClient: client,
	}
}

/*
PaymentGatewayActivitiesUpdateParams contains all the parameters to send to the API endpoint

	for the payment gateway activities update operation.

	Typically these are written to a http.Request.
*/
type PaymentGatewayActivitiesUpdateParams struct {

	// PaymentGatewayActivity.
	PaymentGatewayActivity *models.PaymentGatewayActivity

	// PaymentGatewayActivityID.
	PaymentGatewayActivityID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the payment gateway activities update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PaymentGatewayActivitiesUpdateParams) WithDefaults() *PaymentGatewayActivitiesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the payment gateway activities update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PaymentGatewayActivitiesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the payment gateway activities update params
func (o *PaymentGatewayActivitiesUpdateParams) WithTimeout(timeout time.Duration) *PaymentGatewayActivitiesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the payment gateway activities update params
func (o *PaymentGatewayActivitiesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the payment gateway activities update params
func (o *PaymentGatewayActivitiesUpdateParams) WithContext(ctx context.Context) *PaymentGatewayActivitiesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the payment gateway activities update params
func (o *PaymentGatewayActivitiesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the payment gateway activities update params
func (o *PaymentGatewayActivitiesUpdateParams) WithHTTPClient(client *http.Client) *PaymentGatewayActivitiesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the payment gateway activities update params
func (o *PaymentGatewayActivitiesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPaymentGatewayActivity adds the paymentGatewayActivity to the payment gateway activities update params
func (o *PaymentGatewayActivitiesUpdateParams) WithPaymentGatewayActivity(paymentGatewayActivity *models.PaymentGatewayActivity) *PaymentGatewayActivitiesUpdateParams {
	o.SetPaymentGatewayActivity(paymentGatewayActivity)
	return o
}

// SetPaymentGatewayActivity adds the paymentGatewayActivity to the payment gateway activities update params
func (o *PaymentGatewayActivitiesUpdateParams) SetPaymentGatewayActivity(paymentGatewayActivity *models.PaymentGatewayActivity) {
	o.PaymentGatewayActivity = paymentGatewayActivity
}

// WithPaymentGatewayActivityID adds the paymentGatewayActivityID to the payment gateway activities update params
func (o *PaymentGatewayActivitiesUpdateParams) WithPaymentGatewayActivityID(paymentGatewayActivityID string) *PaymentGatewayActivitiesUpdateParams {
	o.SetPaymentGatewayActivityID(paymentGatewayActivityID)
	return o
}

// SetPaymentGatewayActivityID adds the paymentGatewayActivityId to the payment gateway activities update params
func (o *PaymentGatewayActivitiesUpdateParams) SetPaymentGatewayActivityID(paymentGatewayActivityID string) {
	o.PaymentGatewayActivityID = paymentGatewayActivityID
}

// WriteToRequest writes these params to a swagger request
func (o *PaymentGatewayActivitiesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PaymentGatewayActivity != nil {
		if err := r.SetBodyParam(o.PaymentGatewayActivity); err != nil {
			return err
		}
	}

	// path param paymentGatewayActivityId
	if err := r.SetPathParam("paymentGatewayActivityId", o.PaymentGatewayActivityID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}