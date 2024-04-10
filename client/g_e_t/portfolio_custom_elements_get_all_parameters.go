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

// NewPortfolioCustomElementsGetAllParams creates a new PortfolioCustomElementsGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPortfolioCustomElementsGetAllParams() *PortfolioCustomElementsGetAllParams {
	return &PortfolioCustomElementsGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPortfolioCustomElementsGetAllParamsWithTimeout creates a new PortfolioCustomElementsGetAllParams object
// with the ability to set a timeout on a request.
func NewPortfolioCustomElementsGetAllParamsWithTimeout(timeout time.Duration) *PortfolioCustomElementsGetAllParams {
	return &PortfolioCustomElementsGetAllParams{
		timeout: timeout,
	}
}

// NewPortfolioCustomElementsGetAllParamsWithContext creates a new PortfolioCustomElementsGetAllParams object
// with the ability to set a context for a request.
func NewPortfolioCustomElementsGetAllParamsWithContext(ctx context.Context) *PortfolioCustomElementsGetAllParams {
	return &PortfolioCustomElementsGetAllParams{
		Context: ctx,
	}
}

// NewPortfolioCustomElementsGetAllParamsWithHTTPClient creates a new PortfolioCustomElementsGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewPortfolioCustomElementsGetAllParamsWithHTTPClient(client *http.Client) *PortfolioCustomElementsGetAllParams {
	return &PortfolioCustomElementsGetAllParams{
		HTTPClient: client,
	}
}

/*
PortfolioCustomElementsGetAllParams contains all the parameters to send to the API endpoint

	for the portfolio custom elements get all operation.

	Typically these are written to a http.Request.
*/
type PortfolioCustomElementsGetAllParams struct {

	/* Filter.

	   Filter by user access (default: readwrite)
	*/
	Filter *string

	/* MaintenanceMode.

	   Ignore control grouping (default: false)
	*/
	MaintenanceMode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the portfolio custom elements get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PortfolioCustomElementsGetAllParams) WithDefaults() *PortfolioCustomElementsGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the portfolio custom elements get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PortfolioCustomElementsGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the portfolio custom elements get all params
func (o *PortfolioCustomElementsGetAllParams) WithTimeout(timeout time.Duration) *PortfolioCustomElementsGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the portfolio custom elements get all params
func (o *PortfolioCustomElementsGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the portfolio custom elements get all params
func (o *PortfolioCustomElementsGetAllParams) WithContext(ctx context.Context) *PortfolioCustomElementsGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the portfolio custom elements get all params
func (o *PortfolioCustomElementsGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the portfolio custom elements get all params
func (o *PortfolioCustomElementsGetAllParams) WithHTTPClient(client *http.Client) *PortfolioCustomElementsGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the portfolio custom elements get all params
func (o *PortfolioCustomElementsGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the portfolio custom elements get all params
func (o *PortfolioCustomElementsGetAllParams) WithFilter(filter *string) *PortfolioCustomElementsGetAllParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the portfolio custom elements get all params
func (o *PortfolioCustomElementsGetAllParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithMaintenanceMode adds the maintenanceMode to the portfolio custom elements get all params
func (o *PortfolioCustomElementsGetAllParams) WithMaintenanceMode(maintenanceMode *string) *PortfolioCustomElementsGetAllParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the portfolio custom elements get all params
func (o *PortfolioCustomElementsGetAllParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *PortfolioCustomElementsGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
