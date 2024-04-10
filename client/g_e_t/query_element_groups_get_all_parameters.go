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

// NewQueryElementGroupsGetAllParams creates a new QueryElementGroupsGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryElementGroupsGetAllParams() *QueryElementGroupsGetAllParams {
	return &QueryElementGroupsGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryElementGroupsGetAllParamsWithTimeout creates a new QueryElementGroupsGetAllParams object
// with the ability to set a timeout on a request.
func NewQueryElementGroupsGetAllParamsWithTimeout(timeout time.Duration) *QueryElementGroupsGetAllParams {
	return &QueryElementGroupsGetAllParams{
		timeout: timeout,
	}
}

// NewQueryElementGroupsGetAllParamsWithContext creates a new QueryElementGroupsGetAllParams object
// with the ability to set a context for a request.
func NewQueryElementGroupsGetAllParamsWithContext(ctx context.Context) *QueryElementGroupsGetAllParams {
	return &QueryElementGroupsGetAllParams{
		Context: ctx,
	}
}

// NewQueryElementGroupsGetAllParamsWithHTTPClient creates a new QueryElementGroupsGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryElementGroupsGetAllParamsWithHTTPClient(client *http.Client) *QueryElementGroupsGetAllParams {
	return &QueryElementGroupsGetAllParams{
		HTTPClient: client,
	}
}

/*
QueryElementGroupsGetAllParams contains all the parameters to send to the API endpoint

	for the query element groups get all operation.

	Typically these are written to a http.Request.
*/
type QueryElementGroupsGetAllParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query element groups get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryElementGroupsGetAllParams) WithDefaults() *QueryElementGroupsGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query element groups get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryElementGroupsGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query element groups get all params
func (o *QueryElementGroupsGetAllParams) WithTimeout(timeout time.Duration) *QueryElementGroupsGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query element groups get all params
func (o *QueryElementGroupsGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query element groups get all params
func (o *QueryElementGroupsGetAllParams) WithContext(ctx context.Context) *QueryElementGroupsGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query element groups get all params
func (o *QueryElementGroupsGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query element groups get all params
func (o *QueryElementGroupsGetAllParams) WithHTTPClient(client *http.Client) *QueryElementGroupsGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query element groups get all params
func (o *QueryElementGroupsGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *QueryElementGroupsGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
