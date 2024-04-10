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

// NewPhoneIndicatorsCreateParams creates a new PhoneIndicatorsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPhoneIndicatorsCreateParams() *PhoneIndicatorsCreateParams {
	return &PhoneIndicatorsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPhoneIndicatorsCreateParamsWithTimeout creates a new PhoneIndicatorsCreateParams object
// with the ability to set a timeout on a request.
func NewPhoneIndicatorsCreateParamsWithTimeout(timeout time.Duration) *PhoneIndicatorsCreateParams {
	return &PhoneIndicatorsCreateParams{
		timeout: timeout,
	}
}

// NewPhoneIndicatorsCreateParamsWithContext creates a new PhoneIndicatorsCreateParams object
// with the ability to set a context for a request.
func NewPhoneIndicatorsCreateParamsWithContext(ctx context.Context) *PhoneIndicatorsCreateParams {
	return &PhoneIndicatorsCreateParams{
		Context: ctx,
	}
}

// NewPhoneIndicatorsCreateParamsWithHTTPClient creates a new PhoneIndicatorsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPhoneIndicatorsCreateParamsWithHTTPClient(client *http.Client) *PhoneIndicatorsCreateParams {
	return &PhoneIndicatorsCreateParams{
		HTTPClient: client,
	}
}

/*
PhoneIndicatorsCreateParams contains all the parameters to send to the API endpoint

	for the phone indicators create operation.

	Typically these are written to a http.Request.
*/
type PhoneIndicatorsCreateParams struct {

	// Data.
	Data *models.PhoneIndicator

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the phone indicators create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PhoneIndicatorsCreateParams) WithDefaults() *PhoneIndicatorsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the phone indicators create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PhoneIndicatorsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the phone indicators create params
func (o *PhoneIndicatorsCreateParams) WithTimeout(timeout time.Duration) *PhoneIndicatorsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the phone indicators create params
func (o *PhoneIndicatorsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the phone indicators create params
func (o *PhoneIndicatorsCreateParams) WithContext(ctx context.Context) *PhoneIndicatorsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the phone indicators create params
func (o *PhoneIndicatorsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the phone indicators create params
func (o *PhoneIndicatorsCreateParams) WithHTTPClient(client *http.Client) *PhoneIndicatorsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the phone indicators create params
func (o *PhoneIndicatorsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the phone indicators create params
func (o *PhoneIndicatorsCreateParams) WithData(data *models.PhoneIndicator) *PhoneIndicatorsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the phone indicators create params
func (o *PhoneIndicatorsCreateParams) SetData(data *models.PhoneIndicator) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PhoneIndicatorsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
