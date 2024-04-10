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

// NewInterestCategoriesGetAllParams creates a new InterestCategoriesGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInterestCategoriesGetAllParams() *InterestCategoriesGetAllParams {
	return &InterestCategoriesGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInterestCategoriesGetAllParamsWithTimeout creates a new InterestCategoriesGetAllParams object
// with the ability to set a timeout on a request.
func NewInterestCategoriesGetAllParamsWithTimeout(timeout time.Duration) *InterestCategoriesGetAllParams {
	return &InterestCategoriesGetAllParams{
		timeout: timeout,
	}
}

// NewInterestCategoriesGetAllParamsWithContext creates a new InterestCategoriesGetAllParams object
// with the ability to set a context for a request.
func NewInterestCategoriesGetAllParamsWithContext(ctx context.Context) *InterestCategoriesGetAllParams {
	return &InterestCategoriesGetAllParams{
		Context: ctx,
	}
}

// NewInterestCategoriesGetAllParamsWithHTTPClient creates a new InterestCategoriesGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewInterestCategoriesGetAllParamsWithHTTPClient(client *http.Client) *InterestCategoriesGetAllParams {
	return &InterestCategoriesGetAllParams{
		HTTPClient: client,
	}
}

/*
InterestCategoriesGetAllParams contains all the parameters to send to the API endpoint

	for the interest categories get all operation.

	Typically these are written to a http.Request.
*/
type InterestCategoriesGetAllParams struct {

	/* MaintenanceMode.

	   Ignore control grouping (default: false)
	*/
	MaintenanceMode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the interest categories get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InterestCategoriesGetAllParams) WithDefaults() *InterestCategoriesGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the interest categories get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InterestCategoriesGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the interest categories get all params
func (o *InterestCategoriesGetAllParams) WithTimeout(timeout time.Duration) *InterestCategoriesGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the interest categories get all params
func (o *InterestCategoriesGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the interest categories get all params
func (o *InterestCategoriesGetAllParams) WithContext(ctx context.Context) *InterestCategoriesGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the interest categories get all params
func (o *InterestCategoriesGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the interest categories get all params
func (o *InterestCategoriesGetAllParams) WithHTTPClient(client *http.Client) *InterestCategoriesGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the interest categories get all params
func (o *InterestCategoriesGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMaintenanceMode adds the maintenanceMode to the interest categories get all params
func (o *InterestCategoriesGetAllParams) WithMaintenanceMode(maintenanceMode *string) *InterestCategoriesGetAllParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the interest categories get all params
func (o *InterestCategoriesGetAllParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *InterestCategoriesGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
