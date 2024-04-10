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

// NewCampaignDesignationsGetParams creates a new CampaignDesignationsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCampaignDesignationsGetParams() *CampaignDesignationsGetParams {
	return &CampaignDesignationsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCampaignDesignationsGetParamsWithTimeout creates a new CampaignDesignationsGetParams object
// with the ability to set a timeout on a request.
func NewCampaignDesignationsGetParamsWithTimeout(timeout time.Duration) *CampaignDesignationsGetParams {
	return &CampaignDesignationsGetParams{
		timeout: timeout,
	}
}

// NewCampaignDesignationsGetParamsWithContext creates a new CampaignDesignationsGetParams object
// with the ability to set a context for a request.
func NewCampaignDesignationsGetParamsWithContext(ctx context.Context) *CampaignDesignationsGetParams {
	return &CampaignDesignationsGetParams{
		Context: ctx,
	}
}

// NewCampaignDesignationsGetParamsWithHTTPClient creates a new CampaignDesignationsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCampaignDesignationsGetParamsWithHTTPClient(client *http.Client) *CampaignDesignationsGetParams {
	return &CampaignDesignationsGetParams{
		HTTPClient: client,
	}
}

/*
CampaignDesignationsGetParams contains all the parameters to send to the API endpoint

	for the campaign designations get operation.

	Typically these are written to a http.Request.
*/
type CampaignDesignationsGetParams struct {

	// CampaignDesignationID.
	CampaignDesignationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the campaign designations get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CampaignDesignationsGetParams) WithDefaults() *CampaignDesignationsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the campaign designations get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CampaignDesignationsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the campaign designations get params
func (o *CampaignDesignationsGetParams) WithTimeout(timeout time.Duration) *CampaignDesignationsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the campaign designations get params
func (o *CampaignDesignationsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the campaign designations get params
func (o *CampaignDesignationsGetParams) WithContext(ctx context.Context) *CampaignDesignationsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the campaign designations get params
func (o *CampaignDesignationsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the campaign designations get params
func (o *CampaignDesignationsGetParams) WithHTTPClient(client *http.Client) *CampaignDesignationsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the campaign designations get params
func (o *CampaignDesignationsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCampaignDesignationID adds the campaignDesignationID to the campaign designations get params
func (o *CampaignDesignationsGetParams) WithCampaignDesignationID(campaignDesignationID string) *CampaignDesignationsGetParams {
	o.SetCampaignDesignationID(campaignDesignationID)
	return o
}

// SetCampaignDesignationID adds the campaignDesignationId to the campaign designations get params
func (o *CampaignDesignationsGetParams) SetCampaignDesignationID(campaignDesignationID string) {
	o.CampaignDesignationID = campaignDesignationID
}

// WriteToRequest writes these params to a swagger request
func (o *CampaignDesignationsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param campaignDesignationId
	if err := r.SetPathParam("campaignDesignationId", o.CampaignDesignationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
