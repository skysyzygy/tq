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

// InterestTypesGetAllReader is a Reader for the InterestTypesGetAll structure.
type InterestTypesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterestTypesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterestTypesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewInterestTypesGetAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewInterestTypesGetAllOK creates a InterestTypesGetAllOK with default headers values
func NewInterestTypesGetAllOK() *InterestTypesGetAllOK {
	return &InterestTypesGetAllOK{}
}

/*
InterestTypesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type InterestTypesGetAllOK struct {
	Payload []*models.InterestType
}

// IsSuccess returns true when this interest types get all o k response has a 2xx status code
func (o *InterestTypesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this interest types get all o k response has a 3xx status code
func (o *InterestTypesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this interest types get all o k response has a 4xx status code
func (o *InterestTypesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this interest types get all o k response has a 5xx status code
func (o *InterestTypesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this interest types get all o k response a status code equal to that given
func (o *InterestTypesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the interest types get all o k response
func (o *InterestTypesGetAllOK) Code() int {
	return 200
}

func (o *InterestTypesGetAllOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/InterestTypes][%d] interestTypesGetAllOK %s", 200, payload)
}

func (o *InterestTypesGetAllOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/InterestTypes][%d] interestTypesGetAllOK %s", 200, payload)
}

func (o *InterestTypesGetAllOK) GetPayload() []*models.InterestType {
	return o.Payload
}

func (o *InterestTypesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterestTypesGetAllDefault creates a InterestTypesGetAllDefault with default headers values
func NewInterestTypesGetAllDefault(code int) *InterestTypesGetAllDefault {
	return &InterestTypesGetAllDefault{
		_statusCode: code,
	}
}

/*
InterestTypesGetAllDefault describes a response with status code -1, with default header values.

Error
*/
type InterestTypesGetAllDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this interest types get all default response has a 2xx status code
func (o *InterestTypesGetAllDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this interest types get all default response has a 3xx status code
func (o *InterestTypesGetAllDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this interest types get all default response has a 4xx status code
func (o *InterestTypesGetAllDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this interest types get all default response has a 5xx status code
func (o *InterestTypesGetAllDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this interest types get all default response a status code equal to that given
func (o *InterestTypesGetAllDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the interest types get all default response
func (o *InterestTypesGetAllDefault) Code() int {
	return o._statusCode
}

func (o *InterestTypesGetAllDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/InterestTypes][%d] InterestTypes_GetAll default %s", o._statusCode, payload)
}

func (o *InterestTypesGetAllDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/InterestTypes][%d] InterestTypes_GetAll default %s", o._statusCode, payload)
}

func (o *InterestTypesGetAllDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *InterestTypesGetAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
