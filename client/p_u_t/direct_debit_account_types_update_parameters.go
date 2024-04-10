// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

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

// NewDirectDebitAccountTypesUpdateParams creates a new DirectDebitAccountTypesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDirectDebitAccountTypesUpdateParams() *DirectDebitAccountTypesUpdateParams {
	return &DirectDebitAccountTypesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDirectDebitAccountTypesUpdateParamsWithTimeout creates a new DirectDebitAccountTypesUpdateParams object
// with the ability to set a timeout on a request.
func NewDirectDebitAccountTypesUpdateParamsWithTimeout(timeout time.Duration) *DirectDebitAccountTypesUpdateParams {
	return &DirectDebitAccountTypesUpdateParams{
		timeout: timeout,
	}
}

// NewDirectDebitAccountTypesUpdateParamsWithContext creates a new DirectDebitAccountTypesUpdateParams object
// with the ability to set a context for a request.
func NewDirectDebitAccountTypesUpdateParamsWithContext(ctx context.Context) *DirectDebitAccountTypesUpdateParams {
	return &DirectDebitAccountTypesUpdateParams{
		Context: ctx,
	}
}

// NewDirectDebitAccountTypesUpdateParamsWithHTTPClient creates a new DirectDebitAccountTypesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDirectDebitAccountTypesUpdateParamsWithHTTPClient(client *http.Client) *DirectDebitAccountTypesUpdateParams {
	return &DirectDebitAccountTypesUpdateParams{
		HTTPClient: client,
	}
}

/*
DirectDebitAccountTypesUpdateParams contains all the parameters to send to the API endpoint

	for the direct debit account types update operation.

	Typically these are written to a http.Request.
*/
type DirectDebitAccountTypesUpdateParams struct {

	// Data.
	Data *models.DirectDebitAccountType

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the direct debit account types update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DirectDebitAccountTypesUpdateParams) WithDefaults() *DirectDebitAccountTypesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the direct debit account types update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DirectDebitAccountTypesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the direct debit account types update params
func (o *DirectDebitAccountTypesUpdateParams) WithTimeout(timeout time.Duration) *DirectDebitAccountTypesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the direct debit account types update params
func (o *DirectDebitAccountTypesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the direct debit account types update params
func (o *DirectDebitAccountTypesUpdateParams) WithContext(ctx context.Context) *DirectDebitAccountTypesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the direct debit account types update params
func (o *DirectDebitAccountTypesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the direct debit account types update params
func (o *DirectDebitAccountTypesUpdateParams) WithHTTPClient(client *http.Client) *DirectDebitAccountTypesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the direct debit account types update params
func (o *DirectDebitAccountTypesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the direct debit account types update params
func (o *DirectDebitAccountTypesUpdateParams) WithData(data *models.DirectDebitAccountType) *DirectDebitAccountTypesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the direct debit account types update params
func (o *DirectDebitAccountTypesUpdateParams) SetData(data *models.DirectDebitAccountType) {
	o.Data = data
}

// WithID adds the id to the direct debit account types update params
func (o *DirectDebitAccountTypesUpdateParams) WithID(id string) *DirectDebitAccountTypesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the direct debit account types update params
func (o *DirectDebitAccountTypesUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DirectDebitAccountTypesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
