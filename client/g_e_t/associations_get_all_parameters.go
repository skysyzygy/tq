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

// NewAssociationsGetAllParams creates a new AssociationsGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssociationsGetAllParams() *AssociationsGetAllParams {
	return &AssociationsGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAssociationsGetAllParamsWithTimeout creates a new AssociationsGetAllParams object
// with the ability to set a timeout on a request.
func NewAssociationsGetAllParamsWithTimeout(timeout time.Duration) *AssociationsGetAllParams {
	return &AssociationsGetAllParams{
		timeout: timeout,
	}
}

// NewAssociationsGetAllParamsWithContext creates a new AssociationsGetAllParams object
// with the ability to set a context for a request.
func NewAssociationsGetAllParamsWithContext(ctx context.Context) *AssociationsGetAllParams {
	return &AssociationsGetAllParams{
		Context: ctx,
	}
}

// NewAssociationsGetAllParamsWithHTTPClient creates a new AssociationsGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssociationsGetAllParamsWithHTTPClient(client *http.Client) *AssociationsGetAllParams {
	return &AssociationsGetAllParams{
		HTTPClient: client,
	}
}

/*
AssociationsGetAllParams contains all the parameters to send to the API endpoint

	for the associations get all operation.

	Typically these are written to a http.Request.
*/
type AssociationsGetAllParams struct {

	// AssociatedConstituentID.
	AssociatedConstituentID *string

	/* ConstituentID.

	   Limit results by constituent
	*/
	ConstituentID *string

	// EndActiveDate.
	EndActiveDate *string

	// StartActiveDate.
	StartActiveDate *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the associations get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssociationsGetAllParams) WithDefaults() *AssociationsGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the associations get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssociationsGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the associations get all params
func (o *AssociationsGetAllParams) WithTimeout(timeout time.Duration) *AssociationsGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the associations get all params
func (o *AssociationsGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the associations get all params
func (o *AssociationsGetAllParams) WithContext(ctx context.Context) *AssociationsGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the associations get all params
func (o *AssociationsGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the associations get all params
func (o *AssociationsGetAllParams) WithHTTPClient(client *http.Client) *AssociationsGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the associations get all params
func (o *AssociationsGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssociatedConstituentID adds the associatedConstituentID to the associations get all params
func (o *AssociationsGetAllParams) WithAssociatedConstituentID(associatedConstituentID *string) *AssociationsGetAllParams {
	o.SetAssociatedConstituentID(associatedConstituentID)
	return o
}

// SetAssociatedConstituentID adds the associatedConstituentId to the associations get all params
func (o *AssociationsGetAllParams) SetAssociatedConstituentID(associatedConstituentID *string) {
	o.AssociatedConstituentID = associatedConstituentID
}

// WithConstituentID adds the constituentID to the associations get all params
func (o *AssociationsGetAllParams) WithConstituentID(constituentID *string) *AssociationsGetAllParams {
	o.SetConstituentID(constituentID)
	return o
}

// SetConstituentID adds the constituentId to the associations get all params
func (o *AssociationsGetAllParams) SetConstituentID(constituentID *string) {
	o.ConstituentID = constituentID
}

// WithEndActiveDate adds the endActiveDate to the associations get all params
func (o *AssociationsGetAllParams) WithEndActiveDate(endActiveDate *string) *AssociationsGetAllParams {
	o.SetEndActiveDate(endActiveDate)
	return o
}

// SetEndActiveDate adds the endActiveDate to the associations get all params
func (o *AssociationsGetAllParams) SetEndActiveDate(endActiveDate *string) {
	o.EndActiveDate = endActiveDate
}

// WithStartActiveDate adds the startActiveDate to the associations get all params
func (o *AssociationsGetAllParams) WithStartActiveDate(startActiveDate *string) *AssociationsGetAllParams {
	o.SetStartActiveDate(startActiveDate)
	return o
}

// SetStartActiveDate adds the startActiveDate to the associations get all params
func (o *AssociationsGetAllParams) SetStartActiveDate(startActiveDate *string) {
	o.StartActiveDate = startActiveDate
}

// WriteToRequest writes these params to a swagger request
func (o *AssociationsGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AssociatedConstituentID != nil {

		// query param associatedConstituentId
		var qrAssociatedConstituentID string

		if o.AssociatedConstituentID != nil {
			qrAssociatedConstituentID = *o.AssociatedConstituentID
		}
		qAssociatedConstituentID := qrAssociatedConstituentID
		if qAssociatedConstituentID != "" {

			if err := r.SetQueryParam("associatedConstituentId", qAssociatedConstituentID); err != nil {
				return err
			}
		}
	}

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

	if o.EndActiveDate != nil {

		// query param endActiveDate
		var qrEndActiveDate string

		if o.EndActiveDate != nil {
			qrEndActiveDate = *o.EndActiveDate
		}
		qEndActiveDate := qrEndActiveDate
		if qEndActiveDate != "" {

			if err := r.SetQueryParam("endActiveDate", qEndActiveDate); err != nil {
				return err
			}
		}
	}

	if o.StartActiveDate != nil {

		// query param startActiveDate
		var qrStartActiveDate string

		if o.StartActiveDate != nil {
			qrStartActiveDate = *o.StartActiveDate
		}
		qStartActiveDate := qrStartActiveDate
		if qStartActiveDate != "" {

			if err := r.SetQueryParam("startActiveDate", qStartActiveDate); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
