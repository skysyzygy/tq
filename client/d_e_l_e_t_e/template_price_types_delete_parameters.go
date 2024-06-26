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

// NewTemplatePriceTypesDeleteParams creates a new TemplatePriceTypesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTemplatePriceTypesDeleteParams() *TemplatePriceTypesDeleteParams {
	return &TemplatePriceTypesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTemplatePriceTypesDeleteParamsWithTimeout creates a new TemplatePriceTypesDeleteParams object
// with the ability to set a timeout on a request.
func NewTemplatePriceTypesDeleteParamsWithTimeout(timeout time.Duration) *TemplatePriceTypesDeleteParams {
	return &TemplatePriceTypesDeleteParams{
		timeout: timeout,
	}
}

// NewTemplatePriceTypesDeleteParamsWithContext creates a new TemplatePriceTypesDeleteParams object
// with the ability to set a context for a request.
func NewTemplatePriceTypesDeleteParamsWithContext(ctx context.Context) *TemplatePriceTypesDeleteParams {
	return &TemplatePriceTypesDeleteParams{
		Context: ctx,
	}
}

// NewTemplatePriceTypesDeleteParamsWithHTTPClient creates a new TemplatePriceTypesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewTemplatePriceTypesDeleteParamsWithHTTPClient(client *http.Client) *TemplatePriceTypesDeleteParams {
	return &TemplatePriceTypesDeleteParams{
		HTTPClient: client,
	}
}

/*
TemplatePriceTypesDeleteParams contains all the parameters to send to the API endpoint

	for the template price types delete operation.

	Typically these are written to a http.Request.
*/
type TemplatePriceTypesDeleteParams struct {

	// TemplatePriceTypeID.
	TemplatePriceTypeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the template price types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TemplatePriceTypesDeleteParams) WithDefaults() *TemplatePriceTypesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the template price types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TemplatePriceTypesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the template price types delete params
func (o *TemplatePriceTypesDeleteParams) WithTimeout(timeout time.Duration) *TemplatePriceTypesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the template price types delete params
func (o *TemplatePriceTypesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the template price types delete params
func (o *TemplatePriceTypesDeleteParams) WithContext(ctx context.Context) *TemplatePriceTypesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the template price types delete params
func (o *TemplatePriceTypesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the template price types delete params
func (o *TemplatePriceTypesDeleteParams) WithHTTPClient(client *http.Client) *TemplatePriceTypesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the template price types delete params
func (o *TemplatePriceTypesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTemplatePriceTypeID adds the templatePriceTypeID to the template price types delete params
func (o *TemplatePriceTypesDeleteParams) WithTemplatePriceTypeID(templatePriceTypeID string) *TemplatePriceTypesDeleteParams {
	o.SetTemplatePriceTypeID(templatePriceTypeID)
	return o
}

// SetTemplatePriceTypeID adds the templatePriceTypeId to the template price types delete params
func (o *TemplatePriceTypesDeleteParams) SetTemplatePriceTypeID(templatePriceTypeID string) {
	o.TemplatePriceTypeID = templatePriceTypeID
}

// WriteToRequest writes these params to a swagger request
func (o *TemplatePriceTypesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param templatePriceTypeId
	if err := r.SetPathParam("templatePriceTypeId", o.TemplatePriceTypeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
