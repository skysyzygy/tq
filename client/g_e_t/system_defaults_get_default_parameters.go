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

// NewSystemDefaultsGetDefaultParams creates a new SystemDefaultsGetDefaultParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSystemDefaultsGetDefaultParams() *SystemDefaultsGetDefaultParams {
	return &SystemDefaultsGetDefaultParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSystemDefaultsGetDefaultParamsWithTimeout creates a new SystemDefaultsGetDefaultParams object
// with the ability to set a timeout on a request.
func NewSystemDefaultsGetDefaultParamsWithTimeout(timeout time.Duration) *SystemDefaultsGetDefaultParams {
	return &SystemDefaultsGetDefaultParams{
		timeout: timeout,
	}
}

// NewSystemDefaultsGetDefaultParamsWithContext creates a new SystemDefaultsGetDefaultParams object
// with the ability to set a context for a request.
func NewSystemDefaultsGetDefaultParamsWithContext(ctx context.Context) *SystemDefaultsGetDefaultParams {
	return &SystemDefaultsGetDefaultParams{
		Context: ctx,
	}
}

// NewSystemDefaultsGetDefaultParamsWithHTTPClient creates a new SystemDefaultsGetDefaultParams object
// with the ability to set a custom HTTPClient for a request.
func NewSystemDefaultsGetDefaultParamsWithHTTPClient(client *http.Client) *SystemDefaultsGetDefaultParams {
	return &SystemDefaultsGetDefaultParams{
		HTTPClient: client,
	}
}

/*
SystemDefaultsGetDefaultParams contains all the parameters to send to the API endpoint

	for the system defaults get default operation.

	Typically these are written to a http.Request.
*/
type SystemDefaultsGetDefaultParams struct {

	// Keys.
	Keys *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the system defaults get default params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemDefaultsGetDefaultParams) WithDefaults() *SystemDefaultsGetDefaultParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the system defaults get default params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemDefaultsGetDefaultParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the system defaults get default params
func (o *SystemDefaultsGetDefaultParams) WithTimeout(timeout time.Duration) *SystemDefaultsGetDefaultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system defaults get default params
func (o *SystemDefaultsGetDefaultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system defaults get default params
func (o *SystemDefaultsGetDefaultParams) WithContext(ctx context.Context) *SystemDefaultsGetDefaultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system defaults get default params
func (o *SystemDefaultsGetDefaultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system defaults get default params
func (o *SystemDefaultsGetDefaultParams) WithHTTPClient(client *http.Client) *SystemDefaultsGetDefaultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system defaults get default params
func (o *SystemDefaultsGetDefaultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKeys adds the keys to the system defaults get default params
func (o *SystemDefaultsGetDefaultParams) WithKeys(keys *string) *SystemDefaultsGetDefaultParams {
	o.SetKeys(keys)
	return o
}

// SetKeys adds the keys to the system defaults get default params
func (o *SystemDefaultsGetDefaultParams) SetKeys(keys *string) {
	o.Keys = keys
}

// WriteToRequest writes these params to a swagger request
func (o *SystemDefaultsGetDefaultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Keys != nil {

		// query param keys
		var qrKeys string

		if o.Keys != nil {
			qrKeys = *o.Keys
		}
		qKeys := qrKeys
		if qKeys != "" {

			if err := r.SetQueryParam("keys", qKeys); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
