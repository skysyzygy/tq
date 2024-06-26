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

// NewTemplatesCreateParams creates a new TemplatesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTemplatesCreateParams() *TemplatesCreateParams {
	return &TemplatesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTemplatesCreateParamsWithTimeout creates a new TemplatesCreateParams object
// with the ability to set a timeout on a request.
func NewTemplatesCreateParamsWithTimeout(timeout time.Duration) *TemplatesCreateParams {
	return &TemplatesCreateParams{
		timeout: timeout,
	}
}

// NewTemplatesCreateParamsWithContext creates a new TemplatesCreateParams object
// with the ability to set a context for a request.
func NewTemplatesCreateParamsWithContext(ctx context.Context) *TemplatesCreateParams {
	return &TemplatesCreateParams{
		Context: ctx,
	}
}

// NewTemplatesCreateParamsWithHTTPClient creates a new TemplatesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTemplatesCreateParamsWithHTTPClient(client *http.Client) *TemplatesCreateParams {
	return &TemplatesCreateParams{
		HTTPClient: client,
	}
}

/*
TemplatesCreateParams contains all the parameters to send to the API endpoint

	for the templates create operation.

	Typically these are written to a http.Request.
*/
type TemplatesCreateParams struct {

	// Template.
	Template *models.Template

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the templates create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TemplatesCreateParams) WithDefaults() *TemplatesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the templates create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TemplatesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the templates create params
func (o *TemplatesCreateParams) WithTimeout(timeout time.Duration) *TemplatesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the templates create params
func (o *TemplatesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the templates create params
func (o *TemplatesCreateParams) WithContext(ctx context.Context) *TemplatesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the templates create params
func (o *TemplatesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the templates create params
func (o *TemplatesCreateParams) WithHTTPClient(client *http.Client) *TemplatesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the templates create params
func (o *TemplatesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTemplate adds the template to the templates create params
func (o *TemplatesCreateParams) WithTemplate(template *models.Template) *TemplatesCreateParams {
	o.SetTemplate(template)
	return o
}

// SetTemplate adds the template to the templates create params
func (o *TemplatesCreateParams) SetTemplate(template *models.Template) {
	o.Template = template
}

// WriteToRequest writes these params to a swagger request
func (o *TemplatesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Template != nil {
		if err := r.SetBodyParam(o.Template); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
