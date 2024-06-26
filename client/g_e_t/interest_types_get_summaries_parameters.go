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

// NewInterestTypesGetSummariesParams creates a new InterestTypesGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInterestTypesGetSummariesParams() *InterestTypesGetSummariesParams {
	return &InterestTypesGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInterestTypesGetSummariesParamsWithTimeout creates a new InterestTypesGetSummariesParams object
// with the ability to set a timeout on a request.
func NewInterestTypesGetSummariesParamsWithTimeout(timeout time.Duration) *InterestTypesGetSummariesParams {
	return &InterestTypesGetSummariesParams{
		timeout: timeout,
	}
}

// NewInterestTypesGetSummariesParamsWithContext creates a new InterestTypesGetSummariesParams object
// with the ability to set a context for a request.
func NewInterestTypesGetSummariesParamsWithContext(ctx context.Context) *InterestTypesGetSummariesParams {
	return &InterestTypesGetSummariesParams{
		Context: ctx,
	}
}

// NewInterestTypesGetSummariesParamsWithHTTPClient creates a new InterestTypesGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewInterestTypesGetSummariesParamsWithHTTPClient(client *http.Client) *InterestTypesGetSummariesParams {
	return &InterestTypesGetSummariesParams{
		HTTPClient: client,
	}
}

/*
InterestTypesGetSummariesParams contains all the parameters to send to the API endpoint

	for the interest types get summaries operation.

	Typically these are written to a http.Request.
*/
type InterestTypesGetSummariesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the interest types get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InterestTypesGetSummariesParams) WithDefaults() *InterestTypesGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the interest types get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InterestTypesGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the interest types get summaries params
func (o *InterestTypesGetSummariesParams) WithTimeout(timeout time.Duration) *InterestTypesGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the interest types get summaries params
func (o *InterestTypesGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the interest types get summaries params
func (o *InterestTypesGetSummariesParams) WithContext(ctx context.Context) *InterestTypesGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the interest types get summaries params
func (o *InterestTypesGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the interest types get summaries params
func (o *InterestTypesGetSummariesParams) WithHTTPClient(client *http.Client) *InterestTypesGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the interest types get summaries params
func (o *InterestTypesGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *InterestTypesGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
