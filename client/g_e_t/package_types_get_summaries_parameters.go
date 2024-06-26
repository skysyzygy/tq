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

// NewPackageTypesGetSummariesParams creates a new PackageTypesGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPackageTypesGetSummariesParams() *PackageTypesGetSummariesParams {
	return &PackageTypesGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPackageTypesGetSummariesParamsWithTimeout creates a new PackageTypesGetSummariesParams object
// with the ability to set a timeout on a request.
func NewPackageTypesGetSummariesParamsWithTimeout(timeout time.Duration) *PackageTypesGetSummariesParams {
	return &PackageTypesGetSummariesParams{
		timeout: timeout,
	}
}

// NewPackageTypesGetSummariesParamsWithContext creates a new PackageTypesGetSummariesParams object
// with the ability to set a context for a request.
func NewPackageTypesGetSummariesParamsWithContext(ctx context.Context) *PackageTypesGetSummariesParams {
	return &PackageTypesGetSummariesParams{
		Context: ctx,
	}
}

// NewPackageTypesGetSummariesParamsWithHTTPClient creates a new PackageTypesGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPackageTypesGetSummariesParamsWithHTTPClient(client *http.Client) *PackageTypesGetSummariesParams {
	return &PackageTypesGetSummariesParams{
		HTTPClient: client,
	}
}

/*
PackageTypesGetSummariesParams contains all the parameters to send to the API endpoint

	for the package types get summaries operation.

	Typically these are written to a http.Request.
*/
type PackageTypesGetSummariesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the package types get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackageTypesGetSummariesParams) WithDefaults() *PackageTypesGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the package types get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackageTypesGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the package types get summaries params
func (o *PackageTypesGetSummariesParams) WithTimeout(timeout time.Duration) *PackageTypesGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the package types get summaries params
func (o *PackageTypesGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the package types get summaries params
func (o *PackageTypesGetSummariesParams) WithContext(ctx context.Context) *PackageTypesGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the package types get summaries params
func (o *PackageTypesGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the package types get summaries params
func (o *PackageTypesGetSummariesParams) WithHTTPClient(client *http.Client) *PackageTypesGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the package types get summaries params
func (o *PackageTypesGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PackageTypesGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
