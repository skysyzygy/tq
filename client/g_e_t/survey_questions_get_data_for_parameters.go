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

// NewSurveyQuestionsGetDataForParams creates a new SurveyQuestionsGetDataForParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSurveyQuestionsGetDataForParams() *SurveyQuestionsGetDataForParams {
	return &SurveyQuestionsGetDataForParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSurveyQuestionsGetDataForParamsWithTimeout creates a new SurveyQuestionsGetDataForParams object
// with the ability to set a timeout on a request.
func NewSurveyQuestionsGetDataForParamsWithTimeout(timeout time.Duration) *SurveyQuestionsGetDataForParams {
	return &SurveyQuestionsGetDataForParams{
		timeout: timeout,
	}
}

// NewSurveyQuestionsGetDataForParamsWithContext creates a new SurveyQuestionsGetDataForParams object
// with the ability to set a context for a request.
func NewSurveyQuestionsGetDataForParamsWithContext(ctx context.Context) *SurveyQuestionsGetDataForParams {
	return &SurveyQuestionsGetDataForParams{
		Context: ctx,
	}
}

// NewSurveyQuestionsGetDataForParamsWithHTTPClient creates a new SurveyQuestionsGetDataForParams object
// with the ability to set a custom HTTPClient for a request.
func NewSurveyQuestionsGetDataForParamsWithHTTPClient(client *http.Client) *SurveyQuestionsGetDataForParams {
	return &SurveyQuestionsGetDataForParams{
		HTTPClient: client,
	}
}

/*
SurveyQuestionsGetDataForParams contains all the parameters to send to the API endpoint

	for the survey questions get data for operation.

	Typically these are written to a http.Request.
*/
type SurveyQuestionsGetDataForParams struct {

	/* QuestionID.

	   The id of the survey question
	*/
	QuestionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the survey questions get data for params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SurveyQuestionsGetDataForParams) WithDefaults() *SurveyQuestionsGetDataForParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the survey questions get data for params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SurveyQuestionsGetDataForParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the survey questions get data for params
func (o *SurveyQuestionsGetDataForParams) WithTimeout(timeout time.Duration) *SurveyQuestionsGetDataForParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the survey questions get data for params
func (o *SurveyQuestionsGetDataForParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the survey questions get data for params
func (o *SurveyQuestionsGetDataForParams) WithContext(ctx context.Context) *SurveyQuestionsGetDataForParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the survey questions get data for params
func (o *SurveyQuestionsGetDataForParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the survey questions get data for params
func (o *SurveyQuestionsGetDataForParams) WithHTTPClient(client *http.Client) *SurveyQuestionsGetDataForParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the survey questions get data for params
func (o *SurveyQuestionsGetDataForParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuestionID adds the questionID to the survey questions get data for params
func (o *SurveyQuestionsGetDataForParams) WithQuestionID(questionID string) *SurveyQuestionsGetDataForParams {
	o.SetQuestionID(questionID)
	return o
}

// SetQuestionID adds the questionId to the survey questions get data for params
func (o *SurveyQuestionsGetDataForParams) SetQuestionID(questionID string) {
	o.QuestionID = questionID
}

// WriteToRequest writes these params to a swagger request
func (o *SurveyQuestionsGetDataForParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param questionId
	qrQuestionID := o.QuestionID
	qQuestionID := qrQuestionID
	if qQuestionID != "" {

		if err := r.SetQueryParam("questionId", qQuestionID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
