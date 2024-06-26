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

// NewHoldCodesUpdateParams creates a new HoldCodesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHoldCodesUpdateParams() *HoldCodesUpdateParams {
	return &HoldCodesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHoldCodesUpdateParamsWithTimeout creates a new HoldCodesUpdateParams object
// with the ability to set a timeout on a request.
func NewHoldCodesUpdateParamsWithTimeout(timeout time.Duration) *HoldCodesUpdateParams {
	return &HoldCodesUpdateParams{
		timeout: timeout,
	}
}

// NewHoldCodesUpdateParamsWithContext creates a new HoldCodesUpdateParams object
// with the ability to set a context for a request.
func NewHoldCodesUpdateParamsWithContext(ctx context.Context) *HoldCodesUpdateParams {
	return &HoldCodesUpdateParams{
		Context: ctx,
	}
}

// NewHoldCodesUpdateParamsWithHTTPClient creates a new HoldCodesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewHoldCodesUpdateParamsWithHTTPClient(client *http.Client) *HoldCodesUpdateParams {
	return &HoldCodesUpdateParams{
		HTTPClient: client,
	}
}

/*
HoldCodesUpdateParams contains all the parameters to send to the API endpoint

	for the hold codes update operation.

	Typically these are written to a http.Request.
*/
type HoldCodesUpdateParams struct {

	/* HoldCode.

	   The updated resource to be saved
	*/
	HoldCode *models.HoldCode

	/* HoldCodeID.

	   The ID of the resource to be updated
	*/
	HoldCodeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the hold codes update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HoldCodesUpdateParams) WithDefaults() *HoldCodesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hold codes update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HoldCodesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the hold codes update params
func (o *HoldCodesUpdateParams) WithTimeout(timeout time.Duration) *HoldCodesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hold codes update params
func (o *HoldCodesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hold codes update params
func (o *HoldCodesUpdateParams) WithContext(ctx context.Context) *HoldCodesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hold codes update params
func (o *HoldCodesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hold codes update params
func (o *HoldCodesUpdateParams) WithHTTPClient(client *http.Client) *HoldCodesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hold codes update params
func (o *HoldCodesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHoldCode adds the holdCode to the hold codes update params
func (o *HoldCodesUpdateParams) WithHoldCode(holdCode *models.HoldCode) *HoldCodesUpdateParams {
	o.SetHoldCode(holdCode)
	return o
}

// SetHoldCode adds the holdCode to the hold codes update params
func (o *HoldCodesUpdateParams) SetHoldCode(holdCode *models.HoldCode) {
	o.HoldCode = holdCode
}

// WithHoldCodeID adds the holdCodeID to the hold codes update params
func (o *HoldCodesUpdateParams) WithHoldCodeID(holdCodeID string) *HoldCodesUpdateParams {
	o.SetHoldCodeID(holdCodeID)
	return o
}

// SetHoldCodeID adds the holdCodeId to the hold codes update params
func (o *HoldCodesUpdateParams) SetHoldCodeID(holdCodeID string) {
	o.HoldCodeID = holdCodeID
}

// WriteToRequest writes these params to a swagger request
func (o *HoldCodesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.HoldCode != nil {
		if err := r.SetBodyParam(o.HoldCode); err != nil {
			return err
		}
	}

	// path param holdCodeId
	if err := r.SetPathParam("holdCodeId", o.HoldCodeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
