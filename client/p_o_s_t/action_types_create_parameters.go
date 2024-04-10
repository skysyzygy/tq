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

// NewActionTypesCreateParams creates a new ActionTypesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewActionTypesCreateParams() *ActionTypesCreateParams {
	return &ActionTypesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewActionTypesCreateParamsWithTimeout creates a new ActionTypesCreateParams object
// with the ability to set a timeout on a request.
func NewActionTypesCreateParamsWithTimeout(timeout time.Duration) *ActionTypesCreateParams {
	return &ActionTypesCreateParams{
		timeout: timeout,
	}
}

// NewActionTypesCreateParamsWithContext creates a new ActionTypesCreateParams object
// with the ability to set a context for a request.
func NewActionTypesCreateParamsWithContext(ctx context.Context) *ActionTypesCreateParams {
	return &ActionTypesCreateParams{
		Context: ctx,
	}
}

// NewActionTypesCreateParamsWithHTTPClient creates a new ActionTypesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewActionTypesCreateParamsWithHTTPClient(client *http.Client) *ActionTypesCreateParams {
	return &ActionTypesCreateParams{
		HTTPClient: client,
	}
}

/*
ActionTypesCreateParams contains all the parameters to send to the API endpoint

	for the action types create operation.

	Typically these are written to a http.Request.
*/
type ActionTypesCreateParams struct {

	// Data.
	Data *models.ActionType

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the action types create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActionTypesCreateParams) WithDefaults() *ActionTypesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the action types create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActionTypesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the action types create params
func (o *ActionTypesCreateParams) WithTimeout(timeout time.Duration) *ActionTypesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the action types create params
func (o *ActionTypesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the action types create params
func (o *ActionTypesCreateParams) WithContext(ctx context.Context) *ActionTypesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the action types create params
func (o *ActionTypesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the action types create params
func (o *ActionTypesCreateParams) WithHTTPClient(client *http.Client) *ActionTypesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the action types create params
func (o *ActionTypesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the action types create params
func (o *ActionTypesCreateParams) WithData(data *models.ActionType) *ActionTypesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the action types create params
func (o *ActionTypesCreateParams) SetData(data *models.ActionType) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ActionTypesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
