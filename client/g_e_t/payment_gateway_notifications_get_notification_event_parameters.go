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

// NewPaymentGatewayNotificationsGetNotificationEventParams creates a new PaymentGatewayNotificationsGetNotificationEventParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPaymentGatewayNotificationsGetNotificationEventParams() *PaymentGatewayNotificationsGetNotificationEventParams {
	return &PaymentGatewayNotificationsGetNotificationEventParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPaymentGatewayNotificationsGetNotificationEventParamsWithTimeout creates a new PaymentGatewayNotificationsGetNotificationEventParams object
// with the ability to set a timeout on a request.
func NewPaymentGatewayNotificationsGetNotificationEventParamsWithTimeout(timeout time.Duration) *PaymentGatewayNotificationsGetNotificationEventParams {
	return &PaymentGatewayNotificationsGetNotificationEventParams{
		timeout: timeout,
	}
}

// NewPaymentGatewayNotificationsGetNotificationEventParamsWithContext creates a new PaymentGatewayNotificationsGetNotificationEventParams object
// with the ability to set a context for a request.
func NewPaymentGatewayNotificationsGetNotificationEventParamsWithContext(ctx context.Context) *PaymentGatewayNotificationsGetNotificationEventParams {
	return &PaymentGatewayNotificationsGetNotificationEventParams{
		Context: ctx,
	}
}

// NewPaymentGatewayNotificationsGetNotificationEventParamsWithHTTPClient creates a new PaymentGatewayNotificationsGetNotificationEventParams object
// with the ability to set a custom HTTPClient for a request.
func NewPaymentGatewayNotificationsGetNotificationEventParamsWithHTTPClient(client *http.Client) *PaymentGatewayNotificationsGetNotificationEventParams {
	return &PaymentGatewayNotificationsGetNotificationEventParams{
		HTTPClient: client,
	}
}

/*
PaymentGatewayNotificationsGetNotificationEventParams contains all the parameters to send to the API endpoint

	for the payment gateway notifications get notification event operation.

	Typically these are written to a http.Request.
*/
type PaymentGatewayNotificationsGetNotificationEventParams struct {

	// NotificationEventID.
	NotificationEventID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the payment gateway notifications get notification event params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PaymentGatewayNotificationsGetNotificationEventParams) WithDefaults() *PaymentGatewayNotificationsGetNotificationEventParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the payment gateway notifications get notification event params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PaymentGatewayNotificationsGetNotificationEventParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the payment gateway notifications get notification event params
func (o *PaymentGatewayNotificationsGetNotificationEventParams) WithTimeout(timeout time.Duration) *PaymentGatewayNotificationsGetNotificationEventParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the payment gateway notifications get notification event params
func (o *PaymentGatewayNotificationsGetNotificationEventParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the payment gateway notifications get notification event params
func (o *PaymentGatewayNotificationsGetNotificationEventParams) WithContext(ctx context.Context) *PaymentGatewayNotificationsGetNotificationEventParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the payment gateway notifications get notification event params
func (o *PaymentGatewayNotificationsGetNotificationEventParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the payment gateway notifications get notification event params
func (o *PaymentGatewayNotificationsGetNotificationEventParams) WithHTTPClient(client *http.Client) *PaymentGatewayNotificationsGetNotificationEventParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the payment gateway notifications get notification event params
func (o *PaymentGatewayNotificationsGetNotificationEventParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNotificationEventID adds the notificationEventID to the payment gateway notifications get notification event params
func (o *PaymentGatewayNotificationsGetNotificationEventParams) WithNotificationEventID(notificationEventID string) *PaymentGatewayNotificationsGetNotificationEventParams {
	o.SetNotificationEventID(notificationEventID)
	return o
}

// SetNotificationEventID adds the notificationEventId to the payment gateway notifications get notification event params
func (o *PaymentGatewayNotificationsGetNotificationEventParams) SetNotificationEventID(notificationEventID string) {
	o.NotificationEventID = notificationEventID
}

// WriteToRequest writes these params to a swagger request
func (o *PaymentGatewayNotificationsGetNotificationEventParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param notificationEventId
	if err := r.SetPathParam("notificationEventId", o.NotificationEventID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
