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

// NewTemplateCategoriesCreateParams creates a new TemplateCategoriesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTemplateCategoriesCreateParams() *TemplateCategoriesCreateParams {
	return &TemplateCategoriesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTemplateCategoriesCreateParamsWithTimeout creates a new TemplateCategoriesCreateParams object
// with the ability to set a timeout on a request.
func NewTemplateCategoriesCreateParamsWithTimeout(timeout time.Duration) *TemplateCategoriesCreateParams {
	return &TemplateCategoriesCreateParams{
		timeout: timeout,
	}
}

// NewTemplateCategoriesCreateParamsWithContext creates a new TemplateCategoriesCreateParams object
// with the ability to set a context for a request.
func NewTemplateCategoriesCreateParamsWithContext(ctx context.Context) *TemplateCategoriesCreateParams {
	return &TemplateCategoriesCreateParams{
		Context: ctx,
	}
}

// NewTemplateCategoriesCreateParamsWithHTTPClient creates a new TemplateCategoriesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTemplateCategoriesCreateParamsWithHTTPClient(client *http.Client) *TemplateCategoriesCreateParams {
	return &TemplateCategoriesCreateParams{
		HTTPClient: client,
	}
}

/*
TemplateCategoriesCreateParams contains all the parameters to send to the API endpoint

	for the template categories create operation.

	Typically these are written to a http.Request.
*/
type TemplateCategoriesCreateParams struct {

	// TemplateCategory.
	TemplateCategory *models.TemplateCategory

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the template categories create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TemplateCategoriesCreateParams) WithDefaults() *TemplateCategoriesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the template categories create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TemplateCategoriesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the template categories create params
func (o *TemplateCategoriesCreateParams) WithTimeout(timeout time.Duration) *TemplateCategoriesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the template categories create params
func (o *TemplateCategoriesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the template categories create params
func (o *TemplateCategoriesCreateParams) WithContext(ctx context.Context) *TemplateCategoriesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the template categories create params
func (o *TemplateCategoriesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the template categories create params
func (o *TemplateCategoriesCreateParams) WithHTTPClient(client *http.Client) *TemplateCategoriesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the template categories create params
func (o *TemplateCategoriesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTemplateCategory adds the templateCategory to the template categories create params
func (o *TemplateCategoriesCreateParams) WithTemplateCategory(templateCategory *models.TemplateCategory) *TemplateCategoriesCreateParams {
	o.SetTemplateCategory(templateCategory)
	return o
}

// SetTemplateCategory adds the templateCategory to the template categories create params
func (o *TemplateCategoriesCreateParams) SetTemplateCategory(templateCategory *models.TemplateCategory) {
	o.TemplateCategory = templateCategory
}

// WriteToRequest writes these params to a swagger request
func (o *TemplateCategoriesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.TemplateCategory != nil {
		if err := r.SetBodyParam(o.TemplateCategory); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
