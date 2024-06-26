// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

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
)

// NewSecurityHoldCodesGetAllParams creates a new SecurityHoldCodesGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityHoldCodesGetAllParams() *SecurityHoldCodesGetAllParams {
	return &SecurityHoldCodesGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityHoldCodesGetAllParamsWithTimeout creates a new SecurityHoldCodesGetAllParams object
// with the ability to set a timeout on a request.
func NewSecurityHoldCodesGetAllParamsWithTimeout(timeout time.Duration) *SecurityHoldCodesGetAllParams {
	return &SecurityHoldCodesGetAllParams{
		timeout: timeout,
	}
}

// NewSecurityHoldCodesGetAllParamsWithContext creates a new SecurityHoldCodesGetAllParams object
// with the ability to set a context for a request.
func NewSecurityHoldCodesGetAllParamsWithContext(ctx context.Context) *SecurityHoldCodesGetAllParams {
	return &SecurityHoldCodesGetAllParams{
		Context: ctx,
	}
}

// NewSecurityHoldCodesGetAllParamsWithHTTPClient creates a new SecurityHoldCodesGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityHoldCodesGetAllParamsWithHTTPClient(client *http.Client) *SecurityHoldCodesGetAllParams {
	return &SecurityHoldCodesGetAllParams{
		HTTPClient: client,
	}
}

/*
SecurityHoldCodesGetAllParams contains all the parameters to send to the API endpoint

	for the security hold codes get all operation.

	Typically these are written to a http.Request.
*/
type SecurityHoldCodesGetAllParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security hold codes get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityHoldCodesGetAllParams) WithDefaults() *SecurityHoldCodesGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security hold codes get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityHoldCodesGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the security hold codes get all params
func (o *SecurityHoldCodesGetAllParams) WithTimeout(timeout time.Duration) *SecurityHoldCodesGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security hold codes get all params
func (o *SecurityHoldCodesGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security hold codes get all params
func (o *SecurityHoldCodesGetAllParams) WithContext(ctx context.Context) *SecurityHoldCodesGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security hold codes get all params
func (o *SecurityHoldCodesGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security hold codes get all params
func (o *SecurityHoldCodesGetAllParams) WithHTTPClient(client *http.Client) *SecurityHoldCodesGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security hold codes get all params
func (o *SecurityHoldCodesGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityHoldCodesGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
