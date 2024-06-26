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

// NewSeasonTypesGetParams creates a new SeasonTypesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSeasonTypesGetParams() *SeasonTypesGetParams {
	return &SeasonTypesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSeasonTypesGetParamsWithTimeout creates a new SeasonTypesGetParams object
// with the ability to set a timeout on a request.
func NewSeasonTypesGetParamsWithTimeout(timeout time.Duration) *SeasonTypesGetParams {
	return &SeasonTypesGetParams{
		timeout: timeout,
	}
}

// NewSeasonTypesGetParamsWithContext creates a new SeasonTypesGetParams object
// with the ability to set a context for a request.
func NewSeasonTypesGetParamsWithContext(ctx context.Context) *SeasonTypesGetParams {
	return &SeasonTypesGetParams{
		Context: ctx,
	}
}

// NewSeasonTypesGetParamsWithHTTPClient creates a new SeasonTypesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSeasonTypesGetParamsWithHTTPClient(client *http.Client) *SeasonTypesGetParams {
	return &SeasonTypesGetParams{
		HTTPClient: client,
	}
}

/*
SeasonTypesGetParams contains all the parameters to send to the API endpoint

	for the season types get operation.

	Typically these are written to a http.Request.
*/
type SeasonTypesGetParams struct {

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

// WithDefaults hydrates default values in the season types get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SeasonTypesGetParams) WithDefaults() *SeasonTypesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the season types get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SeasonTypesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the season types get params
func (o *SeasonTypesGetParams) WithTimeout(timeout time.Duration) *SeasonTypesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the season types get params
func (o *SeasonTypesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the season types get params
func (o *SeasonTypesGetParams) WithContext(ctx context.Context) *SeasonTypesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the season types get params
func (o *SeasonTypesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the season types get params
func (o *SeasonTypesGetParams) WithHTTPClient(client *http.Client) *SeasonTypesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the season types get params
func (o *SeasonTypesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the season types get params
func (o *SeasonTypesGetParams) WithID(id string) *SeasonTypesGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the season types get params
func (o *SeasonTypesGetParams) SetID(id string) {
	o.ID = id
}

// WithMaintenanceMode adds the maintenanceMode to the season types get params
func (o *SeasonTypesGetParams) WithMaintenanceMode(maintenanceMode *string) *SeasonTypesGetParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the season types get params
func (o *SeasonTypesGetParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *SeasonTypesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
