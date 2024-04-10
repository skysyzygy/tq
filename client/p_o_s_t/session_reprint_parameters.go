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

// NewSessionReprintParams creates a new SessionReprintParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSessionReprintParams() *SessionReprintParams {
	return &SessionReprintParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSessionReprintParamsWithTimeout creates a new SessionReprintParams object
// with the ability to set a timeout on a request.
func NewSessionReprintParamsWithTimeout(timeout time.Duration) *SessionReprintParams {
	return &SessionReprintParams{
		timeout: timeout,
	}
}

// NewSessionReprintParamsWithContext creates a new SessionReprintParams object
// with the ability to set a context for a request.
func NewSessionReprintParamsWithContext(ctx context.Context) *SessionReprintParams {
	return &SessionReprintParams{
		Context: ctx,
	}
}

// NewSessionReprintParamsWithHTTPClient creates a new SessionReprintParams object
// with the ability to set a custom HTTPClient for a request.
func NewSessionReprintParamsWithHTTPClient(client *http.Client) *SessionReprintParams {
	return &SessionReprintParams{
		HTTPClient: client,
	}
}

/*
SessionReprintParams contains all the parameters to send to the API endpoint

	for the session reprint operation.

	Typically these are written to a http.Request.
*/
type SessionReprintParams struct {

	// Request.
	Request *models.ReprintRequest

	// SessionKey.
	SessionKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the session reprint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SessionReprintParams) WithDefaults() *SessionReprintParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the session reprint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SessionReprintParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the session reprint params
func (o *SessionReprintParams) WithTimeout(timeout time.Duration) *SessionReprintParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the session reprint params
func (o *SessionReprintParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the session reprint params
func (o *SessionReprintParams) WithContext(ctx context.Context) *SessionReprintParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the session reprint params
func (o *SessionReprintParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the session reprint params
func (o *SessionReprintParams) WithHTTPClient(client *http.Client) *SessionReprintParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the session reprint params
func (o *SessionReprintParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the session reprint params
func (o *SessionReprintParams) WithRequest(request *models.ReprintRequest) *SessionReprintParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the session reprint params
func (o *SessionReprintParams) SetRequest(request *models.ReprintRequest) {
	o.Request = request
}

// WithSessionKey adds the sessionKey to the session reprint params
func (o *SessionReprintParams) WithSessionKey(sessionKey string) *SessionReprintParams {
	o.SetSessionKey(sessionKey)
	return o
}

// SetSessionKey adds the sessionKey to the session reprint params
func (o *SessionReprintParams) SetSessionKey(sessionKey string) {
	o.SessionKey = sessionKey
}

// WriteToRequest writes these params to a swagger request
func (o *SessionReprintParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
