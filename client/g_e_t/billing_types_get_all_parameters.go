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

// NewBillingTypesGetAllParams creates a new BillingTypesGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBillingTypesGetAllParams() *BillingTypesGetAllParams {
	return &BillingTypesGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBillingTypesGetAllParamsWithTimeout creates a new BillingTypesGetAllParams object
// with the ability to set a timeout on a request.
func NewBillingTypesGetAllParamsWithTimeout(timeout time.Duration) *BillingTypesGetAllParams {
	return &BillingTypesGetAllParams{
		timeout: timeout,
	}
}

// NewBillingTypesGetAllParamsWithContext creates a new BillingTypesGetAllParams object
// with the ability to set a context for a request.
func NewBillingTypesGetAllParamsWithContext(ctx context.Context) *BillingTypesGetAllParams {
	return &BillingTypesGetAllParams{
		Context: ctx,
	}
}

// NewBillingTypesGetAllParamsWithHTTPClient creates a new BillingTypesGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewBillingTypesGetAllParamsWithHTTPClient(client *http.Client) *BillingTypesGetAllParams {
	return &BillingTypesGetAllParams{
		HTTPClient: client,
	}
}

/*
BillingTypesGetAllParams contains all the parameters to send to the API endpoint

	for the billing types get all operation.

	Typically these are written to a http.Request.
*/
type BillingTypesGetAllParams struct {

	/* MaintenanceMode.

	   Ignore control grouping (default: false)
	*/
	MaintenanceMode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the billing types get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BillingTypesGetAllParams) WithDefaults() *BillingTypesGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the billing types get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BillingTypesGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the billing types get all params
func (o *BillingTypesGetAllParams) WithTimeout(timeout time.Duration) *BillingTypesGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the billing types get all params
func (o *BillingTypesGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the billing types get all params
func (o *BillingTypesGetAllParams) WithContext(ctx context.Context) *BillingTypesGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the billing types get all params
func (o *BillingTypesGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the billing types get all params
func (o *BillingTypesGetAllParams) WithHTTPClient(client *http.Client) *BillingTypesGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the billing types get all params
func (o *BillingTypesGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMaintenanceMode adds the maintenanceMode to the billing types get all params
func (o *BillingTypesGetAllParams) WithMaintenanceMode(maintenanceMode *string) *BillingTypesGetAllParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the billing types get all params
func (o *BillingTypesGetAllParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *BillingTypesGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
