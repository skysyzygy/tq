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

// NewAssociationTypesGetParams creates a new AssociationTypesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssociationTypesGetParams() *AssociationTypesGetParams {
	return &AssociationTypesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAssociationTypesGetParamsWithTimeout creates a new AssociationTypesGetParams object
// with the ability to set a timeout on a request.
func NewAssociationTypesGetParamsWithTimeout(timeout time.Duration) *AssociationTypesGetParams {
	return &AssociationTypesGetParams{
		timeout: timeout,
	}
}

// NewAssociationTypesGetParamsWithContext creates a new AssociationTypesGetParams object
// with the ability to set a context for a request.
func NewAssociationTypesGetParamsWithContext(ctx context.Context) *AssociationTypesGetParams {
	return &AssociationTypesGetParams{
		Context: ctx,
	}
}

// NewAssociationTypesGetParamsWithHTTPClient creates a new AssociationTypesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssociationTypesGetParamsWithHTTPClient(client *http.Client) *AssociationTypesGetParams {
	return &AssociationTypesGetParams{
		HTTPClient: client,
	}
}

/*
AssociationTypesGetParams contains all the parameters to send to the API endpoint

	for the association types get operation.

	Typically these are written to a http.Request.
*/
type AssociationTypesGetParams struct {

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

// WithDefaults hydrates default values in the association types get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssociationTypesGetParams) WithDefaults() *AssociationTypesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the association types get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssociationTypesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the association types get params
func (o *AssociationTypesGetParams) WithTimeout(timeout time.Duration) *AssociationTypesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the association types get params
func (o *AssociationTypesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the association types get params
func (o *AssociationTypesGetParams) WithContext(ctx context.Context) *AssociationTypesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the association types get params
func (o *AssociationTypesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the association types get params
func (o *AssociationTypesGetParams) WithHTTPClient(client *http.Client) *AssociationTypesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the association types get params
func (o *AssociationTypesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the association types get params
func (o *AssociationTypesGetParams) WithFilter(filter *string) *AssociationTypesGetParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the association types get params
func (o *AssociationTypesGetParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the association types get params
func (o *AssociationTypesGetParams) WithID(id string) *AssociationTypesGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the association types get params
func (o *AssociationTypesGetParams) SetID(id string) {
	o.ID = id
}

// WithMaintenanceMode adds the maintenanceMode to the association types get params
func (o *AssociationTypesGetParams) WithMaintenanceMode(maintenanceMode *string) *AssociationTypesGetParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the association types get params
func (o *AssociationTypesGetParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *AssociationTypesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
