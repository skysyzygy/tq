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

// NewAffiliationTypesGetSummariesParams creates a new AffiliationTypesGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAffiliationTypesGetSummariesParams() *AffiliationTypesGetSummariesParams {
	return &AffiliationTypesGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAffiliationTypesGetSummariesParamsWithTimeout creates a new AffiliationTypesGetSummariesParams object
// with the ability to set a timeout on a request.
func NewAffiliationTypesGetSummariesParamsWithTimeout(timeout time.Duration) *AffiliationTypesGetSummariesParams {
	return &AffiliationTypesGetSummariesParams{
		timeout: timeout,
	}
}

// NewAffiliationTypesGetSummariesParamsWithContext creates a new AffiliationTypesGetSummariesParams object
// with the ability to set a context for a request.
func NewAffiliationTypesGetSummariesParamsWithContext(ctx context.Context) *AffiliationTypesGetSummariesParams {
	return &AffiliationTypesGetSummariesParams{
		Context: ctx,
	}
}

// NewAffiliationTypesGetSummariesParamsWithHTTPClient creates a new AffiliationTypesGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewAffiliationTypesGetSummariesParamsWithHTTPClient(client *http.Client) *AffiliationTypesGetSummariesParams {
	return &AffiliationTypesGetSummariesParams{
		HTTPClient: client,
	}
}

/*
AffiliationTypesGetSummariesParams contains all the parameters to send to the API endpoint

	for the affiliation types get summaries operation.

	Typically these are written to a http.Request.
*/
type AffiliationTypesGetSummariesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the affiliation types get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AffiliationTypesGetSummariesParams) WithDefaults() *AffiliationTypesGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the affiliation types get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AffiliationTypesGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the affiliation types get summaries params
func (o *AffiliationTypesGetSummariesParams) WithTimeout(timeout time.Duration) *AffiliationTypesGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the affiliation types get summaries params
func (o *AffiliationTypesGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the affiliation types get summaries params
func (o *AffiliationTypesGetSummariesParams) WithContext(ctx context.Context) *AffiliationTypesGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the affiliation types get summaries params
func (o *AffiliationTypesGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the affiliation types get summaries params
func (o *AffiliationTypesGetSummariesParams) WithHTTPClient(client *http.Client) *AffiliationTypesGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the affiliation types get summaries params
func (o *AffiliationTypesGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AffiliationTypesGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
