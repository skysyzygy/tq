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

// NewBookingCategoriesGetParams creates a new BookingCategoriesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBookingCategoriesGetParams() *BookingCategoriesGetParams {
	return &BookingCategoriesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBookingCategoriesGetParamsWithTimeout creates a new BookingCategoriesGetParams object
// with the ability to set a timeout on a request.
func NewBookingCategoriesGetParamsWithTimeout(timeout time.Duration) *BookingCategoriesGetParams {
	return &BookingCategoriesGetParams{
		timeout: timeout,
	}
}

// NewBookingCategoriesGetParamsWithContext creates a new BookingCategoriesGetParams object
// with the ability to set a context for a request.
func NewBookingCategoriesGetParamsWithContext(ctx context.Context) *BookingCategoriesGetParams {
	return &BookingCategoriesGetParams{
		Context: ctx,
	}
}

// NewBookingCategoriesGetParamsWithHTTPClient creates a new BookingCategoriesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewBookingCategoriesGetParamsWithHTTPClient(client *http.Client) *BookingCategoriesGetParams {
	return &BookingCategoriesGetParams{
		HTTPClient: client,
	}
}

/*
BookingCategoriesGetParams contains all the parameters to send to the API endpoint

	for the booking categories get operation.

	Typically these are written to a http.Request.
*/
type BookingCategoriesGetParams struct {

	/* Filter.

	   Filter by user access (default: readwrite)
	*/
	Filter *string

	/* ID.

	   The id of the resource.
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

// WithDefaults hydrates default values in the booking categories get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BookingCategoriesGetParams) WithDefaults() *BookingCategoriesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the booking categories get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BookingCategoriesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the booking categories get params
func (o *BookingCategoriesGetParams) WithTimeout(timeout time.Duration) *BookingCategoriesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the booking categories get params
func (o *BookingCategoriesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the booking categories get params
func (o *BookingCategoriesGetParams) WithContext(ctx context.Context) *BookingCategoriesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the booking categories get params
func (o *BookingCategoriesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the booking categories get params
func (o *BookingCategoriesGetParams) WithHTTPClient(client *http.Client) *BookingCategoriesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the booking categories get params
func (o *BookingCategoriesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the booking categories get params
func (o *BookingCategoriesGetParams) WithFilter(filter *string) *BookingCategoriesGetParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the booking categories get params
func (o *BookingCategoriesGetParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the booking categories get params
func (o *BookingCategoriesGetParams) WithID(id string) *BookingCategoriesGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the booking categories get params
func (o *BookingCategoriesGetParams) SetID(id string) {
	o.ID = id
}

// WithMaintenanceMode adds the maintenanceMode to the booking categories get params
func (o *BookingCategoriesGetParams) WithMaintenanceMode(maintenanceMode *string) *BookingCategoriesGetParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the booking categories get params
func (o *BookingCategoriesGetParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *BookingCategoriesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
