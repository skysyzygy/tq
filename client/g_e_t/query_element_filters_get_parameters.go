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

// NewQueryElementFiltersGetParams creates a new QueryElementFiltersGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryElementFiltersGetParams() *QueryElementFiltersGetParams {
	return &QueryElementFiltersGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryElementFiltersGetParamsWithTimeout creates a new QueryElementFiltersGetParams object
// with the ability to set a timeout on a request.
func NewQueryElementFiltersGetParamsWithTimeout(timeout time.Duration) *QueryElementFiltersGetParams {
	return &QueryElementFiltersGetParams{
		timeout: timeout,
	}
}

// NewQueryElementFiltersGetParamsWithContext creates a new QueryElementFiltersGetParams object
// with the ability to set a context for a request.
func NewQueryElementFiltersGetParamsWithContext(ctx context.Context) *QueryElementFiltersGetParams {
	return &QueryElementFiltersGetParams{
		Context: ctx,
	}
}

// NewQueryElementFiltersGetParamsWithHTTPClient creates a new QueryElementFiltersGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryElementFiltersGetParamsWithHTTPClient(client *http.Client) *QueryElementFiltersGetParams {
	return &QueryElementFiltersGetParams{
		HTTPClient: client,
	}
}

/*
QueryElementFiltersGetParams contains all the parameters to send to the API endpoint

	for the query element filters get operation.

	Typically these are written to a http.Request.
*/
type QueryElementFiltersGetParams struct {

	// QueryElementFilterID.
	QueryElementFilterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query element filters get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryElementFiltersGetParams) WithDefaults() *QueryElementFiltersGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query element filters get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryElementFiltersGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query element filters get params
func (o *QueryElementFiltersGetParams) WithTimeout(timeout time.Duration) *QueryElementFiltersGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query element filters get params
func (o *QueryElementFiltersGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query element filters get params
func (o *QueryElementFiltersGetParams) WithContext(ctx context.Context) *QueryElementFiltersGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query element filters get params
func (o *QueryElementFiltersGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query element filters get params
func (o *QueryElementFiltersGetParams) WithHTTPClient(client *http.Client) *QueryElementFiltersGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query element filters get params
func (o *QueryElementFiltersGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQueryElementFilterID adds the queryElementFilterID to the query element filters get params
func (o *QueryElementFiltersGetParams) WithQueryElementFilterID(queryElementFilterID string) *QueryElementFiltersGetParams {
	o.SetQueryElementFilterID(queryElementFilterID)
	return o
}

// SetQueryElementFilterID adds the queryElementFilterId to the query element filters get params
func (o *QueryElementFiltersGetParams) SetQueryElementFilterID(queryElementFilterID string) {
	o.QueryElementFilterID = queryElementFilterID
}

// WriteToRequest writes these params to a swagger request
func (o *QueryElementFiltersGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param queryElementFilterId
	if err := r.SetPathParam("queryElementFilterId", o.QueryElementFilterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}