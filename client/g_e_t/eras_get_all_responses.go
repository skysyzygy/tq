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

// ErasGetAllReader is a Reader for the ErasGetAll structure.
type ErasGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ErasGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewErasGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/Eras] Eras_GetAll", response, response.Code())
	}
}

// NewErasGetAllOK creates a ErasGetAllOK with default headers values
func NewErasGetAllOK() *ErasGetAllOK {
	return &ErasGetAllOK{}
}

/*
ErasGetAllOK describes a response with status code 200, with default header values.

OK
*/
type ErasGetAllOK struct {
	Payload []*models.Era
}

// IsSuccess returns true when this eras get all o k response has a 2xx status code
func (o *ErasGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this eras get all o k response has a 3xx status code
func (o *ErasGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this eras get all o k response has a 4xx status code
func (o *ErasGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this eras get all o k response has a 5xx status code
func (o *ErasGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this eras get all o k response a status code equal to that given
func (o *ErasGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the eras get all o k response
func (o *ErasGetAllOK) Code() int {
	return 200
}

func (o *ErasGetAllOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/Eras][%d] erasGetAllOK  %+v", 200, o.Payload)
}

func (o *ErasGetAllOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/Eras][%d] erasGetAllOK  %+v", 200, o.Payload)
}

func (o *ErasGetAllOK) GetPayload() []*models.Era {
	return o.Payload
}

func (o *ErasGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
