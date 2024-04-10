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

// NewSecurityServiceResourcesGetAllParams creates a new SecurityServiceResourcesGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityServiceResourcesGetAllParams() *SecurityServiceResourcesGetAllParams {
	return &SecurityServiceResourcesGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityServiceResourcesGetAllParamsWithTimeout creates a new SecurityServiceResourcesGetAllParams object
// with the ability to set a timeout on a request.
func NewSecurityServiceResourcesGetAllParamsWithTimeout(timeout time.Duration) *SecurityServiceResourcesGetAllParams {
	return &SecurityServiceResourcesGetAllParams{
		timeout: timeout,
	}
}

// NewSecurityServiceResourcesGetAllParamsWithContext creates a new SecurityServiceResourcesGetAllParams object
// with the ability to set a context for a request.
func NewSecurityServiceResourcesGetAllParamsWithContext(ctx context.Context) *SecurityServiceResourcesGetAllParams {
	return &SecurityServiceResourcesGetAllParams{
		Context: ctx,
	}
}

// NewSecurityServiceResourcesGetAllParamsWithHTTPClient creates a new SecurityServiceResourcesGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityServiceResourcesGetAllParamsWithHTTPClient(client *http.Client) *SecurityServiceResourcesGetAllParams {
	return &SecurityServiceResourcesGetAllParams{
		HTTPClient: client,
	}
}

/*
SecurityServiceResourcesGetAllParams contains all the parameters to send to the API endpoint

	for the security service resources get all operation.

	Typically these are written to a http.Request.
*/
type SecurityServiceResourcesGetAllParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security service resources get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityServiceResourcesGetAllParams) WithDefaults() *SecurityServiceResourcesGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security service resources get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityServiceResourcesGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the security service resources get all params
func (o *SecurityServiceResourcesGetAllParams) WithTimeout(timeout time.Duration) *SecurityServiceResourcesGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security service resources get all params
func (o *SecurityServiceResourcesGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security service resources get all params
func (o *SecurityServiceResourcesGetAllParams) WithContext(ctx context.Context) *SecurityServiceResourcesGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security service resources get all params
func (o *SecurityServiceResourcesGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security service resources get all params
func (o *SecurityServiceResourcesGetAllParams) WithHTTPClient(client *http.Client) *SecurityServiceResourcesGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security service resources get all params
func (o *SecurityServiceResourcesGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityServiceResourcesGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
