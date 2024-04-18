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

// NewMachineSettingsGetParams creates a new MachineSettingsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMachineSettingsGetParams() *MachineSettingsGetParams {
	return &MachineSettingsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMachineSettingsGetParamsWithTimeout creates a new MachineSettingsGetParams object
// with the ability to set a timeout on a request.
func NewMachineSettingsGetParamsWithTimeout(timeout time.Duration) *MachineSettingsGetParams {
	return &MachineSettingsGetParams{
		timeout: timeout,
	}
}

// NewMachineSettingsGetParamsWithContext creates a new MachineSettingsGetParams object
// with the ability to set a context for a request.
func NewMachineSettingsGetParamsWithContext(ctx context.Context) *MachineSettingsGetParams {
	return &MachineSettingsGetParams{
		Context: ctx,
	}
}

// NewMachineSettingsGetParamsWithHTTPClient creates a new MachineSettingsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewMachineSettingsGetParamsWithHTTPClient(client *http.Client) *MachineSettingsGetParams {
	return &MachineSettingsGetParams{
		HTTPClient: client,
	}
}

/*
MachineSettingsGetParams contains all the parameters to send to the API endpoint

	for the machine settings get operation.

	Typically these are written to a http.Request.
*/
type MachineSettingsGetParams struct {

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

// WithDefaults hydrates default values in the machine settings get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MachineSettingsGetParams) WithDefaults() *MachineSettingsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the machine settings get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MachineSettingsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the machine settings get params
func (o *MachineSettingsGetParams) WithTimeout(timeout time.Duration) *MachineSettingsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the machine settings get params
func (o *MachineSettingsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the machine settings get params
func (o *MachineSettingsGetParams) WithContext(ctx context.Context) *MachineSettingsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the machine settings get params
func (o *MachineSettingsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the machine settings get params
func (o *MachineSettingsGetParams) WithHTTPClient(client *http.Client) *MachineSettingsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the machine settings get params
func (o *MachineSettingsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the machine settings get params
func (o *MachineSettingsGetParams) WithID(id string) *MachineSettingsGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the machine settings get params
func (o *MachineSettingsGetParams) SetID(id string) {
	o.ID = id
}

// WithMaintenanceMode adds the maintenanceMode to the machine settings get params
func (o *MachineSettingsGetParams) WithMaintenanceMode(maintenanceMode *string) *MachineSettingsGetParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the machine settings get params
func (o *MachineSettingsGetParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *MachineSettingsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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