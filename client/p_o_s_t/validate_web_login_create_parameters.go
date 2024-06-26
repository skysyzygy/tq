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

// NewValidateWebLoginCreateParams creates a new ValidateWebLoginCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateWebLoginCreateParams() *ValidateWebLoginCreateParams {
	return &ValidateWebLoginCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateWebLoginCreateParamsWithTimeout creates a new ValidateWebLoginCreateParams object
// with the ability to set a timeout on a request.
func NewValidateWebLoginCreateParamsWithTimeout(timeout time.Duration) *ValidateWebLoginCreateParams {
	return &ValidateWebLoginCreateParams{
		timeout: timeout,
	}
}

// NewValidateWebLoginCreateParamsWithContext creates a new ValidateWebLoginCreateParams object
// with the ability to set a context for a request.
func NewValidateWebLoginCreateParamsWithContext(ctx context.Context) *ValidateWebLoginCreateParams {
	return &ValidateWebLoginCreateParams{
		Context: ctx,
	}
}

// NewValidateWebLoginCreateParamsWithHTTPClient creates a new ValidateWebLoginCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateWebLoginCreateParamsWithHTTPClient(client *http.Client) *ValidateWebLoginCreateParams {
	return &ValidateWebLoginCreateParams{
		HTTPClient: client,
	}
}

/*
ValidateWebLoginCreateParams contains all the parameters to send to the API endpoint

	for the validate web login create operation.

	Typically these are written to a http.Request.
*/
type ValidateWebLoginCreateParams struct {

	// ValidationRequest.
	ValidationRequest *models.WebLoginValidationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate web login create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateWebLoginCreateParams) WithDefaults() *ValidateWebLoginCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate web login create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateWebLoginCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate web login create params
func (o *ValidateWebLoginCreateParams) WithTimeout(timeout time.Duration) *ValidateWebLoginCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate web login create params
func (o *ValidateWebLoginCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate web login create params
func (o *ValidateWebLoginCreateParams) WithContext(ctx context.Context) *ValidateWebLoginCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate web login create params
func (o *ValidateWebLoginCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate web login create params
func (o *ValidateWebLoginCreateParams) WithHTTPClient(client *http.Client) *ValidateWebLoginCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate web login create params
func (o *ValidateWebLoginCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithValidationRequest adds the validationRequest to the validate web login create params
func (o *ValidateWebLoginCreateParams) WithValidationRequest(validationRequest *models.WebLoginValidationRequest) *ValidateWebLoginCreateParams {
	o.SetValidationRequest(validationRequest)
	return o
}

// SetValidationRequest adds the validationRequest to the validate web login create params
func (o *ValidateWebLoginCreateParams) SetValidationRequest(validationRequest *models.WebLoginValidationRequest) {
	o.ValidationRequest = validationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateWebLoginCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ValidationRequest != nil {
		if err := r.SetBodyParam(o.ValidationRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
