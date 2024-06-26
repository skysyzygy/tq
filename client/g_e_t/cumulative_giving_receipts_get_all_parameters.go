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

// NewCumulativeGivingReceiptsGetAllParams creates a new CumulativeGivingReceiptsGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCumulativeGivingReceiptsGetAllParams() *CumulativeGivingReceiptsGetAllParams {
	return &CumulativeGivingReceiptsGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCumulativeGivingReceiptsGetAllParamsWithTimeout creates a new CumulativeGivingReceiptsGetAllParams object
// with the ability to set a timeout on a request.
func NewCumulativeGivingReceiptsGetAllParamsWithTimeout(timeout time.Duration) *CumulativeGivingReceiptsGetAllParams {
	return &CumulativeGivingReceiptsGetAllParams{
		timeout: timeout,
	}
}

// NewCumulativeGivingReceiptsGetAllParamsWithContext creates a new CumulativeGivingReceiptsGetAllParams object
// with the ability to set a context for a request.
func NewCumulativeGivingReceiptsGetAllParamsWithContext(ctx context.Context) *CumulativeGivingReceiptsGetAllParams {
	return &CumulativeGivingReceiptsGetAllParams{
		Context: ctx,
	}
}

// NewCumulativeGivingReceiptsGetAllParamsWithHTTPClient creates a new CumulativeGivingReceiptsGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewCumulativeGivingReceiptsGetAllParamsWithHTTPClient(client *http.Client) *CumulativeGivingReceiptsGetAllParams {
	return &CumulativeGivingReceiptsGetAllParams{
		HTTPClient: client,
	}
}

/*
CumulativeGivingReceiptsGetAllParams contains all the parameters to send to the API endpoint

	for the cumulative giving receipts get all operation.

	Typically these are written to a http.Request.
*/
type CumulativeGivingReceiptsGetAllParams struct {

	// ConstituentID.
	ConstituentID *string

	// IncludeAffiliations.
	IncludeAffiliations *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cumulative giving receipts get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CumulativeGivingReceiptsGetAllParams) WithDefaults() *CumulativeGivingReceiptsGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cumulative giving receipts get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CumulativeGivingReceiptsGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cumulative giving receipts get all params
func (o *CumulativeGivingReceiptsGetAllParams) WithTimeout(timeout time.Duration) *CumulativeGivingReceiptsGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cumulative giving receipts get all params
func (o *CumulativeGivingReceiptsGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cumulative giving receipts get all params
func (o *CumulativeGivingReceiptsGetAllParams) WithContext(ctx context.Context) *CumulativeGivingReceiptsGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cumulative giving receipts get all params
func (o *CumulativeGivingReceiptsGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cumulative giving receipts get all params
func (o *CumulativeGivingReceiptsGetAllParams) WithHTTPClient(client *http.Client) *CumulativeGivingReceiptsGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cumulative giving receipts get all params
func (o *CumulativeGivingReceiptsGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConstituentID adds the constituentID to the cumulative giving receipts get all params
func (o *CumulativeGivingReceiptsGetAllParams) WithConstituentID(constituentID *string) *CumulativeGivingReceiptsGetAllParams {
	o.SetConstituentID(constituentID)
	return o
}

// SetConstituentID adds the constituentId to the cumulative giving receipts get all params
func (o *CumulativeGivingReceiptsGetAllParams) SetConstituentID(constituentID *string) {
	o.ConstituentID = constituentID
}

// WithIncludeAffiliations adds the includeAffiliations to the cumulative giving receipts get all params
func (o *CumulativeGivingReceiptsGetAllParams) WithIncludeAffiliations(includeAffiliations *string) *CumulativeGivingReceiptsGetAllParams {
	o.SetIncludeAffiliations(includeAffiliations)
	return o
}

// SetIncludeAffiliations adds the includeAffiliations to the cumulative giving receipts get all params
func (o *CumulativeGivingReceiptsGetAllParams) SetIncludeAffiliations(includeAffiliations *string) {
	o.IncludeAffiliations = includeAffiliations
}

// WriteToRequest writes these params to a swagger request
func (o *CumulativeGivingReceiptsGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
