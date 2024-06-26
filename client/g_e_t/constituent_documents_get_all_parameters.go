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

// NewConstituentDocumentsGetAllParams creates a new ConstituentDocumentsGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConstituentDocumentsGetAllParams() *ConstituentDocumentsGetAllParams {
	return &ConstituentDocumentsGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConstituentDocumentsGetAllParamsWithTimeout creates a new ConstituentDocumentsGetAllParams object
// with the ability to set a timeout on a request.
func NewConstituentDocumentsGetAllParamsWithTimeout(timeout time.Duration) *ConstituentDocumentsGetAllParams {
	return &ConstituentDocumentsGetAllParams{
		timeout: timeout,
	}
}

// NewConstituentDocumentsGetAllParamsWithContext creates a new ConstituentDocumentsGetAllParams object
// with the ability to set a context for a request.
func NewConstituentDocumentsGetAllParamsWithContext(ctx context.Context) *ConstituentDocumentsGetAllParams {
	return &ConstituentDocumentsGetAllParams{
		Context: ctx,
	}
}

// NewConstituentDocumentsGetAllParamsWithHTTPClient creates a new ConstituentDocumentsGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewConstituentDocumentsGetAllParamsWithHTTPClient(client *http.Client) *ConstituentDocumentsGetAllParams {
	return &ConstituentDocumentsGetAllParams{
		HTTPClient: client,
	}
}

/*
ConstituentDocumentsGetAllParams contains all the parameters to send to the API endpoint

	for the constituent documents get all operation.

	Typically these are written to a http.Request.
*/
type ConstituentDocumentsGetAllParams struct {

	// ConstituentID.
	ConstituentID string

	// IncludeAffiliations.
	IncludeAffiliations *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the constituent documents get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConstituentDocumentsGetAllParams) WithDefaults() *ConstituentDocumentsGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the constituent documents get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConstituentDocumentsGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the constituent documents get all params
func (o *ConstituentDocumentsGetAllParams) WithTimeout(timeout time.Duration) *ConstituentDocumentsGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the constituent documents get all params
func (o *ConstituentDocumentsGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the constituent documents get all params
func (o *ConstituentDocumentsGetAllParams) WithContext(ctx context.Context) *ConstituentDocumentsGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the constituent documents get all params
func (o *ConstituentDocumentsGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the constituent documents get all params
func (o *ConstituentDocumentsGetAllParams) WithHTTPClient(client *http.Client) *ConstituentDocumentsGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the constituent documents get all params
func (o *ConstituentDocumentsGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConstituentID adds the constituentID to the constituent documents get all params
func (o *ConstituentDocumentsGetAllParams) WithConstituentID(constituentID string) *ConstituentDocumentsGetAllParams {
	o.SetConstituentID(constituentID)
	return o
}

// SetConstituentID adds the constituentId to the constituent documents get all params
func (o *ConstituentDocumentsGetAllParams) SetConstituentID(constituentID string) {
	o.ConstituentID = constituentID
}

// WithIncludeAffiliations adds the includeAffiliations to the constituent documents get all params
func (o *ConstituentDocumentsGetAllParams) WithIncludeAffiliations(includeAffiliations *string) *ConstituentDocumentsGetAllParams {
	o.SetIncludeAffiliations(includeAffiliations)
	return o
}

// SetIncludeAffiliations adds the includeAffiliations to the constituent documents get all params
func (o *ConstituentDocumentsGetAllParams) SetIncludeAffiliations(includeAffiliations *string) {
	o.IncludeAffiliations = includeAffiliations
}

// WriteToRequest writes these params to a swagger request
func (o *ConstituentDocumentsGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param constituentId
	qrConstituentID := o.ConstituentID
	qConstituentID := qrConstituentID
	if qConstituentID != "" {

		if err := r.SetQueryParam("constituentId", qConstituentID); err != nil {
			return err
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
