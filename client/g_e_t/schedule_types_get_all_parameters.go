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

// NewScheduleTypesGetAllParams creates a new ScheduleTypesGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewScheduleTypesGetAllParams() *ScheduleTypesGetAllParams {
	return &ScheduleTypesGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewScheduleTypesGetAllParamsWithTimeout creates a new ScheduleTypesGetAllParams object
// with the ability to set a timeout on a request.
func NewScheduleTypesGetAllParamsWithTimeout(timeout time.Duration) *ScheduleTypesGetAllParams {
	return &ScheduleTypesGetAllParams{
		timeout: timeout,
	}
}

// NewScheduleTypesGetAllParamsWithContext creates a new ScheduleTypesGetAllParams object
// with the ability to set a context for a request.
func NewScheduleTypesGetAllParamsWithContext(ctx context.Context) *ScheduleTypesGetAllParams {
	return &ScheduleTypesGetAllParams{
		Context: ctx,
	}
}

// NewScheduleTypesGetAllParamsWithHTTPClient creates a new ScheduleTypesGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewScheduleTypesGetAllParamsWithHTTPClient(client *http.Client) *ScheduleTypesGetAllParams {
	return &ScheduleTypesGetAllParams{
		HTTPClient: client,
	}
}

/*
ScheduleTypesGetAllParams contains all the parameters to send to the API endpoint

	for the schedule types get all operation.

	Typically these are written to a http.Request.
*/
type ScheduleTypesGetAllParams struct {

	/* MaintenanceMode.

	   Ignore control grouping (default: false)
	*/
	MaintenanceMode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the schedule types get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScheduleTypesGetAllParams) WithDefaults() *ScheduleTypesGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the schedule types get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScheduleTypesGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the schedule types get all params
func (o *ScheduleTypesGetAllParams) WithTimeout(timeout time.Duration) *ScheduleTypesGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schedule types get all params
func (o *ScheduleTypesGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schedule types get all params
func (o *ScheduleTypesGetAllParams) WithContext(ctx context.Context) *ScheduleTypesGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schedule types get all params
func (o *ScheduleTypesGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schedule types get all params
func (o *ScheduleTypesGetAllParams) WithHTTPClient(client *http.Client) *ScheduleTypesGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schedule types get all params
func (o *ScheduleTypesGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMaintenanceMode adds the maintenanceMode to the schedule types get all params
func (o *ScheduleTypesGetAllParams) WithMaintenanceMode(maintenanceMode *string) *ScheduleTypesGetAllParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the schedule types get all params
func (o *ScheduleTypesGetAllParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *ScheduleTypesGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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