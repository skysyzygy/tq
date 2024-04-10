// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

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

	"github.com/skysyzygy/tq/models"
)

// NewBulkDailyCopyExclusionsCreateParams creates a new BulkDailyCopyExclusionsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBulkDailyCopyExclusionsCreateParams() *BulkDailyCopyExclusionsCreateParams {
	return &BulkDailyCopyExclusionsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBulkDailyCopyExclusionsCreateParamsWithTimeout creates a new BulkDailyCopyExclusionsCreateParams object
// with the ability to set a timeout on a request.
func NewBulkDailyCopyExclusionsCreateParamsWithTimeout(timeout time.Duration) *BulkDailyCopyExclusionsCreateParams {
	return &BulkDailyCopyExclusionsCreateParams{
		timeout: timeout,
	}
}

// NewBulkDailyCopyExclusionsCreateParamsWithContext creates a new BulkDailyCopyExclusionsCreateParams object
// with the ability to set a context for a request.
func NewBulkDailyCopyExclusionsCreateParamsWithContext(ctx context.Context) *BulkDailyCopyExclusionsCreateParams {
	return &BulkDailyCopyExclusionsCreateParams{
		Context: ctx,
	}
}

// NewBulkDailyCopyExclusionsCreateParamsWithHTTPClient creates a new BulkDailyCopyExclusionsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewBulkDailyCopyExclusionsCreateParamsWithHTTPClient(client *http.Client) *BulkDailyCopyExclusionsCreateParams {
	return &BulkDailyCopyExclusionsCreateParams{
		HTTPClient: client,
	}
}

/*
BulkDailyCopyExclusionsCreateParams contains all the parameters to send to the API endpoint

	for the bulk daily copy exclusions create operation.

	Typically these are written to a http.Request.
*/
type BulkDailyCopyExclusionsCreateParams struct {

	// BulkDailyCopyExclusion.
	BulkDailyCopyExclusion *models.BulkDailyCopyExclusion

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bulk daily copy exclusions create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkDailyCopyExclusionsCreateParams) WithDefaults() *BulkDailyCopyExclusionsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bulk daily copy exclusions create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkDailyCopyExclusionsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the bulk daily copy exclusions create params
func (o *BulkDailyCopyExclusionsCreateParams) WithTimeout(timeout time.Duration) *BulkDailyCopyExclusionsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bulk daily copy exclusions create params
func (o *BulkDailyCopyExclusionsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bulk daily copy exclusions create params
func (o *BulkDailyCopyExclusionsCreateParams) WithContext(ctx context.Context) *BulkDailyCopyExclusionsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bulk daily copy exclusions create params
func (o *BulkDailyCopyExclusionsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bulk daily copy exclusions create params
func (o *BulkDailyCopyExclusionsCreateParams) WithHTTPClient(client *http.Client) *BulkDailyCopyExclusionsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bulk daily copy exclusions create params
func (o *BulkDailyCopyExclusionsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBulkDailyCopyExclusion adds the bulkDailyCopyExclusion to the bulk daily copy exclusions create params
func (o *BulkDailyCopyExclusionsCreateParams) WithBulkDailyCopyExclusion(bulkDailyCopyExclusion *models.BulkDailyCopyExclusion) *BulkDailyCopyExclusionsCreateParams {
	o.SetBulkDailyCopyExclusion(bulkDailyCopyExclusion)
	return o
}

// SetBulkDailyCopyExclusion adds the bulkDailyCopyExclusion to the bulk daily copy exclusions create params
func (o *BulkDailyCopyExclusionsCreateParams) SetBulkDailyCopyExclusion(bulkDailyCopyExclusion *models.BulkDailyCopyExclusion) {
	o.BulkDailyCopyExclusion = bulkDailyCopyExclusion
}

// WriteToRequest writes these params to a swagger request
func (o *BulkDailyCopyExclusionsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.BulkDailyCopyExclusion != nil {
		if err := r.SetBodyParam(o.BulkDailyCopyExclusion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
