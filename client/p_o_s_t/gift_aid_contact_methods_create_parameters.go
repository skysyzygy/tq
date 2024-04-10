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

// NewGiftAidContactMethodsCreateParams creates a new GiftAidContactMethodsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGiftAidContactMethodsCreateParams() *GiftAidContactMethodsCreateParams {
	return &GiftAidContactMethodsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGiftAidContactMethodsCreateParamsWithTimeout creates a new GiftAidContactMethodsCreateParams object
// with the ability to set a timeout on a request.
func NewGiftAidContactMethodsCreateParamsWithTimeout(timeout time.Duration) *GiftAidContactMethodsCreateParams {
	return &GiftAidContactMethodsCreateParams{
		timeout: timeout,
	}
}

// NewGiftAidContactMethodsCreateParamsWithContext creates a new GiftAidContactMethodsCreateParams object
// with the ability to set a context for a request.
func NewGiftAidContactMethodsCreateParamsWithContext(ctx context.Context) *GiftAidContactMethodsCreateParams {
	return &GiftAidContactMethodsCreateParams{
		Context: ctx,
	}
}

// NewGiftAidContactMethodsCreateParamsWithHTTPClient creates a new GiftAidContactMethodsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGiftAidContactMethodsCreateParamsWithHTTPClient(client *http.Client) *GiftAidContactMethodsCreateParams {
	return &GiftAidContactMethodsCreateParams{
		HTTPClient: client,
	}
}

/*
GiftAidContactMethodsCreateParams contains all the parameters to send to the API endpoint

	for the gift aid contact methods create operation.

	Typically these are written to a http.Request.
*/
type GiftAidContactMethodsCreateParams struct {

	// Data.
	Data *models.GiftAidContactMethod

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the gift aid contact methods create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GiftAidContactMethodsCreateParams) WithDefaults() *GiftAidContactMethodsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the gift aid contact methods create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GiftAidContactMethodsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the gift aid contact methods create params
func (o *GiftAidContactMethodsCreateParams) WithTimeout(timeout time.Duration) *GiftAidContactMethodsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gift aid contact methods create params
func (o *GiftAidContactMethodsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gift aid contact methods create params
func (o *GiftAidContactMethodsCreateParams) WithContext(ctx context.Context) *GiftAidContactMethodsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gift aid contact methods create params
func (o *GiftAidContactMethodsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gift aid contact methods create params
func (o *GiftAidContactMethodsCreateParams) WithHTTPClient(client *http.Client) *GiftAidContactMethodsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gift aid contact methods create params
func (o *GiftAidContactMethodsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the gift aid contact methods create params
func (o *GiftAidContactMethodsCreateParams) WithData(data *models.GiftAidContactMethod) *GiftAidContactMethodsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the gift aid contact methods create params
func (o *GiftAidContactMethodsCreateParams) SetData(data *models.GiftAidContactMethod) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *GiftAidContactMethodsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
