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

// NewRankingsGetParams creates a new RankingsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRankingsGetParams() *RankingsGetParams {
	return &RankingsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRankingsGetParamsWithTimeout creates a new RankingsGetParams object
// with the ability to set a timeout on a request.
func NewRankingsGetParamsWithTimeout(timeout time.Duration) *RankingsGetParams {
	return &RankingsGetParams{
		timeout: timeout,
	}
}

// NewRankingsGetParamsWithContext creates a new RankingsGetParams object
// with the ability to set a context for a request.
func NewRankingsGetParamsWithContext(ctx context.Context) *RankingsGetParams {
	return &RankingsGetParams{
		Context: ctx,
	}
}

// NewRankingsGetParamsWithHTTPClient creates a new RankingsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewRankingsGetParamsWithHTTPClient(client *http.Client) *RankingsGetParams {
	return &RankingsGetParams{
		HTTPClient: client,
	}
}

/*
RankingsGetParams contains all the parameters to send to the API endpoint

	for the rankings get operation.

	Typically these are written to a http.Request.
*/
type RankingsGetParams struct {

	// RankingID.
	RankingID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rankings get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RankingsGetParams) WithDefaults() *RankingsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rankings get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RankingsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the rankings get params
func (o *RankingsGetParams) WithTimeout(timeout time.Duration) *RankingsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rankings get params
func (o *RankingsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rankings get params
func (o *RankingsGetParams) WithContext(ctx context.Context) *RankingsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rankings get params
func (o *RankingsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rankings get params
func (o *RankingsGetParams) WithHTTPClient(client *http.Client) *RankingsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rankings get params
func (o *RankingsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRankingID adds the rankingID to the rankings get params
func (o *RankingsGetParams) WithRankingID(rankingID string) *RankingsGetParams {
	o.SetRankingID(rankingID)
	return o
}

// SetRankingID adds the rankingId to the rankings get params
func (o *RankingsGetParams) SetRankingID(rankingID string) {
	o.RankingID = rankingID
}

// WriteToRequest writes these params to a swagger request
func (o *RankingsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param rankingId
	if err := r.SetPathParam("rankingId", o.RankingID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
