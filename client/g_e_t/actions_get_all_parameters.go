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

// NewActionsGetAllParams creates a new ActionsGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewActionsGetAllParams() *ActionsGetAllParams {
	return &ActionsGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewActionsGetAllParamsWithTimeout creates a new ActionsGetAllParams object
// with the ability to set a timeout on a request.
func NewActionsGetAllParamsWithTimeout(timeout time.Duration) *ActionsGetAllParams {
	return &ActionsGetAllParams{
		timeout: timeout,
	}
}

// NewActionsGetAllParamsWithContext creates a new ActionsGetAllParams object
// with the ability to set a context for a request.
func NewActionsGetAllParamsWithContext(ctx context.Context) *ActionsGetAllParams {
	return &ActionsGetAllParams{
		Context: ctx,
	}
}

// NewActionsGetAllParamsWithHTTPClient creates a new ActionsGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewActionsGetAllParamsWithHTTPClient(client *http.Client) *ActionsGetAllParams {
	return &ActionsGetAllParams{
		HTTPClient: client,
	}
}

/*
ActionsGetAllParams contains all the parameters to send to the API endpoint

	for the actions get all operation.

	Typically these are written to a http.Request.
*/
type ActionsGetAllParams struct {

	/* ConstituentID.

	   Limit results by constituent
	*/
	ConstituentID *string

	// IssueID.
	IssueID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the actions get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActionsGetAllParams) WithDefaults() *ActionsGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the actions get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActionsGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the actions get all params
func (o *ActionsGetAllParams) WithTimeout(timeout time.Duration) *ActionsGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the actions get all params
func (o *ActionsGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the actions get all params
func (o *ActionsGetAllParams) WithContext(ctx context.Context) *ActionsGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the actions get all params
func (o *ActionsGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the actions get all params
func (o *ActionsGetAllParams) WithHTTPClient(client *http.Client) *ActionsGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the actions get all params
func (o *ActionsGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConstituentID adds the constituentID to the actions get all params
func (o *ActionsGetAllParams) WithConstituentID(constituentID *string) *ActionsGetAllParams {
	o.SetConstituentID(constituentID)
	return o
}

// SetConstituentID adds the constituentId to the actions get all params
func (o *ActionsGetAllParams) SetConstituentID(constituentID *string) {
	o.ConstituentID = constituentID
}

// WithIssueID adds the issueID to the actions get all params
func (o *ActionsGetAllParams) WithIssueID(issueID *string) *ActionsGetAllParams {
	o.SetIssueID(issueID)
	return o
}

// SetIssueID adds the issueId to the actions get all params
func (o *ActionsGetAllParams) SetIssueID(issueID *string) {
	o.IssueID = issueID
}

// WriteToRequest writes these params to a swagger request
func (o *ActionsGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ConstituentID != nil {

		// query param constituentId
		var qrConstituentID string

		if o.ConstituentID != nil {
			qrConstituentID = *o.ConstituentID
		}
		qConstituentID := qrConstituentID
		if qConstituentID != "" {

			if err := r.SetQueryParam("constituentId", qConstituentID); err != nil {
				return err
			}
		}
	}

	if o.IssueID != nil {

		// query param issueId
		var qrIssueID string

		if o.IssueID != nil {
			qrIssueID = *o.IssueID
		}
		qIssueID := qrIssueID
		if qIssueID != "" {

			if err := r.SetQueryParam("issueId", qIssueID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
