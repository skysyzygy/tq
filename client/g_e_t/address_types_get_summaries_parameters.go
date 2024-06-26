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

// NewAddressTypesGetSummariesParams creates a new AddressTypesGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddressTypesGetSummariesParams() *AddressTypesGetSummariesParams {
	return &AddressTypesGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddressTypesGetSummariesParamsWithTimeout creates a new AddressTypesGetSummariesParams object
// with the ability to set a timeout on a request.
func NewAddressTypesGetSummariesParamsWithTimeout(timeout time.Duration) *AddressTypesGetSummariesParams {
	return &AddressTypesGetSummariesParams{
		timeout: timeout,
	}
}

// NewAddressTypesGetSummariesParamsWithContext creates a new AddressTypesGetSummariesParams object
// with the ability to set a context for a request.
func NewAddressTypesGetSummariesParamsWithContext(ctx context.Context) *AddressTypesGetSummariesParams {
	return &AddressTypesGetSummariesParams{
		Context: ctx,
	}
}

// NewAddressTypesGetSummariesParamsWithHTTPClient creates a new AddressTypesGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddressTypesGetSummariesParamsWithHTTPClient(client *http.Client) *AddressTypesGetSummariesParams {
	return &AddressTypesGetSummariesParams{
		HTTPClient: client,
	}
}

/*
AddressTypesGetSummariesParams contains all the parameters to send to the API endpoint

	for the address types get summaries operation.

	Typically these are written to a http.Request.
*/
type AddressTypesGetSummariesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the address types get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressTypesGetSummariesParams) WithDefaults() *AddressTypesGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the address types get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddressTypesGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the address types get summaries params
func (o *AddressTypesGetSummariesParams) WithTimeout(timeout time.Duration) *AddressTypesGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the address types get summaries params
func (o *AddressTypesGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the address types get summaries params
func (o *AddressTypesGetSummariesParams) WithContext(ctx context.Context) *AddressTypesGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the address types get summaries params
func (o *AddressTypesGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the address types get summaries params
func (o *AddressTypesGetSummariesParams) WithHTTPClient(client *http.Client) *AddressTypesGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the address types get summaries params
func (o *AddressTypesGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AddressTypesGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
