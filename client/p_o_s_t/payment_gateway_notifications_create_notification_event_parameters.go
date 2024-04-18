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

// NewPaymentGatewayNotificationsCreateNotificationEventParams creates a new PaymentGatewayNotificationsCreateNotificationEventParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPaymentGatewayNotificationsCreateNotificationEventParams() *PaymentGatewayNotificationsCreateNotificationEventParams {
	return &PaymentGatewayNotificationsCreateNotificationEventParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPaymentGatewayNotificationsCreateNotificationEventParamsWithTimeout creates a new PaymentGatewayNotificationsCreateNotificationEventParams object
// with the ability to set a timeout on a request.
func NewPaymentGatewayNotificationsCreateNotificationEventParamsWithTimeout(timeout time.Duration) *PaymentGatewayNotificationsCreateNotificationEventParams {
	return &PaymentGatewayNotificationsCreateNotificationEventParams{
		timeout: timeout,
	}
}

// NewPaymentGatewayNotificationsCreateNotificationEventParamsWithContext creates a new PaymentGatewayNotificationsCreateNotificationEventParams object
// with the ability to set a context for a request.
func NewPaymentGatewayNotificationsCreateNotificationEventParamsWithContext(ctx context.Context) *PaymentGatewayNotificationsCreateNotificationEventParams {
	return &PaymentGatewayNotificationsCreateNotificationEventParams{
		Context: ctx,
	}
}

// NewPaymentGatewayNotificationsCreateNotificationEventParamsWithHTTPClient creates a new PaymentGatewayNotificationsCreateNotificationEventParams object
// with the ability to set a custom HTTPClient for a request.
func NewPaymentGatewayNotificationsCreateNotificationEventParamsWithHTTPClient(client *http.Client) *PaymentGatewayNotificationsCreateNotificationEventParams {
	return &PaymentGatewayNotificationsCreateNotificationEventParams{
		HTTPClient: client,
	}
}

/*
PaymentGatewayNotificationsCreateNotificationEventParams contains all the parameters to send to the API endpoint

	for the payment gateway notifications create notification event operation.

	Typically these are written to a http.Request.
*/
type PaymentGatewayNotificationsCreateNotificationEventParams struct {

	// NotificationEvent.
	NotificationEvent *models.NotificationEvent

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the payment gateway notifications create notification event params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PaymentGatewayNotificationsCreateNotificationEventParams) WithDefaults() *PaymentGatewayNotificationsCreateNotificationEventParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the payment gateway notifications create notification event params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PaymentGatewayNotificationsCreateNotificationEventParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the payment gateway notifications create notification event params
func (o *PaymentGatewayNotificationsCreateNotificationEventParams) WithTimeout(timeout time.Duration) *PaymentGatewayNotificationsCreateNotificationEventParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the payment gateway notifications create notification event params
func (o *PaymentGatewayNotificationsCreateNotificationEventParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the payment gateway notifications create notification event params
func (o *PaymentGatewayNotificationsCreateNotificationEventParams) WithContext(ctx context.Context) *PaymentGatewayNotificationsCreateNotificationEventParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the payment gateway notifications create notification event params
func (o *PaymentGatewayNotificationsCreateNotificationEventParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the payment gateway notifications create notification event params
func (o *PaymentGatewayNotificationsCreateNotificationEventParams) WithHTTPClient(client *http.Client) *PaymentGatewayNotificationsCreateNotificationEventParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the payment gateway notifications create notification event params
func (o *PaymentGatewayNotificationsCreateNotificationEventParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNotificationEvent adds the notificationEvent to the payment gateway notifications create notification event params
func (o *PaymentGatewayNotificationsCreateNotificationEventParams) WithNotificationEvent(notificationEvent *models.NotificationEvent) *PaymentGatewayNotificationsCreateNotificationEventParams {
	o.SetNotificationEvent(notificationEvent)
	return o
}

// SetNotificationEvent adds the notificationEvent to the payment gateway notifications create notification event params
func (o *PaymentGatewayNotificationsCreateNotificationEventParams) SetNotificationEvent(notificationEvent *models.NotificationEvent) {
	o.NotificationEvent = notificationEvent
}

// WriteToRequest writes these params to a swagger request
func (o *PaymentGatewayNotificationsCreateNotificationEventParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.NotificationEvent != nil {
		if err := r.SetBodyParam(o.NotificationEvent); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}