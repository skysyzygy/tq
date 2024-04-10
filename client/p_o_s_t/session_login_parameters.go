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

// NewSessionLoginParams creates a new SessionLoginParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSessionLoginParams() *SessionLoginParams {
	return &SessionLoginParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSessionLoginParamsWithTimeout creates a new SessionLoginParams object
// with the ability to set a timeout on a request.
func NewSessionLoginParamsWithTimeout(timeout time.Duration) *SessionLoginParams {
	return &SessionLoginParams{
		timeout: timeout,
	}
}

// NewSessionLoginParamsWithContext creates a new SessionLoginParams object
// with the ability to set a context for a request.
func NewSessionLoginParamsWithContext(ctx context.Context) *SessionLoginParams {
	return &SessionLoginParams{
		Context: ctx,
	}
}

// NewSessionLoginParamsWithHTTPClient creates a new SessionLoginParams object
// with the ability to set a custom HTTPClient for a request.
func NewSessionLoginParamsWithHTTPClient(client *http.Client) *SessionLoginParams {
	return &SessionLoginParams{
		HTTPClient: client,
	}
}

/*
SessionLoginParams contains all the parameters to send to the API endpoint

	for the session login operation.

	Typically these are written to a http.Request.
*/
type SessionLoginParams struct {

	// Request.
	Request *models.LoginRequest

	// SessionKey.
	SessionKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the session login params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SessionLoginParams) WithDefaults() *SessionLoginParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the session login params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SessionLoginParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the session login params
func (o *SessionLoginParams) WithTimeout(timeout time.Duration) *SessionLoginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the session login params
func (o *SessionLoginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the session login params
func (o *SessionLoginParams) WithContext(ctx context.Context) *SessionLoginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the session login params
func (o *SessionLoginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the session login params
func (o *SessionLoginParams) WithHTTPClient(client *http.Client) *SessionLoginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the session login params
func (o *SessionLoginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the session login params
func (o *SessionLoginParams) WithRequest(request *models.LoginRequest) *SessionLoginParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the session login params
func (o *SessionLoginParams) SetRequest(request *models.LoginRequest) {
	o.Request = request
}

// WithSessionKey adds the sessionKey to the session login params
func (o *SessionLoginParams) WithSessionKey(sessionKey string) *SessionLoginParams {
	o.SetSessionKey(sessionKey)
	return o
}

// SetSessionKey adds the sessionKey to the session login params
func (o *SessionLoginParams) SetSessionKey(sessionKey string) {
	o.SessionKey = sessionKey
}

// WriteToRequest writes these params to a swagger request
func (o *SessionLoginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
