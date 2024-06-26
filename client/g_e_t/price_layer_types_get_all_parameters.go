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

// NewPriceLayerTypesGetAllParams creates a new PriceLayerTypesGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPriceLayerTypesGetAllParams() *PriceLayerTypesGetAllParams {
	return &PriceLayerTypesGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPriceLayerTypesGetAllParamsWithTimeout creates a new PriceLayerTypesGetAllParams object
// with the ability to set a timeout on a request.
func NewPriceLayerTypesGetAllParamsWithTimeout(timeout time.Duration) *PriceLayerTypesGetAllParams {
	return &PriceLayerTypesGetAllParams{
		timeout: timeout,
	}
}

// NewPriceLayerTypesGetAllParamsWithContext creates a new PriceLayerTypesGetAllParams object
// with the ability to set a context for a request.
func NewPriceLayerTypesGetAllParamsWithContext(ctx context.Context) *PriceLayerTypesGetAllParams {
	return &PriceLayerTypesGetAllParams{
		Context: ctx,
	}
}

// NewPriceLayerTypesGetAllParamsWithHTTPClient creates a new PriceLayerTypesGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewPriceLayerTypesGetAllParamsWithHTTPClient(client *http.Client) *PriceLayerTypesGetAllParams {
	return &PriceLayerTypesGetAllParams{
		HTTPClient: client,
	}
}

/*
PriceLayerTypesGetAllParams contains all the parameters to send to the API endpoint

	for the price layer types get all operation.

	Typically these are written to a http.Request.
*/
type PriceLayerTypesGetAllParams struct {

	/* MaintenanceMode.

	   Ignore control grouping (default: false)
	*/
	MaintenanceMode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the price layer types get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PriceLayerTypesGetAllParams) WithDefaults() *PriceLayerTypesGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the price layer types get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PriceLayerTypesGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the price layer types get all params
func (o *PriceLayerTypesGetAllParams) WithTimeout(timeout time.Duration) *PriceLayerTypesGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the price layer types get all params
func (o *PriceLayerTypesGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the price layer types get all params
func (o *PriceLayerTypesGetAllParams) WithContext(ctx context.Context) *PriceLayerTypesGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the price layer types get all params
func (o *PriceLayerTypesGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the price layer types get all params
func (o *PriceLayerTypesGetAllParams) WithHTTPClient(client *http.Client) *PriceLayerTypesGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the price layer types get all params
func (o *PriceLayerTypesGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMaintenanceMode adds the maintenanceMode to the price layer types get all params
func (o *PriceLayerTypesGetAllParams) WithMaintenanceMode(maintenanceMode *string) *PriceLayerTypesGetAllParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the price layer types get all params
func (o *PriceLayerTypesGetAllParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *PriceLayerTypesGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
