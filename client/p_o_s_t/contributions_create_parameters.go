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

// NewContributionsCreateParams creates a new ContributionsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContributionsCreateParams() *ContributionsCreateParams {
	return &ContributionsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContributionsCreateParamsWithTimeout creates a new ContributionsCreateParams object
// with the ability to set a timeout on a request.
func NewContributionsCreateParamsWithTimeout(timeout time.Duration) *ContributionsCreateParams {
	return &ContributionsCreateParams{
		timeout: timeout,
	}
}

// NewContributionsCreateParamsWithContext creates a new ContributionsCreateParams object
// with the ability to set a context for a request.
func NewContributionsCreateParamsWithContext(ctx context.Context) *ContributionsCreateParams {
	return &ContributionsCreateParams{
		Context: ctx,
	}
}

// NewContributionsCreateParamsWithHTTPClient creates a new ContributionsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewContributionsCreateParamsWithHTTPClient(client *http.Client) *ContributionsCreateParams {
	return &ContributionsCreateParams{
		HTTPClient: client,
	}
}

/*
ContributionsCreateParams contains all the parameters to send to the API endpoint

	for the contributions create operation.

	Typically these are written to a http.Request.
*/
type ContributionsCreateParams struct {

	/* Contribution.

	   The contribution entity. Only the Id property is provided at this point.
	*/
	Contribution *models.Contribution

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the contributions create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContributionsCreateParams) WithDefaults() *ContributionsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the contributions create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContributionsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the contributions create params
func (o *ContributionsCreateParams) WithTimeout(timeout time.Duration) *ContributionsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contributions create params
func (o *ContributionsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contributions create params
func (o *ContributionsCreateParams) WithContext(ctx context.Context) *ContributionsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contributions create params
func (o *ContributionsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contributions create params
func (o *ContributionsCreateParams) WithHTTPClient(client *http.Client) *ContributionsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contributions create params
func (o *ContributionsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContribution adds the contribution to the contributions create params
func (o *ContributionsCreateParams) WithContribution(contribution *models.Contribution) *ContributionsCreateParams {
	o.SetContribution(contribution)
	return o
}

// SetContribution adds the contribution to the contributions create params
func (o *ContributionsCreateParams) SetContribution(contribution *models.Contribution) {
	o.Contribution = contribution
}

// WriteToRequest writes these params to a swagger request
func (o *ContributionsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Contribution != nil {
		if err := r.SetBodyParam(o.Contribution); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
