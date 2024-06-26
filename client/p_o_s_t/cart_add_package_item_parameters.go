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

// NewCartAddPackageItemParams creates a new CartAddPackageItemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCartAddPackageItemParams() *CartAddPackageItemParams {
	return &CartAddPackageItemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCartAddPackageItemParamsWithTimeout creates a new CartAddPackageItemParams object
// with the ability to set a timeout on a request.
func NewCartAddPackageItemParamsWithTimeout(timeout time.Duration) *CartAddPackageItemParams {
	return &CartAddPackageItemParams{
		timeout: timeout,
	}
}

// NewCartAddPackageItemParamsWithContext creates a new CartAddPackageItemParams object
// with the ability to set a context for a request.
func NewCartAddPackageItemParamsWithContext(ctx context.Context) *CartAddPackageItemParams {
	return &CartAddPackageItemParams{
		Context: ctx,
	}
}

// NewCartAddPackageItemParamsWithHTTPClient creates a new CartAddPackageItemParams object
// with the ability to set a custom HTTPClient for a request.
func NewCartAddPackageItemParamsWithHTTPClient(client *http.Client) *CartAddPackageItemParams {
	return &CartAddPackageItemParams{
		HTTPClient: client,
	}
}

/*
CartAddPackageItemParams contains all the parameters to send to the API endpoint

	for the cart add package item operation.

	Typically these are written to a http.Request.
*/
type CartAddPackageItemParams struct {

	// Request.
	Request *models.AddPackageItemRequest

	// SessionKey.
	SessionKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cart add package item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartAddPackageItemParams) WithDefaults() *CartAddPackageItemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cart add package item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartAddPackageItemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cart add package item params
func (o *CartAddPackageItemParams) WithTimeout(timeout time.Duration) *CartAddPackageItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cart add package item params
func (o *CartAddPackageItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cart add package item params
func (o *CartAddPackageItemParams) WithContext(ctx context.Context) *CartAddPackageItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cart add package item params
func (o *CartAddPackageItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cart add package item params
func (o *CartAddPackageItemParams) WithHTTPClient(client *http.Client) *CartAddPackageItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cart add package item params
func (o *CartAddPackageItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the cart add package item params
func (o *CartAddPackageItemParams) WithRequest(request *models.AddPackageItemRequest) *CartAddPackageItemParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the cart add package item params
func (o *CartAddPackageItemParams) SetRequest(request *models.AddPackageItemRequest) {
	o.Request = request
}

// WithSessionKey adds the sessionKey to the cart add package item params
func (o *CartAddPackageItemParams) WithSessionKey(sessionKey string) *CartAddPackageItemParams {
	o.SetSessionKey(sessionKey)
	return o
}

// SetSessionKey adds the sessionKey to the cart add package item params
func (o *CartAddPackageItemParams) SetSessionKey(sessionKey string) {
	o.SessionKey = sessionKey
}

// WriteToRequest writes these params to a swagger request
func (o *CartAddPackageItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	// path param sessionKey
	if err := r.SetPathParam("sessionKey", o.SessionKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
