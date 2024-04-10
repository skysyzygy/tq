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

// NewAuditLogsGetAllParams creates a new AuditLogsGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuditLogsGetAllParams() *AuditLogsGetAllParams {
	return &AuditLogsGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuditLogsGetAllParamsWithTimeout creates a new AuditLogsGetAllParams object
// with the ability to set a timeout on a request.
func NewAuditLogsGetAllParamsWithTimeout(timeout time.Duration) *AuditLogsGetAllParams {
	return &AuditLogsGetAllParams{
		timeout: timeout,
	}
}

// NewAuditLogsGetAllParamsWithContext creates a new AuditLogsGetAllParams object
// with the ability to set a context for a request.
func NewAuditLogsGetAllParamsWithContext(ctx context.Context) *AuditLogsGetAllParams {
	return &AuditLogsGetAllParams{
		Context: ctx,
	}
}

// NewAuditLogsGetAllParamsWithHTTPClient creates a new AuditLogsGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuditLogsGetAllParamsWithHTTPClient(client *http.Client) *AuditLogsGetAllParams {
	return &AuditLogsGetAllParams{
		HTTPClient: client,
	}
}

/*
AuditLogsGetAllParams contains all the parameters to send to the API endpoint

	for the audit logs get all operation.

	Typically these are written to a http.Request.
*/
type AuditLogsGetAllParams struct {

	// Action.
	Action *string

	// AuditEndTime.
	AuditEndTime *string

	// AuditStartTime.
	AuditStartTime *string

	// EntityIds.
	EntityIds *string

	// IncludeSubEntities.
	IncludeSubEntities *string

	// PageIndex.
	PageIndex *string

	// PageSize.
	PageSize *string

	// Table.
	Table *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the audit logs get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuditLogsGetAllParams) WithDefaults() *AuditLogsGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the audit logs get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuditLogsGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the audit logs get all params
