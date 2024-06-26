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

// NewOrganizationsGetSummariesParams creates a new OrganizationsGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrganizationsGetSummariesParams() *OrganizationsGetSummariesParams {
	return &OrganizationsGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationsGetSummariesParamsWithTimeout creates a new OrganizationsGetSummariesParams object
// with the ability to set a timeout on a request.
func NewOrganizationsGetSummariesParamsWithTimeout(timeout time.Duration) *OrganizationsGetSummariesParams {
	return &OrganizationsGetSummariesParams{
		timeout: timeout,
	}
}

// NewOrganizationsGetSummariesParamsWithContext creates a new OrganizationsGetSummariesParams object
// with the ability to set a context for a request.
func NewOrganizationsGetSummariesParamsWithContext(ctx context.Context) *OrganizationsGetSummariesParams {
	return &OrganizationsGetSummariesParams{
		Context: ctx,
	}
}

// NewOrganizationsGetSummariesParamsWithHTTPClient creates a new OrganizationsGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrganizationsGetSummariesParamsWithHTTPClient(client *http.Client) *OrganizationsGetSummariesParams {
	return &OrganizationsGetSummariesParams{
		HTTPClient: client,
	}
}

/*
OrganizationsGetSummariesParams contains all the parameters to send to the API endpoint

	for the organizations get summaries operation.

	Typically these are written to a http.Request.
*/
type OrganizationsGetSummariesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the organizations get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationsGetSummariesParams) WithDefaults() *OrganizationsGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the organizations get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationsGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the organizations get summaries params
func (o *OrganizationsGetSummariesParams) WithTimeout(timeout time.Duration) *OrganizationsGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organizations get summaries params
func (o *OrganizationsGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organizations get summaries params
func (o *OrganizationsGetSummariesParams) WithContext(ctx context.Context) *OrganizationsGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organizations get summaries params
func (o *OrganizationsGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organizations get summaries params
func (o *OrganizationsGetSummariesParams) WithHTTPClient(client *http.Client) *OrganizationsGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organizations get summaries params
func (o *OrganizationsGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationsGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
