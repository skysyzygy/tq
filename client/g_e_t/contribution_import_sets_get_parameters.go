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

// NewContributionImportSetsGetParams creates a new ContributionImportSetsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContributionImportSetsGetParams() *ContributionImportSetsGetParams {
	return &ContributionImportSetsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContributionImportSetsGetParamsWithTimeout creates a new ContributionImportSetsGetParams object
// with the ability to set a timeout on a request.
func NewContributionImportSetsGetParamsWithTimeout(timeout time.Duration) *ContributionImportSetsGetParams {
	return &ContributionImportSetsGetParams{
		timeout: timeout,
	}
}

// NewContributionImportSetsGetParamsWithContext creates a new ContributionImportSetsGetParams object
// with the ability to set a context for a request.
func NewContributionImportSetsGetParamsWithContext(ctx context.Context) *ContributionImportSetsGetParams {
	return &ContributionImportSetsGetParams{
		Context: ctx,
	}
}

// NewContributionImportSetsGetParamsWithHTTPClient creates a new ContributionImportSetsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewContributionImportSetsGetParamsWithHTTPClient(client *http.Client) *ContributionImportSetsGetParams {
	return &ContributionImportSetsGetParams{
		HTTPClient: client,
	}
}

/*
ContributionImportSetsGetParams contains all the parameters to send to the API endpoint

	for the contribution import sets get operation.

	Typically these are written to a http.Request.
*/
type ContributionImportSetsGetParams struct {

	// ID.
	ID string

	// MaintenanceMode.
	MaintenanceMode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the contribution import sets get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContributionImportSetsGetParams) WithDefaults() *ContributionImportSetsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the contribution import sets get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContributionImportSetsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the contribution import sets get params
func (o *ContributionImportSetsGetParams) WithTimeout(timeout time.Duration) *ContributionImportSetsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contribution import sets get params
func (o *ContributionImportSetsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contribution import sets get params
func (o *ContributionImportSetsGetParams) WithContext(ctx context.Context) *ContributionImportSetsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contribution import sets get params
func (o *ContributionImportSetsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contribution import sets get params
func (o *ContributionImportSetsGetParams) WithHTTPClient(client *http.Client) *ContributionImportSetsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contribution import sets get params
func (o *ContributionImportSetsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the contribution import sets get params
func (o *ContributionImportSetsGetParams) WithID(id string) *ContributionImportSetsGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the contribution import sets get params
func (o *ContributionImportSetsGetParams) SetID(id string) {
	o.ID = id
}

// WithMaintenanceMode adds the maintenanceMode to the contribution import sets get params
func (o *ContributionImportSetsGetParams) WithMaintenanceMode(maintenanceMode *string) *ContributionImportSetsGetParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the contribution import sets get params
func (o *ContributionImportSetsGetParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *ContributionImportSetsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
