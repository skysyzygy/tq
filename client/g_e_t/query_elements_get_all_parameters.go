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

// NewQueryElementsGetAllParams creates a new QueryElementsGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryElementsGetAllParams() *QueryElementsGetAllParams {
	return &QueryElementsGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryElementsGetAllParamsWithTimeout creates a new QueryElementsGetAllParams object
// with the ability to set a timeout on a request.
func NewQueryElementsGetAllParamsWithTimeout(timeout time.Duration) *QueryElementsGetAllParams {
	return &QueryElementsGetAllParams{
		timeout: timeout,
	}
}

// NewQueryElementsGetAllParamsWithContext creates a new QueryElementsGetAllParams object
// with the ability to set a context for a request.
func NewQueryElementsGetAllParamsWithContext(ctx context.Context) *QueryElementsGetAllParams {
	return &QueryElementsGetAllParams{
		Context: ctx,
	}
}

// NewQueryElementsGetAllParamsWithHTTPClient creates a new QueryElementsGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryElementsGetAllParamsWithHTTPClient(client *http.Client) *QueryElementsGetAllParams {
	return &QueryElementsGetAllParams{
		HTTPClient: client,
	}
}

/*
QueryElementsGetAllParams contains all the parameters to send to the API endpoint

	for the query elements get all operation.

	Typically these are written to a http.Request.
*/
type QueryElementsGetAllParams struct {

	// ActiveOnly.
	ActiveOnly *string

	// GroupIds.
	GroupIds *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query elements get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryElementsGetAllParams) WithDefaults() *QueryElementsGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query elements get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryElementsGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query elements get all params
func (o *QueryElementsGetAllParams) WithTimeout(timeout time.Duration) *QueryElementsGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query elements get all params
func (o *QueryElementsGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query elements get all params
func (o *QueryElementsGetAllParams) WithContext(ctx context.Context) *QueryElementsGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query elements get all params
func (o *QueryElementsGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query elements get all params
func (o *QueryElementsGetAllParams) WithHTTPClient(client *http.Client) *QueryElementsGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query elements get all params
func (o *QueryElementsGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActiveOnly adds the activeOnly to the query elements get all params
func (o *QueryElementsGetAllParams) WithActiveOnly(activeOnly *string) *QueryElementsGetAllParams {
	o.SetActiveOnly(activeOnly)
	return o
}

// SetActiveOnly adds the activeOnly to the query elements get all params
func (o *QueryElementsGetAllParams) SetActiveOnly(activeOnly *string) {
	o.ActiveOnly = activeOnly
}

// WithGroupIds adds the groupIds to the query elements get all params
func (o *QueryElementsGetAllParams) WithGroupIds(groupIds *string) *QueryElementsGetAllParams {
	o.SetGroupIds(groupIds)
	return o
}

// SetGroupIds adds the groupIds to the query elements get all params
func (o *QueryElementsGetAllParams) SetGroupIds(groupIds *string) {
	o.GroupIds = groupIds
}

// WriteToRequest writes these params to a swagger request
func (o *QueryElementsGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ActiveOnly != nil {

		// query param activeOnly
		var qrActiveOnly string

		if o.ActiveOnly != nil {
			qrActiveOnly = *o.ActiveOnly
		}
		qActiveOnly := qrActiveOnly
		if qActiveOnly != "" {

			if err := r.SetQueryParam("activeOnly", qActiveOnly); err != nil {
				return err
			}
		}
	}

	if o.GroupIds != nil {

		// query param groupIds
		var qrGroupIds string

		if o.GroupIds != nil {
			qrGroupIds = *o.GroupIds
		}
		qGroupIds := qrGroupIds
		if qGroupIds != "" {

			if err := r.SetQueryParam("groupIds", qGroupIds); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}