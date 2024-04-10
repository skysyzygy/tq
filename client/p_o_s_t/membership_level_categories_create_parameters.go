// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

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

// NewMembershipLevelCategoriesCreateParams creates a new MembershipLevelCategoriesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMembershipLevelCategoriesCreateParams() *MembershipLevelCategoriesCreateParams {
	return &MembershipLevelCategoriesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMembershipLevelCategoriesCreateParamsWithTimeout creates a new MembershipLevelCategoriesCreateParams object
// with the ability to set a timeout on a request.
func NewMembershipLevelCategoriesCreateParamsWithTimeout(timeout time.Duration) *MembershipLevelCategoriesCreateParams {
	return &MembershipLevelCategoriesCreateParams{
		timeout: timeout,
	}
}

// NewMembershipLevelCategoriesCreateParamsWithContext creates a new MembershipLevelCategoriesCreateParams object
// with the ability to set a context for a request.
func NewMembershipLevelCategoriesCreateParamsWithContext(ctx context.Context) *MembershipLevelCategoriesCreateParams {
	return &MembershipLevelCategoriesCreateParams{
		Context: ctx,
	}
}

// NewMembershipLevelCategoriesCreateParamsWithHTTPClient creates a new MembershipLevelCategoriesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewMembershipLevelCategoriesCreateParamsWithHTTPClient(client *http.Client) *MembershipLevelCategoriesCreateParams {
	return &MembershipLevelCategoriesCreateParams{
		HTTPClient: client,
	}
}

/*
MembershipLevelCategoriesCreateParams contains all the parameters to send to the API endpoint

	for the membership level categories create operation.

	Typically these are written to a http.Request.
*/
type MembershipLevelCategoriesCreateParams struct {

	// Data.
	Data *models.MembershipLevelCategory

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the membership level categories create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MembershipLevelCategoriesCreateParams) WithDefaults() *MembershipLevelCategoriesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the membership level categories create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MembershipLevelCategoriesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the membership level categories create params
func (o *MembershipLevelCategoriesCreateParams) WithTimeout(timeout time.Duration) *MembershipLevelCategoriesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the membership level categories create params
func (o *MembershipLevelCategoriesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the membership level categories create params
func (o *MembershipLevelCategoriesCreateParams) WithContext(ctx context.Context) *MembershipLevelCategoriesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the membership level categories create params
func (o *MembershipLevelCategoriesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the membership level categories create params
func (o *MembershipLevelCategoriesCreateParams) WithHTTPClient(client *http.Client) *MembershipLevelCategoriesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the membership level categories create params
func (o *MembershipLevelCategoriesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the membership level categories create params
func (o *MembershipLevelCategoriesCreateParams) WithData(data *models.MembershipLevelCategory) *MembershipLevelCategoriesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the membership level categories create params
func (o *MembershipLevelCategoriesCreateParams) SetData(data *models.MembershipLevelCategory) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *MembershipLevelCategoriesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
