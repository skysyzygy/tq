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

// NewEmailProfilesGetParams creates a new EmailProfilesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEmailProfilesGetParams() *EmailProfilesGetParams {
	return &EmailProfilesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEmailProfilesGetParamsWithTimeout creates a new EmailProfilesGetParams object
// with the ability to set a timeout on a request.
func NewEmailProfilesGetParamsWithTimeout(timeout time.Duration) *EmailProfilesGetParams {
	return &EmailProfilesGetParams{
		timeout: timeout,
	}
}

// NewEmailProfilesGetParamsWithContext creates a new EmailProfilesGetParams object
// with the ability to set a context for a request.
func NewEmailProfilesGetParamsWithContext(ctx context.Context) *EmailProfilesGetParams {
	return &EmailProfilesGetParams{
		Context: ctx,
	}
}

// NewEmailProfilesGetParamsWithHTTPClient creates a new EmailProfilesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewEmailProfilesGetParamsWithHTTPClient(client *http.Client) *EmailProfilesGetParams {
	return &EmailProfilesGetParams{
		HTTPClient: client,
	}
}

/*
EmailProfilesGetParams contains all the parameters to send to the API endpoint

	for the email profiles get operation.

	Typically these are written to a http.Request.
*/
type EmailProfilesGetParams struct {

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

// WithDefaults hydrates default values in the email profiles get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmailProfilesGetParams) WithDefaults() *EmailProfilesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the email profiles get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmailProfilesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the email profiles get params
func (o *EmailProfilesGetParams) WithTimeout(timeout time.Duration) *EmailProfilesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the email profiles get params
func (o *EmailProfilesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the email profiles get params
func (o *EmailProfilesGetParams) WithContext(ctx context.Context) *EmailProfilesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the email profiles get params
func (o *EmailProfilesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the email profiles get params
func (o *EmailProfilesGetParams) WithHTTPClient(client *http.Client) *EmailProfilesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the email profiles get params
func (o *EmailProfilesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the email profiles get params
func (o *EmailProfilesGetParams) WithFilter(filter *string) *EmailProfilesGetParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the email profiles get params
func (o *EmailProfilesGetParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the email profiles get params
func (o *EmailProfilesGetParams) WithID(id string) *EmailProfilesGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the email profiles get params
func (o *EmailProfilesGetParams) SetID(id string) {
	o.ID = id
}

// WithMaintenanceMode adds the maintenanceMode to the email profiles get params
func (o *EmailProfilesGetParams) WithMaintenanceMode(maintenanceMode *string) *EmailProfilesGetParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the email profiles get params
func (o *EmailProfilesGetParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *EmailProfilesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
