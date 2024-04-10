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

// NewTriPOSCloudConfigurationsCreateParams creates a new TriPOSCloudConfigurationsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTriPOSCloudConfigurationsCreateParams() *TriPOSCloudConfigurationsCreateParams {
	return &TriPOSCloudConfigurationsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTriPOSCloudConfigurationsCreateParamsWithTimeout creates a new TriPOSCloudConfigurationsCreateParams object
// with the ability to set a timeout on a request.
func NewTriPOSCloudConfigurationsCreateParamsWithTimeout(timeout time.Duration) *TriPOSCloudConfigurationsCreateParams {
	return &TriPOSCloudConfigurationsCreateParams{
		timeout: timeout,
	}
}

// NewTriPOSCloudConfigurationsCreateParamsWithContext creates a new TriPOSCloudConfigurationsCreateParams object
// with the ability to set a context for a request.
func NewTriPOSCloudConfigurationsCreateParamsWithContext(ctx context.Context) *TriPOSCloudConfigurationsCreateParams {
	return &TriPOSCloudConfigurationsCreateParams{
		Context: ctx,
	}
}

// NewTriPOSCloudConfigurationsCreateParamsWithHTTPClient creates a new TriPOSCloudConfigurationsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTriPOSCloudConfigurationsCreateParamsWithHTTPClient(client *http.Client) *TriPOSCloudConfigurationsCreateParams {
	return &TriPOSCloudConfigurationsCreateParams{
		HTTPClient: client,
	}
}

/*
TriPOSCloudConfigurationsCreateParams contains all the parameters to send to the API endpoint

	for the tri p o s cloud configurations create operation.

	Typically these are written to a http.Request.
*/
type TriPOSCloudConfigurationsCreateParams struct {

	// Data.
	Data *models.TriPOSCloudConfiguration

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tri p o s cloud configurations create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TriPOSCloudConfigurationsCreateParams) WithDefaults() *TriPOSCloudConfigurationsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tri p o s cloud configurations create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TriPOSCloudConfigurationsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tri p o s cloud configurations create params
func (o *TriPOSCloudConfigurationsCreateParams) WithTimeout(timeout time.Duration) *TriPOSCloudConfigurationsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tri p o s cloud configurations create params
func (o *TriPOSCloudConfigurationsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tri p o s cloud configurations create params
func (o *TriPOSCloudConfigurationsCreateParams) WithContext(ctx context.Context) *TriPOSCloudConfigurationsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tri p o s cloud configurations create params
func (o *TriPOSCloudConfigurationsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tri p o s cloud configurations create params
func (o *TriPOSCloudConfigurationsCreateParams) WithHTTPClient(client *http.Client) *TriPOSCloudConfigurationsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tri p o s cloud configurations create params
func (o *TriPOSCloudConfigurationsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the tri p o s cloud configurations create params
func (o *TriPOSCloudConfigurationsCreateParams) WithData(data *models.TriPOSCloudConfiguration) *TriPOSCloudConfigurationsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the tri p o s cloud configurations create params
func (o *TriPOSCloudConfigurationsCreateParams) SetData(data *models.TriPOSCloudConfiguration) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *TriPOSCloudConfigurationsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
