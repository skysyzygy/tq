// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

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

// NewCustomUpdateParams creates a new CustomUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomUpdateParams() *CustomUpdateParams {
	return &CustomUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomUpdateParamsWithTimeout creates a new CustomUpdateParams object
// with the ability to set a timeout on a request.
func NewCustomUpdateParamsWithTimeout(timeout time.Duration) *CustomUpdateParams {
	return &CustomUpdateParams{
		timeout: timeout,
	}
}

// NewCustomUpdateParamsWithContext creates a new CustomUpdateParams object
// with the ability to set a context for a request.
func NewCustomUpdateParamsWithContext(ctx context.Context) *CustomUpdateParams {
	return &CustomUpdateParams{
		Context: ctx,
	}
}

// NewCustomUpdateParamsWithHTTPClient creates a new CustomUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomUpdateParamsWithHTTPClient(client *http.Client) *CustomUpdateParams {
	return &CustomUpdateParams{
		HTTPClient: client,
	}
}

/*
CustomUpdateParams contains all the parameters to send to the API endpoint

	for the custom update operation.

	Typically these are written to a http.Request.
*/
type CustomUpdateParams struct {

	// ID.
	ID string

	// Request.
	Request object

	// ResourceName.
	ResourceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the custom update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomUpdateParams) WithDefaults() *CustomUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the custom update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the custom update params
func (o *CustomUpdateParams) WithTimeout(timeout time.Duration) *CustomUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom update params
func (o *CustomUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom update params
func (o *CustomUpdateParams) WithContext(ctx context.Context) *CustomUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom update params
func (o *CustomUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom update params
func (o *CustomUpdateParams) WithHTTPClient(client *http.Client) *CustomUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom update params
func (o *CustomUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the custom update params
func (o *CustomUpdateParams) WithID(id string) *CustomUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the custom update params
func (o *CustomUpdateParams) SetID(id string) {
	o.ID = id
}

// WithRequest adds the request to the custom update params
func (o *CustomUpdateParams) WithRequest(request object) *CustomUpdateParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the custom update params
func (o *CustomUpdateParams) SetRequest(request object) {
	o.Request = request
}

// WithResourceName adds the resourceName to the custom update params
func (o *CustomUpdateParams) WithResourceName(resourceName string) *CustomUpdateParams {
	o.SetResourceName(resourceName)
	return o
}

// SetResourceName adds the resourceName to the custom update params
func (o *CustomUpdateParams) SetResourceName(resourceName string) {
	o.ResourceName = resourceName
}

// WriteToRequest writes these params to a swagger request
func (o *CustomUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

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
