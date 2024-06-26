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

// NewCartApplySubLineItemDiscountParams creates a new CartApplySubLineItemDiscountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCartApplySubLineItemDiscountParams() *CartApplySubLineItemDiscountParams {
	return &CartApplySubLineItemDiscountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCartApplySubLineItemDiscountParamsWithTimeout creates a new CartApplySubLineItemDiscountParams object
// with the ability to set a timeout on a request.
func NewCartApplySubLineItemDiscountParamsWithTimeout(timeout time.Duration) *CartApplySubLineItemDiscountParams {
	return &CartApplySubLineItemDiscountParams{
		timeout: timeout,
	}
}

// NewCartApplySubLineItemDiscountParamsWithContext creates a new CartApplySubLineItemDiscountParams object
// with the ability to set a context for a request.
func NewCartApplySubLineItemDiscountParamsWithContext(ctx context.Context) *CartApplySubLineItemDiscountParams {
	return &CartApplySubLineItemDiscountParams{
		Context: ctx,
	}
}

// NewCartApplySubLineItemDiscountParamsWithHTTPClient creates a new CartApplySubLineItemDiscountParams object
// with the ability to set a custom HTTPClient for a request.
func NewCartApplySubLineItemDiscountParamsWithHTTPClient(client *http.Client) *CartApplySubLineItemDiscountParams {
	return &CartApplySubLineItemDiscountParams{
		HTTPClient: client,
	}
}

/*
CartApplySubLineItemDiscountParams contains all the parameters to send to the API endpoint

	for the cart apply sub line item discount operation.

	Typically these are written to a http.Request.
*/
type CartApplySubLineItemDiscountParams struct {

	// Request.
	Request *models.ApplyDiscountRequest

	// SessionKey.
	SessionKey string

	/* SubLineItemID.

	   The Id of the sub lineitem to apply the discount to. Must be a valid sub lineitem in the cart.
	*/
	SubLineItemID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cart apply sub line item discount params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartApplySubLineItemDiscountParams) WithDefaults() *CartApplySubLineItemDiscountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cart apply sub line item discount params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartApplySubLineItemDiscountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cart apply sub line item discount params
func (o *CartApplySubLineItemDiscountParams) WithTimeout(timeout time.Duration) *CartApplySubLineItemDiscountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cart apply sub line item discount params
func (o *CartApplySubLineItemDiscountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cart apply sub line item discount params
func (o *CartApplySubLineItemDiscountParams) WithContext(ctx context.Context) *CartApplySubLineItemDiscountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cart apply sub line item discount params
func (o *CartApplySubLineItemDiscountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cart apply sub line item discount params
func (o *CartApplySubLineItemDiscountParams) WithHTTPClient(client *http.Client) *CartApplySubLineItemDiscountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cart apply sub line item discount params
func (o *CartApplySubLineItemDiscountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the cart apply sub line item discount params
func (o *CartApplySubLineItemDiscountParams) WithRequest(request *models.ApplyDiscountRequest) *CartApplySubLineItemDiscountParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the cart apply sub line item discount params
func (o *CartApplySubLineItemDiscountParams) SetRequest(request *models.ApplyDiscountRequest) {
	o.Request = request
}

// WithSessionKey adds the sessionKey to the cart apply sub line item discount params
func (o *CartApplySubLineItemDiscountParams) WithSessionKey(sessionKey string) *CartApplySubLineItemDiscountParams {
	o.SetSessionKey(sessionKey)
	return o
}

// SetSessionKey adds the sessionKey to the cart apply sub line item discount params
func (o *CartApplySubLineItemDiscountParams) SetSessionKey(sessionKey string) {
	o.SessionKey = sessionKey
}

// WithSubLineItemID adds the subLineItemID to the cart apply sub line item discount params
func (o *CartApplySubLineItemDiscountParams) WithSubLineItemID(subLineItemID string) *CartApplySubLineItemDiscountParams {
	o.SetSubLineItemID(subLineItemID)
	return o
}

// SetSubLineItemID adds the subLineItemId to the cart apply sub line item discount params
func (o *CartApplySubLineItemDiscountParams) SetSubLineItemID(subLineItemID string) {
	o.SubLineItemID = subLineItemID
}

// WriteToRequest writes these params to a swagger request
func (o *CartApplySubLineItemDiscountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param subLineItemId
	if err := r.SetPathParam("subLineItemId", o.SubLineItemID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
