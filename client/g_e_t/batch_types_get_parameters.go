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

// NewBatchTypesGetParams creates a new BatchTypesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBatchTypesGetParams() *BatchTypesGetParams {
	return &BatchTypesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBatchTypesGetParamsWithTimeout creates a new BatchTypesGetParams object
// with the ability to set a timeout on a request.
func NewBatchTypesGetParamsWithTimeout(timeout time.Duration) *BatchTypesGetParams {
	return &BatchTypesGetParams{
		timeout: timeout,
	}
}

// NewBatchTypesGetParamsWithContext creates a new BatchTypesGetParams object
// with the ability to set a context for a request.
func NewBatchTypesGetParamsWithContext(ctx context.Context) *BatchTypesGetParams {
	return &BatchTypesGetParams{
		Context: ctx,
	}
}

// NewBatchTypesGetParamsWithHTTPClient creates a new BatchTypesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewBatchTypesGetParamsWithHTTPClient(client *http.Client) *BatchTypesGetParams {
	return &BatchTypesGetParams{
		HTTPClient: client,
	}
}

/*
BatchTypesGetParams contains all the parameters to send to the API endpoint

	for the batch types get operation.

	Typically these are written to a http.Request.
*/
type BatchTypesGetParams struct {

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

// WithDefaults hydrates default values in the batch types get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BatchTypesGetParams) WithDefaults() *BatchTypesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the batch types get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BatchTypesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the batch types get params
func (o *BatchTypesGetParams) WithTimeout(timeout time.Duration) *BatchTypesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the batch types get params
func (o *BatchTypesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the batch types get params
func (o *BatchTypesGetParams) WithContext(ctx context.Context) *BatchTypesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the batch types get params
func (o *BatchTypesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the batch types get params
func (o *BatchTypesGetParams) WithHTTPClient(client *http.Client) *BatchTypesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the batch types get params
func (o *BatchTypesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the batch types get params
func (o *BatchTypesGetParams) WithID(id string) *BatchTypesGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the batch types get params
func (o *BatchTypesGetParams) SetID(id string) {
	o.ID = id
}

// WithMaintenanceMode adds the maintenanceMode to the batch types get params
func (o *BatchTypesGetParams) WithMaintenanceMode(maintenanceMode *string) *BatchTypesGetParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the batch types get params
func (o *BatchTypesGetParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *BatchTypesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
