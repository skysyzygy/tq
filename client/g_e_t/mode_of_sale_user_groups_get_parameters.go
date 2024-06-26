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

// NewModeOfSaleUserGroupsGetParams creates a new ModeOfSaleUserGroupsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewModeOfSaleUserGroupsGetParams() *ModeOfSaleUserGroupsGetParams {
	return &ModeOfSaleUserGroupsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewModeOfSaleUserGroupsGetParamsWithTimeout creates a new ModeOfSaleUserGroupsGetParams object
// with the ability to set a timeout on a request.
func NewModeOfSaleUserGroupsGetParamsWithTimeout(timeout time.Duration) *ModeOfSaleUserGroupsGetParams {
	return &ModeOfSaleUserGroupsGetParams{
		timeout: timeout,
	}
}

// NewModeOfSaleUserGroupsGetParamsWithContext creates a new ModeOfSaleUserGroupsGetParams object
// with the ability to set a context for a request.
func NewModeOfSaleUserGroupsGetParamsWithContext(ctx context.Context) *ModeOfSaleUserGroupsGetParams {
	return &ModeOfSaleUserGroupsGetParams{
		Context: ctx,
	}
}

// NewModeOfSaleUserGroupsGetParamsWithHTTPClient creates a new ModeOfSaleUserGroupsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewModeOfSaleUserGroupsGetParamsWithHTTPClient(client *http.Client) *ModeOfSaleUserGroupsGetParams {
	return &ModeOfSaleUserGroupsGetParams{
		HTTPClient: client,
	}
}

/*
ModeOfSaleUserGroupsGetParams contains all the parameters to send to the API endpoint

	for the mode of sale user groups get operation.

	Typically these are written to a http.Request.
*/
type ModeOfSaleUserGroupsGetParams struct {

	// ModeOfSaleUserGroupID.
	ModeOfSaleUserGroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the mode of sale user groups get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModeOfSaleUserGroupsGetParams) WithDefaults() *ModeOfSaleUserGroupsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the mode of sale user groups get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModeOfSaleUserGroupsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the mode of sale user groups get params
func (o *ModeOfSaleUserGroupsGetParams) WithTimeout(timeout time.Duration) *ModeOfSaleUserGroupsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the mode of sale user groups get params
func (o *ModeOfSaleUserGroupsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the mode of sale user groups get params
func (o *ModeOfSaleUserGroupsGetParams) WithContext(ctx context.Context) *ModeOfSaleUserGroupsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the mode of sale user groups get params
func (o *ModeOfSaleUserGroupsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the mode of sale user groups get params
func (o *ModeOfSaleUserGroupsGetParams) WithHTTPClient(client *http.Client) *ModeOfSaleUserGroupsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the mode of sale user groups get params
func (o *ModeOfSaleUserGroupsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModeOfSaleUserGroupID adds the modeOfSaleUserGroupID to the mode of sale user groups get params
func (o *ModeOfSaleUserGroupsGetParams) WithModeOfSaleUserGroupID(modeOfSaleUserGroupID string) *ModeOfSaleUserGroupsGetParams {
	o.SetModeOfSaleUserGroupID(modeOfSaleUserGroupID)
	return o
}

// SetModeOfSaleUserGroupID adds the modeOfSaleUserGroupId to the mode of sale user groups get params
func (o *ModeOfSaleUserGroupsGetParams) SetModeOfSaleUserGroupID(modeOfSaleUserGroupID string) {
	o.ModeOfSaleUserGroupID = modeOfSaleUserGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *ModeOfSaleUserGroupsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
