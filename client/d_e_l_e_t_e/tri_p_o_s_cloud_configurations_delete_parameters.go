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

// NewTriPOSCloudConfigurationsDeleteParams creates a new TriPOSCloudConfigurationsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTriPOSCloudConfigurationsDeleteParams() *TriPOSCloudConfigurationsDeleteParams {
	return &TriPOSCloudConfigurationsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTriPOSCloudConfigurationsDeleteParamsWithTimeout creates a new TriPOSCloudConfigurationsDeleteParams object
// with the ability to set a timeout on a request.
func NewTriPOSCloudConfigurationsDeleteParamsWithTimeout(timeout time.Duration) *TriPOSCloudConfigurationsDeleteParams {
	return &TriPOSCloudConfigurationsDeleteParams{
		timeout: timeout,
	}
}

// NewTriPOSCloudConfigurationsDeleteParamsWithContext creates a new TriPOSCloudConfigurationsDeleteParams object
// with the ability to set a context for a request.
func NewTriPOSCloudConfigurationsDeleteParamsWithContext(ctx context.Context) *TriPOSCloudConfigurationsDeleteParams {
	return &TriPOSCloudConfigurationsDeleteParams{
		Context: ctx,
	}
}

// NewTriPOSCloudConfigurationsDeleteParamsWithHTTPClient creates a new TriPOSCloudConfigurationsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewTriPOSCloudConfigurationsDeleteParamsWithHTTPClient(client *http.Client) *TriPOSCloudConfigurationsDeleteParams {
	return &TriPOSCloudConfigurationsDeleteParams{
		HTTPClient: client,
	}
}

/*
TriPOSCloudConfigurationsDeleteParams contains all the parameters to send to the API endpoint

	for the tri p o s cloud configurations delete operation.

	Typically these are written to a http.Request.
*/
type TriPOSCloudConfigurationsDeleteParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tri p o s cloud configurations delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TriPOSCloudConfigurationsDeleteParams) WithDefaults() *TriPOSCloudConfigurationsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tri p o s cloud configurations delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TriPOSCloudConfigurationsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tri p o s cloud configurations delete params
func (o *TriPOSCloudConfigurationsDeleteParams) WithTimeout(timeout time.Duration) *TriPOSCloudConfigurationsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tri p o s cloud configurations delete params
func (o *TriPOSCloudConfigurationsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tri p o s cloud configurations delete params
func (o *TriPOSCloudConfigurationsDeleteParams) WithContext(ctx context.Context) *TriPOSCloudConfigurationsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tri p o s cloud configurations delete params
func (o *TriPOSCloudConfigurationsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tri p o s cloud configurations delete params
func (o *TriPOSCloudConfigurationsDeleteParams) WithHTTPClient(client *http.Client) *TriPOSCloudConfigurationsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tri p o s cloud configurations delete params
func (o *TriPOSCloudConfigurationsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the tri p o s cloud configurations delete params
func (o *TriPOSCloudConfigurationsDeleteParams) WithID(id string) *TriPOSCloudConfigurationsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the tri p o s cloud configurations delete params
func (o *TriPOSCloudConfigurationsDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TriPOSCloudConfigurationsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
