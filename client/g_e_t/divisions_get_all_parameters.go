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

// NewDivisionsGetAllParams creates a new DivisionsGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDivisionsGetAllParams() *DivisionsGetAllParams {
	return &DivisionsGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDivisionsGetAllParamsWithTimeout creates a new DivisionsGetAllParams object
// with the ability to set a timeout on a request.
func NewDivisionsGetAllParamsWithTimeout(timeout time.Duration) *DivisionsGetAllParams {
	return &DivisionsGetAllParams{
		timeout: timeout,
	}
}

// NewDivisionsGetAllParamsWithContext creates a new DivisionsGetAllParams object
// with the ability to set a context for a request.
func NewDivisionsGetAllParamsWithContext(ctx context.Context) *DivisionsGetAllParams {
	return &DivisionsGetAllParams{
		Context: ctx,
	}
}

// NewDivisionsGetAllParamsWithHTTPClient creates a new DivisionsGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewDivisionsGetAllParamsWithHTTPClient(client *http.Client) *DivisionsGetAllParams {
	return &DivisionsGetAllParams{
		HTTPClient: client,
	}
}

/*
DivisionsGetAllParams contains all the parameters to send to the API endpoint

	for the divisions get all operation.

	Typically these are written to a http.Request.
*/
type DivisionsGetAllParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the divisions get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DivisionsGetAllParams) WithDefaults() *DivisionsGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the divisions get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DivisionsGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the divisions get all params
func (o *DivisionsGetAllParams) WithTimeout(timeout time.Duration) *DivisionsGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the divisions get all params
func (o *DivisionsGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the divisions get all params
func (o *DivisionsGetAllParams) WithContext(ctx context.Context) *DivisionsGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the divisions get all params
func (o *DivisionsGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the divisions get all params
func (o *DivisionsGetAllParams) WithHTTPClient(client *http.Client) *DivisionsGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the divisions get all params
func (o *DivisionsGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DivisionsGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
