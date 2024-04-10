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

// NewDiscountTypesUpdateParams creates a new DiscountTypesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDiscountTypesUpdateParams() *DiscountTypesUpdateParams {
	return &DiscountTypesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDiscountTypesUpdateParamsWithTimeout creates a new DiscountTypesUpdateParams object
// with the ability to set a timeout on a request.
func NewDiscountTypesUpdateParamsWithTimeout(timeout time.Duration) *DiscountTypesUpdateParams {
	return &DiscountTypesUpdateParams{
		timeout: timeout,
	}
}

// NewDiscountTypesUpdateParamsWithContext creates a new DiscountTypesUpdateParams object
// with the ability to set a context for a request.
func NewDiscountTypesUpdateParamsWithContext(ctx context.Context) *DiscountTypesUpdateParams {
	return &DiscountTypesUpdateParams{
		Context: ctx,
	}
}

// NewDiscountTypesUpdateParamsWithHTTPClient creates a new DiscountTypesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDiscountTypesUpdateParamsWithHTTPClient(client *http.Client) *DiscountTypesUpdateParams {
	return &DiscountTypesUpdateParams{
		HTTPClient: client,
	}
}

/*
DiscountTypesUpdateParams contains all the parameters to send to the API endpoint

	for the discount types update operation.

	Typically these are written to a http.Request.
*/
type DiscountTypesUpdateParams struct {

	// Data.
	Data *models.DiscountType

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the discount types update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DiscountTypesUpdateParams) WithDefaults() *DiscountTypesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the discount types update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DiscountTypesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the discount types update params
func (o *DiscountTypesUpdateParams) WithTimeout(timeout time.Duration) *DiscountTypesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the discount types update params
func (o *DiscountTypesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the discount types update params
func (o *DiscountTypesUpdateParams) WithContext(ctx context.Context) *DiscountTypesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the discount types update params
func (o *DiscountTypesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the discount types update params
func (o *DiscountTypesUpdateParams) WithHTTPClient(client *http.Client) *DiscountTypesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the discount types update params
func (o *DiscountTypesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the discount types update params
func (o *DiscountTypesUpdateParams) WithData(data *models.DiscountType) *DiscountTypesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the discount types update params
func (o *DiscountTypesUpdateParams) SetData(data *models.DiscountType) {
	o.Data = data
}

// WithID adds the id to the discount types update params
func (o *DiscountTypesUpdateParams) WithID(id string) *DiscountTypesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the discount types update params
func (o *DiscountTypesUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DiscountTypesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
