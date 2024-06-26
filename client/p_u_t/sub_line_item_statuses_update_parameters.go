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

// NewSubLineItemStatusesUpdateParams creates a new SubLineItemStatusesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubLineItemStatusesUpdateParams() *SubLineItemStatusesUpdateParams {
	return &SubLineItemStatusesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubLineItemStatusesUpdateParamsWithTimeout creates a new SubLineItemStatusesUpdateParams object
// with the ability to set a timeout on a request.
func NewSubLineItemStatusesUpdateParamsWithTimeout(timeout time.Duration) *SubLineItemStatusesUpdateParams {
	return &SubLineItemStatusesUpdateParams{
		timeout: timeout,
	}
}

// NewSubLineItemStatusesUpdateParamsWithContext creates a new SubLineItemStatusesUpdateParams object
// with the ability to set a context for a request.
func NewSubLineItemStatusesUpdateParamsWithContext(ctx context.Context) *SubLineItemStatusesUpdateParams {
	return &SubLineItemStatusesUpdateParams{
		Context: ctx,
	}
}

// NewSubLineItemStatusesUpdateParamsWithHTTPClient creates a new SubLineItemStatusesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubLineItemStatusesUpdateParamsWithHTTPClient(client *http.Client) *SubLineItemStatusesUpdateParams {
	return &SubLineItemStatusesUpdateParams{
		HTTPClient: client,
	}
}

/*
SubLineItemStatusesUpdateParams contains all the parameters to send to the API endpoint

	for the sub line item statuses update operation.

	Typically these are written to a http.Request.
*/
type SubLineItemStatusesUpdateParams struct {

	// Data.
	Data *models.SubLineItemStatus

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the sub line item statuses update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubLineItemStatusesUpdateParams) WithDefaults() *SubLineItemStatusesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the sub line item statuses update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubLineItemStatusesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the sub line item statuses update params
func (o *SubLineItemStatusesUpdateParams) WithTimeout(timeout time.Duration) *SubLineItemStatusesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sub line item statuses update params
func (o *SubLineItemStatusesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sub line item statuses update params
func (o *SubLineItemStatusesUpdateParams) WithContext(ctx context.Context) *SubLineItemStatusesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sub line item statuses update params
func (o *SubLineItemStatusesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sub line item statuses update params
func (o *SubLineItemStatusesUpdateParams) WithHTTPClient(client *http.Client) *SubLineItemStatusesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sub line item statuses update params
func (o *SubLineItemStatusesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the sub line item statuses update params
func (o *SubLineItemStatusesUpdateParams) WithData(data *models.SubLineItemStatus) *SubLineItemStatusesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the sub line item statuses update params
func (o *SubLineItemStatusesUpdateParams) SetData(data *models.SubLineItemStatus) {
	o.Data = data
}

// WithID adds the id to the sub line item statuses update params
func (o *SubLineItemStatusesUpdateParams) WithID(id string) *SubLineItemStatusesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the sub line item statuses update params
func (o *SubLineItemStatusesUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SubLineItemStatusesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
