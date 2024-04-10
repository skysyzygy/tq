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

// NewBatchTypesUpdateParams creates a new BatchTypesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBatchTypesUpdateParams() *BatchTypesUpdateParams {
	return &BatchTypesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBatchTypesUpdateParamsWithTimeout creates a new BatchTypesUpdateParams object
// with the ability to set a timeout on a request.
func NewBatchTypesUpdateParamsWithTimeout(timeout time.Duration) *BatchTypesUpdateParams {
	return &BatchTypesUpdateParams{
		timeout: timeout,
	}
}

// NewBatchTypesUpdateParamsWithContext creates a new BatchTypesUpdateParams object
// with the ability to set a context for a request.
func NewBatchTypesUpdateParamsWithContext(ctx context.Context) *BatchTypesUpdateParams {
	return &BatchTypesUpdateParams{
		Context: ctx,
	}
}

// NewBatchTypesUpdateParamsWithHTTPClient creates a new BatchTypesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewBatchTypesUpdateParamsWithHTTPClient(client *http.Client) *BatchTypesUpdateParams {
	return &BatchTypesUpdateParams{
		HTTPClient: client,
	}
}

/*
BatchTypesUpdateParams contains all the parameters to send to the API endpoint

	for the batch types update operation.

	Typically these are written to a http.Request.
*/
type BatchTypesUpdateParams struct {

	// Data.
	Data *models.BatchType

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the batch types update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BatchTypesUpdateParams) WithDefaults() *BatchTypesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the batch types update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BatchTypesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the batch types update params
func (o *BatchTypesUpdateParams) WithTimeout(timeout time.Duration) *BatchTypesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the batch types update params
func (o *BatchTypesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the batch types update params
func (o *BatchTypesUpdateParams) WithContext(ctx context.Context) *BatchTypesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the batch types update params
func (o *BatchTypesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the batch types update params
func (o *BatchTypesUpdateParams) WithHTTPClient(client *http.Client) *BatchTypesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the batch types update params
func (o *BatchTypesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the batch types update params
func (o *BatchTypesUpdateParams) WithData(data *models.BatchType) *BatchTypesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the batch types update params
func (o *BatchTypesUpdateParams) SetData(data *models.BatchType) {
	o.Data = data
}

// WithID adds the id to the batch types update params
func (o *BatchTypesUpdateParams) WithID(id string) *BatchTypesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the batch types update params
func (o *BatchTypesUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *BatchTypesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
