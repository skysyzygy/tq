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

// NewResearchTypesUpdateParams creates a new ResearchTypesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResearchTypesUpdateParams() *ResearchTypesUpdateParams {
	return &ResearchTypesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResearchTypesUpdateParamsWithTimeout creates a new ResearchTypesUpdateParams object
// with the ability to set a timeout on a request.
func NewResearchTypesUpdateParamsWithTimeout(timeout time.Duration) *ResearchTypesUpdateParams {
	return &ResearchTypesUpdateParams{
		timeout: timeout,
	}
}

// NewResearchTypesUpdateParamsWithContext creates a new ResearchTypesUpdateParams object
// with the ability to set a context for a request.
func NewResearchTypesUpdateParamsWithContext(ctx context.Context) *ResearchTypesUpdateParams {
	return &ResearchTypesUpdateParams{
		Context: ctx,
	}
}

// NewResearchTypesUpdateParamsWithHTTPClient creates a new ResearchTypesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewResearchTypesUpdateParamsWithHTTPClient(client *http.Client) *ResearchTypesUpdateParams {
	return &ResearchTypesUpdateParams{
		HTTPClient: client,
	}
}

/*
ResearchTypesUpdateParams contains all the parameters to send to the API endpoint

	for the research types update operation.

	Typically these are written to a http.Request.
*/
type ResearchTypesUpdateParams struct {

	// Data.
	Data *models.ResearchType

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the research types update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResearchTypesUpdateParams) WithDefaults() *ResearchTypesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the research types update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResearchTypesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the research types update params
func (o *ResearchTypesUpdateParams) WithTimeout(timeout time.Duration) *ResearchTypesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the research types update params
func (o *ResearchTypesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the research types update params
func (o *ResearchTypesUpdateParams) WithContext(ctx context.Context) *ResearchTypesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the research types update params
func (o *ResearchTypesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the research types update params
func (o *ResearchTypesUpdateParams) WithHTTPClient(client *http.Client) *ResearchTypesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the research types update params
func (o *ResearchTypesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the research types update params
func (o *ResearchTypesUpdateParams) WithData(data *models.ResearchType) *ResearchTypesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the research types update params
func (o *ResearchTypesUpdateParams) SetData(data *models.ResearchType) {
	o.Data = data
}

// WithID adds the id to the research types update params
func (o *ResearchTypesUpdateParams) WithID(id string) *ResearchTypesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the research types update params
func (o *ResearchTypesUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ResearchTypesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
