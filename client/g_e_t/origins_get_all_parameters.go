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

// NewOriginsGetAllParams creates a new OriginsGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOriginsGetAllParams() *OriginsGetAllParams {
	return &OriginsGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOriginsGetAllParamsWithTimeout creates a new OriginsGetAllParams object
// with the ability to set a timeout on a request.
func NewOriginsGetAllParamsWithTimeout(timeout time.Duration) *OriginsGetAllParams {
	return &OriginsGetAllParams{
		timeout: timeout,
	}
}

// NewOriginsGetAllParamsWithContext creates a new OriginsGetAllParams object
// with the ability to set a context for a request.
func NewOriginsGetAllParamsWithContext(ctx context.Context) *OriginsGetAllParams {
	return &OriginsGetAllParams{
		Context: ctx,
	}
}

// NewOriginsGetAllParamsWithHTTPClient creates a new OriginsGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewOriginsGetAllParamsWithHTTPClient(client *http.Client) *OriginsGetAllParams {
	return &OriginsGetAllParams{
		HTTPClient: client,
	}
}

/*
OriginsGetAllParams contains all the parameters to send to the API endpoint

	for the origins get all operation.

	Typically these are written to a http.Request.
*/
type OriginsGetAllParams struct {

	/* MaintenanceMode.

	   Ignore control grouping (default: false)
	*/
	MaintenanceMode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the origins get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OriginsGetAllParams) WithDefaults() *OriginsGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the origins get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OriginsGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the origins get all params
func (o *OriginsGetAllParams) WithTimeout(timeout time.Duration) *OriginsGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the origins get all params
func (o *OriginsGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the origins get all params
func (o *OriginsGetAllParams) WithContext(ctx context.Context) *OriginsGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the origins get all params
func (o *OriginsGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the origins get all params
func (o *OriginsGetAllParams) WithHTTPClient(client *http.Client) *OriginsGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the origins get all params
func (o *OriginsGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMaintenanceMode adds the maintenanceMode to the origins get all params
func (o *OriginsGetAllParams) WithMaintenanceMode(maintenanceMode *string) *OriginsGetAllParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the origins get all params
func (o *OriginsGetAllParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *OriginsGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
