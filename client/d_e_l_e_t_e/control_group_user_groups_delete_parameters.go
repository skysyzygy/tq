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

// NewControlGroupUserGroupsDeleteParams creates a new ControlGroupUserGroupsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewControlGroupUserGroupsDeleteParams() *ControlGroupUserGroupsDeleteParams {
	return &ControlGroupUserGroupsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewControlGroupUserGroupsDeleteParamsWithTimeout creates a new ControlGroupUserGroupsDeleteParams object
// with the ability to set a timeout on a request.
func NewControlGroupUserGroupsDeleteParamsWithTimeout(timeout time.Duration) *ControlGroupUserGroupsDeleteParams {
	return &ControlGroupUserGroupsDeleteParams{
		timeout: timeout,
	}
}

// NewControlGroupUserGroupsDeleteParamsWithContext creates a new ControlGroupUserGroupsDeleteParams object
// with the ability to set a context for a request.
func NewControlGroupUserGroupsDeleteParamsWithContext(ctx context.Context) *ControlGroupUserGroupsDeleteParams {
	return &ControlGroupUserGroupsDeleteParams{
		Context: ctx,
	}
}

// NewControlGroupUserGroupsDeleteParamsWithHTTPClient creates a new ControlGroupUserGroupsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewControlGroupUserGroupsDeleteParamsWithHTTPClient(client *http.Client) *ControlGroupUserGroupsDeleteParams {
	return &ControlGroupUserGroupsDeleteParams{
		HTTPClient: client,
	}
}

/*
ControlGroupUserGroupsDeleteParams contains all the parameters to send to the API endpoint

	for the control group user groups delete operation.

	Typically these are written to a http.Request.
*/
type ControlGroupUserGroupsDeleteParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the control group user groups delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ControlGroupUserGroupsDeleteParams) WithDefaults() *ControlGroupUserGroupsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the control group user groups delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ControlGroupUserGroupsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the control group user groups delete params
func (o *ControlGroupUserGroupsDeleteParams) WithTimeout(timeout time.Duration) *ControlGroupUserGroupsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the control group user groups delete params
func (o *ControlGroupUserGroupsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the control group user groups delete params
func (o *ControlGroupUserGroupsDeleteParams) WithContext(ctx context.Context) *ControlGroupUserGroupsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the control group user groups delete params
func (o *ControlGroupUserGroupsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the control group user groups delete params
func (o *ControlGroupUserGroupsDeleteParams) WithHTTPClient(client *http.Client) *ControlGroupUserGroupsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the control group user groups delete params
func (o *ControlGroupUserGroupsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the control group user groups delete params
func (o *ControlGroupUserGroupsDeleteParams) WithID(id string) *ControlGroupUserGroupsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the control group user groups delete params
func (o *ControlGroupUserGroupsDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ControlGroupUserGroupsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
