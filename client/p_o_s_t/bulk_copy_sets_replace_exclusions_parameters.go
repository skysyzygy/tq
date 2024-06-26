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

// NewBulkCopySetsReplaceExclusionsParams creates a new BulkCopySetsReplaceExclusionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBulkCopySetsReplaceExclusionsParams() *BulkCopySetsReplaceExclusionsParams {
	return &BulkCopySetsReplaceExclusionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBulkCopySetsReplaceExclusionsParamsWithTimeout creates a new BulkCopySetsReplaceExclusionsParams object
// with the ability to set a timeout on a request.
func NewBulkCopySetsReplaceExclusionsParamsWithTimeout(timeout time.Duration) *BulkCopySetsReplaceExclusionsParams {
	return &BulkCopySetsReplaceExclusionsParams{
		timeout: timeout,
	}
}

// NewBulkCopySetsReplaceExclusionsParamsWithContext creates a new BulkCopySetsReplaceExclusionsParams object
// with the ability to set a context for a request.
func NewBulkCopySetsReplaceExclusionsParamsWithContext(ctx context.Context) *BulkCopySetsReplaceExclusionsParams {
	return &BulkCopySetsReplaceExclusionsParams{
		Context: ctx,
	}
}

// NewBulkCopySetsReplaceExclusionsParamsWithHTTPClient creates a new BulkCopySetsReplaceExclusionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewBulkCopySetsReplaceExclusionsParamsWithHTTPClient(client *http.Client) *BulkCopySetsReplaceExclusionsParams {
	return &BulkCopySetsReplaceExclusionsParams{
		HTTPClient: client,
	}
}

/*
BulkCopySetsReplaceExclusionsParams contains all the parameters to send to the API endpoint

	for the bulk copy sets replace exclusions operation.

	Typically these are written to a http.Request.
*/
type BulkCopySetsReplaceExclusionsParams struct {

	// BulkCopySetID.
	BulkCopySetID string

	// BulkDailyCopyExclusions.
	BulkDailyCopyExclusions []*models.BulkDailyCopyExclusion

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bulk copy sets replace exclusions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkCopySetsReplaceExclusionsParams) WithDefaults() *BulkCopySetsReplaceExclusionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bulk copy sets replace exclusions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkCopySetsReplaceExclusionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the bulk copy sets replace exclusions params
func (o *BulkCopySetsReplaceExclusionsParams) WithTimeout(timeout time.Duration) *BulkCopySetsReplaceExclusionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bulk copy sets replace exclusions params
func (o *BulkCopySetsReplaceExclusionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bulk copy sets replace exclusions params
func (o *BulkCopySetsReplaceExclusionsParams) WithContext(ctx context.Context) *BulkCopySetsReplaceExclusionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bulk copy sets replace exclusions params
func (o *BulkCopySetsReplaceExclusionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bulk copy sets replace exclusions params
func (o *BulkCopySetsReplaceExclusionsParams) WithHTTPClient(client *http.Client) *BulkCopySetsReplaceExclusionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bulk copy sets replace exclusions params
func (o *BulkCopySetsReplaceExclusionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBulkCopySetID adds the bulkCopySetID to the bulk copy sets replace exclusions params
func (o *BulkCopySetsReplaceExclusionsParams) WithBulkCopySetID(bulkCopySetID string) *BulkCopySetsReplaceExclusionsParams {
	o.SetBulkCopySetID(bulkCopySetID)
	return o
}

// SetBulkCopySetID adds the bulkCopySetId to the bulk copy sets replace exclusions params
func (o *BulkCopySetsReplaceExclusionsParams) SetBulkCopySetID(bulkCopySetID string) {
	o.BulkCopySetID = bulkCopySetID
}

// WithBulkDailyCopyExclusions adds the bulkDailyCopyExclusions to the bulk copy sets replace exclusions params
func (o *BulkCopySetsReplaceExclusionsParams) WithBulkDailyCopyExclusions(bulkDailyCopyExclusions []*models.BulkDailyCopyExclusion) *BulkCopySetsReplaceExclusionsParams {
	o.SetBulkDailyCopyExclusions(bulkDailyCopyExclusions)
	return o
}

// SetBulkDailyCopyExclusions adds the bulkDailyCopyExclusions to the bulk copy sets replace exclusions params
func (o *BulkCopySetsReplaceExclusionsParams) SetBulkDailyCopyExclusions(bulkDailyCopyExclusions []*models.BulkDailyCopyExclusion) {
	o.BulkDailyCopyExclusions = bulkDailyCopyExclusions
}

// WriteToRequest writes these params to a swagger request
func (o *BulkCopySetsReplaceExclusionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bulkCopySetId
	if err := r.SetPathParam("bulkCopySetId", o.BulkCopySetID); err != nil {
		return err
	}
	if o.BulkDailyCopyExclusions != nil {
		if err := r.SetBodyParam(o.BulkDailyCopyExclusions); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
