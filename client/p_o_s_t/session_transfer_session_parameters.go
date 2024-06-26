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

// NewSessionTransferSessionParams creates a new SessionTransferSessionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSessionTransferSessionParams() *SessionTransferSessionParams {
	return &SessionTransferSessionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSessionTransferSessionParamsWithTimeout creates a new SessionTransferSessionParams object
// with the ability to set a timeout on a request.
func NewSessionTransferSessionParamsWithTimeout(timeout time.Duration) *SessionTransferSessionParams {
	return &SessionTransferSessionParams{
		timeout: timeout,
	}
}

// NewSessionTransferSessionParamsWithContext creates a new SessionTransferSessionParams object
// with the ability to set a context for a request.
func NewSessionTransferSessionParamsWithContext(ctx context.Context) *SessionTransferSessionParams {
	return &SessionTransferSessionParams{
		Context: ctx,
	}
}

// NewSessionTransferSessionParamsWithHTTPClient creates a new SessionTransferSessionParams object
// with the ability to set a custom HTTPClient for a request.
func NewSessionTransferSessionParamsWithHTTPClient(client *http.Client) *SessionTransferSessionParams {
	return &SessionTransferSessionParams{
		HTTPClient: client,
	}
}

/*
SessionTransferSessionParams contains all the parameters to send to the API endpoint

	for the session transfer session operation.

	Typically these are written to a http.Request.
*/
type SessionTransferSessionParams struct {

	// Request.
	Request *models.TransferSessionRequest

	// SessionKey.
	SessionKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the session transfer session params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SessionTransferSessionParams) WithDefaults() *SessionTransferSessionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the session transfer session params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SessionTransferSessionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the session transfer session params
func (o *SessionTransferSessionParams) WithTimeout(timeout time.Duration) *SessionTransferSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the session transfer session params
func (o *SessionTransferSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the session transfer session params
func (o *SessionTransferSessionParams) WithContext(ctx context.Context) *SessionTransferSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the session transfer session params
func (o *SessionTransferSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the session transfer session params
func (o *SessionTransferSessionParams) WithHTTPClient(client *http.Client) *SessionTransferSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the session transfer session params
func (o *SessionTransferSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the session transfer session params
func (o *SessionTransferSessionParams) WithRequest(request *models.TransferSessionRequest) *SessionTransferSessionParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the session transfer session params
func (o *SessionTransferSessionParams) SetRequest(request *models.TransferSessionRequest) {
	o.Request = request
}

// WithSessionKey adds the sessionKey to the session transfer session params
func (o *SessionTransferSessionParams) WithSessionKey(sessionKey string) *SessionTransferSessionParams {
	o.SetSessionKey(sessionKey)
	return o
}

// SetSessionKey adds the sessionKey to the session transfer session params
func (o *SessionTransferSessionParams) SetSessionKey(sessionKey string) {
	o.SessionKey = sessionKey
}

// WriteToRequest writes these params to a swagger request
func (o *SessionTransferSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
