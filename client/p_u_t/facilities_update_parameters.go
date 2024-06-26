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

// NewFacilitiesUpdateParams creates a new FacilitiesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFacilitiesUpdateParams() *FacilitiesUpdateParams {
	return &FacilitiesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFacilitiesUpdateParamsWithTimeout creates a new FacilitiesUpdateParams object
// with the ability to set a timeout on a request.
func NewFacilitiesUpdateParamsWithTimeout(timeout time.Duration) *FacilitiesUpdateParams {
	return &FacilitiesUpdateParams{
		timeout: timeout,
	}
}

// NewFacilitiesUpdateParamsWithContext creates a new FacilitiesUpdateParams object
// with the ability to set a context for a request.
func NewFacilitiesUpdateParamsWithContext(ctx context.Context) *FacilitiesUpdateParams {
	return &FacilitiesUpdateParams{
		Context: ctx,
	}
}

// NewFacilitiesUpdateParamsWithHTTPClient creates a new FacilitiesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewFacilitiesUpdateParamsWithHTTPClient(client *http.Client) *FacilitiesUpdateParams {
	return &FacilitiesUpdateParams{
		HTTPClient: client,
	}
}

/*
FacilitiesUpdateParams contains all the parameters to send to the API endpoint

	for the facilities update operation.

	Typically these are written to a http.Request.
*/
type FacilitiesUpdateParams struct {

	// Facility.
	Facility *models.Facility

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the facilities update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FacilitiesUpdateParams) WithDefaults() *FacilitiesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the facilities update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FacilitiesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the facilities update params
func (o *FacilitiesUpdateParams) WithTimeout(timeout time.Duration) *FacilitiesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the facilities update params
func (o *FacilitiesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the facilities update params
func (o *FacilitiesUpdateParams) WithContext(ctx context.Context) *FacilitiesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the facilities update params
func (o *FacilitiesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the facilities update params
func (o *FacilitiesUpdateParams) WithHTTPClient(client *http.Client) *FacilitiesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the facilities update params
func (o *FacilitiesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFacility adds the facility to the facilities update params
func (o *FacilitiesUpdateParams) WithFacility(facility *models.Facility) *FacilitiesUpdateParams {
	o.SetFacility(facility)
	return o
}

// SetFacility adds the facility to the facilities update params
func (o *FacilitiesUpdateParams) SetFacility(facility *models.Facility) {
	o.Facility = facility
}

// WithID adds the id to the facilities update params
func (o *FacilitiesUpdateParams) WithID(id string) *FacilitiesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the facilities update params
func (o *FacilitiesUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FacilitiesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Facility != nil {
		if err := r.SetBodyParam(o.Facility); err != nil {
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
