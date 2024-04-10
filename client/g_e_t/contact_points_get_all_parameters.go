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

// NewContactPointsGetAllParams creates a new ContactPointsGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContactPointsGetAllParams() *ContactPointsGetAllParams {
	return &ContactPointsGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContactPointsGetAllParamsWithTimeout creates a new ContactPointsGetAllParams object
// with the ability to set a timeout on a request.
func NewContactPointsGetAllParamsWithTimeout(timeout time.Duration) *ContactPointsGetAllParams {
	return &ContactPointsGetAllParams{
		timeout: timeout,
	}
}

// NewContactPointsGetAllParamsWithContext creates a new ContactPointsGetAllParams object
// with the ability to set a context for a request.
func NewContactPointsGetAllParamsWithContext(ctx context.Context) *ContactPointsGetAllParams {
	return &ContactPointsGetAllParams{
		Context: ctx,
	}
}

// NewContactPointsGetAllParamsWithHTTPClient creates a new ContactPointsGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewContactPointsGetAllParamsWithHTTPClient(client *http.Client) *ContactPointsGetAllParams {
	return &ContactPointsGetAllParams{
		HTTPClient: client,
	}
}

/*
ContactPointsGetAllParams contains all the parameters to send to the API endpoint

	for the contact points get all operation.

	Typically these are written to a http.Request.
*/
type ContactPointsGetAllParams struct {

	/* ConstituentID.

	   Limit results by constituent
	*/
	ConstituentID *string

	/* IncludeAffiliations.

	   Include all of the constituent's affiliates in the results (default: true)
	*/
	IncludeAffiliations *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the contact points get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactPointsGetAllParams) WithDefaults() *ContactPointsGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the contact points get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactPointsGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the contact points get all params
func (o *ContactPointsGetAllParams) WithTimeout(timeout time.Duration) *ContactPointsGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contact points get all params
func (o *ContactPointsGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contact points get all params
func (o *ContactPointsGetAllParams) WithContext(ctx context.Context) *ContactPointsGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contact points get all params
func (o *ContactPointsGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contact points get all params
func (o *ContactPointsGetAllParams) WithHTTPClient(client *http.Client) *ContactPointsGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contact points get all params
func (o *ContactPointsGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConstituentID adds the constituentID to the contact points get all params
func (o *ContactPointsGetAllParams) WithConstituentID(constituentID *string) *ContactPointsGetAllParams {
	o.SetConstituentID(constituentID)
	return o
}

// SetConstituentID adds the constituentId to the contact points get all params
func (o *ContactPointsGetAllParams) SetConstituentID(constituentID *string) {
	o.ConstituentID = constituentID
}

// WithIncludeAffiliations adds the includeAffiliations to the contact points get all params
func (o *ContactPointsGetAllParams) WithIncludeAffiliations(includeAffiliations *string) *ContactPointsGetAllParams {
	o.SetIncludeAffiliations(includeAffiliations)
	return o
}

// SetIncludeAffiliations adds the includeAffiliations to the contact points get all params
func (o *ContactPointsGetAllParams) SetIncludeAffiliations(includeAffiliations *string) {
	o.IncludeAffiliations = includeAffiliations
}

// WriteToRequest writes these params to a swagger request
func (o *ContactPointsGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
