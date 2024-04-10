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

// NewOriginalSourcesGetParams creates a new OriginalSourcesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOriginalSourcesGetParams() *OriginalSourcesGetParams {
	return &OriginalSourcesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOriginalSourcesGetParamsWithTimeout creates a new OriginalSourcesGetParams object
// with the ability to set a timeout on a request.
func NewOriginalSourcesGetParamsWithTimeout(timeout time.Duration) *OriginalSourcesGetParams {
	return &OriginalSourcesGetParams{
		timeout: timeout,
	}
}

// NewOriginalSourcesGetParamsWithContext creates a new OriginalSourcesGetParams object
// with the ability to set a context for a request.
func NewOriginalSourcesGetParamsWithContext(ctx context.Context) *OriginalSourcesGetParams {
	return &OriginalSourcesGetParams{
		Context: ctx,
	}
}

// NewOriginalSourcesGetParamsWithHTTPClient creates a new OriginalSourcesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewOriginalSourcesGetParamsWithHTTPClient(client *http.Client) *OriginalSourcesGetParams {
	return &OriginalSourcesGetParams{
		HTTPClient: client,
	}
}

/*
OriginalSourcesGetParams contains all the parameters to send to the API endpoint

	for the original sources get operation.

	Typically these are written to a http.Request.
*/
type OriginalSourcesGetParams struct {

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

// WithDefaults hydrates default values in the original sources get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OriginalSourcesGetParams) WithDefaults() *OriginalSourcesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the original sources get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OriginalSourcesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the original sources get params
func (o *OriginalSourcesGetParams) WithTimeout(timeout time.Duration) *OriginalSourcesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the original sources get params
func (o *OriginalSourcesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the original sources get params
func (o *OriginalSourcesGetParams) WithContext(ctx context.Context) *OriginalSourcesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the original sources get params
func (o *OriginalSourcesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the original sources get params
func (o *OriginalSourcesGetParams) WithHTTPClient(client *http.Client) *OriginalSourcesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the original sources get params
func (o *OriginalSourcesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the original sources get params
func (o *OriginalSourcesGetParams) WithID(id string) *OriginalSourcesGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the original sources get params
func (o *OriginalSourcesGetParams) SetID(id string) {
	o.ID = id
}

// WithMaintenanceMode adds the maintenanceMode to the original sources get params
func (o *OriginalSourcesGetParams) WithMaintenanceMode(maintenanceMode *string) *OriginalSourcesGetParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the original sources get params
func (o *OriginalSourcesGetParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *OriginalSourcesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
