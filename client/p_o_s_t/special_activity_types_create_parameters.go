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

// NewSpecialActivityTypesCreateParams creates a new SpecialActivityTypesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSpecialActivityTypesCreateParams() *SpecialActivityTypesCreateParams {
	return &SpecialActivityTypesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSpecialActivityTypesCreateParamsWithTimeout creates a new SpecialActivityTypesCreateParams object
// with the ability to set a timeout on a request.
func NewSpecialActivityTypesCreateParamsWithTimeout(timeout time.Duration) *SpecialActivityTypesCreateParams {
	return &SpecialActivityTypesCreateParams{
		timeout: timeout,
	}
}

// NewSpecialActivityTypesCreateParamsWithContext creates a new SpecialActivityTypesCreateParams object
// with the ability to set a context for a request.
func NewSpecialActivityTypesCreateParamsWithContext(ctx context.Context) *SpecialActivityTypesCreateParams {
	return &SpecialActivityTypesCreateParams{
		Context: ctx,
	}
}

// NewSpecialActivityTypesCreateParamsWithHTTPClient creates a new SpecialActivityTypesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSpecialActivityTypesCreateParamsWithHTTPClient(client *http.Client) *SpecialActivityTypesCreateParams {
	return &SpecialActivityTypesCreateParams{
		HTTPClient: client,
	}
}

/*
SpecialActivityTypesCreateParams contains all the parameters to send to the API endpoint

	for the special activity types create operation.

	Typically these are written to a http.Request.
*/
type SpecialActivityTypesCreateParams struct {

	/* Data.

	   The resource to be created
	*/
	Data *models.SpecialActivityType

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the special activity types create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SpecialActivityTypesCreateParams) WithDefaults() *SpecialActivityTypesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the special activity types create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SpecialActivityTypesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the special activity types create params
func (o *SpecialActivityTypesCreateParams) WithTimeout(timeout time.Duration) *SpecialActivityTypesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the special activity types create params
func (o *SpecialActivityTypesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the special activity types create params
func (o *SpecialActivityTypesCreateParams) WithContext(ctx context.Context) *SpecialActivityTypesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the special activity types create params
func (o *SpecialActivityTypesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the special activity types create params
func (o *SpecialActivityTypesCreateParams) WithHTTPClient(client *http.Client) *SpecialActivityTypesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the special activity types create params
func (o *SpecialActivityTypesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the special activity types create params
func (o *SpecialActivityTypesCreateParams) WithData(data *models.SpecialActivityType) *SpecialActivityTypesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the special activity types create params
func (o *SpecialActivityTypesCreateParams) SetData(data *models.SpecialActivityType) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *SpecialActivityTypesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
