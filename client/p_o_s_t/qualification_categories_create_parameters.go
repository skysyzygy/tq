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

// NewQualificationCategoriesCreateParams creates a new QualificationCategoriesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQualificationCategoriesCreateParams() *QualificationCategoriesCreateParams {
	return &QualificationCategoriesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQualificationCategoriesCreateParamsWithTimeout creates a new QualificationCategoriesCreateParams object
// with the ability to set a timeout on a request.
func NewQualificationCategoriesCreateParamsWithTimeout(timeout time.Duration) *QualificationCategoriesCreateParams {
	return &QualificationCategoriesCreateParams{
		timeout: timeout,
	}
}

// NewQualificationCategoriesCreateParamsWithContext creates a new QualificationCategoriesCreateParams object
// with the ability to set a context for a request.
func NewQualificationCategoriesCreateParamsWithContext(ctx context.Context) *QualificationCategoriesCreateParams {
	return &QualificationCategoriesCreateParams{
		Context: ctx,
	}
}

// NewQualificationCategoriesCreateParamsWithHTTPClient creates a new QualificationCategoriesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewQualificationCategoriesCreateParamsWithHTTPClient(client *http.Client) *QualificationCategoriesCreateParams {
	return &QualificationCategoriesCreateParams{
		HTTPClient: client,
	}
}

/*
QualificationCategoriesCreateParams contains all the parameters to send to the API endpoint

	for the qualification categories create operation.

	Typically these are written to a http.Request.
*/
type QualificationCategoriesCreateParams struct {

	/* Data.

	   The resource to be created
	*/
	Data *models.QualificationCategory

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the qualification categories create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QualificationCategoriesCreateParams) WithDefaults() *QualificationCategoriesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the qualification categories create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QualificationCategoriesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the qualification categories create params
func (o *QualificationCategoriesCreateParams) WithTimeout(timeout time.Duration) *QualificationCategoriesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the qualification categories create params
func (o *QualificationCategoriesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the qualification categories create params
func (o *QualificationCategoriesCreateParams) WithContext(ctx context.Context) *QualificationCategoriesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the qualification categories create params
func (o *QualificationCategoriesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the qualification categories create params
func (o *QualificationCategoriesCreateParams) WithHTTPClient(client *http.Client) *QualificationCategoriesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the qualification categories create params
func (o *QualificationCategoriesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the qualification categories create params
func (o *QualificationCategoriesCreateParams) WithData(data *models.QualificationCategory) *QualificationCategoriesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the qualification categories create params
func (o *QualificationCategoriesCreateParams) SetData(data *models.QualificationCategory) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *QualificationCategoriesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
