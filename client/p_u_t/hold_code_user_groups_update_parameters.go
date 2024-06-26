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

// NewHoldCodeUserGroupsUpdateParams creates a new HoldCodeUserGroupsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHoldCodeUserGroupsUpdateParams() *HoldCodeUserGroupsUpdateParams {
	return &HoldCodeUserGroupsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHoldCodeUserGroupsUpdateParamsWithTimeout creates a new HoldCodeUserGroupsUpdateParams object
// with the ability to set a timeout on a request.
func NewHoldCodeUserGroupsUpdateParamsWithTimeout(timeout time.Duration) *HoldCodeUserGroupsUpdateParams {
	return &HoldCodeUserGroupsUpdateParams{
		timeout: timeout,
	}
}

// NewHoldCodeUserGroupsUpdateParamsWithContext creates a new HoldCodeUserGroupsUpdateParams object
// with the ability to set a context for a request.
func NewHoldCodeUserGroupsUpdateParamsWithContext(ctx context.Context) *HoldCodeUserGroupsUpdateParams {
	return &HoldCodeUserGroupsUpdateParams{
		Context: ctx,
	}
}

// NewHoldCodeUserGroupsUpdateParamsWithHTTPClient creates a new HoldCodeUserGroupsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewHoldCodeUserGroupsUpdateParamsWithHTTPClient(client *http.Client) *HoldCodeUserGroupsUpdateParams {
	return &HoldCodeUserGroupsUpdateParams{
		HTTPClient: client,
	}
}

/*
HoldCodeUserGroupsUpdateParams contains all the parameters to send to the API endpoint

	for the hold code user groups update operation.

	Typically these are written to a http.Request.
*/
type HoldCodeUserGroupsUpdateParams struct {

	// HoldCodeUserGroup.
	HoldCodeUserGroup *models.HoldCodeUserGroup

	// HoldCodeUserGroupID.
	HoldCodeUserGroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the hold code user groups update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HoldCodeUserGroupsUpdateParams) WithDefaults() *HoldCodeUserGroupsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hold code user groups update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HoldCodeUserGroupsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the hold code user groups update params
func (o *HoldCodeUserGroupsUpdateParams) WithTimeout(timeout time.Duration) *HoldCodeUserGroupsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hold code user groups update params
func (o *HoldCodeUserGroupsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hold code user groups update params
func (o *HoldCodeUserGroupsUpdateParams) WithContext(ctx context.Context) *HoldCodeUserGroupsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hold code user groups update params
func (o *HoldCodeUserGroupsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hold code user groups update params
func (o *HoldCodeUserGroupsUpdateParams) WithHTTPClient(client *http.Client) *HoldCodeUserGroupsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hold code user groups update params
func (o *HoldCodeUserGroupsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHoldCodeUserGroup adds the holdCodeUserGroup to the hold code user groups update params
func (o *HoldCodeUserGroupsUpdateParams) WithHoldCodeUserGroup(holdCodeUserGroup *models.HoldCodeUserGroup) *HoldCodeUserGroupsUpdateParams {
	o.SetHoldCodeUserGroup(holdCodeUserGroup)
	return o
}

// SetHoldCodeUserGroup adds the holdCodeUserGroup to the hold code user groups update params
func (o *HoldCodeUserGroupsUpdateParams) SetHoldCodeUserGroup(holdCodeUserGroup *models.HoldCodeUserGroup) {
	o.HoldCodeUserGroup = holdCodeUserGroup
}

// WithHoldCodeUserGroupID adds the holdCodeUserGroupID to the hold code user groups update params
func (o *HoldCodeUserGroupsUpdateParams) WithHoldCodeUserGroupID(holdCodeUserGroupID string) *HoldCodeUserGroupsUpdateParams {
	o.SetHoldCodeUserGroupID(holdCodeUserGroupID)
	return o
}

// SetHoldCodeUserGroupID adds the holdCodeUserGroupId to the hold code user groups update params
func (o *HoldCodeUserGroupsUpdateParams) SetHoldCodeUserGroupID(holdCodeUserGroupID string) {
	o.HoldCodeUserGroupID = holdCodeUserGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *HoldCodeUserGroupsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.HoldCodeUserGroup != nil {
		if err := r.SetBodyParam(o.HoldCodeUserGroup); err != nil {
			return err
		}
	}

	// path param holdCodeUserGroupId
	if err := r.SetPathParam("holdCodeUserGroupId", o.HoldCodeUserGroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
