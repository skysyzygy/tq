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

// NewCustomDefaultCategoriesGetSummariesParams creates a new CustomDefaultCategoriesGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomDefaultCategoriesGetSummariesParams() *CustomDefaultCategoriesGetSummariesParams {
	return &CustomDefaultCategoriesGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomDefaultCategoriesGetSummariesParamsWithTimeout creates a new CustomDefaultCategoriesGetSummariesParams object
// with the ability to set a timeout on a request.
func NewCustomDefaultCategoriesGetSummariesParamsWithTimeout(timeout time.Duration) *CustomDefaultCategoriesGetSummariesParams {
	return &CustomDefaultCategoriesGetSummariesParams{
		timeout: timeout,
	}
}

// NewCustomDefaultCategoriesGetSummariesParamsWithContext creates a new CustomDefaultCategoriesGetSummariesParams object
// with the ability to set a context for a request.
func NewCustomDefaultCategoriesGetSummariesParamsWithContext(ctx context.Context) *CustomDefaultCategoriesGetSummariesParams {
	return &CustomDefaultCategoriesGetSummariesParams{
		Context: ctx,
	}
}

// NewCustomDefaultCategoriesGetSummariesParamsWithHTTPClient creates a new CustomDefaultCategoriesGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomDefaultCategoriesGetSummariesParamsWithHTTPClient(client *http.Client) *CustomDefaultCategoriesGetSummariesParams {
	return &CustomDefaultCategoriesGetSummariesParams{
		HTTPClient: client,
	}
}

/*
CustomDefaultCategoriesGetSummariesParams contains all the parameters to send to the API endpoint

	for the custom default categories get summaries operation.

	Typically these are written to a http.Request.
*/
type CustomDefaultCategoriesGetSummariesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the custom default categories get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomDefaultCategoriesGetSummariesParams) WithDefaults() *CustomDefaultCategoriesGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the custom default categories get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomDefaultCategoriesGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the custom default categories get summaries params
func (o *CustomDefaultCategoriesGetSummariesParams) WithTimeout(timeout time.Duration) *CustomDefaultCategoriesGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom default categories get summaries params
func (o *CustomDefaultCategoriesGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom default categories get summaries params
func (o *CustomDefaultCategoriesGetSummariesParams) WithContext(ctx context.Context) *CustomDefaultCategoriesGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom default categories get summaries params
func (o *CustomDefaultCategoriesGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom default categories get summaries params
func (o *CustomDefaultCategoriesGetSummariesParams) WithHTTPClient(client *http.Client) *CustomDefaultCategoriesGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom default categories get summaries params
func (o *CustomDefaultCategoriesGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CustomDefaultCategoriesGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
