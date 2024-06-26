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

// NewPerformancePackageModeOfSalesGetAllParams creates a new PerformancePackageModeOfSalesGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformancePackageModeOfSalesGetAllParams() *PerformancePackageModeOfSalesGetAllParams {
	return &PerformancePackageModeOfSalesGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformancePackageModeOfSalesGetAllParamsWithTimeout creates a new PerformancePackageModeOfSalesGetAllParams object
// with the ability to set a timeout on a request.
func NewPerformancePackageModeOfSalesGetAllParamsWithTimeout(timeout time.Duration) *PerformancePackageModeOfSalesGetAllParams {
	return &PerformancePackageModeOfSalesGetAllParams{
		timeout: timeout,
	}
}

// NewPerformancePackageModeOfSalesGetAllParamsWithContext creates a new PerformancePackageModeOfSalesGetAllParams object
// with the ability to set a context for a request.
func NewPerformancePackageModeOfSalesGetAllParamsWithContext(ctx context.Context) *PerformancePackageModeOfSalesGetAllParams {
	return &PerformancePackageModeOfSalesGetAllParams{
		Context: ctx,
	}
}

// NewPerformancePackageModeOfSalesGetAllParamsWithHTTPClient creates a new PerformancePackageModeOfSalesGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformancePackageModeOfSalesGetAllParamsWithHTTPClient(client *http.Client) *PerformancePackageModeOfSalesGetAllParams {
	return &PerformancePackageModeOfSalesGetAllParams{
		HTTPClient: client,
	}
}

/*
PerformancePackageModeOfSalesGetAllParams contains all the parameters to send to the API endpoint

	for the performance package mode of sales get all operation.

	Typically these are written to a http.Request.
*/
type PerformancePackageModeOfSalesGetAllParams struct {

	/* AsOfDateTime.

	   A date/time to determine active mode of sale assignments.
	*/
	AsOfDateTime *string

	/* ModeOfSaleID.

	   A single mode of sale to filter results.
	*/
	ModeOfSaleID *string

	/* PackageIds.

	   A comma separated list of package ids.
	*/
	PackageIds *string

	/* PerformanceIds.

	   A comma separated list of performance ids.
	*/
	PerformanceIds *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the performance package mode of sales get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformancePackageModeOfSalesGetAllParams) WithDefaults() *PerformancePackageModeOfSalesGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the performance package mode of sales get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformancePackageModeOfSalesGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the performance package mode of sales get all params
func (o *PerformancePackageModeOfSalesGetAllParams) WithTimeout(timeout time.Duration) *PerformancePackageModeOfSalesGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the performance package mode of sales get all params
func (o *PerformancePackageModeOfSalesGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the performance package mode of sales get all params
func (o *PerformancePackageModeOfSalesGetAllParams) WithContext(ctx context.Context) *PerformancePackageModeOfSalesGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the performance package mode of sales get all params
func (o *PerformancePackageModeOfSalesGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the performance package mode of sales get all params
func (o *PerformancePackageModeOfSalesGetAllParams) WithHTTPClient(client *http.Client) *PerformancePackageModeOfSalesGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the performance package mode of sales get all params
func (o *PerformancePackageModeOfSalesGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAsOfDateTime adds the asOfDateTime to the performance package mode of sales get all params
func (o *PerformancePackageModeOfSalesGetAllParams) WithAsOfDateTime(asOfDateTime *string) *PerformancePackageModeOfSalesGetAllParams {
	o.SetAsOfDateTime(asOfDateTime)
	return o
}

// SetAsOfDateTime adds the asOfDateTime to the performance package mode of sales get all params
func (o *PerformancePackageModeOfSalesGetAllParams) SetAsOfDateTime(asOfDateTime *string) {
	o.AsOfDateTime = asOfDateTime
}

// WithModeOfSaleID adds the modeOfSaleID to the performance package mode of sales get all params
func (o *PerformancePackageModeOfSalesGetAllParams) WithModeOfSaleID(modeOfSaleID *string) *PerformancePackageModeOfSalesGetAllParams {
	o.SetModeOfSaleID(modeOfSaleID)
	return o
}

// SetModeOfSaleID adds the modeOfSaleId to the performance package mode of sales get all params
func (o *PerformancePackageModeOfSalesGetAllParams) SetModeOfSaleID(modeOfSaleID *string) {
	o.ModeOfSaleID = modeOfSaleID
}

// WithPackageIds adds the packageIds to the performance package mode of sales get all params
func (o *PerformancePackageModeOfSalesGetAllParams) WithPackageIds(packageIds *string) *PerformancePackageModeOfSalesGetAllParams {
	o.SetPackageIds(packageIds)
	return o
}

// SetPackageIds adds the packageIds to the performance package mode of sales get all params
func (o *PerformancePackageModeOfSalesGetAllParams) SetPackageIds(packageIds *string) {
	o.PackageIds = packageIds
}

// WithPerformanceIds adds the performanceIds to the performance package mode of sales get all params
func (o *PerformancePackageModeOfSalesGetAllParams) WithPerformanceIds(performanceIds *string) *PerformancePackageModeOfSalesGetAllParams {
	o.SetPerformanceIds(performanceIds)
	return o
}

// SetPerformanceIds adds the performanceIds to the performance package mode of sales get all params
func (o *PerformancePackageModeOfSalesGetAllParams) SetPerformanceIds(performanceIds *string) {
	o.PerformanceIds = performanceIds
}

// WriteToRequest writes these params to a swagger request
func (o *PerformancePackageModeOfSalesGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.PerformanceIds != nil {

		// query param performanceIds
		var qrPerformanceIds string

		if o.PerformanceIds != nil {
			qrPerformanceIds = *o.PerformanceIds
		}
		qPerformanceIds := qrPerformanceIds
		if qPerformanceIds != "" {

			if err := r.SetQueryParam("performanceIds", qPerformanceIds); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
