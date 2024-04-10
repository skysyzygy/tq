// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

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

// NewBookingsUpdateParams creates a new BookingsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBookingsUpdateParams() *BookingsUpdateParams {
	return &BookingsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBookingsUpdateParamsWithTimeout creates a new BookingsUpdateParams object
// with the ability to set a timeout on a request.
func NewBookingsUpdateParamsWithTimeout(timeout time.Duration) *BookingsUpdateParams {
	return &BookingsUpdateParams{
		timeout: timeout,
	}
}

// NewBookingsUpdateParamsWithContext creates a new BookingsUpdateParams object
// with the ability to set a context for a request.
func NewBookingsUpdateParamsWithContext(ctx context.Context) *BookingsUpdateParams {
	return &BookingsUpdateParams{
		Context: ctx,
	}
}

// NewBookingsUpdateParamsWithHTTPClient creates a new BookingsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewBookingsUpdateParamsWithHTTPClient(client *http.Client) *BookingsUpdateParams {
	return &BookingsUpdateParams{
		HTTPClient: client,
	}
}

/*
BookingsUpdateParams contains all the parameters to send to the API endpoint

	for the bookings update operation.

	Typically these are written to a http.Request.
*/
type BookingsUpdateParams struct {

	// Booking.
	Booking *models.Booking

	// BookingID.
	BookingID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bookings update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BookingsUpdateParams) WithDefaults() *BookingsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bookings update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BookingsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the bookings update params
func (o *BookingsUpdateParams) WithTimeout(timeout time.Duration) *BookingsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bookings update params
func (o *BookingsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bookings update params
func (o *BookingsUpdateParams) WithContext(ctx context.Context) *BookingsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bookings update params
func (o *BookingsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bookings update params
func (o *BookingsUpdateParams) WithHTTPClient(client *http.Client) *BookingsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bookings update params
func (o *BookingsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBooking adds the booking to the bookings update params
func (o *BookingsUpdateParams) WithBooking(booking *models.Booking) *BookingsUpdateParams {
	o.SetBooking(booking)
	return o
}

// SetBooking adds the booking to the bookings update params
func (o *BookingsUpdateParams) SetBooking(booking *models.Booking) {
	o.Booking = booking
}

// WithBookingID adds the bookingID to the bookings update params
func (o *BookingsUpdateParams) WithBookingID(bookingID string) *BookingsUpdateParams {
	o.SetBookingID(bookingID)
	return o
}

// SetBookingID adds the bookingId to the bookings update params
func (o *BookingsUpdateParams) SetBookingID(bookingID string) {
	o.BookingID = bookingID
}

// WriteToRequest writes these params to a swagger request
func (o *BookingsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Booking != nil {
		if err := r.SetBodyParam(o.Booking); err != nil {
			return err
		}
	}

	// path param bookingId
	if err := r.SetPathParam("bookingId", o.BookingID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
