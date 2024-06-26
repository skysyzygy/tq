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

// NewResearchGetAllParams creates a new ResearchGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResearchGetAllParams() *ResearchGetAllParams {
	return &ResearchGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResearchGetAllParamsWithTimeout creates a new ResearchGetAllParams object
// with the ability to set a timeout on a request.
func NewResearchGetAllParamsWithTimeout(timeout time.Duration) *ResearchGetAllParams {
	return &ResearchGetAllParams{
		timeout: timeout,
	}
}

// NewResearchGetAllParamsWithContext creates a new ResearchGetAllParams object
// with the ability to set a context for a request.
func NewResearchGetAllParamsWithContext(ctx context.Context) *ResearchGetAllParams {
	return &ResearchGetAllParams{
		Context: ctx,
	}
}

// NewResearchGetAllParamsWithHTTPClient creates a new ResearchGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewResearchGetAllParamsWithHTTPClient(client *http.Client) *ResearchGetAllParams {
	return &ResearchGetAllParams{
		HTTPClient: client,
	}
}

/*
ResearchGetAllParams contains all the parameters to send to the API endpoint

	for the research get all operation.

	Typically these are written to a http.Request.
*/
type ResearchGetAllParams struct {

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

// WithDefaults hydrates default values in the research get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResearchGetAllParams) WithDefaults() *ResearchGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the research get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResearchGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the research get all params
func (o *ResearchGetAllParams) WithTimeout(timeout time.Duration) *ResearchGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the research get all params
func (o *ResearchGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the research get all params
func (o *ResearchGetAllParams) WithContext(ctx context.Context) *ResearchGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the research get all params
func (o *ResearchGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the research get all params
func (o *ResearchGetAllParams) WithHTTPClient(client *http.Client) *ResearchGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the research get all params
func (o *ResearchGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConstituentID adds the constituentID to the research get all params
func (o *ResearchGetAllParams) WithConstituentID(constituentID *string) *ResearchGetAllParams {
	o.SetConstituentID(constituentID)
	return o
}

// SetConstituentID adds the constituentId to the research get all params
func (o *ResearchGetAllParams) SetConstituentID(constituentID *string) {
	o.ConstituentID = constituentID
}

// WithIncludeAffiliations adds the includeAffiliations to the research get all params
func (o *ResearchGetAllParams) WithIncludeAffiliations(includeAffiliations *string) *ResearchGetAllParams {
	o.SetIncludeAffiliations(includeAffiliations)
	return o
}

// SetIncludeAffiliations adds the includeAffiliations to the research get all params
func (o *ResearchGetAllParams) SetIncludeAffiliations(includeAffiliations *string) {
	o.IncludeAffiliations = includeAffiliations
}

// WriteToRequest writes these params to a swagger request
func (o *ResearchGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
