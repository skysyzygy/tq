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

// NewPrintersCreateParams creates a new PrintersCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPrintersCreateParams() *PrintersCreateParams {
	return &PrintersCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPrintersCreateParamsWithTimeout creates a new PrintersCreateParams object
// with the ability to set a timeout on a request.
func NewPrintersCreateParamsWithTimeout(timeout time.Duration) *PrintersCreateParams {
	return &PrintersCreateParams{
		timeout: timeout,
	}
}

// NewPrintersCreateParamsWithContext creates a new PrintersCreateParams object
// with the ability to set a context for a request.
func NewPrintersCreateParamsWithContext(ctx context.Context) *PrintersCreateParams {
	return &PrintersCreateParams{
		Context: ctx,
	}
}

// NewPrintersCreateParamsWithHTTPClient creates a new PrintersCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPrintersCreateParamsWithHTTPClient(client *http.Client) *PrintersCreateParams {
	return &PrintersCreateParams{
		HTTPClient: client,
	}
}

/*
PrintersCreateParams contains all the parameters to send to the API endpoint

	for the printers create operation.

	Typically these are written to a http.Request.
*/
type PrintersCreateParams struct {

	// Data.
	Data *models.Printer

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the printers create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PrintersCreateParams) WithDefaults() *PrintersCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the printers create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PrintersCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the printers create params
func (o *PrintersCreateParams) WithTimeout(timeout time.Duration) *PrintersCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the printers create params
func (o *PrintersCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the printers create params
func (o *PrintersCreateParams) WithContext(ctx context.Context) *PrintersCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the printers create params
func (o *PrintersCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the printers create params
func (o *PrintersCreateParams) WithHTTPClient(client *http.Client) *PrintersCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the printers create params
func (o *PrintersCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the printers create params
func (o *PrintersCreateParams) WithData(data *models.Printer) *PrintersCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the printers create params
func (o *PrintersCreateParams) SetData(data *models.Printer) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PrintersCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
