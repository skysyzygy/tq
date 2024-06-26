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

// NewSessionCreateSessionParams creates a new SessionCreateSessionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSessionCreateSessionParams() *SessionCreateSessionParams {
	return &SessionCreateSessionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSessionCreateSessionParamsWithTimeout creates a new SessionCreateSessionParams object
// with the ability to set a timeout on a request.
func NewSessionCreateSessionParamsWithTimeout(timeout time.Duration) *SessionCreateSessionParams {
	return &SessionCreateSessionParams{
		timeout: timeout,
	}
}

// NewSessionCreateSessionParamsWithContext creates a new SessionCreateSessionParams object
// with the ability to set a context for a request.
func NewSessionCreateSessionParamsWithContext(ctx context.Context) *SessionCreateSessionParams {
	return &SessionCreateSessionParams{
		Context: ctx,
	}
}

// NewSessionCreateSessionParamsWithHTTPClient creates a new SessionCreateSessionParams object
// with the ability to set a custom HTTPClient for a request.
func NewSessionCreateSessionParamsWithHTTPClient(client *http.Client) *SessionCreateSessionParams {
	return &SessionCreateSessionParams{
		HTTPClient: client,
	}
}

/*
SessionCreateSessionParams contains all the parameters to send to the API endpoint

	for the session create session operation.

	Typically these are written to a http.Request.
*/
type SessionCreateSessionParams struct {

	// Request.
	Request *models.SessionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the session create session params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SessionCreateSessionParams) WithDefaults() *SessionCreateSessionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the session create session params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SessionCreateSessionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the session create session params
func (o *SessionCreateSessionParams) WithTimeout(timeout time.Duration) *SessionCreateSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the session create session params
func (o *SessionCreateSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the session create session params
func (o *SessionCreateSessionParams) WithContext(ctx context.Context) *SessionCreateSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the session create session params
func (o *SessionCreateSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the session create session params
func (o *SessionCreateSessionParams) WithHTTPClient(client *http.Client) *SessionCreateSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the session create session params
func (o *SessionCreateSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the session create session params
func (o *SessionCreateSessionParams) WithRequest(request *models.SessionRequest) *SessionCreateSessionParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the session create session params
func (o *SessionCreateSessionParams) SetRequest(request *models.SessionRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *SessionCreateSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
