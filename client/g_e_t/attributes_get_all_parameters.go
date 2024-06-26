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

// NewAttributesGetAllParams creates a new AttributesGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAttributesGetAllParams() *AttributesGetAllParams {
	return &AttributesGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAttributesGetAllParamsWithTimeout creates a new AttributesGetAllParams object
// with the ability to set a timeout on a request.
func NewAttributesGetAllParamsWithTimeout(timeout time.Duration) *AttributesGetAllParams {
	return &AttributesGetAllParams{
		timeout: timeout,
	}
}

// NewAttributesGetAllParamsWithContext creates a new AttributesGetAllParams object
// with the ability to set a context for a request.
func NewAttributesGetAllParamsWithContext(ctx context.Context) *AttributesGetAllParams {
	return &AttributesGetAllParams{
		Context: ctx,
	}
}

// NewAttributesGetAllParamsWithHTTPClient creates a new AttributesGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewAttributesGetAllParamsWithHTTPClient(client *http.Client) *AttributesGetAllParams {
	return &AttributesGetAllParams{
		HTTPClient: client,
	}
}

/*
AttributesGetAllParams contains all the parameters to send to the API endpoint

	for the attributes get all operation.

	Typically these are written to a http.Request.
*/
type AttributesGetAllParams struct {

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

// WithDefaults hydrates default values in the attributes get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AttributesGetAllParams) WithDefaults() *AttributesGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the attributes get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AttributesGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the attributes get all params
func (o *AttributesGetAllParams) WithTimeout(timeout time.Duration) *AttributesGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the attributes get all params
func (o *AttributesGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the attributes get all params
func (o *AttributesGetAllParams) WithContext(ctx context.Context) *AttributesGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the attributes get all params
func (o *AttributesGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the attributes get all params
func (o *AttributesGetAllParams) WithHTTPClient(client *http.Client) *AttributesGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the attributes get all params
func (o *AttributesGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConstituentID adds the constituentID to the attributes get all params
func (o *AttributesGetAllParams) WithConstituentID(constituentID *string) *AttributesGetAllParams {
	o.SetConstituentID(constituentID)
	return o
}

// SetConstituentID adds the constituentId to the attributes get all params
func (o *AttributesGetAllParams) SetConstituentID(constituentID *string) {
	o.ConstituentID = constituentID
}

// WithIncludeAffiliations adds the includeAffiliations to the attributes get all params
func (o *AttributesGetAllParams) WithIncludeAffiliations(includeAffiliations *string) *AttributesGetAllParams {
	o.SetIncludeAffiliations(includeAffiliations)
	return o
}

// SetIncludeAffiliations adds the includeAffiliations to the attributes get all params
func (o *AttributesGetAllParams) SetIncludeAffiliations(includeAffiliations *string) {
	o.IncludeAffiliations = includeAffiliations
}

// WriteToRequest writes these params to a swagger request
func (o *AttributesGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
