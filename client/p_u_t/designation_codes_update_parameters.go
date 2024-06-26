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

// NewDesignationCodesUpdateParams creates a new DesignationCodesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDesignationCodesUpdateParams() *DesignationCodesUpdateParams {
	return &DesignationCodesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDesignationCodesUpdateParamsWithTimeout creates a new DesignationCodesUpdateParams object
// with the ability to set a timeout on a request.
func NewDesignationCodesUpdateParamsWithTimeout(timeout time.Duration) *DesignationCodesUpdateParams {
	return &DesignationCodesUpdateParams{
		timeout: timeout,
	}
}

// NewDesignationCodesUpdateParamsWithContext creates a new DesignationCodesUpdateParams object
// with the ability to set a context for a request.
func NewDesignationCodesUpdateParamsWithContext(ctx context.Context) *DesignationCodesUpdateParams {
	return &DesignationCodesUpdateParams{
		Context: ctx,
	}
}

// NewDesignationCodesUpdateParamsWithHTTPClient creates a new DesignationCodesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDesignationCodesUpdateParamsWithHTTPClient(client *http.Client) *DesignationCodesUpdateParams {
	return &DesignationCodesUpdateParams{
		HTTPClient: client,
	}
}

/*
DesignationCodesUpdateParams contains all the parameters to send to the API endpoint

	for the designation codes update operation.

	Typically these are written to a http.Request.
*/
type DesignationCodesUpdateParams struct {

	// Data.
	Data *models.DesignationCode

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the designation codes update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DesignationCodesUpdateParams) WithDefaults() *DesignationCodesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the designation codes update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DesignationCodesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the designation codes update params
func (o *DesignationCodesUpdateParams) WithTimeout(timeout time.Duration) *DesignationCodesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the designation codes update params
func (o *DesignationCodesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the designation codes update params
func (o *DesignationCodesUpdateParams) WithContext(ctx context.Context) *DesignationCodesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the designation codes update params
func (o *DesignationCodesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the designation codes update params
func (o *DesignationCodesUpdateParams) WithHTTPClient(client *http.Client) *DesignationCodesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the designation codes update params
func (o *DesignationCodesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the designation codes update params
func (o *DesignationCodesUpdateParams) WithData(data *models.DesignationCode) *DesignationCodesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the designation codes update params
func (o *DesignationCodesUpdateParams) SetData(data *models.DesignationCode) {
	o.Data = data
}

// WithID adds the id to the designation codes update params
func (o *DesignationCodesUpdateParams) WithID(id string) *DesignationCodesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the designation codes update params
func (o *DesignationCodesUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DesignationCodesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
