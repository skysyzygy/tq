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

// NewContactPermissionTypesGetParams creates a new ContactPermissionTypesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContactPermissionTypesGetParams() *ContactPermissionTypesGetParams {
	return &ContactPermissionTypesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContactPermissionTypesGetParamsWithTimeout creates a new ContactPermissionTypesGetParams object
// with the ability to set a timeout on a request.
func NewContactPermissionTypesGetParamsWithTimeout(timeout time.Duration) *ContactPermissionTypesGetParams {
	return &ContactPermissionTypesGetParams{
		timeout: timeout,
	}
}

// NewContactPermissionTypesGetParamsWithContext creates a new ContactPermissionTypesGetParams object
// with the ability to set a context for a request.
func NewContactPermissionTypesGetParamsWithContext(ctx context.Context) *ContactPermissionTypesGetParams {
	return &ContactPermissionTypesGetParams{
		Context: ctx,
	}
}

// NewContactPermissionTypesGetParamsWithHTTPClient creates a new ContactPermissionTypesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewContactPermissionTypesGetParamsWithHTTPClient(client *http.Client) *ContactPermissionTypesGetParams {
	return &ContactPermissionTypesGetParams{
		HTTPClient: client,
	}
}

/*
ContactPermissionTypesGetParams contains all the parameters to send to the API endpoint

	for the contact permission types get operation.

	Typically these are written to a http.Request.
*/
type ContactPermissionTypesGetParams struct {

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

// WithDefaults hydrates default values in the contact permission types get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactPermissionTypesGetParams) WithDefaults() *ContactPermissionTypesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the contact permission types get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactPermissionTypesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the contact permission types get params
func (o *ContactPermissionTypesGetParams) WithTimeout(timeout time.Duration) *ContactPermissionTypesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contact permission types get params
func (o *ContactPermissionTypesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contact permission types get params
func (o *ContactPermissionTypesGetParams) WithContext(ctx context.Context) *ContactPermissionTypesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contact permission types get params
func (o *ContactPermissionTypesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contact permission types get params
func (o *ContactPermissionTypesGetParams) WithHTTPClient(client *http.Client) *ContactPermissionTypesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contact permission types get params
func (o *ContactPermissionTypesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the contact permission types get params
func (o *ContactPermissionTypesGetParams) WithFilter(filter *string) *ContactPermissionTypesGetParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the contact permission types get params
func (o *ContactPermissionTypesGetParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the contact permission types get params
func (o *ContactPermissionTypesGetParams) WithID(id string) *ContactPermissionTypesGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the contact permission types get params
func (o *ContactPermissionTypesGetParams) SetID(id string) {
	o.ID = id
}

// WithMaintenanceMode adds the maintenanceMode to the contact permission types get params
func (o *ContactPermissionTypesGetParams) WithMaintenanceMode(maintenanceMode *string) *ContactPermissionTypesGetParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the contact permission types get params
func (o *ContactPermissionTypesGetParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *ContactPermissionTypesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
