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

// NewComposersGetParams creates a new ComposersGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewComposersGetParams() *ComposersGetParams {
	return &ComposersGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewComposersGetParamsWithTimeout creates a new ComposersGetParams object
// with the ability to set a timeout on a request.
func NewComposersGetParamsWithTimeout(timeout time.Duration) *ComposersGetParams {
	return &ComposersGetParams{
		timeout: timeout,
	}
}

// NewComposersGetParamsWithContext creates a new ComposersGetParams object
// with the ability to set a context for a request.
func NewComposersGetParamsWithContext(ctx context.Context) *ComposersGetParams {
	return &ComposersGetParams{
		Context: ctx,
	}
}

// NewComposersGetParamsWithHTTPClient creates a new ComposersGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewComposersGetParamsWithHTTPClient(client *http.Client) *ComposersGetParams {
	return &ComposersGetParams{
		HTTPClient: client,
	}
}

/*
ComposersGetParams contains all the parameters to send to the API endpoint

	for the composers get operation.

	Typically these are written to a http.Request.
*/
type ComposersGetParams struct {

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

// WithDefaults hydrates default values in the composers get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ComposersGetParams) WithDefaults() *ComposersGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the composers get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ComposersGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the composers get params
func (o *ComposersGetParams) WithTimeout(timeout time.Duration) *ComposersGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the composers get params
func (o *ComposersGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the composers get params
func (o *ComposersGetParams) WithContext(ctx context.Context) *ComposersGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the composers get params
func (o *ComposersGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the composers get params
func (o *ComposersGetParams) WithHTTPClient(client *http.Client) *ComposersGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the composers get params
func (o *ComposersGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the composers get params
func (o *ComposersGetParams) WithID(id string) *ComposersGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the composers get params
func (o *ComposersGetParams) SetID(id string) {
	o.ID = id
}

// WithMaintenanceMode adds the maintenanceMode to the composers get params
func (o *ComposersGetParams) WithMaintenanceMode(maintenanceMode *string) *ComposersGetParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the composers get params
func (o *ComposersGetParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *ComposersGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
