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

// NewPrefixesUpdateParams creates a new PrefixesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPrefixesUpdateParams() *PrefixesUpdateParams {
	return &PrefixesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPrefixesUpdateParamsWithTimeout creates a new PrefixesUpdateParams object
// with the ability to set a timeout on a request.
func NewPrefixesUpdateParamsWithTimeout(timeout time.Duration) *PrefixesUpdateParams {
	return &PrefixesUpdateParams{
		timeout: timeout,
	}
}

// NewPrefixesUpdateParamsWithContext creates a new PrefixesUpdateParams object
// with the ability to set a context for a request.
func NewPrefixesUpdateParamsWithContext(ctx context.Context) *PrefixesUpdateParams {
	return &PrefixesUpdateParams{
		Context: ctx,
	}
}

// NewPrefixesUpdateParamsWithHTTPClient creates a new PrefixesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPrefixesUpdateParamsWithHTTPClient(client *http.Client) *PrefixesUpdateParams {
	return &PrefixesUpdateParams{
		HTTPClient: client,
	}
}

/*
PrefixesUpdateParams contains all the parameters to send to the API endpoint

	for the prefixes update operation.

	Typically these are written to a http.Request.
*/
type PrefixesUpdateParams struct {

	// Data.
	Data *models.Prefix

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the prefixes update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PrefixesUpdateParams) WithDefaults() *PrefixesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the prefixes update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PrefixesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the prefixes update params
func (o *PrefixesUpdateParams) WithTimeout(timeout time.Duration) *PrefixesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the prefixes update params
func (o *PrefixesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the prefixes update params
func (o *PrefixesUpdateParams) WithContext(ctx context.Context) *PrefixesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the prefixes update params
func (o *PrefixesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the prefixes update params
func (o *PrefixesUpdateParams) WithHTTPClient(client *http.Client) *PrefixesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the prefixes update params
func (o *PrefixesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the prefixes update params
func (o *PrefixesUpdateParams) WithData(data *models.Prefix) *PrefixesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the prefixes update params
func (o *PrefixesUpdateParams) SetData(data *models.Prefix) {
	o.Data = data
}

// WithID adds the id to the prefixes update params
func (o *PrefixesUpdateParams) WithID(id string) *PrefixesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the prefixes update params
func (o *PrefixesUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PrefixesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
