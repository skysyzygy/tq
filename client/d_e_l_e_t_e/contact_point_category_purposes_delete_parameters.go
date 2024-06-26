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

// NewContactPointCategoryPurposesDeleteParams creates a new ContactPointCategoryPurposesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContactPointCategoryPurposesDeleteParams() *ContactPointCategoryPurposesDeleteParams {
	return &ContactPointCategoryPurposesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContactPointCategoryPurposesDeleteParamsWithTimeout creates a new ContactPointCategoryPurposesDeleteParams object
// with the ability to set a timeout on a request.
func NewContactPointCategoryPurposesDeleteParamsWithTimeout(timeout time.Duration) *ContactPointCategoryPurposesDeleteParams {
	return &ContactPointCategoryPurposesDeleteParams{
		timeout: timeout,
	}
}

// NewContactPointCategoryPurposesDeleteParamsWithContext creates a new ContactPointCategoryPurposesDeleteParams object
// with the ability to set a context for a request.
func NewContactPointCategoryPurposesDeleteParamsWithContext(ctx context.Context) *ContactPointCategoryPurposesDeleteParams {
	return &ContactPointCategoryPurposesDeleteParams{
		Context: ctx,
	}
}

// NewContactPointCategoryPurposesDeleteParamsWithHTTPClient creates a new ContactPointCategoryPurposesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewContactPointCategoryPurposesDeleteParamsWithHTTPClient(client *http.Client) *ContactPointCategoryPurposesDeleteParams {
	return &ContactPointCategoryPurposesDeleteParams{
		HTTPClient: client,
	}
}

/*
ContactPointCategoryPurposesDeleteParams contains all the parameters to send to the API endpoint

	for the contact point category purposes delete operation.

	Typically these are written to a http.Request.
*/
type ContactPointCategoryPurposesDeleteParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the contact point category purposes delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactPointCategoryPurposesDeleteParams) WithDefaults() *ContactPointCategoryPurposesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the contact point category purposes delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactPointCategoryPurposesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the contact point category purposes delete params
func (o *ContactPointCategoryPurposesDeleteParams) WithTimeout(timeout time.Duration) *ContactPointCategoryPurposesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contact point category purposes delete params
func (o *ContactPointCategoryPurposesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contact point category purposes delete params
func (o *ContactPointCategoryPurposesDeleteParams) WithContext(ctx context.Context) *ContactPointCategoryPurposesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contact point category purposes delete params
func (o *ContactPointCategoryPurposesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contact point category purposes delete params
func (o *ContactPointCategoryPurposesDeleteParams) WithHTTPClient(client *http.Client) *ContactPointCategoryPurposesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contact point category purposes delete params
func (o *ContactPointCategoryPurposesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the contact point category purposes delete params
func (o *ContactPointCategoryPurposesDeleteParams) WithID(id string) *ContactPointCategoryPurposesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the contact point category purposes delete params
func (o *ContactPointCategoryPurposesDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ContactPointCategoryPurposesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
