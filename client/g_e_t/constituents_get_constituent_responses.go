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

// ConstituentsGetConstituentReader is a Reader for the ConstituentsGetConstituent structure.
type ConstituentsGetConstituentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConstituentsGetConstituentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConstituentsGetConstituentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConstituentsGetConstituentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConstituentsGetConstituentOK creates a ConstituentsGetConstituentOK with default headers values
func NewConstituentsGetConstituentOK() *ConstituentsGetConstituentOK {
	return &ConstituentsGetConstituentOK{}
}

/*
ConstituentsGetConstituentOK describes a response with status code 200, with default header values.

OK
*/
type ConstituentsGetConstituentOK struct {
	Payload *models.ConstituentDetail
}

// IsSuccess returns true when this constituents get constituent o k response has a 2xx status code
func (o *ConstituentsGetConstituentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this constituents get constituent o k response has a 3xx status code
func (o *ConstituentsGetConstituentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this constituents get constituent o k response has a 4xx status code
func (o *ConstituentsGetConstituentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this constituents get constituent o k response has a 5xx status code
func (o *ConstituentsGetConstituentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this constituents get constituent o k response a status code equal to that given
func (o *ConstituentsGetConstituentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the constituents get constituent o k response
func (o *ConstituentsGetConstituentOK) Code() int {
	return 200
}

func (o *ConstituentsGetConstituentOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /CRM/Constituents/{constituentId}/Detail][%d] constituentsGetConstituentOK %s", 200, payload)
}

func (o *ConstituentsGetConstituentOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /CRM/Constituents/{constituentId}/Detail][%d] constituentsGetConstituentOK %s", 200, payload)
}

func (o *ConstituentsGetConstituentOK) GetPayload() *models.ConstituentDetail {
	return o.Payload
}

func (o *ConstituentsGetConstituentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConstituentDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConstituentsGetConstituentDefault creates a ConstituentsGetConstituentDefault with default headers values
func NewConstituentsGetConstituentDefault(code int) *ConstituentsGetConstituentDefault {
	return &ConstituentsGetConstituentDefault{
		_statusCode: code,
	}
}

/*
ConstituentsGetConstituentDefault describes a response with status code -1, with default header values.

Error
*/
type ConstituentsGetConstituentDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this constituents get constituent default response has a 2xx status code
func (o *ConstituentsGetConstituentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this constituents get constituent default response has a 3xx status code
func (o *ConstituentsGetConstituentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this constituents get constituent default response has a 4xx status code
func (o *ConstituentsGetConstituentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this constituents get constituent default response has a 5xx status code
func (o *ConstituentsGetConstituentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this constituents get constituent default response a status code equal to that given
func (o *ConstituentsGetConstituentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the constituents get constituent default response
func (o *ConstituentsGetConstituentDefault) Code() int {
	return o._statusCode
}

func (o *ConstituentsGetConstituentDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /CRM/Constituents/{constituentId}/Detail][%d] Constituents_GetConstituent default %s", o._statusCode, payload)
}

func (o *ConstituentsGetConstituentDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /CRM/Constituents/{constituentId}/Detail][%d] Constituents_GetConstituent default %s", o._statusCode, payload)
}

func (o *ConstituentsGetConstituentDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ConstituentsGetConstituentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
