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

// NewModeOfSaleOffersUpdateParams creates a new ModeOfSaleOffersUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewModeOfSaleOffersUpdateParams() *ModeOfSaleOffersUpdateParams {
	return &ModeOfSaleOffersUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewModeOfSaleOffersUpdateParamsWithTimeout creates a new ModeOfSaleOffersUpdateParams object
// with the ability to set a timeout on a request.
func NewModeOfSaleOffersUpdateParamsWithTimeout(timeout time.Duration) *ModeOfSaleOffersUpdateParams {
	return &ModeOfSaleOffersUpdateParams{
		timeout: timeout,
	}
}

// NewModeOfSaleOffersUpdateParamsWithContext creates a new ModeOfSaleOffersUpdateParams object
// with the ability to set a context for a request.
func NewModeOfSaleOffersUpdateParamsWithContext(ctx context.Context) *ModeOfSaleOffersUpdateParams {
	return &ModeOfSaleOffersUpdateParams{
		Context: ctx,
	}
}

// NewModeOfSaleOffersUpdateParamsWithHTTPClient creates a new ModeOfSaleOffersUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewModeOfSaleOffersUpdateParamsWithHTTPClient(client *http.Client) *ModeOfSaleOffersUpdateParams {
	return &ModeOfSaleOffersUpdateParams{
		HTTPClient: client,
	}
}

/*
ModeOfSaleOffersUpdateParams contains all the parameters to send to the API endpoint

	for the mode of sale offers update operation.

	Typically these are written to a http.Request.
*/
type ModeOfSaleOffersUpdateParams struct {

	// ModeOfSaleOffer.
	ModeOfSaleOffer *models.ModeOfSaleOffer

	// ModeOfSaleOfferID.
	ModeOfSaleOfferID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the mode of sale offers update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModeOfSaleOffersUpdateParams) WithDefaults() *ModeOfSaleOffersUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the mode of sale offers update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModeOfSaleOffersUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the mode of sale offers update params
func (o *ModeOfSaleOffersUpdateParams) WithTimeout(timeout time.Duration) *ModeOfSaleOffersUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the mode of sale offers update params
func (o *ModeOfSaleOffersUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the mode of sale offers update params
func (o *ModeOfSaleOffersUpdateParams) WithContext(ctx context.Context) *ModeOfSaleOffersUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the mode of sale offers update params
func (o *ModeOfSaleOffersUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the mode of sale offers update params
func (o *ModeOfSaleOffersUpdateParams) WithHTTPClient(client *http.Client) *ModeOfSaleOffersUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the mode of sale offers update params
func (o *ModeOfSaleOffersUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModeOfSaleOffer adds the modeOfSaleOffer to the mode of sale offers update params
func (o *ModeOfSaleOffersUpdateParams) WithModeOfSaleOffer(modeOfSaleOffer *models.ModeOfSaleOffer) *ModeOfSaleOffersUpdateParams {
	o.SetModeOfSaleOffer(modeOfSaleOffer)
	return o
}

// SetModeOfSaleOffer adds the modeOfSaleOffer to the mode of sale offers update params
func (o *ModeOfSaleOffersUpdateParams) SetModeOfSaleOffer(modeOfSaleOffer *models.ModeOfSaleOffer) {
	o.ModeOfSaleOffer = modeOfSaleOffer
}

// WithModeOfSaleOfferID adds the modeOfSaleOfferID to the mode of sale offers update params
func (o *ModeOfSaleOffersUpdateParams) WithModeOfSaleOfferID(modeOfSaleOfferID string) *ModeOfSaleOffersUpdateParams {
	o.SetModeOfSaleOfferID(modeOfSaleOfferID)
	return o
}

// SetModeOfSaleOfferID adds the modeOfSaleOfferId to the mode of sale offers update params
func (o *ModeOfSaleOffersUpdateParams) SetModeOfSaleOfferID(modeOfSaleOfferID string) {
	o.ModeOfSaleOfferID = modeOfSaleOfferID
}

// WriteToRequest writes these params to a swagger request
func (o *ModeOfSaleOffersUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ModeOfSaleOffer != nil {
		if err := r.SetBodyParam(o.ModeOfSaleOffer); err != nil {
			return err
		}
	}

	// path param modeOfSaleOfferId
	if err := r.SetPathParam("modeOfSaleOfferId", o.ModeOfSaleOfferID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
