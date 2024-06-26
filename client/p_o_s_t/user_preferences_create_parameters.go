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

// NewUserPreferencesCreateParams creates a new UserPreferencesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserPreferencesCreateParams() *UserPreferencesCreateParams {
	return &UserPreferencesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserPreferencesCreateParamsWithTimeout creates a new UserPreferencesCreateParams object
// with the ability to set a timeout on a request.
func NewUserPreferencesCreateParamsWithTimeout(timeout time.Duration) *UserPreferencesCreateParams {
	return &UserPreferencesCreateParams{
		timeout: timeout,
	}
}

// NewUserPreferencesCreateParamsWithContext creates a new UserPreferencesCreateParams object
// with the ability to set a context for a request.
func NewUserPreferencesCreateParamsWithContext(ctx context.Context) *UserPreferencesCreateParams {
	return &UserPreferencesCreateParams{
		Context: ctx,
	}
}

// NewUserPreferencesCreateParamsWithHTTPClient creates a new UserPreferencesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserPreferencesCreateParamsWithHTTPClient(client *http.Client) *UserPreferencesCreateParams {
	return &UserPreferencesCreateParams{
		HTTPClient: client,
	}
}

/*
UserPreferencesCreateParams contains all the parameters to send to the API endpoint

	for the user preferences create operation.

	Typically these are written to a http.Request.
*/
type UserPreferencesCreateParams struct {

	// UserPreference.
	UserPreference *models.UserPreference

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user preferences create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserPreferencesCreateParams) WithDefaults() *UserPreferencesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user preferences create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserPreferencesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user preferences create params
func (o *UserPreferencesCreateParams) WithTimeout(timeout time.Duration) *UserPreferencesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user preferences create params
func (o *UserPreferencesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user preferences create params
func (o *UserPreferencesCreateParams) WithContext(ctx context.Context) *UserPreferencesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user preferences create params
func (o *UserPreferencesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user preferences create params
func (o *UserPreferencesCreateParams) WithHTTPClient(client *http.Client) *UserPreferencesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user preferences create params
func (o *UserPreferencesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserPreference adds the userPreference to the user preferences create params
func (o *UserPreferencesCreateParams) WithUserPreference(userPreference *models.UserPreference) *UserPreferencesCreateParams {
	o.SetUserPreference(userPreference)
	return o
}

// SetUserPreference adds the userPreference to the user preferences create params
func (o *UserPreferencesCreateParams) SetUserPreference(userPreference *models.UserPreference) {
	o.UserPreference = userPreference
}

// WriteToRequest writes these params to a swagger request
func (o *UserPreferencesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.UserPreference != nil {
		if err := r.SetBodyParam(o.UserPreference); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
