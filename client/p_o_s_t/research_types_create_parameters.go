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

// NewResearchTypesCreateParams creates a new ResearchTypesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResearchTypesCreateParams() *ResearchTypesCreateParams {
	return &ResearchTypesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResearchTypesCreateParamsWithTimeout creates a new ResearchTypesCreateParams object
// with the ability to set a timeout on a request.
func NewResearchTypesCreateParamsWithTimeout(timeout time.Duration) *ResearchTypesCreateParams {
	return &ResearchTypesCreateParams{
		timeout: timeout,
	}
}

// NewResearchTypesCreateParamsWithContext creates a new ResearchTypesCreateParams object
// with the ability to set a context for a request.
func NewResearchTypesCreateParamsWithContext(ctx context.Context) *ResearchTypesCreateParams {
	return &ResearchTypesCreateParams{
		Context: ctx,
	}
}

// NewResearchTypesCreateParamsWithHTTPClient creates a new ResearchTypesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewResearchTypesCreateParamsWithHTTPClient(client *http.Client) *ResearchTypesCreateParams {
	return &ResearchTypesCreateParams{
		HTTPClient: client,
	}
}

/*
ResearchTypesCreateParams contains all the parameters to send to the API endpoint

	for the research types create operation.

	Typically these are written to a http.Request.
*/
type ResearchTypesCreateParams struct {

	// Data.
	Data *models.ResearchType

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the research types create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResearchTypesCreateParams) WithDefaults() *ResearchTypesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the research types create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResearchTypesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the research types create params
func (o *ResearchTypesCreateParams) WithTimeout(timeout time.Duration) *ResearchTypesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the research types create params
func (o *ResearchTypesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the research types create params
func (o *ResearchTypesCreateParams) WithContext(ctx context.Context) *ResearchTypesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the research types create params
func (o *ResearchTypesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the research types create params
func (o *ResearchTypesCreateParams) WithHTTPClient(client *http.Client) *ResearchTypesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the research types create params
func (o *ResearchTypesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the research types create params
func (o *ResearchTypesCreateParams) WithData(data *models.ResearchType) *ResearchTypesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the research types create params
func (o *ResearchTypesCreateParams) SetData(data *models.ResearchType) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ResearchTypesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
