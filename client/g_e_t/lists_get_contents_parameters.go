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

// NewListsGetContentsParams creates a new ListsGetContentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListsGetContentsParams() *ListsGetContentsParams {
	return &ListsGetContentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListsGetContentsParamsWithTimeout creates a new ListsGetContentsParams object
// with the ability to set a timeout on a request.
func NewListsGetContentsParamsWithTimeout(timeout time.Duration) *ListsGetContentsParams {
	return &ListsGetContentsParams{
		timeout: timeout,
	}
}

// NewListsGetContentsParamsWithContext creates a new ListsGetContentsParams object
// with the ability to set a context for a request.
func NewListsGetContentsParamsWithContext(ctx context.Context) *ListsGetContentsParams {
	return &ListsGetContentsParams{
		Context: ctx,
	}
}

// NewListsGetContentsParamsWithHTTPClient creates a new ListsGetContentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListsGetContentsParamsWithHTTPClient(client *http.Client) *ListsGetContentsParams {
	return &ListsGetContentsParams{
		HTTPClient: client,
	}
}

/*
ListsGetContentsParams contains all the parameters to send to the API endpoint

	for the lists get contents operation.

	Typically these are written to a http.Request.
*/
type ListsGetContentsParams struct {

	/* ListID.

	   Required, id of list for contents
	*/
	ListID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lists get contents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListsGetContentsParams) WithDefaults() *ListsGetContentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lists get contents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListsGetContentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the lists get contents params
func (o *ListsGetContentsParams) WithTimeout(timeout time.Duration) *ListsGetContentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lists get contents params
func (o *ListsGetContentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lists get contents params
func (o *ListsGetContentsParams) WithContext(ctx context.Context) *ListsGetContentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lists get contents params
func (o *ListsGetContentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lists get contents params
func (o *ListsGetContentsParams) WithHTTPClient(client *http.Client) *ListsGetContentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lists get contents params
func (o *ListsGetContentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithListID adds the listID to the lists get contents params
func (o *ListsGetContentsParams) WithListID(listID string) *ListsGetContentsParams {
	o.SetListID(listID)
	return o
}

// SetListID adds the listId to the lists get contents params
func (o *ListsGetContentsParams) SetListID(listID string) {
	o.ListID = listID
}

// WriteToRequest writes these params to a swagger request
func (o *ListsGetContentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
