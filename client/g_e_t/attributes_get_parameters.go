// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

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

// NewAttributesGetParams creates a new AttributesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAttributesGetParams() *AttributesGetParams {
	return &AttributesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAttributesGetParamsWithTimeout creates a new AttributesGetParams object
// with the ability to set a timeout on a request.
func NewAttributesGetParamsWithTimeout(timeout time.Duration) *AttributesGetParams {
	return &AttributesGetParams{
		timeout: timeout,
	}
}

// NewAttributesGetParamsWithContext creates a new AttributesGetParams object
// with the ability to set a context for a request.
func NewAttributesGetParamsWithContext(ctx context.Context) *AttributesGetParams {
	return &AttributesGetParams{
		Context: ctx,
	}
}

// NewAttributesGetParamsWithHTTPClient creates a new AttributesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewAttributesGetParamsWithHTTPClient(client *http.Client) *AttributesGetParams {
	return &AttributesGetParams{
		HTTPClient: client,
	}
}

/*
AttributesGetParams contains all the parameters to send to the API endpoint

	for the attributes get operation.

	Typically these are written to a http.Request.
*/
type AttributesGetParams struct {

	// AttributeID.
	AttributeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the attributes get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AttributesGetParams) WithDefaults() *AttributesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the attributes get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AttributesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the attributes get params
func (o *AttributesGetParams) WithTimeout(timeout time.Duration) *AttributesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the attributes get params
func (o *AttributesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the attributes get params
func (o *AttributesGetParams) WithContext(ctx context.Context) *AttributesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the attributes get params
func (o *AttributesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the attributes get params
func (o *AttributesGetParams) WithHTTPClient(client *http.Client) *AttributesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the attributes get params
func (o *AttributesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttributeID adds the attributeID to the attributes get params
func (o *AttributesGetParams) WithAttributeID(attributeID string) *AttributesGetParams {
	o.SetAttributeID(attributeID)
	return o
}

// SetAttributeID adds the attributeId to the attributes get params
func (o *AttributesGetParams) SetAttributeID(attributeID string) {
	o.AttributeID = attributeID
}

// WriteToRequest writes these params to a swagger request
func (o *AttributesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param attributeId
	if err := r.SetPathParam("attributeId", o.AttributeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
