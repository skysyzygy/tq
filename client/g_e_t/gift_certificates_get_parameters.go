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

// NewGiftCertificatesGetParams creates a new GiftCertificatesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGiftCertificatesGetParams() *GiftCertificatesGetParams {
	return &GiftCertificatesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGiftCertificatesGetParamsWithTimeout creates a new GiftCertificatesGetParams object
// with the ability to set a timeout on a request.
func NewGiftCertificatesGetParamsWithTimeout(timeout time.Duration) *GiftCertificatesGetParams {
	return &GiftCertificatesGetParams{
		timeout: timeout,
	}
}

// NewGiftCertificatesGetParamsWithContext creates a new GiftCertificatesGetParams object
// with the ability to set a context for a request.
func NewGiftCertificatesGetParamsWithContext(ctx context.Context) *GiftCertificatesGetParams {
	return &GiftCertificatesGetParams{
		Context: ctx,
	}
}

// NewGiftCertificatesGetParamsWithHTTPClient creates a new GiftCertificatesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewGiftCertificatesGetParamsWithHTTPClient(client *http.Client) *GiftCertificatesGetParams {
	return &GiftCertificatesGetParams{
		HTTPClient: client,
	}
}

/*
GiftCertificatesGetParams contains all the parameters to send to the API endpoint

	for the gift certificates get operation.

	Typically these are written to a http.Request.
*/
type GiftCertificatesGetParams struct {

	// GiftCertificateNumber.
	GiftCertificateNumber string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the gift certificates get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GiftCertificatesGetParams) WithDefaults() *GiftCertificatesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the gift certificates get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GiftCertificatesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the gift certificates get params
func (o *GiftCertificatesGetParams) WithTimeout(timeout time.Duration) *GiftCertificatesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gift certificates get params
func (o *GiftCertificatesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gift certificates get params
func (o *GiftCertificatesGetParams) WithContext(ctx context.Context) *GiftCertificatesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gift certificates get params
func (o *GiftCertificatesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gift certificates get params
func (o *GiftCertificatesGetParams) WithHTTPClient(client *http.Client) *GiftCertificatesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gift certificates get params
func (o *GiftCertificatesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGiftCertificateNumber adds the giftCertificateNumber to the gift certificates get params
func (o *GiftCertificatesGetParams) WithGiftCertificateNumber(giftCertificateNumber string) *GiftCertificatesGetParams {
	o.SetGiftCertificateNumber(giftCertificateNumber)
	return o
}

// SetGiftCertificateNumber adds the giftCertificateNumber to the gift certificates get params
func (o *GiftCertificatesGetParams) SetGiftCertificateNumber(giftCertificateNumber string) {
	o.GiftCertificateNumber = giftCertificateNumber
}

// WriteToRequest writes these params to a swagger request
func (o *GiftCertificatesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param giftCertificateNumber
	if err := r.SetPathParam("giftCertificateNumber", o.GiftCertificateNumber); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
