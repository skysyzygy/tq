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

// NewCartUpdateCartPropertiesParams creates a new CartUpdateCartPropertiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCartUpdateCartPropertiesParams() *CartUpdateCartPropertiesParams {
	return &CartUpdateCartPropertiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCartUpdateCartPropertiesParamsWithTimeout creates a new CartUpdateCartPropertiesParams object
// with the ability to set a timeout on a request.
func NewCartUpdateCartPropertiesParamsWithTimeout(timeout time.Duration) *CartUpdateCartPropertiesParams {
	return &CartUpdateCartPropertiesParams{
		timeout: timeout,
	}
}

// NewCartUpdateCartPropertiesParamsWithContext creates a new CartUpdateCartPropertiesParams object
// with the ability to set a context for a request.
func NewCartUpdateCartPropertiesParamsWithContext(ctx context.Context) *CartUpdateCartPropertiesParams {
	return &CartUpdateCartPropertiesParams{
		Context: ctx,
	}
}

// NewCartUpdateCartPropertiesParamsWithHTTPClient creates a new CartUpdateCartPropertiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewCartUpdateCartPropertiesParamsWithHTTPClient(client *http.Client) *CartUpdateCartPropertiesParams {
	return &CartUpdateCartPropertiesParams{
		HTTPClient: client,
	}
}

/*
CartUpdateCartPropertiesParams contains all the parameters to send to the API endpoint

	for the cart update cart properties operation.

	Typically these are written to a http.Request.
*/
type CartUpdateCartPropertiesParams struct {

	// CartProperties.
	CartProperties *models.CartProperties

	// SessionKey.
	SessionKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cart update cart properties params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartUpdateCartPropertiesParams) WithDefaults() *CartUpdateCartPropertiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cart update cart properties params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartUpdateCartPropertiesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cart update cart properties params
func (o *CartUpdateCartPropertiesParams) WithTimeout(timeout time.Duration) *CartUpdateCartPropertiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cart update cart properties params
func (o *CartUpdateCartPropertiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cart update cart properties params
func (o *CartUpdateCartPropertiesParams) WithContext(ctx context.Context) *CartUpdateCartPropertiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cart update cart properties params
func (o *CartUpdateCartPropertiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cart update cart properties params
func (o *CartUpdateCartPropertiesParams) WithHTTPClient(client *http.Client) *CartUpdateCartPropertiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cart update cart properties params
func (o *CartUpdateCartPropertiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartProperties adds the cartProperties to the cart update cart properties params
func (o *CartUpdateCartPropertiesParams) WithCartProperties(cartProperties *models.CartProperties) *CartUpdateCartPropertiesParams {
	o.SetCartProperties(cartProperties)
	return o
}

// SetCartProperties adds the cartProperties to the cart update cart properties params
func (o *CartUpdateCartPropertiesParams) SetCartProperties(cartProperties *models.CartProperties) {
	o.CartProperties = cartProperties
}

// WithSessionKey adds the sessionKey to the cart update cart properties params
func (o *CartUpdateCartPropertiesParams) WithSessionKey(sessionKey string) *CartUpdateCartPropertiesParams {
	o.SetSessionKey(sessionKey)
	return o
}

// SetSessionKey adds the sessionKey to the cart update cart properties params
func (o *CartUpdateCartPropertiesParams) SetSessionKey(sessionKey string) {
	o.SessionKey = sessionKey
}

// WriteToRequest writes these params to a swagger request
func (o *CartUpdateCartPropertiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.CartProperties != nil {
		if err := r.SetBodyParam(o.CartProperties); err != nil {
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
