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

// NewCustomGetParams creates a new CustomGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomGetParams() *CustomGetParams {
	return &CustomGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomGetParamsWithTimeout creates a new CustomGetParams object
// with the ability to set a timeout on a request.
func NewCustomGetParamsWithTimeout(timeout time.Duration) *CustomGetParams {
	return &CustomGetParams{
		timeout: timeout,
	}
}

// NewCustomGetParamsWithContext creates a new CustomGetParams object
// with the ability to set a context for a request.
func NewCustomGetParamsWithContext(ctx context.Context) *CustomGetParams {
	return &CustomGetParams{
		Context: ctx,
	}
}

// NewCustomGetParamsWithHTTPClient creates a new CustomGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomGetParamsWithHTTPClient(client *http.Client) *CustomGetParams {
	return &CustomGetParams{
		HTTPClient: client,
	}
}

/*
CustomGetParams contains all the parameters to send to the API endpoint

	for the custom get operation.

	Typically these are written to a http.Request.
*/
type CustomGetParams struct {

	// ID.
	ID string

	// ResourceName.
	ResourceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the custom get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomGetParams) WithDefaults() *CustomGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the custom get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the custom get params
func (o *CustomGetParams) WithTimeout(timeout time.Duration) *CustomGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom get params
func (o *CustomGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom get params
func (o *CustomGetParams) WithContext(ctx context.Context) *CustomGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom get params
func (o *CustomGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom get params
func (o *CustomGetParams) WithHTTPClient(client *http.Client) *CustomGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom get params
func (o *CustomGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the custom get params
func (o *CustomGetParams) WithID(id string) *CustomGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the custom get params
func (o *CustomGetParams) SetID(id string) {
	o.ID = id
}

// WithResourceName adds the resourceName to the custom get params
func (o *CustomGetParams) WithResourceName(resourceName string) *CustomGetParams {
	o.SetResourceName(resourceName)
	return o
}

// SetResourceName adds the resourceName to the custom get params
func (o *CustomGetParams) SetResourceName(resourceName string) {
	o.ResourceName = resourceName
}

// WriteToRequest writes these params to a swagger request
func (o *CustomGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param resourceName
	if err := r.SetPathParam("resourceName", o.ResourceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
