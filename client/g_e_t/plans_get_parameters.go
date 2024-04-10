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

// NewPlansGetParams creates a new PlansGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPlansGetParams() *PlansGetParams {
	return &PlansGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPlansGetParamsWithTimeout creates a new PlansGetParams object
// with the ability to set a timeout on a request.
func NewPlansGetParamsWithTimeout(timeout time.Duration) *PlansGetParams {
	return &PlansGetParams{
		timeout: timeout,
	}
}

// NewPlansGetParamsWithContext creates a new PlansGetParams object
// with the ability to set a context for a request.
func NewPlansGetParamsWithContext(ctx context.Context) *PlansGetParams {
	return &PlansGetParams{
		Context: ctx,
	}
}

// NewPlansGetParamsWithHTTPClient creates a new PlansGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPlansGetParamsWithHTTPClient(client *http.Client) *PlansGetParams {
	return &PlansGetParams{
		HTTPClient: client,
	}
}

/*
PlansGetParams contains all the parameters to send to the API endpoint

	for the plans get operation.

	Typically these are written to a http.Request.
*/
type PlansGetParams struct {

	// PlanID.
	PlanID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plans get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PlansGetParams) WithDefaults() *PlansGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plans get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PlansGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plans get params
func (o *PlansGetParams) WithTimeout(timeout time.Duration) *PlansGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plans get params
func (o *PlansGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plans get params
func (o *PlansGetParams) WithContext(ctx context.Context) *PlansGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plans get params
func (o *PlansGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plans get params
func (o *PlansGetParams) WithHTTPClient(client *http.Client) *PlansGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plans get params
func (o *PlansGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPlanID adds the planID to the plans get params
func (o *PlansGetParams) WithPlanID(planID string) *PlansGetParams {
	o.SetPlanID(planID)
	return o
}

// SetPlanID adds the planId to the plans get params
func (o *PlansGetParams) SetPlanID(planID string) {
	o.PlanID = planID
}

// WriteToRequest writes these params to a swagger request
func (o *PlansGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param planId
	if err := r.SetPathParam("planId", o.PlanID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
