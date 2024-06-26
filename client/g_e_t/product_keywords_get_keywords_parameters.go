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

// NewProductKeywordsGetKeywordsParams creates a new ProductKeywordsGetKeywordsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProductKeywordsGetKeywordsParams() *ProductKeywordsGetKeywordsParams {
	return &ProductKeywordsGetKeywordsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProductKeywordsGetKeywordsParamsWithTimeout creates a new ProductKeywordsGetKeywordsParams object
// with the ability to set a timeout on a request.
func NewProductKeywordsGetKeywordsParamsWithTimeout(timeout time.Duration) *ProductKeywordsGetKeywordsParams {
	return &ProductKeywordsGetKeywordsParams{
		timeout: timeout,
	}
}

// NewProductKeywordsGetKeywordsParamsWithContext creates a new ProductKeywordsGetKeywordsParams object
// with the ability to set a context for a request.
func NewProductKeywordsGetKeywordsParamsWithContext(ctx context.Context) *ProductKeywordsGetKeywordsParams {
	return &ProductKeywordsGetKeywordsParams{
		Context: ctx,
	}
}

// NewProductKeywordsGetKeywordsParamsWithHTTPClient creates a new ProductKeywordsGetKeywordsParams object
// with the ability to set a custom HTTPClient for a request.
func NewProductKeywordsGetKeywordsParamsWithHTTPClient(client *http.Client) *ProductKeywordsGetKeywordsParams {
	return &ProductKeywordsGetKeywordsParams{
		HTTPClient: client,
	}
}

/*
ProductKeywordsGetKeywordsParams contains all the parameters to send to the API endpoint

	for the product keywords get keywords operation.

	Typically these are written to a http.Request.
*/
type ProductKeywordsGetKeywordsParams struct {

	/* KeywordIds.

	   Must be a comma delimited string of valid keyword ids.
	*/
	KeywordIds *string

	/* PackageIds.

	   Must be a list of valid package ids (must be supplied if productionElementIds is not supplied)
	*/
	PackageIds *string

	/* ProductionElementIds.

	   Must be a list of valid perf, prod season, prod or title Ids (must be supplied if packageIds is not supplied)
	*/
	ProductionElementIds *string

	/* ShowAll.

	   Indicates if keywords from ‘parent’ requested element(s) is desired. This parameter is ignored for packages.
	*/
	ShowAll *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the product keywords get keywords params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProductKeywordsGetKeywordsParams) WithDefaults() *ProductKeywordsGetKeywordsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the product keywords get keywords params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProductKeywordsGetKeywordsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the product keywords get keywords params
func (o *ProductKeywordsGetKeywordsParams) WithTimeout(timeout time.Duration) *ProductKeywordsGetKeywordsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the product keywords get keywords params
func (o *ProductKeywordsGetKeywordsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the product keywords get keywords params
func (o *ProductKeywordsGetKeywordsParams) WithContext(ctx context.Context) *ProductKeywordsGetKeywordsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the product keywords get keywords params
func (o *ProductKeywordsGetKeywordsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the product keywords get keywords params
func (o *ProductKeywordsGetKeywordsParams) WithHTTPClient(client *http.Client) *ProductKeywordsGetKeywordsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the product keywords get keywords params
func (o *ProductKeywordsGetKeywordsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKeywordIds adds the keywordIds to the product keywords get keywords params
func (o *ProductKeywordsGetKeywordsParams) WithKeywordIds(keywordIds *string) *ProductKeywordsGetKeywordsParams {
	o.SetKeywordIds(keywordIds)
	return o
}

// SetKeywordIds adds the keywordIds to the product keywords get keywords params
func (o *ProductKeywordsGetKeywordsParams) SetKeywordIds(keywordIds *string) {
	o.KeywordIds = keywordIds
}

// WithPackageIds adds the packageIds to the product keywords get keywords params
func (o *ProductKeywordsGetKeywordsParams) WithPackageIds(packageIds *string) *ProductKeywordsGetKeywordsParams {
	o.SetPackageIds(packageIds)
	return o
}

// SetPackageIds adds the packageIds to the product keywords get keywords params
func (o *ProductKeywordsGetKeywordsParams) SetPackageIds(packageIds *string) {
	o.PackageIds = packageIds
}

// WithProductionElementIds adds the productionElementIds to the product keywords get keywords params
func (o *ProductKeywordsGetKeywordsParams) WithProductionElementIds(productionElementIds *string) *ProductKeywordsGetKeywordsParams {
	o.SetProductionElementIds(productionElementIds)
	return o
}

// SetProductionElementIds adds the productionElementIds to the product keywords get keywords params
func (o *ProductKeywordsGetKeywordsParams) SetProductionElementIds(productionElementIds *string) {
	o.ProductionElementIds = productionElementIds
}

// WithShowAll adds the showAll to the product keywords get keywords params
func (o *ProductKeywordsGetKeywordsParams) WithShowAll(showAll *string) *ProductKeywordsGetKeywordsParams {
	o.SetShowAll(showAll)
	return o
}

// SetShowAll adds the showAll to the product keywords get keywords params
func (o *ProductKeywordsGetKeywordsParams) SetShowAll(showAll *string) {
	o.ShowAll = showAll
}

// WriteToRequest writes these params to a swagger request
func (o *ProductKeywordsGetKeywordsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.KeywordIds != nil {

		// query param keywordIds
		var qrKeywordIds string

		if o.KeywordIds != nil {
			qrKeywordIds = *o.KeywordIds
		}
		qKeywordIds := qrKeywordIds
		if qKeywordIds != "" {

			if err := r.SetQueryParam("keywordIds", qKeywordIds); err != nil {
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

	if o.ProductionElementIds != nil {

		// query param productionElementIds
		var qrProductionElementIds string

		if o.ProductionElementIds != nil {
			qrProductionElementIds = *o.ProductionElementIds
		}
		qProductionElementIds := qrProductionElementIds
		if qProductionElementIds != "" {

			if err := r.SetQueryParam("productionElementIds", qProductionElementIds); err != nil {
				return err
			}
		}
	}

	if o.ShowAll != nil {

		// query param showAll
		var qrShowAll string

		if o.ShowAll != nil {
			qrShowAll = *o.ShowAll
		}
		qShowAll := qrShowAll
		if qShowAll != "" {

			if err := r.SetQueryParam("showAll", qShowAll); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
