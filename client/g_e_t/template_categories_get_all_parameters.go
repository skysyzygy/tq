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

// NewTemplateCategoriesGetAllParams creates a new TemplateCategoriesGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTemplateCategoriesGetAllParams() *TemplateCategoriesGetAllParams {
	return &TemplateCategoriesGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTemplateCategoriesGetAllParamsWithTimeout creates a new TemplateCategoriesGetAllParams object
// with the ability to set a timeout on a request.
func NewTemplateCategoriesGetAllParamsWithTimeout(timeout time.Duration) *TemplateCategoriesGetAllParams {
	return &TemplateCategoriesGetAllParams{
		timeout: timeout,
	}
}

// NewTemplateCategoriesGetAllParamsWithContext creates a new TemplateCategoriesGetAllParams object
// with the ability to set a context for a request.
func NewTemplateCategoriesGetAllParamsWithContext(ctx context.Context) *TemplateCategoriesGetAllParams {
	return &TemplateCategoriesGetAllParams{
		Context: ctx,
	}
}

// NewTemplateCategoriesGetAllParamsWithHTTPClient creates a new TemplateCategoriesGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewTemplateCategoriesGetAllParamsWithHTTPClient(client *http.Client) *TemplateCategoriesGetAllParams {
	return &TemplateCategoriesGetAllParams{
		HTTPClient: client,
	}
}

/*
TemplateCategoriesGetAllParams contains all the parameters to send to the API endpoint

	for the template categories get all operation.

	Typically these are written to a http.Request.
*/
type TemplateCategoriesGetAllParams struct {

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

// WithDefaults hydrates default values in the template categories get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TemplateCategoriesGetAllParams) WithDefaults() *TemplateCategoriesGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the template categories get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TemplateCategoriesGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the template categories get all params
func (o *TemplateCategoriesGetAllParams) WithTimeout(timeout time.Duration) *TemplateCategoriesGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the template categories get all params
func (o *TemplateCategoriesGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the template categories get all params
func (o *TemplateCategoriesGetAllParams) WithContext(ctx context.Context) *TemplateCategoriesGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the template categories get all params
func (o *TemplateCategoriesGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the template categories get all params
func (o *TemplateCategoriesGetAllParams) WithHTTPClient(client *http.Client) *TemplateCategoriesGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the template categories get all params
func (o *TemplateCategoriesGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the template categories get all params
func (o *TemplateCategoriesGetAllParams) WithFilter(filter *string) *TemplateCategoriesGetAllParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the template categories get all params
func (o *TemplateCategoriesGetAllParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithMaintenanceMode adds the maintenanceMode to the template categories get all params
func (o *TemplateCategoriesGetAllParams) WithMaintenanceMode(maintenanceMode *string) *TemplateCategoriesGetAllParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the template categories get all params
func (o *TemplateCategoriesGetAllParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *TemplateCategoriesGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
