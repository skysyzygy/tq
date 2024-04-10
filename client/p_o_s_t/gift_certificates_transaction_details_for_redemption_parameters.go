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

// NewGiftCertificatesTransactionDetailsForRedemptionParams creates a new GiftCertificatesTransactionDetailsForRedemptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGiftCertificatesTransactionDetailsForRedemptionParams() *GiftCertificatesTransactionDetailsForRedemptionParams {
	return &GiftCertificatesTransactionDetailsForRedemptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGiftCertificatesTransactionDetailsForRedemptionParamsWithTimeout creates a new GiftCertificatesTransactionDetailsForRedemptionParams object
// with the ability to set a timeout on a request.
func NewGiftCertificatesTransactionDetailsForRedemptionParamsWithTimeout(timeout time.Duration) *GiftCertificatesTransactionDetailsForRedemptionParams {
	return &GiftCertificatesTransactionDetailsForRedemptionParams{
		timeout: timeout,
	}
}

// NewGiftCertificatesTransactionDetailsForRedemptionParamsWithContext creates a new GiftCertificatesTransactionDetailsForRedemptionParams object
// with the ability to set a context for a request.
func NewGiftCertificatesTransactionDetailsForRedemptionParamsWithContext(ctx context.Context) *GiftCertificatesTransactionDetailsForRedemptionParams {
	return &GiftCertificatesTransactionDetailsForRedemptionParams{
		Context: ctx,
	}
}

// NewGiftCertificatesTransactionDetailsForRedemptionParamsWithHTTPClient creates a new GiftCertificatesTransactionDetailsForRedemptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGiftCertificatesTransactionDetailsForRedemptionParamsWithHTTPClient(client *http.Client) *GiftCertificatesTransactionDetailsForRedemptionParams {
	return &GiftCertificatesTransactionDetailsForRedemptionParams{
		HTTPClient: client,
	}
}

/*
GiftCertificatesTransactionDetailsForRedemptionParams contains all the parameters to send to the API endpoint

	for the gift certificates transaction details for redemption operation.

	Typically these are written to a http.Request.
*/
type GiftCertificatesTransactionDetailsForRedemptionParams struct {

	// GiftCertificateRedemptionRequest.
	GiftCertificateRedemptionRequest *models.GiftCertificateRedemptionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the gift certificates transaction details for redemption params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GiftCertificatesTransactionDetailsForRedemptionParams) WithDefaults() *GiftCertificatesTransactionDetailsForRedemptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the gift certificates transaction details for redemption params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GiftCertificatesTransactionDetailsForRedemptionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the gift certificates transaction details for redemption params
func (o *GiftCertificatesTransactionDetailsForRedemptionParams) WithTimeout(timeout time.Duration) *GiftCertificatesTransactionDetailsForRedemptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gift certificates transaction details for redemption params
func (o *GiftCertificatesTransactionDetailsForRedemptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gift certificates transaction details for redemption params
func (o *GiftCertificatesTransactionDetailsForRedemptionParams) WithContext(ctx context.Context) *GiftCertificatesTransactionDetailsForRedemptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gift certificates transaction details for redemption params
func (o *GiftCertificatesTransactionDetailsForRedemptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gift certificates transaction details for redemption params
func (o *GiftCertificatesTransactionDetailsForRedemptionParams) WithHTTPClient(client *http.Client) *GiftCertificatesTransactionDetailsForRedemptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gift certificates transaction details for redemption params
func (o *GiftCertificatesTransactionDetailsForRedemptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGiftCertificateRedemptionRequest adds the giftCertificateRedemptionRequest to the gift certificates transaction details for redemption params
func (o *GiftCertificatesTransactionDetailsForRedemptionParams) WithGiftCertificateRedemptionRequest(giftCertificateRedemptionRequest *models.GiftCertificateRedemptionRequest) *GiftCertificatesTransactionDetailsForRedemptionParams {
	o.SetGiftCertificateRedemptionRequest(giftCertificateRedemptionRequest)
	return o
}

// SetGiftCertificateRedemptionRequest adds the giftCertificateRedemptionRequest to the gift certificates transaction details for redemption params
func (o *GiftCertificatesTransactionDetailsForRedemptionParams) SetGiftCertificateRedemptionRequest(giftCertificateRedemptionRequest *models.GiftCertificateRedemptionRequest) {
	o.GiftCertificateRedemptionRequest = giftCertificateRedemptionRequest
}

// WriteToRequest writes these params to a swagger request
func (o *GiftCertificatesTransactionDetailsForRedemptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.GiftCertificateRedemptionRequest != nil {
		if err := r.SetBodyParam(o.GiftCertificateRedemptionRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
