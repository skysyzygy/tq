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

// NewInventoryWebContentsGetAllParams creates a new InventoryWebContentsGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInventoryWebContentsGetAllParams() *InventoryWebContentsGetAllParams {
	return &InventoryWebContentsGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInventoryWebContentsGetAllParamsWithTimeout creates a new InventoryWebContentsGetAllParams object
// with the ability to set a timeout on a request.
func NewInventoryWebContentsGetAllParamsWithTimeout(timeout time.Duration) *InventoryWebContentsGetAllParams {
	return &InventoryWebContentsGetAllParams{
		timeout: timeout,
	}
}

// NewInventoryWebContentsGetAllParamsWithContext creates a new InventoryWebContentsGetAllParams object
// with the ability to set a context for a request.
func NewInventoryWebContentsGetAllParamsWithContext(ctx context.Context) *InventoryWebContentsGetAllParams {
	return &InventoryWebContentsGetAllParams{
		Context: ctx,
	}
}

// NewInventoryWebContentsGetAllParamsWithHTTPClient creates a new InventoryWebContentsGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewInventoryWebContentsGetAllParamsWithHTTPClient(client *http.Client) *InventoryWebContentsGetAllParams {
	return &InventoryWebContentsGetAllParams{
		HTTPClient: client,
	}
}

/*
InventoryWebContentsGetAllParams contains all the parameters to send to the API endpoint

	for the inventory web contents get all operation.

	Typically these are written to a http.Request.
*/
type InventoryWebContentsGetAllParams struct {

	/* ContentTypeIds.

	   A comma separated list of content type ids.  Required if productionElementIds is not used.
	*/
	ContentTypeIds *string

	/* ProductionElementIds.

	   A comma separated list of production element ids (Title, Production, Production Season, or Performance). Required if contentTypeIds is not used.
	*/
	ProductionElementIds *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the inventory web contents get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InventoryWebContentsGetAllParams) WithDefaults() *InventoryWebContentsGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the inventory web contents get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InventoryWebContentsGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the inventory web contents get all params
func (o *InventoryWebContentsGetAllParams) WithTimeout(timeout time.Duration) *InventoryWebContentsGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inventory web contents get all params
func (o *InventoryWebContentsGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inventory web contents get all params
func (o *InventoryWebContentsGetAllParams) WithContext(ctx context.Context) *InventoryWebContentsGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inventory web contents get all params
func (o *InventoryWebContentsGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inventory web contents get all params
func (o *InventoryWebContentsGetAllParams) WithHTTPClient(client *http.Client) *InventoryWebContentsGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inventory web contents get all params
func (o *InventoryWebContentsGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentTypeIds adds the contentTypeIds to the inventory web contents get all params
func (o *InventoryWebContentsGetAllParams) WithContentTypeIds(contentTypeIds *string) *InventoryWebContentsGetAllParams {
	o.SetContentTypeIds(contentTypeIds)
	return o
}

// SetContentTypeIds adds the contentTypeIds to the inventory web contents get all params
func (o *InventoryWebContentsGetAllParams) SetContentTypeIds(contentTypeIds *string) {
	o.ContentTypeIds = contentTypeIds
}

// WithProductionElementIds adds the productionElementIds to the inventory web contents get all params
func (o *InventoryWebContentsGetAllParams) WithProductionElementIds(productionElementIds *string) *InventoryWebContentsGetAllParams {
	o.SetProductionElementIds(productionElementIds)
	return o
}

// SetProductionElementIds adds the productionElementIds to the inventory web contents get all params
func (o *InventoryWebContentsGetAllParams) SetProductionElementIds(productionElementIds *string) {
	o.ProductionElementIds = productionElementIds
}

// WriteToRequest writes these params to a swagger request
func (o *InventoryWebContentsGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentTypeIds != nil {

		// query param contentTypeIds
		var qrContentTypeIds string

		if o.ContentTypeIds != nil {
			qrContentTypeIds = *o.ContentTypeIds
		}
		qContentTypeIds := qrContentTypeIds
		if qContentTypeIds != "" {

			if err := r.SetQueryParam("contentTypeIds", qContentTypeIds); err != nil {
				return err
			}
		}
	}

	if o.ProductionElementIds != nil {

		// query param productionElementIds
		var qrProductionElementIds string

		if o.ProductionElementIds != nil {
			qrProductionElementIds = *o.ProductionElementIds
		}
		qProductionElementIds := qrProductionElementIds
		if qProductionElementIds != "" {

			if err := r.SetQueryParam("productionElementIds", qProductionElementIds); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
