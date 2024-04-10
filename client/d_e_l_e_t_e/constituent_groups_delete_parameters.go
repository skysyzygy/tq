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

// NewConstituentGroupsDeleteParams creates a new ConstituentGroupsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConstituentGroupsDeleteParams() *ConstituentGroupsDeleteParams {
	return &ConstituentGroupsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConstituentGroupsDeleteParamsWithTimeout creates a new ConstituentGroupsDeleteParams object
// with the ability to set a timeout on a request.
func NewConstituentGroupsDeleteParamsWithTimeout(timeout time.Duration) *ConstituentGroupsDeleteParams {
	return &ConstituentGroupsDeleteParams{
		timeout: timeout,
	}
}

// NewConstituentGroupsDeleteParamsWithContext creates a new ConstituentGroupsDeleteParams object
// with the ability to set a context for a request.
func NewConstituentGroupsDeleteParamsWithContext(ctx context.Context) *ConstituentGroupsDeleteParams {
	return &ConstituentGroupsDeleteParams{
		Context: ctx,
	}
}

// NewConstituentGroupsDeleteParamsWithHTTPClient creates a new ConstituentGroupsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewConstituentGroupsDeleteParamsWithHTTPClient(client *http.Client) *ConstituentGroupsDeleteParams {
	return &ConstituentGroupsDeleteParams{
		HTTPClient: client,
	}
}

/*
ConstituentGroupsDeleteParams contains all the parameters to send to the API endpoint

	for the constituent groups delete operation.

	Typically these are written to a http.Request.
*/
type ConstituentGroupsDeleteParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the constituent groups delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConstituentGroupsDeleteParams) WithDefaults() *ConstituentGroupsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the constituent groups delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConstituentGroupsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the constituent groups delete params
func (o *ConstituentGroupsDeleteParams) WithTimeout(timeout time.Duration) *ConstituentGroupsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the constituent groups delete params
func (o *ConstituentGroupsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the constituent groups delete params
func (o *ConstituentGroupsDeleteParams) WithContext(ctx context.Context) *ConstituentGroupsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the constituent groups delete params
func (o *ConstituentGroupsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the constituent groups delete params
func (o *ConstituentGroupsDeleteParams) WithHTTPClient(client *http.Client) *ConstituentGroupsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the constituent groups delete params
func (o *ConstituentGroupsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the constituent groups delete params
func (o *ConstituentGroupsDeleteParams) WithID(id string) *ConstituentGroupsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the constituent groups delete params
func (o *ConstituentGroupsDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ConstituentGroupsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
