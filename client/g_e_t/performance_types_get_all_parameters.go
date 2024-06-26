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

// NewPerformanceTypesGetAllParams creates a new PerformanceTypesGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformanceTypesGetAllParams() *PerformanceTypesGetAllParams {
	return &PerformanceTypesGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformanceTypesGetAllParamsWithTimeout creates a new PerformanceTypesGetAllParams object
// with the ability to set a timeout on a request.
func NewPerformanceTypesGetAllParamsWithTimeout(timeout time.Duration) *PerformanceTypesGetAllParams {
	return &PerformanceTypesGetAllParams{
		timeout: timeout,
	}
}

// NewPerformanceTypesGetAllParamsWithContext creates a new PerformanceTypesGetAllParams object
// with the ability to set a context for a request.
func NewPerformanceTypesGetAllParamsWithContext(ctx context.Context) *PerformanceTypesGetAllParams {
	return &PerformanceTypesGetAllParams{
		Context: ctx,
	}
}

// NewPerformanceTypesGetAllParamsWithHTTPClient creates a new PerformanceTypesGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformanceTypesGetAllParamsWithHTTPClient(client *http.Client) *PerformanceTypesGetAllParams {
	return &PerformanceTypesGetAllParams{
		HTTPClient: client,
	}
}

/*
PerformanceTypesGetAllParams contains all the parameters to send to the API endpoint

	for the performance types get all operation.

	Typically these are written to a http.Request.
*/
type PerformanceTypesGetAllParams struct {

	/* MaintenanceMode.

	   Ignore control grouping (default: false)
	*/
	MaintenanceMode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the performance types get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformanceTypesGetAllParams) WithDefaults() *PerformanceTypesGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the performance types get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformanceTypesGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the performance types get all params
func (o *PerformanceTypesGetAllParams) WithTimeout(timeout time.Duration) *PerformanceTypesGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the performance types get all params
func (o *PerformanceTypesGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the performance types get all params
func (o *PerformanceTypesGetAllParams) WithContext(ctx context.Context) *PerformanceTypesGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the performance types get all params
func (o *PerformanceTypesGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the performance types get all params
func (o *PerformanceTypesGetAllParams) WithHTTPClient(client *http.Client) *PerformanceTypesGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the performance types get all params
func (o *PerformanceTypesGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMaintenanceMode adds the maintenanceMode to the performance types get all params
func (o *PerformanceTypesGetAllParams) WithMaintenanceMode(maintenanceMode *string) *PerformanceTypesGetAllParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the performance types get all params
func (o *PerformanceTypesGetAllParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *PerformanceTypesGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
