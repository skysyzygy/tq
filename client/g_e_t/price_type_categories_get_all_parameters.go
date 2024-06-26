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

// NewPriceTypeCategoriesGetAllParams creates a new PriceTypeCategoriesGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPriceTypeCategoriesGetAllParams() *PriceTypeCategoriesGetAllParams {
	return &PriceTypeCategoriesGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPriceTypeCategoriesGetAllParamsWithTimeout creates a new PriceTypeCategoriesGetAllParams object
// with the ability to set a timeout on a request.
func NewPriceTypeCategoriesGetAllParamsWithTimeout(timeout time.Duration) *PriceTypeCategoriesGetAllParams {
	return &PriceTypeCategoriesGetAllParams{
		timeout: timeout,
	}
}

// NewPriceTypeCategoriesGetAllParamsWithContext creates a new PriceTypeCategoriesGetAllParams object
// with the ability to set a context for a request.
func NewPriceTypeCategoriesGetAllParamsWithContext(ctx context.Context) *PriceTypeCategoriesGetAllParams {
	return &PriceTypeCategoriesGetAllParams{
		Context: ctx,
	}
}

// NewPriceTypeCategoriesGetAllParamsWithHTTPClient creates a new PriceTypeCategoriesGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewPriceTypeCategoriesGetAllParamsWithHTTPClient(client *http.Client) *PriceTypeCategoriesGetAllParams {
	return &PriceTypeCategoriesGetAllParams{
		HTTPClient: client,
	}
}

/*
PriceTypeCategoriesGetAllParams contains all the parameters to send to the API endpoint

	for the price type categories get all operation.

	Typically these are written to a http.Request.
*/
type PriceTypeCategoriesGetAllParams struct {

	/* MaintenanceMode.

	   Ignore control grouping (default: false)
	*/
	MaintenanceMode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the price type categories get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PriceTypeCategoriesGetAllParams) WithDefaults() *PriceTypeCategoriesGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the price type categories get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PriceTypeCategoriesGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the price type categories get all params
func (o *PriceTypeCategoriesGetAllParams) WithTimeout(timeout time.Duration) *PriceTypeCategoriesGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the price type categories get all params
func (o *PriceTypeCategoriesGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the price type categories get all params
func (o *PriceTypeCategoriesGetAllParams) WithContext(ctx context.Context) *PriceTypeCategoriesGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the price type categories get all params
func (o *PriceTypeCategoriesGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the price type categories get all params
func (o *PriceTypeCategoriesGetAllParams) WithHTTPClient(client *http.Client) *PriceTypeCategoriesGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the price type categories get all params
func (o *PriceTypeCategoriesGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMaintenanceMode adds the maintenanceMode to the price type categories get all params
func (o *PriceTypeCategoriesGetAllParams) WithMaintenanceMode(maintenanceMode *string) *PriceTypeCategoriesGetAllParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the price type categories get all params
func (o *PriceTypeCategoriesGetAllParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *PriceTypeCategoriesGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
