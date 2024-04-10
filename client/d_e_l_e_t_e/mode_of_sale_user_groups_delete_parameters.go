// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

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

// NewModeOfSaleUserGroupsDeleteParams creates a new ModeOfSaleUserGroupsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewModeOfSaleUserGroupsDeleteParams() *ModeOfSaleUserGroupsDeleteParams {
	return &ModeOfSaleUserGroupsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewModeOfSaleUserGroupsDeleteParamsWithTimeout creates a new ModeOfSaleUserGroupsDeleteParams object
// with the ability to set a timeout on a request.
func NewModeOfSaleUserGroupsDeleteParamsWithTimeout(timeout time.Duration) *ModeOfSaleUserGroupsDeleteParams {
	return &ModeOfSaleUserGroupsDeleteParams{
		timeout: timeout,
	}
}

// NewModeOfSaleUserGroupsDeleteParamsWithContext creates a new ModeOfSaleUserGroupsDeleteParams object
// with the ability to set a context for a request.
func NewModeOfSaleUserGroupsDeleteParamsWithContext(ctx context.Context) *ModeOfSaleUserGroupsDeleteParams {
	return &ModeOfSaleUserGroupsDeleteParams{
		Context: ctx,
	}
}

// NewModeOfSaleUserGroupsDeleteParamsWithHTTPClient creates a new ModeOfSaleUserGroupsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewModeOfSaleUserGroupsDeleteParamsWithHTTPClient(client *http.Client) *ModeOfSaleUserGroupsDeleteParams {
	return &ModeOfSaleUserGroupsDeleteParams{
		HTTPClient: client,
	}
}

/*
ModeOfSaleUserGroupsDeleteParams contains all the parameters to send to the API endpoint

	for the mode of sale user groups delete operation.

	Typically these are written to a http.Request.
*/
type ModeOfSaleUserGroupsDeleteParams struct {

	// ModeOfSaleUserGroupID.
	ModeOfSaleUserGroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the mode of sale user groups delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModeOfSaleUserGroupsDeleteParams) WithDefaults() *ModeOfSaleUserGroupsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the mode of sale user groups delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModeOfSaleUserGroupsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the mode of sale user groups delete params
func (o *ModeOfSaleUserGroupsDeleteParams) WithTimeout(timeout time.Duration) *ModeOfSaleUserGroupsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the mode of sale user groups delete params
func (o *ModeOfSaleUserGroupsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the mode of sale user groups delete params
func (o *ModeOfSaleUserGroupsDeleteParams) WithContext(ctx context.Context) *ModeOfSaleUserGroupsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the mode of sale user groups delete params
func (o *ModeOfSaleUserGroupsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the mode of sale user groups delete params
func (o *ModeOfSaleUserGroupsDeleteParams) WithHTTPClient(client *http.Client) *ModeOfSaleUserGroupsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the mode of sale user groups delete params
func (o *ModeOfSaleUserGroupsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModeOfSaleUserGroupID adds the modeOfSaleUserGroupID to the mode of sale user groups delete params
func (o *ModeOfSaleUserGroupsDeleteParams) WithModeOfSaleUserGroupID(modeOfSaleUserGroupID string) *ModeOfSaleUserGroupsDeleteParams {
	o.SetModeOfSaleUserGroupID(modeOfSaleUserGroupID)
	return o
}

// SetModeOfSaleUserGroupID adds the modeOfSaleUserGroupId to the mode of sale user groups delete params
func (o *ModeOfSaleUserGroupsDeleteParams) SetModeOfSaleUserGroupID(modeOfSaleUserGroupID string) {
	o.ModeOfSaleUserGroupID = modeOfSaleUserGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *ModeOfSaleUserGroupsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param modeOfSaleUserGroupId
	if err := r.SetPathParam("modeOfSaleUserGroupId", o.ModeOfSaleUserGroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
