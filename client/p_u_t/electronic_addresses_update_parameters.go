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

// NewElectronicAddressesUpdateParams creates a new ElectronicAddressesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewElectronicAddressesUpdateParams() *ElectronicAddressesUpdateParams {
	return &ElectronicAddressesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewElectronicAddressesUpdateParamsWithTimeout creates a new ElectronicAddressesUpdateParams object
// with the ability to set a timeout on a request.
func NewElectronicAddressesUpdateParamsWithTimeout(timeout time.Duration) *ElectronicAddressesUpdateParams {
	return &ElectronicAddressesUpdateParams{
		timeout: timeout,
	}
}

// NewElectronicAddressesUpdateParamsWithContext creates a new ElectronicAddressesUpdateParams object
// with the ability to set a context for a request.
func NewElectronicAddressesUpdateParamsWithContext(ctx context.Context) *ElectronicAddressesUpdateParams {
	return &ElectronicAddressesUpdateParams{
		Context: ctx,
	}
}

// NewElectronicAddressesUpdateParamsWithHTTPClient creates a new ElectronicAddressesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewElectronicAddressesUpdateParamsWithHTTPClient(client *http.Client) *ElectronicAddressesUpdateParams {
	return &ElectronicAddressesUpdateParams{
		HTTPClient: client,
	}
}

/*
ElectronicAddressesUpdateParams contains all the parameters to send to the API endpoint

	for the electronic addresses update operation.

	Typically these are written to a http.Request.
*/
type ElectronicAddressesUpdateParams struct {

	// ElectronicAddress.
	ElectronicAddress *models.ElectronicAddress

	// ElectronicAddressID.
	ElectronicAddressID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the electronic addresses update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ElectronicAddressesUpdateParams) WithDefaults() *ElectronicAddressesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the electronic addresses update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ElectronicAddressesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the electronic addresses update params
func (o *ElectronicAddressesUpdateParams) WithTimeout(timeout time.Duration) *ElectronicAddressesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the electronic addresses update params
func (o *ElectronicAddressesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the electronic addresses update params
func (o *ElectronicAddressesUpdateParams) WithContext(ctx context.Context) *ElectronicAddressesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the electronic addresses update params
func (o *ElectronicAddressesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the electronic addresses update params
func (o *ElectronicAddressesUpdateParams) WithHTTPClient(client *http.Client) *ElectronicAddressesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the electronic addresses update params
func (o *ElectronicAddressesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithElectronicAddress adds the electronicAddress to the electronic addresses update params
func (o *ElectronicAddressesUpdateParams) WithElectronicAddress(electronicAddress *models.ElectronicAddress) *ElectronicAddressesUpdateParams {
	o.SetElectronicAddress(electronicAddress)
	return o
}

// SetElectronicAddress adds the electronicAddress to the electronic addresses update params
func (o *ElectronicAddressesUpdateParams) SetElectronicAddress(electronicAddress *models.ElectronicAddress) {
	o.ElectronicAddress = electronicAddress
}

// WithElectronicAddressID adds the electronicAddressID to the electronic addresses update params
func (o *ElectronicAddressesUpdateParams) WithElectronicAddressID(electronicAddressID string) *ElectronicAddressesUpdateParams {
	o.SetElectronicAddressID(electronicAddressID)
	return o
}

// SetElectronicAddressID adds the electronicAddressId to the electronic addresses update params
func (o *ElectronicAddressesUpdateParams) SetElectronicAddressID(electronicAddressID string) {
	o.ElectronicAddressID = electronicAddressID
}

// WriteToRequest writes these params to a swagger request
func (o *ElectronicAddressesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ElectronicAddress != nil {
		if err := r.SetBodyParam(o.ElectronicAddress); err != nil {
			return err
		}
	}

	// path param electronicAddressId
	if err := r.SetPathParam("electronicAddressId", o.ElectronicAddressID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
