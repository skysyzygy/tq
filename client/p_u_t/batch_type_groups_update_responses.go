// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// BatchTypeGroupsUpdateReader is a Reader for the BatchTypeGroupsUpdate structure.
type BatchTypeGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BatchTypeGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBatchTypeGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ReferenceData/BatchTypeGroups/{id}] BatchTypeGroups_Update", response, response.Code())
	}
}

// NewBatchTypeGroupsUpdateOK creates a BatchTypeGroupsUpdateOK with default headers values
func NewBatchTypeGroupsUpdateOK() *BatchTypeGroupsUpdateOK {
	return &BatchTypeGroupsUpdateOK{}
}

/*
BatchTypeGroupsUpdateOK describes a response with status code 200, with default header values.

OK
*/
type BatchTypeGroupsUpdateOK struct {
	Payload *models.BatchTypeGroup
}

// IsSuccess returns true when this batch type groups update o k response has a 2xx status code
func (o *BatchTypeGroupsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this batch type groups update o k response has a 3xx status code
func (o *BatchTypeGroupsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this batch type groups update o k response has a 4xx status code
func (o *BatchTypeGroupsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this batch type groups update o k response has a 5xx status code
func (o *BatchTypeGroupsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this batch type groups update o k response a status code equal to that given
func (o *BatchTypeGroupsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the batch type groups update o k response
func (o *BatchTypeGroupsUpdateOK) Code() int {
	return 200
}

func (o *BatchTypeGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ReferenceData/BatchTypeGroups/{id}][%d] batchTypeGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *BatchTypeGroupsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ReferenceData/BatchTypeGroups/{id}][%d] batchTypeGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *BatchTypeGroupsUpdateOK) GetPayload() *models.BatchTypeGroup {
	return o.Payload
}

func (o *BatchTypeGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BatchTypeGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
