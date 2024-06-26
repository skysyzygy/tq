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

// NewBillingTypesCreateParams creates a new BillingTypesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBillingTypesCreateParams() *BillingTypesCreateParams {
	return &BillingTypesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBillingTypesCreateParamsWithTimeout creates a new BillingTypesCreateParams object
// with the ability to set a timeout on a request.
func NewBillingTypesCreateParamsWithTimeout(timeout time.Duration) *BillingTypesCreateParams {
	return &BillingTypesCreateParams{
		timeout: timeout,
	}
}

// NewBillingTypesCreateParamsWithContext creates a new BillingTypesCreateParams object
// with the ability to set a context for a request.
func NewBillingTypesCreateParamsWithContext(ctx context.Context) *BillingTypesCreateParams {
	return &BillingTypesCreateParams{
		Context: ctx,
	}
}

// NewBillingTypesCreateParamsWithHTTPClient creates a new BillingTypesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewBillingTypesCreateParamsWithHTTPClient(client *http.Client) *BillingTypesCreateParams {
	return &BillingTypesCreateParams{
		HTTPClient: client,
	}
}

/*
BillingTypesCreateParams contains all the parameters to send to the API endpoint

	for the billing types create operation.

	Typically these are written to a http.Request.
*/
type BillingTypesCreateParams struct {

	/* Data.

	   The resource to be created
	*/
	Data *models.BillingType

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the billing types create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BillingTypesCreateParams) WithDefaults() *BillingTypesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the billing types create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BillingTypesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the billing types create params
func (o *BillingTypesCreateParams) WithTimeout(timeout time.Duration) *BillingTypesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the billing types create params
func (o *BillingTypesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the billing types create params
func (o *BillingTypesCreateParams) WithContext(ctx context.Context) *BillingTypesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the billing types create params
func (o *BillingTypesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the billing types create params
func (o *BillingTypesCreateParams) WithHTTPClient(client *http.Client) *BillingTypesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the billing types create params
func (o *BillingTypesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the billing types create params
func (o *BillingTypesCreateParams) WithData(data *models.BillingType) *BillingTypesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the billing types create params
func (o *BillingTypesCreateParams) SetData(data *models.BillingType) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *BillingTypesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
