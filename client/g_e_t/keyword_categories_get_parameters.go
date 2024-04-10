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

// NewKeywordCategoriesGetParams creates a new KeywordCategoriesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKeywordCategoriesGetParams() *KeywordCategoriesGetParams {
	return &KeywordCategoriesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKeywordCategoriesGetParamsWithTimeout creates a new KeywordCategoriesGetParams object
// with the ability to set a timeout on a request.
func NewKeywordCategoriesGetParamsWithTimeout(timeout time.Duration) *KeywordCategoriesGetParams {
	return &KeywordCategoriesGetParams{
		timeout: timeout,
	}
}

// NewKeywordCategoriesGetParamsWithContext creates a new KeywordCategoriesGetParams object
// with the ability to set a context for a request.
func NewKeywordCategoriesGetParamsWithContext(ctx context.Context) *KeywordCategoriesGetParams {
	return &KeywordCategoriesGetParams{
		Context: ctx,
	}
}

// NewKeywordCategoriesGetParamsWithHTTPClient creates a new KeywordCategoriesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewKeywordCategoriesGetParamsWithHTTPClient(client *http.Client) *KeywordCategoriesGetParams {
	return &KeywordCategoriesGetParams{
		HTTPClient: client,
	}
}

/*
KeywordCategoriesGetParams contains all the parameters to send to the API endpoint

	for the keyword categories get operation.

	Typically these are written to a http.Request.
*/
type KeywordCategoriesGetParams struct {

	// ID.
	ID string

	/* MaintenanceMode.

	   Ignore control grouping (default: false)
	*/
	MaintenanceMode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the keyword categories get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KeywordCategoriesGetParams) WithDefaults() *KeywordCategoriesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the keyword categories get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KeywordCategoriesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the keyword categories get params
func (o *KeywordCategoriesGetParams) WithTimeout(timeout time.Duration) *KeywordCategoriesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the keyword categories get params
func (o *KeywordCategoriesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the keyword categories get params
func (o *KeywordCategoriesGetParams) WithContext(ctx context.Context) *KeywordCategoriesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the keyword categories get params
func (o *KeywordCategoriesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the keyword categories get params
func (o *KeywordCategoriesGetParams) WithHTTPClient(client *http.Client) *KeywordCategoriesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the keyword categories get params
func (o *KeywordCategoriesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the keyword categories get params
func (o *KeywordCategoriesGetParams) WithID(id string) *KeywordCategoriesGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the keyword categories get params
func (o *KeywordCategoriesGetParams) SetID(id string) {
	o.ID = id
}

// WithMaintenanceMode adds the maintenanceMode to the keyword categories get params
func (o *KeywordCategoriesGetParams) WithMaintenanceMode(maintenanceMode *string) *KeywordCategoriesGetParams {
	o.SetMaintenanceMode(maintenanceMode)
	return o
}

// SetMaintenanceMode adds the maintenanceMode to the keyword categories get params
func (o *KeywordCategoriesGetParams) SetMaintenanceMode(maintenanceMode *string) {
	o.MaintenanceMode = maintenanceMode
}

// WriteToRequest writes these params to a swagger request
func (o *KeywordCategoriesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.MaintenanceMode != nil {

		// query param maintenanceMode
		var qrMaintenanceMode string

		if o.MaintenanceMode != nil {
			qrMaintenanceMode = *o.MaintenanceMode
		}
		qMaintenanceMode := qrMaintenanceMode
		if qMaintenanceMode != "" {

			if err := r.SetQueryParam("maintenanceMode", qMaintenanceMode); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
