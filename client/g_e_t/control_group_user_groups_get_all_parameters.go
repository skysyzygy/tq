// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

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

// NewControlGroupUserGroupsGetAllParams creates a new ControlGroupUserGroupsGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewControlGroupUserGroupsGetAllParams() *ControlGroupUserGroupsGetAllParams {
	return &ControlGroupUserGroupsGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewControlGroupUserGroupsGetAllParamsWithTimeout creates a new ControlGroupUserGroupsGetAllParams object
// with the ability to set a timeout on a request.
func NewControlGroupUserGroupsGetAllParamsWithTimeout(timeout time.Duration) *ControlGroupUserGroupsGetAllParams {
	return &ControlGroupUserGroupsGetAllParams{
		timeout: timeout,
	}
}

// NewControlGroupUserGroupsGetAllParamsWithContext creates a new ControlGroupUserGroupsGetAllParams object
// with the ability to set a context for a request.
func NewControlGroupUserGroupsGetAllParamsWithContext(ctx context.Context) *ControlGroupUserGroupsGetAllParams {
	return &ControlGroupUserGroupsGetAllParams{
		Context: ctx,
	}
}

// NewControlGroupUserGroupsGetAllParamsWithHTTPClient creates a new ControlGroupUserGroupsGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewControlGroupUserGroupsGetAllParamsWithHTTPClient(client *http.Client) *ControlGroupUserGroupsGetAllParams {
	return &ControlGroupUserGroupsGetAllParams{
		HTTPClient: client,
	}
}

/*
ControlGroupUserGroupsGetAllParams contains all the parameters to send to the API endpoint

	for the control group user groups get all operation.

	Typically these are written to a http.Request.
*/
type ControlGroupUserGroupsGetAllParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the control group user groups get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ControlGroupUserGroupsGetAllParams) WithDefaults() *ControlGroupUserGroupsGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the control group user groups get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ControlGroupUserGroupsGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the control group user groups get all params
func (o *ControlGroupUserGroupsGetAllParams) WithTimeout(timeout time.Duration) *ControlGroupUserGroupsGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the control group user groups get all params
func (o *ControlGroupUserGroupsGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the control group user groups get all params
func (o *ControlGroupUserGroupsGetAllParams) WithContext(ctx context.Context) *ControlGroupUserGroupsGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the control group user groups get all params
func (o *ControlGroupUserGroupsGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the control group user groups get all params
func (o *ControlGroupUserGroupsGetAllParams) WithHTTPClient(client *http.Client) *ControlGroupUserGroupsGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the control group user groups get all params
func (o *ControlGroupUserGroupsGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ControlGroupUserGroupsGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
