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

// CrediteeTypesGetReader is a Reader for the CrediteeTypesGet structure.
type CrediteeTypesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CrediteeTypesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCrediteeTypesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCrediteeTypesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCrediteeTypesGetOK creates a CrediteeTypesGetOK with default headers values
func NewCrediteeTypesGetOK() *CrediteeTypesGetOK {
	return &CrediteeTypesGetOK{}
}

/*
CrediteeTypesGetOK describes a response with status code 200, with default header values.

OK
*/
type CrediteeTypesGetOK struct {
	Payload *models.CrediteeType
}

// IsSuccess returns true when this creditee types get o k response has a 2xx status code
func (o *CrediteeTypesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this creditee types get o k response has a 3xx status code
func (o *CrediteeTypesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this creditee types get o k response has a 4xx status code
func (o *CrediteeTypesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this creditee types get o k response has a 5xx status code
func (o *CrediteeTypesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this creditee types get o k response a status code equal to that given
func (o *CrediteeTypesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the creditee types get o k response
func (o *CrediteeTypesGetOK) Code() int {
	return 200
}

func (o *CrediteeTypesGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/CrediteeTypes/{id}][%d] crediteeTypesGetOK %s", 200, payload)
}

func (o *CrediteeTypesGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/CrediteeTypes/{id}][%d] crediteeTypesGetOK %s", 200, payload)
}

func (o *CrediteeTypesGetOK) GetPayload() *models.CrediteeType {
	return o.Payload
}

func (o *CrediteeTypesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CrediteeType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCrediteeTypesGetDefault creates a CrediteeTypesGetDefault with default headers values
func NewCrediteeTypesGetDefault(code int) *CrediteeTypesGetDefault {
	return &CrediteeTypesGetDefault{
		_statusCode: code,
	}
}

/*
CrediteeTypesGetDefault describes a response with status code -1, with default header values.

Error
*/
type CrediteeTypesGetDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this creditee types get default response has a 2xx status code
func (o *CrediteeTypesGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this creditee types get default response has a 3xx status code
func (o *CrediteeTypesGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this creditee types get default response has a 4xx status code
func (o *CrediteeTypesGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this creditee types get default response has a 5xx status code
func (o *CrediteeTypesGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this creditee types get default response a status code equal to that given
func (o *CrediteeTypesGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the creditee types get default response
func (o *CrediteeTypesGetDefault) Code() int {
	return o._statusCode
}

func (o *CrediteeTypesGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/CrediteeTypes/{id}][%d] CrediteeTypes_Get default %s", o._statusCode, payload)
}

func (o *CrediteeTypesGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/CrediteeTypes/{id}][%d] CrediteeTypes_Get default %s", o._statusCode, payload)
}

func (o *CrediteeTypesGetDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CrediteeTypesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
