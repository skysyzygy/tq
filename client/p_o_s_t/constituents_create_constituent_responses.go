// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

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

// ConstituentsCreateConstituentReader is a Reader for the ConstituentsCreateConstituent structure.
type ConstituentsCreateConstituentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConstituentsCreateConstituentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConstituentsCreateConstituentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConstituentsCreateConstituentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConstituentsCreateConstituentOK creates a ConstituentsCreateConstituentOK with default headers values
func NewConstituentsCreateConstituentOK() *ConstituentsCreateConstituentOK {
	return &ConstituentsCreateConstituentOK{}
}

/*
ConstituentsCreateConstituentOK describes a response with status code 200, with default header values.

OK
*/
type ConstituentsCreateConstituentOK struct {
	Payload *models.ConstituentDetail
}

// IsSuccess returns true when this constituents create constituent o k response has a 2xx status code
func (o *ConstituentsCreateConstituentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this constituents create constituent o k response has a 3xx status code
func (o *ConstituentsCreateConstituentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this constituents create constituent o k response has a 4xx status code
func (o *ConstituentsCreateConstituentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this constituents create constituent o k response has a 5xx status code
func (o *ConstituentsCreateConstituentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this constituents create constituent o k response a status code equal to that given
func (o *ConstituentsCreateConstituentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the constituents create constituent o k response
func (o *ConstituentsCreateConstituentOK) Code() int {
	return 200
}

func (o *ConstituentsCreateConstituentOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /CRM/Constituents/Detail][%d] constituentsCreateConstituentOK %s", 200, payload)
}

func (o *ConstituentsCreateConstituentOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /CRM/Constituents/Detail][%d] constituentsCreateConstituentOK %s", 200, payload)
}

func (o *ConstituentsCreateConstituentOK) GetPayload() *models.ConstituentDetail {
	return o.Payload
}

func (o *ConstituentsCreateConstituentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConstituentDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConstituentsCreateConstituentDefault creates a ConstituentsCreateConstituentDefault with default headers values
func NewConstituentsCreateConstituentDefault(code int) *ConstituentsCreateConstituentDefault {
	return &ConstituentsCreateConstituentDefault{
		_statusCode: code,
	}
}

/*
ConstituentsCreateConstituentDefault describes a response with status code -1, with default header values.

Error
*/
type ConstituentsCreateConstituentDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this constituents create constituent default response has a 2xx status code
func (o *ConstituentsCreateConstituentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this constituents create constituent default response has a 3xx status code
func (o *ConstituentsCreateConstituentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this constituents create constituent default response has a 4xx status code
func (o *ConstituentsCreateConstituentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this constituents create constituent default response has a 5xx status code
func (o *ConstituentsCreateConstituentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this constituents create constituent default response a status code equal to that given
func (o *ConstituentsCreateConstituentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the constituents create constituent default response
func (o *ConstituentsCreateConstituentDefault) Code() int {
	return o._statusCode
}

func (o *ConstituentsCreateConstituentDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /CRM/Constituents/Detail][%d] Constituents_CreateConstituent default %s", o._statusCode, payload)
}

func (o *ConstituentsCreateConstituentDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /CRM/Constituents/Detail][%d] Constituents_CreateConstituent default %s", o._statusCode, payload)
}

func (o *ConstituentsCreateConstituentDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ConstituentsCreateConstituentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
