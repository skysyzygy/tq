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

// NewAuthenticateAuthenticateParams creates a new AuthenticateAuthenticateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthenticateAuthenticateParams() *AuthenticateAuthenticateParams {
	return &AuthenticateAuthenticateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthenticateAuthenticateParamsWithTimeout creates a new AuthenticateAuthenticateParams object
// with the ability to set a timeout on a request.
func NewAuthenticateAuthenticateParamsWithTimeout(timeout time.Duration) *AuthenticateAuthenticateParams {
	return &AuthenticateAuthenticateParams{
		timeout: timeout,
	}
}

// NewAuthenticateAuthenticateParamsWithContext creates a new AuthenticateAuthenticateParams object
// with the ability to set a context for a request.
func NewAuthenticateAuthenticateParamsWithContext(ctx context.Context) *AuthenticateAuthenticateParams {
	return &AuthenticateAuthenticateParams{
		Context: ctx,
	}
}

// NewAuthenticateAuthenticateParamsWithHTTPClient creates a new AuthenticateAuthenticateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuthenticateAuthenticateParamsWithHTTPClient(client *http.Client) *AuthenticateAuthenticateParams {
	return &AuthenticateAuthenticateParams{
		HTTPClient: client,
	}
}

/*
AuthenticateAuthenticateParams contains all the parameters to send to the API endpoint

	for the authenticate authenticate operation.

	Typically these are written to a http.Request.
*/
type AuthenticateAuthenticateParams struct {

	// AuthenticationRequest.
	AuthenticationRequest *models.AuthenticationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the authenticate authenticate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthenticateAuthenticateParams) WithDefaults() *AuthenticateAuthenticateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the authenticate authenticate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthenticateAuthenticateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the authenticate authenticate params
func (o *AuthenticateAuthenticateParams) WithTimeout(timeout time.Duration) *AuthenticateAuthenticateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the authenticate authenticate params
func (o *AuthenticateAuthenticateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the authenticate authenticate params
func (o *AuthenticateAuthenticateParams) WithContext(ctx context.Context) *AuthenticateAuthenticateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the authenticate authenticate params
func (o *AuthenticateAuthenticateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the authenticate authenticate params
func (o *AuthenticateAuthenticateParams) WithHTTPClient(client *http.Client) *AuthenticateAuthenticateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the authenticate authenticate params
func (o *AuthenticateAuthenticateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthenticationRequest adds the authenticationRequest to the authenticate authenticate params
func (o *AuthenticateAuthenticateParams) WithAuthenticationRequest(authenticationRequest *models.AuthenticationRequest) *AuthenticateAuthenticateParams {
	o.SetAuthenticationRequest(authenticationRequest)
	return o
}

// SetAuthenticationRequest adds the authenticationRequest to the authenticate authenticate params
func (o *AuthenticateAuthenticateParams) SetAuthenticationRequest(authenticationRequest *models.AuthenticationRequest) {
	o.AuthenticationRequest = authenticationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *AuthenticateAuthenticateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.AuthenticationRequest != nil {
		if err := r.SetBodyParam(o.AuthenticationRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
