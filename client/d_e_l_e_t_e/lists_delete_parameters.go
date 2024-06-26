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

// NewListsDeleteParams creates a new ListsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListsDeleteParams() *ListsDeleteParams {
	return &ListsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListsDeleteParamsWithTimeout creates a new ListsDeleteParams object
// with the ability to set a timeout on a request.
func NewListsDeleteParamsWithTimeout(timeout time.Duration) *ListsDeleteParams {
	return &ListsDeleteParams{
		timeout: timeout,
	}
}

// NewListsDeleteParamsWithContext creates a new ListsDeleteParams object
// with the ability to set a context for a request.
func NewListsDeleteParamsWithContext(ctx context.Context) *ListsDeleteParams {
	return &ListsDeleteParams{
		Context: ctx,
	}
}

// NewListsDeleteParamsWithHTTPClient creates a new ListsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewListsDeleteParamsWithHTTPClient(client *http.Client) *ListsDeleteParams {
	return &ListsDeleteParams{
		HTTPClient: client,
	}
}

/*
ListsDeleteParams contains all the parameters to send to the API endpoint

	for the lists delete operation.

	Typically these are written to a http.Request.
*/
type ListsDeleteParams struct {

	// ListID.
	ListID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lists delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListsDeleteParams) WithDefaults() *ListsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lists delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the lists delete params
func (o *ListsDeleteParams) WithTimeout(timeout time.Duration) *ListsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lists delete params
func (o *ListsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lists delete params
func (o *ListsDeleteParams) WithContext(ctx context.Context) *ListsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lists delete params
func (o *ListsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lists delete params
func (o *ListsDeleteParams) WithHTTPClient(client *http.Client) *ListsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lists delete params
func (o *ListsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithListID adds the listID to the lists delete params
func (o *ListsDeleteParams) WithListID(listID string) *ListsDeleteParams {
	o.SetListID(listID)
	return o
}

// SetListID adds the listId to the lists delete params
func (o *ListsDeleteParams) SetListID(listID string) {
	o.ListID = listID
}

// WriteToRequest writes these params to a swagger request
func (o *ListsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param listId
	if err := r.SetPathParam("listId", o.ListID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
