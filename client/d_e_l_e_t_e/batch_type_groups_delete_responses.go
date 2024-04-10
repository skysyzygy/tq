// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// BatchTypeGroupsDeleteReader is a Reader for the BatchTypeGroupsDelete structure.
type BatchTypeGroupsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BatchTypeGroupsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewBatchTypeGroupsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ReferenceData/BatchTypeGroups/{id}] BatchTypeGroups_Delete", response, response.Code())
	}
}

// NewBatchTypeGroupsDeleteNoContent creates a BatchTypeGroupsDeleteNoContent with default headers values
func NewBatchTypeGroupsDeleteNoContent() *BatchTypeGroupsDeleteNoContent {
	return &BatchTypeGroupsDeleteNoContent{}
}

/*
BatchTypeGroupsDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type BatchTypeGroupsDeleteNoContent struct {
}

// IsSuccess returns true when this batch type groups delete no content response has a 2xx status code
func (o *BatchTypeGroupsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this batch type groups delete no content response has a 3xx status code
func (o *BatchTypeGroupsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this batch type groups delete no content response has a 4xx status code
func (o *BatchTypeGroupsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this batch type groups delete no content response has a 5xx status code
func (o *BatchTypeGroupsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this batch type groups delete no content response a status code equal to that given
func (o *BatchTypeGroupsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the batch type groups delete no content response
func (o *BatchTypeGroupsDeleteNoContent) Code() int {
	return 204
}

func (o *BatchTypeGroupsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/BatchTypeGroups/{id}][%d] batchTypeGroupsDeleteNoContent ", 204)
}

func (o *BatchTypeGroupsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/BatchTypeGroups/{id}][%d] batchTypeGroupsDeleteNoContent ", 204)
}

func (o *BatchTypeGroupsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
