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

// WorkerRolesGetAllReader is a Reader for the WorkerRolesGetAll structure.
type WorkerRolesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkerRolesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkerRolesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/WorkerRoles] WorkerRoles_GetAll", response, response.Code())
	}
}

// NewWorkerRolesGetAllOK creates a WorkerRolesGetAllOK with default headers values
func NewWorkerRolesGetAllOK() *WorkerRolesGetAllOK {
	return &WorkerRolesGetAllOK{}
}

/*
WorkerRolesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type WorkerRolesGetAllOK struct {
	Payload []*models.WorkerRole
}

// IsSuccess returns true when this worker roles get all o k response has a 2xx status code
func (o *WorkerRolesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this worker roles get all o k response has a 3xx status code
func (o *WorkerRolesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this worker roles get all o k response has a 4xx status code
func (o *WorkerRolesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this worker roles get all o k response has a 5xx status code
func (o *WorkerRolesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this worker roles get all o k response a status code equal to that given
func (o *WorkerRolesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the worker roles get all o k response
func (o *WorkerRolesGetAllOK) Code() int {
	return 200
}

func (o *WorkerRolesGetAllOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/WorkerRoles][%d] workerRolesGetAllOK  %+v", 200, o.Payload)
}

func (o *WorkerRolesGetAllOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/WorkerRoles][%d] workerRolesGetAllOK  %+v", 200, o.Payload)
}

func (o *WorkerRolesGetAllOK) GetPayload() []*models.WorkerRole {
	return o.Payload
}

func (o *WorkerRolesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
