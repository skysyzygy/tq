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

// NewInterestsGetAllParams creates a new InterestsGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInterestsGetAllParams() *InterestsGetAllParams {
	return &InterestsGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInterestsGetAllParamsWithTimeout creates a new InterestsGetAllParams object
// with the ability to set a timeout on a request.
func NewInterestsGetAllParamsWithTimeout(timeout time.Duration) *InterestsGetAllParams {
	return &InterestsGetAllParams{
		timeout: timeout,
	}
}

// NewInterestsGetAllParamsWithContext creates a new InterestsGetAllParams object
// with the ability to set a context for a request.
func NewInterestsGetAllParamsWithContext(ctx context.Context) *InterestsGetAllParams {
	return &InterestsGetAllParams{
		Context: ctx,
	}
}

// NewInterestsGetAllParamsWithHTTPClient creates a new InterestsGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewInterestsGetAllParamsWithHTTPClient(client *http.Client) *InterestsGetAllParams {
	return &InterestsGetAllParams{
		HTTPClient: client,
	}
}

/*
InterestsGetAllParams contains all the parameters to send to the API endpoint

	for the interests get all operation.

	Typically these are written to a http.Request.
*/
type InterestsGetAllParams struct {

	// CategoryIds.
	CategoryIds *string

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

// WithDefaults hydrates default values in the interests get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InterestsGetAllParams) WithDefaults() *InterestsGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the interests get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InterestsGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the interests get all params
func (o *InterestsGetAllParams) WithTimeout(timeout time.Duration) *InterestsGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the interests get all params
func (o *InterestsGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the interests get all params
func (o *InterestsGetAllParams) WithContext(ctx context.Context) *InterestsGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the interests get all params
func (o *InterestsGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the interests get all params
func (o *InterestsGetAllParams) WithHTTPClient(client *http.Client) *InterestsGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the interests get all params
func (o *InterestsGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategoryIds adds the categoryIds to the interests get all params
func (o *InterestsGetAllParams) WithCategoryIds(categoryIds *string) *InterestsGetAllParams {
	o.SetCategoryIds(categoryIds)
	return o
}

// SetCategoryIds adds the categoryIds to the interests get all params
func (o *InterestsGetAllParams) SetCategoryIds(categoryIds *string) {
	o.CategoryIds = categoryIds
}

// WithConstituentID adds the constituentID to the interests get all params
func (o *InterestsGetAllParams) WithConstituentID(constituentID *string) *InterestsGetAllParams {
	o.SetConstituentID(constituentID)
	return o
}

// SetConstituentID adds the constituentId to the interests get all params
func (o *InterestsGetAllParams) SetConstituentID(constituentID *string) {
	o.ConstituentID = constituentID
}

// WithIncludeAffiliations adds the includeAffiliations to the interests get all params
func (o *InterestsGetAllParams) WithIncludeAffiliations(includeAffiliations *string) *InterestsGetAllParams {
	o.SetIncludeAffiliations(includeAffiliations)
	return o
}

// SetIncludeAffiliations adds the includeAffiliations to the interests get all params
func (o *InterestsGetAllParams) SetIncludeAffiliations(includeAffiliations *string) {
	o.IncludeAffiliations = includeAffiliations
}

// WriteToRequest writes these params to a swagger request
func (o *InterestsGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CategoryIds != nil {

		// query param categoryIds
		var qrCategoryIds string

		if o.CategoryIds != nil {
			qrCategoryIds = *o.CategoryIds
		}
		qCategoryIds := qrCategoryIds
		if qCategoryIds != "" {

			if err := r.SetQueryParam("categoryIds", qCategoryIds); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
