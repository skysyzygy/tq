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

// NewAliasTypesGetSummariesParams creates a new AliasTypesGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAliasTypesGetSummariesParams() *AliasTypesGetSummariesParams {
	return &AliasTypesGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAliasTypesGetSummariesParamsWithTimeout creates a new AliasTypesGetSummariesParams object
// with the ability to set a timeout on a request.
func NewAliasTypesGetSummariesParamsWithTimeout(timeout time.Duration) *AliasTypesGetSummariesParams {
	return &AliasTypesGetSummariesParams{
		timeout: timeout,
	}
}

// NewAliasTypesGetSummariesParamsWithContext creates a new AliasTypesGetSummariesParams object
// with the ability to set a context for a request.
func NewAliasTypesGetSummariesParamsWithContext(ctx context.Context) *AliasTypesGetSummariesParams {
	return &AliasTypesGetSummariesParams{
		Context: ctx,
	}
}

// NewAliasTypesGetSummariesParamsWithHTTPClient creates a new AliasTypesGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewAliasTypesGetSummariesParamsWithHTTPClient(client *http.Client) *AliasTypesGetSummariesParams {
	return &AliasTypesGetSummariesParams{
		HTTPClient: client,
	}
}

/*
AliasTypesGetSummariesParams contains all the parameters to send to the API endpoint

	for the alias types get summaries operation.

	Typically these are written to a http.Request.
*/
type AliasTypesGetSummariesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the alias types get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AliasTypesGetSummariesParams) WithDefaults() *AliasTypesGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the alias types get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AliasTypesGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the alias types get summaries params
func (o *AliasTypesGetSummariesParams) WithTimeout(timeout time.Duration) *AliasTypesGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the alias types get summaries params
func (o *AliasTypesGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the alias types get summaries params
func (o *AliasTypesGetSummariesParams) WithContext(ctx context.Context) *AliasTypesGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the alias types get summaries params
func (o *AliasTypesGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the alias types get summaries params
func (o *AliasTypesGetSummariesParams) WithHTTPClient(client *http.Client) *AliasTypesGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the alias types get summaries params
func (o *AliasTypesGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AliasTypesGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
