// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

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

	"github.com/skysyzygy/tq/models"
)

// NewPackageTypesUpdateParams creates a new PackageTypesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPackageTypesUpdateParams() *PackageTypesUpdateParams {
	return &PackageTypesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPackageTypesUpdateParamsWithTimeout creates a new PackageTypesUpdateParams object
// with the ability to set a timeout on a request.
func NewPackageTypesUpdateParamsWithTimeout(timeout time.Duration) *PackageTypesUpdateParams {
	return &PackageTypesUpdateParams{
		timeout: timeout,
	}
}

// NewPackageTypesUpdateParamsWithContext creates a new PackageTypesUpdateParams object
// with the ability to set a context for a request.
func NewPackageTypesUpdateParamsWithContext(ctx context.Context) *PackageTypesUpdateParams {
	return &PackageTypesUpdateParams{
		Context: ctx,
	}
}

// NewPackageTypesUpdateParamsWithHTTPClient creates a new PackageTypesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPackageTypesUpdateParamsWithHTTPClient(client *http.Client) *PackageTypesUpdateParams {
	return &PackageTypesUpdateParams{
		HTTPClient: client,
	}
}

/*
PackageTypesUpdateParams contains all the parameters to send to the API endpoint

	for the package types update operation.

	Typically these are written to a http.Request.
*/
type PackageTypesUpdateParams struct {

	// Data.
	Data *models.PackageType

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the package types update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackageTypesUpdateParams) WithDefaults() *PackageTypesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the package types update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackageTypesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the package types update params
func (o *PackageTypesUpdateParams) WithTimeout(timeout time.Duration) *PackageTypesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the package types update params
func (o *PackageTypesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the package types update params
func (o *PackageTypesUpdateParams) WithContext(ctx context.Context) *PackageTypesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the package types update params
func (o *PackageTypesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the package types update params
func (o *PackageTypesUpdateParams) WithHTTPClient(client *http.Client) *PackageTypesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the package types update params
func (o *PackageTypesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the package types update params
func (o *PackageTypesUpdateParams) WithData(data *models.PackageType) *PackageTypesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the package types update params
func (o *PackageTypesUpdateParams) SetData(data *models.PackageType) {
	o.Data = data
}

// WithID adds the id to the package types update params
func (o *PackageTypesUpdateParams) WithID(id string) *PackageTypesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the package types update params
func (o *PackageTypesUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PackageTypesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
