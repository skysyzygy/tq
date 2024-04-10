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

// NewReceiptSettingsUpdateParams creates a new ReceiptSettingsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReceiptSettingsUpdateParams() *ReceiptSettingsUpdateParams {
	return &ReceiptSettingsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReceiptSettingsUpdateParamsWithTimeout creates a new ReceiptSettingsUpdateParams object
// with the ability to set a timeout on a request.
func NewReceiptSettingsUpdateParamsWithTimeout(timeout time.Duration) *ReceiptSettingsUpdateParams {
	return &ReceiptSettingsUpdateParams{
		timeout: timeout,
	}
}

// NewReceiptSettingsUpdateParamsWithContext creates a new ReceiptSettingsUpdateParams object
// with the ability to set a context for a request.
func NewReceiptSettingsUpdateParamsWithContext(ctx context.Context) *ReceiptSettingsUpdateParams {
	return &ReceiptSettingsUpdateParams{
		Context: ctx,
	}
}

// NewReceiptSettingsUpdateParamsWithHTTPClient creates a new ReceiptSettingsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewReceiptSettingsUpdateParamsWithHTTPClient(client *http.Client) *ReceiptSettingsUpdateParams {
	return &ReceiptSettingsUpdateParams{
		HTTPClient: client,
	}
}

/*
ReceiptSettingsUpdateParams contains all the parameters to send to the API endpoint

	for the receipt settings update operation.

	Typically these are written to a http.Request.
*/
type ReceiptSettingsUpdateParams struct {

	/* Data.

	   The updated resource to be saved
	*/
	Data *models.ReceiptSetting

	/* ID.

	   The id of the resource to be updated
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the receipt settings update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReceiptSettingsUpdateParams) WithDefaults() *ReceiptSettingsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the receipt settings update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReceiptSettingsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the receipt settings update params
func (o *ReceiptSettingsUpdateParams) WithTimeout(timeout time.Duration) *ReceiptSettingsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the receipt settings update params
func (o *ReceiptSettingsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the receipt settings update params
func (o *ReceiptSettingsUpdateParams) WithContext(ctx context.Context) *ReceiptSettingsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the receipt settings update params
func (o *ReceiptSettingsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the receipt settings update params
func (o *ReceiptSettingsUpdateParams) WithHTTPClient(client *http.Client) *ReceiptSettingsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the receipt settings update params
func (o *ReceiptSettingsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the receipt settings update params
func (o *ReceiptSettingsUpdateParams) WithData(data *models.ReceiptSetting) *ReceiptSettingsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the receipt settings update params
func (o *ReceiptSettingsUpdateParams) SetData(data *models.ReceiptSetting) {
	o.Data = data
}

// WithID adds the id to the receipt settings update params
func (o *ReceiptSettingsUpdateParams) WithID(id string) *ReceiptSettingsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the receipt settings update params
func (o *ReceiptSettingsUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReceiptSettingsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
