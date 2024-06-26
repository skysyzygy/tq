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

// PerformancePriceTypesUpdateReader is a Reader for the PerformancePriceTypesUpdate structure.
type PerformancePriceTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformancePriceTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformancePriceTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPerformancePriceTypesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPerformancePriceTypesUpdateOK creates a PerformancePriceTypesUpdateOK with default headers values
func NewPerformancePriceTypesUpdateOK() *PerformancePriceTypesUpdateOK {
	return &PerformancePriceTypesUpdateOK{}
}

/*
PerformancePriceTypesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type PerformancePriceTypesUpdateOK struct {
	Payload *models.PerformancePriceType
}

// IsSuccess returns true when this performance price types update o k response has a 2xx status code
func (o *PerformancePriceTypesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this performance price types update o k response has a 3xx status code
func (o *PerformancePriceTypesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this performance price types update o k response has a 4xx status code
func (o *PerformancePriceTypesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this performance price types update o k response has a 5xx status code
func (o *PerformancePriceTypesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this performance price types update o k response a status code equal to that given
func (o *PerformancePriceTypesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the performance price types update o k response
func (o *PerformancePriceTypesUpdateOK) Code() int {
	return 200
}

func (o *PerformancePriceTypesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /TXN/PerformancePriceTypes/{performancePriceTypeId}][%d] performancePriceTypesUpdateOK %s", 200, payload)
}

func (o *PerformancePriceTypesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /TXN/PerformancePriceTypes/{performancePriceTypeId}][%d] performancePriceTypesUpdateOK %s", 200, payload)
}

func (o *PerformancePriceTypesUpdateOK) GetPayload() *models.PerformancePriceType {
	return o.Payload
}

func (o *PerformancePriceTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PerformancePriceType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformancePriceTypesUpdateDefault creates a PerformancePriceTypesUpdateDefault with default headers values
func NewPerformancePriceTypesUpdateDefault(code int) *PerformancePriceTypesUpdateDefault {
	return &PerformancePriceTypesUpdateDefault{
		_statusCode: code,
	}
}

/*
PerformancePriceTypesUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type PerformancePriceTypesUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this performance price types update default response has a 2xx status code
func (o *PerformancePriceTypesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this performance price types update default response has a 3xx status code
func (o *PerformancePriceTypesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this performance price types update default response has a 4xx status code
func (o *PerformancePriceTypesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this performance price types update default response has a 5xx status code
func (o *PerformancePriceTypesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this performance price types update default response a status code equal to that given
func (o *PerformancePriceTypesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the performance price types update default response
func (o *PerformancePriceTypesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *PerformancePriceTypesUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /TXN/PerformancePriceTypes/{performancePriceTypeId}][%d] PerformancePriceTypes_Update default %s", o._statusCode, payload)
}

func (o *PerformancePriceTypesUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /TXN/PerformancePriceTypes/{performancePriceTypeId}][%d] PerformancePriceTypes_Update default %s", o._statusCode, payload)
}

func (o *PerformancePriceTypesUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PerformancePriceTypesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
