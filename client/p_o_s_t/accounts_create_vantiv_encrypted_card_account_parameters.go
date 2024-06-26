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

// NewAccountsCreateVantivEncryptedCardAccountParams creates a new AccountsCreateVantivEncryptedCardAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountsCreateVantivEncryptedCardAccountParams() *AccountsCreateVantivEncryptedCardAccountParams {
	return &AccountsCreateVantivEncryptedCardAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountsCreateVantivEncryptedCardAccountParamsWithTimeout creates a new AccountsCreateVantivEncryptedCardAccountParams object
// with the ability to set a timeout on a request.
func NewAccountsCreateVantivEncryptedCardAccountParamsWithTimeout(timeout time.Duration) *AccountsCreateVantivEncryptedCardAccountParams {
	return &AccountsCreateVantivEncryptedCardAccountParams{
		timeout: timeout,
	}
}

// NewAccountsCreateVantivEncryptedCardAccountParamsWithContext creates a new AccountsCreateVantivEncryptedCardAccountParams object
// with the ability to set a context for a request.
func NewAccountsCreateVantivEncryptedCardAccountParamsWithContext(ctx context.Context) *AccountsCreateVantivEncryptedCardAccountParams {
	return &AccountsCreateVantivEncryptedCardAccountParams{
		Context: ctx,
	}
}

// NewAccountsCreateVantivEncryptedCardAccountParamsWithHTTPClient creates a new AccountsCreateVantivEncryptedCardAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccountsCreateVantivEncryptedCardAccountParamsWithHTTPClient(client *http.Client) *AccountsCreateVantivEncryptedCardAccountParams {
	return &AccountsCreateVantivEncryptedCardAccountParams{
		HTTPClient: client,
	}
}

/*
AccountsCreateVantivEncryptedCardAccountParams contains all the parameters to send to the API endpoint

	for the accounts create vantiv encrypted card account operation.

	Typically these are written to a http.Request.
*/
type AccountsCreateVantivEncryptedCardAccountParams struct {

	// Request.
	Request *models.CreateVantivEncryptedCardAccountRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the accounts create vantiv encrypted card account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountsCreateVantivEncryptedCardAccountParams) WithDefaults() *AccountsCreateVantivEncryptedCardAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the accounts create vantiv encrypted card account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountsCreateVantivEncryptedCardAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the accounts create vantiv encrypted card account params
func (o *AccountsCreateVantivEncryptedCardAccountParams) WithTimeout(timeout time.Duration) *AccountsCreateVantivEncryptedCardAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the accounts create vantiv encrypted card account params
func (o *AccountsCreateVantivEncryptedCardAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the accounts create vantiv encrypted card account params
func (o *AccountsCreateVantivEncryptedCardAccountParams) WithContext(ctx context.Context) *AccountsCreateVantivEncryptedCardAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the accounts create vantiv encrypted card account params
func (o *AccountsCreateVantivEncryptedCardAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the accounts create vantiv encrypted card account params
func (o *AccountsCreateVantivEncryptedCardAccountParams) WithHTTPClient(client *http.Client) *AccountsCreateVantivEncryptedCardAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the accounts create vantiv encrypted card account params
func (o *AccountsCreateVantivEncryptedCardAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the accounts create vantiv encrypted card account params
func (o *AccountsCreateVantivEncryptedCardAccountParams) WithRequest(request *models.CreateVantivEncryptedCardAccountRequest) *AccountsCreateVantivEncryptedCardAccountParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the accounts create vantiv encrypted card account params
func (o *AccountsCreateVantivEncryptedCardAccountParams) SetRequest(request *models.CreateVantivEncryptedCardAccountRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *AccountsCreateVantivEncryptedCardAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
