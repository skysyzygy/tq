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

// NewBookingTemplatesGetSummariesParams creates a new BookingTemplatesGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBookingTemplatesGetSummariesParams() *BookingTemplatesGetSummariesParams {
	return &BookingTemplatesGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBookingTemplatesGetSummariesParamsWithTimeout creates a new BookingTemplatesGetSummariesParams object
// with the ability to set a timeout on a request.
func NewBookingTemplatesGetSummariesParamsWithTimeout(timeout time.Duration) *BookingTemplatesGetSummariesParams {
	return &BookingTemplatesGetSummariesParams{
		timeout: timeout,
	}
}

// NewBookingTemplatesGetSummariesParamsWithContext creates a new BookingTemplatesGetSummariesParams object
// with the ability to set a context for a request.
func NewBookingTemplatesGetSummariesParamsWithContext(ctx context.Context) *BookingTemplatesGetSummariesParams {
	return &BookingTemplatesGetSummariesParams{
		Context: ctx,
	}
}

// NewBookingTemplatesGetSummariesParamsWithHTTPClient creates a new BookingTemplatesGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewBookingTemplatesGetSummariesParamsWithHTTPClient(client *http.Client) *BookingTemplatesGetSummariesParams {
	return &BookingTemplatesGetSummariesParams{
		HTTPClient: client,
	}
}

/*
BookingTemplatesGetSummariesParams contains all the parameters to send to the API endpoint

	for the booking templates get summaries operation.

	Typically these are written to a http.Request.
*/
type BookingTemplatesGetSummariesParams struct {

	// CategoryID.
	CategoryID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the booking templates get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BookingTemplatesGetSummariesParams) WithDefaults() *BookingTemplatesGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the booking templates get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BookingTemplatesGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the booking templates get summaries params
func (o *BookingTemplatesGetSummariesParams) WithTimeout(timeout time.Duration) *BookingTemplatesGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the booking templates get summaries params
func (o *BookingTemplatesGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the booking templates get summaries params
func (o *BookingTemplatesGetSummariesParams) WithContext(ctx context.Context) *BookingTemplatesGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the booking templates get summaries params
func (o *BookingTemplatesGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the booking templates get summaries params
func (o *BookingTemplatesGetSummariesParams) WithHTTPClient(client *http.Client) *BookingTemplatesGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the booking templates get summaries params
func (o *BookingTemplatesGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategoryID adds the categoryID to the booking templates get summaries params
func (o *BookingTemplatesGetSummariesParams) WithCategoryID(categoryID *string) *BookingTemplatesGetSummariesParams {
	o.SetCategoryID(categoryID)
	return o
}

// SetCategoryID adds the categoryId to the booking templates get summaries params
func (o *BookingTemplatesGetSummariesParams) SetCategoryID(categoryID *string) {
	o.CategoryID = categoryID
}

// WriteToRequest writes these params to a swagger request
func (o *BookingTemplatesGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CategoryID != nil {

		// query param categoryId
		var qrCategoryID string

		if o.CategoryID != nil {
			qrCategoryID = *o.CategoryID
		}
		qCategoryID := qrCategoryID
		if qCategoryID != "" {

			if err := r.SetQueryParam("categoryId", qCategoryID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
