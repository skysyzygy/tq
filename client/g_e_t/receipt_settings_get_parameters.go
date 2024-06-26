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

// NewReceiptSettingsGetParams creates a new ReceiptSettingsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReceiptSettingsGetParams() *ReceiptSettingsGetParams {
	return &ReceiptSettingsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReceiptSettingsGetParamsWithTimeout creates a new ReceiptSettingsGetParams object
// with the ability to set a timeout on a request.
func NewReceiptSettingsGetParamsWithTimeout(timeout time.Duration) *ReceiptSettingsGetParams {
	return &ReceiptSettingsGetParams{
		timeout: timeout,
	}
}

// NewReceiptSettingsGetParamsWithContext creates a new ReceiptSettingsGetParams object
// with the ability to set a context for a request.
func NewReceiptSettingsGetParamsWithContext(ctx context.Context) *ReceiptSettingsGetParams {
	return &ReceiptSettingsGetParams{
		Context: ctx,
	}
}

// NewReceiptSettingsGetParamsWithHTTPClient creates a new ReceiptSettingsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewReceiptSettingsGetParamsWithHTTPClient(client *http.Client) *ReceiptSettingsGetParams {
	return &ReceiptSettingsGetParams{
		HTTPClient: client,
	}
}

/*
ReceiptSettingsGetParams contains all the parameters to send to the API endpoint

	for the receipt settings get operation.

	Typically these are written to a http.Request.
*/
type ReceiptSettingsGetParams struct {

	/* Filter.

	   Filter by user access (default: readwrite)
	*/
	Filter *string

	/* ID.

	   The id of the resource
	*/
	ID string

	/* MaintenanceMode.

	   Ignore control grouping (default: false)
	*/
	MaintenanceMode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the receipt settings get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReceiptSettingsGetParams) WithDefaults() *ReceiptSettingsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the receipt settings get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReceiptSettingsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the receipt settings get params
func (o *ReceiptSettingsGetParams) WithTimeout(timeout time.Duration) *ReceiptSettingsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the receipt settings get params
func (o *ReceiptSettingsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the receipt settings get params
func (o *ReceiptSettingsGetParams) WithContext(ctx context.Context) *ReceiptSettingsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the receipt settings get params
func (o *ReceiptSettingsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the receipt settings get params
func (o *ReceiptSettingsGetParams) WithHTTPClient(client *http.Client) *ReceiptSettingsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the receipt settings get params
func (o *ReceiptSettingsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the receipt settings get params
func (o *ReceiptSettingsGetParams) WithFilter(filter *string) *ReceiptSettingsGetParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the receipt settings get params
func (o *ReceiptSettingsGetParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the receipt settings get params
func (o *ReceiptSettingsGetParams) WithID(id string) *ReceiptSettingsGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the receipt settings get params
func (o *ReceiptSettingsGetParams) SetID(id string) {
	o.ID = id
}

// WithMaintenanceMode adds the maintenanceMode to the receipt settings get params
func (o *ReceiptSettingsGetParams) WithMaintenanceMode(maintenanceMode *string) *ReceiptSettingsGetParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the receipt settings get params
func (o *ReceiptSettingsGetParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *ReceiptSettingsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
