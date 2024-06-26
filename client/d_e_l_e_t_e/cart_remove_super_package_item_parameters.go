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

// NewCartRemoveSuperPackageItemParams creates a new CartRemoveSuperPackageItemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCartRemoveSuperPackageItemParams() *CartRemoveSuperPackageItemParams {
	return &CartRemoveSuperPackageItemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCartRemoveSuperPackageItemParamsWithTimeout creates a new CartRemoveSuperPackageItemParams object
// with the ability to set a timeout on a request.
func NewCartRemoveSuperPackageItemParamsWithTimeout(timeout time.Duration) *CartRemoveSuperPackageItemParams {
	return &CartRemoveSuperPackageItemParams{
		timeout: timeout,
	}
}

// NewCartRemoveSuperPackageItemParamsWithContext creates a new CartRemoveSuperPackageItemParams object
// with the ability to set a context for a request.
func NewCartRemoveSuperPackageItemParamsWithContext(ctx context.Context) *CartRemoveSuperPackageItemParams {
	return &CartRemoveSuperPackageItemParams{
		Context: ctx,
	}
}

// NewCartRemoveSuperPackageItemParamsWithHTTPClient creates a new CartRemoveSuperPackageItemParams object
// with the ability to set a custom HTTPClient for a request.
func NewCartRemoveSuperPackageItemParamsWithHTTPClient(client *http.Client) *CartRemoveSuperPackageItemParams {
	return &CartRemoveSuperPackageItemParams{
		HTTPClient: client,
	}
}

/*
CartRemoveSuperPackageItemParams contains all the parameters to send to the API endpoint

	for the cart remove super package item operation.

	Typically these are written to a http.Request.
*/
type CartRemoveSuperPackageItemParams struct {

	// SessionKey.
	SessionKey string

	/* SuperPackageID.

	   Id of the package to remove
	*/
	SuperPackageID string

	/* SuperPackageLineItemID.

	   Lineitem ID of the package line to remove
	*/
	SuperPackageLineItemID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cart remove super package item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartRemoveSuperPackageItemParams) WithDefaults() *CartRemoveSuperPackageItemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cart remove super package item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartRemoveSuperPackageItemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cart remove super package item params
func (o *CartRemoveSuperPackageItemParams) WithTimeout(timeout time.Duration) *CartRemoveSuperPackageItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cart remove super package item params
func (o *CartRemoveSuperPackageItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cart remove super package item params
func (o *CartRemoveSuperPackageItemParams) WithContext(ctx context.Context) *CartRemoveSuperPackageItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cart remove super package item params
func (o *CartRemoveSuperPackageItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cart remove super package item params
func (o *CartRemoveSuperPackageItemParams) WithHTTPClient(client *http.Client) *CartRemoveSuperPackageItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cart remove super package item params
func (o *CartRemoveSuperPackageItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSessionKey adds the sessionKey to the cart remove super package item params
func (o *CartRemoveSuperPackageItemParams) WithSessionKey(sessionKey string) *CartRemoveSuperPackageItemParams {
	o.SetSessionKey(sessionKey)
	return o
}

// SetSessionKey adds the sessionKey to the cart remove super package item params
func (o *CartRemoveSuperPackageItemParams) SetSessionKey(sessionKey string) {
	o.SessionKey = sessionKey
}

// WithSuperPackageID adds the superPackageID to the cart remove super package item params
func (o *CartRemoveSuperPackageItemParams) WithSuperPackageID(superPackageID string) *CartRemoveSuperPackageItemParams {
	o.SetSuperPackageID(superPackageID)
	return o
}

// SetSuperPackageID adds the superPackageId to the cart remove super package item params
func (o *CartRemoveSuperPackageItemParams) SetSuperPackageID(superPackageID string) {
	o.SuperPackageID = superPackageID
}

// WithSuperPackageLineItemID adds the superPackageLineItemID to the cart remove super package item params
func (o *CartRemoveSuperPackageItemParams) WithSuperPackageLineItemID(superPackageLineItemID string) *CartRemoveSuperPackageItemParams {
	o.SetSuperPackageLineItemID(superPackageLineItemID)
	return o
}

// SetSuperPackageLineItemID adds the superPackageLineItemId to the cart remove super package item params
func (o *CartRemoveSuperPackageItemParams) SetSuperPackageLineItemID(superPackageLineItemID string) {
	o.SuperPackageLineItemID = superPackageLineItemID
}

// WriteToRequest writes these params to a swagger request
func (o *CartRemoveSuperPackageItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param sessionKey
	if err := r.SetPathParam("sessionKey", o.SessionKey); err != nil {
		return err
	}

	// path param superPackageId
	if err := r.SetPathParam("superPackageId", o.SuperPackageID); err != nil {
		return err
	}

	// path param superPackageLineItemId
	if err := r.SetPathParam("superPackageLineItemId", o.SuperPackageLineItemID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
