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

// NewConstituentsConvertIndividualToHouseholdParams creates a new ConstituentsConvertIndividualToHouseholdParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConstituentsConvertIndividualToHouseholdParams() *ConstituentsConvertIndividualToHouseholdParams {
	return &ConstituentsConvertIndividualToHouseholdParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConstituentsConvertIndividualToHouseholdParamsWithTimeout creates a new ConstituentsConvertIndividualToHouseholdParams object
// with the ability to set a timeout on a request.
func NewConstituentsConvertIndividualToHouseholdParamsWithTimeout(timeout time.Duration) *ConstituentsConvertIndividualToHouseholdParams {
	return &ConstituentsConvertIndividualToHouseholdParams{
		timeout: timeout,
	}
}

// NewConstituentsConvertIndividualToHouseholdParamsWithContext creates a new ConstituentsConvertIndividualToHouseholdParams object
// with the ability to set a context for a request.
func NewConstituentsConvertIndividualToHouseholdParamsWithContext(ctx context.Context) *ConstituentsConvertIndividualToHouseholdParams {
	return &ConstituentsConvertIndividualToHouseholdParams{
		Context: ctx,
	}
}

// NewConstituentsConvertIndividualToHouseholdParamsWithHTTPClient creates a new ConstituentsConvertIndividualToHouseholdParams object
// with the ability to set a custom HTTPClient for a request.
func NewConstituentsConvertIndividualToHouseholdParamsWithHTTPClient(client *http.Client) *ConstituentsConvertIndividualToHouseholdParams {
	return &ConstituentsConvertIndividualToHouseholdParams{
		HTTPClient: client,
	}
}

/*
ConstituentsConvertIndividualToHouseholdParams contains all the parameters to send to the API endpoint

	for the constituents convert individual to household operation.

	Typically these are written to a http.Request.
*/
type ConstituentsConvertIndividualToHouseholdParams struct {

	// ConstituentID.
	ConstituentID string

	// Request.
	Request *models.IndividualToHouseholdRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the constituents convert individual to household params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConstituentsConvertIndividualToHouseholdParams) WithDefaults() *ConstituentsConvertIndividualToHouseholdParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the constituents convert individual to household params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConstituentsConvertIndividualToHouseholdParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the constituents convert individual to household params
func (o *ConstituentsConvertIndividualToHouseholdParams) WithTimeout(timeout time.Duration) *ConstituentsConvertIndividualToHouseholdParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the constituents convert individual to household params
func (o *ConstituentsConvertIndividualToHouseholdParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the constituents convert individual to household params
func (o *ConstituentsConvertIndividualToHouseholdParams) WithContext(ctx context.Context) *ConstituentsConvertIndividualToHouseholdParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the constituents convert individual to household params
func (o *ConstituentsConvertIndividualToHouseholdParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the constituents convert individual to household params
func (o *ConstituentsConvertIndividualToHouseholdParams) WithHTTPClient(client *http.Client) *ConstituentsConvertIndividualToHouseholdParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the constituents convert individual to household params
func (o *ConstituentsConvertIndividualToHouseholdParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConstituentID adds the constituentID to the constituents convert individual to household params
func (o *ConstituentsConvertIndividualToHouseholdParams) WithConstituentID(constituentID string) *ConstituentsConvertIndividualToHouseholdParams {
	o.SetConstituentID(constituentID)
	return o
}

// SetConstituentID adds the constituentId to the constituents convert individual to household params
func (o *ConstituentsConvertIndividualToHouseholdParams) SetConstituentID(constituentID string) {
	o.ConstituentID = constituentID
}

// WithRequest adds the request to the constituents convert individual to household params
func (o *ConstituentsConvertIndividualToHouseholdParams) WithRequest(request *models.IndividualToHouseholdRequest) *ConstituentsConvertIndividualToHouseholdParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the constituents convert individual to household params
func (o *ConstituentsConvertIndividualToHouseholdParams) SetRequest(request *models.IndividualToHouseholdRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *ConstituentsConvertIndividualToHouseholdParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param constituentId
	if err := r.SetPathParam("constituentId", o.ConstituentID); err != nil {
		return err
	}
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
