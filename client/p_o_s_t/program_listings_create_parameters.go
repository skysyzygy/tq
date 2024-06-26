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

// NewProgramListingsCreateParams creates a new ProgramListingsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProgramListingsCreateParams() *ProgramListingsCreateParams {
	return &ProgramListingsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProgramListingsCreateParamsWithTimeout creates a new ProgramListingsCreateParams object
// with the ability to set a timeout on a request.
func NewProgramListingsCreateParamsWithTimeout(timeout time.Duration) *ProgramListingsCreateParams {
	return &ProgramListingsCreateParams{
		timeout: timeout,
	}
}

// NewProgramListingsCreateParamsWithContext creates a new ProgramListingsCreateParams object
// with the ability to set a context for a request.
func NewProgramListingsCreateParamsWithContext(ctx context.Context) *ProgramListingsCreateParams {
	return &ProgramListingsCreateParams{
		Context: ctx,
	}
}

// NewProgramListingsCreateParamsWithHTTPClient creates a new ProgramListingsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewProgramListingsCreateParamsWithHTTPClient(client *http.Client) *ProgramListingsCreateParams {
	return &ProgramListingsCreateParams{
		HTTPClient: client,
	}
}

/*
ProgramListingsCreateParams contains all the parameters to send to the API endpoint

	for the program listings create operation.

	Typically these are written to a http.Request.
*/
type ProgramListingsCreateParams struct {

	// ProgramListing.
	ProgramListing *models.ProgramListing

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the program listings create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProgramListingsCreateParams) WithDefaults() *ProgramListingsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the program listings create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProgramListingsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the program listings create params
func (o *ProgramListingsCreateParams) WithTimeout(timeout time.Duration) *ProgramListingsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the program listings create params
func (o *ProgramListingsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the program listings create params
func (o *ProgramListingsCreateParams) WithContext(ctx context.Context) *ProgramListingsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the program listings create params
func (o *ProgramListingsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the program listings create params
func (o *ProgramListingsCreateParams) WithHTTPClient(client *http.Client) *ProgramListingsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the program listings create params
func (o *ProgramListingsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProgramListing adds the programListing to the program listings create params
func (o *ProgramListingsCreateParams) WithProgramListing(programListing *models.ProgramListing) *ProgramListingsCreateParams {
	o.SetProgramListing(programListing)
	return o
}

// SetProgramListing adds the programListing to the program listings create params
func (o *ProgramListingsCreateParams) SetProgramListing(programListing *models.ProgramListing) {
	o.ProgramListing = programListing
}

// WriteToRequest writes these params to a swagger request
func (o *ProgramListingsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ProgramListing != nil {
		if err := r.SetBodyParam(o.ProgramListing); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
