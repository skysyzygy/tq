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

// NewEmailProfilesGetSummariesParams creates a new EmailProfilesGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEmailProfilesGetSummariesParams() *EmailProfilesGetSummariesParams {
	return &EmailProfilesGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEmailProfilesGetSummariesParamsWithTimeout creates a new EmailProfilesGetSummariesParams object
// with the ability to set a timeout on a request.
func NewEmailProfilesGetSummariesParamsWithTimeout(timeout time.Duration) *EmailProfilesGetSummariesParams {
	return &EmailProfilesGetSummariesParams{
		timeout: timeout,
	}
}

// NewEmailProfilesGetSummariesParamsWithContext creates a new EmailProfilesGetSummariesParams object
// with the ability to set a context for a request.
func NewEmailProfilesGetSummariesParamsWithContext(ctx context.Context) *EmailProfilesGetSummariesParams {
	return &EmailProfilesGetSummariesParams{
		Context: ctx,
	}
}

// NewEmailProfilesGetSummariesParamsWithHTTPClient creates a new EmailProfilesGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewEmailProfilesGetSummariesParamsWithHTTPClient(client *http.Client) *EmailProfilesGetSummariesParams {
	return &EmailProfilesGetSummariesParams{
		HTTPClient: client,
	}
}

/*
EmailProfilesGetSummariesParams contains all the parameters to send to the API endpoint

	for the email profiles get summaries operation.

	Typically these are written to a http.Request.
*/
type EmailProfilesGetSummariesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the email profiles get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmailProfilesGetSummariesParams) WithDefaults() *EmailProfilesGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the email profiles get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmailProfilesGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the email profiles get summaries params
func (o *EmailProfilesGetSummariesParams) WithTimeout(timeout time.Duration) *EmailProfilesGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the email profiles get summaries params
func (o *EmailProfilesGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the email profiles get summaries params
func (o *EmailProfilesGetSummariesParams) WithContext(ctx context.Context) *EmailProfilesGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the email profiles get summaries params
func (o *EmailProfilesGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the email profiles get summaries params
func (o *EmailProfilesGetSummariesParams) WithHTTPClient(client *http.Client) *EmailProfilesGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the email profiles get summaries params
func (o *EmailProfilesGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *EmailProfilesGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
