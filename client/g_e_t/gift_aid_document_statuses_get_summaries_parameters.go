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

// NewGiftAidDocumentStatusesGetSummariesParams creates a new GiftAidDocumentStatusesGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGiftAidDocumentStatusesGetSummariesParams() *GiftAidDocumentStatusesGetSummariesParams {
	return &GiftAidDocumentStatusesGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGiftAidDocumentStatusesGetSummariesParamsWithTimeout creates a new GiftAidDocumentStatusesGetSummariesParams object
// with the ability to set a timeout on a request.
func NewGiftAidDocumentStatusesGetSummariesParamsWithTimeout(timeout time.Duration) *GiftAidDocumentStatusesGetSummariesParams {
	return &GiftAidDocumentStatusesGetSummariesParams{
		timeout: timeout,
	}
}

// NewGiftAidDocumentStatusesGetSummariesParamsWithContext creates a new GiftAidDocumentStatusesGetSummariesParams object
// with the ability to set a context for a request.
func NewGiftAidDocumentStatusesGetSummariesParamsWithContext(ctx context.Context) *GiftAidDocumentStatusesGetSummariesParams {
	return &GiftAidDocumentStatusesGetSummariesParams{
		Context: ctx,
	}
}

// NewGiftAidDocumentStatusesGetSummariesParamsWithHTTPClient creates a new GiftAidDocumentStatusesGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGiftAidDocumentStatusesGetSummariesParamsWithHTTPClient(client *http.Client) *GiftAidDocumentStatusesGetSummariesParams {
	return &GiftAidDocumentStatusesGetSummariesParams{
		HTTPClient: client,
	}
}

/*
GiftAidDocumentStatusesGetSummariesParams contains all the parameters to send to the API endpoint

	for the gift aid document statuses get summaries operation.

	Typically these are written to a http.Request.
*/
type GiftAidDocumentStatusesGetSummariesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the gift aid document statuses get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GiftAidDocumentStatusesGetSummariesParams) WithDefaults() *GiftAidDocumentStatusesGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the gift aid document statuses get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GiftAidDocumentStatusesGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the gift aid document statuses get summaries params
func (o *GiftAidDocumentStatusesGetSummariesParams) WithTimeout(timeout time.Duration) *GiftAidDocumentStatusesGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gift aid document statuses get summaries params
func (o *GiftAidDocumentStatusesGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gift aid document statuses get summaries params
func (o *GiftAidDocumentStatusesGetSummariesParams) WithContext(ctx context.Context) *GiftAidDocumentStatusesGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gift aid document statuses get summaries params
func (o *GiftAidDocumentStatusesGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gift aid document statuses get summaries params
func (o *GiftAidDocumentStatusesGetSummariesParams) WithHTTPClient(client *http.Client) *GiftAidDocumentStatusesGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gift aid document statuses get summaries params
func (o *GiftAidDocumentStatusesGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GiftAidDocumentStatusesGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
