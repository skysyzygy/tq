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

// NewPhilanthropyTypesUpdateParams creates a new PhilanthropyTypesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPhilanthropyTypesUpdateParams() *PhilanthropyTypesUpdateParams {
	return &PhilanthropyTypesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPhilanthropyTypesUpdateParamsWithTimeout creates a new PhilanthropyTypesUpdateParams object
// with the ability to set a timeout on a request.
func NewPhilanthropyTypesUpdateParamsWithTimeout(timeout time.Duration) *PhilanthropyTypesUpdateParams {
	return &PhilanthropyTypesUpdateParams{
		timeout: timeout,
	}
}

// NewPhilanthropyTypesUpdateParamsWithContext creates a new PhilanthropyTypesUpdateParams object
// with the ability to set a context for a request.
func NewPhilanthropyTypesUpdateParamsWithContext(ctx context.Context) *PhilanthropyTypesUpdateParams {
	return &PhilanthropyTypesUpdateParams{
		Context: ctx,
	}
}

// NewPhilanthropyTypesUpdateParamsWithHTTPClient creates a new PhilanthropyTypesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPhilanthropyTypesUpdateParamsWithHTTPClient(client *http.Client) *PhilanthropyTypesUpdateParams {
	return &PhilanthropyTypesUpdateParams{
		HTTPClient: client,
	}
}

/*
PhilanthropyTypesUpdateParams contains all the parameters to send to the API endpoint

	for the philanthropy types update operation.

	Typically these are written to a http.Request.
*/
type PhilanthropyTypesUpdateParams struct {

	// Data.
	Data *models.PhilanthropyType

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the philanthropy types update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PhilanthropyTypesUpdateParams) WithDefaults() *PhilanthropyTypesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the philanthropy types update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PhilanthropyTypesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the philanthropy types update params
func (o *PhilanthropyTypesUpdateParams) WithTimeout(timeout time.Duration) *PhilanthropyTypesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the philanthropy types update params
func (o *PhilanthropyTypesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the philanthropy types update params
func (o *PhilanthropyTypesUpdateParams) WithContext(ctx context.Context) *PhilanthropyTypesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the philanthropy types update params
func (o *PhilanthropyTypesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the philanthropy types update params
func (o *PhilanthropyTypesUpdateParams) WithHTTPClient(client *http.Client) *PhilanthropyTypesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the philanthropy types update params
func (o *PhilanthropyTypesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the philanthropy types update params
func (o *PhilanthropyTypesUpdateParams) WithData(data *models.PhilanthropyType) *PhilanthropyTypesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the philanthropy types update params
func (o *PhilanthropyTypesUpdateParams) SetData(data *models.PhilanthropyType) {
	o.Data = data
}

// WithID adds the id to the philanthropy types update params
func (o *PhilanthropyTypesUpdateParams) WithID(id string) *PhilanthropyTypesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the philanthropy types update params
func (o *PhilanthropyTypesUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PhilanthropyTypesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
