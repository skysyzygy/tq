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

// NewDirectDebitAccountTypesGetParams creates a new DirectDebitAccountTypesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDirectDebitAccountTypesGetParams() *DirectDebitAccountTypesGetParams {
	return &DirectDebitAccountTypesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDirectDebitAccountTypesGetParamsWithTimeout creates a new DirectDebitAccountTypesGetParams object
// with the ability to set a timeout on a request.
func NewDirectDebitAccountTypesGetParamsWithTimeout(timeout time.Duration) *DirectDebitAccountTypesGetParams {
	return &DirectDebitAccountTypesGetParams{
		timeout: timeout,
	}
}

// NewDirectDebitAccountTypesGetParamsWithContext creates a new DirectDebitAccountTypesGetParams object
// with the ability to set a context for a request.
func NewDirectDebitAccountTypesGetParamsWithContext(ctx context.Context) *DirectDebitAccountTypesGetParams {
	return &DirectDebitAccountTypesGetParams{
		Context: ctx,
	}
}

// NewDirectDebitAccountTypesGetParamsWithHTTPClient creates a new DirectDebitAccountTypesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewDirectDebitAccountTypesGetParamsWithHTTPClient(client *http.Client) *DirectDebitAccountTypesGetParams {
	return &DirectDebitAccountTypesGetParams{
		HTTPClient: client,
	}
}

/*
DirectDebitAccountTypesGetParams contains all the parameters to send to the API endpoint

	for the direct debit account types get operation.

	Typically these are written to a http.Request.
*/
type DirectDebitAccountTypesGetParams struct {

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

// WithDefaults hydrates default values in the direct debit account types get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DirectDebitAccountTypesGetParams) WithDefaults() *DirectDebitAccountTypesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the direct debit account types get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DirectDebitAccountTypesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the direct debit account types get params
func (o *DirectDebitAccountTypesGetParams) WithTimeout(timeout time.Duration) *DirectDebitAccountTypesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the direct debit account types get params
func (o *DirectDebitAccountTypesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the direct debit account types get params
func (o *DirectDebitAccountTypesGetParams) WithContext(ctx context.Context) *DirectDebitAccountTypesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the direct debit account types get params
func (o *DirectDebitAccountTypesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the direct debit account types get params
func (o *DirectDebitAccountTypesGetParams) WithHTTPClient(client *http.Client) *DirectDebitAccountTypesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the direct debit account types get params
func (o *DirectDebitAccountTypesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the direct debit account types get params
func (o *DirectDebitAccountTypesGetParams) WithID(id string) *DirectDebitAccountTypesGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the direct debit account types get params
func (o *DirectDebitAccountTypesGetParams) SetID(id string) {
	o.ID = id
}

// WithMaintenanceMode adds the maintenanceMode to the direct debit account types get params
func (o *DirectDebitAccountTypesGetParams) WithMaintenanceMode(maintenanceMode *string) *DirectDebitAccountTypesGetParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the direct debit account types get params
func (o *DirectDebitAccountTypesGetParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *DirectDebitAccountTypesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
