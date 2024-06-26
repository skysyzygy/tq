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

// NewIssuesDeleteParams creates a new IssuesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIssuesDeleteParams() *IssuesDeleteParams {
	return &IssuesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIssuesDeleteParamsWithTimeout creates a new IssuesDeleteParams object
// with the ability to set a timeout on a request.
func NewIssuesDeleteParamsWithTimeout(timeout time.Duration) *IssuesDeleteParams {
	return &IssuesDeleteParams{
		timeout: timeout,
	}
}

// NewIssuesDeleteParamsWithContext creates a new IssuesDeleteParams object
// with the ability to set a context for a request.
func NewIssuesDeleteParamsWithContext(ctx context.Context) *IssuesDeleteParams {
	return &IssuesDeleteParams{
		Context: ctx,
	}
}

// NewIssuesDeleteParamsWithHTTPClient creates a new IssuesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIssuesDeleteParamsWithHTTPClient(client *http.Client) *IssuesDeleteParams {
	return &IssuesDeleteParams{
		HTTPClient: client,
	}
}

/*
IssuesDeleteParams contains all the parameters to send to the API endpoint

	for the issues delete operation.

	Typically these are written to a http.Request.
*/
type IssuesDeleteParams struct {

	// IssueID.
	IssueID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the issues delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssuesDeleteParams) WithDefaults() *IssuesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the issues delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssuesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the issues delete params
func (o *IssuesDeleteParams) WithTimeout(timeout time.Duration) *IssuesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issues delete params
func (o *IssuesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issues delete params
func (o *IssuesDeleteParams) WithContext(ctx context.Context) *IssuesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issues delete params
func (o *IssuesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issues delete params
func (o *IssuesDeleteParams) WithHTTPClient(client *http.Client) *IssuesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issues delete params
func (o *IssuesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIssueID adds the issueID to the issues delete params
func (o *IssuesDeleteParams) WithIssueID(issueID string) *IssuesDeleteParams {
	o.SetIssueID(issueID)
	return o
}

// SetIssueID adds the issueId to the issues delete params
func (o *IssuesDeleteParams) SetIssueID(issueID string) {
	o.IssueID = issueID
}

// WriteToRequest writes these params to a swagger request
func (o *IssuesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param issueId
	if err := r.SetPathParam("issueId", o.IssueID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
