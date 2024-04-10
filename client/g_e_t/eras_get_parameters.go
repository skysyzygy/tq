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

// NewErasGetParams creates a new ErasGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewErasGetParams() *ErasGetParams {
	return &ErasGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewErasGetParamsWithTimeout creates a new ErasGetParams object
// with the ability to set a timeout on a request.
func NewErasGetParamsWithTimeout(timeout time.Duration) *ErasGetParams {
	return &ErasGetParams{
		timeout: timeout,
	}
}

// NewErasGetParamsWithContext creates a new ErasGetParams object
// with the ability to set a context for a request.
func NewErasGetParamsWithContext(ctx context.Context) *ErasGetParams {
	return &ErasGetParams{
		Context: ctx,
	}
}

// NewErasGetParamsWithHTTPClient creates a new ErasGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewErasGetParamsWithHTTPClient(client *http.Client) *ErasGetParams {
	return &ErasGetParams{
		HTTPClient: client,
	}
}

/*
ErasGetParams contains all the parameters to send to the API endpoint

	for the eras get operation.

	Typically these are written to a http.Request.
*/
type ErasGetParams struct {

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

// WithDefaults hydrates default values in the eras get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ErasGetParams) WithDefaults() *ErasGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the eras get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ErasGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the eras get params
func (o *ErasGetParams) WithTimeout(timeout time.Duration) *ErasGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the eras get params
func (o *ErasGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the eras get params
func (o *ErasGetParams) WithContext(ctx context.Context) *ErasGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the eras get params
func (o *ErasGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the eras get params
func (o *ErasGetParams) WithHTTPClient(client *http.Client) *ErasGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the eras get params
func (o *ErasGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the eras get params
func (o *ErasGetParams) WithID(id string) *ErasGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the eras get params
func (o *ErasGetParams) SetID(id string) {
	o.ID = id
}

// WithMaintenanceMode adds the maintenanceMode to the eras get params
func (o *ErasGetParams) WithMaintenanceMode(maintenanceMode *string) *ErasGetParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the eras get params
func (o *ErasGetParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *ErasGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
