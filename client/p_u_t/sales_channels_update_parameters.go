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

// NewSalesChannelsUpdateParams creates a new SalesChannelsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSalesChannelsUpdateParams() *SalesChannelsUpdateParams {
	return &SalesChannelsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSalesChannelsUpdateParamsWithTimeout creates a new SalesChannelsUpdateParams object
// with the ability to set a timeout on a request.
func NewSalesChannelsUpdateParamsWithTimeout(timeout time.Duration) *SalesChannelsUpdateParams {
	return &SalesChannelsUpdateParams{
		timeout: timeout,
	}
}

// NewSalesChannelsUpdateParamsWithContext creates a new SalesChannelsUpdateParams object
// with the ability to set a context for a request.
func NewSalesChannelsUpdateParamsWithContext(ctx context.Context) *SalesChannelsUpdateParams {
	return &SalesChannelsUpdateParams{
		Context: ctx,
	}
}

// NewSalesChannelsUpdateParamsWithHTTPClient creates a new SalesChannelsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSalesChannelsUpdateParamsWithHTTPClient(client *http.Client) *SalesChannelsUpdateParams {
	return &SalesChannelsUpdateParams{
		HTTPClient: client,
	}
}

/*
SalesChannelsUpdateParams contains all the parameters to send to the API endpoint

	for the sales channels update operation.

	Typically these are written to a http.Request.
*/
type SalesChannelsUpdateParams struct {

	// Data.
	Data *models.SalesChannel

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the sales channels update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SalesChannelsUpdateParams) WithDefaults() *SalesChannelsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the sales channels update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SalesChannelsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the sales channels update params
func (o *SalesChannelsUpdateParams) WithTimeout(timeout time.Duration) *SalesChannelsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sales channels update params
func (o *SalesChannelsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sales channels update params
func (o *SalesChannelsUpdateParams) WithContext(ctx context.Context) *SalesChannelsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sales channels update params
func (o *SalesChannelsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sales channels update params
func (o *SalesChannelsUpdateParams) WithHTTPClient(client *http.Client) *SalesChannelsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sales channels update params
func (o *SalesChannelsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the sales channels update params
func (o *SalesChannelsUpdateParams) WithData(data *models.SalesChannel) *SalesChannelsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the sales channels update params
func (o *SalesChannelsUpdateParams) SetData(data *models.SalesChannel) {
	o.Data = data
}

// WithID adds the id to the sales channels update params
func (o *SalesChannelsUpdateParams) WithID(id string) *SalesChannelsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the sales channels update params
func (o *SalesChannelsUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SalesChannelsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
