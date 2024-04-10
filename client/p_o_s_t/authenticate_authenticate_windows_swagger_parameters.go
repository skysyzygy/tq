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

// NewAuthenticateAuthenticateWindowsParams creates a new AuthenticateAuthenticateWindowsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthenticateAuthenticateWindowsParams() *AuthenticateAuthenticateWindowsParams {
	return &AuthenticateAuthenticateWindowsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthenticateAuthenticateWindowsParamsWithTimeout creates a new AuthenticateAuthenticateWindowsParams object
// with the ability to set a timeout on a request.
func NewAuthenticateAuthenticateWindowsParamsWithTimeout(timeout time.Duration) *AuthenticateAuthenticateWindowsParams {
	return &AuthenticateAuthenticateWindowsParams{
		timeout: timeout,
	}
}

// NewAuthenticateAuthenticateWindowsParamsWithContext creates a new AuthenticateAuthenticateWindowsParams object
// with the ability to set a context for a request.
func NewAuthenticateAuthenticateWindowsParamsWithContext(ctx context.Context) *AuthenticateAuthenticateWindowsParams {
	return &AuthenticateAuthenticateWindowsParams{
		Context: ctx,
	}
}

// NewAuthenticateAuthenticateWindowsParamsWithHTTPClient creates a new AuthenticateAuthenticateWindowsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuthenticateAuthenticateWindowsParamsWithHTTPClient(client *http.Client) *AuthenticateAuthenticateWindowsParams {
	return &AuthenticateAuthenticateWindowsParams{
		HTTPClient: client,
	}
}

/*
AuthenticateAuthenticateWindowsParams contains all the parameters to send to the API endpoint

	for the authenticate authenticate windows operation.

	Typically these are written to a http.Request.
*/
type AuthenticateAuthenticateWindowsParams struct {

	// AuthenticationRequest.
	AuthenticationRequest *models.WindowsAuthenticationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the authenticate authenticate windows params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthenticateAuthenticateWindowsParams) WithDefaults() *AuthenticateAuthenticateWindowsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the authenticate authenticate windows params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthenticateAuthenticateWindowsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the authenticate authenticate windows params
func (o *AuthenticateAuthenticateWindowsParams) WithTimeout(timeout time.Duration) *AuthenticateAuthenticateWindowsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the authenticate authenticate windows params
func (o *AuthenticateAuthenticateWindowsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the authenticate authenticate windows params
func (o *AuthenticateAuthenticateWindowsParams) WithContext(ctx context.Context) *AuthenticateAuthenticateWindowsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the authenticate authenticate windows params
func (o *AuthenticateAuthenticateWindowsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the authenticate authenticate windows params
func (o *AuthenticateAuthenticateWindowsParams) WithHTTPClient(client *http.Client) *AuthenticateAuthenticateWindowsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the authenticate authenticate windows params
func (o *AuthenticateAuthenticateWindowsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthenticationRequest adds the authenticationRequest to the authenticate authenticate windows params
func (o *AuthenticateAuthenticateWindowsParams) WithAuthenticationRequest(authenticationRequest *models.WindowsAuthenticationRequest) *AuthenticateAuthenticateWindowsParams {
	o.SetAuthenticationRequest(authenticationRequest)
	return o
}

// SetAuthenticationRequest adds the authenticationRequest to the authenticate authenticate windows params
func (o *AuthenticateAuthenticateWindowsParams) SetAuthenticationRequest(authenticationRequest *models.WindowsAuthenticationRequest) {
	o.AuthenticationRequest = authenticationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *AuthenticateAuthenticateWindowsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
