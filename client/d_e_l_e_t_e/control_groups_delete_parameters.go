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

// NewControlGroupsDeleteParams creates a new ControlGroupsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewControlGroupsDeleteParams() *ControlGroupsDeleteParams {
	return &ControlGroupsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewControlGroupsDeleteParamsWithTimeout creates a new ControlGroupsDeleteParams object
// with the ability to set a timeout on a request.
func NewControlGroupsDeleteParamsWithTimeout(timeout time.Duration) *ControlGroupsDeleteParams {
	return &ControlGroupsDeleteParams{
		timeout: timeout,
	}
}

// NewControlGroupsDeleteParamsWithContext creates a new ControlGroupsDeleteParams object
// with the ability to set a context for a request.
func NewControlGroupsDeleteParamsWithContext(ctx context.Context) *ControlGroupsDeleteParams {
	return &ControlGroupsDeleteParams{
		Context: ctx,
	}
}

// NewControlGroupsDeleteParamsWithHTTPClient creates a new ControlGroupsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewControlGroupsDeleteParamsWithHTTPClient(client *http.Client) *ControlGroupsDeleteParams {
	return &ControlGroupsDeleteParams{
		HTTPClient: client,
	}
}

/*
ControlGroupsDeleteParams contains all the parameters to send to the API endpoint

	for the control groups delete operation.

	Typically these are written to a http.Request.
*/
type ControlGroupsDeleteParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the control groups delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ControlGroupsDeleteParams) WithDefaults() *ControlGroupsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the control groups delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ControlGroupsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the control groups delete params
func (o *ControlGroupsDeleteParams) WithTimeout(timeout time.Duration) *ControlGroupsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the control groups delete params
func (o *ControlGroupsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the control groups delete params
func (o *ControlGroupsDeleteParams) WithContext(ctx context.Context) *ControlGroupsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the control groups delete params
func (o *ControlGroupsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the control groups delete params
func (o *ControlGroupsDeleteParams) WithHTTPClient(client *http.Client) *ControlGroupsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the control groups delete params
func (o *ControlGroupsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the control groups delete params
func (o *ControlGroupsDeleteParams) WithID(id string) *ControlGroupsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the control groups delete params
func (o *ControlGroupsDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ControlGroupsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
