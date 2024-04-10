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

// NewSurveyResponsesUpdateParams creates a new SurveyResponsesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSurveyResponsesUpdateParams() *SurveyResponsesUpdateParams {
	return &SurveyResponsesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSurveyResponsesUpdateParamsWithTimeout creates a new SurveyResponsesUpdateParams object
// with the ability to set a timeout on a request.
func NewSurveyResponsesUpdateParamsWithTimeout(timeout time.Duration) *SurveyResponsesUpdateParams {
	return &SurveyResponsesUpdateParams{
		timeout: timeout,
	}
}

// NewSurveyResponsesUpdateParamsWithContext creates a new SurveyResponsesUpdateParams object
// with the ability to set a context for a request.
func NewSurveyResponsesUpdateParamsWithContext(ctx context.Context) *SurveyResponsesUpdateParams {
	return &SurveyResponsesUpdateParams{
		Context: ctx,
	}
}

// NewSurveyResponsesUpdateParamsWithHTTPClient creates a new SurveyResponsesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSurveyResponsesUpdateParamsWithHTTPClient(client *http.Client) *SurveyResponsesUpdateParams {
	return &SurveyResponsesUpdateParams{
		HTTPClient: client,
	}
}

/*
SurveyResponsesUpdateParams contains all the parameters to send to the API endpoint

	for the survey responses update operation.

	Typically these are written to a http.Request.
*/
type SurveyResponsesUpdateParams struct {

	// SurveyResponse.
	SurveyResponse *models.SurveyResponse

	// SurveyResponseID.
	SurveyResponseID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the survey responses update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SurveyResponsesUpdateParams) WithDefaults() *SurveyResponsesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the survey responses update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SurveyResponsesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the survey responses update params
func (o *SurveyResponsesUpdateParams) WithTimeout(timeout time.Duration) *SurveyResponsesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the survey responses update params
func (o *SurveyResponsesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the survey responses update params
func (o *SurveyResponsesUpdateParams) WithContext(ctx context.Context) *SurveyResponsesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the survey responses update params
func (o *SurveyResponsesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the survey responses update params
func (o *SurveyResponsesUpdateParams) WithHTTPClient(client *http.Client) *SurveyResponsesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the survey responses update params
func (o *SurveyResponsesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSurveyResponse adds the surveyResponse to the survey responses update params
func (o *SurveyResponsesUpdateParams) WithSurveyResponse(surveyResponse *models.SurveyResponse) *SurveyResponsesUpdateParams {
	o.SetSurveyResponse(surveyResponse)
	return o
}

// SetSurveyResponse adds the surveyResponse to the survey responses update params
func (o *SurveyResponsesUpdateParams) SetSurveyResponse(surveyResponse *models.SurveyResponse) {
	o.SurveyResponse = surveyResponse
}

// WithSurveyResponseID adds the surveyResponseID to the survey responses update params
func (o *SurveyResponsesUpdateParams) WithSurveyResponseID(surveyResponseID string) *SurveyResponsesUpdateParams {
	o.SetSurveyResponseID(surveyResponseID)
	return o
}

// SetSurveyResponseID adds the surveyResponseId to the survey responses update params
func (o *SurveyResponsesUpdateParams) SetSurveyResponseID(surveyResponseID string) {
	o.SurveyResponseID = surveyResponseID
}

// WriteToRequest writes these params to a swagger request
func (o *SurveyResponsesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.SurveyResponse != nil {
		if err := r.SetBodyParam(o.SurveyResponse); err != nil {
			return err
		}
	}

	// path param surveyResponseId
	if err := r.SetPathParam("surveyResponseId", o.SurveyResponseID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
