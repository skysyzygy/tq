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

// NewActionTypesDeleteParams creates a new ActionTypesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewActionTypesDeleteParams() *ActionTypesDeleteParams {
	return &ActionTypesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewActionTypesDeleteParamsWithTimeout creates a new ActionTypesDeleteParams object
// with the ability to set a timeout on a request.
func NewActionTypesDeleteParamsWithTimeout(timeout time.Duration) *ActionTypesDeleteParams {
	return &ActionTypesDeleteParams{
		timeout: timeout,
	}
}

// NewActionTypesDeleteParamsWithContext creates a new ActionTypesDeleteParams object
// with the ability to set a context for a request.
func NewActionTypesDeleteParamsWithContext(ctx context.Context) *ActionTypesDeleteParams {
	return &ActionTypesDeleteParams{
		Context: ctx,
	}
}

// NewActionTypesDeleteParamsWithHTTPClient creates a new ActionTypesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewActionTypesDeleteParamsWithHTTPClient(client *http.Client) *ActionTypesDeleteParams {
	return &ActionTypesDeleteParams{
		HTTPClient: client,
	}
}

/*
ActionTypesDeleteParams contains all the parameters to send to the API endpoint

	for the action types delete operation.

	Typically these are written to a http.Request.
*/
type ActionTypesDeleteParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the action types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActionTypesDeleteParams) WithDefaults() *ActionTypesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the action types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActionTypesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the action types delete params
func (o *ActionTypesDeleteParams) WithTimeout(timeout time.Duration) *ActionTypesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the action types delete params
func (o *ActionTypesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the action types delete params
func (o *ActionTypesDeleteParams) WithContext(ctx context.Context) *ActionTypesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the action types delete params
func (o *ActionTypesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the action types delete params
func (o *ActionTypesDeleteParams) WithHTTPClient(client *http.Client) *ActionTypesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the action types delete params
func (o *ActionTypesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the action types delete params
func (o *ActionTypesDeleteParams) WithID(id string) *ActionTypesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the action types delete params
func (o *ActionTypesDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ActionTypesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
