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

// NewGiftAidDeclarationsUpdateParams creates a new GiftAidDeclarationsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGiftAidDeclarationsUpdateParams() *GiftAidDeclarationsUpdateParams {
	return &GiftAidDeclarationsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGiftAidDeclarationsUpdateParamsWithTimeout creates a new GiftAidDeclarationsUpdateParams object
// with the ability to set a timeout on a request.
func NewGiftAidDeclarationsUpdateParamsWithTimeout(timeout time.Duration) *GiftAidDeclarationsUpdateParams {
	return &GiftAidDeclarationsUpdateParams{
		timeout: timeout,
	}
}

// NewGiftAidDeclarationsUpdateParamsWithContext creates a new GiftAidDeclarationsUpdateParams object
// with the ability to set a context for a request.
func NewGiftAidDeclarationsUpdateParamsWithContext(ctx context.Context) *GiftAidDeclarationsUpdateParams {
	return &GiftAidDeclarationsUpdateParams{
		Context: ctx,
	}
}

// NewGiftAidDeclarationsUpdateParamsWithHTTPClient creates a new GiftAidDeclarationsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGiftAidDeclarationsUpdateParamsWithHTTPClient(client *http.Client) *GiftAidDeclarationsUpdateParams {
	return &GiftAidDeclarationsUpdateParams{
		HTTPClient: client,
	}
}

/*
GiftAidDeclarationsUpdateParams contains all the parameters to send to the API endpoint

	for the gift aid declarations update operation.

	Typically these are written to a http.Request.
*/
type GiftAidDeclarationsUpdateParams struct {

	// GiftAidDeclaration.
	GiftAidDeclaration *models.GiftAidDeclaration

	// GiftAidDeclarationID.
	GiftAidDeclarationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the gift aid declarations update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GiftAidDeclarationsUpdateParams) WithDefaults() *GiftAidDeclarationsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the gift aid declarations update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GiftAidDeclarationsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the gift aid declarations update params
func (o *GiftAidDeclarationsUpdateParams) WithTimeout(timeout time.Duration) *GiftAidDeclarationsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gift aid declarations update params
func (o *GiftAidDeclarationsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gift aid declarations update params
func (o *GiftAidDeclarationsUpdateParams) WithContext(ctx context.Context) *GiftAidDeclarationsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gift aid declarations update params
func (o *GiftAidDeclarationsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gift aid declarations update params
func (o *GiftAidDeclarationsUpdateParams) WithHTTPClient(client *http.Client) *GiftAidDeclarationsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gift aid declarations update params
func (o *GiftAidDeclarationsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGiftAidDeclaration adds the giftAidDeclaration to the gift aid declarations update params
func (o *GiftAidDeclarationsUpdateParams) WithGiftAidDeclaration(giftAidDeclaration *models.GiftAidDeclaration) *GiftAidDeclarationsUpdateParams {
	o.SetGiftAidDeclaration(giftAidDeclaration)
	return o
}

// SetGiftAidDeclaration adds the giftAidDeclaration to the gift aid declarations update params
func (o *GiftAidDeclarationsUpdateParams) SetGiftAidDeclaration(giftAidDeclaration *models.GiftAidDeclaration) {
	o.GiftAidDeclaration = giftAidDeclaration
}

// WithGiftAidDeclarationID adds the giftAidDeclarationID to the gift aid declarations update params
func (o *GiftAidDeclarationsUpdateParams) WithGiftAidDeclarationID(giftAidDeclarationID string) *GiftAidDeclarationsUpdateParams {
	o.SetGiftAidDeclarationID(giftAidDeclarationID)
	return o
}

// SetGiftAidDeclarationID adds the giftAidDeclarationId to the gift aid declarations update params
func (o *GiftAidDeclarationsUpdateParams) SetGiftAidDeclarationID(giftAidDeclarationID string) {
	o.GiftAidDeclarationID = giftAidDeclarationID
}

// WriteToRequest writes these params to a swagger request
func (o *GiftAidDeclarationsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.GiftAidDeclaration != nil {
		if err := r.SetBodyParam(o.GiftAidDeclaration); err != nil {
			return err
		}
	}

	// path param giftAidDeclarationId
	if err := r.SetPathParam("giftAidDeclarationId", o.GiftAidDeclarationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
