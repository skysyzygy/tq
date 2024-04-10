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

// NewInventoryWebContentsDeleteParams creates a new InventoryWebContentsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInventoryWebContentsDeleteParams() *InventoryWebContentsDeleteParams {
	return &InventoryWebContentsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInventoryWebContentsDeleteParamsWithTimeout creates a new InventoryWebContentsDeleteParams object
// with the ability to set a timeout on a request.
func NewInventoryWebContentsDeleteParamsWithTimeout(timeout time.Duration) *InventoryWebContentsDeleteParams {
	return &InventoryWebContentsDeleteParams{
		timeout: timeout,
	}
}

// NewInventoryWebContentsDeleteParamsWithContext creates a new InventoryWebContentsDeleteParams object
// with the ability to set a context for a request.
func NewInventoryWebContentsDeleteParamsWithContext(ctx context.Context) *InventoryWebContentsDeleteParams {
	return &InventoryWebContentsDeleteParams{
		Context: ctx,
	}
}

// NewInventoryWebContentsDeleteParamsWithHTTPClient creates a new InventoryWebContentsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewInventoryWebContentsDeleteParamsWithHTTPClient(client *http.Client) *InventoryWebContentsDeleteParams {
	return &InventoryWebContentsDeleteParams{
		HTTPClient: client,
	}
}

/*
InventoryWebContentsDeleteParams contains all the parameters to send to the API endpoint

	for the inventory web contents delete operation.

	Typically these are written to a http.Request.
*/
type InventoryWebContentsDeleteParams struct {

	/* InventoryWebContentID.

	   Id of InventoryWebContent.
	*/
	InventoryWebContentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the inventory web contents delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InventoryWebContentsDeleteParams) WithDefaults() *InventoryWebContentsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the inventory web contents delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InventoryWebContentsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the inventory web contents delete params
func (o *InventoryWebContentsDeleteParams) WithTimeout(timeout time.Duration) *InventoryWebContentsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventory web contents delete params
func (o *InventoryWebContentsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventory web contents delete params
func (o *InventoryWebContentsDeleteParams) WithContext(ctx context.Context) *InventoryWebContentsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventory web contents delete params
func (o *InventoryWebContentsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventory web contents delete params
func (o *InventoryWebContentsDeleteParams) WithHTTPClient(client *http.Client) *InventoryWebContentsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventory web contents delete params
func (o *InventoryWebContentsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInventoryWebContentID adds the inventoryWebContentID to the inventory web contents delete params
func (o *InventoryWebContentsDeleteParams) WithInventoryWebContentID(inventoryWebContentID string) *InventoryWebContentsDeleteParams {
	o.SetInventoryWebContentID(inventoryWebContentID)
	return o
}

// SetInventoryWebContentID adds the inventoryWebContentId to the inventory web contents delete params
func (o *InventoryWebContentsDeleteParams) SetInventoryWebContentID(inventoryWebContentID string) {
	o.InventoryWebContentID = inventoryWebContentID
}

// WriteToRequest writes these params to a swagger request
func (o *InventoryWebContentsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param inventoryWebContentId
	if err := r.SetPathParam("inventoryWebContentId", o.InventoryWebContentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
