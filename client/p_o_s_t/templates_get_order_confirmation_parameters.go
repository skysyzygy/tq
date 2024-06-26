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

// NewTemplatesGetOrderConfirmationParams creates a new TemplatesGetOrderConfirmationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTemplatesGetOrderConfirmationParams() *TemplatesGetOrderConfirmationParams {
	return &TemplatesGetOrderConfirmationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTemplatesGetOrderConfirmationParamsWithTimeout creates a new TemplatesGetOrderConfirmationParams object
// with the ability to set a timeout on a request.
func NewTemplatesGetOrderConfirmationParamsWithTimeout(timeout time.Duration) *TemplatesGetOrderConfirmationParams {
	return &TemplatesGetOrderConfirmationParams{
		timeout: timeout,
	}
}

// NewTemplatesGetOrderConfirmationParamsWithContext creates a new TemplatesGetOrderConfirmationParams object
// with the ability to set a context for a request.
func NewTemplatesGetOrderConfirmationParamsWithContext(ctx context.Context) *TemplatesGetOrderConfirmationParams {
	return &TemplatesGetOrderConfirmationParams{
		Context: ctx,
	}
}

// NewTemplatesGetOrderConfirmationParamsWithHTTPClient creates a new TemplatesGetOrderConfirmationParams object
// with the ability to set a custom HTTPClient for a request.
func NewTemplatesGetOrderConfirmationParamsWithHTTPClient(client *http.Client) *TemplatesGetOrderConfirmationParams {
	return &TemplatesGetOrderConfirmationParams{
		HTTPClient: client,
	}
}

/*
TemplatesGetOrderConfirmationParams contains all the parameters to send to the API endpoint

	for the templates get order confirmation operation.

	Typically these are written to a http.Request.
*/
type TemplatesGetOrderConfirmationParams struct {

	/* OrderID.

	   Id of the order to retrieve information from
	*/
	OrderID string

	/* Request.

	   A list of name/value pairs to be used in the template
	*/
	Request []*models.NameValue

	/* TemplateID.

	   Id of the template used to format the html
	*/
	TemplateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the templates get order confirmation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TemplatesGetOrderConfirmationParams) WithDefaults() *TemplatesGetOrderConfirmationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the templates get order confirmation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TemplatesGetOrderConfirmationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the templates get order confirmation params
func (o *TemplatesGetOrderConfirmationParams) WithTimeout(timeout time.Duration) *TemplatesGetOrderConfirmationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the templates get order confirmation params
func (o *TemplatesGetOrderConfirmationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the templates get order confirmation params
func (o *TemplatesGetOrderConfirmationParams) WithContext(ctx context.Context) *TemplatesGetOrderConfirmationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the templates get order confirmation params
func (o *TemplatesGetOrderConfirmationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the templates get order confirmation params
func (o *TemplatesGetOrderConfirmationParams) WithHTTPClient(client *http.Client) *TemplatesGetOrderConfirmationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the templates get order confirmation params
func (o *TemplatesGetOrderConfirmationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderID adds the orderID to the templates get order confirmation params
func (o *TemplatesGetOrderConfirmationParams) WithOrderID(orderID string) *TemplatesGetOrderConfirmationParams {
	o.SetOrderID(orderID)
	return o
}

// SetOrderID adds the orderId to the templates get order confirmation params
func (o *TemplatesGetOrderConfirmationParams) SetOrderID(orderID string) {
	o.OrderID = orderID
}

// WithRequest adds the request to the templates get order confirmation params
func (o *TemplatesGetOrderConfirmationParams) WithRequest(request []*models.NameValue) *TemplatesGetOrderConfirmationParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the templates get order confirmation params
func (o *TemplatesGetOrderConfirmationParams) SetRequest(request []*models.NameValue) {
	o.Request = request
}

// WithTemplateID adds the templateID to the templates get order confirmation params
func (o *TemplatesGetOrderConfirmationParams) WithTemplateID(templateID string) *TemplatesGetOrderConfirmationParams {
	o.SetTemplateID(templateID)
	return o
}

// SetTemplateID adds the templateId to the templates get order confirmation params
func (o *TemplatesGetOrderConfirmationParams) SetTemplateID(templateID string) {
	o.TemplateID = templateID
}

// WriteToRequest writes these params to a swagger request
func (o *TemplatesGetOrderConfirmationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orderId
	if err := r.SetPathParam("orderId", o.OrderID); err != nil {
		return err
	}
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	// path param templateId
	if err := r.SetPathParam("templateId", o.TemplateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
