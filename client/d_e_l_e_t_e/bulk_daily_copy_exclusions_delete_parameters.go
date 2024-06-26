// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

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

// NewBulkDailyCopyExclusionsDeleteParams creates a new BulkDailyCopyExclusionsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBulkDailyCopyExclusionsDeleteParams() *BulkDailyCopyExclusionsDeleteParams {
	return &BulkDailyCopyExclusionsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBulkDailyCopyExclusionsDeleteParamsWithTimeout creates a new BulkDailyCopyExclusionsDeleteParams object
// with the ability to set a timeout on a request.
func NewBulkDailyCopyExclusionsDeleteParamsWithTimeout(timeout time.Duration) *BulkDailyCopyExclusionsDeleteParams {
	return &BulkDailyCopyExclusionsDeleteParams{
		timeout: timeout,
	}
}

// NewBulkDailyCopyExclusionsDeleteParamsWithContext creates a new BulkDailyCopyExclusionsDeleteParams object
// with the ability to set a context for a request.
func NewBulkDailyCopyExclusionsDeleteParamsWithContext(ctx context.Context) *BulkDailyCopyExclusionsDeleteParams {
	return &BulkDailyCopyExclusionsDeleteParams{
		Context: ctx,
	}
}

// NewBulkDailyCopyExclusionsDeleteParamsWithHTTPClient creates a new BulkDailyCopyExclusionsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewBulkDailyCopyExclusionsDeleteParamsWithHTTPClient(client *http.Client) *BulkDailyCopyExclusionsDeleteParams {
	return &BulkDailyCopyExclusionsDeleteParams{
		HTTPClient: client,
	}
}

/*
BulkDailyCopyExclusionsDeleteParams contains all the parameters to send to the API endpoint

	for the bulk daily copy exclusions delete operation.

	Typically these are written to a http.Request.
*/
type BulkDailyCopyExclusionsDeleteParams struct {

	// BulkDailyCopyExclusionID.
	BulkDailyCopyExclusionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bulk daily copy exclusions delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkDailyCopyExclusionsDeleteParams) WithDefaults() *BulkDailyCopyExclusionsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bulk daily copy exclusions delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkDailyCopyExclusionsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the bulk daily copy exclusions delete params
func (o *BulkDailyCopyExclusionsDeleteParams) WithTimeout(timeout time.Duration) *BulkDailyCopyExclusionsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bulk daily copy exclusions delete params
func (o *BulkDailyCopyExclusionsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bulk daily copy exclusions delete params
func (o *BulkDailyCopyExclusionsDeleteParams) WithContext(ctx context.Context) *BulkDailyCopyExclusionsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bulk daily copy exclusions delete params
func (o *BulkDailyCopyExclusionsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bulk daily copy exclusions delete params
func (o *BulkDailyCopyExclusionsDeleteParams) WithHTTPClient(client *http.Client) *BulkDailyCopyExclusionsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bulk daily copy exclusions delete params
func (o *BulkDailyCopyExclusionsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBulkDailyCopyExclusionID adds the bulkDailyCopyExclusionID to the bulk daily copy exclusions delete params
func (o *BulkDailyCopyExclusionsDeleteParams) WithBulkDailyCopyExclusionID(bulkDailyCopyExclusionID string) *BulkDailyCopyExclusionsDeleteParams {
	o.SetBulkDailyCopyExclusionID(bulkDailyCopyExclusionID)
	return o
}

// SetBulkDailyCopyExclusionID adds the bulkDailyCopyExclusionId to the bulk daily copy exclusions delete params
func (o *BulkDailyCopyExclusionsDeleteParams) SetBulkDailyCopyExclusionID(bulkDailyCopyExclusionID string) {
	o.BulkDailyCopyExclusionID = bulkDailyCopyExclusionID
}

// WriteToRequest writes these params to a swagger request
func (o *BulkDailyCopyExclusionsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bulkDailyCopyExclusionId
	if err := r.SetPathParam("bulkDailyCopyExclusionId", o.BulkDailyCopyExclusionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
