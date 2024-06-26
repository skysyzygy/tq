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

// NewRecordAttendanceRecordTicketParams creates a new RecordAttendanceRecordTicketParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRecordAttendanceRecordTicketParams() *RecordAttendanceRecordTicketParams {
	return &RecordAttendanceRecordTicketParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRecordAttendanceRecordTicketParamsWithTimeout creates a new RecordAttendanceRecordTicketParams object
// with the ability to set a timeout on a request.
func NewRecordAttendanceRecordTicketParamsWithTimeout(timeout time.Duration) *RecordAttendanceRecordTicketParams {
	return &RecordAttendanceRecordTicketParams{
		timeout: timeout,
	}
}

// NewRecordAttendanceRecordTicketParamsWithContext creates a new RecordAttendanceRecordTicketParams object
// with the ability to set a context for a request.
func NewRecordAttendanceRecordTicketParamsWithContext(ctx context.Context) *RecordAttendanceRecordTicketParams {
	return &RecordAttendanceRecordTicketParams{
		Context: ctx,
	}
}

// NewRecordAttendanceRecordTicketParamsWithHTTPClient creates a new RecordAttendanceRecordTicketParams object
// with the ability to set a custom HTTPClient for a request.
func NewRecordAttendanceRecordTicketParamsWithHTTPClient(client *http.Client) *RecordAttendanceRecordTicketParams {
	return &RecordAttendanceRecordTicketParams{
		HTTPClient: client,
	}
}

/*
RecordAttendanceRecordTicketParams contains all the parameters to send to the API endpoint

	for the record attendance record ticket operation.

	Typically these are written to a http.Request.
*/
type RecordAttendanceRecordTicketParams struct {

	// Request.
	Request *models.RecordAttendanceTicketRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the record attendance record ticket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RecordAttendanceRecordTicketParams) WithDefaults() *RecordAttendanceRecordTicketParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the record attendance record ticket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RecordAttendanceRecordTicketParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the record attendance record ticket params
func (o *RecordAttendanceRecordTicketParams) WithTimeout(timeout time.Duration) *RecordAttendanceRecordTicketParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the record attendance record ticket params
func (o *RecordAttendanceRecordTicketParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the record attendance record ticket params
func (o *RecordAttendanceRecordTicketParams) WithContext(ctx context.Context) *RecordAttendanceRecordTicketParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the record attendance record ticket params
func (o *RecordAttendanceRecordTicketParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the record attendance record ticket params
func (o *RecordAttendanceRecordTicketParams) WithHTTPClient(client *http.Client) *RecordAttendanceRecordTicketParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the record attendance record ticket params
func (o *RecordAttendanceRecordTicketParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the record attendance record ticket params
func (o *RecordAttendanceRecordTicketParams) WithRequest(request *models.RecordAttendanceTicketRequest) *RecordAttendanceRecordTicketParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the record attendance record ticket params
func (o *RecordAttendanceRecordTicketParams) SetRequest(request *models.RecordAttendanceTicketRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *RecordAttendanceRecordTicketParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
