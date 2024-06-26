// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

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

// AppealsGetReader is a Reader for the AppealsGet structure.
type AppealsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppealsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAppealsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAppealsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAppealsGetOK creates a AppealsGetOK with default headers values
func NewAppealsGetOK() *AppealsGetOK {
	return &AppealsGetOK{}
}

/*
AppealsGetOK describes a response with status code 200, with default header values.

OK
*/
type AppealsGetOK struct {
	Payload *models.Appeal
}

// IsSuccess returns true when this appeals get o k response has a 2xx status code
func (o *AppealsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this appeals get o k response has a 3xx status code
func (o *AppealsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this appeals get o k response has a 4xx status code
func (o *AppealsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this appeals get o k response has a 5xx status code
func (o *AppealsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this appeals get o k response a status code equal to that given
func (o *AppealsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the appeals get o k response
func (o *AppealsGetOK) Code() int {
	return 200
}

func (o *AppealsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Finance/Appeals/{appealId}][%d] appealsGetOK %s", 200, payload)
}

func (o *AppealsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Finance/Appeals/{appealId}][%d] appealsGetOK %s", 200, payload)
}

func (o *AppealsGetOK) GetPayload() *models.Appeal {
	return o.Payload
}

func (o *AppealsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Appeal)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppealsGetDefault creates a AppealsGetDefault with default headers values
func NewAppealsGetDefault(code int) *AppealsGetDefault {
	return &AppealsGetDefault{
		_statusCode: code,
	}
}

/*
AppealsGetDefault describes a response with status code -1, with default header values.

Error
*/
type AppealsGetDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this appeals get default response has a 2xx status code
func (o *AppealsGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this appeals get default response has a 3xx status code
func (o *AppealsGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this appeals get default response has a 4xx status code
func (o *AppealsGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this appeals get default response has a 5xx status code
func (o *AppealsGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this appeals get default response a status code equal to that given
func (o *AppealsGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the appeals get default response
func (o *AppealsGetDefault) Code() int {
	return o._statusCode
}

func (o *AppealsGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Finance/Appeals/{appealId}][%d] Appeals_Get default %s", o._statusCode, payload)
}

func (o *AppealsGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Finance/Appeals/{appealId}][%d] Appeals_Get default %s", o._statusCode, payload)
}

func (o *AppealsGetDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *AppealsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
