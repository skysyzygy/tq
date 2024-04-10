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

// NewTemplatePriceTypesBatchCreateParams creates a new TemplatePriceTypesBatchCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTemplatePriceTypesBatchCreateParams() *TemplatePriceTypesBatchCreateParams {
	return &TemplatePriceTypesBatchCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTemplatePriceTypesBatchCreateParamsWithTimeout creates a new TemplatePriceTypesBatchCreateParams object
// with the ability to set a timeout on a request.
func NewTemplatePriceTypesBatchCreateParamsWithTimeout(timeout time.Duration) *TemplatePriceTypesBatchCreateParams {
	return &TemplatePriceTypesBatchCreateParams{
		timeout: timeout,
	}
}

// NewTemplatePriceTypesBatchCreateParamsWithContext creates a new TemplatePriceTypesBatchCreateParams object
// with the ability to set a context for a request.
func NewTemplatePriceTypesBatchCreateParamsWithContext(ctx context.Context) *TemplatePriceTypesBatchCreateParams {
	return &TemplatePriceTypesBatchCreateParams{
		Context: ctx,
	}
}

// NewTemplatePriceTypesBatchCreateParamsWithHTTPClient creates a new TemplatePriceTypesBatchCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTemplatePriceTypesBatchCreateParamsWithHTTPClient(client *http.Client) *TemplatePriceTypesBatchCreateParams {
	return &TemplatePriceTypesBatchCreateParams{
		HTTPClient: client,
	}
}

/*
TemplatePriceTypesBatchCreateParams contains all the parameters to send to the API endpoint

	for the template price types batch create operation.

	Typically these are written to a http.Request.
*/
type TemplatePriceTypesBatchCreateParams struct {

	// TemplatePriceTypes.
	TemplatePriceTypes []*models.TemplatePriceType

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the template price types batch create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TemplatePriceTypesBatchCreateParams) WithDefaults() *TemplatePriceTypesBatchCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the template price types batch create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TemplatePriceTypesBatchCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the template price types batch create params
func (o *TemplatePriceTypesBatchCreateParams) WithTimeout(timeout time.Duration) *TemplatePriceTypesBatchCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the template price types batch create params
func (o *TemplatePriceTypesBatchCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the template price types batch create params
func (o *TemplatePriceTypesBatchCreateParams) WithContext(ctx context.Context) *TemplatePriceTypesBatchCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the template price types batch create params
func (o *TemplatePriceTypesBatchCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the template price types batch create params
func (o *TemplatePriceTypesBatchCreateParams) WithHTTPClient(client *http.Client) *TemplatePriceTypesBatchCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the template price types batch create params
func (o *TemplatePriceTypesBatchCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTemplatePriceTypes adds the templatePriceTypes to the template price types batch create params
func (o *TemplatePriceTypesBatchCreateParams) WithTemplatePriceTypes(templatePriceTypes []*models.TemplatePriceType) *TemplatePriceTypesBatchCreateParams {
	o.SetTemplatePriceTypes(templatePriceTypes)
	return o
}

// SetTemplatePriceTypes adds the templatePriceTypes to the template price types batch create params
func (o *TemplatePriceTypesBatchCreateParams) SetTemplatePriceTypes(templatePriceTypes []*models.TemplatePriceType) {
	o.TemplatePriceTypes = templatePriceTypes
}

// WriteToRequest writes these params to a swagger request
func (o *TemplatePriceTypesBatchCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.TemplatePriceTypes != nil {
		if err := r.SetBodyParam(o.TemplatePriceTypes); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
