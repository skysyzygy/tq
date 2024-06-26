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

// NewAssociationsUpdateParams creates a new AssociationsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssociationsUpdateParams() *AssociationsUpdateParams {
	return &AssociationsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAssociationsUpdateParamsWithTimeout creates a new AssociationsUpdateParams object
// with the ability to set a timeout on a request.
func NewAssociationsUpdateParamsWithTimeout(timeout time.Duration) *AssociationsUpdateParams {
	return &AssociationsUpdateParams{
		timeout: timeout,
	}
}

// NewAssociationsUpdateParamsWithContext creates a new AssociationsUpdateParams object
// with the ability to set a context for a request.
func NewAssociationsUpdateParamsWithContext(ctx context.Context) *AssociationsUpdateParams {
	return &AssociationsUpdateParams{
		Context: ctx,
	}
}

// NewAssociationsUpdateParamsWithHTTPClient creates a new AssociationsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssociationsUpdateParamsWithHTTPClient(client *http.Client) *AssociationsUpdateParams {
	return &AssociationsUpdateParams{
		HTTPClient: client,
	}
}

/*
AssociationsUpdateParams contains all the parameters to send to the API endpoint

	for the associations update operation.

	Typically these are written to a http.Request.
*/
type AssociationsUpdateParams struct {

	// Association.
	Association *models.Association

	// AssociationID.
	AssociationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the associations update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssociationsUpdateParams) WithDefaults() *AssociationsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the associations update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssociationsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the associations update params
func (o *AssociationsUpdateParams) WithTimeout(timeout time.Duration) *AssociationsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the associations update params
func (o *AssociationsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the associations update params
func (o *AssociationsUpdateParams) WithContext(ctx context.Context) *AssociationsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the associations update params
func (o *AssociationsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the associations update params
func (o *AssociationsUpdateParams) WithHTTPClient(client *http.Client) *AssociationsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the associations update params
func (o *AssociationsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssociation adds the association to the associations update params
func (o *AssociationsUpdateParams) WithAssociation(association *models.Association) *AssociationsUpdateParams {
	o.SetAssociation(association)
	return o
}

// SetAssociation adds the association to the associations update params
func (o *AssociationsUpdateParams) SetAssociation(association *models.Association) {
	o.Association = association
}

// WithAssociationID adds the associationID to the associations update params
func (o *AssociationsUpdateParams) WithAssociationID(associationID string) *AssociationsUpdateParams {
	o.SetAssociationID(associationID)
	return o
}

// SetAssociationID adds the associationId to the associations update params
func (o *AssociationsUpdateParams) SetAssociationID(associationID string) {
	o.AssociationID = associationID
}

// WriteToRequest writes these params to a swagger request
func (o *AssociationsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Association != nil {
		if err := r.SetBodyParam(o.Association); err != nil {
			return err
		}
	}

	// path param associationId
	if err := r.SetPathParam("associationId", o.AssociationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
