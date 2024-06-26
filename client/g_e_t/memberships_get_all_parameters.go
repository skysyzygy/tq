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

// NewMembershipsGetAllParams creates a new MembershipsGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMembershipsGetAllParams() *MembershipsGetAllParams {
	return &MembershipsGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMembershipsGetAllParamsWithTimeout creates a new MembershipsGetAllParams object
// with the ability to set a timeout on a request.
func NewMembershipsGetAllParamsWithTimeout(timeout time.Duration) *MembershipsGetAllParams {
	return &MembershipsGetAllParams{
		timeout: timeout,
	}
}

// NewMembershipsGetAllParamsWithContext creates a new MembershipsGetAllParams object
// with the ability to set a context for a request.
func NewMembershipsGetAllParamsWithContext(ctx context.Context) *MembershipsGetAllParams {
	return &MembershipsGetAllParams{
		Context: ctx,
	}
}

// NewMembershipsGetAllParamsWithHTTPClient creates a new MembershipsGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewMembershipsGetAllParamsWithHTTPClient(client *http.Client) *MembershipsGetAllParams {
	return &MembershipsGetAllParams{
		HTTPClient: client,
	}
}

/*
MembershipsGetAllParams contains all the parameters to send to the API endpoint

	for the memberships get all operation.

	Typically these are written to a http.Request.
*/
type MembershipsGetAllParams struct {

	/* ConstituentID.

	   Id of constituent
	*/
	ConstituentID string

	/* IncludeAffiliations.

	   If true, will also return details for affiliates of the constituent
	*/
	IncludeAffiliations *string

	/* MembershipOrgIds.

	   Only return memberships for the specified organizations in the comma-delimited list.
	*/
	MembershipOrgIds *string

	/* OnlyShowCurrent.

	   Specifies if only current memberships should be returned.  Cannot be true if onlyShowDefault is true.
	*/
	OnlyShowCurrent *string

	/* OnlyShowDefault.

	   Specifies if only the default membership should be returned.  Cannot be true if onlyShowCurrent is true.
	*/
	OnlyShowDefault *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the memberships get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MembershipsGetAllParams) WithDefaults() *MembershipsGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the memberships get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MembershipsGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the memberships get all params
func (o *MembershipsGetAllParams) WithTimeout(timeout time.Duration) *MembershipsGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the memberships get all params
func (o *MembershipsGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the memberships get all params
func (o *MembershipsGetAllParams) WithContext(ctx context.Context) *MembershipsGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the memberships get all params
func (o *MembershipsGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the memberships get all params
func (o *MembershipsGetAllParams) WithHTTPClient(client *http.Client) *MembershipsGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the memberships get all params
func (o *MembershipsGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConstituentID adds the constituentID to the memberships get all params
func (o *MembershipsGetAllParams) WithConstituentID(constituentID string) *MembershipsGetAllParams {
	o.SetConstituentID(constituentID)
	return o
}

// SetConstituentID adds the constituentId to the memberships get all params
func (o *MembershipsGetAllParams) SetConstituentID(constituentID string) {
	o.ConstituentID = constituentID
}

// WithIncludeAffiliations adds the includeAffiliations to the memberships get all params
func (o *MembershipsGetAllParams) WithIncludeAffiliations(includeAffiliations *string) *MembershipsGetAllParams {
	o.SetIncludeAffiliations(includeAffiliations)
	return o
}

// SetIncludeAffiliations adds the includeAffiliations to the memberships get all params
func (o *MembershipsGetAllParams) SetIncludeAffiliations(includeAffiliations *string) {
	o.IncludeAffiliations = includeAffiliations
}

// WithMembershipOrgIds adds the membershipOrgIds to the memberships get all params
func (o *MembershipsGetAllParams) WithMembershipOrgIds(membershipOrgIds *string) *MembershipsGetAllParams {
	o.SetMembershipOrgIds(membershipOrgIds)
	return o
}

// SetMembershipOrgIds adds the membershipOrgIds to the memberships get all params
func (o *MembershipsGetAllParams) SetMembershipOrgIds(membershipOrgIds *string) {
	o.MembershipOrgIds = membershipOrgIds
}

// WithOnlyShowCurrent adds the onlyShowCurrent to the memberships get all params
func (o *MembershipsGetAllParams) WithOnlyShowCurrent(onlyShowCurrent *string) *MembershipsGetAllParams {
	o.SetOnlyShowCurrent(onlyShowCurrent)
	return o
}

// SetOnlyShowCurrent adds the onlyShowCurrent to the memberships get all params
func (o *MembershipsGetAllParams) SetOnlyShowCurrent(onlyShowCurrent *string) {
	o.OnlyShowCurrent = onlyShowCurrent
}

// WithOnlyShowDefault adds the onlyShowDefault to the memberships get all params
func (o *MembershipsGetAllParams) WithOnlyShowDefault(onlyShowDefault *string) *MembershipsGetAllParams {
	o.SetOnlyShowDefault(onlyShowDefault)
	return o
}

// SetOnlyShowDefault adds the onlyShowDefault to the memberships get all params
func (o *MembershipsGetAllParams) SetOnlyShowDefault(onlyShowDefault *string) {
	o.OnlyShowDefault = onlyShowDefault
}

// WriteToRequest writes these params to a swagger request
func (o *MembershipsGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param constituentId
	qrConstituentID := o.ConstituentID
	qConstituentID := qrConstituentID
	if qConstituentID != "" {

		if err := r.SetQueryParam("constituentId", qConstituentID); err != nil {
			return err
		}
	}

	if o.IncludeAffiliations != nil {

		// query param includeAffiliations
		var qrIncludeAffiliations string

		if o.IncludeAffiliations != nil {
			qrIncludeAffiliations = *o.IncludeAffiliations
		}
		qIncludeAffiliations := qrIncludeAffiliations
		if qIncludeAffiliations != "" {

			if err := r.SetQueryParam("includeAffiliations", qIncludeAffiliations); err != nil {
				return err
			}
		}
	}

	if o.MembershipOrgIds != nil {

		// query param membershipOrgIds
		var qrMembershipOrgIds string

		if o.MembershipOrgIds != nil {
			qrMembershipOrgIds = *o.MembershipOrgIds
		}
		qMembershipOrgIds := qrMembershipOrgIds
		if qMembershipOrgIds != "" {

			if err := r.SetQueryParam("membershipOrgIds", qMembershipOrgIds); err != nil {
				return err
			}
		}
	}

	if o.OnlyShowCurrent != nil {

		// query param onlyShowCurrent
		var qrOnlyShowCurrent string

		if o.OnlyShowCurrent != nil {
			qrOnlyShowCurrent = *o.OnlyShowCurrent
		}
		qOnlyShowCurrent := qrOnlyShowCurrent
		if qOnlyShowCurrent != "" {

			if err := r.SetQueryParam("onlyShowCurrent", qOnlyShowCurrent); err != nil {
				return err
			}
		}
	}

	if o.OnlyShowDefault != nil {

		// query param onlyShowDefault
		var qrOnlyShowDefault string

		if o.OnlyShowDefault != nil {
			qrOnlyShowDefault = *o.OnlyShowDefault
		}
		qOnlyShowDefault := qrOnlyShowDefault
		if qOnlyShowDefault != "" {

			if err := r.SetQueryParam("onlyShowDefault", qOnlyShowDefault); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
