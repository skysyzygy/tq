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

// DesignationCodesCreateReader is a Reader for the DesignationCodesCreate structure.
type DesignationCodesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DesignationCodesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDesignationCodesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDesignationCodesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDesignationCodesCreateOK creates a DesignationCodesCreateOK with default headers values
func NewDesignationCodesCreateOK() *DesignationCodesCreateOK {
	return &DesignationCodesCreateOK{}
}

/*
DesignationCodesCreateOK describes a response with status code 200, with default header values.

OK
*/
type DesignationCodesCreateOK struct {
	Payload *models.DesignationCode
}

// IsSuccess returns true when this designation codes create o k response has a 2xx status code
func (o *DesignationCodesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this designation codes create o k response has a 3xx status code
func (o *DesignationCodesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this designation codes create o k response has a 4xx status code
func (o *DesignationCodesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this designation codes create o k response has a 5xx status code
func (o *DesignationCodesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this designation codes create o k response a status code equal to that given
func (o *DesignationCodesCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the designation codes create o k response
func (o *DesignationCodesCreateOK) Code() int {
	return 200
}

func (o *DesignationCodesCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ReferenceData/DesignationCodes][%d] designationCodesCreateOK %s", 200, payload)
}

func (o *DesignationCodesCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ReferenceData/DesignationCodes][%d] designationCodesCreateOK %s", 200, payload)
}

func (o *DesignationCodesCreateOK) GetPayload() *models.DesignationCode {
	return o.Payload
}

func (o *DesignationCodesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DesignationCode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesignationCodesCreateDefault creates a DesignationCodesCreateDefault with default headers values
func NewDesignationCodesCreateDefault(code int) *DesignationCodesCreateDefault {
	return &DesignationCodesCreateDefault{
		_statusCode: code,
	}
}

/*
DesignationCodesCreateDefault describes a response with status code -1, with default header values.

Error
*/
type DesignationCodesCreateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this designation codes create default response has a 2xx status code
func (o *DesignationCodesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this designation codes create default response has a 3xx status code
func (o *DesignationCodesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this designation codes create default response has a 4xx status code
func (o *DesignationCodesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this designation codes create default response has a 5xx status code
func (o *DesignationCodesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this designation codes create default response a status code equal to that given
func (o *DesignationCodesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the designation codes create default response
func (o *DesignationCodesCreateDefault) Code() int {
	return o._statusCode
}

func (o *DesignationCodesCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ReferenceData/DesignationCodes][%d] DesignationCodes_Create default %s", o._statusCode, payload)
}

func (o *DesignationCodesCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ReferenceData/DesignationCodes][%d] DesignationCodes_Create default %s", o._statusCode, payload)
}

func (o *DesignationCodesCreateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *DesignationCodesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
