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

// NewCartReturnTicketParams creates a new CartReturnTicketParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCartReturnTicketParams() *CartReturnTicketParams {
	return &CartReturnTicketParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCartReturnTicketParamsWithTimeout creates a new CartReturnTicketParams object
// with the ability to set a timeout on a request.
func NewCartReturnTicketParamsWithTimeout(timeout time.Duration) *CartReturnTicketParams {
	return &CartReturnTicketParams{
		timeout: timeout,
	}
}

// NewCartReturnTicketParamsWithContext creates a new CartReturnTicketParams object
// with the ability to set a context for a request.
func NewCartReturnTicketParamsWithContext(ctx context.Context) *CartReturnTicketParams {
	return &CartReturnTicketParams{
		Context: ctx,
	}
}

// NewCartReturnTicketParamsWithHTTPClient creates a new CartReturnTicketParams object
// with the ability to set a custom HTTPClient for a request.
func NewCartReturnTicketParamsWithHTTPClient(client *http.Client) *CartReturnTicketParams {
	return &CartReturnTicketParams{
		HTTPClient: client,
	}
}

/*
CartReturnTicketParams contains all the parameters to send to the API endpoint

	for the cart return ticket operation.

	Typically these are written to a http.Request.
*/
type CartReturnTicketParams struct {

	// Request.
	Request *models.ReturnTicketWithNumberRequest

	// SessionKey.
	SessionKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cart return ticket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartReturnTicketParams) WithDefaults() *CartReturnTicketParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cart return ticket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartReturnTicketParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cart return ticket params
func (o *CartReturnTicketParams) WithTimeout(timeout time.Duration) *CartReturnTicketParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cart return ticket params
func (o *CartReturnTicketParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cart return ticket params
func (o *CartReturnTicketParams) WithContext(ctx context.Context) *CartReturnTicketParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cart return ticket params
func (o *CartReturnTicketParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cart return ticket params
func (o *CartReturnTicketParams) WithHTTPClient(client *http.Client) *CartReturnTicketParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cart return ticket params
func (o *CartReturnTicketParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the cart return ticket params
func (o *CartReturnTicketParams) WithRequest(request *models.ReturnTicketWithNumberRequest) *CartReturnTicketParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the cart return ticket params
func (o *CartReturnTicketParams) SetRequest(request *models.ReturnTicketWithNumberRequest) {
	o.Request = request
}

// WithSessionKey adds the sessionKey to the cart return ticket params
func (o *CartReturnTicketParams) WithSessionKey(sessionKey string) *CartReturnTicketParams {
	o.SetSessionKey(sessionKey)
	return o
}

// SetSessionKey adds the sessionKey to the cart return ticket params
func (o *CartReturnTicketParams) SetSessionKey(sessionKey string) {
	o.SessionKey = sessionKey
}

// WriteToRequest writes these params to a swagger request
func (o *CartReturnTicketParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
