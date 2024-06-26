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

// NewSessionAddVariableParams creates a new SessionAddVariableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSessionAddVariableParams() *SessionAddVariableParams {
	return &SessionAddVariableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSessionAddVariableParamsWithTimeout creates a new SessionAddVariableParams object
// with the ability to set a timeout on a request.
func NewSessionAddVariableParamsWithTimeout(timeout time.Duration) *SessionAddVariableParams {
	return &SessionAddVariableParams{
		timeout: timeout,
	}
}

// NewSessionAddVariableParamsWithContext creates a new SessionAddVariableParams object
// with the ability to set a context for a request.
func NewSessionAddVariableParamsWithContext(ctx context.Context) *SessionAddVariableParams {
	return &SessionAddVariableParams{
		Context: ctx,
	}
}

// NewSessionAddVariableParamsWithHTTPClient creates a new SessionAddVariableParams object
// with the ability to set a custom HTTPClient for a request.
func NewSessionAddVariableParamsWithHTTPClient(client *http.Client) *SessionAddVariableParams {
	return &SessionAddVariableParams{
		HTTPClient: client,
	}
}

/*
SessionAddVariableParams contains all the parameters to send to the API endpoint

	for the session add variable operation.

	Typically these are written to a http.Request.
*/
type SessionAddVariableParams struct {

	// SessionKey.
	SessionKey string

	// SessionVariable.
	SessionVariable *models.SessionVariable

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the session add variable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SessionAddVariableParams) WithDefaults() *SessionAddVariableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the session add variable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SessionAddVariableParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the session add variable params
func (o *SessionAddVariableParams) WithTimeout(timeout time.Duration) *SessionAddVariableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the session add variable params
func (o *SessionAddVariableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the session add variable params
func (o *SessionAddVariableParams) WithContext(ctx context.Context) *SessionAddVariableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the session add variable params
func (o *SessionAddVariableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the session add variable params
func (o *SessionAddVariableParams) WithHTTPClient(client *http.Client) *SessionAddVariableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the session add variable params
func (o *SessionAddVariableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSessionKey adds the sessionKey to the session add variable params
func (o *SessionAddVariableParams) WithSessionKey(sessionKey string) *SessionAddVariableParams {
	o.SetSessionKey(sessionKey)
	return o
}

// SetSessionKey adds the sessionKey to the session add variable params
func (o *SessionAddVariableParams) SetSessionKey(sessionKey string) {
	o.SessionKey = sessionKey
}

// WithSessionVariable adds the sessionVariable to the session add variable params
func (o *SessionAddVariableParams) WithSessionVariable(sessionVariable *models.SessionVariable) *SessionAddVariableParams {
	o.SetSessionVariable(sessionVariable)
	return o
}

// SetSessionVariable adds the sessionVariable to the session add variable params
func (o *SessionAddVariableParams) SetSessionVariable(sessionVariable *models.SessionVariable) {
	o.SessionVariable = sessionVariable
}

// WriteToRequest writes these params to a swagger request
func (o *SessionAddVariableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param sessionKey
	if err := r.SetPathParam("sessionKey", o.SessionKey); err != nil {
		return err
	}
	if o.SessionVariable != nil {
		if err := r.SetBodyParam(o.SessionVariable); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
