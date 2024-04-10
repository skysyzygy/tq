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

// NewCartUpdateSubLineItemRecipientParams creates a new CartUpdateSubLineItemRecipientParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCartUpdateSubLineItemRecipientParams() *CartUpdateSubLineItemRecipientParams {
	return &CartUpdateSubLineItemRecipientParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCartUpdateSubLineItemRecipientParamsWithTimeout creates a new CartUpdateSubLineItemRecipientParams object
// with the ability to set a timeout on a request.
func NewCartUpdateSubLineItemRecipientParamsWithTimeout(timeout time.Duration) *CartUpdateSubLineItemRecipientParams {
	return &CartUpdateSubLineItemRecipientParams{
		timeout: timeout,
	}
}

// NewCartUpdateSubLineItemRecipientParamsWithContext creates a new CartUpdateSubLineItemRecipientParams object
// with the ability to set a context for a request.
func NewCartUpdateSubLineItemRecipientParamsWithContext(ctx context.Context) *CartUpdateSubLineItemRecipientParams {
	return &CartUpdateSubLineItemRecipientParams{
		Context: ctx,
	}
}

// NewCartUpdateSubLineItemRecipientParamsWithHTTPClient creates a new CartUpdateSubLineItemRecipientParams object
// with the ability to set a custom HTTPClient for a request.
func NewCartUpdateSubLineItemRecipientParamsWithHTTPClient(client *http.Client) *CartUpdateSubLineItemRecipientParams {
	return &CartUpdateSubLineItemRecipientParams{
		HTTPClient: client,
	}
}

/*
CartUpdateSubLineItemRecipientParams contains all the parameters to send to the API endpoint

	for the cart update sub line item recipient operation.

	Typically these are written to a http.Request.
*/
type CartUpdateSubLineItemRecipientParams struct {

	// Request.
	Request *models.UpdateRecipientRequest

	// SessionKey.
	SessionKey string

	// SubLineItemID.
	SubLineItemID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cart update sub line item recipient params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartUpdateSubLineItemRecipientParams) WithDefaults() *CartUpdateSubLineItemRecipientParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cart update sub line item recipient params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartUpdateSubLineItemRecipientParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cart update sub line item recipient params
func (o *CartUpdateSubLineItemRecipientParams) WithTimeout(timeout time.Duration) *CartUpdateSubLineItemRecipientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cart update sub line item recipient params
func (o *CartUpdateSubLineItemRecipientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cart update sub line item recipient params
func (o *CartUpdateSubLineItemRecipientParams) WithContext(ctx context.Context) *CartUpdateSubLineItemRecipientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cart update sub line item recipient params
func (o *CartUpdateSubLineItemRecipientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cart update sub line item recipient params
func (o *CartUpdateSubLineItemRecipientParams) WithHTTPClient(client *http.Client) *CartUpdateSubLineItemRecipientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cart update sub line item recipient params
func (o *CartUpdateSubLineItemRecipientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the cart update sub line item recipient params
func (o *CartUpdateSubLineItemRecipientParams) WithRequest(request *models.UpdateRecipientRequest) *CartUpdateSubLineItemRecipientParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the cart update sub line item recipient params
func (o *CartUpdateSubLineItemRecipientParams) SetRequest(request *models.UpdateRecipientRequest) {
	o.Request = request
}

// WithSessionKey adds the sessionKey to the cart update sub line item recipient params
func (o *CartUpdateSubLineItemRecipientParams) WithSessionKey(sessionKey string) *CartUpdateSubLineItemRecipientParams {
	o.SetSessionKey(sessionKey)
	return o
}

// SetSessionKey adds the sessionKey to the cart update sub line item recipient params
func (o *CartUpdateSubLineItemRecipientParams) SetSessionKey(sessionKey string) {
	o.SessionKey = sessionKey
}

// WithSubLineItemID adds the subLineItemID to the cart update sub line item recipient params
func (o *CartUpdateSubLineItemRecipientParams) WithSubLineItemID(subLineItemID string) *CartUpdateSubLineItemRecipientParams {
	o.SetSubLineItemID(subLineItemID)
	return o
}

// SetSubLineItemID adds the subLineItemId to the cart update sub line item recipient params
func (o *CartUpdateSubLineItemRecipientParams) SetSubLineItemID(subLineItemID string) {
	o.SubLineItemID = subLineItemID
}

// WriteToRequest writes these params to a swagger request
func (o *CartUpdateSubLineItemRecipientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
