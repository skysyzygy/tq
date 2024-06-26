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

// NewBusinessUnitsGetAllParams creates a new BusinessUnitsGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBusinessUnitsGetAllParams() *BusinessUnitsGetAllParams {
	return &BusinessUnitsGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBusinessUnitsGetAllParamsWithTimeout creates a new BusinessUnitsGetAllParams object
// with the ability to set a timeout on a request.
func NewBusinessUnitsGetAllParamsWithTimeout(timeout time.Duration) *BusinessUnitsGetAllParams {
	return &BusinessUnitsGetAllParams{
		timeout: timeout,
	}
}

// NewBusinessUnitsGetAllParamsWithContext creates a new BusinessUnitsGetAllParams object
// with the ability to set a context for a request.
func NewBusinessUnitsGetAllParamsWithContext(ctx context.Context) *BusinessUnitsGetAllParams {
	return &BusinessUnitsGetAllParams{
		Context: ctx,
	}
}

// NewBusinessUnitsGetAllParamsWithHTTPClient creates a new BusinessUnitsGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewBusinessUnitsGetAllParamsWithHTTPClient(client *http.Client) *BusinessUnitsGetAllParams {
	return &BusinessUnitsGetAllParams{
		HTTPClient: client,
	}
}

/*
BusinessUnitsGetAllParams contains all the parameters to send to the API endpoint

	for the business units get all operation.

	Typically these are written to a http.Request.
*/
type BusinessUnitsGetAllParams struct {

	/* MaintenanceMode.

	   Ignore control grouping (default: false)
	*/
	MaintenanceMode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the business units get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BusinessUnitsGetAllParams) WithDefaults() *BusinessUnitsGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the business units get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BusinessUnitsGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the business units get all params
func (o *BusinessUnitsGetAllParams) WithTimeout(timeout time.Duration) *BusinessUnitsGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the business units get all params
func (o *BusinessUnitsGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the business units get all params
func (o *BusinessUnitsGetAllParams) WithContext(ctx context.Context) *BusinessUnitsGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the business units get all params
func (o *BusinessUnitsGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the business units get all params
func (o *BusinessUnitsGetAllParams) WithHTTPClient(client *http.Client) *BusinessUnitsGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the business units get all params
func (o *BusinessUnitsGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMaintenanceMode adds the maintenanceMode to the business units get all params
func (o *BusinessUnitsGetAllParams) WithMaintenanceMode(maintenanceMode *string) *BusinessUnitsGetAllParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the business units get all params
func (o *BusinessUnitsGetAllParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *BusinessUnitsGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
