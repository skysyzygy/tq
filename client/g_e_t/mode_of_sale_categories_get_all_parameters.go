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

// NewModeOfSaleCategoriesGetAllParams creates a new ModeOfSaleCategoriesGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewModeOfSaleCategoriesGetAllParams() *ModeOfSaleCategoriesGetAllParams {
	return &ModeOfSaleCategoriesGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewModeOfSaleCategoriesGetAllParamsWithTimeout creates a new ModeOfSaleCategoriesGetAllParams object
// with the ability to set a timeout on a request.
func NewModeOfSaleCategoriesGetAllParamsWithTimeout(timeout time.Duration) *ModeOfSaleCategoriesGetAllParams {
	return &ModeOfSaleCategoriesGetAllParams{
		timeout: timeout,
	}
}

// NewModeOfSaleCategoriesGetAllParamsWithContext creates a new ModeOfSaleCategoriesGetAllParams object
// with the ability to set a context for a request.
func NewModeOfSaleCategoriesGetAllParamsWithContext(ctx context.Context) *ModeOfSaleCategoriesGetAllParams {
	return &ModeOfSaleCategoriesGetAllParams{
		Context: ctx,
	}
}

// NewModeOfSaleCategoriesGetAllParamsWithHTTPClient creates a new ModeOfSaleCategoriesGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewModeOfSaleCategoriesGetAllParamsWithHTTPClient(client *http.Client) *ModeOfSaleCategoriesGetAllParams {
	return &ModeOfSaleCategoriesGetAllParams{
		HTTPClient: client,
	}
}

/*
ModeOfSaleCategoriesGetAllParams contains all the parameters to send to the API endpoint

	for the mode of sale categories get all operation.

	Typically these are written to a http.Request.
*/
type ModeOfSaleCategoriesGetAllParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the mode of sale categories get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModeOfSaleCategoriesGetAllParams) WithDefaults() *ModeOfSaleCategoriesGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the mode of sale categories get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModeOfSaleCategoriesGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the mode of sale categories get all params
func (o *ModeOfSaleCategoriesGetAllParams) WithTimeout(timeout time.Duration) *ModeOfSaleCategoriesGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the mode of sale categories get all params
func (o *ModeOfSaleCategoriesGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the mode of sale categories get all params
func (o *ModeOfSaleCategoriesGetAllParams) WithContext(ctx context.Context) *ModeOfSaleCategoriesGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the mode of sale categories get all params
func (o *ModeOfSaleCategoriesGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the mode of sale categories get all params
func (o *ModeOfSaleCategoriesGetAllParams) WithHTTPClient(client *http.Client) *ModeOfSaleCategoriesGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the mode of sale categories get all params
func (o *ModeOfSaleCategoriesGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ModeOfSaleCategoriesGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
