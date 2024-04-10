// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// BatchTypeUserGroupGetAllReader is a Reader for the BatchTypeUserGroupGetAll structure.
type BatchTypeUserGroupGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BatchTypeUserGroupGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBatchTypeUserGroupGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /Finance/BatchTypeUserGroups] BatchTypeUserGroup_GetAll", response, response.Code())
	}
}

// NewBatchTypeUserGroupGetAllOK creates a BatchTypeUserGroupGetAllOK with default headers values
func NewBatchTypeUserGroupGetAllOK() *BatchTypeUserGroupGetAllOK {
	return &BatchTypeUserGroupGetAllOK{}
}

/*
BatchTypeUserGroupGetAllOK describes a response with status code 200, with default header values.

OK
*/
type BatchTypeUserGroupGetAllOK struct {
	Payload []*models.BatchTypeUserGroup
}

// IsSuccess returns true when this batch type user group get all o k response has a 2xx status code
func (o *BatchTypeUserGroupGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this batch type user group get all o k response has a 3xx status code
func (o *BatchTypeUserGroupGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this batch type user group get all o k response has a 4xx status code
func (o *BatchTypeUserGroupGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this batch type user group get all o k response has a 5xx status code
func (o *BatchTypeUserGroupGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this batch type user group get all o k response a status code equal to that given
func (o *BatchTypeUserGroupGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the batch type user group get all o k response
func (o *BatchTypeUserGroupGetAllOK) Code() int {
	return 200
}

func (o *BatchTypeUserGroupGetAllOK) Error() string {
	return fmt.Sprintf("[GET /Finance/BatchTypeUserGroups][%d] batchTypeUserGroupGetAllOK  %+v", 200, o.Payload)
}

func (o *BatchTypeUserGroupGetAllOK) String() string {
	return fmt.Sprintf("[GET /Finance/BatchTypeUserGroups][%d] batchTypeUserGroupGetAllOK  %+v", 200, o.Payload)
}

func (o *BatchTypeUserGroupGetAllOK) GetPayload() []*models.BatchTypeUserGroup {
	return o.Payload
}

func (o *BatchTypeUserGroupGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
