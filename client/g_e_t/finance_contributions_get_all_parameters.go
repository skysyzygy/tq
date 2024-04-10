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

// NewFinanceContributionsGetAllParams creates a new FinanceContributionsGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFinanceContributionsGetAllParams() *FinanceContributionsGetAllParams {
	return &FinanceContributionsGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFinanceContributionsGetAllParamsWithTimeout creates a new FinanceContributionsGetAllParams object
// with the ability to set a timeout on a request.
func NewFinanceContributionsGetAllParamsWithTimeout(timeout time.Duration) *FinanceContributionsGetAllParams {
	return &FinanceContributionsGetAllParams{
		timeout: timeout,
	}
}

// NewFinanceContributionsGetAllParamsWithContext creates a new FinanceContributionsGetAllParams object
// with the ability to set a context for a request.
func NewFinanceContributionsGetAllParamsWithContext(ctx context.Context) *FinanceContributionsGetAllParams {
	return &FinanceContributionsGetAllParams{
		Context: ctx,
	}
}

// NewFinanceContributionsGetAllParamsWithHTTPClient creates a new FinanceContributionsGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewFinanceContributionsGetAllParamsWithHTTPClient(client *http.Client) *FinanceContributionsGetAllParams {
	return &FinanceContributionsGetAllParams{
		HTTPClient: client,
	}
}

/*
FinanceContributionsGetAllParams contains all the parameters to send to the API endpoint

	for the finance contributions get all operation.

	Typically these are written to a http.Request.
*/
type FinanceContributionsGetAllParams struct {

	// PlanID.
	PlanID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the finance contributions get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FinanceContributionsGetAllParams) WithDefaults() *FinanceContributionsGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the finance contributions get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FinanceContributionsGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the finance contributions get all params
func (o *FinanceContributionsGetAllParams) WithTimeout(timeout time.Duration) *FinanceContributionsGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the finance contributions get all params
func (o *FinanceContributionsGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the finance contributions get all params
func (o *FinanceContributionsGetAllParams) WithContext(ctx context.Context) *FinanceContributionsGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the finance contributions get all params
func (o *FinanceContributionsGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the finance contributions get all params
func (o *FinanceContributionsGetAllParams) WithHTTPClient(client *http.Client) *FinanceContributionsGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the finance contributions get all params
func (o *FinanceContributionsGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPlanID adds the planID to the finance contributions get all params
func (o *FinanceContributionsGetAllParams) WithPlanID(planID *string) *FinanceContributionsGetAllParams {
	o.SetPlanID(planID)
	return o
}

// SetPlanID adds the planId to the finance contributions get all params
func (o *FinanceContributionsGetAllParams) SetPlanID(planID *string) {
	o.PlanID = planID
}

// WriteToRequest writes these params to a swagger request
func (o *FinanceContributionsGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PlanID != nil {

		// query param planId
		var qrPlanID string

		if o.PlanID != nil {
			qrPlanID = *o.PlanID
		}
		qPlanID := qrPlanID
		if qPlanID != "" {

			if err := r.SetQueryParam("planId", qPlanID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
