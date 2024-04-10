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

// WorkersGetAllReader is a Reader for the WorkersGetAll structure.
type WorkersGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkersGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkersGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /Finance/Workers] Workers_GetAll", response, response.Code())
	}
}

// NewWorkersGetAllOK creates a WorkersGetAllOK with default headers values
func NewWorkersGetAllOK() *WorkersGetAllOK {
	return &WorkersGetAllOK{}
}

/*
WorkersGetAllOK describes a response with status code 200, with default header values.

OK
*/
type WorkersGetAllOK struct {
	Payload []*models.Worker
}

// IsSuccess returns true when this workers get all o k response has a 2xx status code
func (o *WorkersGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this workers get all o k response has a 3xx status code
func (o *WorkersGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this workers get all o k response has a 4xx status code
func (o *WorkersGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this workers get all o k response has a 5xx status code
func (o *WorkersGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this workers get all o k response a status code equal to that given
func (o *WorkersGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the workers get all o k response
func (o *WorkersGetAllOK) Code() int {
	return 200
}

func (o *WorkersGetAllOK) Error() string {
	return fmt.Sprintf("[GET /Finance/Workers][%d] workersGetAllOK  %+v", 200, o.Payload)
}

func (o *WorkersGetAllOK) String() string {
	return fmt.Sprintf("[GET /Finance/Workers][%d] workersGetAllOK  %+v", 200, o.Payload)
}

func (o *WorkersGetAllOK) GetPayload() []*models.Worker {
	return o.Payload
}

func (o *WorkersGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
