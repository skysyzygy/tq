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

// NewPriceTemplatesDeleteParams creates a new PriceTemplatesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPriceTemplatesDeleteParams() *PriceTemplatesDeleteParams {
	return &PriceTemplatesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPriceTemplatesDeleteParamsWithTimeout creates a new PriceTemplatesDeleteParams object
// with the ability to set a timeout on a request.
func NewPriceTemplatesDeleteParamsWithTimeout(timeout time.Duration) *PriceTemplatesDeleteParams {
	return &PriceTemplatesDeleteParams{
		timeout: timeout,
	}
}

// NewPriceTemplatesDeleteParamsWithContext creates a new PriceTemplatesDeleteParams object
// with the ability to set a context for a request.
func NewPriceTemplatesDeleteParamsWithContext(ctx context.Context) *PriceTemplatesDeleteParams {
	return &PriceTemplatesDeleteParams{
		Context: ctx,
	}
}

// NewPriceTemplatesDeleteParamsWithHTTPClient creates a new PriceTemplatesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPriceTemplatesDeleteParamsWithHTTPClient(client *http.Client) *PriceTemplatesDeleteParams {
	return &PriceTemplatesDeleteParams{
		HTTPClient: client,
	}
}

/*
PriceTemplatesDeleteParams contains all the parameters to send to the API endpoint

	for the price templates delete operation.

	Typically these are written to a http.Request.
*/
type PriceTemplatesDeleteParams struct {

	// PriceTemplateID.
	PriceTemplateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the price templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PriceTemplatesDeleteParams) WithDefaults() *PriceTemplatesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the price templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PriceTemplatesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the price templates delete params
func (o *PriceTemplatesDeleteParams) WithTimeout(timeout time.Duration) *PriceTemplatesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the price templates delete params
func (o *PriceTemplatesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the price templates delete params
func (o *PriceTemplatesDeleteParams) WithContext(ctx context.Context) *PriceTemplatesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the price templates delete params
func (o *PriceTemplatesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the price templates delete params
func (o *PriceTemplatesDeleteParams) WithHTTPClient(client *http.Client) *PriceTemplatesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the price templates delete params
func (o *PriceTemplatesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPriceTemplateID adds the priceTemplateID to the price templates delete params
func (o *PriceTemplatesDeleteParams) WithPriceTemplateID(priceTemplateID string) *PriceTemplatesDeleteParams {
	o.SetPriceTemplateID(priceTemplateID)
	return o
}

// SetPriceTemplateID adds the priceTemplateId to the price templates delete params
func (o *PriceTemplatesDeleteParams) SetPriceTemplateID(priceTemplateID string) {
	o.PriceTemplateID = priceTemplateID
}

// WriteToRequest writes these params to a swagger request
func (o *PriceTemplatesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param priceTemplateId
	if err := r.SetPathParam("priceTemplateId", o.PriceTemplateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
