// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

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

// NewBatchTypeGroupsGetParams creates a new BatchTypeGroupsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBatchTypeGroupsGetParams() *BatchTypeGroupsGetParams {
	return &BatchTypeGroupsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBatchTypeGroupsGetParamsWithTimeout creates a new BatchTypeGroupsGetParams object
// with the ability to set a timeout on a request.
func NewBatchTypeGroupsGetParamsWithTimeout(timeout time.Duration) *BatchTypeGroupsGetParams {
	return &BatchTypeGroupsGetParams{
		timeout: timeout,
	}
}

// NewBatchTypeGroupsGetParamsWithContext creates a new BatchTypeGroupsGetParams object
// with the ability to set a context for a request.
func NewBatchTypeGroupsGetParamsWithContext(ctx context.Context) *BatchTypeGroupsGetParams {
	return &BatchTypeGroupsGetParams{
		Context: ctx,
	}
}

// NewBatchTypeGroupsGetParamsWithHTTPClient creates a new BatchTypeGroupsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewBatchTypeGroupsGetParamsWithHTTPClient(client *http.Client) *BatchTypeGroupsGetParams {
	return &BatchTypeGroupsGetParams{
		HTTPClient: client,
	}
}

/*
BatchTypeGroupsGetParams contains all the parameters to send to the API endpoint

	for the batch type groups get operation.

	Typically these are written to a http.Request.
*/
type BatchTypeGroupsGetParams struct {

	/* Filter.

	   Filter by user access (default: readwrite)
	*/
	Filter *string

	// ID.
	ID string

	/* MaintenanceMode.

	   Ignore control grouping (default: false)
	*/
	MaintenanceMode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the batch type groups get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BatchTypeGroupsGetParams) WithDefaults() *BatchTypeGroupsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the batch type groups get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BatchTypeGroupsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the batch type groups get params
func (o *BatchTypeGroupsGetParams) WithTimeout(timeout time.Duration) *BatchTypeGroupsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the batch type groups get params
func (o *BatchTypeGroupsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the batch type groups get params
func (o *BatchTypeGroupsGetParams) WithContext(ctx context.Context) *BatchTypeGroupsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the batch type groups get params
func (o *BatchTypeGroupsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the batch type groups get params
func (o *BatchTypeGroupsGetParams) WithHTTPClient(client *http.Client) *BatchTypeGroupsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the batch type groups get params
func (o *BatchTypeGroupsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the batch type groups get params
func (o *BatchTypeGroupsGetParams) WithFilter(filter *string) *BatchTypeGroupsGetParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the batch type groups get params
func (o *BatchTypeGroupsGetParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the batch type groups get params
func (o *BatchTypeGroupsGetParams) WithID(id string) *BatchTypeGroupsGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the batch type groups get params
func (o *BatchTypeGroupsGetParams) SetID(id string) {
	o.ID = id
}

// WithMaintenanceMode adds the maintenanceMode to the batch type groups get params
func (o *BatchTypeGroupsGetParams) WithMaintenanceMode(maintenanceMode *string) *BatchTypeGroupsGetParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the batch type groups get params
func (o *BatchTypeGroupsGetParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *BatchTypeGroupsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.MaintenanceMode != nil {

		// query param maintenanceMode
		var qrMaintenanceMode string

		if o.MaintenanceMode != nil {
			qrMaintenanceMode = *o.MaintenanceMode
		}
		qMaintenanceMode := qrMaintenanceMode
		if qMaintenanceMode != "" {

			if err := r.SetQueryParam("maintenanceMode", qMaintenanceMode); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
