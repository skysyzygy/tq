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

// NewQualificationsDeleteParams creates a new QualificationsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQualificationsDeleteParams() *QualificationsDeleteParams {
	return &QualificationsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQualificationsDeleteParamsWithTimeout creates a new QualificationsDeleteParams object
// with the ability to set a timeout on a request.
func NewQualificationsDeleteParamsWithTimeout(timeout time.Duration) *QualificationsDeleteParams {
	return &QualificationsDeleteParams{
		timeout: timeout,
	}
}

// NewQualificationsDeleteParamsWithContext creates a new QualificationsDeleteParams object
// with the ability to set a context for a request.
func NewQualificationsDeleteParamsWithContext(ctx context.Context) *QualificationsDeleteParams {
	return &QualificationsDeleteParams{
		Context: ctx,
	}
}

// NewQualificationsDeleteParamsWithHTTPClient creates a new QualificationsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewQualificationsDeleteParamsWithHTTPClient(client *http.Client) *QualificationsDeleteParams {
	return &QualificationsDeleteParams{
		HTTPClient: client,
	}
}

/*
QualificationsDeleteParams contains all the parameters to send to the API endpoint

	for the qualifications delete operation.

	Typically these are written to a http.Request.
*/
type QualificationsDeleteParams struct {

	/* ID.

	   The id of the resource to be deleted
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the qualifications delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QualificationsDeleteParams) WithDefaults() *QualificationsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the qualifications delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QualificationsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the qualifications delete params
func (o *QualificationsDeleteParams) WithTimeout(timeout time.Duration) *QualificationsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the qualifications delete params
func (o *QualificationsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the qualifications delete params
func (o *QualificationsDeleteParams) WithContext(ctx context.Context) *QualificationsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the qualifications delete params
func (o *QualificationsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the qualifications delete params
func (o *QualificationsDeleteParams) WithHTTPClient(client *http.Client) *QualificationsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the qualifications delete params
func (o *QualificationsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the qualifications delete params
func (o *QualificationsDeleteParams) WithID(id string) *QualificationsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the qualifications delete params
func (o *QualificationsDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *QualificationsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
