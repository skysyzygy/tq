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

// NewResourceSchedulesDeleteParams creates a new ResourceSchedulesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResourceSchedulesDeleteParams() *ResourceSchedulesDeleteParams {
	return &ResourceSchedulesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResourceSchedulesDeleteParamsWithTimeout creates a new ResourceSchedulesDeleteParams object
// with the ability to set a timeout on a request.
func NewResourceSchedulesDeleteParamsWithTimeout(timeout time.Duration) *ResourceSchedulesDeleteParams {
	return &ResourceSchedulesDeleteParams{
		timeout: timeout,
	}
}

// NewResourceSchedulesDeleteParamsWithContext creates a new ResourceSchedulesDeleteParams object
// with the ability to set a context for a request.
func NewResourceSchedulesDeleteParamsWithContext(ctx context.Context) *ResourceSchedulesDeleteParams {
	return &ResourceSchedulesDeleteParams{
		Context: ctx,
	}
}

// NewResourceSchedulesDeleteParamsWithHTTPClient creates a new ResourceSchedulesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewResourceSchedulesDeleteParamsWithHTTPClient(client *http.Client) *ResourceSchedulesDeleteParams {
	return &ResourceSchedulesDeleteParams{
		HTTPClient: client,
	}
}

/*
ResourceSchedulesDeleteParams contains all the parameters to send to the API endpoint

	for the resource schedules delete operation.

	Typically these are written to a http.Request.
*/
type ResourceSchedulesDeleteParams struct {

	/* OverrideConflicts.

	   When passed as true and the user group has rights to Override Resource Assignments, any overlapping will be updated to Ignore Schedule Conflicts
	*/
	OverrideConflicts *string

	// ResourceScheduleID.
	ResourceScheduleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the resource schedules delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResourceSchedulesDeleteParams) WithDefaults() *ResourceSchedulesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the resource schedules delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResourceSchedulesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the resource schedules delete params
func (o *ResourceSchedulesDeleteParams) WithTimeout(timeout time.Duration) *ResourceSchedulesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resource schedules delete params
func (o *ResourceSchedulesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resource schedules delete params
func (o *ResourceSchedulesDeleteParams) WithContext(ctx context.Context) *ResourceSchedulesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resource schedules delete params
func (o *ResourceSchedulesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resource schedules delete params
func (o *ResourceSchedulesDeleteParams) WithHTTPClient(client *http.Client) *ResourceSchedulesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resource schedules delete params
func (o *ResourceSchedulesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOverrideConflicts adds the overrideConflicts to the resource schedules delete params
func (o *ResourceSchedulesDeleteParams) WithOverrideConflicts(overrideConflicts *string) *ResourceSchedulesDeleteParams {
	o.SetOverrideConflicts(overrideConflicts)
	return o
}

// SetOverrideConflicts adds the overrideConflicts to the resource schedules delete params
func (o *ResourceSchedulesDeleteParams) SetOverrideConflicts(overrideConflicts *string) {
	o.OverrideConflicts = overrideConflicts
}

// WithResourceScheduleID adds the resourceScheduleID to the resource schedules delete params
func (o *ResourceSchedulesDeleteParams) WithResourceScheduleID(resourceScheduleID string) *ResourceSchedulesDeleteParams {
	o.SetResourceScheduleID(resourceScheduleID)
	return o
}

// SetResourceScheduleID adds the resourceScheduleId to the resource schedules delete params
func (o *ResourceSchedulesDeleteParams) SetResourceScheduleID(resourceScheduleID string) {
	o.ResourceScheduleID = resourceScheduleID
}

// WriteToRequest writes these params to a swagger request
func (o *ResourceSchedulesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.OverrideConflicts != nil {

		// query param overrideConflicts
		var qrOverrideConflicts string

		if o.OverrideConflicts != nil {
			qrOverrideConflicts = *o.OverrideConflicts
		}
		qOverrideConflicts := qrOverrideConflicts
		if qOverrideConflicts != "" {

			if err := r.SetQueryParam("overrideConflicts", qOverrideConflicts); err != nil {
				return err
			}
		}
	}

	// path param resourceScheduleId
	if err := r.SetPathParam("resourceScheduleId", o.ResourceScheduleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
