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

// NewCartReturnTicketWithSeatParams creates a new CartReturnTicketWithSeatParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCartReturnTicketWithSeatParams() *CartReturnTicketWithSeatParams {
	return &CartReturnTicketWithSeatParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCartReturnTicketWithSeatParamsWithTimeout creates a new CartReturnTicketWithSeatParams object
// with the ability to set a timeout on a request.
func NewCartReturnTicketWithSeatParamsWithTimeout(timeout time.Duration) *CartReturnTicketWithSeatParams {
	return &CartReturnTicketWithSeatParams{
		timeout: timeout,
	}
}

// NewCartReturnTicketWithSeatParamsWithContext creates a new CartReturnTicketWithSeatParams object
// with the ability to set a context for a request.
func NewCartReturnTicketWithSeatParamsWithContext(ctx context.Context) *CartReturnTicketWithSeatParams {
	return &CartReturnTicketWithSeatParams{
		Context: ctx,
	}
}

// NewCartReturnTicketWithSeatParamsWithHTTPClient creates a new CartReturnTicketWithSeatParams object
// with the ability to set a custom HTTPClient for a request.
func NewCartReturnTicketWithSeatParamsWithHTTPClient(client *http.Client) *CartReturnTicketWithSeatParams {
	return &CartReturnTicketWithSeatParams{
		HTTPClient: client,
	}
}

/*
CartReturnTicketWithSeatParams contains all the parameters to send to the API endpoint

	for the cart return ticket with seat operation.

	Typically these are written to a http.Request.
*/
type CartReturnTicketWithSeatParams struct {

	// Request.
	Request *models.ReturnTicketWithSeatRequest

	// SessionKey.
	SessionKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cart return ticket with seat params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartReturnTicketWithSeatParams) WithDefaults() *CartReturnTicketWithSeatParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cart return ticket with seat params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartReturnTicketWithSeatParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cart return ticket with seat params
func (o *CartReturnTicketWithSeatParams) WithTimeout(timeout time.Duration) *CartReturnTicketWithSeatParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cart return ticket with seat params
func (o *CartReturnTicketWithSeatParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cart return ticket with seat params
func (o *CartReturnTicketWithSeatParams) WithContext(ctx context.Context) *CartReturnTicketWithSeatParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cart return ticket with seat params
func (o *CartReturnTicketWithSeatParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cart return ticket with seat params
func (o *CartReturnTicketWithSeatParams) WithHTTPClient(client *http.Client) *CartReturnTicketWithSeatParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cart return ticket with seat params
func (o *CartReturnTicketWithSeatParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the cart return ticket with seat params
func (o *CartReturnTicketWithSeatParams) WithRequest(request *models.ReturnTicketWithSeatRequest) *CartReturnTicketWithSeatParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the cart return ticket with seat params
func (o *CartReturnTicketWithSeatParams) SetRequest(request *models.ReturnTicketWithSeatRequest) {
	o.Request = request
}

// WithSessionKey adds the sessionKey to the cart return ticket with seat params
func (o *CartReturnTicketWithSeatParams) WithSessionKey(sessionKey string) *CartReturnTicketWithSeatParams {
	o.SetSessionKey(sessionKey)
	return o
}

// SetSessionKey adds the sessionKey to the cart return ticket with seat params
func (o *CartReturnTicketWithSeatParams) SetSessionKey(sessionKey string) {
	o.SessionKey = sessionKey
}

// WriteToRequest writes these params to a swagger request
func (o *CartReturnTicketWithSeatParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	// path param sessionKey
	if err := r.SetPathParam("sessionKey", o.SessionKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
