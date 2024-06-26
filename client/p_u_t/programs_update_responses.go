// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// ProgramsUpdateReader is a Reader for the ProgramsUpdate structure.
type ProgramsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProgramsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProgramsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProgramsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProgramsUpdateOK creates a ProgramsUpdateOK with default headers values
func NewProgramsUpdateOK() *ProgramsUpdateOK {
	return &ProgramsUpdateOK{}
}

/*
ProgramsUpdateOK describes a response with status code 200, with default header values.

OK
*/
type ProgramsUpdateOK struct {
	Payload *models.Program
}

// IsSuccess returns true when this programs update o k response has a 2xx status code
func (o *ProgramsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this programs update o k response has a 3xx status code
func (o *ProgramsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this programs update o k response has a 4xx status code
func (o *ProgramsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this programs update o k response has a 5xx status code
func (o *ProgramsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this programs update o k response a status code equal to that given
func (o *ProgramsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the programs update o k response
func (o *ProgramsUpdateOK) Code() int {
	return 200
}

func (o *ProgramsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/Programs/{id}][%d] programsUpdateOK %s", 200, payload)
}

func (o *ProgramsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/Programs/{id}][%d] programsUpdateOK %s", 200, payload)
}

func (o *ProgramsUpdateOK) GetPayload() *models.Program {
	return o.Payload
}

func (o *ProgramsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Program)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProgramsUpdateDefault creates a ProgramsUpdateDefault with default headers values
func NewProgramsUpdateDefault(code int) *ProgramsUpdateDefault {
	return &ProgramsUpdateDefault{
		_statusCode: code,
	}
}

/*
ProgramsUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type ProgramsUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this programs update default response has a 2xx status code
func (o *ProgramsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this programs update default response has a 3xx status code
func (o *ProgramsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this programs update default response has a 4xx status code
func (o *ProgramsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this programs update default response has a 5xx status code
func (o *ProgramsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this programs update default response a status code equal to that given
func (o *ProgramsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the programs update default response
func (o *ProgramsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ProgramsUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/Programs/{id}][%d] Programs_Update default %s", o._statusCode, payload)
}

func (o *ProgramsUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/Programs/{id}][%d] Programs_Update default %s", o._statusCode, payload)
}

func (o *ProgramsUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ProgramsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
