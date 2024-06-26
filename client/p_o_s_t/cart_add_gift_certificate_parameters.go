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

// NewCartAddGiftCertificateParams creates a new CartAddGiftCertificateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCartAddGiftCertificateParams() *CartAddGiftCertificateParams {
	return &CartAddGiftCertificateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCartAddGiftCertificateParamsWithTimeout creates a new CartAddGiftCertificateParams object
// with the ability to set a timeout on a request.
func NewCartAddGiftCertificateParamsWithTimeout(timeout time.Duration) *CartAddGiftCertificateParams {
	return &CartAddGiftCertificateParams{
		timeout: timeout,
	}
}

// NewCartAddGiftCertificateParamsWithContext creates a new CartAddGiftCertificateParams object
// with the ability to set a context for a request.
func NewCartAddGiftCertificateParamsWithContext(ctx context.Context) *CartAddGiftCertificateParams {
	return &CartAddGiftCertificateParams{
		Context: ctx,
	}
}

// NewCartAddGiftCertificateParamsWithHTTPClient creates a new CartAddGiftCertificateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCartAddGiftCertificateParamsWithHTTPClient(client *http.Client) *CartAddGiftCertificateParams {
	return &CartAddGiftCertificateParams{
		HTTPClient: client,
	}
}

/*
CartAddGiftCertificateParams contains all the parameters to send to the API endpoint

	for the cart add gift certificate operation.

	Typically these are written to a http.Request.
*/
type CartAddGiftCertificateParams struct {

	// Request.
	Request *models.AddGiftCertificateRequest

	// SessionKey.
	SessionKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cart add gift certificate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartAddGiftCertificateParams) WithDefaults() *CartAddGiftCertificateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cart add gift certificate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartAddGiftCertificateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cart add gift certificate params
func (o *CartAddGiftCertificateParams) WithTimeout(timeout time.Duration) *CartAddGiftCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cart add gift certificate params
func (o *CartAddGiftCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cart add gift certificate params
func (o *CartAddGiftCertificateParams) WithContext(ctx context.Context) *CartAddGiftCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cart add gift certificate params
func (o *CartAddGiftCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cart add gift certificate params
func (o *CartAddGiftCertificateParams) WithHTTPClient(client *http.Client) *CartAddGiftCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cart add gift certificate params
func (o *CartAddGiftCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the cart add gift certificate params
func (o *CartAddGiftCertificateParams) WithRequest(request *models.AddGiftCertificateRequest) *CartAddGiftCertificateParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the cart add gift certificate params
func (o *CartAddGiftCertificateParams) SetRequest(request *models.AddGiftCertificateRequest) {
	o.Request = request
}

// WithSessionKey adds the sessionKey to the cart add gift certificate params
func (o *CartAddGiftCertificateParams) WithSessionKey(sessionKey string) *CartAddGiftCertificateParams {
	o.SetSessionKey(sessionKey)
	return o
}

// SetSessionKey adds the sessionKey to the cart add gift certificate params
func (o *CartAddGiftCertificateParams) SetSessionKey(sessionKey string) {
	o.SessionKey = sessionKey
}

// WriteToRequest writes these params to a swagger request
func (o *CartAddGiftCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
