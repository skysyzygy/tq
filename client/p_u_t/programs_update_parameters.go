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

// NewProgramsUpdateParams creates a new ProgramsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProgramsUpdateParams() *ProgramsUpdateParams {
	return &ProgramsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProgramsUpdateParamsWithTimeout creates a new ProgramsUpdateParams object
// with the ability to set a timeout on a request.
func NewProgramsUpdateParamsWithTimeout(timeout time.Duration) *ProgramsUpdateParams {
	return &ProgramsUpdateParams{
		timeout: timeout,
	}
}

// NewProgramsUpdateParamsWithContext creates a new ProgramsUpdateParams object
// with the ability to set a context for a request.
func NewProgramsUpdateParamsWithContext(ctx context.Context) *ProgramsUpdateParams {
	return &ProgramsUpdateParams{
		Context: ctx,
	}
}

// NewProgramsUpdateParamsWithHTTPClient creates a new ProgramsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewProgramsUpdateParamsWithHTTPClient(client *http.Client) *ProgramsUpdateParams {
	return &ProgramsUpdateParams{
		HTTPClient: client,
	}
}

/*
ProgramsUpdateParams contains all the parameters to send to the API endpoint

	for the programs update operation.

	Typically these are written to a http.Request.
*/
type ProgramsUpdateParams struct {

	// Data.
	Data *models.Program

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the programs update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProgramsUpdateParams) WithDefaults() *ProgramsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the programs update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProgramsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the programs update params
func (o *ProgramsUpdateParams) WithTimeout(timeout time.Duration) *ProgramsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the programs update params
func (o *ProgramsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the programs update params
func (o *ProgramsUpdateParams) WithContext(ctx context.Context) *ProgramsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the programs update params
func (o *ProgramsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the programs update params
func (o *ProgramsUpdateParams) WithHTTPClient(client *http.Client) *ProgramsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the programs update params
func (o *ProgramsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the programs update params
func (o *ProgramsUpdateParams) WithData(data *models.Program) *ProgramsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the programs update params
func (o *ProgramsUpdateParams) SetData(data *models.Program) {
	o.Data = data
}

// WithID adds the id to the programs update params
func (o *ProgramsUpdateParams) WithID(id string) *ProgramsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the programs update params
func (o *ProgramsUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ProgramsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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