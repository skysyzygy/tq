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

// NewSurveyQuestionsGetSummariesParams creates a new SurveyQuestionsGetSummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSurveyQuestionsGetSummariesParams() *SurveyQuestionsGetSummariesParams {
	return &SurveyQuestionsGetSummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSurveyQuestionsGetSummariesParamsWithTimeout creates a new SurveyQuestionsGetSummariesParams object
// with the ability to set a timeout on a request.
func NewSurveyQuestionsGetSummariesParamsWithTimeout(timeout time.Duration) *SurveyQuestionsGetSummariesParams {
	return &SurveyQuestionsGetSummariesParams{
		timeout: timeout,
	}
}

// NewSurveyQuestionsGetSummariesParamsWithContext creates a new SurveyQuestionsGetSummariesParams object
// with the ability to set a context for a request.
func NewSurveyQuestionsGetSummariesParamsWithContext(ctx context.Context) *SurveyQuestionsGetSummariesParams {
	return &SurveyQuestionsGetSummariesParams{
		Context: ctx,
	}
}

// NewSurveyQuestionsGetSummariesParamsWithHTTPClient creates a new SurveyQuestionsGetSummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewSurveyQuestionsGetSummariesParamsWithHTTPClient(client *http.Client) *SurveyQuestionsGetSummariesParams {
	return &SurveyQuestionsGetSummariesParams{
		HTTPClient: client,
	}
}

/*
SurveyQuestionsGetSummariesParams contains all the parameters to send to the API endpoint

	for the survey questions get summaries operation.

	Typically these are written to a http.Request.
*/
type SurveyQuestionsGetSummariesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the survey questions get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SurveyQuestionsGetSummariesParams) WithDefaults() *SurveyQuestionsGetSummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the survey questions get summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SurveyQuestionsGetSummariesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the survey questions get summaries params
func (o *SurveyQuestionsGetSummariesParams) WithTimeout(timeout time.Duration) *SurveyQuestionsGetSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the survey questions get summaries params
func (o *SurveyQuestionsGetSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the survey questions get summaries params
func (o *SurveyQuestionsGetSummariesParams) WithContext(ctx context.Context) *SurveyQuestionsGetSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the survey questions get summaries params
func (o *SurveyQuestionsGetSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the survey questions get summaries params
func (o *SurveyQuestionsGetSummariesParams) WithHTTPClient(client *http.Client) *SurveyQuestionsGetSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the survey questions get summaries params
func (o *SurveyQuestionsGetSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SurveyQuestionsGetSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
