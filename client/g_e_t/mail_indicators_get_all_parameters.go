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

// NewMailIndicatorsGetAllParams creates a new MailIndicatorsGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMailIndicatorsGetAllParams() *MailIndicatorsGetAllParams {
	return &MailIndicatorsGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMailIndicatorsGetAllParamsWithTimeout creates a new MailIndicatorsGetAllParams object
// with the ability to set a timeout on a request.
func NewMailIndicatorsGetAllParamsWithTimeout(timeout time.Duration) *MailIndicatorsGetAllParams {
	return &MailIndicatorsGetAllParams{
		timeout: timeout,
	}
}

// NewMailIndicatorsGetAllParamsWithContext creates a new MailIndicatorsGetAllParams object
// with the ability to set a context for a request.
func NewMailIndicatorsGetAllParamsWithContext(ctx context.Context) *MailIndicatorsGetAllParams {
	return &MailIndicatorsGetAllParams{
		Context: ctx,
	}
}

// NewMailIndicatorsGetAllParamsWithHTTPClient creates a new MailIndicatorsGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewMailIndicatorsGetAllParamsWithHTTPClient(client *http.Client) *MailIndicatorsGetAllParams {
	return &MailIndicatorsGetAllParams{
		HTTPClient: client,
	}
}

/*
MailIndicatorsGetAllParams contains all the parameters to send to the API endpoint

	for the mail indicators get all operation.

	Typically these are written to a http.Request.
*/
type MailIndicatorsGetAllParams struct {

	/* MaintenanceMode.

	   Ignore control grouping (default: false)
	*/
	MaintenanceMode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the mail indicators get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MailIndicatorsGetAllParams) WithDefaults() *MailIndicatorsGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the mail indicators get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MailIndicatorsGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the mail indicators get all params
func (o *MailIndicatorsGetAllParams) WithTimeout(timeout time.Duration) *MailIndicatorsGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the mail indicators get all params
func (o *MailIndicatorsGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the mail indicators get all params
func (o *MailIndicatorsGetAllParams) WithContext(ctx context.Context) *MailIndicatorsGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the mail indicators get all params
func (o *MailIndicatorsGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the mail indicators get all params
func (o *MailIndicatorsGetAllParams) WithHTTPClient(client *http.Client) *MailIndicatorsGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the mail indicators get all params
func (o *MailIndicatorsGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMaintenanceMode adds the maintenanceMode to the mail indicators get all params
func (o *MailIndicatorsGetAllParams) WithMaintenanceMode(maintenanceMode *string) *MailIndicatorsGetAllParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the mail indicators get all params
func (o *MailIndicatorsGetAllParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *MailIndicatorsGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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