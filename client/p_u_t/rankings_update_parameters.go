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

// NewRankingsUpdateParams creates a new RankingsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRankingsUpdateParams() *RankingsUpdateParams {
	return &RankingsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRankingsUpdateParamsWithTimeout creates a new RankingsUpdateParams object
// with the ability to set a timeout on a request.
func NewRankingsUpdateParamsWithTimeout(timeout time.Duration) *RankingsUpdateParams {
	return &RankingsUpdateParams{
		timeout: timeout,
	}
}

// NewRankingsUpdateParamsWithContext creates a new RankingsUpdateParams object
// with the ability to set a context for a request.
func NewRankingsUpdateParamsWithContext(ctx context.Context) *RankingsUpdateParams {
	return &RankingsUpdateParams{
		Context: ctx,
	}
}

// NewRankingsUpdateParamsWithHTTPClient creates a new RankingsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewRankingsUpdateParamsWithHTTPClient(client *http.Client) *RankingsUpdateParams {
	return &RankingsUpdateParams{
		HTTPClient: client,
	}
}

/*
RankingsUpdateParams contains all the parameters to send to the API endpoint

	for the rankings update operation.

	Typically these are written to a http.Request.
*/
type RankingsUpdateParams struct {

	// Ranking.
	Ranking *models.Ranking

	// RankingID.
	RankingID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rankings update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RankingsUpdateParams) WithDefaults() *RankingsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rankings update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RankingsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the rankings update params
func (o *RankingsUpdateParams) WithTimeout(timeout time.Duration) *RankingsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rankings update params
func (o *RankingsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rankings update params
func (o *RankingsUpdateParams) WithContext(ctx context.Context) *RankingsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rankings update params
func (o *RankingsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rankings update params
func (o *RankingsUpdateParams) WithHTTPClient(client *http.Client) *RankingsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rankings update params
func (o *RankingsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRanking adds the ranking to the rankings update params
func (o *RankingsUpdateParams) WithRanking(ranking *models.Ranking) *RankingsUpdateParams {
	o.SetRanking(ranking)
	return o
}

// SetRanking adds the ranking to the rankings update params
func (o *RankingsUpdateParams) SetRanking(ranking *models.Ranking) {
	o.Ranking = ranking
}

// WithRankingID adds the rankingID to the rankings update params
func (o *RankingsUpdateParams) WithRankingID(rankingID string) *RankingsUpdateParams {
	o.SetRankingID(rankingID)
	return o
}

// SetRankingID adds the rankingId to the rankings update params
func (o *RankingsUpdateParams) SetRankingID(rankingID string) {
	o.RankingID = rankingID
}

// WriteToRequest writes these params to a swagger request
func (o *RankingsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Ranking != nil {
		if err := r.SetBodyParam(o.Ranking); err != nil {
			return err
		}
	}

	// path param rankingId
	if err := r.SetPathParam("rankingId", o.RankingID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
