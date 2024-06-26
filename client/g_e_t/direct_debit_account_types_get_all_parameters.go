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

// NewDirectDebitAccountTypesGetAllParams creates a new DirectDebitAccountTypesGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDirectDebitAccountTypesGetAllParams() *DirectDebitAccountTypesGetAllParams {
	return &DirectDebitAccountTypesGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDirectDebitAccountTypesGetAllParamsWithTimeout creates a new DirectDebitAccountTypesGetAllParams object
// with the ability to set a timeout on a request.
func NewDirectDebitAccountTypesGetAllParamsWithTimeout(timeout time.Duration) *DirectDebitAccountTypesGetAllParams {
	return &DirectDebitAccountTypesGetAllParams{
		timeout: timeout,
	}
}

// NewDirectDebitAccountTypesGetAllParamsWithContext creates a new DirectDebitAccountTypesGetAllParams object
// with the ability to set a context for a request.
func NewDirectDebitAccountTypesGetAllParamsWithContext(ctx context.Context) *DirectDebitAccountTypesGetAllParams {
	return &DirectDebitAccountTypesGetAllParams{
		Context: ctx,
	}
}

// NewDirectDebitAccountTypesGetAllParamsWithHTTPClient creates a new DirectDebitAccountTypesGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewDirectDebitAccountTypesGetAllParamsWithHTTPClient(client *http.Client) *DirectDebitAccountTypesGetAllParams {
	return &DirectDebitAccountTypesGetAllParams{
		HTTPClient: client,
	}
}

/*
DirectDebitAccountTypesGetAllParams contains all the parameters to send to the API endpoint

	for the direct debit account types get all operation.

	Typically these are written to a http.Request.
*/
type DirectDebitAccountTypesGetAllParams struct {

	/* MaintenanceMode.

	   Ignore control grouping (default: false)
	*/
	MaintenanceMode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the direct debit account types get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DirectDebitAccountTypesGetAllParams) WithDefaults() *DirectDebitAccountTypesGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the direct debit account types get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DirectDebitAccountTypesGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the direct debit account types get all params
func (o *DirectDebitAccountTypesGetAllParams) WithTimeout(timeout time.Duration) *DirectDebitAccountTypesGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the direct debit account types get all params
func (o *DirectDebitAccountTypesGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the direct debit account types get all params
func (o *DirectDebitAccountTypesGetAllParams) WithContext(ctx context.Context) *DirectDebitAccountTypesGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the direct debit account types get all params
func (o *DirectDebitAccountTypesGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the direct debit account types get all params
func (o *DirectDebitAccountTypesGetAllParams) WithHTTPClient(client *http.Client) *DirectDebitAccountTypesGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the direct debit account types get all params
func (o *DirectDebitAccountTypesGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMaintenanceMode adds the maintenanceMode to the direct debit account types get all params
func (o *DirectDebitAccountTypesGetAllParams) WithMaintenanceMode(maintenanceMode *string) *DirectDebitAccountTypesGetAllParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the direct debit account types get all params
func (o *DirectDebitAccountTypesGetAllParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *DirectDebitAccountTypesGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
