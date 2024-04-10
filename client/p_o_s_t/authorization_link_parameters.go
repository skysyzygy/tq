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

// NewAuthorizationLinkParams creates a new AuthorizationLinkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthorizationLinkParams() *AuthorizationLinkParams {
	return &AuthorizationLinkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthorizationLinkParamsWithTimeout creates a new AuthorizationLinkParams object
// with the ability to set a timeout on a request.
func NewAuthorizationLinkParamsWithTimeout(timeout time.Duration) *AuthorizationLinkParams {
	return &AuthorizationLinkParams{
		timeout: timeout,
	}
}

// NewAuthorizationLinkParamsWithContext creates a new AuthorizationLinkParams object
// with the ability to set a context for a request.
func NewAuthorizationLinkParamsWithContext(ctx context.Context) *AuthorizationLinkParams {
	return &AuthorizationLinkParams{
		Context: ctx,
	}
}

// NewAuthorizationLinkParamsWithHTTPClient creates a new AuthorizationLinkParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuthorizationLinkParamsWithHTTPClient(client *http.Client) *AuthorizationLinkParams {
	return &AuthorizationLinkParams{
		HTTPClient: client,
	}
}

/*
AuthorizationLinkParams contains all the parameters to send to the API endpoint

	for the authorization link operation.

	Typically these are written to a http.Request.
*/
type AuthorizationLinkParams struct {

	// PayByLinkRequest.
	PayByLinkRequest *models.PayByLinkRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the authorization link params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthorizationLinkParams) WithDefaults() *AuthorizationLinkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the authorization link params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthorizationLinkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the authorization link params
func (o *AuthorizationLinkParams) WithTimeout(timeout time.Duration) *AuthorizationLinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the authorization link params
func (o *AuthorizationLinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the authorization link params
func (o *AuthorizationLinkParams) WithContext(ctx context.Context) *AuthorizationLinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the authorization link params
func (o *AuthorizationLinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the authorization link params
func (o *AuthorizationLinkParams) WithHTTPClient(client *http.Client) *AuthorizationLinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the authorization link params
func (o *AuthorizationLinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPayByLinkRequest adds the payByLinkRequest to the authorization link params
func (o *AuthorizationLinkParams) WithPayByLinkRequest(payByLinkRequest *models.PayByLinkRequest) *AuthorizationLinkParams {
	o.SetPayByLinkRequest(payByLinkRequest)
	return o
}

// SetPayByLinkRequest adds the payByLinkRequest to the authorization link params
func (o *AuthorizationLinkParams) SetPayByLinkRequest(payByLinkRequest *models.PayByLinkRequest) {
	o.PayByLinkRequest = payByLinkRequest
}

// WriteToRequest writes these params to a swagger request
func (o *AuthorizationLinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PayByLinkRequest != nil {
		if err := r.SetBodyParam(o.PayByLinkRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
