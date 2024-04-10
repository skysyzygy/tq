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

// NewAffiliationsGetParams creates a new AffiliationsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAffiliationsGetParams() *AffiliationsGetParams {
	return &AffiliationsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAffiliationsGetParamsWithTimeout creates a new AffiliationsGetParams object
// with the ability to set a timeout on a request.
func NewAffiliationsGetParamsWithTimeout(timeout time.Duration) *AffiliationsGetParams {
	return &AffiliationsGetParams{
		timeout: timeout,
	}
}

// NewAffiliationsGetParamsWithContext creates a new AffiliationsGetParams object
// with the ability to set a context for a request.
func NewAffiliationsGetParamsWithContext(ctx context.Context) *AffiliationsGetParams {
	return &AffiliationsGetParams{
		Context: ctx,
	}
}

// NewAffiliationsGetParamsWithHTTPClient creates a new AffiliationsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewAffiliationsGetParamsWithHTTPClient(client *http.Client) *AffiliationsGetParams {
	return &AffiliationsGetParams{
		HTTPClient: client,
	}
}

/*
AffiliationsGetParams contains all the parameters to send to the API endpoint

	for the affiliations get operation.

	Typically these are written to a http.Request.
*/
type AffiliationsGetParams struct {

	// AffiliationID.
	AffiliationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the affiliations get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AffiliationsGetParams) WithDefaults() *AffiliationsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the affiliations get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AffiliationsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the affiliations get params
func (o *AffiliationsGetParams) WithTimeout(timeout time.Duration) *AffiliationsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the affiliations get params
func (o *AffiliationsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the affiliations get params
func (o *AffiliationsGetParams) WithContext(ctx context.Context) *AffiliationsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the affiliations get params
func (o *AffiliationsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the affiliations get params
func (o *AffiliationsGetParams) WithHTTPClient(client *http.Client) *AffiliationsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the affiliations get params
func (o *AffiliationsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAffiliationID adds the affiliationID to the affiliations get params
func (o *AffiliationsGetParams) WithAffiliationID(affiliationID string) *AffiliationsGetParams {
	o.SetAffiliationID(affiliationID)
	return o
}

// SetAffiliationID adds the affiliationId to the affiliations get params
func (o *AffiliationsGetParams) SetAffiliationID(affiliationID string) {
	o.AffiliationID = affiliationID
}

// WriteToRequest writes these params to a swagger request
func (o *AffiliationsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param affiliationId
	if err := r.SetPathParam("affiliationId", o.AffiliationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
