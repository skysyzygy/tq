// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

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

// NewSuffixesDeleteParams creates a new SuffixesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSuffixesDeleteParams() *SuffixesDeleteParams {
	return &SuffixesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSuffixesDeleteParamsWithTimeout creates a new SuffixesDeleteParams object
// with the ability to set a timeout on a request.
func NewSuffixesDeleteParamsWithTimeout(timeout time.Duration) *SuffixesDeleteParams {
	return &SuffixesDeleteParams{
		timeout: timeout,
	}
}

// NewSuffixesDeleteParamsWithContext creates a new SuffixesDeleteParams object
// with the ability to set a context for a request.
func NewSuffixesDeleteParamsWithContext(ctx context.Context) *SuffixesDeleteParams {
	return &SuffixesDeleteParams{
		Context: ctx,
	}
}

// NewSuffixesDeleteParamsWithHTTPClient creates a new SuffixesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewSuffixesDeleteParamsWithHTTPClient(client *http.Client) *SuffixesDeleteParams {
	return &SuffixesDeleteParams{
		HTTPClient: client,
	}
}

/*
SuffixesDeleteParams contains all the parameters to send to the API endpoint

	for the suffixes delete operation.

	Typically these are written to a http.Request.
*/
type SuffixesDeleteParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the suffixes delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SuffixesDeleteParams) WithDefaults() *SuffixesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the suffixes delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SuffixesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the suffixes delete params
func (o *SuffixesDeleteParams) WithTimeout(timeout time.Duration) *SuffixesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the suffixes delete params
func (o *SuffixesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the suffixes delete params
func (o *SuffixesDeleteParams) WithContext(ctx context.Context) *SuffixesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the suffixes delete params
func (o *SuffixesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the suffixes delete params
func (o *SuffixesDeleteParams) WithHTTPClient(client *http.Client) *SuffixesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the suffixes delete params
func (o *SuffixesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the suffixes delete params
func (o *SuffixesDeleteParams) WithID(id string) *SuffixesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the suffixes delete params
func (o *SuffixesDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SuffixesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
