// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

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
)

// NewProductionSeasonMembershipOrganizationsGetParams creates a new ProductionSeasonMembershipOrganizationsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProductionSeasonMembershipOrganizationsGetParams() *ProductionSeasonMembershipOrganizationsGetParams {
	return &ProductionSeasonMembershipOrganizationsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProductionSeasonMembershipOrganizationsGetParamsWithTimeout creates a new ProductionSeasonMembershipOrganizationsGetParams object
// with the ability to set a timeout on a request.
func NewProductionSeasonMembershipOrganizationsGetParamsWithTimeout(timeout time.Duration) *ProductionSeasonMembershipOrganizationsGetParams {
	return &ProductionSeasonMembershipOrganizationsGetParams{
		timeout: timeout,
	}
}

// NewProductionSeasonMembershipOrganizationsGetParamsWithContext creates a new ProductionSeasonMembershipOrganizationsGetParams object
// with the ability to set a context for a request.
func NewProductionSeasonMembershipOrganizationsGetParamsWithContext(ctx context.Context) *ProductionSeasonMembershipOrganizationsGetParams {
	return &ProductionSeasonMembershipOrganizationsGetParams{
		Context: ctx,
	}
}

// NewProductionSeasonMembershipOrganizationsGetParamsWithHTTPClient creates a new ProductionSeasonMembershipOrganizationsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewProductionSeasonMembershipOrganizationsGetParamsWithHTTPClient(client *http.Client) *ProductionSeasonMembershipOrganizationsGetParams {
	return &ProductionSeasonMembershipOrganizationsGetParams{
		HTTPClient: client,
	}
}

/*
ProductionSeasonMembershipOrganizationsGetParams contains all the parameters to send to the API endpoint

	for the production season membership organizations get operation.

	Typically these are written to a http.Request.
*/
type ProductionSeasonMembershipOrganizationsGetParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the production season membership organizations get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProductionSeasonMembershipOrganizationsGetParams) WithDefaults() *ProductionSeasonMembershipOrganizationsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the production season membership organizations get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProductionSeasonMembershipOrganizationsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the production season membership organizations get params
func (o *ProductionSeasonMembershipOrganizationsGetParams) WithTimeout(timeout time.Duration) *ProductionSeasonMembershipOrganizationsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the production season membership organizations get params
func (o *ProductionSeasonMembershipOrganizationsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the production season membership organizations get params
func (o *ProductionSeasonMembershipOrganizationsGetParams) WithContext(ctx context.Context) *ProductionSeasonMembershipOrganizationsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the production season membership organizations get params
func (o *ProductionSeasonMembershipOrganizationsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the production season membership organizations get params
func (o *ProductionSeasonMembershipOrganizationsGetParams) WithHTTPClient(client *http.Client) *ProductionSeasonMembershipOrganizationsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the production season membership organizations get params
func (o *ProductionSeasonMembershipOrganizationsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the production season membership organizations get params
func (o *ProductionSeasonMembershipOrganizationsGetParams) WithID(id string) *ProductionSeasonMembershipOrganizationsGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the production season membership organizations get params
func (o *ProductionSeasonMembershipOrganizationsGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ProductionSeasonMembershipOrganizationsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
