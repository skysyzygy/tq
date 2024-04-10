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

// NewPhonesGetAllParams creates a new PhonesGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPhonesGetAllParams() *PhonesGetAllParams {
	return &PhonesGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPhonesGetAllParamsWithTimeout creates a new PhonesGetAllParams object
// with the ability to set a timeout on a request.
func NewPhonesGetAllParamsWithTimeout(timeout time.Duration) *PhonesGetAllParams {
	return &PhonesGetAllParams{
		timeout: timeout,
	}
}

// NewPhonesGetAllParamsWithContext creates a new PhonesGetAllParams object
// with the ability to set a context for a request.
func NewPhonesGetAllParamsWithContext(ctx context.Context) *PhonesGetAllParams {
	return &PhonesGetAllParams{
		Context: ctx,
	}
}

// NewPhonesGetAllParamsWithHTTPClient creates a new PhonesGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewPhonesGetAllParamsWithHTTPClient(client *http.Client) *PhonesGetAllParams {
	return &PhonesGetAllParams{
		HTTPClient: client,
	}
}

/*
PhonesGetAllParams contains all the parameters to send to the API endpoint

	for the phones get all operation.

	Typically these are written to a http.Request.
*/
type PhonesGetAllParams struct {

	// AddressID.
	AddressID *string

	/* ConstituentID.

	   Limit results by constituent
	*/
	ConstituentID *string

	/* IncludeAffiliations.

	   Include all of the constituent's affiliates in the results (default: true)
	*/
	IncludeAffiliations *string

	// PrimaryOnly.
	PrimaryOnly *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the phones get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PhonesGetAllParams) WithDefaults() *PhonesGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the phones get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PhonesGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the phones get all params
func (o *PhonesGetAllParams) WithTimeout(timeout time.Duration) *PhonesGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the phones get all params
func (o *PhonesGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the phones get all params
func (o *PhonesGetAllParams) WithContext(ctx context.Context) *PhonesGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the phones get all params
func (o *PhonesGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the phones get all params
func (o *PhonesGetAllParams) WithHTTPClient(client *http.Client) *PhonesGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the phones get all params
func (o *PhonesGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddressID adds the addressID to the phones get all params
func (o *PhonesGetAllParams) WithAddressID(addressID *string) *PhonesGetAllParams {
	o.SetAddressID(addressID)
	return o
}

// SetAddressID adds the addressId to the phones get all params
func (o *PhonesGetAllParams) SetAddressID(addressID *string) {
	o.AddressID = addressID
}

// WithConstituentID adds the constituentID to the phones get all params
func (o *PhonesGetAllParams) WithConstituentID(constituentID *string) *PhonesGetAllParams {
	o.SetConstituentID(constituentID)
	return o
}

// SetConstituentID adds the constituentId to the phones get all params
func (o *PhonesGetAllParams) SetConstituentID(constituentID *string) {
	o.ConstituentID = constituentID
}

// WithIncludeAffiliations adds the includeAffiliations to the phones get all params
func (o *PhonesGetAllParams) WithIncludeAffiliations(includeAffiliations *string) *PhonesGetAllParams {
	o.SetIncludeAffiliations(includeAffiliations)
	return o
}

// SetIncludeAffiliations adds the includeAffiliations to the phones get all params
func (o *PhonesGetAllParams) SetIncludeAffiliations(includeAffiliations *string) {
	o.IncludeAffiliations = includeAffiliations
}

// WithPrimaryOnly adds the primaryOnly to the phones get all params
func (o *PhonesGetAllParams) WithPrimaryOnly(primaryOnly *string) *PhonesGetAllParams {
	o.SetPrimaryOnly(primaryOnly)
	return o
}

// SetPrimaryOnly adds the primaryOnly to the phones get all params
func (o *PhonesGetAllParams) SetPrimaryOnly(primaryOnly *string) {
	o.PrimaryOnly = primaryOnly
}

// WriteToRequest writes these params to a swagger request
func (o *PhonesGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AddressID != nil {

		// query param addressId
		var qrAddressID string

		if o.AddressID != nil {
			qrAddressID = *o.AddressID
		}
		qAddressID := qrAddressID
		if qAddressID != "" {

			if err := r.SetQueryParam("addressId", qAddressID); err != nil {
				return err
			}
		}
	}

	if o.ConstituentID != nil {

		// query param constituentId
		var qrConstituentID string

		if o.ConstituentID != nil {
			qrConstituentID = *o.ConstituentID
		}
		qConstituentID := qrConstituentID
		if qConstituentID != "" {

			if err := r.SetQueryParam("constituentId", qConstituentID); err != nil {
				return err
			}
		}
	}

	if o.IncludeAffiliations != nil {

		// query param includeAffiliations
		var qrIncludeAffiliations string

		if o.IncludeAffiliations != nil {
			qrIncludeAffiliations = *o.IncludeAffiliations
		}
		qIncludeAffiliations := qrIncludeAffiliations
		if qIncludeAffiliations != "" {

			if err := r.SetQueryParam("includeAffiliations", qIncludeAffiliations); err != nil {
				return err
			}
		}
	}

	if o.PrimaryOnly != nil {

		// query param primaryOnly
		var qrPrimaryOnly string

		if o.PrimaryOnly != nil {
			qrPrimaryOnly = *o.PrimaryOnly
		}
		qPrimaryOnly := qrPrimaryOnly
		if qPrimaryOnly != "" {

			if err := r.SetQueryParam("primaryOnly", qPrimaryOnly); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
