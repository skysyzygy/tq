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

// NewSeatCodesUpdateParams creates a new SeatCodesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSeatCodesUpdateParams() *SeatCodesUpdateParams {
	return &SeatCodesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSeatCodesUpdateParamsWithTimeout creates a new SeatCodesUpdateParams object
// with the ability to set a timeout on a request.
func NewSeatCodesUpdateParamsWithTimeout(timeout time.Duration) *SeatCodesUpdateParams {
	return &SeatCodesUpdateParams{
		timeout: timeout,
	}
}

// NewSeatCodesUpdateParamsWithContext creates a new SeatCodesUpdateParams object
// with the ability to set a context for a request.
func NewSeatCodesUpdateParamsWithContext(ctx context.Context) *SeatCodesUpdateParams {
	return &SeatCodesUpdateParams{
		Context: ctx,
	}
}

// NewSeatCodesUpdateParamsWithHTTPClient creates a new SeatCodesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSeatCodesUpdateParamsWithHTTPClient(client *http.Client) *SeatCodesUpdateParams {
	return &SeatCodesUpdateParams{
		HTTPClient: client,
	}
}

/*
SeatCodesUpdateParams contains all the parameters to send to the API endpoint

	for the seat codes update operation.

	Typically these are written to a http.Request.
*/
type SeatCodesUpdateParams struct {

	// Data.
	Data *models.SeatCode

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the seat codes update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SeatCodesUpdateParams) WithDefaults() *SeatCodesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the seat codes update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SeatCodesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the seat codes update params
func (o *SeatCodesUpdateParams) WithTimeout(timeout time.Duration) *SeatCodesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the seat codes update params
func (o *SeatCodesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the seat codes update params
func (o *SeatCodesUpdateParams) WithContext(ctx context.Context) *SeatCodesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the seat codes update params
func (o *SeatCodesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the seat codes update params
func (o *SeatCodesUpdateParams) WithHTTPClient(client *http.Client) *SeatCodesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the seat codes update params
func (o *SeatCodesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the seat codes update params
func (o *SeatCodesUpdateParams) WithData(data *models.SeatCode) *SeatCodesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the seat codes update params
func (o *SeatCodesUpdateParams) SetData(data *models.SeatCode) {
	o.Data = data
}

// WithID adds the id to the seat codes update params
func (o *SeatCodesUpdateParams) WithID(id string) *SeatCodesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the seat codes update params
func (o *SeatCodesUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SeatCodesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
