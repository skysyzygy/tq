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

// NewSalutationsGetParams creates a new SalutationsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSalutationsGetParams() *SalutationsGetParams {
	return &SalutationsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSalutationsGetParamsWithTimeout creates a new SalutationsGetParams object
// with the ability to set a timeout on a request.
func NewSalutationsGetParamsWithTimeout(timeout time.Duration) *SalutationsGetParams {
	return &SalutationsGetParams{
		timeout: timeout,
	}
}

// NewSalutationsGetParamsWithContext creates a new SalutationsGetParams object
// with the ability to set a context for a request.
func NewSalutationsGetParamsWithContext(ctx context.Context) *SalutationsGetParams {
	return &SalutationsGetParams{
		Context: ctx,
	}
}

// NewSalutationsGetParamsWithHTTPClient creates a new SalutationsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSalutationsGetParamsWithHTTPClient(client *http.Client) *SalutationsGetParams {
	return &SalutationsGetParams{
		HTTPClient: client,
	}
}

/*
SalutationsGetParams contains all the parameters to send to the API endpoint

	for the salutations get operation.

	Typically these are written to a http.Request.
*/
type SalutationsGetParams struct {

	// SalutationID.
	SalutationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the salutations get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SalutationsGetParams) WithDefaults() *SalutationsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the salutations get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SalutationsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the salutations get params
func (o *SalutationsGetParams) WithTimeout(timeout time.Duration) *SalutationsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the salutations get params
func (o *SalutationsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the salutations get params
func (o *SalutationsGetParams) WithContext(ctx context.Context) *SalutationsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the salutations get params
func (o *SalutationsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the salutations get params
func (o *SalutationsGetParams) WithHTTPClient(client *http.Client) *SalutationsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the salutations get params
func (o *SalutationsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSalutationID adds the salutationID to the salutations get params
func (o *SalutationsGetParams) WithSalutationID(salutationID string) *SalutationsGetParams {
	o.SetSalutationID(salutationID)
	return o
}

// SetSalutationID adds the salutationId to the salutations get params
func (o *SalutationsGetParams) SetSalutationID(salutationID string) {
	o.SalutationID = salutationID
}

// WriteToRequest writes these params to a swagger request
func (o *SalutationsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param salutationId
	if err := r.SetPathParam("salutationId", o.SalutationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
