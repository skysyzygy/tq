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

// NewConstituencyTypesCreateParams creates a new ConstituencyTypesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConstituencyTypesCreateParams() *ConstituencyTypesCreateParams {
	return &ConstituencyTypesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConstituencyTypesCreateParamsWithTimeout creates a new ConstituencyTypesCreateParams object
// with the ability to set a timeout on a request.
func NewConstituencyTypesCreateParamsWithTimeout(timeout time.Duration) *ConstituencyTypesCreateParams {
	return &ConstituencyTypesCreateParams{
		timeout: timeout,
	}
}

// NewConstituencyTypesCreateParamsWithContext creates a new ConstituencyTypesCreateParams object
// with the ability to set a context for a request.
func NewConstituencyTypesCreateParamsWithContext(ctx context.Context) *ConstituencyTypesCreateParams {
	return &ConstituencyTypesCreateParams{
		Context: ctx,
	}
}

// NewConstituencyTypesCreateParamsWithHTTPClient creates a new ConstituencyTypesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewConstituencyTypesCreateParamsWithHTTPClient(client *http.Client) *ConstituencyTypesCreateParams {
	return &ConstituencyTypesCreateParams{
		HTTPClient: client,
	}
}

/*
ConstituencyTypesCreateParams contains all the parameters to send to the API endpoint

	for the constituency types create operation.

	Typically these are written to a http.Request.
*/
type ConstituencyTypesCreateParams struct {

	// Data.
	Data *models.ConstituencyType

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the constituency types create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConstituencyTypesCreateParams) WithDefaults() *ConstituencyTypesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the constituency types create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConstituencyTypesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the constituency types create params
func (o *ConstituencyTypesCreateParams) WithTimeout(timeout time.Duration) *ConstituencyTypesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the constituency types create params
func (o *ConstituencyTypesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the constituency types create params
func (o *ConstituencyTypesCreateParams) WithContext(ctx context.Context) *ConstituencyTypesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the constituency types create params
func (o *ConstituencyTypesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the constituency types create params
func (o *ConstituencyTypesCreateParams) WithHTTPClient(client *http.Client) *ConstituencyTypesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the constituency types create params
func (o *ConstituencyTypesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the constituency types create params
func (o *ConstituencyTypesCreateParams) WithData(data *models.ConstituencyType) *ConstituencyTypesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the constituency types create params
func (o *ConstituencyTypesCreateParams) SetData(data *models.ConstituencyType) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ConstituencyTypesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
