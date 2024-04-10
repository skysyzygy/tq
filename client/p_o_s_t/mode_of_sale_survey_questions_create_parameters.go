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

// NewModeOfSaleSurveyQuestionsCreateParams creates a new ModeOfSaleSurveyQuestionsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewModeOfSaleSurveyQuestionsCreateParams() *ModeOfSaleSurveyQuestionsCreateParams {
	return &ModeOfSaleSurveyQuestionsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewModeOfSaleSurveyQuestionsCreateParamsWithTimeout creates a new ModeOfSaleSurveyQuestionsCreateParams object
// with the ability to set a timeout on a request.
func NewModeOfSaleSurveyQuestionsCreateParamsWithTimeout(timeout time.Duration) *ModeOfSaleSurveyQuestionsCreateParams {
	return &ModeOfSaleSurveyQuestionsCreateParams{
		timeout: timeout,
	}
}

// NewModeOfSaleSurveyQuestionsCreateParamsWithContext creates a new ModeOfSaleSurveyQuestionsCreateParams object
// with the ability to set a context for a request.
func NewModeOfSaleSurveyQuestionsCreateParamsWithContext(ctx context.Context) *ModeOfSaleSurveyQuestionsCreateParams {
	return &ModeOfSaleSurveyQuestionsCreateParams{
		Context: ctx,
	}
}

// NewModeOfSaleSurveyQuestionsCreateParamsWithHTTPClient creates a new ModeOfSaleSurveyQuestionsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewModeOfSaleSurveyQuestionsCreateParamsWithHTTPClient(client *http.Client) *ModeOfSaleSurveyQuestionsCreateParams {
	return &ModeOfSaleSurveyQuestionsCreateParams{
		HTTPClient: client,
	}
}

/*
ModeOfSaleSurveyQuestionsCreateParams contains all the parameters to send to the API endpoint

	for the mode of sale survey questions create operation.

	Typically these are written to a http.Request.
*/
type ModeOfSaleSurveyQuestionsCreateParams struct {

	// ModeOfSaleSurveyQuestion.
	ModeOfSaleSurveyQuestion *models.ModeOfSaleSurveyQuestion

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the mode of sale survey questions create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModeOfSaleSurveyQuestionsCreateParams) WithDefaults() *ModeOfSaleSurveyQuestionsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the mode of sale survey questions create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModeOfSaleSurveyQuestionsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the mode of sale survey questions create params
func (o *ModeOfSaleSurveyQuestionsCreateParams) WithTimeout(timeout time.Duration) *ModeOfSaleSurveyQuestionsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the mode of sale survey questions create params
func (o *ModeOfSaleSurveyQuestionsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the mode of sale survey questions create params
func (o *ModeOfSaleSurveyQuestionsCreateParams) WithContext(ctx context.Context) *ModeOfSaleSurveyQuestionsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the mode of sale survey questions create params
func (o *ModeOfSaleSurveyQuestionsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the mode of sale survey questions create params
func (o *ModeOfSaleSurveyQuestionsCreateParams) WithHTTPClient(client *http.Client) *ModeOfSaleSurveyQuestionsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the mode of sale survey questions create params
func (o *ModeOfSaleSurveyQuestionsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModeOfSaleSurveyQuestion adds the modeOfSaleSurveyQuestion to the mode of sale survey questions create params
func (o *ModeOfSaleSurveyQuestionsCreateParams) WithModeOfSaleSurveyQuestion(modeOfSaleSurveyQuestion *models.ModeOfSaleSurveyQuestion) *ModeOfSaleSurveyQuestionsCreateParams {
	o.SetModeOfSaleSurveyQuestion(modeOfSaleSurveyQuestion)
	return o
}

// SetModeOfSaleSurveyQuestion adds the modeOfSaleSurveyQuestion to the mode of sale survey questions create params
func (o *ModeOfSaleSurveyQuestionsCreateParams) SetModeOfSaleSurveyQuestion(modeOfSaleSurveyQuestion *models.ModeOfSaleSurveyQuestion) {
	o.ModeOfSaleSurveyQuestion = modeOfSaleSurveyQuestion
}

// WriteToRequest writes these params to a swagger request
func (o *ModeOfSaleSurveyQuestionsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ModeOfSaleSurveyQuestion != nil {
		if err := r.SetBodyParam(o.ModeOfSaleSurveyQuestion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
