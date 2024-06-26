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

// ConstituentsGetPrimaryHouseholdReader is a Reader for the ConstituentsGetPrimaryHousehold structure.
type ConstituentsGetPrimaryHouseholdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConstituentsGetPrimaryHouseholdReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConstituentsGetPrimaryHouseholdOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConstituentsGetPrimaryHouseholdDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConstituentsGetPrimaryHouseholdOK creates a ConstituentsGetPrimaryHouseholdOK with default headers values
func NewConstituentsGetPrimaryHouseholdOK() *ConstituentsGetPrimaryHouseholdOK {
	return &ConstituentsGetPrimaryHouseholdOK{}
}

/*
ConstituentsGetPrimaryHouseholdOK describes a response with status code 200, with default header values.

OK
*/
type ConstituentsGetPrimaryHouseholdOK struct {
	Payload *models.ConstituentDetail
}

// IsSuccess returns true when this constituents get primary household o k response has a 2xx status code
func (o *ConstituentsGetPrimaryHouseholdOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this constituents get primary household o k response has a 3xx status code
func (o *ConstituentsGetPrimaryHouseholdOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this constituents get primary household o k response has a 4xx status code
func (o *ConstituentsGetPrimaryHouseholdOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this constituents get primary household o k response has a 5xx status code
func (o *ConstituentsGetPrimaryHouseholdOK) IsServerError() bool {
	return false
}

// IsCode returns true when this constituents get primary household o k response a status code equal to that given
func (o *ConstituentsGetPrimaryHouseholdOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the constituents get primary household o k response
func (o *ConstituentsGetPrimaryHouseholdOK) Code() int {
	return 200
}

func (o *ConstituentsGetPrimaryHouseholdOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /CRM/Constituents/PrimaryHousehold][%d] constituentsGetPrimaryHouseholdOK %s", 200, payload)
}

func (o *ConstituentsGetPrimaryHouseholdOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /CRM/Constituents/PrimaryHousehold][%d] constituentsGetPrimaryHouseholdOK %s", 200, payload)
}

func (o *ConstituentsGetPrimaryHouseholdOK) GetPayload() *models.ConstituentDetail {
	return o.Payload
}

func (o *ConstituentsGetPrimaryHouseholdOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConstituentDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConstituentsGetPrimaryHouseholdDefault creates a ConstituentsGetPrimaryHouseholdDefault with default headers values
func NewConstituentsGetPrimaryHouseholdDefault(code int) *ConstituentsGetPrimaryHouseholdDefault {
	return &ConstituentsGetPrimaryHouseholdDefault{
		_statusCode: code,
	}
}

/*
ConstituentsGetPrimaryHouseholdDefault describes a response with status code -1, with default header values.

Error
*/
type ConstituentsGetPrimaryHouseholdDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this constituents get primary household default response has a 2xx status code
func (o *ConstituentsGetPrimaryHouseholdDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this constituents get primary household default response has a 3xx status code
func (o *ConstituentsGetPrimaryHouseholdDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this constituents get primary household default response has a 4xx status code
func (o *ConstituentsGetPrimaryHouseholdDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this constituents get primary household default response has a 5xx status code
func (o *ConstituentsGetPrimaryHouseholdDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this constituents get primary household default response a status code equal to that given
func (o *ConstituentsGetPrimaryHouseholdDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the constituents get primary household default response
func (o *ConstituentsGetPrimaryHouseholdDefault) Code() int {
	return o._statusCode
}

func (o *ConstituentsGetPrimaryHouseholdDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /CRM/Constituents/PrimaryHousehold][%d] Constituents_GetPrimaryHousehold default %s", o._statusCode, payload)
}

func (o *ConstituentsGetPrimaryHouseholdDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /CRM/Constituents/PrimaryHousehold][%d] Constituents_GetPrimaryHousehold default %s", o._statusCode, payload)
}

func (o *ConstituentsGetPrimaryHouseholdDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ConstituentsGetPrimaryHouseholdDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
