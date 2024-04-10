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

// NewSessionLoginUsingTokenParams creates a new SessionLoginUsingTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSessionLoginUsingTokenParams() *SessionLoginUsingTokenParams {
	return &SessionLoginUsingTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSessionLoginUsingTokenParamsWithTimeout creates a new SessionLoginUsingTokenParams object
// with the ability to set a timeout on a request.
func NewSessionLoginUsingTokenParamsWithTimeout(timeout time.Duration) *SessionLoginUsingTokenParams {
	return &SessionLoginUsingTokenParams{
		timeout: timeout,
	}
}

// NewSessionLoginUsingTokenParamsWithContext creates a new SessionLoginUsingTokenParams object
// with the ability to set a context for a request.
func NewSessionLoginUsingTokenParamsWithContext(ctx context.Context) *SessionLoginUsingTokenParams {
	return &SessionLoginUsingTokenParams{
		Context: ctx,
	}
}

// NewSessionLoginUsingTokenParamsWithHTTPClient creates a new SessionLoginUsingTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewSessionLoginUsingTokenParamsWithHTTPClient(client *http.Client) *SessionLoginUsingTokenParams {
	return &SessionLoginUsingTokenParams{
		HTTPClient: client,
	}
}

/*
SessionLoginUsingTokenParams contains all the parameters to send to the API endpoint

	for the session login using token operation.

	Typically these are written to a http.Request.
*/
type SessionLoginUsingTokenParams struct {

	// Request.
	Request *models.LoginTokenRequest

	// SessionKey.
	SessionKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the session login using token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SessionLoginUsingTokenParams) WithDefaults() *SessionLoginUsingTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the session login using token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SessionLoginUsingTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the session login using token params
func (o *SessionLoginUsingTokenParams) WithTimeout(timeout time.Duration) *SessionLoginUsingTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the session login using token params
func (o *SessionLoginUsingTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the session login using token params
func (o *SessionLoginUsingTokenParams) WithContext(ctx context.Context) *SessionLoginUsingTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the session login using token params
func (o *SessionLoginUsingTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the session login using token params
func (o *SessionLoginUsingTokenParams) WithHTTPClient(client *http.Client) *SessionLoginUsingTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the session login using token params
func (o *SessionLoginUsingTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the session login using token params
func (o *SessionLoginUsingTokenParams) WithRequest(request *models.LoginTokenRequest) *SessionLoginUsingTokenParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the session login using token params
func (o *SessionLoginUsingTokenParams) SetRequest(request *models.LoginTokenRequest) {
	o.Request = request
}

// WithSessionKey adds the sessionKey to the session login using token params
func (o *SessionLoginUsingTokenParams) WithSessionKey(sessionKey string) *SessionLoginUsingTokenParams {
	o.SetSessionKey(sessionKey)
	return o
}

// SetSessionKey adds the sessionKey to the session login using token params
func (o *SessionLoginUsingTokenParams) SetSessionKey(sessionKey string) {
	o.SessionKey = sessionKey
}

// WriteToRequest writes these params to a swagger request
func (o *SessionLoginUsingTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
