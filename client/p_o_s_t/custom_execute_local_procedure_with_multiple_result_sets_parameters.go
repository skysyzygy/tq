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

// NewCustomExecuteLocalProcedureWithMultipleResultSetsParams creates a new CustomExecuteLocalProcedureWithMultipleResultSetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomExecuteLocalProcedureWithMultipleResultSetsParams() *CustomExecuteLocalProcedureWithMultipleResultSetsParams {
	return &CustomExecuteLocalProcedureWithMultipleResultSetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomExecuteLocalProcedureWithMultipleResultSetsParamsWithTimeout creates a new CustomExecuteLocalProcedureWithMultipleResultSetsParams object
// with the ability to set a timeout on a request.
func NewCustomExecuteLocalProcedureWithMultipleResultSetsParamsWithTimeout(timeout time.Duration) *CustomExecuteLocalProcedureWithMultipleResultSetsParams {
	return &CustomExecuteLocalProcedureWithMultipleResultSetsParams{
		timeout: timeout,
	}
}

// NewCustomExecuteLocalProcedureWithMultipleResultSetsParamsWithContext creates a new CustomExecuteLocalProcedureWithMultipleResultSetsParams object
// with the ability to set a context for a request.
func NewCustomExecuteLocalProcedureWithMultipleResultSetsParamsWithContext(ctx context.Context) *CustomExecuteLocalProcedureWithMultipleResultSetsParams {
	return &CustomExecuteLocalProcedureWithMultipleResultSetsParams{
		Context: ctx,
	}
}

// NewCustomExecuteLocalProcedureWithMultipleResultSetsParamsWithHTTPClient creates a new CustomExecuteLocalProcedureWithMultipleResultSetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomExecuteLocalProcedureWithMultipleResultSetsParamsWithHTTPClient(client *http.Client) *CustomExecuteLocalProcedureWithMultipleResultSetsParams {
	return &CustomExecuteLocalProcedureWithMultipleResultSetsParams{
		HTTPClient: client,
	}
}

/*
CustomExecuteLocalProcedureWithMultipleResultSetsParams contains all the parameters to send to the API endpoint

	for the custom execute local procedure with multiple result sets operation.

	Typically these are written to a http.Request.
*/
type CustomExecuteLocalProcedureWithMultipleResultSetsParams struct {

	// LocalProcedureRequest.
	LocalProcedureRequest *models.LocalProcedureRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the custom execute local procedure with multiple result sets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomExecuteLocalProcedureWithMultipleResultSetsParams) WithDefaults() *CustomExecuteLocalProcedureWithMultipleResultSetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the custom execute local procedure with multiple result sets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomExecuteLocalProcedureWithMultipleResultSetsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the custom execute local procedure with multiple result sets params
func (o *CustomExecuteLocalProcedureWithMultipleResultSetsParams) WithTimeout(timeout time.Duration) *CustomExecuteLocalProcedureWithMultipleResultSetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom execute local procedure with multiple result sets params
func (o *CustomExecuteLocalProcedureWithMultipleResultSetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom execute local procedure with multiple result sets params
func (o *CustomExecuteLocalProcedureWithMultipleResultSetsParams) WithContext(ctx context.Context) *CustomExecuteLocalProcedureWithMultipleResultSetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom execute local procedure with multiple result sets params
func (o *CustomExecuteLocalProcedureWithMultipleResultSetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom execute local procedure with multiple result sets params
func (o *CustomExecuteLocalProcedureWithMultipleResultSetsParams) WithHTTPClient(client *http.Client) *CustomExecuteLocalProcedureWithMultipleResultSetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom execute local procedure with multiple result sets params
func (o *CustomExecuteLocalProcedureWithMultipleResultSetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLocalProcedureRequest adds the localProcedureRequest to the custom execute local procedure with multiple result sets params
func (o *CustomExecuteLocalProcedureWithMultipleResultSetsParams) WithLocalProcedureRequest(localProcedureRequest *models.LocalProcedureRequest) *CustomExecuteLocalProcedureWithMultipleResultSetsParams {
	o.SetLocalProcedureRequest(localProcedureRequest)
	return o
}

// SetLocalProcedureRequest adds the localProcedureRequest to the custom execute local procedure with multiple result sets params
func (o *CustomExecuteLocalProcedureWithMultipleResultSetsParams) SetLocalProcedureRequest(localProcedureRequest *models.LocalProcedureRequest) {
	o.LocalProcedureRequest = localProcedureRequest
}

// WriteToRequest writes these params to a swagger request
func (o *CustomExecuteLocalProcedureWithMultipleResultSetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.LocalProcedureRequest != nil {
		if err := r.SetBodyParam(o.LocalProcedureRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
