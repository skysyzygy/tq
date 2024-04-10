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

// NewProductionsGetParams creates a new ProductionsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProductionsGetParams() *ProductionsGetParams {
	return &ProductionsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProductionsGetParamsWithTimeout creates a new ProductionsGetParams object
// with the ability to set a timeout on a request.
func NewProductionsGetParamsWithTimeout(timeout time.Duration) *ProductionsGetParams {
	return &ProductionsGetParams{
		timeout: timeout,
	}
}

// NewProductionsGetParamsWithContext creates a new ProductionsGetParams object
// with the ability to set a context for a request.
func NewProductionsGetParamsWithContext(ctx context.Context) *ProductionsGetParams {
	return &ProductionsGetParams{
		Context: ctx,
	}
}

// NewProductionsGetParamsWithHTTPClient creates a new ProductionsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewProductionsGetParamsWithHTTPClient(client *http.Client) *ProductionsGetParams {
	return &ProductionsGetParams{
		HTTPClient: client,
	}
}

/*
ProductionsGetParams contains all the parameters to send to the API endpoint

	for the productions get operation.

	Typically these are written to a http.Request.
*/
type ProductionsGetParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the productions get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProductionsGetParams) WithDefaults() *ProductionsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the productions get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProductionsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the productions get params
func (o *ProductionsGetParams) WithTimeout(timeout time.Duration) *ProductionsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the productions get params
func (o *ProductionsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the productions get params
func (o *ProductionsGetParams) WithContext(ctx context.Context) *ProductionsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the productions get params
func (o *ProductionsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the productions get params
func (o *ProductionsGetParams) WithHTTPClient(client *http.Client) *ProductionsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the productions get params
func (o *ProductionsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the productions get params
func (o *ProductionsGetParams) WithID(id string) *ProductionsGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the productions get params
func (o *ProductionsGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ProductionsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
