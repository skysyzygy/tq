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

	"github.com/skysyzygy/tq/models"
)

// NewProductsDescribeParams creates a new ProductsDescribeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProductsDescribeParams() *ProductsDescribeParams {
	return &ProductsDescribeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProductsDescribeParamsWithTimeout creates a new ProductsDescribeParams object
// with the ability to set a timeout on a request.
func NewProductsDescribeParamsWithTimeout(timeout time.Duration) *ProductsDescribeParams {
	return &ProductsDescribeParams{
		timeout: timeout,
	}
}

// NewProductsDescribeParamsWithContext creates a new ProductsDescribeParams object
// with the ability to set a context for a request.
func NewProductsDescribeParamsWithContext(ctx context.Context) *ProductsDescribeParams {
	return &ProductsDescribeParams{
		Context: ctx,
	}
}

// NewProductsDescribeParamsWithHTTPClient creates a new ProductsDescribeParams object
// with the ability to set a custom HTTPClient for a request.
func NewProductsDescribeParamsWithHTTPClient(client *http.Client) *ProductsDescribeParams {
	return &ProductsDescribeParams{
		HTTPClient: client,
	}
}

/*
ProductsDescribeParams contains all the parameters to send to the API endpoint

	for the products describe operation.

	Typically these are written to a http.Request.
*/
type ProductsDescribeParams struct {

	// Request.
	Request *models.ProductDescribeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the products describe params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProductsDescribeParams) WithDefaults() *ProductsDescribeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the products describe params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProductsDescribeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the products describe params
func (o *ProductsDescribeParams) WithTimeout(timeout time.Duration) *ProductsDescribeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the products describe params
func (o *ProductsDescribeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the products describe params
func (o *ProductsDescribeParams) WithContext(ctx context.Context) *ProductsDescribeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the products describe params
func (o *ProductsDescribeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the products describe params
func (o *ProductsDescribeParams) WithHTTPClient(client *http.Client) *ProductsDescribeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the products describe params
func (o *ProductsDescribeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the products describe params
func (o *ProductsDescribeParams) WithRequest(request *models.ProductDescribeRequest) *ProductsDescribeParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the products describe params
func (o *ProductsDescribeParams) SetRequest(request *models.ProductDescribeRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *ProductsDescribeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
