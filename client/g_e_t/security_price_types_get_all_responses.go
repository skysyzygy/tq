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

// SecurityPriceTypesGetAllReader is a Reader for the SecurityPriceTypesGetAll structure.
type SecurityPriceTypesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityPriceTypesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityPriceTypesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityPriceTypesGetAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityPriceTypesGetAllOK creates a SecurityPriceTypesGetAllOK with default headers values
func NewSecurityPriceTypesGetAllOK() *SecurityPriceTypesGetAllOK {
	return &SecurityPriceTypesGetAllOK{}
}

/*
SecurityPriceTypesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type SecurityPriceTypesGetAllOK struct {
	Payload []*models.PriceTypeUserGroup
}

// IsSuccess returns true when this security price types get all o k response has a 2xx status code
func (o *SecurityPriceTypesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security price types get all o k response has a 3xx status code
func (o *SecurityPriceTypesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security price types get all o k response has a 4xx status code
func (o *SecurityPriceTypesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security price types get all o k response has a 5xx status code
func (o *SecurityPriceTypesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security price types get all o k response a status code equal to that given
func (o *SecurityPriceTypesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security price types get all o k response
func (o *SecurityPriceTypesGetAllOK) Code() int {
	return 200
}

func (o *SecurityPriceTypesGetAllOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Security/PriceTypes][%d] securityPriceTypesGetAllOK %s", 200, payload)
}

func (o *SecurityPriceTypesGetAllOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Security/PriceTypes][%d] securityPriceTypesGetAllOK %s", 200, payload)
}

func (o *SecurityPriceTypesGetAllOK) GetPayload() []*models.PriceTypeUserGroup {
	return o.Payload
}

func (o *SecurityPriceTypesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityPriceTypesGetAllDefault creates a SecurityPriceTypesGetAllDefault with default headers values
func NewSecurityPriceTypesGetAllDefault(code int) *SecurityPriceTypesGetAllDefault {
	return &SecurityPriceTypesGetAllDefault{
		_statusCode: code,
	}
}

/*
SecurityPriceTypesGetAllDefault describes a response with status code -1, with default header values.

Error
*/
type SecurityPriceTypesGetAllDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this security price types get all default response has a 2xx status code
func (o *SecurityPriceTypesGetAllDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security price types get all default response has a 3xx status code
func (o *SecurityPriceTypesGetAllDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security price types get all default response has a 4xx status code
func (o *SecurityPriceTypesGetAllDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security price types get all default response has a 5xx status code
func (o *SecurityPriceTypesGetAllDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security price types get all default response a status code equal to that given
func (o *SecurityPriceTypesGetAllDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security price types get all default response
func (o *SecurityPriceTypesGetAllDefault) Code() int {
	return o._statusCode
}

func (o *SecurityPriceTypesGetAllDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Security/PriceTypes][%d] SecurityPriceTypes_GetAll default %s", o._statusCode, payload)
}

func (o *SecurityPriceTypesGetAllDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Security/PriceTypes][%d] SecurityPriceTypes_GetAll default %s", o._statusCode, payload)
}

func (o *SecurityPriceTypesGetAllDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SecurityPriceTypesGetAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
