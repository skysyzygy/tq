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

// PerformancesGetAllReader is a Reader for the PerformancesGetAll structure.
type PerformancesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformancesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformancesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPerformancesGetAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPerformancesGetAllOK creates a PerformancesGetAllOK with default headers values
func NewPerformancesGetAllOK() *PerformancesGetAllOK {
	return &PerformancesGetAllOK{}
}

/*
PerformancesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type PerformancesGetAllOK struct {
	Payload []*models.Performance
}

// IsSuccess returns true when this performances get all o k response has a 2xx status code
func (o *PerformancesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this performances get all o k response has a 3xx status code
func (o *PerformancesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this performances get all o k response has a 4xx status code
func (o *PerformancesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this performances get all o k response has a 5xx status code
func (o *PerformancesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this performances get all o k response a status code equal to that given
func (o *PerformancesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the performances get all o k response
func (o *PerformancesGetAllOK) Code() int {
	return 200
}

func (o *PerformancesGetAllOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/Performances][%d] performancesGetAllOK %s", 200, payload)
}

func (o *PerformancesGetAllOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/Performances][%d] performancesGetAllOK %s", 200, payload)
}

func (o *PerformancesGetAllOK) GetPayload() []*models.Performance {
	return o.Payload
}

func (o *PerformancesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformancesGetAllDefault creates a PerformancesGetAllDefault with default headers values
func NewPerformancesGetAllDefault(code int) *PerformancesGetAllDefault {
	return &PerformancesGetAllDefault{
		_statusCode: code,
	}
}

/*
PerformancesGetAllDefault describes a response with status code -1, with default header values.

Error
*/
type PerformancesGetAllDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this performances get all default response has a 2xx status code
func (o *PerformancesGetAllDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this performances get all default response has a 3xx status code
func (o *PerformancesGetAllDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this performances get all default response has a 4xx status code
func (o *PerformancesGetAllDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this performances get all default response has a 5xx status code
func (o *PerformancesGetAllDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this performances get all default response a status code equal to that given
func (o *PerformancesGetAllDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the performances get all default response
func (o *PerformancesGetAllDefault) Code() int {
	return o._statusCode
}

func (o *PerformancesGetAllDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/Performances][%d] Performances_GetAll default %s", o._statusCode, payload)
}

func (o *PerformancesGetAllDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/Performances][%d] Performances_GetAll default %s", o._statusCode, payload)
}

func (o *PerformancesGetAllDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PerformancesGetAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
