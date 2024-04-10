// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// BatchMaintenanceCreateReader is a Reader for the BatchMaintenanceCreate structure.
type BatchMaintenanceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BatchMaintenanceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBatchMaintenanceCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /Finance/BatchMaintenance] BatchMaintenance_Create", response, response.Code())
	}
}

// NewBatchMaintenanceCreateOK creates a BatchMaintenanceCreateOK with default headers values
func NewBatchMaintenanceCreateOK() *BatchMaintenanceCreateOK {
	return &BatchMaintenanceCreateOK{}
}

/*
BatchMaintenanceCreateOK describes a response with status code 200, with default header values.

OK
*/
type BatchMaintenanceCreateOK struct {
	Payload *models.Batch
}

// IsSuccess returns true when this batch maintenance create o k response has a 2xx status code
func (o *BatchMaintenanceCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this batch maintenance create o k response has a 3xx status code
func (o *BatchMaintenanceCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this batch maintenance create o k response has a 4xx status code
func (o *BatchMaintenanceCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this batch maintenance create o k response has a 5xx status code
func (o *BatchMaintenanceCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this batch maintenance create o k response a status code equal to that given
func (o *BatchMaintenanceCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the batch maintenance create o k response
func (o *BatchMaintenanceCreateOK) Code() int {
	return 200
}

func (o *BatchMaintenanceCreateOK) Error() string {
	return fmt.Sprintf("[POST /Finance/BatchMaintenance][%d] batchMaintenanceCreateOK  %+v", 200, o.Payload)
}

func (o *BatchMaintenanceCreateOK) String() string {
	return fmt.Sprintf("[POST /Finance/BatchMaintenance][%d] batchMaintenanceCreateOK  %+v", 200, o.Payload)
}

func (o *BatchMaintenanceCreateOK) GetPayload() *models.Batch {
	return o.Payload
}

func (o *BatchMaintenanceCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Batch)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
