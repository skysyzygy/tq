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

// NewCartRemovePackageItemParams creates a new CartRemovePackageItemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCartRemovePackageItemParams() *CartRemovePackageItemParams {
	return &CartRemovePackageItemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCartRemovePackageItemParamsWithTimeout creates a new CartRemovePackageItemParams object
// with the ability to set a timeout on a request.
func NewCartRemovePackageItemParamsWithTimeout(timeout time.Duration) *CartRemovePackageItemParams {
	return &CartRemovePackageItemParams{
		timeout: timeout,
	}
}

// NewCartRemovePackageItemParamsWithContext creates a new CartRemovePackageItemParams object
// with the ability to set a context for a request.
func NewCartRemovePackageItemParamsWithContext(ctx context.Context) *CartRemovePackageItemParams {
	return &CartRemovePackageItemParams{
		Context: ctx,
	}
}

// NewCartRemovePackageItemParamsWithHTTPClient creates a new CartRemovePackageItemParams object
// with the ability to set a custom HTTPClient for a request.
func NewCartRemovePackageItemParamsWithHTTPClient(client *http.Client) *CartRemovePackageItemParams {
	return &CartRemovePackageItemParams{
		HTTPClient: client,
	}
}

/*
CartRemovePackageItemParams contains all the parameters to send to the API endpoint

	for the cart remove package item operation.

	Typically these are written to a http.Request.
*/
type CartRemovePackageItemParams struct {

	/* LineItemID.

	   Lineitem ID for the package line to remove
	*/
	LineItemID string

	/* PackageID.

	   Id of the packge to remove
	*/
	PackageID string

	// SessionKey.
	SessionKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cart remove package item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartRemovePackageItemParams) WithDefaults() *CartRemovePackageItemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cart remove package item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartRemovePackageItemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cart remove package item params
func (o *CartRemovePackageItemParams) WithTimeout(timeout time.Duration) *CartRemovePackageItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cart remove package item params
func (o *CartRemovePackageItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cart remove package item params
func (o *CartRemovePackageItemParams) WithContext(ctx context.Context) *CartRemovePackageItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cart remove package item params
func (o *CartRemovePackageItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cart remove package item params
func (o *CartRemovePackageItemParams) WithHTTPClient(client *http.Client) *CartRemovePackageItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cart remove package item params
func (o *CartRemovePackageItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLineItemID adds the lineItemID to the cart remove package item params
func (o *CartRemovePackageItemParams) WithLineItemID(lineItemID string) *CartRemovePackageItemParams {
	o.SetLineItemID(lineItemID)
	return o
}

// SetLineItemID adds the lineItemId to the cart remove package item params
func (o *CartRemovePackageItemParams) SetLineItemID(lineItemID string) {
	o.LineItemID = lineItemID
}

// WithPackageID adds the packageID to the cart remove package item params
func (o *CartRemovePackageItemParams) WithPackageID(packageID string) *CartRemovePackageItemParams {
	o.SetPackageID(packageID)
	return o
}

// SetPackageID adds the packageId to the cart remove package item params
func (o *CartRemovePackageItemParams) SetPackageID(packageID string) {
	o.PackageID = packageID
}

// WithSessionKey adds the sessionKey to the cart remove package item params
func (o *CartRemovePackageItemParams) WithSessionKey(sessionKey string) *CartRemovePackageItemParams {
	o.SetSessionKey(sessionKey)
	return o
}

// SetSessionKey adds the sessionKey to the cart remove package item params
func (o *CartRemovePackageItemParams) SetSessionKey(sessionKey string) {
	o.SessionKey = sessionKey
}

// WriteToRequest writes these params to a swagger request
func (o *CartRemovePackageItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param lineItemId
	if err := r.SetPathParam("lineItemId", o.LineItemID); err != nil {
		return err
	}

	// path param packageId
	if err := r.SetPathParam("packageId", o.PackageID); err != nil {
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
