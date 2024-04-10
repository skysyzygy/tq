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

// NewPronounsUpdateParams creates a new PronounsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPronounsUpdateParams() *PronounsUpdateParams {
	return &PronounsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPronounsUpdateParamsWithTimeout creates a new PronounsUpdateParams object
// with the ability to set a timeout on a request.
func NewPronounsUpdateParamsWithTimeout(timeout time.Duration) *PronounsUpdateParams {
	return &PronounsUpdateParams{
		timeout: timeout,
	}
}

// NewPronounsUpdateParamsWithContext creates a new PronounsUpdateParams object
// with the ability to set a context for a request.
func NewPronounsUpdateParamsWithContext(ctx context.Context) *PronounsUpdateParams {
	return &PronounsUpdateParams{
		Context: ctx,
	}
}

// NewPronounsUpdateParamsWithHTTPClient creates a new PronounsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPronounsUpdateParamsWithHTTPClient(client *http.Client) *PronounsUpdateParams {
	return &PronounsUpdateParams{
		HTTPClient: client,
	}
}

/*
PronounsUpdateParams contains all the parameters to send to the API endpoint

	for the pronouns update operation.

	Typically these are written to a http.Request.
*/
type PronounsUpdateParams struct {

	// Data.
	Data *models.Pronoun

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pronouns update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PronounsUpdateParams) WithDefaults() *PronounsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pronouns update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PronounsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pronouns update params
func (o *PronounsUpdateParams) WithTimeout(timeout time.Duration) *PronounsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pronouns update params
func (o *PronounsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pronouns update params
func (o *PronounsUpdateParams) WithContext(ctx context.Context) *PronounsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pronouns update params
func (o *PronounsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pronouns update params
func (o *PronounsUpdateParams) WithHTTPClient(client *http.Client) *PronounsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pronouns update params
func (o *PronounsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the pronouns update params
func (o *PronounsUpdateParams) WithData(data *models.Pronoun) *PronounsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the pronouns update params
func (o *PronounsUpdateParams) SetData(data *models.Pronoun) {
	o.Data = data
}

// WithID adds the id to the pronouns update params
func (o *PronounsUpdateParams) WithID(id string) *PronounsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the pronouns update params
func (o *PronounsUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PronounsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
