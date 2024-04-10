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

// NewComposersGetSummariesParams creates a new ComposersGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewComposersGetSummariesParams() *ComposersGetSummariesParams {
	return &ComposersGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewComposersGetSummariesParamsWithTimeout creates a new ComposersGetSummariesParams object
// with the ability to set a timeout on a request.
func NewComposersGetSummariesParamsWithTimeout(timeout time.Duration) *ComposersGetSummariesParams {
	return &ComposersGetSummariesParams{
		timeout: timeout,
	}
}

// NewComposersGetSummariesParamsWithContext creates a new ComposersGetSummariesParams object
// with the ability to set a context for a request.
func NewComposersGetSummariesParamsWithContext(ctx context.Context) *ComposersGetSummariesParams {
	return &ComposersGetSummariesParams{
		Context: ctx,
	}
}

// NewComposersGetSummariesParamsWithHTTPClient creates a new ComposersGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewComposersGetSummariesParamsWithHTTPClient(client *http.Client) *ComposersGetSummariesParams {
	return &ComposersGetSummariesParams{
		HTTPClient: client,
	}
}

/*
ComposersGetSummariesParams contains all the parameters to send to the API endpoint

	for the composers get summaries operation.

	Typically these are written to a http.Request.
*/
type ComposersGetSummariesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the composers get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ComposersGetSummariesParams) WithDefaults() *ComposersGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the composers get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ComposersGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the composers get summaries params
func (o *ComposersGetSummariesParams) WithTimeout(timeout time.Duration) *ComposersGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the composers get summaries params
func (o *ComposersGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the composers get summaries params
func (o *ComposersGetSummariesParams) WithContext(ctx context.Context) *ComposersGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the composers get summaries params
func (o *ComposersGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the composers get summaries params
func (o *ComposersGetSummariesParams) WithHTTPClient(client *http.Client) *ComposersGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the composers get summaries params
func (o *ComposersGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ComposersGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
