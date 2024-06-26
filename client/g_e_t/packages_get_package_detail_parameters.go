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

// NewPackagesGetPackageDetailParams creates a new PackagesGetPackageDetailParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPackagesGetPackageDetailParams() *PackagesGetPackageDetailParams {
	return &PackagesGetPackageDetailParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPackagesGetPackageDetailParamsWithTimeout creates a new PackagesGetPackageDetailParams object
// with the ability to set a timeout on a request.
func NewPackagesGetPackageDetailParamsWithTimeout(timeout time.Duration) *PackagesGetPackageDetailParams {
	return &PackagesGetPackageDetailParams{
		timeout: timeout,
	}
}

// NewPackagesGetPackageDetailParamsWithContext creates a new PackagesGetPackageDetailParams object
// with the ability to set a context for a request.
func NewPackagesGetPackageDetailParamsWithContext(ctx context.Context) *PackagesGetPackageDetailParams {
	return &PackagesGetPackageDetailParams{
		Context: ctx,
	}
}

// NewPackagesGetPackageDetailParamsWithHTTPClient creates a new PackagesGetPackageDetailParams object
// with the ability to set a custom HTTPClient for a request.
func NewPackagesGetPackageDetailParamsWithHTTPClient(client *http.Client) *PackagesGetPackageDetailParams {
	return &PackagesGetPackageDetailParams{
		HTTPClient: client,
	}
}

/*
PackagesGetPackageDetailParams contains all the parameters to send to the API endpoint

	for the packages get package detail operation.

	Typically these are written to a http.Request.
*/
type PackagesGetPackageDetailParams struct {

	/* AsOfDateTime.

	   If not supplied and there isn't a package date, asOfDateTime will automatically be set to the current date. If not supplied and there is a package date, asOfDateTime will automatically be set to the package date.
	*/
	AsOfDateTime *string

	// ModeOfSaleID.
	ModeOfSaleID *string

	// PackageID.
	PackageID string

	// PriceTypeID.
	PriceTypeID *string

	// SourceID.
	SourceID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the packages get package detail params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackagesGetPackageDetailParams) WithDefaults() *PackagesGetPackageDetailParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the packages get package detail params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackagesGetPackageDetailParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the packages get package detail params
func (o *PackagesGetPackageDetailParams) WithTimeout(timeout time.Duration) *PackagesGetPackageDetailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the packages get package detail params
func (o *PackagesGetPackageDetailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the packages get package detail params
func (o *PackagesGetPackageDetailParams) WithContext(ctx context.Context) *PackagesGetPackageDetailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the packages get package detail params
func (o *PackagesGetPackageDetailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the packages get package detail params
func (o *PackagesGetPackageDetailParams) WithHTTPClient(client *http.Client) *PackagesGetPackageDetailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the packages get package detail params
func (o *PackagesGetPackageDetailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAsOfDateTime adds the asOfDateTime to the packages get package detail params
func (o *PackagesGetPackageDetailParams) WithAsOfDateTime(asOfDateTime *string) *PackagesGetPackageDetailParams {
	o.SetAsOfDateTime(asOfDateTime)
	return o
}

// SetAsOfDateTime adds the asOfDateTime to the packages get package detail params
func (o *PackagesGetPackageDetailParams) SetAsOfDateTime(asOfDateTime *string) {
	o.AsOfDateTime = asOfDateTime
}

// WithModeOfSaleID adds the modeOfSaleID to the packages get package detail params
func (o *PackagesGetPackageDetailParams) WithModeOfSaleID(modeOfSaleID *string) *PackagesGetPackageDetailParams {
	o.SetModeOfSaleID(modeOfSaleID)
	return o
}

// SetModeOfSaleID adds the modeOfSaleId to the packages get package detail params
func (o *PackagesGetPackageDetailParams) SetModeOfSaleID(modeOfSaleID *string) {
	o.ModeOfSaleID = modeOfSaleID
}

// WithPackageID adds the packageID to the packages get package detail params
func (o *PackagesGetPackageDetailParams) WithPackageID(packageID string) *PackagesGetPackageDetailParams {
	o.SetPackageID(packageID)
	return o
}

// SetPackageID adds the packageId to the packages get package detail params
func (o *PackagesGetPackageDetailParams) SetPackageID(packageID string) {
	o.PackageID = packageID
}

// WithPriceTypeID adds the priceTypeID to the packages get package detail params
func (o *PackagesGetPackageDetailParams) WithPriceTypeID(priceTypeID *string) *PackagesGetPackageDetailParams {
	o.SetPriceTypeID(priceTypeID)
	return o
}

// SetPriceTypeID adds the priceTypeId to the packages get package detail params
func (o *PackagesGetPackageDetailParams) SetPriceTypeID(priceTypeID *string) {
	o.PriceTypeID = priceTypeID
}

// WithSourceID adds the sourceID to the packages get package detail params
func (o *PackagesGetPackageDetailParams) WithSourceID(sourceID *string) *PackagesGetPackageDetailParams {
	o.SetSourceID(sourceID)
	return o
}

// SetSourceID adds the sourceId to the packages get package detail params
func (o *PackagesGetPackageDetailParams) SetSourceID(sourceID *string) {
	o.SourceID = sourceID
}

// WriteToRequest writes these params to a swagger request
func (o *PackagesGetPackageDetailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AsOfDateTime != nil {

		// query param asOfDateTime
		var qrAsOfDateTime string

		if o.AsOfDateTime != nil {
			qrAsOfDateTime = *o.AsOfDateTime
		}
		qAsOfDateTime := qrAsOfDateTime
		if qAsOfDateTime != "" {

			if err := r.SetQueryParam("asOfDateTime", qAsOfDateTime); err != nil {
				return err
			}
		}
	}

	if o.ModeOfSaleID != nil {

		// query param modeOfSaleId
		var qrModeOfSaleID string

		if o.ModeOfSaleID != nil {
			qrModeOfSaleID = *o.ModeOfSaleID
		}
		qModeOfSaleID := qrModeOfSaleID
		if qModeOfSaleID != "" {

			if err := r.SetQueryParam("modeOfSaleId", qModeOfSaleID); err != nil {
				return err
			}
		}
	}

	// path param packageId
	if err := r.SetPathParam("packageId", o.PackageID); err != nil {
		return err
	}

	if o.PriceTypeID != nil {

		// query param priceTypeId
		var qrPriceTypeID string

		if o.PriceTypeID != nil {
			qrPriceTypeID = *o.PriceTypeID
		}
		qPriceTypeID := qrPriceTypeID
		if qPriceTypeID != "" {

			if err := r.SetQueryParam("priceTypeId", qPriceTypeID); err != nil {
				return err
			}
		}
	}

	if o.SourceID != nil {

		// query param sourceId
		var qrSourceID string

		if o.SourceID != nil {
			qrSourceID = *o.SourceID
		}
		qSourceID := qrSourceID
		if qSourceID != "" {

			if err := r.SetQueryParam("sourceId", qSourceID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
