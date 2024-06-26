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

// NewCartUpdateCartPropertiesCustomDataParams creates a new CartUpdateCartPropertiesCustomDataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCartUpdateCartPropertiesCustomDataParams() *CartUpdateCartPropertiesCustomDataParams {
	return &CartUpdateCartPropertiesCustomDataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCartUpdateCartPropertiesCustomDataParamsWithTimeout creates a new CartUpdateCartPropertiesCustomDataParams object
// with the ability to set a timeout on a request.
func NewCartUpdateCartPropertiesCustomDataParamsWithTimeout(timeout time.Duration) *CartUpdateCartPropertiesCustomDataParams {
	return &CartUpdateCartPropertiesCustomDataParams{
		timeout: timeout,
	}
}

// NewCartUpdateCartPropertiesCustomDataParamsWithContext creates a new CartUpdateCartPropertiesCustomDataParams object
// with the ability to set a context for a request.
func NewCartUpdateCartPropertiesCustomDataParamsWithContext(ctx context.Context) *CartUpdateCartPropertiesCustomDataParams {
	return &CartUpdateCartPropertiesCustomDataParams{
		Context: ctx,
	}
}

// NewCartUpdateCartPropertiesCustomDataParamsWithHTTPClient creates a new CartUpdateCartPropertiesCustomDataParams object
// with the ability to set a custom HTTPClient for a request.
func NewCartUpdateCartPropertiesCustomDataParamsWithHTTPClient(client *http.Client) *CartUpdateCartPropertiesCustomDataParams {
	return &CartUpdateCartPropertiesCustomDataParams{
		HTTPClient: client,
	}
}

/*
CartUpdateCartPropertiesCustomDataParams contains all the parameters to send to the API endpoint

	for the cart update cart properties custom data operation.

	Typically these are written to a http.Request.
*/
type CartUpdateCartPropertiesCustomDataParams struct {

	// CustomDataItem.
	CustomDataItem *models.CustomDataItem

	// CustomID.
	CustomID string

	// SessionKey.
	SessionKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cart update cart properties custom data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartUpdateCartPropertiesCustomDataParams) WithDefaults() *CartUpdateCartPropertiesCustomDataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cart update cart properties custom data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartUpdateCartPropertiesCustomDataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cart update cart properties custom data params
func (o *CartUpdateCartPropertiesCustomDataParams) WithTimeout(timeout time.Duration) *CartUpdateCartPropertiesCustomDataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cart update cart properties custom data params
func (o *CartUpdateCartPropertiesCustomDataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cart update cart properties custom data params
func (o *CartUpdateCartPropertiesCustomDataParams) WithContext(ctx context.Context) *CartUpdateCartPropertiesCustomDataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cart update cart properties custom data params
func (o *CartUpdateCartPropertiesCustomDataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cart update cart properties custom data params
func (o *CartUpdateCartPropertiesCustomDataParams) WithHTTPClient(client *http.Client) *CartUpdateCartPropertiesCustomDataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cart update cart properties custom data params
func (o *CartUpdateCartPropertiesCustomDataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomDataItem adds the customDataItem to the cart update cart properties custom data params
func (o *CartUpdateCartPropertiesCustomDataParams) WithCustomDataItem(customDataItem *models.CustomDataItem) *CartUpdateCartPropertiesCustomDataParams {
	o.SetCustomDataItem(customDataItem)
	return o
}

// SetCustomDataItem adds the customDataItem to the cart update cart properties custom data params
func (o *CartUpdateCartPropertiesCustomDataParams) SetCustomDataItem(customDataItem *models.CustomDataItem) {
	o.CustomDataItem = customDataItem
}

// WithCustomID adds the customID to the cart update cart properties custom data params
func (o *CartUpdateCartPropertiesCustomDataParams) WithCustomID(customID string) *CartUpdateCartPropertiesCustomDataParams {
	o.SetCustomID(customID)
	return o
}

// SetCustomID adds the customId to the cart update cart properties custom data params
func (o *CartUpdateCartPropertiesCustomDataParams) SetCustomID(customID string) {
	o.CustomID = customID
}

// WithSessionKey adds the sessionKey to the cart update cart properties custom data params
func (o *CartUpdateCartPropertiesCustomDataParams) WithSessionKey(sessionKey string) *CartUpdateCartPropertiesCustomDataParams {
	o.SetSessionKey(sessionKey)
	return o
}

// SetSessionKey adds the sessionKey to the cart update cart properties custom data params
func (o *CartUpdateCartPropertiesCustomDataParams) SetSessionKey(sessionKey string) {
	o.SessionKey = sessionKey
}

// WriteToRequest writes these params to a swagger request
func (o *CartUpdateCartPropertiesCustomDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.CustomDataItem != nil {
		if err := r.SetBodyParam(o.CustomDataItem); err != nil {
			return err
		}
	}

	// path param customId
	if err := r.SetPathParam("customId", o.CustomID); err != nil {
		return err
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
