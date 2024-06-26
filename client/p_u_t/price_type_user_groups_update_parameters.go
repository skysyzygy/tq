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

// NewPriceTypeUserGroupsUpdateParams creates a new PriceTypeUserGroupsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPriceTypeUserGroupsUpdateParams() *PriceTypeUserGroupsUpdateParams {
	return &PriceTypeUserGroupsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPriceTypeUserGroupsUpdateParamsWithTimeout creates a new PriceTypeUserGroupsUpdateParams object
// with the ability to set a timeout on a request.
func NewPriceTypeUserGroupsUpdateParamsWithTimeout(timeout time.Duration) *PriceTypeUserGroupsUpdateParams {
	return &PriceTypeUserGroupsUpdateParams{
		timeout: timeout,
	}
}

// NewPriceTypeUserGroupsUpdateParamsWithContext creates a new PriceTypeUserGroupsUpdateParams object
// with the ability to set a context for a request.
func NewPriceTypeUserGroupsUpdateParamsWithContext(ctx context.Context) *PriceTypeUserGroupsUpdateParams {
	return &PriceTypeUserGroupsUpdateParams{
		Context: ctx,
	}
}

// NewPriceTypeUserGroupsUpdateParamsWithHTTPClient creates a new PriceTypeUserGroupsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPriceTypeUserGroupsUpdateParamsWithHTTPClient(client *http.Client) *PriceTypeUserGroupsUpdateParams {
	return &PriceTypeUserGroupsUpdateParams{
		HTTPClient: client,
	}
}

/*
PriceTypeUserGroupsUpdateParams contains all the parameters to send to the API endpoint

	for the price type user groups update operation.

	Typically these are written to a http.Request.
*/
type PriceTypeUserGroupsUpdateParams struct {

	// PriceTypeUserGroup.
	PriceTypeUserGroup *models.PriceTypeUserGroup

	// PriceTypeUserGroupID.
	PriceTypeUserGroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the price type user groups update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PriceTypeUserGroupsUpdateParams) WithDefaults() *PriceTypeUserGroupsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the price type user groups update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PriceTypeUserGroupsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the price type user groups update params
func (o *PriceTypeUserGroupsUpdateParams) WithTimeout(timeout time.Duration) *PriceTypeUserGroupsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the price type user groups update params
func (o *PriceTypeUserGroupsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the price type user groups update params
func (o *PriceTypeUserGroupsUpdateParams) WithContext(ctx context.Context) *PriceTypeUserGroupsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the price type user groups update params
func (o *PriceTypeUserGroupsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the price type user groups update params
func (o *PriceTypeUserGroupsUpdateParams) WithHTTPClient(client *http.Client) *PriceTypeUserGroupsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the price type user groups update params
func (o *PriceTypeUserGroupsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPriceTypeUserGroup adds the priceTypeUserGroup to the price type user groups update params
func (o *PriceTypeUserGroupsUpdateParams) WithPriceTypeUserGroup(priceTypeUserGroup *models.PriceTypeUserGroup) *PriceTypeUserGroupsUpdateParams {
	o.SetPriceTypeUserGroup(priceTypeUserGroup)
	return o
}

// SetPriceTypeUserGroup adds the priceTypeUserGroup to the price type user groups update params
func (o *PriceTypeUserGroupsUpdateParams) SetPriceTypeUserGroup(priceTypeUserGroup *models.PriceTypeUserGroup) {
	o.PriceTypeUserGroup = priceTypeUserGroup
}

// WithPriceTypeUserGroupID adds the priceTypeUserGroupID to the price type user groups update params
func (o *PriceTypeUserGroupsUpdateParams) WithPriceTypeUserGroupID(priceTypeUserGroupID string) *PriceTypeUserGroupsUpdateParams {
	o.SetPriceTypeUserGroupID(priceTypeUserGroupID)
	return o
}

// SetPriceTypeUserGroupID adds the priceTypeUserGroupId to the price type user groups update params
func (o *PriceTypeUserGroupsUpdateParams) SetPriceTypeUserGroupID(priceTypeUserGroupID string) {
	o.PriceTypeUserGroupID = priceTypeUserGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *PriceTypeUserGroupsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PriceTypeUserGroup != nil {
		if err := r.SetBodyParam(o.PriceTypeUserGroup); err != nil {
			return err
		}
	}

	// path param priceTypeUserGroupId
	if err := r.SetPathParam("priceTypeUserGroupId", o.PriceTypeUserGroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
