// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

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

// NewListsUpdateContentsParams creates a new ListsUpdateContentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListsUpdateContentsParams() *ListsUpdateContentsParams {
	return &ListsUpdateContentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListsUpdateContentsParamsWithTimeout creates a new ListsUpdateContentsParams object
// with the ability to set a timeout on a request.
func NewListsUpdateContentsParamsWithTimeout(timeout time.Duration) *ListsUpdateContentsParams {
	return &ListsUpdateContentsParams{
		timeout: timeout,
	}
}

// NewListsUpdateContentsParamsWithContext creates a new ListsUpdateContentsParams object
// with the ability to set a context for a request.
func NewListsUpdateContentsParamsWithContext(ctx context.Context) *ListsUpdateContentsParams {
	return &ListsUpdateContentsParams{
		Context: ctx,
	}
}

// NewListsUpdateContentsParamsWithHTTPClient creates a new ListsUpdateContentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListsUpdateContentsParamsWithHTTPClient(client *http.Client) *ListsUpdateContentsParams {
	return &ListsUpdateContentsParams{
		HTTPClient: client,
	}
}

/*
ListsUpdateContentsParams contains all the parameters to send to the API endpoint

	for the lists update contents operation.

	Typically these are written to a http.Request.
*/
type ListsUpdateContentsParams struct {

	// ListID.
	ListID string

	// ListImportRequest.
	ListImportRequest *models.ListImportRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lists update contents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListsUpdateContentsParams) WithDefaults() *ListsUpdateContentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lists update contents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListsUpdateContentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the lists update contents params
func (o *ListsUpdateContentsParams) WithTimeout(timeout time.Duration) *ListsUpdateContentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lists update contents params
func (o *ListsUpdateContentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lists update contents params
func (o *ListsUpdateContentsParams) WithContext(ctx context.Context) *ListsUpdateContentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lists update contents params
func (o *ListsUpdateContentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lists update contents params
func (o *ListsUpdateContentsParams) WithHTTPClient(client *http.Client) *ListsUpdateContentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lists update contents params
func (o *ListsUpdateContentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithListID adds the listID to the lists update contents params
func (o *ListsUpdateContentsParams) WithListID(listID string) *ListsUpdateContentsParams {
	o.SetListID(listID)
	return o
}

// SetListID adds the listId to the lists update contents params
func (o *ListsUpdateContentsParams) SetListID(listID string) {
	o.ListID = listID
}

// WithListImportRequest adds the listImportRequest to the lists update contents params
func (o *ListsUpdateContentsParams) WithListImportRequest(listImportRequest *models.ListImportRequest) *ListsUpdateContentsParams {
	o.SetListImportRequest(listImportRequest)
	return o
}

// SetListImportRequest adds the listImportRequest to the lists update contents params
func (o *ListsUpdateContentsParams) SetListImportRequest(listImportRequest *models.ListImportRequest) {
	o.ListImportRequest = listImportRequest
}

// WriteToRequest writes these params to a swagger request
func (o *ListsUpdateContentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param listId
	if err := r.SetPathParam("listId", o.ListID); err != nil {
		return err
	}
	if o.ListImportRequest != nil {
		if err := r.SetBodyParam(o.ListImportRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
