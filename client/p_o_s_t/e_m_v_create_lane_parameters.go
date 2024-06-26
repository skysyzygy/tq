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

// NewEMVCreateLaneParams creates a new EMVCreateLaneParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEMVCreateLaneParams() *EMVCreateLaneParams {
	return &EMVCreateLaneParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEMVCreateLaneParamsWithTimeout creates a new EMVCreateLaneParams object
// with the ability to set a timeout on a request.
func NewEMVCreateLaneParamsWithTimeout(timeout time.Duration) *EMVCreateLaneParams {
	return &EMVCreateLaneParams{
		timeout: timeout,
	}
}

// NewEMVCreateLaneParamsWithContext creates a new EMVCreateLaneParams object
// with the ability to set a context for a request.
func NewEMVCreateLaneParamsWithContext(ctx context.Context) *EMVCreateLaneParams {
	return &EMVCreateLaneParams{
		Context: ctx,
	}
}

// NewEMVCreateLaneParamsWithHTTPClient creates a new EMVCreateLaneParams object
// with the ability to set a custom HTTPClient for a request.
func NewEMVCreateLaneParamsWithHTTPClient(client *http.Client) *EMVCreateLaneParams {
	return &EMVCreateLaneParams{
		HTTPClient: client,
	}
}

/*
EMVCreateLaneParams contains all the parameters to send to the API endpoint

	for the e m v create lane operation.

	Typically these are written to a http.Request.
*/
type EMVCreateLaneParams struct {

	// Cert.
	Cert *string

	// Lane.
	Lane *models.LaneCreate

	// Merchant.
	Merchant *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the e m v create lane params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EMVCreateLaneParams) WithDefaults() *EMVCreateLaneParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the e m v create lane params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EMVCreateLaneParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the e m v create lane params
func (o *EMVCreateLaneParams) WithTimeout(timeout time.Duration) *EMVCreateLaneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the e m v create lane params
func (o *EMVCreateLaneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the e m v create lane params
func (o *EMVCreateLaneParams) WithContext(ctx context.Context) *EMVCreateLaneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the e m v create lane params
func (o *EMVCreateLaneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the e m v create lane params
func (o *EMVCreateLaneParams) WithHTTPClient(client *http.Client) *EMVCreateLaneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the e m v create lane params
func (o *EMVCreateLaneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCert adds the cert to the e m v create lane params
func (o *EMVCreateLaneParams) WithCert(cert *string) *EMVCreateLaneParams {
	o.SetCert(cert)
	return o
}

// SetCert adds the cert to the e m v create lane params
func (o *EMVCreateLaneParams) SetCert(cert *string) {
	o.Cert = cert
}

// WithLane adds the lane to the e m v create lane params
func (o *EMVCreateLaneParams) WithLane(lane *models.LaneCreate) *EMVCreateLaneParams {
	o.SetLane(lane)
	return o
}

// SetLane adds the lane to the e m v create lane params
func (o *EMVCreateLaneParams) SetLane(lane *models.LaneCreate) {
	o.Lane = lane
}

// WithMerchant adds the merchant to the e m v create lane params
func (o *EMVCreateLaneParams) WithMerchant(merchant *string) *EMVCreateLaneParams {
	o.SetMerchant(merchant)
	return o
}

// SetMerchant adds the merchant to the e m v create lane params
func (o *EMVCreateLaneParams) SetMerchant(merchant *string) {
	o.Merchant = merchant
}

// WriteToRequest writes these params to a swagger request
func (o *EMVCreateLaneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cert != nil {

		// query param cert
		var qrCert string

		if o.Cert != nil {
			qrCert = *o.Cert
		}
		qCert := qrCert
		if qCert != "" {

			if err := r.SetQueryParam("cert", qCert); err != nil {
				return err
			}
		}
	}
	if o.Lane != nil {
		if err := r.SetBodyParam(o.Lane); err != nil {
			return err
		}
	}

	if o.Merchant != nil {

		// query param merchant
		var qrMerchant string

		if o.Merchant != nil {
			qrMerchant = *o.Merchant
		}
		qMerchant := qrMerchant
		if qMerchant != "" {

			if err := r.SetQueryParam("merchant", qMerchant); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
