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

// NewTemplatesRenderLoginCredentialsParams creates a new TemplatesRenderLoginCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTemplatesRenderLoginCredentialsParams() *TemplatesRenderLoginCredentialsParams {
	return &TemplatesRenderLoginCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTemplatesRenderLoginCredentialsParamsWithTimeout creates a new TemplatesRenderLoginCredentialsParams object
// with the ability to set a timeout on a request.
func NewTemplatesRenderLoginCredentialsParamsWithTimeout(timeout time.Duration) *TemplatesRenderLoginCredentialsParams {
	return &TemplatesRenderLoginCredentialsParams{
		timeout: timeout,
	}
}

// NewTemplatesRenderLoginCredentialsParamsWithContext creates a new TemplatesRenderLoginCredentialsParams object
// with the ability to set a context for a request.
func NewTemplatesRenderLoginCredentialsParamsWithContext(ctx context.Context) *TemplatesRenderLoginCredentialsParams {
	return &TemplatesRenderLoginCredentialsParams{
		Context: ctx,
	}
}

// NewTemplatesRenderLoginCredentialsParamsWithHTTPClient creates a new TemplatesRenderLoginCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewTemplatesRenderLoginCredentialsParamsWithHTTPClient(client *http.Client) *TemplatesRenderLoginCredentialsParams {
	return &TemplatesRenderLoginCredentialsParams{
		HTTPClient: client,
	}
}

/*
TemplatesRenderLoginCredentialsParams contains all the parameters to send to the API endpoint

	for the templates render login credentials operation.

	Typically these are written to a http.Request.
*/
type TemplatesRenderLoginCredentialsParams struct {

	/* LoginID.

	   Login id of the order to retrieve information from
	*/
	LoginID string

	/* Request.

	   Template body used to render the template
	*/
	Request *models.RenderRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the templates render login credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TemplatesRenderLoginCredentialsParams) WithDefaults() *TemplatesRenderLoginCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the templates render login credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TemplatesRenderLoginCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the templates render login credentials params
func (o *TemplatesRenderLoginCredentialsParams) WithTimeout(timeout time.Duration) *TemplatesRenderLoginCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the templates render login credentials params
func (o *TemplatesRenderLoginCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the templates render login credentials params
func (o *TemplatesRenderLoginCredentialsParams) WithContext(ctx context.Context) *TemplatesRenderLoginCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the templates render login credentials params
func (o *TemplatesRenderLoginCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the templates render login credentials params
func (o *TemplatesRenderLoginCredentialsParams) WithHTTPClient(client *http.Client) *TemplatesRenderLoginCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the templates render login credentials params
func (o *TemplatesRenderLoginCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLoginID adds the loginID to the templates render login credentials params
func (o *TemplatesRenderLoginCredentialsParams) WithLoginID(loginID string) *TemplatesRenderLoginCredentialsParams {
	o.SetLoginID(loginID)
	return o
}

// SetLoginID adds the loginId to the templates render login credentials params
func (o *TemplatesRenderLoginCredentialsParams) SetLoginID(loginID string) {
	o.LoginID = loginID
}

// WithRequest adds the request to the templates render login credentials params
func (o *TemplatesRenderLoginCredentialsParams) WithRequest(request *models.RenderRequest) *TemplatesRenderLoginCredentialsParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the templates render login credentials params
func (o *TemplatesRenderLoginCredentialsParams) SetRequest(request *models.RenderRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *TemplatesRenderLoginCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param loginId
	if err := r.SetPathParam("loginId", o.LoginID); err != nil {
		return err
	}
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
