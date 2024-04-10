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

// NewAppealCategoriesGetSummariesParams creates a new AppealCategoriesGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAppealCategoriesGetSummariesParams() *AppealCategoriesGetSummariesParams {
	return &AppealCategoriesGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAppealCategoriesGetSummariesParamsWithTimeout creates a new AppealCategoriesGetSummariesParams object
// with the ability to set a timeout on a request.
func NewAppealCategoriesGetSummariesParamsWithTimeout(timeout time.Duration) *AppealCategoriesGetSummariesParams {
	return &AppealCategoriesGetSummariesParams{
		timeout: timeout,
	}
}

// NewAppealCategoriesGetSummariesParamsWithContext creates a new AppealCategoriesGetSummariesParams object
// with the ability to set a context for a request.
func NewAppealCategoriesGetSummariesParamsWithContext(ctx context.Context) *AppealCategoriesGetSummariesParams {
	return &AppealCategoriesGetSummariesParams{
		Context: ctx,
	}
}

// NewAppealCategoriesGetSummariesParamsWithHTTPClient creates a new AppealCategoriesGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewAppealCategoriesGetSummariesParamsWithHTTPClient(client *http.Client) *AppealCategoriesGetSummariesParams {
	return &AppealCategoriesGetSummariesParams{
		HTTPClient: client,
	}
}

/*
AppealCategoriesGetSummariesParams contains all the parameters to send to the API endpoint

	for the appeal categories get summaries operation.

	Typically these are written to a http.Request.
*/
type AppealCategoriesGetSummariesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the appeal categories get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppealCategoriesGetSummariesParams) WithDefaults() *AppealCategoriesGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the appeal categories get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppealCategoriesGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the appeal categories get summaries params
func (o *AppealCategoriesGetSummariesParams) WithTimeout(timeout time.Duration) *AppealCategoriesGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the appeal categories get summaries params
func (o *AppealCategoriesGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the appeal categories get summaries params
func (o *AppealCategoriesGetSummariesParams) WithContext(ctx context.Context) *AppealCategoriesGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the appeal categories get summaries params
func (o *AppealCategoriesGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the appeal categories get summaries params
func (o *AppealCategoriesGetSummariesParams) WithHTTPClient(client *http.Client) *AppealCategoriesGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the appeal categories get summaries params
func (o *AppealCategoriesGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AppealCategoriesGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
