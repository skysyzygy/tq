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

// NewCartUpdateLineItemSpecialRequestParams creates a new CartUpdateLineItemSpecialRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCartUpdateLineItemSpecialRequestParams() *CartUpdateLineItemSpecialRequestParams {
	return &CartUpdateLineItemSpecialRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCartUpdateLineItemSpecialRequestParamsWithTimeout creates a new CartUpdateLineItemSpecialRequestParams object
// with the ability to set a timeout on a request.
func NewCartUpdateLineItemSpecialRequestParamsWithTimeout(timeout time.Duration) *CartUpdateLineItemSpecialRequestParams {
	return &CartUpdateLineItemSpecialRequestParams{
		timeout: timeout,
	}
}

// NewCartUpdateLineItemSpecialRequestParamsWithContext creates a new CartUpdateLineItemSpecialRequestParams object
// with the ability to set a context for a request.
func NewCartUpdateLineItemSpecialRequestParamsWithContext(ctx context.Context) *CartUpdateLineItemSpecialRequestParams {
	return &CartUpdateLineItemSpecialRequestParams{
		Context: ctx,
	}
}

// NewCartUpdateLineItemSpecialRequestParamsWithHTTPClient creates a new CartUpdateLineItemSpecialRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewCartUpdateLineItemSpecialRequestParamsWithHTTPClient(client *http.Client) *CartUpdateLineItemSpecialRequestParams {
	return &CartUpdateLineItemSpecialRequestParams{
		HTTPClient: client,
	}
}

/*
CartUpdateLineItemSpecialRequestParams contains all the parameters to send to the API endpoint

	for the cart update line item special request operation.

	Typically these are written to a http.Request.
*/
type CartUpdateLineItemSpecialRequestParams struct {

	/* LineItemID.

	   The lineitem ID to update the special request info for
	*/
	LineItemID string

	// Request.
	Request *models.SpecialRequest

	// SessionKey.
	SessionKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cart update line item special request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartUpdateLineItemSpecialRequestParams) WithDefaults() *CartUpdateLineItemSpecialRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cart update line item special request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CartUpdateLineItemSpecialRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cart update line item special request params
func (o *CartUpdateLineItemSpecialRequestParams) WithTimeout(timeout time.Duration) *CartUpdateLineItemSpecialRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cart update line item special request params
func (o *CartUpdateLineItemSpecialRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cart update line item special request params
func (o *CartUpdateLineItemSpecialRequestParams) WithContext(ctx context.Context) *CartUpdateLineItemSpecialRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cart update line item special request params
func (o *CartUpdateLineItemSpecialRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cart update line item special request params
func (o *CartUpdateLineItemSpecialRequestParams) WithHTTPClient(client *http.Client) *CartUpdateLineItemSpecialRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cart update line item special request params
func (o *CartUpdateLineItemSpecialRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLineItemID adds the lineItemID to the cart update line item special request params
func (o *CartUpdateLineItemSpecialRequestParams) WithLineItemID(lineItemID string) *CartUpdateLineItemSpecialRequestParams {
	o.SetLineItemID(lineItemID)
	return o
}

// SetLineItemID adds the lineItemId to the cart update line item special request params
func (o *CartUpdateLineItemSpecialRequestParams) SetLineItemID(lineItemID string) {
	o.LineItemID = lineItemID
}

// WithRequest adds the request to the cart update line item special request params
func (o *CartUpdateLineItemSpecialRequestParams) WithRequest(request *models.SpecialRequest) *CartUpdateLineItemSpecialRequestParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the cart update line item special request params
func (o *CartUpdateLineItemSpecialRequestParams) SetRequest(request *models.SpecialRequest) {
	o.Request = request
}

// WithSessionKey adds the sessionKey to the cart update line item special request params
func (o *CartUpdateLineItemSpecialRequestParams) WithSessionKey(sessionKey string) *CartUpdateLineItemSpecialRequestParams {
	o.SetSessionKey(sessionKey)
	return o
}

// SetSessionKey adds the sessionKey to the cart update line item special request params
func (o *CartUpdateLineItemSpecialRequestParams) SetSessionKey(sessionKey string) {
	o.SessionKey = sessionKey
}

// WriteToRequest writes these params to a swagger request
func (o *CartUpdateLineItemSpecialRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param lineItemId
	if err := r.SetPathParam("lineItemId", o.LineItemID); err != nil {
		return err
	}
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