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

// NewListCategoriesDeleteParams creates a new ListCategoriesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListCategoriesDeleteParams() *ListCategoriesDeleteParams {
	return &ListCategoriesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListCategoriesDeleteParamsWithTimeout creates a new ListCategoriesDeleteParams object
// with the ability to set a timeout on a request.
func NewListCategoriesDeleteParamsWithTimeout(timeout time.Duration) *ListCategoriesDeleteParams {
	return &ListCategoriesDeleteParams{
		timeout: timeout,
	}
}

// NewListCategoriesDeleteParamsWithContext creates a new ListCategoriesDeleteParams object
// with the ability to set a context for a request.
func NewListCategoriesDeleteParamsWithContext(ctx context.Context) *ListCategoriesDeleteParams {
	return &ListCategoriesDeleteParams{
		Context: ctx,
	}
}

// NewListCategoriesDeleteParamsWithHTTPClient creates a new ListCategoriesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewListCategoriesDeleteParamsWithHTTPClient(client *http.Client) *ListCategoriesDeleteParams {
	return &ListCategoriesDeleteParams{
		HTTPClient: client,
	}
}

/*
ListCategoriesDeleteParams contains all the parameters to send to the API endpoint

	for the list categories delete operation.

	Typically these are written to a http.Request.
*/
type ListCategoriesDeleteParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list categories delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListCategoriesDeleteParams) WithDefaults() *ListCategoriesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list categories delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListCategoriesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list categories delete params
func (o *ListCategoriesDeleteParams) WithTimeout(timeout time.Duration) *ListCategoriesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list categories delete params
func (o *ListCategoriesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list categories delete params
func (o *ListCategoriesDeleteParams) WithContext(ctx context.Context) *ListCategoriesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list categories delete params
func (o *ListCategoriesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list categories delete params
func (o *ListCategoriesDeleteParams) WithHTTPClient(client *http.Client) *ListCategoriesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list categories delete params
func (o *ListCategoriesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the list categories delete params
func (o *ListCategoriesDeleteParams) WithID(id string) *ListCategoriesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the list categories delete params
func (o *ListCategoriesDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ListCategoriesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
