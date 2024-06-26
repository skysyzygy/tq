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

// NewBatchTypeGroupsUpdateParams creates a new BatchTypeGroupsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBatchTypeGroupsUpdateParams() *BatchTypeGroupsUpdateParams {
	return &BatchTypeGroupsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBatchTypeGroupsUpdateParamsWithTimeout creates a new BatchTypeGroupsUpdateParams object
// with the ability to set a timeout on a request.
func NewBatchTypeGroupsUpdateParamsWithTimeout(timeout time.Duration) *BatchTypeGroupsUpdateParams {
	return &BatchTypeGroupsUpdateParams{
		timeout: timeout,
	}
}

// NewBatchTypeGroupsUpdateParamsWithContext creates a new BatchTypeGroupsUpdateParams object
// with the ability to set a context for a request.
func NewBatchTypeGroupsUpdateParamsWithContext(ctx context.Context) *BatchTypeGroupsUpdateParams {
	return &BatchTypeGroupsUpdateParams{
		Context: ctx,
	}
}

// NewBatchTypeGroupsUpdateParamsWithHTTPClient creates a new BatchTypeGroupsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewBatchTypeGroupsUpdateParamsWithHTTPClient(client *http.Client) *BatchTypeGroupsUpdateParams {
	return &BatchTypeGroupsUpdateParams{
		HTTPClient: client,
	}
}

/*
BatchTypeGroupsUpdateParams contains all the parameters to send to the API endpoint

	for the batch type groups update operation.

	Typically these are written to a http.Request.
*/
type BatchTypeGroupsUpdateParams struct {

	// Data.
	Data *models.BatchTypeGroup

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the batch type groups update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BatchTypeGroupsUpdateParams) WithDefaults() *BatchTypeGroupsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the batch type groups update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BatchTypeGroupsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the batch type groups update params
func (o *BatchTypeGroupsUpdateParams) WithTimeout(timeout time.Duration) *BatchTypeGroupsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the batch type groups update params
func (o *BatchTypeGroupsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the batch type groups update params
func (o *BatchTypeGroupsUpdateParams) WithContext(ctx context.Context) *BatchTypeGroupsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the batch type groups update params
func (o *BatchTypeGroupsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the batch type groups update params
func (o *BatchTypeGroupsUpdateParams) WithHTTPClient(client *http.Client) *BatchTypeGroupsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the batch type groups update params
func (o *BatchTypeGroupsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the batch type groups update params
func (o *BatchTypeGroupsUpdateParams) WithData(data *models.BatchTypeGroup) *BatchTypeGroupsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the batch type groups update params
func (o *BatchTypeGroupsUpdateParams) SetData(data *models.BatchTypeGroup) {
	o.Data = data
}

// WithID adds the id to the batch type groups update params
func (o *BatchTypeGroupsUpdateParams) WithID(id string) *BatchTypeGroupsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the batch type groups update params
func (o *BatchTypeGroupsUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *BatchTypeGroupsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
