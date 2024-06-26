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

// NewResourceTypesDeleteParams creates a new ResourceTypesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResourceTypesDeleteParams() *ResourceTypesDeleteParams {
	return &ResourceTypesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResourceTypesDeleteParamsWithTimeout creates a new ResourceTypesDeleteParams object
// with the ability to set a timeout on a request.
func NewResourceTypesDeleteParamsWithTimeout(timeout time.Duration) *ResourceTypesDeleteParams {
	return &ResourceTypesDeleteParams{
		timeout: timeout,
	}
}

// NewResourceTypesDeleteParamsWithContext creates a new ResourceTypesDeleteParams object
// with the ability to set a context for a request.
func NewResourceTypesDeleteParamsWithContext(ctx context.Context) *ResourceTypesDeleteParams {
	return &ResourceTypesDeleteParams{
		Context: ctx,
	}
}

// NewResourceTypesDeleteParamsWithHTTPClient creates a new ResourceTypesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewResourceTypesDeleteParamsWithHTTPClient(client *http.Client) *ResourceTypesDeleteParams {
	return &ResourceTypesDeleteParams{
		HTTPClient: client,
	}
}

/*
ResourceTypesDeleteParams contains all the parameters to send to the API endpoint

	for the resource types delete operation.

	Typically these are written to a http.Request.
*/
type ResourceTypesDeleteParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the resource types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResourceTypesDeleteParams) WithDefaults() *ResourceTypesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the resource types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResourceTypesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the resource types delete params
func (o *ResourceTypesDeleteParams) WithTimeout(timeout time.Duration) *ResourceTypesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resource types delete params
func (o *ResourceTypesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resource types delete params
func (o *ResourceTypesDeleteParams) WithContext(ctx context.Context) *ResourceTypesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resource types delete params
func (o *ResourceTypesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resource types delete params
func (o *ResourceTypesDeleteParams) WithHTTPClient(client *http.Client) *ResourceTypesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resource types delete params
func (o *ResourceTypesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the resource types delete params
func (o *ResourceTypesDeleteParams) WithID(id string) *ResourceTypesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the resource types delete params
func (o *ResourceTypesDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ResourceTypesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
