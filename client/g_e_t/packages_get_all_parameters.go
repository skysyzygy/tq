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

// NewPackagesGetAllParams creates a new PackagesGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPackagesGetAllParams() *PackagesGetAllParams {
	return &PackagesGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPackagesGetAllParamsWithTimeout creates a new PackagesGetAllParams object
// with the ability to set a timeout on a request.
func NewPackagesGetAllParamsWithTimeout(timeout time.Duration) *PackagesGetAllParams {
	return &PackagesGetAllParams{
		timeout: timeout,
	}
}

// NewPackagesGetAllParamsWithContext creates a new PackagesGetAllParams object
// with the ability to set a context for a request.
func NewPackagesGetAllParamsWithContext(ctx context.Context) *PackagesGetAllParams {
	return &PackagesGetAllParams{
		Context: ctx,
	}
}

// NewPackagesGetAllParamsWithHTTPClient creates a new PackagesGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewPackagesGetAllParamsWithHTTPClient(client *http.Client) *PackagesGetAllParams {
	return &PackagesGetAllParams{
		HTTPClient: client,
	}
}

/*
PackagesGetAllParams contains all the parameters to send to the API endpoint

	for the packages get all operation.

	Typically these are written to a http.Request.
*/
type PackagesGetAllParams struct {

	/* SeasonID.

	   A single season Id to filter the packages
	*/
	SeasonID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the packages get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackagesGetAllParams) WithDefaults() *PackagesGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the packages get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackagesGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the packages get all params
func (o *PackagesGetAllParams) WithTimeout(timeout time.Duration) *PackagesGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the packages get all params
func (o *PackagesGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the packages get all params
func (o *PackagesGetAllParams) WithContext(ctx context.Context) *PackagesGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the packages get all params
func (o *PackagesGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the packages get all params
func (o *PackagesGetAllParams) WithHTTPClient(client *http.Client) *PackagesGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the packages get all params
func (o *PackagesGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSeasonID adds the seasonID to the packages get all params
func (o *PackagesGetAllParams) WithSeasonID(seasonID *string) *PackagesGetAllParams {
	o.SetSeasonID(seasonID)
	return o
}

// SetSeasonID adds the seasonId to the packages get all params
func (o *PackagesGetAllParams) SetSeasonID(seasonID *string) {
	o.SeasonID = seasonID
}

// WriteToRequest writes these params to a swagger request
func (o *PackagesGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SeasonID != nil {

		// query param seasonId
		var qrSeasonID string

		if o.SeasonID != nil {
			qrSeasonID = *o.SeasonID
		}
		qSeasonID := qrSeasonID
		if qSeasonID != "" {

			if err := r.SetQueryParam("seasonId", qSeasonID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
