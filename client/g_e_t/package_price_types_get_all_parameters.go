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

// NewPackagePriceTypesGetAllParams creates a new PackagePriceTypesGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPackagePriceTypesGetAllParams() *PackagePriceTypesGetAllParams {
	return &PackagePriceTypesGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPackagePriceTypesGetAllParamsWithTimeout creates a new PackagePriceTypesGetAllParams object
// with the ability to set a timeout on a request.
func NewPackagePriceTypesGetAllParamsWithTimeout(timeout time.Duration) *PackagePriceTypesGetAllParams {
	return &PackagePriceTypesGetAllParams{
		timeout: timeout,
	}
}

// NewPackagePriceTypesGetAllParamsWithContext creates a new PackagePriceTypesGetAllParams object
// with the ability to set a context for a request.
func NewPackagePriceTypesGetAllParamsWithContext(ctx context.Context) *PackagePriceTypesGetAllParams {
	return &PackagePriceTypesGetAllParams{
		Context: ctx,
	}
}

// NewPackagePriceTypesGetAllParamsWithHTTPClient creates a new PackagePriceTypesGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewPackagePriceTypesGetAllParamsWithHTTPClient(client *http.Client) *PackagePriceTypesGetAllParams {
	return &PackagePriceTypesGetAllParams{
		HTTPClient: client,
	}
}

/*
PackagePriceTypesGetAllParams contains all the parameters to send to the API endpoint

	for the package price types get all operation.

	Typically these are written to a http.Request.
*/
type PackagePriceTypesGetAllParams struct {

	/* PackageIds.

	   Required to filter package pricetypes to one or more packages.
	*/
	PackageIds *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the package price types get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackagePriceTypesGetAllParams) WithDefaults() *PackagePriceTypesGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the package price types get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackagePriceTypesGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the package price types get all params
func (o *PackagePriceTypesGetAllParams) WithTimeout(timeout time.Duration) *PackagePriceTypesGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the package price types get all params
func (o *PackagePriceTypesGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the package price types get all params
func (o *PackagePriceTypesGetAllParams) WithContext(ctx context.Context) *PackagePriceTypesGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the package price types get all params
func (o *PackagePriceTypesGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the package price types get all params
func (o *PackagePriceTypesGetAllParams) WithHTTPClient(client *http.Client) *PackagePriceTypesGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the package price types get all params
func (o *PackagePriceTypesGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageIds adds the packageIds to the package price types get all params
func (o *PackagePriceTypesGetAllParams) WithPackageIds(packageIds *string) *PackagePriceTypesGetAllParams {
	o.SetPackageIds(packageIds)
	return o
}

// SetPackageIds adds the packageIds to the package price types get all params
func (o *PackagePriceTypesGetAllParams) SetPackageIds(packageIds *string) {
	o.PackageIds = packageIds
}

// WriteToRequest writes these params to a swagger request
func (o *PackagePriceTypesGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PackageIds != nil {

		// query param packageIds
		var qrPackageIds string

		if o.PackageIds != nil {
			qrPackageIds = *o.PackageIds
		}
		qPackageIds := qrPackageIds
		if qPackageIds != "" {

			if err := r.SetQueryParam("packageIds", qPackageIds); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
