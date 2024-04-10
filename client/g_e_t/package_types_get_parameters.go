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

// NewPackageTypesGetParams creates a new PackageTypesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPackageTypesGetParams() *PackageTypesGetParams {
	return &PackageTypesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPackageTypesGetParamsWithTimeout creates a new PackageTypesGetParams object
// with the ability to set a timeout on a request.
func NewPackageTypesGetParamsWithTimeout(timeout time.Duration) *PackageTypesGetParams {
	return &PackageTypesGetParams{
		timeout: timeout,
	}
}

// NewPackageTypesGetParamsWithContext creates a new PackageTypesGetParams object
// with the ability to set a context for a request.
func NewPackageTypesGetParamsWithContext(ctx context.Context) *PackageTypesGetParams {
	return &PackageTypesGetParams{
		Context: ctx,
	}
}

// NewPackageTypesGetParamsWithHTTPClient creates a new PackageTypesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPackageTypesGetParamsWithHTTPClient(client *http.Client) *PackageTypesGetParams {
	return &PackageTypesGetParams{
		HTTPClient: client,
	}
}

/*
PackageTypesGetParams contains all the parameters to send to the API endpoint

	for the package types get operation.

	Typically these are written to a http.Request.
*/
type PackageTypesGetParams struct {

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

// WithDefaults hydrates default values in the package types get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackageTypesGetParams) WithDefaults() *PackageTypesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the package types get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackageTypesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the package types get params
func (o *PackageTypesGetParams) WithTimeout(timeout time.Duration) *PackageTypesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the package types get params
func (o *PackageTypesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the package types get params
func (o *PackageTypesGetParams) WithContext(ctx context.Context) *PackageTypesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the package types get params
func (o *PackageTypesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the package types get params
func (o *PackageTypesGetParams) WithHTTPClient(client *http.Client) *PackageTypesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the package types get params
func (o *PackageTypesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the package types get params
func (o *PackageTypesGetParams) WithID(id string) *PackageTypesGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the package types get params
func (o *PackageTypesGetParams) SetID(id string) {
	o.ID = id
}

// WithMaintenanceMode adds the maintenanceMode to the package types get params
func (o *PackageTypesGetParams) WithMaintenanceMode(maintenanceMode *string) *PackageTypesGetParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the package types get params
func (o *PackageTypesGetParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *PackageTypesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