func (o *AuditLogsGetAllParams) WithTimeout(timeout time.Duration) *AuditLogsGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the audit logs get all params
func (o *AuditLogsGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the audit logs get all params
func (o *AuditLogsGetAllParams) WithContext(ctx context.Context) *AuditLogsGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the audit logs get all params
func (o *AuditLogsGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the audit logs get all params
func (o *AuditLogsGetAllParams) WithHTTPClient(client *http.Client) *AuditLogsGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the audit logs get all params
func (o *AuditLogsGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAction adds the action to the audit logs get all params
func (o *AuditLogsGetAllParams) WithAction(action *string) *AuditLogsGetAllParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the audit logs get all params
func (o *AuditLogsGetAllParams) SetAction(action *string) {
	o.Action = action
}

// WithAuditEndTime adds the auditEndTime to the audit logs get all params
func (o *AuditLogsGetAllParams) WithAuditEndTime(auditEndTime *string) *AuditLogsGetAllParams {
	o.SetAuditEndTime(auditEndTime)
	return o
}

// SetAuditEndTime adds the auditEndTime to the audit logs get all params
func (o *AuditLogsGetAllParams) SetAuditEndTime(auditEndTime *string) {
	o.AuditEndTime = auditEndTime
}

// WithAuditStartTime adds the auditStartTime to the audit logs get all params
func (o *AuditLogsGetAllParams) WithAuditStartTime(auditStartTime *string) *AuditLogsGetAllParams {
	o.SetAuditStartTime(auditStartTime)
	return o
}

// SetAuditStartTime adds the auditStartTime to the audit logs get all params
func (o *AuditLogsGetAllParams) SetAuditStartTime(auditStartTime *string) {
	o.AuditStartTime = auditStartTime
}

// WithEntityIds adds the entityIds to the audit logs get all params
func (o *AuditLogsGetAllParams) WithEntityIds(entityIds *string) *AuditLogsGetAllParams {
	o.SetEntityIds(entityIds)
	return o
}

// SetEntityIds adds the entityIds to the audit logs get all params
func (o *AuditLogsGetAllParams) SetEntityIds(entityIds *string) {
	o.EntityIds = entityIds
}

// WithIncludeSubEntities adds the includeSubEntities to the audit logs get all params
func (o *AuditLogsGetAllParams) WithIncludeSubEntities(includeSubEntities *string) *AuditLogsGetAllParams {
	o.SetIncludeSubEntities(includeSubEntities)
	return o
}

// SetIncludeSubEntities adds the includeSubEntities to the audit logs get all params
func (o *AuditLogsGetAllParams) SetIncludeSubEntities(includeSubEntities *string) {
	o.IncludeSubEntities = includeSubEntities
}

// WithPageIndex adds the pageIndex to the audit logs get all params
func (o *AuditLogsGetAllParams) WithPageIndex(pageIndex *string) *AuditLogsGetAllParams {
	o.SetPageIndex(pageIndex)
	return o
}

// SetPageIndex adds the pageIndex to the audit logs get all params
func (o *AuditLogsGetAllParams) SetPageIndex(pageIndex *string) {
	o.PageIndex = pageIndex
}

// WithPageSize adds the pageSize to the audit logs get all params
func (o *AuditLogsGetAllParams) WithPageSize(pageSize *string) *AuditLogsGetAllParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the audit logs get all params
func (o *AuditLogsGetAllParams) SetPageSize(pageSize *string) {
	o.PageSize = pageSize
}

// WithTable adds the table to the audit logs get all params
func (o *AuditLogsGetAllParams) WithTable(table *string) *AuditLogsGetAllParams {
	o.SetTable(table)
	return o
}

// SetTable adds the table to the audit logs get all params
func (o *AuditLogsGetAllParams) SetTable(table *string) {
	o.Table = table
}

// WriteToRequest writes these params to a swagger request
func (o *AuditLogsGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Action != nil {

		// query param action
		var qrAction string

		if o.Action != nil {
			qrAction = *o.Action
		}
		qAction := qrAction
		if qAction != "" {

			if err := r.SetQueryParam("action", qAction); err != nil {
				return err
			}
		}
	}

	if o.AuditEndTime != nil {

		// query param auditEndTime
		var qrAuditEndTime string

		if o.AuditEndTime != nil {
			qrAuditEndTime = *o.AuditEndTime
		}
		qAuditEndTime := qrAuditEndTime
		if qAuditEndTime != "" {

			if err := r.SetQueryParam("auditEndTime", qAuditEndTime); err != nil {
				return err
			}
		}
	}

	if o.AuditStartTime != nil {

		// query param auditStartTime
		var qrAuditStartTime string

		if o.AuditStartTime != nil {
			qrAuditStartTime = *o.AuditStartTime
		}
		qAuditStartTime := qrAuditStartTime
		if qAuditStartTime != "" {

			if err := r.SetQueryParam("auditStartTime", qAuditStartTime); err != nil {
				return err
			}
		}
	}

	if o.EntityIds != nil {

		// query param entityIds
		var qrEntityIds string

		if o.EntityIds != nil {
			qrEntityIds = *o.EntityIds
		}
		qEntityIds := qrEntityIds
		if qEntityIds != "" {

			if err := r.SetQueryParam("entityIds", qEntityIds); err != nil {
				return err
			}
		}
	}

	if o.IncludeSubEntities != nil {

		// query param includeSubEntities
		var qrIncludeSubEntities string

		if o.IncludeSubEntities != nil {
			qrIncludeSubEntities = *o.IncludeSubEntities
		}
		qIncludeSubEntities := qrIncludeSubEntities
		if qIncludeSubEntities != "" {

			if err := r.SetQueryParam("includeSubEntities", qIncludeSubEntities); err != nil {
				return err
			}
		}
	}

	if o.PageIndex != nil {

		// query param pageIndex
		var qrPageIndex string

		if o.PageIndex != nil {
			qrPageIndex = *o.PageIndex
		}
		qPageIndex := qrPageIndex
		if qPageIndex != "" {

			if err := r.SetQueryParam("pageIndex", qPageIndex); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize string

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := qrPageSize
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.Table != nil {

		// query param table
		var qrTable string

		if o.Table != nil {
			qrTable = *o.Table
		}
		qTable := qrTable
		if qTable != "" {

			if err := r.SetQueryParam("table", qTable); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
