// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

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

// NewArtistsCreateParams creates a new ArtistsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewArtistsCreateParams() *ArtistsCreateParams {
	return &ArtistsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewArtistsCreateParamsWithTimeout creates a new ArtistsCreateParams object
// with the ability to set a timeout on a request.
func NewArtistsCreateParamsWithTimeout(timeout time.Duration) *ArtistsCreateParams {
	return &ArtistsCreateParams{
		timeout: timeout,
	}
}

// NewArtistsCreateParamsWithContext creates a new ArtistsCreateParams object
// with the ability to set a context for a request.
func NewArtistsCreateParamsWithContext(ctx context.Context) *ArtistsCreateParams {
	return &ArtistsCreateParams{
		Context: ctx,
	}
}

// NewArtistsCreateParamsWithHTTPClient creates a new ArtistsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewArtistsCreateParamsWithHTTPClient(client *http.Client) *ArtistsCreateParams {
	return &ArtistsCreateParams{
		HTTPClient: client,
	}
}

/*
ArtistsCreateParams contains all the parameters to send to the API endpoint

	for the artists create operation.

	Typically these are written to a http.Request.
*/
type ArtistsCreateParams struct {

	// Artist.
	Artist *models.Artist

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the artists create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ArtistsCreateParams) WithDefaults() *ArtistsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the artists create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ArtistsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the artists create params
func (o *ArtistsCreateParams) WithTimeout(timeout time.Duration) *ArtistsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the artists create params
func (o *ArtistsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the artists create params
func (o *ArtistsCreateParams) WithContext(ctx context.Context) *ArtistsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the artists create params
func (o *ArtistsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the artists create params
func (o *ArtistsCreateParams) WithHTTPClient(client *http.Client) *ArtistsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the artists create params
func (o *ArtistsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArtist adds the artist to the artists create params
func (o *ArtistsCreateParams) WithArtist(artist *models.Artist) *ArtistsCreateParams {
	o.SetArtist(artist)
	return o
}

// SetArtist adds the artist to the artists create params
func (o *ArtistsCreateParams) SetArtist(artist *models.Artist) {
	o.Artist = artist
}

// WriteToRequest writes these params to a swagger request
func (o *ArtistsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Artist != nil {
		if err := r.SetBodyParam(o.Artist); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
