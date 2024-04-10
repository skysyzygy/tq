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

// NewGiftAidRatesUpdateParams creates a new GiftAidRatesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGiftAidRatesUpdateParams() *GiftAidRatesUpdateParams {
	return &GiftAidRatesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGiftAidRatesUpdateParamsWithTimeout creates a new GiftAidRatesUpdateParams object
// with the ability to set a timeout on a request.
func NewGiftAidRatesUpdateParamsWithTimeout(timeout time.Duration) *GiftAidRatesUpdateParams {
	return &GiftAidRatesUpdateParams{
		timeout: timeout,
	}
}

// NewGiftAidRatesUpdateParamsWithContext creates a new GiftAidRatesUpdateParams object
// with the ability to set a context for a request.
func NewGiftAidRatesUpdateParamsWithContext(ctx context.Context) *GiftAidRatesUpdateParams {
	return &GiftAidRatesUpdateParams{
		Context: ctx,
	}
}

// NewGiftAidRatesUpdateParamsWithHTTPClient creates a new GiftAidRatesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGiftAidRatesUpdateParamsWithHTTPClient(client *http.Client) *GiftAidRatesUpdateParams {
	return &GiftAidRatesUpdateParams{
		HTTPClient: client,
	}
}

/*
GiftAidRatesUpdateParams contains all the parameters to send to the API endpoint

	for the gift aid rates update operation.

	Typically these are written to a http.Request.
*/
type GiftAidRatesUpdateParams struct {

	// Data.
	Data *models.GiftAidRate

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the gift aid rates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GiftAidRatesUpdateParams) WithDefaults() *GiftAidRatesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the gift aid rates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GiftAidRatesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the gift aid rates update params
func (o *GiftAidRatesUpdateParams) WithTimeout(timeout time.Duration) *GiftAidRatesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gift aid rates update params
func (o *GiftAidRatesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gift aid rates update params
func (o *GiftAidRatesUpdateParams) WithContext(ctx context.Context) *GiftAidRatesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gift aid rates update params
func (o *GiftAidRatesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gift aid rates update params
func (o *GiftAidRatesUpdateParams) WithHTTPClient(client *http.Client) *GiftAidRatesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gift aid rates update params
func (o *GiftAidRatesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the gift aid rates update params
func (o *GiftAidRatesUpdateParams) WithData(data *models.GiftAidRate) *GiftAidRatesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the gift aid rates update params
func (o *GiftAidRatesUpdateParams) SetData(data *models.GiftAidRate) {
	o.Data = data
}

// WithID adds the id to the gift aid rates update params
func (o *GiftAidRatesUpdateParams) WithID(id string) *GiftAidRatesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the gift aid rates update params
func (o *GiftAidRatesUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GiftAidRatesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
