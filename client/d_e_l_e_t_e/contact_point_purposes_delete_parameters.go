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

// NewContactPointPurposesDeleteParams creates a new ContactPointPurposesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContactPointPurposesDeleteParams() *ContactPointPurposesDeleteParams {
	return &ContactPointPurposesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContactPointPurposesDeleteParamsWithTimeout creates a new ContactPointPurposesDeleteParams object
// with the ability to set a timeout on a request.
func NewContactPointPurposesDeleteParamsWithTimeout(timeout time.Duration) *ContactPointPurposesDeleteParams {
	return &ContactPointPurposesDeleteParams{
		timeout: timeout,
	}
}

// NewContactPointPurposesDeleteParamsWithContext creates a new ContactPointPurposesDeleteParams object
// with the ability to set a context for a request.
func NewContactPointPurposesDeleteParamsWithContext(ctx context.Context) *ContactPointPurposesDeleteParams {
	return &ContactPointPurposesDeleteParams{
		Context: ctx,
	}
}

// NewContactPointPurposesDeleteParamsWithHTTPClient creates a new ContactPointPurposesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewContactPointPurposesDeleteParamsWithHTTPClient(client *http.Client) *ContactPointPurposesDeleteParams {
	return &ContactPointPurposesDeleteParams{
		HTTPClient: client,
	}
}

/*
ContactPointPurposesDeleteParams contains all the parameters to send to the API endpoint

	for the contact point purposes delete operation.

	Typically these are written to a http.Request.
*/
type ContactPointPurposesDeleteParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the contact point purposes delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactPointPurposesDeleteParams) WithDefaults() *ContactPointPurposesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the contact point purposes delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContactPointPurposesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the contact point purposes delete params
func (o *ContactPointPurposesDeleteParams) WithTimeout(timeout time.Duration) *ContactPointPurposesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contact point purposes delete params
func (o *ContactPointPurposesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contact point purposes delete params
func (o *ContactPointPurposesDeleteParams) WithContext(ctx context.Context) *ContactPointPurposesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contact point purposes delete params
func (o *ContactPointPurposesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contact point purposes delete params
func (o *ContactPointPurposesDeleteParams) WithHTTPClient(client *http.Client) *ContactPointPurposesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contact point purposes delete params
func (o *ContactPointPurposesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the contact point purposes delete params
func (o *ContactPointPurposesDeleteParams) WithID(id string) *ContactPointPurposesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the contact point purposes delete params
func (o *ContactPointPurposesDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ContactPointPurposesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
