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

// NewPortfolioCustomElementsGetSummariesParams creates a new PortfolioCustomElementsGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPortfolioCustomElementsGetSummariesParams() *PortfolioCustomElementsGetSummariesParams {
	return &PortfolioCustomElementsGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPortfolioCustomElementsGetSummariesParamsWithTimeout creates a new PortfolioCustomElementsGetSummariesParams object
// with the ability to set a timeout on a request.
func NewPortfolioCustomElementsGetSummariesParamsWithTimeout(timeout time.Duration) *PortfolioCustomElementsGetSummariesParams {
	return &PortfolioCustomElementsGetSummariesParams{
		timeout: timeout,
	}
}

// NewPortfolioCustomElementsGetSummariesParamsWithContext creates a new PortfolioCustomElementsGetSummariesParams object
// with the ability to set a context for a request.
func NewPortfolioCustomElementsGetSummariesParamsWithContext(ctx context.Context) *PortfolioCustomElementsGetSummariesParams {
	return &PortfolioCustomElementsGetSummariesParams{
		Context: ctx,
	}
}

// NewPortfolioCustomElementsGetSummariesParamsWithHTTPClient creates a new PortfolioCustomElementsGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPortfolioCustomElementsGetSummariesParamsWithHTTPClient(client *http.Client) *PortfolioCustomElementsGetSummariesParams {
	return &PortfolioCustomElementsGetSummariesParams{
		HTTPClient: client,
	}
}

/*
PortfolioCustomElementsGetSummariesParams contains all the parameters to send to the API endpoint

	for the portfolio custom elements get summaries operation.

	Typically these are written to a http.Request.
*/
type PortfolioCustomElementsGetSummariesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the portfolio custom elements get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PortfolioCustomElementsGetSummariesParams) WithDefaults() *PortfolioCustomElementsGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the portfolio custom elements get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PortfolioCustomElementsGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the portfolio custom elements get summaries params
func (o *PortfolioCustomElementsGetSummariesParams) WithTimeout(timeout time.Duration) *PortfolioCustomElementsGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the portfolio custom elements get summaries params
func (o *PortfolioCustomElementsGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the portfolio custom elements get summaries params
func (o *PortfolioCustomElementsGetSummariesParams) WithContext(ctx context.Context) *PortfolioCustomElementsGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the portfolio custom elements get summaries params
func (o *PortfolioCustomElementsGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the portfolio custom elements get summaries params
func (o *PortfolioCustomElementsGetSummariesParams) WithHTTPClient(client *http.Client) *PortfolioCustomElementsGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the portfolio custom elements get summaries params
func (o *PortfolioCustomElementsGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PortfolioCustomElementsGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
