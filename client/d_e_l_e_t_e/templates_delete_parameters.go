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

// NewTemplatesDeleteParams creates a new TemplatesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTemplatesDeleteParams() *TemplatesDeleteParams {
	return &TemplatesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTemplatesDeleteParamsWithTimeout creates a new TemplatesDeleteParams object
// with the ability to set a timeout on a request.
func NewTemplatesDeleteParamsWithTimeout(timeout time.Duration) *TemplatesDeleteParams {
	return &TemplatesDeleteParams{
		timeout: timeout,
	}
}

// NewTemplatesDeleteParamsWithContext creates a new TemplatesDeleteParams object
// with the ability to set a context for a request.
func NewTemplatesDeleteParamsWithContext(ctx context.Context) *TemplatesDeleteParams {
	return &TemplatesDeleteParams{
		Context: ctx,
	}
}

// NewTemplatesDeleteParamsWithHTTPClient creates a new TemplatesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewTemplatesDeleteParamsWithHTTPClient(client *http.Client) *TemplatesDeleteParams {
	return &TemplatesDeleteParams{
		HTTPClient: client,
	}
}

/*
TemplatesDeleteParams contains all the parameters to send to the API endpoint

	for the templates delete operation.

	Typically these are written to a http.Request.
*/
type TemplatesDeleteParams struct {

	// TemplateID.
	TemplateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TemplatesDeleteParams) WithDefaults() *TemplatesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TemplatesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the templates delete params
func (o *TemplatesDeleteParams) WithTimeout(timeout time.Duration) *TemplatesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the templates delete params
func (o *TemplatesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the templates delete params
func (o *TemplatesDeleteParams) WithContext(ctx context.Context) *TemplatesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the templates delete params
func (o *TemplatesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the templates delete params
func (o *TemplatesDeleteParams) WithHTTPClient(client *http.Client) *TemplatesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the templates delete params
func (o *TemplatesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTemplateID adds the templateID to the templates delete params
func (o *TemplatesDeleteParams) WithTemplateID(templateID string) *TemplatesDeleteParams {
	o.SetTemplateID(templateID)
	return o
}

// SetTemplateID adds the templateId to the templates delete params
func (o *TemplatesDeleteParams) SetTemplateID(templateID string) {
	o.TemplateID = templateID
}

// WriteToRequest writes these params to a swagger request
func (o *TemplatesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param templateId
	if err := r.SetPathParam("templateId", o.TemplateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
