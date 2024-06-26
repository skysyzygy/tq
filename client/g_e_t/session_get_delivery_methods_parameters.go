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

// NewSessionGetDeliveryMethodsParams creates a new SessionGetDeliveryMethodsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSessionGetDeliveryMethodsParams() *SessionGetDeliveryMethodsParams {
	return &SessionGetDeliveryMethodsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSessionGetDeliveryMethodsParamsWithTimeout creates a new SessionGetDeliveryMethodsParams object
// with the ability to set a timeout on a request.
func NewSessionGetDeliveryMethodsParamsWithTimeout(timeout time.Duration) *SessionGetDeliveryMethodsParams {
	return &SessionGetDeliveryMethodsParams{
		timeout: timeout,
	}
}

// NewSessionGetDeliveryMethodsParamsWithContext creates a new SessionGetDeliveryMethodsParams object
// with the ability to set a context for a request.
func NewSessionGetDeliveryMethodsParamsWithContext(ctx context.Context) *SessionGetDeliveryMethodsParams {
	return &SessionGetDeliveryMethodsParams{
		Context: ctx,
	}
}

// NewSessionGetDeliveryMethodsParamsWithHTTPClient creates a new SessionGetDeliveryMethodsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSessionGetDeliveryMethodsParamsWithHTTPClient(client *http.Client) *SessionGetDeliveryMethodsParams {
	return &SessionGetDeliveryMethodsParams{
		HTTPClient: client,
	}
}

/*
SessionGetDeliveryMethodsParams contains all the parameters to send to the API endpoint

	for the session get delivery methods operation.

	Typically these are written to a http.Request.
*/
type SessionGetDeliveryMethodsParams struct {

	// SessionKey.
	SessionKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the session get delivery methods params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SessionGetDeliveryMethodsParams) WithDefaults() *SessionGetDeliveryMethodsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the session get delivery methods params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SessionGetDeliveryMethodsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the session get delivery methods params
func (o *SessionGetDeliveryMethodsParams) WithTimeout(timeout time.Duration) *SessionGetDeliveryMethodsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the session get delivery methods params
func (o *SessionGetDeliveryMethodsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the session get delivery methods params
func (o *SessionGetDeliveryMethodsParams) WithContext(ctx context.Context) *SessionGetDeliveryMethodsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the session get delivery methods params
func (o *SessionGetDeliveryMethodsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the session get delivery methods params
func (o *SessionGetDeliveryMethodsParams) WithHTTPClient(client *http.Client) *SessionGetDeliveryMethodsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the session get delivery methods params
func (o *SessionGetDeliveryMethodsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSessionKey adds the sessionKey to the session get delivery methods params
func (o *SessionGetDeliveryMethodsParams) WithSessionKey(sessionKey string) *SessionGetDeliveryMethodsParams {
	o.SetSessionKey(sessionKey)
	return o
}

// SetSessionKey adds the sessionKey to the session get delivery methods params
func (o *SessionGetDeliveryMethodsParams) SetSessionKey(sessionKey string) {
	o.SessionKey = sessionKey
}

// WriteToRequest writes these params to a swagger request
func (o *SessionGetDeliveryMethodsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param sessionKey
	if err := r.SetPathParam("sessionKey", o.SessionKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
