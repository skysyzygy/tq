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

// NewHoldCodeUserGroupsGetSummariesParams creates a new HoldCodeUserGroupsGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHoldCodeUserGroupsGetSummariesParams() *HoldCodeUserGroupsGetSummariesParams {
	return &HoldCodeUserGroupsGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHoldCodeUserGroupsGetSummariesParamsWithTimeout creates a new HoldCodeUserGroupsGetSummariesParams object
// with the ability to set a timeout on a request.
func NewHoldCodeUserGroupsGetSummariesParamsWithTimeout(timeout time.Duration) *HoldCodeUserGroupsGetSummariesParams {
	return &HoldCodeUserGroupsGetSummariesParams{
		timeout: timeout,
	}
}

// NewHoldCodeUserGroupsGetSummariesParamsWithContext creates a new HoldCodeUserGroupsGetSummariesParams object
// with the ability to set a context for a request.
func NewHoldCodeUserGroupsGetSummariesParamsWithContext(ctx context.Context) *HoldCodeUserGroupsGetSummariesParams {
	return &HoldCodeUserGroupsGetSummariesParams{
		Context: ctx,
	}
}

// NewHoldCodeUserGroupsGetSummariesParamsWithHTTPClient creates a new HoldCodeUserGroupsGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewHoldCodeUserGroupsGetSummariesParamsWithHTTPClient(client *http.Client) *HoldCodeUserGroupsGetSummariesParams {
	return &HoldCodeUserGroupsGetSummariesParams{
		HTTPClient: client,
	}
}

/*
HoldCodeUserGroupsGetSummariesParams contains all the parameters to send to the API endpoint

	for the hold code user groups get summaries operation.

	Typically these are written to a http.Request.
*/
type HoldCodeUserGroupsGetSummariesParams struct {

	// HoldCode.
	HoldCode *string

	// UserGroup.
	UserGroup *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the hold code user groups get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HoldCodeUserGroupsGetSummariesParams) WithDefaults() *HoldCodeUserGroupsGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hold code user groups get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HoldCodeUserGroupsGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the hold code user groups get summaries params
func (o *HoldCodeUserGroupsGetSummariesParams) WithTimeout(timeout time.Duration) *HoldCodeUserGroupsGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hold code user groups get summaries params
func (o *HoldCodeUserGroupsGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hold code user groups get summaries params
func (o *HoldCodeUserGroupsGetSummariesParams) WithContext(ctx context.Context) *HoldCodeUserGroupsGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hold code user groups get summaries params
func (o *HoldCodeUserGroupsGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hold code user groups get summaries params
func (o *HoldCodeUserGroupsGetSummariesParams) WithHTTPClient(client *http.Client) *HoldCodeUserGroupsGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hold code user groups get summaries params
func (o *HoldCodeUserGroupsGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHoldCode adds the holdCode to the hold code user groups get summaries params
func (o *HoldCodeUserGroupsGetSummariesParams) WithHoldCode(holdCode *string) *HoldCodeUserGroupsGetSummariesParams {
	o.SetHoldCode(holdCode)
	return o
}

// SetHoldCode adds the holdCode to the hold code user groups get summaries params
func (o *HoldCodeUserGroupsGetSummariesParams) SetHoldCode(holdCode *string) {
	o.HoldCode = holdCode
}

// WithUserGroup adds the userGroup to the hold code user groups get summaries params
func (o *HoldCodeUserGroupsGetSummariesParams) WithUserGroup(userGroup *string) *HoldCodeUserGroupsGetSummariesParams {
	o.SetUserGroup(userGroup)
	return o
}

// SetUserGroup adds the userGroup to the hold code user groups get summaries params
func (o *HoldCodeUserGroupsGetSummariesParams) SetUserGroup(userGroup *string) {
	o.UserGroup = userGroup
}

// WriteToRequest writes these params to a swagger request
func (o *HoldCodeUserGroupsGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.HoldCode != nil {

		// query param holdCode
		var qrHoldCode string

		if o.HoldCode != nil {
			qrHoldCode = *o.HoldCode
		}
		qHoldCode := qrHoldCode
		if qHoldCode != "" {

			if err := r.SetQueryParam("holdCode", qHoldCode); err != nil {
				return err
			}
		}
	}

	if o.UserGroup != nil {

		// query param userGroup
		var qrUserGroup string

		if o.UserGroup != nil {
			qrUserGroup = *o.UserGroup
		}
		qUserGroup := qrUserGroup
		if qUserGroup != "" {

			if err := r.SetQueryParam("userGroup", qUserGroup); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
