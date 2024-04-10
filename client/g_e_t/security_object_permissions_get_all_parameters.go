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

// NewSecurityObjectPermissionsGetAllParams creates a new SecurityObjectPermissionsGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityObjectPermissionsGetAllParams() *SecurityObjectPermissionsGetAllParams {
	return &SecurityObjectPermissionsGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityObjectPermissionsGetAllParamsWithTimeout creates a new SecurityObjectPermissionsGetAllParams object
// with the ability to set a timeout on a request.
func NewSecurityObjectPermissionsGetAllParamsWithTimeout(timeout time.Duration) *SecurityObjectPermissionsGetAllParams {
	return &SecurityObjectPermissionsGetAllParams{
		timeout: timeout,
	}
}

// NewSecurityObjectPermissionsGetAllParamsWithContext creates a new SecurityObjectPermissionsGetAllParams object
// with the ability to set a context for a request.
func NewSecurityObjectPermissionsGetAllParamsWithContext(ctx context.Context) *SecurityObjectPermissionsGetAllParams {
	return &SecurityObjectPermissionsGetAllParams{
		Context: ctx,
	}
}

// NewSecurityObjectPermissionsGetAllParamsWithHTTPClient creates a new SecurityObjectPermissionsGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityObjectPermissionsGetAllParamsWithHTTPClient(client *http.Client) *SecurityObjectPermissionsGetAllParams {
	return &SecurityObjectPermissionsGetAllParams{
		HTTPClient: client,
	}
}

/*
SecurityObjectPermissionsGetAllParams contains all the parameters to send to the API endpoint

	for the security object permissions get all operation.

	Typically these are written to a http.Request.
*/
type SecurityObjectPermissionsGetAllParams struct {

	/* ConstituencyID.

	   Filter object permissions by the specified constituency.
	*/
	ConstituencyID *string

	/* ObjectIds.

	   Get permissions for the specified (Comma separated list of objectIds) objects.
	*/
	ObjectIds *string

	/* Objectid.

	   Get permissions for a specific object id.
	*/
	Objectid *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security object permissions get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityObjectPermissionsGetAllParams) WithDefaults() *SecurityObjectPermissionsGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security object permissions get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityObjectPermissionsGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the security object permissions get all params
func (o *SecurityObjectPermissionsGetAllParams) WithTimeout(timeout time.Duration) *SecurityObjectPermissionsGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security object permissions get all params
func (o *SecurityObjectPermissionsGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security object permissions get all params
func (o *SecurityObjectPermissionsGetAllParams) WithContext(ctx context.Context) *SecurityObjectPermissionsGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security object permissions get all params
func (o *SecurityObjectPermissionsGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security object permissions get all params
func (o *SecurityObjectPermissionsGetAllParams) WithHTTPClient(client *http.Client) *SecurityObjectPermissionsGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security object permissions get all params
func (o *SecurityObjectPermissionsGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConstituencyID adds the constituencyID to the security object permissions get all params
func (o *SecurityObjectPermissionsGetAllParams) WithConstituencyID(constituencyID *string) *SecurityObjectPermissionsGetAllParams {
	o.SetConstituencyID(constituencyID)
	return o
}

// SetConstituencyID adds the constituencyId to the security object permissions get all params
func (o *SecurityObjectPermissionsGetAllParams) SetConstituencyID(constituencyID *string) {
	o.ConstituencyID = constituencyID
}

// WithObjectIds adds the objectIds to the security object permissions get all params
func (o *SecurityObjectPermissionsGetAllParams) WithObjectIds(objectIds *string) *SecurityObjectPermissionsGetAllParams {
	o.SetObjectIds(objectIds)
	return o
}

// SetObjectIds adds the objectIds to the security object permissions get all params
func (o *SecurityObjectPermissionsGetAllParams) SetObjectIds(objectIds *string) {
	o.ObjectIds = objectIds
}

// WithObjectid adds the objectid to the security object permissions get all params
func (o *SecurityObjectPermissionsGetAllParams) WithObjectid(objectid *string) *SecurityObjectPermissionsGetAllParams {
	o.SetObjectid(objectid)
	return o
}

// SetObjectid adds the objectid to the security object permissions get all params
func (o *SecurityObjectPermissionsGetAllParams) SetObjectid(objectid *string) {
	o.Objectid = objectid
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityObjectPermissionsGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ConstituencyID != nil {

		// query param constituencyId
		var qrConstituencyID string

		if o.ConstituencyID != nil {
			qrConstituencyID = *o.ConstituencyID
		}
		qConstituencyID := qrConstituencyID
		if qConstituencyID != "" {

			if err := r.SetQueryParam("constituencyId", qConstituencyID); err != nil {
				return err
			}
		}
	}

	if o.ObjectIds != nil {

		// query param objectIds
		var qrObjectIds string

		if o.ObjectIds != nil {
			qrObjectIds = *o.ObjectIds
		}
		qObjectIds := qrObjectIds
		if qObjectIds != "" {

			if err := r.SetQueryParam("objectIds", qObjectIds); err != nil {
				return err
			}
		}
	}

	if o.Objectid != nil {

		// query param objectid
		var qrObjectid string

		if o.Objectid != nil {
			qrObjectid = *o.Objectid
		}
		qObjectid := qrObjectid
		if qObjectid != "" {

			if err := r.SetQueryParam("objectid", qObjectid); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
