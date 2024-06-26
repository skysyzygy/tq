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

// NewPaymentGatewayNotificationsUpdateParams creates a new PaymentGatewayNotificationsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPaymentGatewayNotificationsUpdateParams() *PaymentGatewayNotificationsUpdateParams {
	return &PaymentGatewayNotificationsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPaymentGatewayNotificationsUpdateParamsWithTimeout creates a new PaymentGatewayNotificationsUpdateParams object
// with the ability to set a timeout on a request.
func NewPaymentGatewayNotificationsUpdateParamsWithTimeout(timeout time.Duration) *PaymentGatewayNotificationsUpdateParams {
	return &PaymentGatewayNotificationsUpdateParams{
		timeout: timeout,
	}
}

// NewPaymentGatewayNotificationsUpdateParamsWithContext creates a new PaymentGatewayNotificationsUpdateParams object
// with the ability to set a context for a request.
func NewPaymentGatewayNotificationsUpdateParamsWithContext(ctx context.Context) *PaymentGatewayNotificationsUpdateParams {
	return &PaymentGatewayNotificationsUpdateParams{
		Context: ctx,
	}
}

// NewPaymentGatewayNotificationsUpdateParamsWithHTTPClient creates a new PaymentGatewayNotificationsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPaymentGatewayNotificationsUpdateParamsWithHTTPClient(client *http.Client) *PaymentGatewayNotificationsUpdateParams {
	return &PaymentGatewayNotificationsUpdateParams{
		HTTPClient: client,
	}
}

/*
PaymentGatewayNotificationsUpdateParams contains all the parameters to send to the API endpoint

	for the payment gateway notifications update operation.

	Typically these are written to a http.Request.
*/
type PaymentGatewayNotificationsUpdateParams struct {

	// NotificationEvent.
	NotificationEvent *models.NotificationEvent

	// NotificationEventID.
	NotificationEventID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the payment gateway notifications update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PaymentGatewayNotificationsUpdateParams) WithDefaults() *PaymentGatewayNotificationsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the payment gateway notifications update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PaymentGatewayNotificationsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the payment gateway notifications update params
func (o *PaymentGatewayNotificationsUpdateParams) WithTimeout(timeout time.Duration) *PaymentGatewayNotificationsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the payment gateway notifications update params
func (o *PaymentGatewayNotificationsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the payment gateway notifications update params
func (o *PaymentGatewayNotificationsUpdateParams) WithContext(ctx context.Context) *PaymentGatewayNotificationsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the payment gateway notifications update params
func (o *PaymentGatewayNotificationsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the payment gateway notifications update params
func (o *PaymentGatewayNotificationsUpdateParams) WithHTTPClient(client *http.Client) *PaymentGatewayNotificationsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the payment gateway notifications update params
func (o *PaymentGatewayNotificationsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNotificationEvent adds the notificationEvent to the payment gateway notifications update params
func (o *PaymentGatewayNotificationsUpdateParams) WithNotificationEvent(notificationEvent *models.NotificationEvent) *PaymentGatewayNotificationsUpdateParams {
	o.SetNotificationEvent(notificationEvent)
	return o
}

// SetNotificationEvent adds the notificationEvent to the payment gateway notifications update params
func (o *PaymentGatewayNotificationsUpdateParams) SetNotificationEvent(notificationEvent *models.NotificationEvent) {
	o.NotificationEvent = notificationEvent
}

// WithNotificationEventID adds the notificationEventID to the payment gateway notifications update params
func (o *PaymentGatewayNotificationsUpdateParams) WithNotificationEventID(notificationEventID string) *PaymentGatewayNotificationsUpdateParams {
	o.SetNotificationEventID(notificationEventID)
	return o
}

// SetNotificationEventID adds the notificationEventId to the payment gateway notifications update params
func (o *PaymentGatewayNotificationsUpdateParams) SetNotificationEventID(notificationEventID string) {
	o.NotificationEventID = notificationEventID
}

// WriteToRequest writes these params to a swagger request
func (o *PaymentGatewayNotificationsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.NotificationEvent != nil {
		if err := r.SetBodyParam(o.NotificationEvent); err != nil {
			return err
		}
	}

	// path param notificationEventId
	if err := r.SetPathParam("notificationEventId", o.NotificationEventID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
