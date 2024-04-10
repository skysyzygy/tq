// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

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

// NewZonesCreateParams creates a new ZonesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewZonesCreateParams() *ZonesCreateParams {
	return &ZonesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewZonesCreateParamsWithTimeout creates a new ZonesCreateParams object
// with the ability to set a timeout on a request.
func NewZonesCreateParamsWithTimeout(timeout time.Duration) *ZonesCreateParams {
	return &ZonesCreateParams{
		timeout: timeout,
	}
}

// NewZonesCreateParamsWithContext creates a new ZonesCreateParams object
// with the ability to set a context for a request.
func NewZonesCreateParamsWithContext(ctx context.Context) *ZonesCreateParams {
	return &ZonesCreateParams{
		Context: ctx,
	}
}

// NewZonesCreateParamsWithHTTPClient creates a new ZonesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewZonesCreateParamsWithHTTPClient(client *http.Client) *ZonesCreateParams {
	return &ZonesCreateParams{
		HTTPClient: client,
	}
}

/*
ZonesCreateParams contains all the parameters to send to the API endpoint

	for the zones create operation.

	Typically these are written to a http.Request.
*/
type ZonesCreateParams struct {

	// Data.
	Data *models.Zone

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the zones create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ZonesCreateParams) WithDefaults() *ZonesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the zones create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ZonesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the zones create params
func (o *ZonesCreateParams) WithTimeout(timeout time.Duration) *ZonesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the zones create params
func (o *ZonesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the zones create params
func (o *ZonesCreateParams) WithContext(ctx context.Context) *ZonesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the zones create params
func (o *ZonesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the zones create params
func (o *ZonesCreateParams) WithHTTPClient(client *http.Client) *ZonesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the zones create params
func (o *ZonesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the zones create params
func (o *ZonesCreateParams) WithData(data *models.Zone) *ZonesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the zones create params
func (o *ZonesCreateParams) SetData(data *models.Zone) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ZonesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
