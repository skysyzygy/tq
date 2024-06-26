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

// NewSpecialActivityTypesGetAllParams creates a new SpecialActivityTypesGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSpecialActivityTypesGetAllParams() *SpecialActivityTypesGetAllParams {
	return &SpecialActivityTypesGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSpecialActivityTypesGetAllParamsWithTimeout creates a new SpecialActivityTypesGetAllParams object
// with the ability to set a timeout on a request.
func NewSpecialActivityTypesGetAllParamsWithTimeout(timeout time.Duration) *SpecialActivityTypesGetAllParams {
	return &SpecialActivityTypesGetAllParams{
		timeout: timeout,
	}
}

// NewSpecialActivityTypesGetAllParamsWithContext creates a new SpecialActivityTypesGetAllParams object
// with the ability to set a context for a request.
func NewSpecialActivityTypesGetAllParamsWithContext(ctx context.Context) *SpecialActivityTypesGetAllParams {
	return &SpecialActivityTypesGetAllParams{
		Context: ctx,
	}
}

// NewSpecialActivityTypesGetAllParamsWithHTTPClient creates a new SpecialActivityTypesGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewSpecialActivityTypesGetAllParamsWithHTTPClient(client *http.Client) *SpecialActivityTypesGetAllParams {
	return &SpecialActivityTypesGetAllParams{
		HTTPClient: client,
	}
}

/*
SpecialActivityTypesGetAllParams contains all the parameters to send to the API endpoint

	for the special activity types get all operation.

	Typically these are written to a http.Request.
*/
type SpecialActivityTypesGetAllParams struct {

	/* Filter.

	   Filter by user access (default: readwrite)
	*/
	Filter *string

	/* MaintenanceMode.

	   Ignore control grouping (default: false)
	*/
	MaintenanceMode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the special activity types get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SpecialActivityTypesGetAllParams) WithDefaults() *SpecialActivityTypesGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the special activity types get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SpecialActivityTypesGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the special activity types get all params
func (o *SpecialActivityTypesGetAllParams) WithTimeout(timeout time.Duration) *SpecialActivityTypesGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the special activity types get all params
func (o *SpecialActivityTypesGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the special activity types get all params
func (o *SpecialActivityTypesGetAllParams) WithContext(ctx context.Context) *SpecialActivityTypesGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the special activity types get all params
func (o *SpecialActivityTypesGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the special activity types get all params
func (o *SpecialActivityTypesGetAllParams) WithHTTPClient(client *http.Client) *SpecialActivityTypesGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the special activity types get all params
func (o *SpecialActivityTypesGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the special activity types get all params
func (o *SpecialActivityTypesGetAllParams) WithFilter(filter *string) *SpecialActivityTypesGetAllParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the special activity types get all params
func (o *SpecialActivityTypesGetAllParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithMaintenanceMode adds the maintenanceMode to the special activity types get all params
func (o *SpecialActivityTypesGetAllParams) WithMaintenanceMode(maintenanceMode *string) *SpecialActivityTypesGetAllParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the special activity types get all params
func (o *SpecialActivityTypesGetAllParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *SpecialActivityTypesGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
