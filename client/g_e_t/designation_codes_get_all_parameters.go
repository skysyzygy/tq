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

// NewDesignationCodesGetAllParams creates a new DesignationCodesGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDesignationCodesGetAllParams() *DesignationCodesGetAllParams {
	return &DesignationCodesGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDesignationCodesGetAllParamsWithTimeout creates a new DesignationCodesGetAllParams object
// with the ability to set a timeout on a request.
func NewDesignationCodesGetAllParamsWithTimeout(timeout time.Duration) *DesignationCodesGetAllParams {
	return &DesignationCodesGetAllParams{
		timeout: timeout,
	}
}

// NewDesignationCodesGetAllParamsWithContext creates a new DesignationCodesGetAllParams object
// with the ability to set a context for a request.
func NewDesignationCodesGetAllParamsWithContext(ctx context.Context) *DesignationCodesGetAllParams {
	return &DesignationCodesGetAllParams{
		Context: ctx,
	}
}

// NewDesignationCodesGetAllParamsWithHTTPClient creates a new DesignationCodesGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewDesignationCodesGetAllParamsWithHTTPClient(client *http.Client) *DesignationCodesGetAllParams {
	return &DesignationCodesGetAllParams{
		HTTPClient: client,
	}
}

/*
DesignationCodesGetAllParams contains all the parameters to send to the API endpoint

	for the designation codes get all operation.

	Typically these are written to a http.Request.
*/
type DesignationCodesGetAllParams struct {

	/* MaintenanceMode.

	   Ignore control grouping (default: false)
	*/
	MaintenanceMode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the designation codes get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DesignationCodesGetAllParams) WithDefaults() *DesignationCodesGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the designation codes get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DesignationCodesGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the designation codes get all params
func (o *DesignationCodesGetAllParams) WithTimeout(timeout time.Duration) *DesignationCodesGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the designation codes get all params
func (o *DesignationCodesGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the designation codes get all params
func (o *DesignationCodesGetAllParams) WithContext(ctx context.Context) *DesignationCodesGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the designation codes get all params
func (o *DesignationCodesGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the designation codes get all params
func (o *DesignationCodesGetAllParams) WithHTTPClient(client *http.Client) *DesignationCodesGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the designation codes get all params
func (o *DesignationCodesGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMaintenanceMode adds the maintenanceMode to the designation codes get all params
func (o *DesignationCodesGetAllParams) WithMaintenanceMode(maintenanceMode *string) *DesignationCodesGetAllParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the designation codes get all params
func (o *DesignationCodesGetAllParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *DesignationCodesGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
