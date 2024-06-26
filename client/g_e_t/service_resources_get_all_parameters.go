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

// NewServiceResourcesGetAllParams creates a new ServiceResourcesGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServiceResourcesGetAllParams() *ServiceResourcesGetAllParams {
	return &ServiceResourcesGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServiceResourcesGetAllParamsWithTimeout creates a new ServiceResourcesGetAllParams object
// with the ability to set a timeout on a request.
func NewServiceResourcesGetAllParamsWithTimeout(timeout time.Duration) *ServiceResourcesGetAllParams {
	return &ServiceResourcesGetAllParams{
		timeout: timeout,
	}
}

// NewServiceResourcesGetAllParamsWithContext creates a new ServiceResourcesGetAllParams object
// with the ability to set a context for a request.
func NewServiceResourcesGetAllParamsWithContext(ctx context.Context) *ServiceResourcesGetAllParams {
	return &ServiceResourcesGetAllParams{
		Context: ctx,
	}
}

// NewServiceResourcesGetAllParamsWithHTTPClient creates a new ServiceResourcesGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewServiceResourcesGetAllParamsWithHTTPClient(client *http.Client) *ServiceResourcesGetAllParams {
	return &ServiceResourcesGetAllParams{
		HTTPClient: client,
	}
}

/*
ServiceResourcesGetAllParams contains all the parameters to send to the API endpoint

	for the service resources get all operation.

	Typically these are written to a http.Request.
*/
type ServiceResourcesGetAllParams struct {

	/* MaintenanceMode.

	   Ignore control grouping (default: false)
	*/
	MaintenanceMode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service resources get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceResourcesGetAllParams) WithDefaults() *ServiceResourcesGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service resources get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceResourcesGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service resources get all params
func (o *ServiceResourcesGetAllParams) WithTimeout(timeout time.Duration) *ServiceResourcesGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service resources get all params
func (o *ServiceResourcesGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service resources get all params
func (o *ServiceResourcesGetAllParams) WithContext(ctx context.Context) *ServiceResourcesGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service resources get all params
func (o *ServiceResourcesGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service resources get all params
func (o *ServiceResourcesGetAllParams) WithHTTPClient(client *http.Client) *ServiceResourcesGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service resources get all params
func (o *ServiceResourcesGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMaintenanceMode adds the maintenanceMode to the service resources get all params
func (o *ServiceResourcesGetAllParams) WithMaintenanceMode(maintenanceMode *string) *ServiceResourcesGetAllParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the service resources get all params
func (o *ServiceResourcesGetAllParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceResourcesGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
