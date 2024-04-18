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

// NewConstituentsGetPrimaryHouseholdParams creates a new ConstituentsGetPrimaryHouseholdParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConstituentsGetPrimaryHouseholdParams() *ConstituentsGetPrimaryHouseholdParams {
	return &ConstituentsGetPrimaryHouseholdParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConstituentsGetPrimaryHouseholdParamsWithTimeout creates a new ConstituentsGetPrimaryHouseholdParams object
// with the ability to set a timeout on a request.
func NewConstituentsGetPrimaryHouseholdParamsWithTimeout(timeout time.Duration) *ConstituentsGetPrimaryHouseholdParams {
	return &ConstituentsGetPrimaryHouseholdParams{
		timeout: timeout,
	}
}

// NewConstituentsGetPrimaryHouseholdParamsWithContext creates a new ConstituentsGetPrimaryHouseholdParams object
// with the ability to set a context for a request.
func NewConstituentsGetPrimaryHouseholdParamsWithContext(ctx context.Context) *ConstituentsGetPrimaryHouseholdParams {
	return &ConstituentsGetPrimaryHouseholdParams{
		Context: ctx,
	}
}

// NewConstituentsGetPrimaryHouseholdParamsWithHTTPClient creates a new ConstituentsGetPrimaryHouseholdParams object
// with the ability to set a custom HTTPClient for a request.
func NewConstituentsGetPrimaryHouseholdParamsWithHTTPClient(client *http.Client) *ConstituentsGetPrimaryHouseholdParams {
	return &ConstituentsGetPrimaryHouseholdParams{
		HTTPClient: client,
	}
}

/*
ConstituentsGetPrimaryHouseholdParams contains all the parameters to send to the API endpoint

	for the constituents get primary household operation.

	Typically these are written to a http.Request.
*/
type ConstituentsGetPrimaryHouseholdParams struct {

	// ConstituentID.
	ConstituentID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the constituents get primary household params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConstituentsGetPrimaryHouseholdParams) WithDefaults() *ConstituentsGetPrimaryHouseholdParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the constituents get primary household params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConstituentsGetPrimaryHouseholdParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the constituents get primary household params
func (o *ConstituentsGetPrimaryHouseholdParams) WithTimeout(timeout time.Duration) *ConstituentsGetPrimaryHouseholdParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the constituents get primary household params
func (o *ConstituentsGetPrimaryHouseholdParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the constituents get primary household params
func (o *ConstituentsGetPrimaryHouseholdParams) WithContext(ctx context.Context) *ConstituentsGetPrimaryHouseholdParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the constituents get primary household params
func (o *ConstituentsGetPrimaryHouseholdParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the constituents get primary household params
func (o *ConstituentsGetPrimaryHouseholdParams) WithHTTPClient(client *http.Client) *ConstituentsGetPrimaryHouseholdParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the constituents get primary household params
func (o *ConstituentsGetPrimaryHouseholdParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConstituentID adds the constituentID to the constituents get primary household params
func (o *ConstituentsGetPrimaryHouseholdParams) WithConstituentID(constituentID *string) *ConstituentsGetPrimaryHouseholdParams {
	o.SetConstituentID(constituentID)
	return o
}

// SetConstituentID adds the constituentId to the constituents get primary household params
func (o *ConstituentsGetPrimaryHouseholdParams) SetConstituentID(constituentID *string) {
	o.ConstituentID = constituentID
}

// WriteToRequest writes these params to a swagger request
func (o *ConstituentsGetPrimaryHouseholdParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}