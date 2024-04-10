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

// NewTemplatePriceTypesUpdateParams creates a new TemplatePriceTypesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTemplatePriceTypesUpdateParams() *TemplatePriceTypesUpdateParams {
	return &TemplatePriceTypesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTemplatePriceTypesUpdateParamsWithTimeout creates a new TemplatePriceTypesUpdateParams object
// with the ability to set a timeout on a request.
func NewTemplatePriceTypesUpdateParamsWithTimeout(timeout time.Duration) *TemplatePriceTypesUpdateParams {
	return &TemplatePriceTypesUpdateParams{
		timeout: timeout,
	}
}

// NewTemplatePriceTypesUpdateParamsWithContext creates a new TemplatePriceTypesUpdateParams object
// with the ability to set a context for a request.
func NewTemplatePriceTypesUpdateParamsWithContext(ctx context.Context) *TemplatePriceTypesUpdateParams {
	return &TemplatePriceTypesUpdateParams{
		Context: ctx,
	}
}

// NewTemplatePriceTypesUpdateParamsWithHTTPClient creates a new TemplatePriceTypesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTemplatePriceTypesUpdateParamsWithHTTPClient(client *http.Client) *TemplatePriceTypesUpdateParams {
	return &TemplatePriceTypesUpdateParams{
		HTTPClient: client,
	}
}

/*
TemplatePriceTypesUpdateParams contains all the parameters to send to the API endpoint

	for the template price types update operation.

	Typically these are written to a http.Request.
*/
type TemplatePriceTypesUpdateParams struct {

	// TemplatePriceType.
	TemplatePriceType *models.TemplatePriceType

	// TemplatePriceTypeID.
	TemplatePriceTypeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the template price types update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TemplatePriceTypesUpdateParams) WithDefaults() *TemplatePriceTypesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the template price types update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TemplatePriceTypesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the template price types update params
func (o *TemplatePriceTypesUpdateParams) WithTimeout(timeout time.Duration) *TemplatePriceTypesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the template price types update params
func (o *TemplatePriceTypesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the template price types update params
func (o *TemplatePriceTypesUpdateParams) WithContext(ctx context.Context) *TemplatePriceTypesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the template price types update params
func (o *TemplatePriceTypesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the template price types update params
func (o *TemplatePriceTypesUpdateParams) WithHTTPClient(client *http.Client) *TemplatePriceTypesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the template price types update params
func (o *TemplatePriceTypesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTemplatePriceType adds the templatePriceType to the template price types update params
func (o *TemplatePriceTypesUpdateParams) WithTemplatePriceType(templatePriceType *models.TemplatePriceType) *TemplatePriceTypesUpdateParams {
	o.SetTemplatePriceType(templatePriceType)
	return o
}

// SetTemplatePriceType adds the templatePriceType to the template price types update params
func (o *TemplatePriceTypesUpdateParams) SetTemplatePriceType(templatePriceType *models.TemplatePriceType) {
	o.TemplatePriceType = templatePriceType
}

// WithTemplatePriceTypeID adds the templatePriceTypeID to the template price types update params
func (o *TemplatePriceTypesUpdateParams) WithTemplatePriceTypeID(templatePriceTypeID string) *TemplatePriceTypesUpdateParams {
	o.SetTemplatePriceTypeID(templatePriceTypeID)
	return o
}

// SetTemplatePriceTypeID adds the templatePriceTypeId to the template price types update params
func (o *TemplatePriceTypesUpdateParams) SetTemplatePriceTypeID(templatePriceTypeID string) {
	o.TemplatePriceTypeID = templatePriceTypeID
}

// WriteToRequest writes these params to a swagger request
func (o *TemplatePriceTypesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.TemplatePriceType != nil {
		if err := r.SetBodyParam(o.TemplatePriceType); err != nil {
			return err
		}
	}

	// path param templatePriceTypeId
	if err := r.SetPathParam("templatePriceTypeId", o.TemplatePriceTypeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
