// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

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

// NewOriginalSourcesUpdateParams creates a new OriginalSourcesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOriginalSourcesUpdateParams() *OriginalSourcesUpdateParams {
	return &OriginalSourcesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOriginalSourcesUpdateParamsWithTimeout creates a new OriginalSourcesUpdateParams object
// with the ability to set a timeout on a request.
func NewOriginalSourcesUpdateParamsWithTimeout(timeout time.Duration) *OriginalSourcesUpdateParams {
	return &OriginalSourcesUpdateParams{
		timeout: timeout,
	}
}

// NewOriginalSourcesUpdateParamsWithContext creates a new OriginalSourcesUpdateParams object
// with the ability to set a context for a request.
func NewOriginalSourcesUpdateParamsWithContext(ctx context.Context) *OriginalSourcesUpdateParams {
	return &OriginalSourcesUpdateParams{
		Context: ctx,
	}
}

// NewOriginalSourcesUpdateParamsWithHTTPClient creates a new OriginalSourcesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewOriginalSourcesUpdateParamsWithHTTPClient(client *http.Client) *OriginalSourcesUpdateParams {
	return &OriginalSourcesUpdateParams{
		HTTPClient: client,
	}
}

/*
OriginalSourcesUpdateParams contains all the parameters to send to the API endpoint

	for the original sources update operation.

	Typically these are written to a http.Request.
*/
type OriginalSourcesUpdateParams struct {

	// Data.
	Data *models.OriginalSource

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the original sources update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OriginalSourcesUpdateParams) WithDefaults() *OriginalSourcesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the original sources update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OriginalSourcesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the original sources update params
func (o *OriginalSourcesUpdateParams) WithTimeout(timeout time.Duration) *OriginalSourcesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the original sources update params
func (o *OriginalSourcesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the original sources update params
func (o *OriginalSourcesUpdateParams) WithContext(ctx context.Context) *OriginalSourcesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the original sources update params
func (o *OriginalSourcesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the original sources update params
func (o *OriginalSourcesUpdateParams) WithHTTPClient(client *http.Client) *OriginalSourcesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the original sources update params
func (o *OriginalSourcesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the original sources update params
func (o *OriginalSourcesUpdateParams) WithData(data *models.OriginalSource) *OriginalSourcesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the original sources update params
func (o *OriginalSourcesUpdateParams) SetData(data *models.OriginalSource) {
	o.Data = data
}

// WithID adds the id to the original sources update params
func (o *OriginalSourcesUpdateParams) WithID(id string) *OriginalSourcesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the original sources update params
func (o *OriginalSourcesUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *OriginalSourcesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
