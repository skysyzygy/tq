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

// NewContactTypesGetParams creates a new ContactTypesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContactTypesGetParams() *ContactTypesGetParams {
	return &ContactTypesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContactTypesGetParamsWithTimeout creates a new ContactTypesGetParams object
// with the ability to set a timeout on a request.
func NewContactTypesGetParamsWithTimeout(timeout time.Duration) *ContactTypesGetParams {
	return &ContactTypesGetParams{
		timeout: timeout,
	}
}

// NewContactTypesGetParamsWithContext creates a new ContactTypesGetParams object
// with the ability to set a context for a request.
func NewContactTypesGetParamsWithContext(ctx context.Context) *ContactTypesGetParams {
	return &ContactTypesGetParams{
		Context: ctx,
	}
}

// NewContactTypesGetParamsWithHTTPClient creates a new ContactTypesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewContactTypesGetParamsWithHTTPClient(client *http.Client) *ContactTypesGetParams {
	return &ContactTypesGetParams{
		HTTPClient: client,
	}
}

/*
ContactTypesGetParams contains all the parameters to send to the API endpoint

	for the contact types get operation.

	Typically these are written to a http.Request.
*/
type ContactTypesGetParams struct {

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

// WithDefaults hydrates default values in the contact types get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactTypesGetParams) WithDefaults() *ContactTypesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the contact types get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactTypesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the contact types get params
func (o *ContactTypesGetParams) WithTimeout(timeout time.Duration) *ContactTypesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contact types get params
func (o *ContactTypesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contact types get params
func (o *ContactTypesGetParams) WithContext(ctx context.Context) *ContactTypesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contact types get params
func (o *ContactTypesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contact types get params
func (o *ContactTypesGetParams) WithHTTPClient(client *http.Client) *ContactTypesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contact types get params
func (o *ContactTypesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the contact types get params
func (o *ContactTypesGetParams) WithID(id string) *ContactTypesGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the contact types get params
func (o *ContactTypesGetParams) SetID(id string) {
	o.ID = id
}

// WithMaintenanceMode adds the maintenanceMode to the contact types get params
func (o *ContactTypesGetParams) WithMaintenanceMode(maintenanceMode *string) *ContactTypesGetParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the contact types get params
func (o *ContactTypesGetParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *ContactTypesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
