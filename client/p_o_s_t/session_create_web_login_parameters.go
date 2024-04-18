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

// NewSessionCreateWebLoginParams creates a new SessionCreateWebLoginParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSessionCreateWebLoginParams() *SessionCreateWebLoginParams {
	return &SessionCreateWebLoginParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSessionCreateWebLoginParamsWithTimeout creates a new SessionCreateWebLoginParams object
// with the ability to set a timeout on a request.
func NewSessionCreateWebLoginParamsWithTimeout(timeout time.Duration) *SessionCreateWebLoginParams {
	return &SessionCreateWebLoginParams{
		timeout: timeout,
	}
}

// NewSessionCreateWebLoginParamsWithContext creates a new SessionCreateWebLoginParams object
// with the ability to set a context for a request.
func NewSessionCreateWebLoginParamsWithContext(ctx context.Context) *SessionCreateWebLoginParams {
	return &SessionCreateWebLoginParams{
		Context: ctx,
	}
}

// NewSessionCreateWebLoginParamsWithHTTPClient creates a new SessionCreateWebLoginParams object
// with the ability to set a custom HTTPClient for a request.
func NewSessionCreateWebLoginParamsWithHTTPClient(client *http.Client) *SessionCreateWebLoginParams {
	return &SessionCreateWebLoginParams{
		HTTPClient: client,
	}
}

/*
SessionCreateWebLoginParams contains all the parameters to send to the API endpoint

	for the session create web login operation.

	Typically these are written to a http.Request.
*/
type SessionCreateWebLoginParams struct {

	// SessionKey.
	SessionKey string

	// SessionWebLogin.
	SessionWebLogin *models.SessionWebLogin

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the session create web login params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SessionCreateWebLoginParams) WithDefaults() *SessionCreateWebLoginParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the session create web login params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SessionCreateWebLoginParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the session create web login params
func (o *SessionCreateWebLoginParams) WithTimeout(timeout time.Duration) *SessionCreateWebLoginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the session create web login params
func (o *SessionCreateWebLoginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the session create web login params
func (o *SessionCreateWebLoginParams) WithContext(ctx context.Context) *SessionCreateWebLoginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the session create web login params
func (o *SessionCreateWebLoginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the session create web login params
func (o *SessionCreateWebLoginParams) WithHTTPClient(client *http.Client) *SessionCreateWebLoginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the session create web login params
func (o *SessionCreateWebLoginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSessionKey adds the sessionKey to the session create web login params
func (o *SessionCreateWebLoginParams) WithSessionKey(sessionKey string) *SessionCreateWebLoginParams {
	o.SetSessionKey(sessionKey)
	return o
}

// SetSessionKey adds the sessionKey to the session create web login params
func (o *SessionCreateWebLoginParams) SetSessionKey(sessionKey string) {
	o.SessionKey = sessionKey
}

// WithSessionWebLogin adds the sessionWebLogin to the session create web login params
func (o *SessionCreateWebLoginParams) WithSessionWebLogin(sessionWebLogin *models.SessionWebLogin) *SessionCreateWebLoginParams {
	o.SetSessionWebLogin(sessionWebLogin)
	return o
}

// SetSessionWebLogin adds the sessionWebLogin to the session create web login params
func (o *SessionCreateWebLoginParams) SetSessionWebLogin(sessionWebLogin *models.SessionWebLogin) {
	o.SessionWebLogin = sessionWebLogin
}

// WriteToRequest writes these params to a swagger request
func (o *SessionCreateWebLoginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param sessionKey
	if err := r.SetPathParam("sessionKey", o.SessionKey); err != nil {
		return err
	}
	if o.SessionWebLogin != nil {
		if err := r.SetBodyParam(o.SessionWebLogin); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}