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

// NewFeesGetAllParams creates a new FeesGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFeesGetAllParams() *FeesGetAllParams {
	return &FeesGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFeesGetAllParamsWithTimeout creates a new FeesGetAllParams object
// with the ability to set a timeout on a request.
func NewFeesGetAllParamsWithTimeout(timeout time.Duration) *FeesGetAllParams {
	return &FeesGetAllParams{
		timeout: timeout,
	}
}

// NewFeesGetAllParamsWithContext creates a new FeesGetAllParams object
// with the ability to set a context for a request.
func NewFeesGetAllParamsWithContext(ctx context.Context) *FeesGetAllParams {
	return &FeesGetAllParams{
		Context: ctx,
	}
}

// NewFeesGetAllParamsWithHTTPClient creates a new FeesGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewFeesGetAllParamsWithHTTPClient(client *http.Client) *FeesGetAllParams {
	return &FeesGetAllParams{
		HTTPClient: client,
	}
}

/*
FeesGetAllParams contains all the parameters to send to the API endpoint

	for the fees get all operation.

	Typically these are written to a http.Request.
*/
type FeesGetAllParams struct {

	// FeeCategoryIds.
	FeeCategoryIds *string

	// SeasonID.
	SeasonID *string

	// UserDefined.
	UserDefined *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fees get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FeesGetAllParams) WithDefaults() *FeesGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fees get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FeesGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the fees get all params
func (o *FeesGetAllParams) WithTimeout(timeout time.Duration) *FeesGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fees get all params
func (o *FeesGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fees get all params
func (o *FeesGetAllParams) WithContext(ctx context.Context) *FeesGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fees get all params
func (o *FeesGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fees get all params
func (o *FeesGetAllParams) WithHTTPClient(client *http.Client) *FeesGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fees get all params
func (o *FeesGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFeeCategoryIds adds the feeCategoryIds to the fees get all params
func (o *FeesGetAllParams) WithFeeCategoryIds(feeCategoryIds *string) *FeesGetAllParams {
	o.SetFeeCategoryIds(feeCategoryIds)
	return o
}

// SetFeeCategoryIds adds the feeCategoryIds to the fees get all params
func (o *FeesGetAllParams) SetFeeCategoryIds(feeCategoryIds *string) {
	o.FeeCategoryIds = feeCategoryIds
}

// WithSeasonID adds the seasonID to the fees get all params
func (o *FeesGetAllParams) WithSeasonID(seasonID *string) *FeesGetAllParams {
	o.SetSeasonID(seasonID)
	return o
}

// SetSeasonID adds the seasonId to the fees get all params
func (o *FeesGetAllParams) SetSeasonID(seasonID *string) {
	o.SeasonID = seasonID
}

// WithUserDefined adds the userDefined to the fees get all params
func (o *FeesGetAllParams) WithUserDefined(userDefined *string) *FeesGetAllParams {
	o.SetUserDefined(userDefined)
	return o
}

// SetUserDefined adds the userDefined to the fees get all params
func (o *FeesGetAllParams) SetUserDefined(userDefined *string) {
	o.UserDefined = userDefined
}

// WriteToRequest writes these params to a swagger request
func (o *FeesGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FeeCategoryIds != nil {

		// query param feeCategoryIds
		var qrFeeCategoryIds string

		if o.FeeCategoryIds != nil {
			qrFeeCategoryIds = *o.FeeCategoryIds
		}
		qFeeCategoryIds := qrFeeCategoryIds
		if qFeeCategoryIds != "" {

			if err := r.SetQueryParam("feeCategoryIds", qFeeCategoryIds); err != nil {
				return err
			}
		}
	}

	if o.SeasonID != nil {

		// query param seasonId
		var qrSeasonID string

		if o.SeasonID != nil {
			qrSeasonID = *o.SeasonID
		}
		qSeasonID := qrSeasonID
		if qSeasonID != "" {

			if err := r.SetQueryParam("seasonId", qSeasonID); err != nil {
				return err
			}
		}
	}

	if o.UserDefined != nil {

		// query param userDefined
		var qrUserDefined string

		if o.UserDefined != nil {
			qrUserDefined = *o.UserDefined
		}
		qUserDefined := qrUserDefined
		if qUserDefined != "" {

			if err := r.SetQueryParam("userDefined", qUserDefined); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
