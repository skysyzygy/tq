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

// NewControlGroupUserGroupsCreateParams creates a new ControlGroupUserGroupsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewControlGroupUserGroupsCreateParams() *ControlGroupUserGroupsCreateParams {
	return &ControlGroupUserGroupsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewControlGroupUserGroupsCreateParamsWithTimeout creates a new ControlGroupUserGroupsCreateParams object
// with the ability to set a timeout on a request.
func NewControlGroupUserGroupsCreateParamsWithTimeout(timeout time.Duration) *ControlGroupUserGroupsCreateParams {
	return &ControlGroupUserGroupsCreateParams{
		timeout: timeout,
	}
}

// NewControlGroupUserGroupsCreateParamsWithContext creates a new ControlGroupUserGroupsCreateParams object
// with the ability to set a context for a request.
func NewControlGroupUserGroupsCreateParamsWithContext(ctx context.Context) *ControlGroupUserGroupsCreateParams {
	return &ControlGroupUserGroupsCreateParams{
		Context: ctx,
	}
}

// NewControlGroupUserGroupsCreateParamsWithHTTPClient creates a new ControlGroupUserGroupsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewControlGroupUserGroupsCreateParamsWithHTTPClient(client *http.Client) *ControlGroupUserGroupsCreateParams {
	return &ControlGroupUserGroupsCreateParams{
		HTTPClient: client,
	}
}

/*
ControlGroupUserGroupsCreateParams contains all the parameters to send to the API endpoint

	for the control group user groups create operation.

	Typically these are written to a http.Request.
*/
type ControlGroupUserGroupsCreateParams struct {

	// Data.
	Data *models.ControlGroupUserGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the control group user groups create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ControlGroupUserGroupsCreateParams) WithDefaults() *ControlGroupUserGroupsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the control group user groups create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ControlGroupUserGroupsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the control group user groups create params
func (o *ControlGroupUserGroupsCreateParams) WithTimeout(timeout time.Duration) *ControlGroupUserGroupsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the control group user groups create params
func (o *ControlGroupUserGroupsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the control group user groups create params
func (o *ControlGroupUserGroupsCreateParams) WithContext(ctx context.Context) *ControlGroupUserGroupsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the control group user groups create params
func (o *ControlGroupUserGroupsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the control group user groups create params
func (o *ControlGroupUserGroupsCreateParams) WithHTTPClient(client *http.Client) *ControlGroupUserGroupsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the control group user groups create params
func (o *ControlGroupUserGroupsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the control group user groups create params
func (o *ControlGroupUserGroupsCreateParams) WithData(data *models.ControlGroupUserGroup) *ControlGroupUserGroupsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the control group user groups create params
func (o *ControlGroupUserGroupsCreateParams) SetData(data *models.ControlGroupUserGroup) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ControlGroupUserGroupsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
