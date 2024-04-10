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

// WorkerTypesGetReader is a Reader for the WorkerTypesGet structure.
type WorkerTypesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkerTypesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkerTypesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/WorkerTypes/{id}] WorkerTypes_Get", response, response.Code())
	}
}

// NewWorkerTypesGetOK creates a WorkerTypesGetOK with default headers values
func NewWorkerTypesGetOK() *WorkerTypesGetOK {
	return &WorkerTypesGetOK{}
}

/*
WorkerTypesGetOK describes a response with status code 200, with default header values.

OK
*/
type WorkerTypesGetOK struct {
	Payload *models.WorkerType
}

// IsSuccess returns true when this worker types get o k response has a 2xx status code
func (o *WorkerTypesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this worker types get o k response has a 3xx status code
func (o *WorkerTypesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this worker types get o k response has a 4xx status code
func (o *WorkerTypesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this worker types get o k response has a 5xx status code
func (o *WorkerTypesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this worker types get o k response a status code equal to that given
func (o *WorkerTypesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the worker types get o k response
func (o *WorkerTypesGetOK) Code() int {
	return 200
}

func (o *WorkerTypesGetOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/WorkerTypes/{id}][%d] workerTypesGetOK  %+v", 200, o.Payload)
}

func (o *WorkerTypesGetOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/WorkerTypes/{id}][%d] workerTypesGetOK  %+v", 200, o.Payload)
}

func (o *WorkerTypesGetOK) GetPayload() *models.WorkerType {
	return o.Payload
}

func (o *WorkerTypesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkerType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
