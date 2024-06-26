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

// NewNameStatusesCreateParams creates a new NameStatusesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNameStatusesCreateParams() *NameStatusesCreateParams {
	return &NameStatusesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNameStatusesCreateParamsWithTimeout creates a new NameStatusesCreateParams object
// with the ability to set a timeout on a request.
func NewNameStatusesCreateParamsWithTimeout(timeout time.Duration) *NameStatusesCreateParams {
	return &NameStatusesCreateParams{
		timeout: timeout,
	}
}

// NewNameStatusesCreateParamsWithContext creates a new NameStatusesCreateParams object
// with the ability to set a context for a request.
func NewNameStatusesCreateParamsWithContext(ctx context.Context) *NameStatusesCreateParams {
	return &NameStatusesCreateParams{
		Context: ctx,
	}
}

// NewNameStatusesCreateParamsWithHTTPClient creates a new NameStatusesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewNameStatusesCreateParamsWithHTTPClient(client *http.Client) *NameStatusesCreateParams {
	return &NameStatusesCreateParams{
		HTTPClient: client,
	}
}

/*
NameStatusesCreateParams contains all the parameters to send to the API endpoint

	for the name statuses create operation.

	Typically these are written to a http.Request.
*/
type NameStatusesCreateParams struct {

	// Data.
	Data *models.NameStatus

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the name statuses create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NameStatusesCreateParams) WithDefaults() *NameStatusesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the name statuses create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NameStatusesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the name statuses create params
func (o *NameStatusesCreateParams) WithTimeout(timeout time.Duration) *NameStatusesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the name statuses create params
func (o *NameStatusesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the name statuses create params
func (o *NameStatusesCreateParams) WithContext(ctx context.Context) *NameStatusesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the name statuses create params
func (o *NameStatusesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the name statuses create params
func (o *NameStatusesCreateParams) WithHTTPClient(client *http.Client) *NameStatusesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the name statuses create params
func (o *NameStatusesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the name statuses create params
func (o *NameStatusesCreateParams) WithData(data *models.NameStatus) *NameStatusesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the name statuses create params
func (o *NameStatusesCreateParams) SetData(data *models.NameStatus) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *NameStatusesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
