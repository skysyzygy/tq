// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

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
)

// NewSeatStatusesDeleteParams creates a new SeatStatusesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSeatStatusesDeleteParams() *SeatStatusesDeleteParams {
	return &SeatStatusesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSeatStatusesDeleteParamsWithTimeout creates a new SeatStatusesDeleteParams object
// with the ability to set a timeout on a request.
func NewSeatStatusesDeleteParamsWithTimeout(timeout time.Duration) *SeatStatusesDeleteParams {
	return &SeatStatusesDeleteParams{
		timeout: timeout,
	}
}

// NewSeatStatusesDeleteParamsWithContext creates a new SeatStatusesDeleteParams object
// with the ability to set a context for a request.
func NewSeatStatusesDeleteParamsWithContext(ctx context.Context) *SeatStatusesDeleteParams {
	return &SeatStatusesDeleteParams{
		Context: ctx,
	}
}

// NewSeatStatusesDeleteParamsWithHTTPClient creates a new SeatStatusesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewSeatStatusesDeleteParamsWithHTTPClient(client *http.Client) *SeatStatusesDeleteParams {
	return &SeatStatusesDeleteParams{
		HTTPClient: client,
	}
}

/*
SeatStatusesDeleteParams contains all the parameters to send to the API endpoint

	for the seat statuses delete operation.

	Typically these are written to a http.Request.
*/
type SeatStatusesDeleteParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the seat statuses delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SeatStatusesDeleteParams) WithDefaults() *SeatStatusesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the seat statuses delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SeatStatusesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the seat statuses delete params
func (o *SeatStatusesDeleteParams) WithTimeout(timeout time.Duration) *SeatStatusesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the seat statuses delete params
func (o *SeatStatusesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the seat statuses delete params
func (o *SeatStatusesDeleteParams) WithContext(ctx context.Context) *SeatStatusesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the seat statuses delete params
func (o *SeatStatusesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the seat statuses delete params
func (o *SeatStatusesDeleteParams) WithHTTPClient(client *http.Client) *SeatStatusesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the seat statuses delete params
func (o *SeatStatusesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the seat statuses delete params
func (o *SeatStatusesDeleteParams) WithID(id string) *SeatStatusesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the seat statuses delete params
func (o *SeatStatusesDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SeatStatusesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
