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

// NewBookingsAddDocumentParams creates a new BookingsAddDocumentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBookingsAddDocumentParams() *BookingsAddDocumentParams {
	return &BookingsAddDocumentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBookingsAddDocumentParamsWithTimeout creates a new BookingsAddDocumentParams object
// with the ability to set a timeout on a request.
func NewBookingsAddDocumentParamsWithTimeout(timeout time.Duration) *BookingsAddDocumentParams {
	return &BookingsAddDocumentParams{
		timeout: timeout,
	}
}

// NewBookingsAddDocumentParamsWithContext creates a new BookingsAddDocumentParams object
// with the ability to set a context for a request.
func NewBookingsAddDocumentParamsWithContext(ctx context.Context) *BookingsAddDocumentParams {
	return &BookingsAddDocumentParams{
		Context: ctx,
	}
}

// NewBookingsAddDocumentParamsWithHTTPClient creates a new BookingsAddDocumentParams object
// with the ability to set a custom HTTPClient for a request.
func NewBookingsAddDocumentParamsWithHTTPClient(client *http.Client) *BookingsAddDocumentParams {
	return &BookingsAddDocumentParams{
		HTTPClient: client,
	}
}

/*
BookingsAddDocumentParams contains all the parameters to send to the API endpoint

	for the bookings add document operation.

	Typically these are written to a http.Request.
*/
type BookingsAddDocumentParams struct {

	// BookingID.
	BookingID string

	// Document.
	Document *models.Document

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bookings add document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BookingsAddDocumentParams) WithDefaults() *BookingsAddDocumentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bookings add document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BookingsAddDocumentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the bookings add document params
func (o *BookingsAddDocumentParams) WithTimeout(timeout time.Duration) *BookingsAddDocumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bookings add document params
func (o *BookingsAddDocumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bookings add document params
func (o *BookingsAddDocumentParams) WithContext(ctx context.Context) *BookingsAddDocumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bookings add document params
func (o *BookingsAddDocumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bookings add document params
func (o *BookingsAddDocumentParams) WithHTTPClient(client *http.Client) *BookingsAddDocumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bookings add document params
func (o *BookingsAddDocumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBookingID adds the bookingID to the bookings add document params
func (o *BookingsAddDocumentParams) WithBookingID(bookingID string) *BookingsAddDocumentParams {
	o.SetBookingID(bookingID)
	return o
}

// SetBookingID adds the bookingId to the bookings add document params
func (o *BookingsAddDocumentParams) SetBookingID(bookingID string) {
	o.BookingID = bookingID
}

// WithDocument adds the document to the bookings add document params
func (o *BookingsAddDocumentParams) WithDocument(document *models.Document) *BookingsAddDocumentParams {
	o.SetDocument(document)
	return o
}

// SetDocument adds the document to the bookings add document params
func (o *BookingsAddDocumentParams) SetDocument(document *models.Document) {
	o.Document = document
}

// WriteToRequest writes these params to a swagger request
func (o *BookingsAddDocumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bookingId
	if err := r.SetPathParam("bookingId", o.BookingID); err != nil {
		return err
	}
	if o.Document != nil {
		if err := r.SetBodyParam(o.Document); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
