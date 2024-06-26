// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

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

// NewCartRemoveContributionCustomDataParams creates a new CartRemoveContributionCustomDataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCartRemoveContributionCustomDataParams() *CartRemoveContributionCustomDataParams {
	return &CartRemoveContributionCustomDataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCartRemoveContributionCustomDataParamsWithTimeout creates a new CartRemoveContributionCustomDataParams object
// with the ability to set a timeout on a request.
func NewCartRemoveContributionCustomDataParamsWithTimeout(timeout time.Duration) *CartRemoveContributionCustomDataParams {
	return &CartRemoveContributionCustomDataParams{
		timeout: timeout,
	}
}

// NewCartRemoveContributionCustomDataParamsWithContext creates a new CartRemoveContributionCustomDataParams object
// with the ability to set a context for a request.
func NewCartRemoveContributionCustomDataParamsWithContext(ctx context.Context) *CartRemoveContributionCustomDataParams {
	return &CartRemoveContributionCustomDataParams{
		Context: ctx,
	}
}

// NewCartRemoveContributionCustomDataParamsWithHTTPClient creates a new CartRemoveContributionCustomDataParams object
// with the ability to set a custom HTTPClient for a request.
func NewCartRemoveContributionCustomDataParamsWithHTTPClient(client *http.Client) *CartRemoveContributionCustomDataParams {
	return &CartRemoveContributionCustomDataParams{
		HTTPClient: client,
	}
}

/*
CartRemoveContributionCustomDataParams contains all the parameters to send to the API endpoint

	for the cart remove contribution custom data operation.

	Typically these are written to a http.Request.
*/
type CartRemoveContributionCustomDataParams struct {

	/* CustomID.

	   The Index of custom data value to remove
	*/
	CustomID string

	/* LineItemID.

	   The Reference Number of the contribution to remove the custom data value for
	*/
	LineItemID string

	// SessionKey.
	SessionKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cart remove contribution custom data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartRemoveContributionCustomDataParams) WithDefaults() *CartRemoveContributionCustomDataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cart remove contribution custom data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartRemoveContributionCustomDataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cart remove contribution custom data params
func (o *CartRemoveContributionCustomDataParams) WithTimeout(timeout time.Duration) *CartRemoveContributionCustomDataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cart remove contribution custom data params
func (o *CartRemoveContributionCustomDataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cart remove contribution custom data params
func (o *CartRemoveContributionCustomDataParams) WithContext(ctx context.Context) *CartRemoveContributionCustomDataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cart remove contribution custom data params
func (o *CartRemoveContributionCustomDataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cart remove contribution custom data params
func (o *CartRemoveContributionCustomDataParams) WithHTTPClient(client *http.Client) *CartRemoveContributionCustomDataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cart remove contribution custom data params
func (o *CartRemoveContributionCustomDataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomID adds the customID to the cart remove contribution custom data params
func (o *CartRemoveContributionCustomDataParams) WithCustomID(customID string) *CartRemoveContributionCustomDataParams {
	o.SetCustomID(customID)
	return o
}

// SetCustomID adds the customId to the cart remove contribution custom data params
func (o *CartRemoveContributionCustomDataParams) SetCustomID(customID string) {
	o.CustomID = customID
}

// WithLineItemID adds the lineItemID to the cart remove contribution custom data params
func (o *CartRemoveContributionCustomDataParams) WithLineItemID(lineItemID string) *CartRemoveContributionCustomDataParams {
	o.SetLineItemID(lineItemID)
	return o
}

// SetLineItemID adds the lineItemId to the cart remove contribution custom data params
func (o *CartRemoveContributionCustomDataParams) SetLineItemID(lineItemID string) {
	o.LineItemID = lineItemID
}

// WithSessionKey adds the sessionKey to the cart remove contribution custom data params
func (o *CartRemoveContributionCustomDataParams) WithSessionKey(sessionKey string) *CartRemoveContributionCustomDataParams {
	o.SetSessionKey(sessionKey)
	return o
}

// SetSessionKey adds the sessionKey to the cart remove contribution custom data params
func (o *CartRemoveContributionCustomDataParams) SetSessionKey(sessionKey string) {
	o.SessionKey = sessionKey
}

// WriteToRequest writes these params to a swagger request
func (o *CartRemoveContributionCustomDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customId
	if err := r.SetPathParam("customId", o.CustomID); err != nil {
		return err
	}

	// path param lineItemId
	if err := r.SetPathParam("lineItemId", o.LineItemID); err != nil {
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
