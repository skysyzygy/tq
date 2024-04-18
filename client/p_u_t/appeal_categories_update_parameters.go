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

// NewAppealCategoriesUpdateParams creates a new AppealCategoriesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAppealCategoriesUpdateParams() *AppealCategoriesUpdateParams {
	return &AppealCategoriesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAppealCategoriesUpdateParamsWithTimeout creates a new AppealCategoriesUpdateParams object
// with the ability to set a timeout on a request.
func NewAppealCategoriesUpdateParamsWithTimeout(timeout time.Duration) *AppealCategoriesUpdateParams {
	return &AppealCategoriesUpdateParams{
		timeout: timeout,
	}
}

// NewAppealCategoriesUpdateParamsWithContext creates a new AppealCategoriesUpdateParams object
// with the ability to set a context for a request.
func NewAppealCategoriesUpdateParamsWithContext(ctx context.Context) *AppealCategoriesUpdateParams {
	return &AppealCategoriesUpdateParams{
		Context: ctx,
	}
}

// NewAppealCategoriesUpdateParamsWithHTTPClient creates a new AppealCategoriesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAppealCategoriesUpdateParamsWithHTTPClient(client *http.Client) *AppealCategoriesUpdateParams {
	return &AppealCategoriesUpdateParams{
		HTTPClient: client,
	}
}

/*
AppealCategoriesUpdateParams contains all the parameters to send to the API endpoint

	for the appeal categories update operation.

	Typically these are written to a http.Request.
*/
type AppealCategoriesUpdateParams struct {

	// Data.
	Data *models.AppealCategory

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the appeal categories update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppealCategoriesUpdateParams) WithDefaults() *AppealCategoriesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the appeal categories update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppealCategoriesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the appeal categories update params
func (o *AppealCategoriesUpdateParams) WithTimeout(timeout time.Duration) *AppealCategoriesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the appeal categories update params
func (o *AppealCategoriesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the appeal categories update params
func (o *AppealCategoriesUpdateParams) WithContext(ctx context.Context) *AppealCategoriesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the appeal categories update params
func (o *AppealCategoriesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the appeal categories update params
func (o *AppealCategoriesUpdateParams) WithHTTPClient(client *http.Client) *AppealCategoriesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the appeal categories update params
func (o *AppealCategoriesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the appeal categories update params
func (o *AppealCategoriesUpdateParams) WithData(data *models.AppealCategory) *AppealCategoriesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the appeal categories update params
func (o *AppealCategoriesUpdateParams) SetData(data *models.AppealCategory) {
	o.Data = data
}

// WithID adds the id to the appeal categories update params
func (o *AppealCategoriesUpdateParams) WithID(id string) *AppealCategoriesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the appeal categories update params
func (o *AppealCategoriesUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AppealCategoriesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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