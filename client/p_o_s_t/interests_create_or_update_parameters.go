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

// NewInterestsCreateOrUpdateParams creates a new InterestsCreateOrUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInterestsCreateOrUpdateParams() *InterestsCreateOrUpdateParams {
	return &InterestsCreateOrUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInterestsCreateOrUpdateParamsWithTimeout creates a new InterestsCreateOrUpdateParams object
// with the ability to set a timeout on a request.
func NewInterestsCreateOrUpdateParamsWithTimeout(timeout time.Duration) *InterestsCreateOrUpdateParams {
	return &InterestsCreateOrUpdateParams{
		timeout: timeout,
	}
}

// NewInterestsCreateOrUpdateParamsWithContext creates a new InterestsCreateOrUpdateParams object
// with the ability to set a context for a request.
func NewInterestsCreateOrUpdateParamsWithContext(ctx context.Context) *InterestsCreateOrUpdateParams {
	return &InterestsCreateOrUpdateParams{
		Context: ctx,
	}
}

// NewInterestsCreateOrUpdateParamsWithHTTPClient creates a new InterestsCreateOrUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewInterestsCreateOrUpdateParamsWithHTTPClient(client *http.Client) *InterestsCreateOrUpdateParams {
	return &InterestsCreateOrUpdateParams{
		HTTPClient: client,
	}
}

/*
InterestsCreateOrUpdateParams contains all the parameters to send to the API endpoint

	for the interests create or update operation.

	Typically these are written to a http.Request.
*/
type InterestsCreateOrUpdateParams struct {

	// Interests.
	Interests []*models.Interest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the interests create or update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InterestsCreateOrUpdateParams) WithDefaults() *InterestsCreateOrUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the interests create or update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InterestsCreateOrUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the interests create or update params
func (o *InterestsCreateOrUpdateParams) WithTimeout(timeout time.Duration) *InterestsCreateOrUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the interests create or update params
func (o *InterestsCreateOrUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the interests create or update params
func (o *InterestsCreateOrUpdateParams) WithContext(ctx context.Context) *InterestsCreateOrUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the interests create or update params
func (o *InterestsCreateOrUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the interests create or update params
func (o *InterestsCreateOrUpdateParams) WithHTTPClient(client *http.Client) *InterestsCreateOrUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the interests create or update params
func (o *InterestsCreateOrUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInterests adds the interests to the interests create or update params
func (o *InterestsCreateOrUpdateParams) WithInterests(interests []*models.Interest) *InterestsCreateOrUpdateParams {
	o.SetInterests(interests)
	return o
}

// SetInterests adds the interests to the interests create or update params
func (o *InterestsCreateOrUpdateParams) SetInterests(interests []*models.Interest) {
	o.Interests = interests
}

// WriteToRequest writes these params to a swagger request
func (o *InterestsCreateOrUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Interests != nil {
		if err := r.SetBodyParam(o.Interests); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
