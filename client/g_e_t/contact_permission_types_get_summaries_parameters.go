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

// NewContactPermissionTypesGetSummariesParams creates a new ContactPermissionTypesGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContactPermissionTypesGetSummariesParams() *ContactPermissionTypesGetSummariesParams {
	return &ContactPermissionTypesGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContactPermissionTypesGetSummariesParamsWithTimeout creates a new ContactPermissionTypesGetSummariesParams object
// with the ability to set a timeout on a request.
func NewContactPermissionTypesGetSummariesParamsWithTimeout(timeout time.Duration) *ContactPermissionTypesGetSummariesParams {
	return &ContactPermissionTypesGetSummariesParams{
		timeout: timeout,
	}
}

// NewContactPermissionTypesGetSummariesParamsWithContext creates a new ContactPermissionTypesGetSummariesParams object
// with the ability to set a context for a request.
func NewContactPermissionTypesGetSummariesParamsWithContext(ctx context.Context) *ContactPermissionTypesGetSummariesParams {
	return &ContactPermissionTypesGetSummariesParams{
		Context: ctx,
	}
}

// NewContactPermissionTypesGetSummariesParamsWithHTTPClient creates a new ContactPermissionTypesGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewContactPermissionTypesGetSummariesParamsWithHTTPClient(client *http.Client) *ContactPermissionTypesGetSummariesParams {
	return &ContactPermissionTypesGetSummariesParams{
		HTTPClient: client,
	}
}

/*
ContactPermissionTypesGetSummariesParams contains all the parameters to send to the API endpoint

	for the contact permission types get summaries operation.

	Typically these are written to a http.Request.
*/
type ContactPermissionTypesGetSummariesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the contact permission types get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactPermissionTypesGetSummariesParams) WithDefaults() *ContactPermissionTypesGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the contact permission types get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactPermissionTypesGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the contact permission types get summaries params
func (o *ContactPermissionTypesGetSummariesParams) WithTimeout(timeout time.Duration) *ContactPermissionTypesGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contact permission types get summaries params
func (o *ContactPermissionTypesGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contact permission types get summaries params
func (o *ContactPermissionTypesGetSummariesParams) WithContext(ctx context.Context) *ContactPermissionTypesGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contact permission types get summaries params
func (o *ContactPermissionTypesGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contact permission types get summaries params
func (o *ContactPermissionTypesGetSummariesParams) WithHTTPClient(client *http.Client) *ContactPermissionTypesGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contact permission types get summaries params
func (o *ContactPermissionTypesGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ContactPermissionTypesGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
