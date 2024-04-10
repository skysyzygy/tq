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

// NewPriceTypeReasonsCreateParams creates a new PriceTypeReasonsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPriceTypeReasonsCreateParams() *PriceTypeReasonsCreateParams {
	return &PriceTypeReasonsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPriceTypeReasonsCreateParamsWithTimeout creates a new PriceTypeReasonsCreateParams object
// with the ability to set a timeout on a request.
func NewPriceTypeReasonsCreateParamsWithTimeout(timeout time.Duration) *PriceTypeReasonsCreateParams {
	return &PriceTypeReasonsCreateParams{
		timeout: timeout,
	}
}

// NewPriceTypeReasonsCreateParamsWithContext creates a new PriceTypeReasonsCreateParams object
// with the ability to set a context for a request.
func NewPriceTypeReasonsCreateParamsWithContext(ctx context.Context) *PriceTypeReasonsCreateParams {
	return &PriceTypeReasonsCreateParams{
		Context: ctx,
	}
}

// NewPriceTypeReasonsCreateParamsWithHTTPClient creates a new PriceTypeReasonsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPriceTypeReasonsCreateParamsWithHTTPClient(client *http.Client) *PriceTypeReasonsCreateParams {
	return &PriceTypeReasonsCreateParams{
		HTTPClient: client,
	}
}

/*
PriceTypeReasonsCreateParams contains all the parameters to send to the API endpoint

	for the price type reasons create operation.

	Typically these are written to a http.Request.
*/
type PriceTypeReasonsCreateParams struct {

	// Data.
	Data *models.PriceTypeReason

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the price type reasons create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PriceTypeReasonsCreateParams) WithDefaults() *PriceTypeReasonsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the price type reasons create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PriceTypeReasonsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the price type reasons create params
func (o *PriceTypeReasonsCreateParams) WithTimeout(timeout time.Duration) *PriceTypeReasonsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the price type reasons create params
func (o *PriceTypeReasonsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the price type reasons create params
func (o *PriceTypeReasonsCreateParams) WithContext(ctx context.Context) *PriceTypeReasonsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the price type reasons create params
func (o *PriceTypeReasonsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the price type reasons create params
func (o *PriceTypeReasonsCreateParams) WithHTTPClient(client *http.Client) *PriceTypeReasonsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the price type reasons create params
func (o *PriceTypeReasonsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the price type reasons create params
func (o *PriceTypeReasonsCreateParams) WithData(data *models.PriceTypeReason) *PriceTypeReasonsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the price type reasons create params
func (o *PriceTypeReasonsCreateParams) SetData(data *models.PriceTypeReason) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PriceTypeReasonsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
