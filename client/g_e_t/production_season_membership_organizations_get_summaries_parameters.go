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

// NewProductionSeasonMembershipOrganizationsGetSummariesParams creates a new ProductionSeasonMembershipOrganizationsGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProductionSeasonMembershipOrganizationsGetSummariesParams() *ProductionSeasonMembershipOrganizationsGetSummariesParams {
	return &ProductionSeasonMembershipOrganizationsGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProductionSeasonMembershipOrganizationsGetSummariesParamsWithTimeout creates a new ProductionSeasonMembershipOrganizationsGetSummariesParams object
// with the ability to set a timeout on a request.
func NewProductionSeasonMembershipOrganizationsGetSummariesParamsWithTimeout(timeout time.Duration) *ProductionSeasonMembershipOrganizationsGetSummariesParams {
	return &ProductionSeasonMembershipOrganizationsGetSummariesParams{
		timeout: timeout,
	}
}

// NewProductionSeasonMembershipOrganizationsGetSummariesParamsWithContext creates a new ProductionSeasonMembershipOrganizationsGetSummariesParams object
// with the ability to set a context for a request.
func NewProductionSeasonMembershipOrganizationsGetSummariesParamsWithContext(ctx context.Context) *ProductionSeasonMembershipOrganizationsGetSummariesParams {
	return &ProductionSeasonMembershipOrganizationsGetSummariesParams{
		Context: ctx,
	}
}

// NewProductionSeasonMembershipOrganizationsGetSummariesParamsWithHTTPClient creates a new ProductionSeasonMembershipOrganizationsGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewProductionSeasonMembershipOrganizationsGetSummariesParamsWithHTTPClient(client *http.Client) *ProductionSeasonMembershipOrganizationsGetSummariesParams {
	return &ProductionSeasonMembershipOrganizationsGetSummariesParams{
		HTTPClient: client,
	}
}

/*
ProductionSeasonMembershipOrganizationsGetSummariesParams contains all the parameters to send to the API endpoint

	for the production season membership organizations get summaries operation.

	Typically these are written to a http.Request.
*/
type ProductionSeasonMembershipOrganizationsGetSummariesParams struct {

	// ProductionSeasonID.
	ProductionSeasonID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the production season membership organizations get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProductionSeasonMembershipOrganizationsGetSummariesParams) WithDefaults() *ProductionSeasonMembershipOrganizationsGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the production season membership organizations get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProductionSeasonMembershipOrganizationsGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the production season membership organizations get summaries params
func (o *ProductionSeasonMembershipOrganizationsGetSummariesParams) WithTimeout(timeout time.Duration) *ProductionSeasonMembershipOrganizationsGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the production season membership organizations get summaries params
func (o *ProductionSeasonMembershipOrganizationsGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the production season membership organizations get summaries params
func (o *ProductionSeasonMembershipOrganizationsGetSummariesParams) WithContext(ctx context.Context) *ProductionSeasonMembershipOrganizationsGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the production season membership organizations get summaries params
func (o *ProductionSeasonMembershipOrganizationsGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the production season membership organizations get summaries params
func (o *ProductionSeasonMembershipOrganizationsGetSummariesParams) WithHTTPClient(client *http.Client) *ProductionSeasonMembershipOrganizationsGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the production season membership organizations get summaries params
func (o *ProductionSeasonMembershipOrganizationsGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProductionSeasonID adds the productionSeasonID to the production season membership organizations get summaries params
func (o *ProductionSeasonMembershipOrganizationsGetSummariesParams) WithProductionSeasonID(productionSeasonID *string) *ProductionSeasonMembershipOrganizationsGetSummariesParams {
	o.SetProductionSeasonID(productionSeasonID)
	return o
}

// SetProductionSeasonID adds the productionSeasonId to the production season membership organizations get summaries params
func (o *ProductionSeasonMembershipOrganizationsGetSummariesParams) SetProductionSeasonID(productionSeasonID *string) {
	o.ProductionSeasonID = productionSeasonID
}

// WriteToRequest writes these params to a swagger request
func (o *ProductionSeasonMembershipOrganizationsGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ProductionSeasonID != nil {

		// query param productionSeasonId
		var qrProductionSeasonID string

		if o.ProductionSeasonID != nil {
			qrProductionSeasonID = *o.ProductionSeasonID
		}
		qProductionSeasonID := qrProductionSeasonID
		if qProductionSeasonID != "" {

			if err := r.SetQueryParam("productionSeasonId", qProductionSeasonID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
