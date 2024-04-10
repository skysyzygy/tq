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

// NewPackageWebContentsGetAllSummariesParams creates a new PackageWebContentsGetAllSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPackageWebContentsGetAllSummariesParams() *PackageWebContentsGetAllSummariesParams {
	return &PackageWebContentsGetAllSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPackageWebContentsGetAllSummariesParamsWithTimeout creates a new PackageWebContentsGetAllSummariesParams object
// with the ability to set a timeout on a request.
func NewPackageWebContentsGetAllSummariesParamsWithTimeout(timeout time.Duration) *PackageWebContentsGetAllSummariesParams {
	return &PackageWebContentsGetAllSummariesParams{
		timeout: timeout,
	}
}

// NewPackageWebContentsGetAllSummariesParamsWithContext creates a new PackageWebContentsGetAllSummariesParams object
// with the ability to set a context for a request.
func NewPackageWebContentsGetAllSummariesParamsWithContext(ctx context.Context) *PackageWebContentsGetAllSummariesParams {
	return &PackageWebContentsGetAllSummariesParams{
		Context: ctx,
	}
}

// NewPackageWebContentsGetAllSummariesParamsWithHTTPClient creates a new PackageWebContentsGetAllSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPackageWebContentsGetAllSummariesParamsWithHTTPClient(client *http.Client) *PackageWebContentsGetAllSummariesParams {
	return &PackageWebContentsGetAllSummariesParams{
		HTTPClient: client,
	}
}

/*
PackageWebContentsGetAllSummariesParams contains all the parameters to send to the API endpoint

	for the package web contents get all summaries operation.

	Typically these are written to a http.Request.
*/
type PackageWebContentsGetAllSummariesParams struct {

	/* ContentTypeIds.

	   A comma separated list of web content type ids. Required if packageIds is not used.
	*/
	ContentTypeIds *string

	/* PackageIds.

	   A comma separated list of package ids.  Required if contentTypeIds is not used.
	*/
	PackageIds *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the package web contents get all summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackageWebContentsGetAllSummariesParams) WithDefaults() *PackageWebContentsGetAllSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the package web contents get all summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackageWebContentsGetAllSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the package web contents get all summaries params
func (o *PackageWebContentsGetAllSummariesParams) WithTimeout(timeout time.Duration) *PackageWebContentsGetAllSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the package web contents get all summaries params
func (o *PackageWebContentsGetAllSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the package web contents get all summaries params
func (o *PackageWebContentsGetAllSummariesParams) WithContext(ctx context.Context) *PackageWebContentsGetAllSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the package web contents get all summaries params
func (o *PackageWebContentsGetAllSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the package web contents get all summaries params
func (o *PackageWebContentsGetAllSummariesParams) WithHTTPClient(client *http.Client) *PackageWebContentsGetAllSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the package web contents get all summaries params
func (o *PackageWebContentsGetAllSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentTypeIds adds the contentTypeIds to the package web contents get all summaries params
func (o *PackageWebContentsGetAllSummariesParams) WithContentTypeIds(contentTypeIds *string) *PackageWebContentsGetAllSummariesParams {
	o.SetContentTypeIds(contentTypeIds)
	return o
}

// SetContentTypeIds adds the contentTypeIds to the package web contents get all summaries params
func (o *PackageWebContentsGetAllSummariesParams) SetContentTypeIds(contentTypeIds *string) {
	o.ContentTypeIds = contentTypeIds
}

// WithPackageIds adds the packageIds to the package web contents get all summaries params
func (o *PackageWebContentsGetAllSummariesParams) WithPackageIds(packageIds *string) *PackageWebContentsGetAllSummariesParams {
	o.SetPackageIds(packageIds)
	return o
}

// SetPackageIds adds the packageIds to the package web contents get all summaries params
func (o *PackageWebContentsGetAllSummariesParams) SetPackageIds(packageIds *string) {
	o.PackageIds = packageIds
}

// WriteToRequest writes these params to a swagger request
func (o *PackageWebContentsGetAllSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentTypeIds != nil {

		// query param contentTypeIds
		var qrContentTypeIds string

		if o.ContentTypeIds != nil {
			qrContentTypeIds = *o.ContentTypeIds
		}
		qContentTypeIds := qrContentTypeIds
		if qContentTypeIds != "" {

			if err := r.SetQueryParam("contentTypeIds", qContentTypeIds); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
