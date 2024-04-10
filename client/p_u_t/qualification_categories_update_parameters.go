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

// NewQualificationCategoriesUpdateParams creates a new QualificationCategoriesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQualificationCategoriesUpdateParams() *QualificationCategoriesUpdateParams {
	return &QualificationCategoriesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQualificationCategoriesUpdateParamsWithTimeout creates a new QualificationCategoriesUpdateParams object
// with the ability to set a timeout on a request.
func NewQualificationCategoriesUpdateParamsWithTimeout(timeout time.Duration) *QualificationCategoriesUpdateParams {
	return &QualificationCategoriesUpdateParams{
		timeout: timeout,
	}
}

// NewQualificationCategoriesUpdateParamsWithContext creates a new QualificationCategoriesUpdateParams object
// with the ability to set a context for a request.
func NewQualificationCategoriesUpdateParamsWithContext(ctx context.Context) *QualificationCategoriesUpdateParams {
	return &QualificationCategoriesUpdateParams{
		Context: ctx,
	}
}

// NewQualificationCategoriesUpdateParamsWithHTTPClient creates a new QualificationCategoriesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewQualificationCategoriesUpdateParamsWithHTTPClient(client *http.Client) *QualificationCategoriesUpdateParams {
	return &QualificationCategoriesUpdateParams{
		HTTPClient: client,
	}
}

/*
QualificationCategoriesUpdateParams contains all the parameters to send to the API endpoint

	for the qualification categories update operation.

	Typically these are written to a http.Request.
*/
type QualificationCategoriesUpdateParams struct {

	/* Data.

	   The updated resource to be saved
	*/
	Data *models.QualificationCategory

	/* ID.

	   The id of the resource to be updated
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the qualification categories update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QualificationCategoriesUpdateParams) WithDefaults() *QualificationCategoriesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the qualification categories update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QualificationCategoriesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the qualification categories update params
func (o *QualificationCategoriesUpdateParams) WithTimeout(timeout time.Duration) *QualificationCategoriesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the qualification categories update params
func (o *QualificationCategoriesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the qualification categories update params
func (o *QualificationCategoriesUpdateParams) WithContext(ctx context.Context) *QualificationCategoriesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the qualification categories update params
func (o *QualificationCategoriesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the qualification categories update params
func (o *QualificationCategoriesUpdateParams) WithHTTPClient(client *http.Client) *QualificationCategoriesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the qualification categories update params
func (o *QualificationCategoriesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the qualification categories update params
func (o *QualificationCategoriesUpdateParams) WithData(data *models.QualificationCategory) *QualificationCategoriesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the qualification categories update params
func (o *QualificationCategoriesUpdateParams) SetData(data *models.QualificationCategory) {
	o.Data = data
}

// WithID adds the id to the qualification categories update params
func (o *QualificationCategoriesUpdateParams) WithID(id string) *QualificationCategoriesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the qualification categories update params
func (o *QualificationCategoriesUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *QualificationCategoriesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
