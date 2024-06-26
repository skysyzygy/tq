// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

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

// NewCustomCreateParams creates a new CustomCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomCreateParams() *CustomCreateParams {
	return &CustomCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomCreateParamsWithTimeout creates a new CustomCreateParams object
// with the ability to set a timeout on a request.
func NewCustomCreateParamsWithTimeout(timeout time.Duration) *CustomCreateParams {
	return &CustomCreateParams{
		timeout: timeout,
	}
}

// NewCustomCreateParamsWithContext creates a new CustomCreateParams object
// with the ability to set a context for a request.
func NewCustomCreateParamsWithContext(ctx context.Context) *CustomCreateParams {
	return &CustomCreateParams{
		Context: ctx,
	}
}

// NewCustomCreateParamsWithHTTPClient creates a new CustomCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomCreateParamsWithHTTPClient(client *http.Client) *CustomCreateParams {
	return &CustomCreateParams{
		HTTPClient: client,
	}
}

/*
CustomCreateParams contains all the parameters to send to the API endpoint

	for the custom create operation.

	Typically these are written to a http.Request.
*/
type CustomCreateParams struct {

	// Request.
	Request object

	// ResourceName.
	ResourceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the custom create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomCreateParams) WithDefaults() *CustomCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the custom create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the custom create params
func (o *CustomCreateParams) WithTimeout(timeout time.Duration) *CustomCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom create params
func (o *CustomCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom create params
func (o *CustomCreateParams) WithContext(ctx context.Context) *CustomCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom create params
func (o *CustomCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom create params
func (o *CustomCreateParams) WithHTTPClient(client *http.Client) *CustomCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom create params
func (o *CustomCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the custom create params
func (o *CustomCreateParams) WithRequest(request object) *CustomCreateParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the custom create params
func (o *CustomCreateParams) SetRequest(request object) {
	o.Request = request
}

// WithResourceName adds the resourceName to the custom create params
func (o *CustomCreateParams) WithResourceName(resourceName string) *CustomCreateParams {
	o.SetResourceName(resourceName)
	return o
}

// SetResourceName adds the resourceName to the custom create params
func (o *CustomCreateParams) SetResourceName(resourceName string) {
	o.ResourceName = resourceName
}

// WriteToRequest writes these params to a swagger request
func (o *CustomCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param request
	qrRequest := o.Request
	qRequest := qrRequest
	if qRequest != "" {

		if err := r.SetQueryParam("request", string(qRequest)); err != nil {
			return err
		}
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
