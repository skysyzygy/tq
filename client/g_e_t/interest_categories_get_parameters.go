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

// NewInterestCategoriesGetParams creates a new InterestCategoriesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInterestCategoriesGetParams() *InterestCategoriesGetParams {
	return &InterestCategoriesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInterestCategoriesGetParamsWithTimeout creates a new InterestCategoriesGetParams object
// with the ability to set a timeout on a request.
func NewInterestCategoriesGetParamsWithTimeout(timeout time.Duration) *InterestCategoriesGetParams {
	return &InterestCategoriesGetParams{
		timeout: timeout,
	}
}

// NewInterestCategoriesGetParamsWithContext creates a new InterestCategoriesGetParams object
// with the ability to set a context for a request.
func NewInterestCategoriesGetParamsWithContext(ctx context.Context) *InterestCategoriesGetParams {
	return &InterestCategoriesGetParams{
		Context: ctx,
	}
}

// NewInterestCategoriesGetParamsWithHTTPClient creates a new InterestCategoriesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewInterestCategoriesGetParamsWithHTTPClient(client *http.Client) *InterestCategoriesGetParams {
	return &InterestCategoriesGetParams{
		HTTPClient: client,
	}
}

/*
InterestCategoriesGetParams contains all the parameters to send to the API endpoint

	for the interest categories get operation.

	Typically these are written to a http.Request.
*/
type InterestCategoriesGetParams struct {

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

// WithDefaults hydrates default values in the interest categories get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InterestCategoriesGetParams) WithDefaults() *InterestCategoriesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the interest categories get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InterestCategoriesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the interest categories get params
func (o *InterestCategoriesGetParams) WithTimeout(timeout time.Duration) *InterestCategoriesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the interest categories get params
func (o *InterestCategoriesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the interest categories get params
func (o *InterestCategoriesGetParams) WithContext(ctx context.Context) *InterestCategoriesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the interest categories get params
func (o *InterestCategoriesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the interest categories get params
func (o *InterestCategoriesGetParams) WithHTTPClient(client *http.Client) *InterestCategoriesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the interest categories get params
func (o *InterestCategoriesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the interest categories get params
func (o *InterestCategoriesGetParams) WithID(id string) *InterestCategoriesGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the interest categories get params
func (o *InterestCategoriesGetParams) SetID(id string) {
	o.ID = id
}

// WithMaintenanceMode adds the maintenanceMode to the interest categories get params
func (o *InterestCategoriesGetParams) WithMaintenanceMode(maintenanceMode *string) *InterestCategoriesGetParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the interest categories get params
func (o *InterestCategoriesGetParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *InterestCategoriesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
