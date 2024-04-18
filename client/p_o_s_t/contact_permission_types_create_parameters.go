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

// NewContactPermissionTypesCreateParams creates a new ContactPermissionTypesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContactPermissionTypesCreateParams() *ContactPermissionTypesCreateParams {
	return &ContactPermissionTypesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContactPermissionTypesCreateParamsWithTimeout creates a new ContactPermissionTypesCreateParams object
// with the ability to set a timeout on a request.
func NewContactPermissionTypesCreateParamsWithTimeout(timeout time.Duration) *ContactPermissionTypesCreateParams {
	return &ContactPermissionTypesCreateParams{
		timeout: timeout,
	}
}

// NewContactPermissionTypesCreateParamsWithContext creates a new ContactPermissionTypesCreateParams object
// with the ability to set a context for a request.
func NewContactPermissionTypesCreateParamsWithContext(ctx context.Context) *ContactPermissionTypesCreateParams {
	return &ContactPermissionTypesCreateParams{
		Context: ctx,
	}
}

// NewContactPermissionTypesCreateParamsWithHTTPClient creates a new ContactPermissionTypesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewContactPermissionTypesCreateParamsWithHTTPClient(client *http.Client) *ContactPermissionTypesCreateParams {
	return &ContactPermissionTypesCreateParams{
		HTTPClient: client,
	}
}

/*
ContactPermissionTypesCreateParams contains all the parameters to send to the API endpoint

	for the contact permission types create operation.

	Typically these are written to a http.Request.
*/
type ContactPermissionTypesCreateParams struct {

	// Data.
	Data *models.ContactPermissionType

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the contact permission types create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactPermissionTypesCreateParams) WithDefaults() *ContactPermissionTypesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the contact permission types create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactPermissionTypesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the contact permission types create params
func (o *ContactPermissionTypesCreateParams) WithTimeout(timeout time.Duration) *ContactPermissionTypesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contact permission types create params
func (o *ContactPermissionTypesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contact permission types create params
func (o *ContactPermissionTypesCreateParams) WithContext(ctx context.Context) *ContactPermissionTypesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contact permission types create params
func (o *ContactPermissionTypesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contact permission types create params
func (o *ContactPermissionTypesCreateParams) WithHTTPClient(client *http.Client) *ContactPermissionTypesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contact permission types create params
func (o *ContactPermissionTypesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the contact permission types create params
func (o *ContactPermissionTypesCreateParams) WithData(data *models.ContactPermissionType) *ContactPermissionTypesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the contact permission types create params
func (o *ContactPermissionTypesCreateParams) SetData(data *models.ContactPermissionType) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ContactPermissionTypesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}